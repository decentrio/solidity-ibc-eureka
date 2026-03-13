package prover

import (
	"crypto/sha512"
	"fmt"
	"math/big"
	"os"

	"0x5ea000000/ecip-gnark/signature/eddsa"
	"0x5ea000000/ecip-gnark/utils"

	"filippo.io/edwards25519"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
)

// Prover generates Groth16 proofs for Ed25519 signature verification.
type Prover struct {
	r1cs constraint.ConstraintSystem
	pk   groth16.ProvingKey
}

// NewProver creates a new Prover by loading the pre-compiled R1CS and proving key.
func NewProver(r1csPath, pkPath string) (*Prover, error) {
	r1csFile, err := os.Open(r1csPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open r1cs file: %w", err)
	}
	defer r1csFile.Close()

	r1cs := groth16.NewCS(ecc.BN254)
	if _, err := r1cs.ReadFrom(r1csFile); err != nil {
		return nil, fmt.Errorf("failed to read r1cs: %w", err)
	}

	pk := groth16.NewProvingKey(ecc.BN254)
	pkFile, err := os.Open(pkPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open proving key file: %w", err)
	}
	defer pkFile.Close()

	if _, err := pk.ReadFrom(pkFile); err != nil {
		return nil, fmt.Errorf("failed to read proving key: %w", err)
	}

	return &Prover{
		r1cs: r1cs,
		pk:   pk,
	}, nil
}

// ProveSignature generates a Groth16 proof for a single Ed25519 signature.
//
// Parameters:
//   - sig: 64-byte Ed25519 signature (R || S)
//   - pub: 32-byte Ed25519 public key (compressed)
//   - msg: the message that was signed (canonical vote sign bytes)
//
// The hash H = SHA512(R || A || msg) is computed off-chain and passed to the circuit.
// Returns proof components ready for Solidity verification.
func (p *Prover) ProveSignature(sig, pub, msg []byte) (
	proof [8]*big.Int,
	commitments [2]*big.Int,
	commitmentPok [2]*big.Int,
	err error,
) {
	if len(sig) != 64 {
		err = fmt.Errorf("invalid signature length: %d, expected 64", len(sig))
		return
	}
	if len(pub) != 32 {
		err = fmt.Errorf("invalid public key length: %d, expected 32", len(pub))
		return
	}

	// Extract R (first 32 bytes) and S (last 32 bytes) from signature
	R := sig[:32]
	S, sErr := edwards25519.NewScalar().SetCanonicalBytes(sig[32:])
	if sErr != nil {
		err = fmt.Errorf("invalid signature scalar S: %w", sErr)
		return
	}

	// Decompress Ed25519 points to Weierstrass coordinates
	aX, aY, decompErr := utils.DecompressPoint(pub)
	if decompErr != nil {
		err = fmt.Errorf("failed to decompress public key: %w", decompErr)
		return
	}

	rX, rY, decompErr := utils.DecompressPoint(R)
	if decompErr != nil {
		err = fmt.Errorf("failed to decompress R point: %w", decompErr)
		return
	}

	// Convert S scalar to big.Int
	s := utils.ScalarToBigInt(S)

	// Compute H = SHA512(R || A || msg) off-chain
	hasher := sha512.New()
	hasher.Write(R)
	hasher.Write(pub)
	hasher.Write(msg)
	sum := hasher.Sum(nil)

	H, hErr := edwards25519.NewScalar().SetUniformBytes(sum)
	if hErr != nil {
		err = fmt.Errorf("failed to set hash scalar: %w", hErr)
		return
	}
	h := utils.ScalarToBigInt(H)

	// Build witness assignment
	assignment := PreHashCircuit[Fp25519, Fr25519]{
		Sig: eddsa.Signature[Fp25519, Fr25519]{
			R: sw_emulated.AffinePoint[Fp25519]{
				X: emulated.ValueOf[Fp25519](rX),
				Y: emulated.ValueOf[Fp25519](rY),
			},
			S: emulated.ValueOf[Fr25519](s),
		},
		Hash: emulated.ValueOf[Fr25519](h),
		Pub: eddsa.PublicKey[Fp25519, Fr25519]{
			A: sw_emulated.AffinePoint[Fp25519]{
				X: emulated.ValueOf[Fp25519](aX),
				Y: emulated.ValueOf[Fp25519](aY),
			},
		},
	}

	// Create witness
	witness, wErr := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if wErr != nil {
		err = fmt.Errorf("failed to create witness: %w", wErr)
		return
	}

	// Generate proof
	gnarkProof, pErr := groth16.Prove(p.r1cs, p.pk, witness)
	if pErr != nil {
		err = fmt.Errorf("failed to generate proof: %w", pErr)
		return
	}

	// Convert to Solidity-compatible format
	return ProofToBigInts(gnarkProof)
}

// ProofToBigInts converts a gnark Groth16 BN254 proof to uint256 arrays for Solidity.
//
// Returns:
//   - proof[8]: [Ar.X, Ar.Y, Bs.X.A0, Bs.X.A1, Bs.Y.A0, Bs.Y.A1, Krs.X, Krs.Y]
//   - commitments[2]: [Commitments[0].X, Commitments[0].Y]
//   - commitmentPok[2]: [CommitmentPok.X, CommitmentPok.Y]
func ProofToBigInts(gnarkProof groth16.Proof) (
	proof [8]*big.Int,
	commitments [2]*big.Int,
	commitmentPok [2]*big.Int,
	err error,
) {
	p, ok := gnarkProof.(*groth16_bn254.Proof)
	if !ok {
		err = fmt.Errorf("expected BN254 proof")
		return
	}

	// Initialize all big.Int pointers
	for i := range proof {
		proof[i] = new(big.Int)
	}
	for i := range commitments {
		commitments[i] = new(big.Int)
	}
	for i := range commitmentPok {
		commitmentPok[i] = new(big.Int)
	}

	// A (G1)
	p.Ar.X.BigInt(proof[0])
	p.Ar.Y.BigInt(proof[1])

	// B (G2)
	p.Bs.X.A0.BigInt(proof[2])
	p.Bs.X.A1.BigInt(proof[3])
	p.Bs.Y.A0.BigInt(proof[4])
	p.Bs.Y.A1.BigInt(proof[5])

	// C (G1)
	p.Krs.X.BigInt(proof[6])
	p.Krs.Y.BigInt(proof[7])

	// Commitment proof of knowledge
	p.CommitmentPok.X.BigInt(commitmentPok[0])
	p.CommitmentPok.Y.BigInt(commitmentPok[1])

	// Commitments
	p.Commitments[0].X.BigInt(commitments[0])
	p.Commitments[0].Y.BigInt(commitments[1])

	return
}
