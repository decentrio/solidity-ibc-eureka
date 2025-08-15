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
        IICS07TendermintMsgs.ClientState clientState,
        IMisbehaviorMsgs.Misbehaviour misbehaviour,
        IICS07TendermintMsgs.ConsensusState trustedConsensusState1,
        IICS07TendermintMsgs.ConsensusState trustedConsensusState2,
        u128 time
    ) public returns (IMisbehaviorMsgs.MisbehaviourOutput memory output) {
        // check client state id and misbehaviour chain id
        require(clientState.chainId == misbehaviour.signedHeader.header.chainId, ChainIdMismatch());

        // validate basic
        require(validateBasic(misbehaviour), "validate basic fail");
    }

    function validateBasic(IMisbehaviorMsgs.Misbehaviour misbehaviour)  {
        if misbehaviour.header1.signedHeader.header.chainId != self.header2.signedHeader.header.chainId
        {
            return Err(TendermintClientError::MismatchedHeaderChainIds {
                expected: self.header1.signed_header.header.chain_id.to_string(),
                actual: self.header2.signed_header.header.chain_id.to_string(),
            });
        }

        if misbehaviour.header1.height() < misbehaviour.header2.height() {
            return Err(
                TendermintClientError::InsufficientMisbehaviourHeaderHeight {
                    height_1: self.header1.height(),
                    height_2: self.header2.height(),
                },
            );
        }
    }

    function validateHeaderBasic(IMisbehaviorMsgs.Header header) {
        if header.height.revisionNumber != header.trustedHeight.revisionNumber {
            return Err(TendermintClientError::MismatchedRevisionHeights {
                expected: self.trusted_height.revision_number(),
                actual: self.height().revision_number(),
            });
        }

        // We need to ensure that the trusted height (representing the
        // height of the header already on chain for which this client update is
        // based on) must be smaller than height of the new header that we're
        // installing.
        if header.trustedHeight >= header.height {
            return Err(TendermintClientError::InvalidHeaderHeight(
                self.height().revision_height(),
            ));
        }

        let validatorsHash = hashValset();

        if validatorsHash != header.signedHeader.header.validatorsHash {
            return Err(TendermintClientError::MismatchedValidatorHashes {
                expected: self.signed_header.header.validators_hash,
                actual: validators_hash,
            });
        }
    }

    function hashValset(IMisbehaviorMsgs.Header header, IMembershipMsgs.HashOp hashOp) internal pure returns (bytes32) {
        
        let validator_bytes: Vec<Vec<u8>> = self
            .validators()
            .iter()
            .map(|validator| validator.hash_bytes())
            .collect();

        Hash::Sha256(merkle::simple_hash_from_byte_vectors::<H>(&validator_bytes))

        uint len = header.validators.length;
        byte32[] validatorBytes = new byte32[](len);

        for (uint i = 0; i < len; i++) {
            validatorBytes[i] = header.validators[i];
        }


        if (hashOp == IMembershipMsgs.HashOp.NO_HASH) {
            return bytesToBytes32(data);
        } else if (hashOp == IMembershipMsgs.HashOp.SHA256) {
            return sha256(data);
        } else if (hashOp == IMembershipMsgs.HashOp.KECCAK256) {
            return keccak256(data);
        } else {
            revert("Unsupported hash operation");
        }
    }
    
}