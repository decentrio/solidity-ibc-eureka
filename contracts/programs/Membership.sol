// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import { IMembershipMsgs } from "../msgs/IMembershipMsgs.sol";
import { IMembership } from "../interfaces/IMembership.sol";
/**
 * @title MerkleTreeMembership
 * @dev Contract to verify membership of key-value pairs in a Merkle tree
 * Converted from Rust zkVM code for Cosmos SDK proof verification
 */
contract MerkleTreeMembership  is IMembership {
     
    // Events
    event MembershipVerified(bytes32 indexed commitmentRoot, uint256 pairsCount);
    event NonMembershipVerified(bytes32 indexed commitmentRoot, bytes key);
    
    // Custom errors
    error InvalidLength();
    error EmptyRequest();
    error VerificationMembershipFailed();
    error VerificationNonMembershipFailed();
    error MissingMerkleProof();
    error MissingMerkleRoot();
    error MismatchedNumberOfProofs(uint256 expected, uint256 actual);
    error MissingVerifiedValue();
    error InvalidMerkleProof();
    error InvalidExistenceProof();
    error FailedToVerifyMembership();

    /**
     * @dev Verify membership of multiple key-value pairs in the Merkle tree
     * @param appHash The root hash of the Merkle tree (32 bytes)
     * @param kvPairs Array of key-value pairs to verify
     * @param merkleProofs Array of corresponding Merkle proofs
     * @return output The membership verification result
     */
    function membership(
        bytes32 appHash,
        IMembershipMsgs.KVPair[] memory kvPairs,
        IMembershipMsgs.MerkleProof[] memory merkleProofs
    ) public returns (IMembershipMsgs.MembershipOutput memory output) {
        if (kvPairs.length == 0) {
            revert EmptyRequest();
        }
        
        if (kvPairs.length != merkleProofs.length) {
            revert InvalidLength();
        }
        
        bytes32 commitmentRoot = appHash;
        IMembershipMsgs.KVPair[] memory verifiedPairs = new KVPair[](kvPairs.length);
        
        for (uint256 i = 0; i < kvPairs.length; i++) {
            IMembershipMsgs.KVPair memory kvPair = kvPairs[i];
            IMembershipMsgs.MerkleProof memory merkleProof = merkleProofs[i];
            
            IMembershipMsgs.ProofSpec[] memory proofSpecs = [
                iavlSpec(),
                tendermintSpec()
            ];
            // Check if this is a non-membership proof (empty value)
            if (kvPair.value.length == 0) {
                // Verify non-membership
                if (!verifyNonMembership(proofSpecs, commitmentRoot, kvPair.path, merkleProof)) {
                    revert VerificationNonMembershipFailed();
                }
                emit NonMembershipVerified(commitmentRoot, kvPair.key);
            } else {
                // Verify membership
                if (!verifyMembership(proofSpecs, commitmentRoot, kvPair.path, kvPair.value, 0, merkleProof)) {
                    revert VerificationMembershipFailed();
                }
            }
            
            verifiedPairs[i] = kvPair;
        }
        
        output = MembershipOutput({
            commitmentRoot: commitmentRoot,
            kvPairs: verifiedPairs
        });
        
        emit MembershipVerified(commitmentRoot, kvPairs.length);
        return output;
    }
    
    function verifyMembership(
        IMembershipMsgs.ProofSpec[] memory proofSpecs,
        bytes32 root,
        bytes[] memory path,
        bytes memory value,
        uint256 startIndex,
        IMembershipMsgs.MerkleProof memory proof
    ) internal pure {
        if (proof.proofs.length == 0) {
            revert MissingMerkleProof();
        }

        if (root.hash == bytes32(0)) {
            revert MissingMerkleRoot();
        }

        uint256 proofLength = proof.proofs.length;
        if (proofSpecs.length != proofLength) {
            revert MismatchedNumberOfProofs(proofSpecs.length, proofLength);
        }

        if (path.length != proofLength) {
            revert MismatchedNumberOfProofs(path.length, proofLength);
        }

        if (value == bytes32(0)) {
            revert MissingVerifiedValue();
        }

        // Process proofs from startIndex onwards
        // Keys are represented from root-to-leaf, so we iterate in reverse
        bytes32 subroot = value;
        bytes32 valueUpdate = value; 

        uint256 pathLength = path.length;
        for (uint256 i = startIndex; i < proofLength; i++) {
            bytes memory keyPath = path[pathLength - i - 1];

            IMembershipMsgs.CommitmentProof memory commitmentProof = proof.proofs[i];
            if (commitmentProof.proofType != IMembershipMsgs.ProofType.EXIST) {
                revert InvalidMerkleProof();
            }

            subroot = calculateExistenceRoot(commitmentProof.existenceProof);
            if (!verifyMembershipInternal(
                commitmentProof.existenceProof,
                specs[i],
                subroot,
                path[keyPath],
                valueUpdate
                )
            ) {
                revert FailedToVerifyMembership();
            }
            valueUpdate = subroot;
        }

        if (root != subroot) {
            revert FailedToVerifyMembership();
        }
    }

    function calculateExistenceRoot(IMembershipMsgs.ExistenceProof memory proof) 
        internal 
        pure 
        returns (bytes32) 
    {
        if (proof.key.length == 0 || proof.value.length == 0) {
            revert InvalidExistenceProof();
        }

        IMembershipMsgs.LeafOp memory leafOp = proof.leaf;

        for (uint256 i = 0; i < proof.path.length; i++) {
            current = keccak256(abi.encodePacked(current, proof.path[i]));
        }
        
        return current;
    }
    
    function verifyMembershipInternal(
        IMembershipMsgs.ExistenceProof memory proof,
        IMembershipMsgs.ProofSpec memory spec,
        bytes memory subroot,
        bytes memory key,
        bytes memory value
    ) internal pure returns (bool) {

        return false;
    }
    
    /**
     * @dev Verify non-membership of a key in the Merkle tree
     * @param root The Merkle root
     * @param key The key to verify non-existence
     * @param proof The Merkle proof for non-membership
     * @return True if the non-membership proof is valid
     */
    function verifyNonMembership(
        bytes32 root,
        bytes memory key,
        MerkleProof memory proof
    ) internal pure returns (bool) {
        
        // For non-membership, we verify that the proof leads to a different key
        // or an empty position in the tree
        bytes32 keyHash = keccak256(key);
        
        // In a simple implementation, we check if the proof doesn't lead to our key
        // This is a simplified version - in practice, you'd need more sophisticated
        // non-membership proof verification depending on your tree structure
        return !verifyProof(root, keyHash, proof.proof, proof.index);
    }
    
    /**
     * @dev Generic Merkle proof verification
     * @param root The Merkle root
     * @param leaf The leaf hash to verify
     * @param proof Array of sibling hashes for the proof path
     * @param index The index of the leaf in the tree
     * @return True if the proof is valid
     */
    function verifyProof(
        bytes32 root,
        bytes32 leaf,
        bytes[] memory proof,
        uint256 index
    ) internal pure returns (bool) {
        
        bytes32 computedHash = leaf;
        uint256 currentIndex = index;
        
        for (uint256 i = 0; i < proof.length; i++) {
            bytes32 proofElement = bytesToBytes32(proof[i]);
            
            if (currentIndex % 2 == 0) {
                // If current index is even, proof element is right sibling
                computedHash = keccak256(abi.encodePacked(computedHash, proofElement));
            } else {
                // If current index is odd, proof element is left sibling
                computedHash = keccak256(abi.encodePacked(proofElement, computedHash));
            }
            
            currentIndex = currentIndex / 2;
        }
        
        return computedHash == root;
    }
    
    /**
     * @dev Convert bytes to bytes32
     * @param data The bytes to convert
     * @return result The bytes32 result
     */
    function bytesToBytes32(bytes memory data) internal pure returns (bytes32 result) {
        if (data.length >= 32) {
            assembly {
                result := mload(add(data, 32))
            }
        } else {
            // Pad with zeros if data is shorter than 32 bytes
            bytes32 temp;
            assembly {
                temp := mload(add(data, 32))
            }
            result = temp >> (8 * (32 - data.length));
        }
    }

    function hashData(bytes memory data, IMembershipMsgs.HashOp hashOp) internal pure returns (bytes32) {
        if (hashOp == IMembershipMsgs.HashOp.NO_HASH) {
            
        }
    }
    function tendermintSpec() public pure returns (IMembershipMsgs.ProofSpec memory) {
        IMembershipMsgs.LeafOp memory leaf = IMembershipMsgs.LeafOp({
            hash: IMembershipMsgs.HashOp.SHA256,
            prehashKey: 0,
            prehashValue: IMembershipMsgs.HashOp.SHA256,
            prefix: hex"00"
        });

        uint32[] memory childOrder = new uint32[](2);
        childOrder[0] = 0;
        childOrder[1] = 1;

        IMembershipMsgs.InnerSpec memory inner = IMembershipMsgs.InnerSpec({
            childOrder: childOrder,
            minPrefixLength: 1,
            maxPrefixLength: 1,
            childSize: 32,
            emptyChild: "",
            hash: IMembershipMsgs.HashOp.SHA256
        });

        return IMembershipMsgs.ProofSpec({
            leafSpec: leaf,
            innerSpec: inner,
            minDepth: 0,
            maxDepth: 0,
            prehashKeyBeforeComparison: false
        });
    }

    function iavlSpec() public pure override returns (IMembershipMsgs.ProofSpec memory) {
        IMembershipMsgs.LeafOp memory leaf = IMembershipMsgs.LeafOp({
            hash: IMembershipMsgs.HashOp.SHA256,
            prehashKey: 0,
            prehashValue: IMembershipMsgs.HashOp.SHA256,
            length: IMembershipMsgs.LengthOp.VAR_PROTO,
            prefix: hex"00"
        });

        uint32[] memory childOrder = new uint32[](2);
        childOrder[0] = 0;
        childOrder[1] = 1;

        IMembershipMsgs.InnerSpec memory inner = IMembershipMsgs.InnerSpec({
            childOrder: childOrder,
            minPrefixLength: 4,
            maxPrefixLength: 12,
            childSize: 33,
            emptyChild: "",
            hash: IMembershipMsgs.HashOp.SHA256
        });

        // Return proof spec
        return IMembershipMsgs.ProofSpec({
            leafSpec: leaf,
            innerSpec: inner,
            minDepth: 0,
            maxDepth: 0,
            prehashKeyBeforeComparison: false
        });
    }
}