package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"crypto/sha256"

	"github.com/cometbft/cometbft/crypto/merkle"
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	cmtversion "github.com/cometbft/cometbft/proto/tendermint/version"
)

// encodeVarint matches Solidity's Encode.encodeVarint exactly.
// Standard protobuf unsigned varint encoding.
func encodeVarint(value uint64) []byte {
	if value < 128 {
		return []byte{byte(value)}
	}
	var buf []byte
	for value >= 128 {
		buf = append(buf, byte((value&0x7F)|0x80))
		value >>= 7
	}
	buf = append(buf, byte(value))
	return buf
}

// encodeString matches Solidity's Encode.encodeString:
// varint(len) + string bytes
func encodeString(s string) []byte {
	b := []byte(s)
	return append(encodeVarint(uint64(len(b))), b...)
}

// encodeValidatorSolidity matches Solidity's Encode.encodeValidator:
//
//	tag 0x0A + length 32 + raw pubkey (32 bytes)
//	tag 0x10 + varint(votingPower)
//
// NOTE: This is NOT the same as proto.Marshal(SimpleValidator) because
// Solidity treats pubKey as raw 32 bytes, while Go's SimpleValidator
// wraps it in a crypto.PublicKey oneof message.
func encodeValidatorSolidity(pubKey [32]byte, votingPower int64) []byte {
	var buf []byte
	// Field 1: pub_key (tag=0x0A, length=32, data)
	buf = append(buf, 0x0A)
	buf = append(buf, 32)
	buf = append(buf, pubKey[:]...)
	// Field 2: voting_power (tag=0x10, varint)
	buf = append(buf, 0x10)
	buf = append(buf, encodeVarint(uint64(votingPower))...)
	return buf
}

// encodeVersionProto uses Go's proto marshal for version.Consensus.
// NOTE: Go proto.Marshal skips zero-value fields. Solidity always encodes them.
func encodeVersionProto(blockVersion, appVersion uint64) ([]byte, error) {
	v := &cmtversion.Consensus{
		Block: blockVersion,
		App:   appVersion,
	}
	return v.Marshal()
}

// encodeVersionSolidity matches Solidity's Encode.encodeVersion exactly.
// Unlike proto.Marshal, Solidity ALWAYS encodes fields even when zero.
func encodeVersionSolidity(blockVersion, appVersion uint64) []byte {
	var buf []byte
	// Field 1: blockVersion (tag=0x08, varint)
	buf = append(buf, 0x08)
	buf = append(buf, encodeVarint(blockVersion)...)
	// Field 2: appVersion (tag=0x10, varint) - ALWAYS encoded, even if 0
	buf = append(buf, 0x10)
	buf = append(buf, encodeVarint(appVersion)...)
	return buf
}

// encodePartSetHeaderProto uses Go's proto marshal.
// Should match Solidity's Encode.encodePartSetHeader.
func encodePartSetHeaderProto(total uint32, hash [32]byte) ([]byte, error) {
	psh := &cmttypes.PartSetHeader{
		Total: total,
		Hash:  hash[:],
	}
	return psh.Marshal()
}

// encodeBlockIDProto uses Go's proto marshal.
// Should match Solidity's Encode.encodeBlockId.
func encodeBlockIDProto(hash [32]byte, pshTotal uint32, pshHash [32]byte) ([]byte, error) {
	bid := &cmttypes.BlockID{
		Hash: hash[:],
		PartSetHeader: cmttypes.PartSetHeader{
			Total: pshTotal,
			Hash:  pshHash[:],
		},
	}
	return bid.Marshal()
}

// encodeSimpleValidatorProto uses Go's proto marshal for comparison.
// This wraps pubkey in crypto.PublicKey{Ed25519: ...} — differs from Solidity's raw encoding.
func encodeSimpleValidatorProto(pubKey [32]byte, votingPower int64) ([]byte, error) {
	sv := &cmttypes.SimpleValidator{
		PubKey: &cmtcrypto.PublicKey{
			Sum: &cmtcrypto.PublicKey_Ed25519{
				Ed25519: pubKey[:],
			},
		},
		VotingPower: votingPower,
	}
	return sv.Marshal()
}

