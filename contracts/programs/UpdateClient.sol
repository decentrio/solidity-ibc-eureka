// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import { IUpdateClientMsgs } from "../light-clients/msgs/IUpdateClientMsgs.sol";
import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { IUpdateClient } from "../interfaces/IUpdateClient.sol";

import { Strings } from "@openzeppelin-contracts/utils/Strings.sol";

contract UpdateClient is IUpdateClient {
    string clientId = "07-tendermint-0";
    function updateClient(
        IICS07TendermintMsgs.ClientState memory clientState,
        IICS07TendermintMsgs.ConsensusState memory trustedConsensusState,
        IUpdateClientMsgs.Header memory proposedHeader,
        uint128 time
    ) external {
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
}