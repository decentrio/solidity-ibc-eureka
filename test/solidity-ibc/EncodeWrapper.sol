// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

// solhint-disable custom-errors

import { Encode } from "../../contracts/utils/Encode.sol";
import { Header } from "../../contracts/utils/Header.sol";
import { IICS07TendermintMsgs } from "../../contracts/light-clients/msgs/IICS07TendermintMsgs.sol";

/// @title EncodeWrapper
/// @notice Wrapper with primitive parameters for cross-validation via cast call.
///         Used by Go test (operator/cmd/encode_debug/cross_test.go) to call
///         Solidity encoding functions and compare output with Go proto.Marshal().
contract EncodeWrapper {
    function encodeVersion(uint64 blockVersion, uint64 appVersion) external pure returns (bytes memory) {
        return Encode.encodeVersion(IICS07TendermintMsgs.Version(blockVersion, appVersion));
    }

    function encodeValidator(bytes32 pubKey, uint64 votingPower) external pure returns (bytes memory) {
        return Encode.encodeValidator(IICS07TendermintMsgs.SimpleValidator(pubKey, votingPower));
    }

    function encodePartSetHeader(uint32 total, bytes32 hashData) external pure returns (bytes memory) {
        return Encode.encodePartSetHeader(IICS07TendermintMsgs.PartSetHeader(total, hashData));
    }

    function encodeBlockId(
        bytes32 hashData,
        uint32 pshTotal,
        bytes32 pshHash
    ) external pure returns (bytes memory) {
        return Encode.encodeBlockId(
            IICS07TendermintMsgs.BlockId(hashData, IICS07TendermintMsgs.PartSetHeader(pshTotal, pshHash))
        );
    }

    function cdcEncodeString(string calldata value) external pure returns (bytes memory) {
        return Encode.cdcEncodeString(value);
    }

    function cdcEncodeInt64(uint256 value) external pure returns (bytes memory) {
        return Encode.cdcEncodeInt64(value);
    }

    function cdcEncodeBytes32(bytes32 value) external pure returns (bytes memory) {
        return Encode.cdcEncodeBytes32(value);
    }

    function encodeTimestamp(uint256 secs) external pure returns (bytes memory) {
        return Encode.encodeTimestamp(secs);
    }

    function merkleHash(bytes[] calldata items) external pure returns (bytes32) {
        return Header.merkleHash(items);
    }

    function hashValSet(
        bytes32[] calldata pubKeys,
        uint64[] calldata votingPowers
    ) external pure returns (bytes32) {
        require(pubKeys.length == votingPowers.length, "length mismatch");

        IICS07TendermintMsgs.ValidatorInfo[] memory vals =
            new IICS07TendermintMsgs.ValidatorInfo[](pubKeys.length);
        for (uint256 i = 0; i < pubKeys.length; i++) {
            vals[i] = IICS07TendermintMsgs.ValidatorInfo({
                valAddress: hex"",
                pubKey: pubKeys[i],
                votingPower: votingPowers[i],
                proposerPriority: 0
            });
        }

        IICS07TendermintMsgs.ValidatorSet memory valSet = IICS07TendermintMsgs.ValidatorSet({
            validators: vals,
            hasProposer: false,
            proposer: IICS07TendermintMsgs.ValidatorInfo({
                valAddress: hex"",
                pubKey: bytes32(0),
                votingPower: 0,
                proposerPriority: 0
            }),
            totalVotingPower: 0
        });

        return Header.hashValSet(valSet);
    }
}