// hashHeaderFields computes Tendermint header hash by encoding each field
// individually and then computing Merkle root.
// This matches Solidity's Header.hashHeader logic.
func hashHeaderFields(
	blockVersion, appVersion uint64,
	chainID string,
	height uint64,
	timeNanos uint64, // simplified: Solidity uses uint128 unix nanos
	hasLastBlockId bool,
	lastBlockIdHash [32]byte,
	pshTotal uint32,
	pshHash [32]byte,
	hasLastCommitHash bool,
	lastCommitHash [32]byte,
	hasDataHash bool,
	dataHash [32]byte,
	validatorsHash [32]byte,
	nextValidatorsHash [32]byte,
	consensusHash [32]byte,
	appHash []byte,
	hasLastResultsHash bool,
	lastResultsHash [32]byte,
	hasEvidenceHash bool,
	evidenceHash [32]byte,
	proposerAddress []byte,
) []byte {
	var fields [][]byte

	// Field 1: Version (always present) — use Solidity encoding (includes zero fields)
	versionBytes := encodeVersionSolidity(blockVersion, appVersion)
	fields = append(fields, versionBytes)

	// Field 2: ChainID (always present)
	fields = append(fields, encodeString(chainID))

	// Field 3: Height (always present)
	fields = append(fields, encodeVarint(height))

	// Field 4: Time (always present)
	fields = append(fields, encodeVarint(timeNanos))

	// Field 5: LastBlockId (conditional)
	if hasLastBlockId {
		blockIDBytes, _ := encodeBlockIDProto(lastBlockIdHash, pshTotal, pshHash)
		fields = append(fields, blockIDBytes)
	}

	// Field 6: LastCommitHash (conditional)
	if hasLastCommitHash {
		buf := make([]byte, 33)
		buf[0] = 32
		copy(buf[1:], lastCommitHash[:])
		fields = append(fields, buf)
	}

	// Field 7: DataHash (conditional)
	if hasDataHash {
		buf := make([]byte, 33)
		buf[0] = 32
		copy(buf[1:], dataHash[:])
		fields = append(fields, buf)
	}

	// Field 8: ValidatorsHash (always present)
	buf8 := make([]byte, 33)
	buf8[0] = 32
	copy(buf8[1:], validatorsHash[:])
	fields = append(fields, buf8)

	// Field 9: NextValidatorsHash (always present)
	buf9 := make([]byte, 33)
	buf9[0] = 32
	copy(buf9[1:], nextValidatorsHash[:])
	fields = append(fields, buf9)

	// Field 10: ConsensusHash (always present)
	buf10 := make([]byte, 33)
	buf10[0] = 32
	copy(buf10[1:], consensusHash[:])
	fields = append(fields, buf10)

	// Field 11: AppHash (always present, variable length)
	appHashField := append(encodeVarint(uint64(len(appHash))), appHash...)
	fields = append(fields, appHashField)

	// Field 12: LastResultsHash (conditional)
	if hasLastResultsHash {
		buf := make([]byte, 33)
		buf[0] = 32
		copy(buf[1:], lastResultsHash[:])
		fields = append(fields, buf)
	}

	// Field 13: EvidenceHash (conditional)
	if hasEvidenceHash {
		buf := make([]byte, 33)
		buf[0] = 32
		copy(buf[1:], evidenceHash[:])
		fields = append(fields, buf)
	}

	// Field 14: ProposerAddress (always present, variable length)
	proposerField := append(encodeVarint(uint64(len(proposerAddress))), proposerAddress...)
	fields = append(fields, proposerField)

	return solidityMerkleHash(fields)
}

