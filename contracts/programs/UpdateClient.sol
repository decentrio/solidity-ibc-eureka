// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import { IUpdateClientMsgs } from "../light-clients/msgs/IUpdateClientMsgs.sol";
import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { IICS02ClientMsgs } from "../msgs/IICS02ClientMsgs.sol";
import { IUpdateClient } from "../interfaces/IUpdateClient.sol";
import { HeightCmp } from "../utils/HeightCmp.sol";

import { Strings } from "@openzeppelin-contracts/utils/Strings.sol";

contract UpdateClient is IUpdateClient {
    string clientId = "07-tendermint-0";

    error MismatchedRevisionHeight(
        uint64 expected,
        uint64 actual
    );
    error InvalidHeaderHeight(
        uint64 height
    );
    error ValSetHashMismatch(
        bytes32 expected,
        bytes32 actual
    );
    error HeaderChainIdMismatch(
        string expected,
        string actual
    );

    function updateClient(
        IICS07TendermintMsgs.ClientState memory clientState,
        IICS07TendermintMsgs.ConsensusState memory trustedConsensusState,
        IUpdateClientMsgs.Header memory proposedHeader,
        uint128 time
    ) view external {
        IUpdateClientMsgs.ChainId memory chainId = getChainId(clientState.chainId);
        IUpdateClientMsgs.Options memory options = IUpdateClientMsgs.Options({
            trustThreshold: clientState.trustLevel,
            trustingPeriod: clientState.trustingPeriod,
            clockDrift: 15
        });

        IUpdateClientMsgs.ClientConsensusStatePath memory path = IUpdateClientMsgs.ClientConsensusStatePath({
            clientId: clientId,
            revisionNumber: proposedHeader.trustedHeight.revisionNumber,
            revisionHeight: proposedHeader.trustedHeight.revisionHeight
        });

        verifyHeader(
            proposedHeader,
            clientId,
            chainId,
            options,
            time,
            path,
            trustedConsensusState
        );
    }

    function verifyHeader(
        IUpdateClientMsgs.Header memory proposedHeader,
        string memory clientId,
        IUpdateClientMsgs.ChainId memory chainId,
        IUpdateClientMsgs.Options memory options,
        uint128 time,
        IUpdateClientMsgs.ClientConsensusStatePath memory path,
        IICS07TendermintMsgs.ConsensusState memory trustedConsensusState
    ) internal pure {
        // Checks that the header fields are valid.
        validateBasic(proposedHeader);

        // The tendermint-light-client crate though works on heights that are assumed
        // to have the same revision number. We ensure this here.
        verifyChainIdVersion(chainId, proposedHeader);

        // Delegate to tendermint-light-client, which contains the required checks
        // of the new header against the trusted consensus state.
        {
            // TODO: implement
        }


    }


    function verifyChainIdVersion(
        IUpdateClientMsgs.ChainId memory chainId,
        IUpdateClientMsgs.Header memory header
    ) internal pure {
        IUpdateClientMsgs.ChainId memory headerChainId = getChainId(header.signedHeader.header.chainId);
        if (chainId.revisionNumber != headerChainId.revisionNumber) {
            revert HeaderChainIdMismatch(
                header.signedHeader.header.chainId,
                chainId.id
            );
        }
    }

    function validateBasic(
        IUpdateClientMsgs.Header memory header
    ) internal pure {
        IUpdateClientMsgs.ChainId memory chainId = getChainId(header.signedHeader.header.chainId);
        if (chainId.revisionNumber != header.trustedHeight.revisionNumber) {
            revert MismatchedRevisionHeight(
                chainId.revisionNumber,
                header.trustedHeight.revisionNumber
            );
        }

        IICS02ClientMsgs.Height memory height = IICS02ClientMsgs.Height({
            revisionNumber: chainId.revisionNumber,
            revisionHeight: header.signedHeader.header.height
        });

        // We need to ensure that the trusted height (representing the
        // height of the header already on chain for which this client update is
        // based on) must be smaller than height of the new header that we're
        // installing.
        if (HeightCmp.ge(header.trustedHeight, height)) {
            revert InvalidHeaderHeight(
                height.revisionHeight
            );
        }

        bytes32 valSetHash = hashValSet(header.validatorSet);
        if (valSetHash != header.signedHeader.header.validatorsHash) {
            revert ValSetHashMismatch(
                header.signedHeader.header.validatorsHash,
                valSetHash
            );
        }
    }


    function hashValSet(
        IUpdateClientMsgs.ValidatorSet memory valset
    ) pure internal returns (bytes32) {
        bytes[] memory validatorBytes = new bytes[](valset.validators.length);
        for (uint256 i = 0; i < valset.validators.length; i++) {
            bytes memory validatorHash = encodeToVec(
                IUpdateClientMsgs.SimpleValidator({
                    pubKey: valset.validators[i].pubKey,
                    votingPower: valset.validators[i].votingPower
                })
            );
            validatorBytes[i] = validatorHash;

        }


        return bytes32(0);
    }

    function merkleHash(
        bytes[] memory bytesArray
    ) internal pure returns (bytes32) {
        if (bytesArray.length == 0) {
            return bytes32(0);
        }

        // tmhash(0x00 || leaf)
        // Pre and post-conditions: the hasher is in the reset state
        // before and after calling this function.
        if (bytesArray.length == 1) {
            return sha256(abi.encode(0x00, bytesArray[0]));
        }

        
        bytes32[] memory leafHashes = new bytes32[](bytesArray.length);
        for (uint256 i = 0; i < bytesArray.length; i++) {
            leafHashes[i] = keccak256(bytesArray[i]);
        }

        return bytes32(0);
    }

    function getChainId(string memory id) pure internal returns (IUpdateClientMsgs.ChainId memory) {
        // TODO: validate id
        bytes memory chainIdBytes = bytes(id);

        // Find the last occurrence of '-'
        int256 lastDashIndex = -1;
        for (uint256 i = chainIdBytes.length-1; i >= 0; i--) {
            if (chainIdBytes[i] == 0x2D) { // '-' character
                lastDashIndex = int256(i);
                break;
            }
        }

        if (lastDashIndex == -1) {
            if (1 <= chainIdBytes.length && chainIdBytes.length < 64) {
                return IUpdateClientMsgs.ChainId({
                    id: id, 
                    revisionNumber: 0
                });
            } else {
                revert("Invalid chain ID length");
            }
        }

        uint256 dashIndex = uint256(lastDashIndex);
        // Extract revision number string
        bytes memory revisionBytes = new bytes(chainIdBytes.length - dashIndex - 1);
        for (uint256 i = 0; i < revisionBytes.length; i++) {
            revisionBytes[i] = chainIdBytes[dashIndex + 1 + i];
        }

        // Validates the revision number not to start with leading zeros, like "01".
        // Zero is the only allowed revision number with leading zero.
        if (revisionBytes.length == 0 || (revisionBytes[0] == 0x30 && revisionBytes.length > 1)) {
            if (1 <= chainIdBytes.length && chainIdBytes.length < 64) {
                return IUpdateClientMsgs.ChainId({
                    id: id, 
                    revisionNumber: 0
                });
            } else {
                revert("Invalid chain ID length");
            }
        }

        (bool success, uint256 parsedRevisionNumber) = Strings.tryParseUint(string(revisionBytes));
        if (! success || parsedRevisionNumber > type(uint64).max) {
            if (1 <= chainIdBytes.length && chainIdBytes.length < 64) {
                return IUpdateClientMsgs.ChainId({
                    id: id, 
                    revisionNumber: 0
                });
            } else {
                revert("Invalid chain ID length");
            }
        }

        uint64 revisionNumber = uint64(parsedRevisionNumber);
        
        // Extract chain name
        bytes memory chainNameBytes = new bytes(dashIndex);
        for (uint256 i = 0; i < dashIndex; i++) {
            chainNameBytes[i] = chainIdBytes[i];
        }

        uint256 min = 1;
        // Prefix must be at most `max_id_length - 21` characters long since the
        // longest identifier we can construct is `{prefix}-{u64::MAX}` which
        // extends prefix by 21 characters.
        uint256 max = 43;
        if (chainNameBytes.length <= min || chainNameBytes.length >= max) {
            revert("Invalid chain prefix length");
        }

        return IUpdateClientMsgs.ChainId({
            id: id, 
            revisionNumber: revisionNumber
        });
    }

    function encodeVarint(uint256 value) internal pure returns (bytes memory) {
        if (value < 128) {
            return abi.encodePacked(uint8(value));
        }
        
        bytes memory result;
        while (value >= 128) {
            result = abi.encodePacked(result, uint8((value & 0x7F) | 0x80));
            value >>= 7;
        }
        result = abi.encodePacked(result, uint8(value));
        return result;
    }

    function encodeToVec(
        IUpdateClientMsgs.SimpleValidator memory validator
    ) internal pure returns (bytes memory) {
        bytes memory encoded = new bytes(0);
        
        // Encode field 1: pub_key (tag = 1, wire type = 2 for length-delimited)
        // Tag byte: (field_number << 3) | wire_type = (1 << 3) | 2 = 0x0A
        encoded = abi.encodePacked(encoded, uint8(0x0A)); // tag
        encoded = abi.encodePacked(encoded, uint8(32));   // length (32 bytes)
        encoded = abi.encodePacked(encoded, validator.pub_key); // data
        
        // Encode field 2: voting_power (tag = 2, wire type = 0 for varint)
        // Tag byte: (field_number << 3) | wire_type = (2 << 3) | 0 = 0x10
        encoded = abi.encodePacked(encoded, uint8(0x10)); // tag
        encoded = abi.encodePacked(encoded, encodeVarint(uint256(validator.voting_power))); // varint-encoded value

        return encoded;
    }

}