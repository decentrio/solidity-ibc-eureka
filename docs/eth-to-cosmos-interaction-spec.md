# Ethereum to Cosmos Interaction Spec

This document describes the Ethereum -> Cosmos IBC Eureka flow for ICS-20 token transfer packets: send, receive on Cosmos, write acknowledgement, acknowledge on Ethereum, and timeout/error handling. It is written as a handoff map for engineers or AI agents that need to continue work in this repo.

## Scope

The primary path is ERC-20 on Ethereum sent through `ICS20Transfer` and `ICS26Router`, relayed by the `eth-to-cosmos` module into Cosmos SDK IBC v2 messages. The reverse Cosmos -> Ethereum leg is included only where it matters for acknowledgements and round trips.

Code entry points:

- Ethereum app contract: [`contracts/ICS20Transfer.sol`](../contracts/ICS20Transfer.sol)
- Ethereum router contract: [`contracts/ICS26Router.sol`](../contracts/ICS26Router.sol)
- Ethereum IBC commitment store: [`contracts/utils/IBCStoreUpgradeable.sol`](../contracts/utils/IBCStoreUpgradeable.sol)
- Ethereum router message types: [`contracts/msgs/IICS26RouterMsgs.sol`](../contracts/msgs/IICS26RouterMsgs.sol)
- ICS-20 app message types: [`contracts/msgs/IICS20TransferMsgs.sol`](../contracts/msgs/IICS20TransferMsgs.sol)
- ETH -> Cosmos relayer module: [`packages/relayer/modules/eth-to-cosmos/src/lib.rs`](../packages/relayer/modules/eth-to-cosmos/src/lib.rs)
- ETH -> Cosmos transaction builder: [`packages/relayer/modules/eth-to-cosmos/src/tx_builder.rs`](../packages/relayer/modules/eth-to-cosmos/src/tx_builder.rs)
- Relayer event/message conversion helpers: [`packages/relayer/lib/src/utils/cosmos.rs`](../packages/relayer/lib/src/utils/cosmos.rs)
- E2E happy path: [`e2e/interchaintestv8/ibc_eureka_test.go`](../e2e/interchaintestv8/ibc_eureka_test.go)
- IFT/native attestor E2E variants: [`e2e/interchaintestv8/cosmos_ethereum_ift_test.go`](../e2e/interchaintestv8/cosmos_ethereum_ift_test.go)

## Actors

- **Ethereum user**: approves and calls `ICS20Transfer.sendTransfer`.
- **ICS20Transfer**: IBC application on Ethereum. Escrows native ERC-20 or burns returning voucher tokens, builds ICS-20 packet data, and handles ack/timeout refunds.
- **ICS26Router**: Ethereum IBC router. Allocates sequence, commits packet commitments, verifies counterparty proofs, calls apps, writes acknowledgements, and clears commitments.
- **Relayer service**: watches events, builds unsigned transaction bodies for the target chain, and returns them through the gRPC API.
- **Cosmos chain**: receives IBC v2 `MsgRecvPacket`, executes transfer module logic, writes ack, and later accepts `MsgAcknowledgement` or `MsgTimeout`.
- **Light client on Cosmos**: verifies Ethereum commitment proofs. It can be the Ethereum wasm light client or native attestor light client depending on relayer mode.

## Data Types

### Ethereum ICS-20 Send Input