// solidityMerkleHash matches Solidity's Header.merkleHash exactly.
// BUG: Solidity uses [0x00] and [0x01] which are uint256[1] = 32-byte values,
// not 1-byte prefixes as Tendermint spec requires.
func solidityMerkleHash(items [][]byte) []byte {
	if len(items) == 0 {
		return make([]byte, 32)
	}
	if len(items) == 1 {
		// tmhash(uint256(0) || leaf) — 32-byte zero prefix (Solidity's [0x00])
		prefix := make([]byte, 32) // 32 zero bytes
		h := sha256.Sum256(append(prefix, items[0]...))
		return h[:]
	}
	split := nextPowerOfTwo(len(items)) / 2
	left := solidityMerkleHash(items[:split])
	right := solidityMerkleHash(items[split:])
	// tmhash(uint256(1) || left || right) — 32-byte prefix (Solidity's [0x01])
	prefix := make([]byte, 32)
	prefix[31] = 0x01 // big-endian uint256(1)
	var data []byte
	data = append(data, prefix...)
	data = append(data, left...)
	data = append(data, right...)
	h := sha256.Sum256(data)
	return h[:]
}

func nextPowerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	if n&(n-1) == 0 {
		return n
	}
	power := 1
	for power < n {
		power <<= 1
	}
	return power
}

// hashValSet computes validator set hash matching Solidity's Header.hashValSet.
func hashValSet(validators [][32]byte, votingPowers []int64) []byte {
	encodedValidators := make([][]byte, len(validators))
	for i := range validators {
		encodedValidators[i] = encodeValidatorSolidity(validators[i], votingPowers[i])
	}
	return solidityMerkleHash(encodedValidators)
}

func printSection(name string) {
	fmt.Printf("\n=== %s ===\n", name)
}

func printHex(label string, data []byte) {
	fmt.Printf("  %s: 0x%s\n", label, hex.EncodeToString(data))
}

