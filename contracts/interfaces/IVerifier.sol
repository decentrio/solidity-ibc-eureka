// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IVerifier {
    /// @notice Verifies a proof with given public values and vkey.
    /// @param proofA πa element of the groth16 proof.
    /// @param proofB πb element of the groth16 proof.
    /// @param proofC πc element of the groth16 proof.
    /// @param publicSignals Public signals of the circuit.

    function verifyProof(
        uint[2] calldata proofA,
        uint[2][2] calldata proofB,
        uint[2] calldata proofC,
        uint[11] calldata publicSignals
    ) external view returns (bool);
}