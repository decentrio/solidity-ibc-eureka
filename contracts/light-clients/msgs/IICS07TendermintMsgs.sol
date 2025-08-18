// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { IICS02ClientMsgs } from "../../msgs/IICS02ClientMsgs.sol";

/// @title ICS07 Tendermint Messages
/// @author srdtrk
/// @notice Defines shared types for ICS07Tendermint implementations.
interface IICS07TendermintMsgs {
    /// @notice Fraction of validator overlap needed to update header
    /// @param numerator Numerator of the fraction
    /// @param denominator Denominator of the fraction
    struct TrustThreshold {
        uint8 numerator;
        uint8 denominator;
    }

    /// @notice Defines the ICS07Tendermint ClientState for ibc-lite
    /// @param chainId Chain ID
    /// @param trustLevel Fraction of validator overlap needed to update header
    /// @param latestHeight Latest height the client was updated to
    /// @param trustingPeriod duration of the period since the LatestTimestamp during which the
    /// submitted headers are valid for upgrade in seconds.
    /// @param unbondingPeriod duration of the staking unbonding period in seconds
    /// @param isFrozen whether or not client is frozen (due to misbehavior)
    /// @param zkAlgorithm The zk algorithm supported by this contract (for the relayers).
    struct ClientState {
        string chainId;
        TrustThreshold trustLevel;
        IICS02ClientMsgs.Height latestHeight;
        uint32 trustingPeriod;
        uint32 unbondingPeriod;
        bool isFrozen;
        SupportedZkAlgorithm zkAlgorithm;
    }

    /// @notice Defines the Tendermint light client's consensus state at some height.
    /// @param timestamp timestamp that corresponds to the counterparty block height
    /// in which the ConsensusState was generated. (in unix nanoseconds)
    /// @param root commitment root (i.e app hash)
    /// @param nextValidatorsHash next validators hash
    struct ConsensusState {
        uint128 timestamp;
        bytes32 root;
        bytes32 nextValidatorsHash;
    }

    /// @notice Defines the supported zk algorithms
    enum SupportedZkAlgorithm {
        Groth16,
        Plonk
    }

     struct Header {
        SignedHeader signedHeader;
        ValidatorSet validatorSet;
        IICS02ClientMsgs.Height trustedHeight;
        ValidatorSet trustedNextValidatorSet;
    }

    struct SignedHeader {
        BlockHeader header;
        BlockCommit commit;
    }

    struct ValidatorSet {
        ValidatorInfo[] validators;
        bool hasProposer;
        ValidatorInfo proposer;
        uint64 totalVotingPower;
    }

    struct ValidatorInfo {
        bytes valAddress;
        bytes pubKey;
        uint64 votingPower;
        int64 proposerPriority;
    }

    struct SimpleValidator {
        bytes32 pubKey;
        uint64 votingPower;
    }

    struct BlockHeader {
        uint32 version;
        string chainId;
        uint64 height;
        uint128 time;
        bool hasLastBlockId;
        BlockId lastBlockId;
        bool hasLastCommitHash;
        bytes32 lastCommitHash;
        bool hasDataHash;
        bytes32 dataHash;
        bytes32 validatorsHash;
        bytes32 nextValidatorsHash;
        bytes32 consensusHash;
        bytes appHash;
        bool hasLastResultsHash;
        bytes32 lastResultsHash;
        bool hasEvidenceHash;
        bytes32 evidenceHash;
        bytes proposerAddress;
    }

    struct BlockCommit {
        uint64 height;
        uint32 round;
        BlockId blockId;
        CommitSig[] commitSigs;
    }

    struct BlockId {
        bytes32 hashData;
        PartSetHeader partSetHeader;
    }

    struct PartSetHeader {
        uint32 total;
        bytes32 hashData;
    }

    enum CommitSigFlag {
        /// no vote was received from a validator.
        BLOCK_ID_FLAG_ABSENT,
        /// voted for the Commit.BlockID.
        BLOCK_ID_FLAG_COMMIT,
        /// voted for nil.
        BLOCK_ID_FLAG_NIL
    }

    struct CommitSigData {
        bytes validatorAddress;
        uint128 timestamp;
        bool hasSignature;
        bytes signature;
    }

    struct CommitSig {
        CommitSigFlag flag;
        CommitSigData data;
    }

    struct ChainId {
        string id;
        uint64 revisionNumber;
    }

    struct Options {
        IICS07TendermintMsgs.TrustThreshold trustThreshold;
        uint32 trustingPeriod;
        uint32 clockDrift;
    }

    struct ClientConsensusStatePath {
        string clientId;
        uint64 revisionNumber;
        uint64 revisionHeight;
    }

    struct TrustedBlockState {
        ChainId chainId;
        uint128 headerTime;
        uint64 height;
        ValidatorSet nextValidators;
        bytes32 nextValidatorsHash;
    }

    struct UnTrustedBlockState {
        SignedHeader signedHeader;
        ValidatorSet validators;
    }

