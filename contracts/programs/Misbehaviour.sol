// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import { IMisbehaviourMsgs } from "../light-clients/msgs/IMisbehaviourMsgs.sol";
import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";

/**
 * @title Misbehavior
 * @dev Contract to verify misbehavior
 * Converted from Rust zkVM code for Cosmos SDK proof verification
 */
contract Misbehavior {

    // Custom errors
    error InvalidClientId();
    error InvalidChainId(string id);
    error ChainIdMismatch();
    error MisbehaviourVerificationFailed();
    error CheckForMisbehaviourFailed();
    error MisbehaviourNotDetected();

    function checkForMisbehaviour (
        client_state ClientState,
    IMisbehaviorMsgs.Misbehaviour misbehaviour,
    IICS07TendermintMsgs.ConsensusState trustedConsensusState1,
    IICS07TendermintMsgs.ConsensusState trustedConsensusState2,
    u128 time
    ) public returns (IMisbehaviorMsgs.MisbehaviourOutput memory output) {
        require(validateBasic(misbehaviour), );
    }

    function validateBasic(IMisbehaviorMsgs.Misbehaviour misbehaviour) {

    }

    function validateHeaderBasic(IMisbehaviorMsgs.Header header) {

    }
    
}