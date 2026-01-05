// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IVerifier {
    function verifyProof(
        uint256[8] calldata proof,
        uint256[1] calldata input
    ) external view;
}