    struct VotingPowerTally {
        /// Total voting power
        uint64 total;
        /// Tallied voting power
        uint64 tallied;
        /// Trust threshold for voting power
        TrustThreshold trustThreshold;
    }

    struct NonAbsentCommitVote {
        SignedVote signedVote;
        /// Flag indicating whether the signature has already been verified.
        verified: bool;
    }

    struct NonAbsentCommitVotes {
        /// Votes sorted by validator address.
        NonAbsentCommitVote[] votes,
        /// Internal buffer for storing sign_bytes.
        ///
        /// The buffer is reused for each canonical vote so that we allocate it
        /// once.
        bytes32 signBytes,
    }

    struct PartSetHeader {
        /// Number of parts in this block
        uint32 total;

        /// Hash of the parts set header,
        bytes32 hash;
    }

    struct BlockId {
        bytes32 hash;

        PartSetHeader partSetHeader;
    }

    struct CanonicalVote {
        /// Type of vote (prevote or precommit)
        VoteType voteType;

        /// Block height
        uint64 height;

        /// Round
        uint64 round;

        /// Block ID
        BlockId blockId;

        /// Timestamp
        uint timestamp,

        /// Chain ID
        ChainId chainId,
    }

    struct SignedVote {
        CanonicalVote vote;
        bytes validatorAddress,
        bytes32 signature,
    }

    enum VoteType {
        /// Votes for blocks which validators observe are valid for a given round
        Prevote = 1,

        /// Votes to commit to a particular block for a given round
        Precommit = 2,
    }

    enum Verdict {
        /// Verification succeeded, the block is valid.
        SUCCESS,
        /// The minimum voting power threshold is not reached,
        /// the block cannot be trusted yet.
        NOT_ENOUGH_TRUST,
        /// Verification failed, the block is invalid.
        INVALID
    }


    function encodeToVec(
        SimpleValidator memory validator
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

    function tallyVote(
        VotingPowerTally power,
        uint64 votingPower,
    ) pure returns (VotingPowerTally) {
        power.tallied += votingPower;
        require(power.tallied <= power.total, "tallied should be less than total voting power")
    }

    function checkTally(
        VotingPowerTally power,
    ) pure returns (bool) {
        power.tallied * power.trustThreshold.denominator > power.totalVotingPower * power.trustThreshold.numerator
    }

    function hashValSet(
        ValidatorSet memory valset
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

    function hashHeader(
        BlockHeader memory header
    ) pure internal returns (bytes32) {
        uint256 fieldCount = 9;
        
        if (header.hasLastBlockId) fieldCount++;
        if (header.hasLastCommitHash) fieldCount++;
        if (header.hasDataHash) fieldCount++;
        if (header.hasLastResultsHash) fieldCount++;
        if (header.hasEvidenceHash) fieldCount++;
        
        bytes[] memory headerBytes = new bytes[](fieldCount);
        uint256 index = 0;
        
        // Always present fields
        headerBytes[index++] = encodeVersion(header.version);
        headerBytes[index++] = encodeString(header.chainId);
        headerBytes[index++] = encodeVarint(uint256(header.height));
        headerBytes[index++] = encodeVarint(uint256(header.time));
        
        // Conditional fields
        if (header.hasLastBlockId) {
            headerBytes[index++] = encodeBlockId(header.lastBlockId);
        }
        
        if (header.hasLastCommitHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.lastCommitHash);
        }
        
        if (header.hasDataHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.dataHash);
        }
        
        // Always present fields
        headerBytes[index++] = abi.encodePacked(uint8(32), header.validatorsHash);
        headerBytes[index++] = abi.encodePacked(uint8(32), header.nextValidatorsHash);
        headerBytes[index++] = abi.encodePacked(uint8(32), header.consensusHash);
        
        // appHash is always present (no has flag in struct)
        uint256 appHashLength = header.appHash.length;
        headerBytes[index++] = abi.encodePacked(encodeVarint(appHashLength), header.appHash);
        
        // Conditional fields
        if (header.hasLastResultsHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.lastResultsHash);
        }
        
        if (header.hasEvidenceHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.evidenceHash);
        }
        
        // proposerAddress is always present (no has flag in struct)
        uint256 proposerAddressLength = header.proposerAddress.length;
        headerBytes[index++] = abi.encodePacked(encodeVarint(proposerAddressLength), header.proposerAddress);
        
        return merkleHash(headerBytes);
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
            return sha256(abi.encodePacked([0x00], bytesArray[0]));
        }

        uint256 split = nextPowerOfTwo(bytesArray.length) / 2;
        bytes32 left = merkleHash(getSlice(bytesArray, 0, split));
        bytes32 right = merkleHash(getSlice(bytesArray, split, bytesArray.length));

        return sha256(abi.encodePacked([0x01], left, right));
    }

    function getChainId(string memory id) pure internal returns (ChainId memory) {
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
                return ChainId({
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
                return ChainId({
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
                return ChainId({
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

        return ChainId({
            id: id, 
            revisionNumber: revisionNumber
        });
    }
}
