pragma solidity ^0.8.0;

import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { VotingPowerCalculator } from "./VotingPowerCalculator.sol";
import { Header } from "./Header.sol";

/// @title Predicates
library Predicates {
    function verifyValSets(
        IICS07TendermintMsgs.UntrustedBlockState memory untrustedState
    ) internal pure {
        // Ensure the header validator hashes match the given validators
        bytes32 valSetHash = Header.hashValSet(untrustedState.validatorSet);
        if (valSetHash != untrustedState.signedHeader.header.validatorsHash) {
            revert("invalid block: validator set hash mismatch");
        }

        // Ensure the header matches the commit
        bytes32 headerHash = Header.hashHeader(untrustedState.signedHeader.header);
        if (headerHash != untrustedState.signedHeader.commit.blockId.hashData) {
            revert("invalid block: header hash mismatch");
        }

        // Additional implementation specific validation
        validateCommit(untrustedState.signedHeader, untrustedState.validatorSet);
    }

    function verifyAgainstTrusted(
        IICS07TendermintMsgs.UntrustedBlockState memory untrustedState,
        IICS07TendermintMsgs.TrustedBlockState memory trustedState,
        uint64 trustingPeriod,
        uint128 time
    ) internal pure {
        // Ensure the latest trusted header hasn't expired
        if (time < trustedState.headerTime || time - trustedState.headerTime > trustingPeriod) {
            revert("invalid block: untrusted state is outside of trusting period");
        }

        // Check that the untrusted block is more recent than the trusted state
        require(untrustedState.signedHeader.header.time > trustedState.headerTime, "invalid block: non monotonic bft time");

        // Check that the chain-id of the untrusted block matches that of the trusted state
        require(keccak256(abi.encodePacked(untrustedState.signedHeader.header.chainId)) == keccak256(abi.encodePacked(trustedState.chainId)), "invalid block: chain-id mismatch");

        uint64 trustedNextHeight = trustedState.height + 1;

        if (untrustedState.signedHeader.header.height == trustedNextHeight) {
            // If the untrusted block is the very next block after the trusted block,
            // check that their (next) validator sets hashes match.
            require(untrustedState.signedHeader.header.validatorsHash == trustedState.nextValidatorHash, "invalid block: next validator set hash mismatch");
        } else {
            // Otherwise, ensure that the untrusted block has a greater height than
            // the trusted block.
            require(untrustedState.signedHeader.header.height > trustedNextHeight, "invalid block: non increasing height");
        }
    }

    /// Verify that a) there is enough overlap between the validator sets of the
    /// trusted and untrusted blocks and b) more than 2/3 of the validators
    /// correctly committed the block.
    function verifyCommitAgainstTrusted(
        IICS07TendermintMsgs.UntrustedBlockState memory untrustedState,
        IICS07TendermintMsgs.TrustedBlockState memory trustedState,
        IICS07TendermintMsgs.Options memory options
    ) internal pure {
        // If the trusted validator set has changed we need to check if there’s
        // overlap between the old trusted set and the new untrested header in
        // addition to checking if the new set correctly signed the header.
        uint64 trustedNextHeight = trustedState.height + 1;
        bool needBoth = untrustedState.signedHeader.header.height != trustedNextHeight;

        if (needBoth) {
            // TODO: check_enough_trust_and_signers
        } else {
            /// Check that there is enough signers overlap between the given, untrusted
            /// validator set and the untrusted signed header.
            IICS07TendermintMsgs.TrustThreshold memory trustThreshold = IICS07TendermintMsgs.TrustThreshold({
                numerator: 2,
                denominator: 3
            });

            // TODO: check enough power
        }
    }

    function validateCommit(
        IICS07TendermintMsgs.SignedHeader memory signedHeader,
        IICS07TendermintMsgs.ValidatorSet memory validators
    ) internal pure {
        IICS07TendermintMsgs.CommitSig[] memory commitSigs = signedHeader.commit.commitSigs;

        bool hasPresentSignature = false;
        for (uint256 i = 0; i < commitSigs.length; i++) {
            if (commitSigs[i].flag != IICS07TendermintMsgs.CommitSigFlag.BLOCK_ID_FLAG_ABSENT) {
                hasPresentSignature = true;
            }
        }

        if (!hasPresentSignature) {
            revert("invalid commit: no present signatures");
        }

        if (commitSigs.length != validators.validators.length) {
            revert("invalid commit: number of signatures does not match number of validators");
        }

        for (uint256 i = 0; i < commitSigs.length; i++) {
            IICS07TendermintMsgs.CommitSig memory sig = commitSigs[i];
            bytes memory validatorAddress;

            if (sig.flag == IICS07TendermintMsgs.CommitSigFlag.BLOCK_ID_FLAG_ABSENT) {
                continue;
            } else {
                validatorAddress = sig.data.validatorAddress;
            }

            bool found = false;
            for (uint256 j = 0; j < validators.validators.length; j++) {
                if (keccak256(abi.encodePacked(validators.validators[j].valAddress)) == keccak256(abi.encodePacked(validatorAddress))) {
                    found = true;
                    break;
                }
            }
            if (!found) {
                revert("invalid commit: faulty signer");
            }
        }
    }

// TODO: calculator
    // /// Checks that there is enough overlap between validators and the untrusted
    // /// signed header.
    // ///
    // /// First of all, checks that enough validators from the
    // /// `trusted_validators` set signed the untrusted header to reach given
    // /// `trust_threshold`.
    // ///
    // /// Second of all, checks that enough validators from the
    // /// `untrusted_validators` set signed the untrusted header to reach a trust
    // /// threshold of ⅔.
    // ///
    // /// If both of those conditions aren’t met, it’s unspecified which error is
    // /// returned.
    // ///
    // /// Note also that the method isn’t guaranteed to verify all the signatures
    // /// present in the signed header.  If there are invalid signatures, the
    // /// method may or may not return an error depending on which validators
    // /// those signatures correspond to.
    // function hasSufficientValidatorsAndSignersOverlap(
    //     IICS07TendermintMsgs.SignedHeader signedHeader,
    //     IICS07TendermintMsgs.ValidatorSet trustedValidators,
    //     IICS07TendermintMsgs.TrustThreshold trustThreshold,
    //     IICS07TendermintMsgs.ValidatorSet untrustedValidators,
    //     VotingPowerCalculator calculator,
    // ) {
    //     calculator.check_enough_trust_and_signers(
    //         untrusted_sh,
    //         trusted_validators,
    //         *trust_threshold,
    //         untrusted_validators,
    //     )?;
    // }

    // /// Check that there is enough signers overlap between the given, untrusted
    // /// validator set and the untrusted signed header.
    // function hasSufficientSignersOverlap(
    //     IICS07TendermintMsgs.SignedHeader untrustedSh,
    //     IICS07TendermintMsgs.ValidatorSet untrustedValidators,
    //     VotingPowerCalculator calculator,
    // ) {
    //     calculator.check_signers_overlap(untrustedSh, untrustedValidators)?;
    // }
}