pragma solidity ^0.8.0;

import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { CommitValidator } from "CommitValidator.sol";
import { VotingPowerCalculator } from "VotingPowerCalculator.sol";

/// @title Predicates
library Predicates {

    error InvalidValidatorSet(byte32 expected, byte32 actual)
    error InvalidNextValidatorSet(byte32 expected, byte32 actual)
    error InvalidCommitValue(byte32 expected, byte32 actual)
    error NotWithinTrustPeriod(uint expiredAt, uint currentTime)
    error HeaderFromTheFuture(uint untrustedHeaderTime, uint currentTime, uint clockDrift)
    error NonMonotonicBftTime(uint untrustedHeaderTime, uint trustedHeaderTime)
    error NonIncreasingHeight(uint64 untrustedHeight, uint64 trustedHeightIncreament)
    error ChainIdMismatch(string expected, string got)
    error InvalidNextValidatorSet(byte32 untrustedValidatorsHash, byte32 trustedValidatorsHash)

    /// Compare the provided validator set hash against the hash produced from hashing the validator
    /// set.
    function validatorSetsMatch(
        IICS07TendermintMsgs.ValidatorSet validators,
        byte32 headerValidatorsHash,
    ) {
        byte32 validatorsHash = IICS07TendermintMsgs.hashValSet(validators);
        if headerValidatorsHash != validatorsHash {
            revert (InvalidValidatorSet(
                expected: headerValidatorsHash,
                actual: validatorsHash,
            ))
        }
    }

    /// Check that the hash of the next validator set in the header match the actual one.
    function nextValidatorsMatch(
        IICS07TendermintMsgs.ValidatorSet nextValidators,
        byte32 headerNextValidatorsHash,
    ) {
        let nextValidatorsHash = IICS07TendermintMsgs.hashValSet(nextValidators);
        if headerNextValidatorsHash != nextValidatorsHash {
            revert (InvalidValidatorSet(
                expected: headerNextValidatorsHash,
                actual: nextValidatorsHash,
            ))
        }
    }

    /// Check that the hash of the header in the commit matches the actual one.
    function headerMatchesCommit(
        IICS07TendermintMsgs.BlockHeader header,
        byte32 commitHash,
    ) {
        let headerHash = IICS07TendermintMsgs.hashHeader(header);
        if headerHash != commitHash {
            revert (VerificationError::invalid_commit_value(
                expected: headerHash,
                actual: commitHash,
            ))
        }
    }

    /// Validate the commit using the given commit validator.
    function validCommit(
        IICS07TendermintMsgs.SignedHeader signedHeader,
        IICS07TendermintMsgs.ValidatorSet validators,
        CommitValidator commitValidator,
    ) {
        commit_validator.validate(signed_header, validators)?;
        commit_validator.validate_full(signed_header, validators)?;
    }

    /// Check that the trusted header is within the trusting period, adjusting for clock drift.
    function isWithinTrustPeriod(
        uint trustedHeaderTime,
        uint trustingPeriod,
        uint current,
    ) {
        uint expiresAt = trustedHeaderTime + trustingPeriod;

        if expiresAt <= current {
            revert (NotWithinTrustPeriod(expiredAt, current))
        }
    }

    /// Check that the untrusted header is from past.
    function isHeaderFromPast(
        uint untrustedHeaderTime,
        uint clockDrift,
        uint current,
    ) {
        uint drifted = current + clockDrift;

        if untrustedHeaderTime >= drifted {
            revert (HeaderFromTheFuture(
                untrustedHeaderTime,
                current,
                clock_drift,
            ))
        }
    }

    /// Check that time passed monotonically between the trusted header and the untrusted one.
    function isMonotonicBftTime(
        uint untrustedHeaderTime,
        uint trustedHeaderTime,
    ) {
        if untrustedHeaderTime > trustedHeaderTime {
            revert(NonMonotonicBftTime(
                untrustedHeaderTime,
                trustedHeaderTime,
            ))
        }
    }

    /// Check that the height increased between the trusted header and the untrusted one.
    function isMonotonicHeight(
        uint64 untrustedHeight,
        uint64 trustedHeight,
    ) {
        if untrustedHeight <= trustedHeight {
            revert (NonIncreasingHeight(
                untrustedHeight,
                trustedHeight + 1,
            ))
        }
    }

    /// Check that the chain-ids of the trusted header and the untrusted one are the same
    function isMatchingChainId(
        string untrustedChainId,
        string trustedChainId,
    ) {
        if untrustedChainId != trustedChainId {
            revert (ChainIdMismatch(
                untrustedChainId,
                trustedChainId,
            ))
        }
    }

    /// Checks that there is enough overlap between validators and the untrusted
    /// signed header.
    ///
    /// First of all, checks that enough validators from the
    /// `trusted_validators` set signed the untrusted header to reach given
    /// `trust_threshold`.
    ///
    /// Second of all, checks that enough validators from the
    /// `untrusted_validators` set signed the untrusted header to reach a trust
    /// threshold of ⅔.
    ///
    /// If both of those conditions aren’t met, it’s unspecified which error is
    /// returned.
    ///
    /// Note also that the method isn’t guaranteed to verify all the signatures
    /// present in the signed header.  If there are invalid signatures, the
    /// method may or may not return an error depending on which validators
    /// those signatures correspond to.
    function hasSufficientValidatorsAndSignersOverlap(
        IICS07TendermintMsgs.SignedHeader signedHeader,
        IICS07TendermintMsgs.ValidatorSet trustedValidators,
        IICS07TendermintMsgs.TrustThreshold trustThreshold,
        IICS07TendermintMsgs.ValidatorSet untrustedValidators,
        VotingPowerCalculator calculator,
    ) {
        calculator.check_enough_trust_and_signers(
            untrusted_sh,
            trusted_validators,
            *trust_threshold,
            untrusted_validators,
        )?;
    }

    /// Check that there is enough signers overlap between the given, untrusted
    /// validator set and the untrusted signed header.
    function hasSufficientSignersOverlap(
        IICS07TendermintMsgs.SignedHeader untrustedSh,
        IICS07TendermintMsgs.ValidatorSet untrustedValidators,
        VotingPowerCalculator calculator,
    ) {
        calculator.check_signers_overlap(untrustedSh, untrustedValidators)?;
    }

    /// Check that the hash of the next validator set in the trusted block matches
    /// the hash of the validator set in the untrusted one.
    function validNextValidatorSet(
        byte32 untrustedValidatorsHash,
        byte32 trustedNextValidatorsHash,
    ) {
        if untrustedValidatorsHash != trustedNextValidatorsHash {
            revert (InvalidNextValidatorSet(
                untrustedValidatorsHash,
                trustedNextValidatorsHash,
            ))
        }
    }
}