`ICS20Transfer.sendTransfer` accepts `SendTransferMsg` from [`IICS20TransferMsgs.sol`](../contracts/msgs/IICS20TransferMsgs.sol#L17):

| Field | Type | Meaning |
| --- | --- | --- |
| `denom` | `address` | ERC-20 token address on Ethereum. |
| `amount` | `uint256` | Token amount. Must be greater than zero. |
| `receiver` | `string` | Cosmos receiver address. |
| `sourceClient` | `string` | Ethereum-side client ID representing the Cosmos counterparty. |
| `destPort` | `string` | Destination app port. For ICS-20 this should be `transfer`. |
| `timeoutTimestamp` | `uint64` | Absolute Unix timestamp in seconds. Router requires it to be in the future and no more than one day ahead. |
| `memo` | `string` | Optional ICS-20 memo. |

The contract currently builds the router payload with the default ICS-20 port for both source and destination; see [`ICS20Transfer.sol#L256`](../contracts/ICS20Transfer.sol#L256) and [`ICS20Transfer.sol#L264`](../contracts/ICS20Transfer.sol#L264).

### ICS-20 Packet Data

The packet payload is `abi.encode(FungibleTokenPacketData)` from [`IICS20TransferMsgs.sol#L35`](../contracts/msgs/IICS20TransferMsgs.sol#L35):

| Field | Type | Meaning |
| --- | --- | --- |
| `denom` | `string` | Full denom path. Native Ethereum ERC-20 starts as the token address hex string. Returning vouchers may carry an IBC denom path. |
| `sender` | `string` | Ethereum sender encoded as hex string. Used for ack callbacks and refunds. |
| `receiver` | `string` | Cosmos receiver address. |
| `amount` | `uint256` | Transfer amount. |
| `memo` | `string` | Optional memo. |

The packet data is built in [`ICS20Transfer._sendTransferFromEscrowWithSender`](../contracts/ICS20Transfer.sol#L232), specifically [`packetData`](../contracts/ICS20Transfer.sol#L256).

### Router Packet

The router packet and payload structs are defined in [`IICS26RouterMsgs.sol#L17`](../contracts/msgs/IICS26RouterMsgs.sol#L17):

```solidity
struct Packet {
    uint64 sequence;
    string sourceClient;
    string destClient;
    uint64 timeoutTimestamp;
    Payload[] payloads;
}

struct Payload {
    string sourcePort;
    string destPort;
    string version;
    string encoding;
    bytes value;
}
```

Important payload values for ICS-20:

- `sourcePort`: `transfer`
- `destPort`: `transfer`
- `version`: ICS-20 v1
- `encoding`: ABI encoding
- `value`: ABI encoded `FungibleTokenPacketData`

The Solidity router currently only supports one payload per packet in `recvPacket`, `ackPacket`, and `timeoutPacket`; see [`ICS26Router.recvPacket`](../contracts/ICS26Router.sol#L144), [`ackPacket`](../contracts/ICS26Router.sol#L211), and [`timeoutPacket`](../contracts/ICS26Router.sol#L263).

### Relayer gRPC Request

The public relayer API is [`proto/relayer/relayer.proto`](../proto/relayer/relayer.proto). `RelayByTxRequest` contains:

- `src_chain`, `dst_chain`
- `source_tx_ids`: source tx hashes to scan for `SendPacket` or `WriteAcknowledgement`
- `timeout_tx_ids`: target tx hashes to scan for expired sends when building timeout messages
- `src_client_id`, `dst_client_id`: filters for packet direction
- `src_packet_sequences`, `dst_packet_sequences`: optional sequence filters

For ETH -> Cosmos receive, `source_tx_ids` are Ethereum transaction hashes and the response `tx` is a Cosmos `TxBody`; `address` is empty. See the relayer response behavior in [`eth-to-cosmos/src/lib.rs#L160`](../packages/relayer/modules/eth-to-cosmos/src/lib.rs#L160).

## Flow 1: Ethereum Send -> Cosmos Receive -> Ethereum Ack

### 1. User approves and sends from Ethereum

The e2e test starts by approving `ICS20Transfer` to spend the ERC-20 and then ABI-packing `sendTransfer`; see [`ibc_eureka_test.go#L675`](../e2e/interchaintestv8/ibc_eureka_test.go#L675).

Contract flow:

1. User calls [`ICS20Transfer.sendTransfer`](../contracts/ICS20Transfer.sol#L162).
2. `amount > 0` is enforced.
3. `_getOrCreateEscrow(sourceClient)` creates or fetches the escrow for that client.
4. `_transferFrom(user, escrow, denom, amount)` moves ERC-20 into escrow.
5. `escrow.recvCallback` records receipt.
6. [`_sendTransferFromEscrowWithSender`](../contracts/ICS20Transfer.sol#L232) determines the full denom path:
   - Native ERC-20: token address hex string.
   - IBC voucher returning to source: detect prefix and burn from escrow.
7. It builds `FungibleTokenPacketData` at [`ICS20Transfer.sol#L256`](../contracts/ICS20Transfer.sol#L256).
8. It calls `ICS26Router.sendPacket` with one `Payload` at [`ICS20Transfer.sol#L264`](../contracts/ICS20Transfer.sol#L264).

Router flow:

1. [`ICS26Router.sendPacket`](../contracts/ICS26Router.sol#L108) checks the caller is the registered IBC app for `payload.sourcePort`.
2. It resolves `destClient` from the counterparty of `sourceClient`.
3. It validates timeout timestamp.
4. It increments sequence via `nextSequenceSend`.
5. It constructs `Packet`.
6. It stores the commitment through [`commitPacket`](../contracts/utils/IBCStoreUpgradeable.sol#L51).
7. It emits `SendPacket`.

Commitment key/value:

- Path: packet commitment path for `(sourceClient, sequence)`.
- Value: `ICS24Host.packetCommitmentBytes32(packet)`.
- Storage location: [`IBCStoreUpgradeable.commitPacket`](../contracts/utils/IBCStoreUpgradeable.sol#L51).

### 2. Relayer builds Cosmos receive transaction

The e2e test calls `RelayByTx` with the Ethereum send tx hash at [`ibc_eureka_test.go#L736`](../e2e/interchaintestv8/ibc_eureka_test.go#L736).

Relayer flow:

1. [`EthToCosmosRelayerModuleService.relay_by_tx`](../packages/relayer/modules/eth-to-cosmos/src/lib.rs#L160) parses Ethereum tx hashes.
2. [`eth_listener.fetch_tx_events`](../packages/relayer/modules/eth-to-cosmos/src/lib.rs#L176) loads logs from the Ethereum tx block.
3. The Ethereum listener decodes only `SendPacket` and `WriteAcknowledgement` events in [`listener/eth_eureka.rs#L63`](../packages/relayer/lib/src/listener/eth_eureka.rs#L63) and [`events/eureka.rs`](../packages/relayer/lib/src/events/eureka.rs).
4. [`tx_builder.relay_events`](../packages/relayer/modules/eth-to-cosmos/src/lib.rs#L210) receives normalized `EurekaEventWithHeight`.
5. [`cosmos::src_events_to_recv_and_ack_msgs`](../packages/relayer/lib/src/utils/cosmos.rs#L91) converts matching `SendPacket` events into Cosmos `MsgRecvPacket`.
6. Proofs are injected:
   - Real Ethereum light client path: [`inject_ethereum_proofs`](../packages/relayer/lib/src/utils/cosmos.rs#L607).
   - Attestor path: [`build_attestor_relay_events_tx`](../packages/relayer/lib/src/utils/cosmos_attested.rs#L311), which fetches attestations, builds an update-client message, and injects attestor proofs.
7. The relayer returns a protobuf-encoded Cosmos `TxBody`.

Filtering rules in [`src_events_to_recv_and_ack_msgs`](../packages/relayer/lib/src/utils/cosmos.rs#L91):

- A `SendPacket` becomes `MsgRecvPacket` only if:
  - `packet.timeoutTimestamp > now`
  - `packet.sourceClient == src_client_id`
  - `packet.destClient == dst_client_id`
  - sequence is included in `src_packet_sequences`, or no sequence filter is provided
- A `WriteAcknowledgement` becomes `MsgAcknowledgement` only if:
  - `packet.sourceClient == dst_client_id`
  - `packet.destClient == src_client_id`
  - sequence is included in `dst_packet_sequences`, or no sequence filter is provided

### 3. Submit Cosmos receive transaction

The caller signs and broadcasts the returned Cosmos `TxBody`; see [`ibc_eureka_test.go#L736`](../e2e/interchaintestv8/ibc_eureka_test.go#L736).

On Cosmos, `MsgRecvPacket` should:

1. Verify the Ethereum packet commitment proof against the Ethereum light client.
2. Write the packet receipt.
3. Execute the transfer app with the ABI encoded packet payload.
4. Mint or unlock the corresponding Cosmos voucher denom.
5. Write an acknowledgement.

The e2e test checks the Cosmos user receives the voucher denom at [`ibc_eureka_test.go#L736`](../e2e/interchaintestv8/ibc_eureka_test.go#L736). The denom is built as token address plus the hop `(transfer, ethereum light client id on Cosmos)`.

### 4. Relayer builds Ethereum acknowledgement transaction

After Cosmos writes the acknowledgement, the e2e test calls `RelayByTx` from Cosmos -> Ethereum at [`ibc_eureka_test.go#L780`](../e2e/interchaintestv8/ibc_eureka_test.go#L780).

This leg is handled by the Cosmos -> Ethereum module, not the ETH -> Cosmos module. The generated Ethereum calldata targets `ICS26Router` and calls [`ackPacket`](../contracts/ICS26Router.sol#L211). The response `address` should be the Ethereum `ICS26Router` address, as asserted in the e2e test.

Router acknowledgement flow:

1. [`ICS26Router.ackPacket`](../contracts/ICS26Router.sol#L211) checks one payload.
2. It validates `sourceClient`/`destClient` counterparty relation.
3. It verifies proof of the acknowledgement commitment on Cosmos through the Ethereum-side light client.
4. It deletes the original packet commitment via [`checkAndDeletePacketCommitment`](../contracts/utils/IBCStoreUpgradeable.sol#L69).
5. It calls [`ICS20Transfer.onAcknowledgementPacket`](../contracts/ICS20Transfer.sol#L373).
6. It emits `AckPacket`.

ICS-20 acknowledgement handling:

- Success ack: sender callback is invoked with success and escrow remains locked for native Ethereum ERC-20 that has left Ethereum.
- Universal error ack: [`_refundTokens`](../contracts/ICS20Transfer.sol#L411) returns tokens to the sender and callback is invoked with failure.

## Flow 2: Cosmos Return Transfer -> Ethereum Receive -> Cosmos Ack

The round-trip test sends the Cosmos voucher back to Ethereum at [`ibc_eureka_test.go#L827`](../e2e/interchaintestv8/ibc_eureka_test.go#L827).

High-level flow:

1. Cosmos user broadcasts IBC v2 `MsgSendPacket` with ABI-encoded `FungibleTokenPacketData`.
2. Cosmos writes a packet commitment.
3. Relayer builds Ethereum calldata for `ICS26Router.recvPacket`; see the e2e receive step at [`ibc_eureka_test.go#L882`](../e2e/interchaintestv8/ibc_eureka_test.go#L882).
4. [`ICS26Router.recvPacket`](../contracts/ICS26Router.sol#L144) verifies the Cosmos packet commitment proof.
5. Router writes packet receipt with [`setPacketReceipt`](../contracts/utils/IBCStoreUpgradeable.sol#L90).
6. Router calls [`ICS20Transfer.onRecvPacket`](../contracts/ICS20Transfer.sol#L295).
7. `ICS20Transfer` detects whether the token is returning to origin:
   - If the denom has this chain's prefix, it removes the first hop and unlocks the native ERC-20 or mapped voucher.
   - Otherwise it creates/mints an `IBCERC20` voucher.
8. Router commits the app acknowledgement via [`commitPacketAcknowledgement`](../contracts/utils/IBCStoreUpgradeable.sol#L108).
9. Router emits `WriteAcknowledgement`.
10. The ETH -> Cosmos relayer converts that Ethereum `WriteAcknowledgement` into Cosmos `MsgAcknowledgement`; see [`ibc_eureka_test.go#L930`](../e2e/interchaintestv8/ibc_eureka_test.go#L930).
11. Cosmos processes `MsgAcknowledgement` and removes its original packet commitment.

This direction is important for ETH -> Cosmos because the final ack back to Cosmos is generated by the ETH -> Cosmos relayer from an Ethereum `WriteAcknowledgement` event.

## Flow 3: Receive Failure and Error Acknowledgement

If the destination application fails during receive, the router/app writes an error acknowledgement instead of a success acknowledgement.

Ethereum-side router behavior:

- [`ICS26Router.recvPacket`](../contracts/ICS26Router.sol#L144) catches app callback errors.
- If `onRecvPacket` reverts with a reason, the router emits `IBCAppRecvPacketCallbackError` and sets `UNIVERSAL_ERROR_ACK`.
- It commits and emits the acknowledgement through `WriteAcknowledgement`.

Ethereum-side source ack behavior:

- [`ICS20Transfer.onAcknowledgementPacket`](../contracts/ICS20Transfer.sol#L373) compares the acknowledgement with `UNIVERSAL_ERROR_ACK`.
- On error, [`_refundTokens`](../contracts/ICS20Transfer.sol#L411) sends tokens back to the original Ethereum sender.

E2E references:

- ETH -> Cosmos receive failure and error ack back to Ethereum: [`cosmos_ethereum_ift_test.go#L867`](../e2e/interchaintestv8/cosmos_ethereum_ift_test.go#L867), relay packet at [`#L940`](../e2e/interchaintestv8/cosmos_ethereum_ift_test.go#L940), relay error ack at [`#L961`](../e2e/interchaintestv8/cosmos_ethereum_ift_test.go#L961).
- Cosmos -> Ethereum receive failure and error ack back to Cosmos: [`cosmos_ethereum_ift_test.go#L998`](../e2e/interchaintestv8/cosmos_ethereum_ift_test.go#L998).

## Flow 4: Timeout

Timeout is used when the destination chain never receives the packet before `timeoutTimestamp`.

ETH -> Cosmos timeout construction:

1. Caller passes Cosmos target-side tx hashes in `timeout_tx_ids` to the ETH -> Cosmos relayer.
2. [`relay_by_tx`](../packages/relayer/modules/eth-to-cosmos/src/lib.rs#L160) fetches timeout-related Cosmos events with `tm_listener.fetch_tx_events`.
3. [`cosmos::target_events_to_timeout_msgs`](../packages/relayer/lib/src/utils/cosmos.rs#L52) selects target-side `SendPacket` events that are expired and match `(src_client_id, dst_client_id)`.
4. Proof injection produces non-membership proof for packet receipt.
5. Cosmos processes `MsgTimeout`.

Ethereum router timeout behavior, used when Cosmos -> Ethereum packets timeout:

1. [`ICS26Router.timeoutPacket`](../contracts/ICS26Router.sol#L263) verifies counterparty relation.
2. It verifies non-membership for the receipt path.
3. It checks the verified counterparty timestamp is at or past `packet.timeoutTimestamp`.
4. It deletes the source packet commitment through [`checkAndDeletePacketCommitment`](../contracts/utils/IBCStoreUpgradeable.sol#L69).
5. It calls [`ICS20Transfer.onTimeoutPacket`](../contracts/ICS20Transfer.sol#L393), which refunds through [`_refundTokens`](../contracts/ICS20Transfer.sol#L411).
6. It emits `TimeoutPacket`.

E2E timeout references:

- ETH -> Cosmos timeout relay: [`cosmos_ethereum_ift_test.go#L723`](../e2e/interchaintestv8/cosmos_ethereum_ift_test.go#L723).
- Cosmos -> Ethereum timeout relay: [`cosmos_ethereum_ift_test.go#L833`](../e2e/interchaintestv8/cosmos_ethereum_ift_test.go#L833).

## Commitment State Machine

For a single packet sequence:

| Step | Chain | State written | Code |
| --- | --- | --- | --- |
| Send | Source | Packet commitment at `(sourceClient, sequence)` | [`commitPacket`](../contracts/utils/IBCStoreUpgradeable.sol#L51) |
| Receive | Destination | Packet receipt at `(destClient, sequence)` | [`setPacketReceipt`](../contracts/utils/IBCStoreUpgradeable.sol#L90) |
| Receive | Destination | Ack commitment at `(destClient, sequence)` | [`commitPacketAcknowledgement`](../contracts/utils/IBCStoreUpgradeable.sol#L108) |
| Acknowledge | Source | Source packet commitment deleted | [`checkAndDeletePacketCommitment`](../contracts/utils/IBCStoreUpgradeable.sol#L69) |
| Timeout | Source | Source packet commitment deleted | [`checkAndDeletePacketCommitment`](../contracts/utils/IBCStoreUpgradeable.sol#L69) |

Operational notes:

- `recvPacket` no-ops if the receipt already exists, but only after proof verification. This lets proof caching still happen.
- `ackPacket` and `timeoutPacket` no-op if the packet commitment is already gone, but only after proof verification.
- Ack and timeout are terminal for the source packet commitment.

## Relayer Modes

The ETH -> Cosmos module supports three transaction builder modes in [`EthToCosmosConfig`](../packages/relayer/modules/eth-to-cosmos/src/lib.rs#L57):

- `real`: uses Ethereum beacon light client proofs. Main builder starts at [`tx_builder.rs#L319`](../packages/relayer/modules/eth-to-cosmos/src/tx_builder.rs#L319).
- `mock`: injects mock proofs for tests. Mock relay path starts at [`tx_builder.rs#L703`](../packages/relayer/modules/eth-to-cosmos/src/tx_builder.rs#L703).
- `attested`: uses an aggregator/attestor light client. Delegates to [`AttestedTxBuilder.relay_events`](../packages/relayer/modules/eth-to-cosmos/src/tx_builder.rs#L854), which calls [`build_attestor_relay_events_tx`](../packages/relayer/lib/src/utils/cosmos_attested.rs#L311).

There is also a compatibility module that switches between v1.2 and current relayers based on the wasm client checksum: [`packages/relayer/modules/eth-to-cosmos-compat/src/lib.rs`](../packages/relayer/modules/eth-to-cosmos-compat/src/lib.rs).

## Handoff Checklist

When debugging ETH -> Cosmos transfer issues, inspect in this order:

1. Ethereum transaction receipt has `SendPacket` from `ICS26Router`.
2. Packet has expected `sourceClient`, `destClient`, `timeoutTimestamp`, one payload, `transfer` ports, ABI encoding, and encoded `FungibleTokenPacketData`.
3. Ethereum escrow balance changed as expected.
4. `RelayByTxRequest` uses Ethereum tx hash in `source_tx_ids`, correct `src_client_id`, and correct Cosmos-side `dst_client_id`.
5. ETH -> Cosmos relayer selected the correct mode and returned non-empty `tx` with empty `address`.
6. Cosmos broadcast included any needed update-client message before packet messages.
7. Cosmos tx emitted write-ack event and minted/unlocked the expected denom.
8. Ack relay back to Ethereum calls `ICS26Router.ackPacket` and emits `AckPacket`.
9. On failure, verify whether the ack is success or `UNIVERSAL_ERROR_ACK`, then follow refund logic in `ICS20Transfer.onAcknowledgementPacket`.
10. On timeout, verify proof target is non-membership of receipt and that verified counterparty time is past `timeoutTimestamp`.

