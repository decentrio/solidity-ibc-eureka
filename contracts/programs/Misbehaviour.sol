// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import { IMisbehaviourMsgs } from "../light-clients/msgs/IMisbehaviourMsgs.sol";
import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { ISP1ICS07TendermintErrors } from "../light-clients/errors/ISP1ICS07TendermintErrors.sol";


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

        _validateBasic(misbehaviour);

        Options options: {
            clientState.trustLevel,
            trustingPeriod: clientState.trustingPeriod,
            clockDrift: uint32(15),
        };

        uint currentTimestamp = block.timestamp;
        IICS07TendermintMsgs.ChainId chainId = getChainId(clientState.chainId)

        _verifyMisbehaviourHeader(
            misbehaviour.header1,
            chainId,
            options,
            trustedConsensusState1.timeStamp,
            trustedConsensusState1.nextValidatorsHash,
            current_timestamp,

        );

        _verifyMisbehaviourHeader(
            misbehaviour.header2,
            chainId,
            options,
            trustedConsensusState2.timeStamp,
            trustedConsensusState2.nextValidatorsHash,
            current_timestamp,
            
        );
    }

    function _validateBasic(IMisbehaviorMsgs.Misbehaviour misbehaviour)  {
        _validateHeaderBasic(misbehaviour.header2);
        _validateHeaderBasic(misbehaviour.header2);

        if misbehaviour.header1.signedHeader.header.chainId != self.header2.signedHeader.header.chainId
        {
            revert (ISP1ICS07TendermintErrors.ChainIdMismatch {
                expected: misbehaviour.header1.signedHeader.header.chainId,
                actual: misbehaviour.header2.signedHeader.header.chainId,
            });
        }

        if misbehaviour.header1.height() < misbehaviour.header2.height() {
            revert (
                ISP1ICS07TendermintErrors.InsufficientMisbehaviourHeaderHeight {
                    height1: misbehaviour.header1.height,
                    height2: misbehaviour.header2.height,
                },
            );
        }


    }

    function _validateHeaderBasic(IICS07TendermintMsgs.Header header) {
        if header.height.revisionNumber != header.trustedHeight.revisionNumber {
            revert (ISP1ICS07TendermintErrors.MismatchedRevisionHeights {
                expected: header.trustedHeight.revisionNumber,
                actual: header.height.revisionNumber,
            });
        }

        // We need to ensure that the trusted height (representing the
        // height of the header already on chain for which this client update is
        // based on) must be smaller than height of the new header that we're
        // installing.
        if header.trustedHeight >= header.height {
            revert (ISP1ICS07TendermintErrors.InvalidHeaderHeight(
                header.height.revisionHeight,
            ));
        }

        byte32 validatorsHash = IICS07TendermintMsgs.hashValSet(header.validatorSet);

        if validatorsHash != header.signedHeader.header.validatorsHash {
            revert (ISP1ICS07TendermintErrors.MismatchedValidatorHashes {
                expected: header.signedHeader.header.validatorsHash,
                actual: validatorsHash,
            });
        }
    }

    function _verifyMisbehaviourHeader (
        IICS07TendermintMsgs.Header header,
        IICS07TendermintMsgs.ChainId chainId,
        IICS07TendermintMsgs.Options options,
        uint trustedTime,
        byte32 trustedNextValidatorHash,
        uint currentTimestamp,
        verifier: &impl Verifier,
    )  {
        // ensure correctness of the trusted next validator set provided by the relayer
        _checkTrustedNextValidatorSet(header, trustedNextValidatorHash)?;

        // ensure trusted consensus state is within trusting period
        {
            let durationSinceConsensusState = currentTimestamp - trustedTime
                
            if durationSinceConsensusState >= options.trustingPeriod {
                revert (ISP1ICS07TendermintErrors.InsufficientTrustingPeriod {
                    durationSinceConsensusState,
                    trustingPeriod: options.trustingPeriod,
                });
            }
        }

        // main header verification, delegated to the tendermint-light-client crate.
        let untrusted_state = header.as_untrusted_block_state();

        let tm_chain_id =
            &chain_id
                .as_str()
                .try_into()
                .map_err(|e| IdentifierError::FailedToParse {
                    description: format!("chain ID `{chain_id}`: {e:?}"),
                })?;

        let trusted_state =
            header.as_trusted_block_state(tm_chain_id, trusted_time, trusted_next_validator_hash)?;

        let current_timestamp = current_timestamp.into_host_time()?;

        verifier
            .verify_misbehaviour_header(untrusted_state, trusted_state, options, current_timestamp)

    }

    function _checkTrustedNextValidatorSet(
        IICS07TendermintMsgs.Header header,
        byte32 trustedNextValidatorHash,
    ) -> Result<(), ClientError> {
        byte32 validatorsHash = IICS07TendermintMsgs.hashValSet(header.trustedNextValidatorSet);

        if (validatorsHash != trustedNextValidatorHash) {
            revert (ISP1ICS07TendermintErrors.FailedToVerifyHeader{
                description: "trusted next validator set hash does not match hash stored on chain"
            })
        }
    }
}