// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IVerifier {
    function verifyProof(
        uint256[8] calldata proof,
        uint256[2] calldata commitments,
        uint256[2] calldata commitmentPok,
        bytes32[2] calldata signature,
        bytes32 pubkey,
        bytes calldata message
    ) external returns (bool);
}

interface IGroth16Verifier {
    function verifyProof(
        uint256[8] calldata proof,
        uint256[2] calldata commitments,
        uint256[2] calldata commitmentPok,
        uint256[24] calldata input
    ) external view;
}