func main() {
	// ─── Test Data ───
	var pubKey [32]byte
	for i := range pubKey {
		pubKey[i] = byte(i + 0xAA)
	}

	var hash1 [32]byte
	for i := range hash1 {
		hash1[i] = byte(i + 0x10)
	}

	var hash2 [32]byte
	for i := range hash2 {
		hash2[i] = byte(i + 0x20)
	}

	var hash3 [32]byte
	for i := range hash3 {
		hash3[i] = byte(i + 0x30)
	}

	var hash4 [32]byte
	for i := range hash4 {
		hash4[i] = byte(i + 0x40)
	}

	var hash5 [32]byte
	for i := range hash5 {
		hash5[i] = byte(i + 0x50)
	}

	var hash6 [32]byte
	for i := range hash6 {
		hash6[i] = byte(i + 0x60)
	}

	var hash7 [32]byte
	for i := range hash7 {
		hash7[i] = byte(i + 0x70)
	}

	var pshHash [32]byte
	for i := range pshHash {
		pshHash[i] = byte(i + 0x80)
	}

	// ─── 1. Varint ───
	printSection("encodeVarint")
	varintTests := []uint64{0, 1, 127, 128, 255, 256, 300, 16384, 1000000}
	for _, v := range varintTests {
		encoded := encodeVarint(v)
		fmt.Printf("  input: %d → hex: 0x%s\n", v, hex.EncodeToString(encoded))
	}

	// ─── 2. encodeString ───
	printSection("encodeString")
	stringTests := []string{"", "cosmoshub-4", "test-chain", "a"}
	for _, s := range stringTests {
		encoded := encodeString(s)
		fmt.Printf("  input: %q → hex: 0x%s\n", s, hex.EncodeToString(encoded))
	}

	// ─── 3. encodeValidator (Solidity-style) ───
	printSection("encodeValidator (Solidity-style: raw pubkey)")
	votingPowers := []int64{100, 1, 0, 1000000}
	for _, vp := range votingPowers {
		encoded := encodeValidatorSolidity(pubKey, vp)
		printHex(fmt.Sprintf("pubKey=0x%s, votingPower=%d", hex.EncodeToString(pubKey[:4]), vp), encoded)
	}

	// ─── 4. encodeValidator (Go proto) ───
	printSection("encodeValidator (Go proto: wrapped PublicKey)")
	for _, vp := range votingPowers {
		encoded, err := encodeSimpleValidatorProto(pubKey, vp)
		if err != nil {
			fmt.Printf("  ERROR: %v\n", err)
			continue
		}
		printHex(fmt.Sprintf("pubKey=0x%s, votingPower=%d", hex.EncodeToString(pubKey[:4]), vp), encoded)
	}
	fmt.Println("  NOTE: Go proto wraps pubkey in PublicKey{Ed25519: ...} oneof,")
	fmt.Println("        so bytes DIFFER from Solidity's raw 32-byte encoding.")
	fmt.Println("        Solidity's encoding is a 'SimpleValidator' without the crypto.PublicKey wrapper.")

	// ─── 5. encodeVersion ───
	printSection("encodeVersion")
	versionTests := [][2]uint64{{11, 0}, {11, 2}, {0, 0}, {1, 1}}
	for _, vt := range versionTests {
		encoded, err := encodeVersionProto(vt[0], vt[1])
		if err != nil {
			fmt.Printf("  ERROR: %v\n", err)
			continue
		}
		fmt.Printf("  block=%d, app=%d → hex: 0x%s\n", vt[0], vt[1], hex.EncodeToString(encoded))
	}

	// ─── 6. encodePartSetHeader ───
	printSection("encodePartSetHeader")
	encoded, err := encodePartSetHeaderProto(1, hash1)
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	} else {
		fmt.Printf("  total=1, hash=0x%s...\n", hex.EncodeToString(hash1[:4]))
		printHex("result", encoded)
	}

	// ─── 7. encodeBlockId ───
	printSection("encodeBlockId")
	encoded, err = encodeBlockIDProto(hash2, 1, hash1)
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	} else {
		fmt.Printf("  hash=0x%s..., psh.total=1, psh.hash=0x%s...\n",
			hex.EncodeToString(hash2[:4]), hex.EncodeToString(hash1[:4]))
		printHex("result", encoded)
	}

	// ─── 8. hashValSet ───
	printSection("hashValSet (Merkle root of encoded validators)")
	var pubKey2 [32]byte
	for i := range pubKey2 {
		pubKey2[i] = byte(i + 0xBB)
	}
	valPubkeys := [][32]byte{pubKey, pubKey2}
	valPowers := []int64{100, 200}
	valSetHash := hashValSet(valPubkeys, valPowers)
	fmt.Printf("  validators: 2 validators\n")
	printHex("valSetHash", valSetHash)
	fmt.Println("  Individual validator encodings:")
	for i, pk := range valPubkeys {
		enc := encodeValidatorSolidity(pk, valPowers[i])
		printHex(fmt.Sprintf("  validator[%d]", i), enc)
	}

	// ─── 9. hashHeader ───
	printSection("hashHeader (full header Merkle root)")
	proposerAddr := make([]byte, 20)
	for i := range proposerAddr {
		proposerAddr[i] = byte(i + 0xF0)
	}
	appHash := make([]byte, 32)
	for i := range appHash {
		appHash[i] = byte(i + 0xA0)
	}

	headerHash := hashHeaderFields(
		11, 0, // version: block=11, app=0
		"cosmoshub-4",     // chainID
		12345,             // height
		1700000000,        // time (unix nanos simplified)
		true,              // hasLastBlockId
		hash2, 1, hash1,   // lastBlockId
		true, hash3,       // hasLastCommitHash
		true, hash4,       // hasDataHash
		hash5,             // validatorsHash
		hash6,             // nextValidatorsHash
		hash7,             // consensusHash
		appHash,           // appHash
		true, hash1,       // hasLastResultsHash
		true, hash2,       // hasEvidenceHash
		proposerAddr,      // proposerAddress
	)
	printHex("headerHash", headerHash)

	// Print all inputs for Solidity test reproduction
	printSection("Solidity Test Inputs")
	fmt.Println("  Use these values in Foundry test to call Encode.sol functions")
	fmt.Println("  and compare hex output with the Go results above.")
	fmt.Printf("  pubKey:           0x%s\n", hex.EncodeToString(pubKey[:]))
	fmt.Printf("  hash1:            0x%s\n", hex.EncodeToString(hash1[:]))
	fmt.Printf("  hash2:            0x%s\n", hex.EncodeToString(hash2[:]))
	fmt.Printf("  proposerAddress:  0x%s\n", hex.EncodeToString(proposerAddr))
	fmt.Printf("  appHash:          0x%s\n", hex.EncodeToString(appHash))

	// ─── 10. Print raw protobuf field-by-field for header ───
	printSection("Header field-by-field encoding (for Solidity hashHeader comparison)")

	versionEncProto, _ := encodeVersionProto(11, 0)
	printHex("field[0] version(11,0) Go proto", versionEncProto)
	versionEncSol := encodeVersionSolidity(11, 0)
	printHex("field[0] version(11,0) Solidity", versionEncSol)

	chainIDEnc := encodeString("cosmoshub-4")
	printHex("field[1] chainId", chainIDEnc)

	heightEnc := encodeVarint(12345)
	printHex("field[2] height=12345", heightEnc)

	timeEnc := encodeVarint(1700000000)
	printHex("field[3] time=1700000000", timeEnc)

	blockIDEnc, _ := encodeBlockIDProto(hash2, 1, hash1)
	printHex("field[4] lastBlockId", blockIDEnc)

	buf := make([]byte, 33)
	buf[0] = 32
	copy(buf[1:], hash3[:])
	printHex("field[5] lastCommitHash", buf)

	buf = make([]byte, 33)
	buf[0] = 32
	copy(buf[1:], hash4[:])
	printHex("field[6] dataHash", buf)

	buf = make([]byte, 33)
	buf[0] = 32
	copy(buf[1:], hash5[:])
	printHex("field[7] validatorsHash", buf)

	buf = make([]byte, 33)
	buf[0] = 32
	copy(buf[1:], hash6[:])
	printHex("field[8] nextValidatorsHash", buf)

	buf = make([]byte, 33)
	buf[0] = 32
	copy(buf[1:], hash7[:])
	printHex("field[9] consensusHash", buf)

	appHashField := append(encodeVarint(uint64(len(appHash))), appHash...)
	printHex("field[10] appHash", appHashField)

	buf = make([]byte, 33)
	buf[0] = 32
	copy(buf[1:], hash1[:])
	printHex("field[11] lastResultsHash", buf)

	buf = make([]byte, 33)
	buf[0] = 32
	copy(buf[1:], hash2[:])
	printHex("field[12] evidenceHash", buf)

	proposerField := append(encodeVarint(uint64(len(proposerAddr))), proposerAddr...)
	printHex("field[13] proposerAddress", proposerField)

	// ─── 11. Compare with CometBFT's correct merkle hash ───
	printSection("CometBFT correct merkle (1-byte prefix) vs Solidity (32-byte prefix)")
	singleItem := [][]byte{[]byte{0xde, 0xad, 0xbe, 0xef}}
	correctHash := merkle.HashFromByteSlices(singleItem)
	solidityHash := solidityMerkleHash(singleItem)
	printHex("CometBFT merkle(deadbeef)", correctHash)
	printHex("Solidity merkle(deadbeef)", solidityHash)
	fmt.Println("  These DIFFER because Solidity uses [0x00] (32-byte uint256) as leaf prefix,")
	fmt.Println("  while Tendermint spec uses 0x00 (1 byte). This is a BUG in Header.sol.")

	// ─── 12. Binary representation helpers ───
	printSection("uint64 big-endian encoding (for Solidity bytes32 comparison)")
	vals := []uint64{12345, 1700000000, 100}
	for _, v := range vals {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, v)
		fmt.Printf("  %d → BE hex: 0x%s\n", v, hex.EncodeToString(buf))
	}
}
