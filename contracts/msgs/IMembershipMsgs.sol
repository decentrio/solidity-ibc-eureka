// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IMembershipMsgs {
    // Struct to represent a key-value pair
    struct KVPair {
        bytes[] path;
        bytes32 value;
    }
    
    struct ProofSpec {
        SpecType specType;
        bool hasLeafSpec;
        LeafOp leafOp;
        bool hasInnerSpec;
        InnerSpec innerSpec;
        uint32 minDepth;
        uint32 maxDepth;
        bool prehashKeyBeforeComparison;
    }

    struct LeafOp {
        HashOp hashOp;
        HashOp prehashKey;
        HashOp prehashValue;
        /// prefix is a fixed bytes that may optionally be included at the beginning to differentiate
        /// a leaf node from an inner node.
        bytes prefix;
    }

    struct InnerSpec {
        /// Child order is the ordering of the children node, must count from 0
        /// iavl tree is \[0, 1\] (left then right)
        /// merk is \[0, 2, 1\] (left, right, here)
        uint32[] childOrder;
        uint32 childSize;
        uint32 minPrefixLength;
        uint32 maxPrefixLength;
        /// empty child is the prehash image that is used when one child is nil (eg. 20 bytes of 0)
        bytes emptyChild;
        HashOp hashOp;
    }

    struct Padding {
        uint32 minPrefix;
        uint32 maxPrefix;
        uint256 suffix;
    }

    struct InnerOp {
        HashOp hashOp;
        bytes prefix;
        bytes suffix;
    }

    enum HashOp {
        NO_HASH,
        SHA256,
        KECCAK256
    }

    enum SpecType {
        IAVL,
        TENDERMINT
    }    
    struct MerkleProof {
        CommitmentProof[] proofs;
    }
    
    
    struct CommitmentProof {
        ProofType proofType;
        ExistenceProof existenceProof;
        NonExistenceProof nonExistenceProof;
    }

    enum ProofType {
        EXIST,
        NON_EXIST
    }

    struct ExistenceProof {
        bytes key;
        bytes32 value;
        LeafOp leaf;
        InnerOp[] path;
    }
    
    struct NonExistenceProof {
        bytes key;
        bool hasLeft;
        ExistenceProof left;
        bool hasRight;
        ExistenceProof right;
    }

    // Struct for membership output
    struct MembershipOutput {
        bytes32 commitmentRoot;
        KVPair[] kvPairs;
    }
}