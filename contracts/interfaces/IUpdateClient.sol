// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { IUpdateClientMsgs } from "../light-clients/msgs/IUpdateClientMsgs.sol";
import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
interface IUpdateClient {
     function updateClient(
        IICS07TendermintMsgs.ClientState calldata clientState,
        IICS07TendermintMsgs.ConsensusState calldata trustedConsensusState,
        IICS07TendermintMsgs.Header calldata proposedHeader,
        uint128 time
    ) view external returns (IUpdateClientMsgs.UpdateClientOutput memory);
}