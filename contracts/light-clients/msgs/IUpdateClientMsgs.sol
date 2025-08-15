// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { ISP1Msgs } from "./ISP1Msgs.sol";
import { IICS07TendermintMsgs } from "./IICS07TendermintMsgs.sol";
import { IICS02ClientMsgs } from "../../msgs/IICS02ClientMsgs.sol";


/// @title Update Client Program Messages
/// @author srdtrk
/// @notice Defines shared types for the update client program.
interface IUpdateClientMsgs {
    /// @notice The message that is submitted to the updateClient function.
    /// @param sp1Proof The SP1 proof for updating the client.
    struct MsgUpdateClient {
        ISP1Msgs.SP1Proof sp1Proof;
    }

    /// @notice The public value output for the sp1 update client program.
    /// @param clientState The client state that was used to verify the header.
    /// @param trustedConsensusState The trusted consensus state.
    /// @param newConsensusState The new consensus state with the verified header.
    /// @param time The time which the header was verified in unix nanoseconds.
    /// @param trustedHeight The trusted height.
    /// @param newHeight The new height.
    struct UpdateClientOutput {
        IICS07TendermintMsgs.ClientState clientState;
        IICS07TendermintMsgs.ConsensusState trustedConsensusState;
        IICS07TendermintMsgs.ConsensusState newConsensusState;
        uint128 time;
        IICS02ClientMsgs.Height trustedHeight;
        IICS02ClientMsgs.Height newHeight;
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

    struct BlockHeader {
        uint32 version;
        string chainId;
        IICS02ClientMsgs.Height height;
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

   struct ConsensusState {
        uint128 timestamp;
        bytes32 root;
        bytes32 nextValidatorsHash;
    }
}
