package main

import (
	"encoding/hex"
	"fmt"
	"time"

	cmtbytes "github.com/cometbft/cometbft/libs/bytes"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/crypto/merkle"
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	cmtversion "github.com/cometbft/cometbft/proto/tendermint/version"
	"github.com/cometbft/cometbft/types"
	gogotypes "github.com/cosmos/gogoproto/types"
)

// This program outputs Go proto.Marshal() reference hex for cross-validating
// Solidity's Encode.sol and Header.sol. Go proto is the SOURCE OF TRUTH.
// Divergences between Go output and Solidity output = bugs in Solidity.

func printSection(name string) {
	fmt.Printf("\n=== %s ===\n", name)
}

func printHex(label string, data []byte) {
	fmt.Printf("  %s: 0x%s\n", label, hex.EncodeToString(data))
}

// cdcEncode mirrors CometBFT's internal cdcEncode for header field hashing.
// It wraps values in protobuf wrapper types (StringValue, Int64Value, BytesValue).
func cdcEncode(item interface{}) []byte {
	if item == nil {
		return nil
	}
	switch v := item.(type) {
	case string:
		if v == "" {
			return nil
		}
		i := gogotypes.StringValue{Value: v}
		bz, _ := i.Marshal()
		return bz
	case int64:
		if v == 0 {
			return nil
		}
		i := gogotypes.Int64Value{Value: v}
		bz, _ := i.Marshal()
		return bz
	case cmtbytes.HexBytes:
		if len(v) == 0 {
			return nil
		}
		i := gogotypes.BytesValue{Value: v}
		bz, _ := i.Marshal()
		return bz
	case []byte:
		if len(v) == 0 {
			return nil
		}
		i := gogotypes.BytesValue{Value: v}
		bz, _ := i.Marshal()
		return bz
	default:
		return nil
	}
}

func main() {
	// ─── Test Data (same as EncodeTest.t.sol) ───
	var pubKey [32]byte
	for i := range pubKey {
		pubKey[i] = byte(i + 0xAA)
	}

	var pubKey2 [32]byte
	for i := range pubKey2 {
		pubKey2[i] = byte(i + 0xBB)
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

	proposerAddr := make([]byte, 20)
	for i := range proposerAddr {
		proposerAddr[i] = byte(i + 0xF0)
	}
	appHash := make([]byte, 32)
	for i := range appHash {
		appHash[i] = byte(i + 0xA0)
	}

	// ═══════════════════════════════════════════════════
	// 1. encodeVersion
	// ═══════════════════════════════════════════════════
	printSection("encodeVersion — proto.Marshal(cmtversion.Consensus)")
	versionTests := []struct{ block, app uint64 }{
		{11, 0}, {11, 2}, {1, 1}, {0, 0},
	}
	for _, vt := range versionTests {
		v := &cmtversion.Consensus{Block: vt.block, App: vt.app}
		encoded, _ := v.Marshal()
		fmt.Printf("  block=%d, app=%d → 0x%s\n", vt.block, vt.app, hex.EncodeToString(encoded))
	}
	fmt.Println("  ⚠ Solidity ALWAYS encodes zero fields (e.g. appVersion=0 → 0x1000)")
	fmt.Println("    Go proto.Marshal SKIPS zero fields")

	// ═══════════════════════════════════════════════════
	// 2. encodePartSetHeader
	// ═══════════════════════════════════════════════════
	printSection("encodePartSetHeader — proto.Marshal(cmttypes.PartSetHeader)")
	psh := &cmttypes.PartSetHeader{Total: 1, Hash: hash1[:]}
	pshBytes, _ := psh.Marshal()
	printHex("total=1, hash=hash1", pshBytes)

	// ═══════════════════════════════════════════════════
	// 3. encodeBlockId
	// ═══════════════════════════════════════════════════
	printSection("encodeBlockId — proto.Marshal(cmttypes.BlockID)")
	bid := &cmttypes.BlockID{
		Hash:          hash2[:],
		PartSetHeader: cmttypes.PartSetHeader{Total: 1, Hash: hash1[:]},
	}
	bidBytes, _ := bid.Marshal()
	printHex("hash=hash2, psh={1, hash1}", bidBytes)

	// ═══════════════════════════════════════════════════
	// 4. encodeValidator (SimpleValidator proto)
	// ═══════════════════════════════════════════════════
	printSection("encodeValidator — proto.Marshal(cmttypes.SimpleValidator)")
	fmt.Println("  Go SimpleValidator wraps pubKey in crypto.PublicKey{Ed25519: ...} oneof.")
	fmt.Println("  Solidity treats pubKey as raw 32 bytes (tag 0x0A + len 32 + raw data).")
	fmt.Println()

	votingPowers := []int64{100, 1, 0, 1000000}
	for _, vp := range votingPowers {
		sv := &cmttypes.SimpleValidator{
			PubKey: &cmtcrypto.PublicKey{
				Sum: &cmtcrypto.PublicKey_Ed25519{Ed25519: pubKey[:]},
			},
			VotingPower: vp,
		}
		svBytes, _ := sv.Marshal()
		printHex(fmt.Sprintf("votingPower=%d", vp), svBytes)
	}

	// Also show what Validator.Bytes() produces (used for ValSet hashing)
	printSection("Validator.Bytes() — used for validator set merkle hash")
	pk1 := ed25519.PubKey(pubKey[:])
	val1 := types.NewValidator(pk1, 100)
	val1Bytes := val1.Bytes()
	printHex("val1 (power=100)", val1Bytes)

	pk2 := ed25519.PubKey(pubKey2[:])
	val2 := types.NewValidator(pk2, 200)
	val2Bytes := val2.Bytes()
	printHex("val2 (power=200)", val2Bytes)

	// ═══════════════════════════════════════════════════
	// 5. hashValSet — merkle.HashFromByteSlices
	// ═══════════════════════════════════════════════════
	printSection("hashValSet — merkle.HashFromByteSlices(validator.Bytes()...)")
	fmt.Println("  CometBFT uses 1-byte merkle prefixes (0x00 leaf, 0x01 inner)")
	fmt.Println("  Solidity uses [0x00]/[0x01] = uint256[1] = 32-byte prefix (BUG)")
	fmt.Println()
	valSetHash := merkle.HashFromByteSlices([][]byte{val1Bytes, val2Bytes})
	printHex("valSetHash", valSetHash)

	// ═══════════════════════════════════════════════════
	// 6. hashHeader — CometBFT types.Header.Hash()
	// ═══════════════════════════════════════════════════
	printSection("hashHeader — types.Header.Hash()")
	fmt.Println("  CometBFT Header.Hash() always hashes 14 fields using merkle.")
	fmt.Println("  Fields wrapped via cdcEncode: StringValue, Int64Value, BytesValue.")
	fmt.Println("  Solidity uses raw encoding (encodeString, encodeVarint, raw bytes).")
	fmt.Println("  Solidity conditionally skips fields (has* flags) → different field count!")
	fmt.Println()

	header := &types.Header{
		Version: cmtversion.Consensus{Block: 11, App: 0},
		ChainID: "cosmoshub-4",
		Height:  12345,
		Time:    time.Unix(1700000000, 0),
		LastBlockID: types.BlockID{
			Hash: hash2[:],
			PartSetHeader: types.PartSetHeader{Total: 1, Hash: hash1[:]},
		},
		LastCommitHash:     hash3[:],
		DataHash:           hash4[:],
		ValidatorsHash:     hash5[:],
		NextValidatorsHash: hash6[:],
		ConsensusHash:      hash7[:],
		AppHash:            appHash,
		LastResultsHash:    hash1[:],
		EvidenceHash:       hash2[:],
		ProposerAddress:    proposerAddr,
	}

	headerHash := header.Hash()
	printHex("headerHash", headerHash)

	// Print individual field encodings as CometBFT does them
	printSection("Header field-by-field encodings (CometBFT)")
	fmt.Println("  These are the exact bytes passed to merkle.HashFromByteSlices()")
	fmt.Println()

	versionBytes, _ := header.Version.Marshal()
	printHex("field[ 0] Version.Marshal()", versionBytes)

	chainIDBytes := cdcEncode(header.ChainID)
	printHex("field[ 1] cdcEncode(ChainID)", chainIDBytes)
	fmt.Println("    → Wrapped in StringValue{Value: chainID}  (NOT raw varint+string)")

	heightBytes := cdcEncode(header.Height)
	printHex("field[ 2] cdcEncode(Height)", heightBytes)
	fmt.Println("    → Wrapped in Int64Value{Value: height}  (NOT raw varint)")

	timeBytes, _ := gogotypes.StdTimeMarshal(header.Time)
	printHex("field[ 3] StdTimeMarshal(Time)", timeBytes)
	fmt.Println("    → google.protobuf.Timestamp{seconds, nanos}  (NOT raw varint)")

	lastBlockIDProto := header.LastBlockID.ToProto()
	lastBlockIDBytes, _ := lastBlockIDProto.Marshal()
	printHex("field[ 4] LastBlockID.Marshal()", lastBlockIDBytes)

	printHex("field[ 5] cdcEncode(LastCommitHash)", cdcEncode(header.LastCommitHash))
	fmt.Println("    → Wrapped in BytesValue{Value: hash}  (NOT uint8(32)+hash)")
	printHex("field[ 6] cdcEncode(DataHash)", cdcEncode(header.DataHash))
	printHex("field[ 7] cdcEncode(ValidatorsHash)", cdcEncode(header.ValidatorsHash))
	printHex("field[ 8] cdcEncode(NextValidatorsHash)", cdcEncode(header.NextValidatorsHash))
	printHex("field[ 9] cdcEncode(ConsensusHash)", cdcEncode(header.ConsensusHash))
	printHex("field[10] cdcEncode(AppHash)", cdcEncode(header.AppHash))
	printHex("field[11] cdcEncode(LastResultsHash)", cdcEncode(header.LastResultsHash))
	printHex("field[12] cdcEncode(EvidenceHash)", cdcEncode(header.EvidenceHash))
	printHex("field[13] cdcEncode(ProposerAddress)", cdcEncode(header.ProposerAddress))

	// ═══════════════════════════════════════════════════
	// 7. Merkle hash comparison
	// ═══════════════════════════════════════════════════
	printSection("merkleHash — CometBFT vs Solidity prefix comparison")
	singleItem := [][]byte{{0xde, 0xad, 0xbe, 0xef}}
	printHex("merkle([deadbeef]) 1-byte prefix", merkle.HashFromByteSlices(singleItem))

	twoItems := [][]byte{{0xaa}, {0xbb}}
	printHex("merkle([aa, bb]) 1-byte prefix", merkle.HashFromByteSlices(twoItems))

	// ═══════════════════════════════════════════════════
	// Summary
	// ═══════════════════════════════════════════════════
	printSection("DIVERGENCES: Solidity Encode.sol vs Go proto.Marshal")
	fmt.Println()
	fmt.Println("  1. encodeVersion: Solidity always encodes zero fields; Go proto skips them")
	fmt.Println("     Go:       Version{11,0} → 0x080b")
	fmt.Println("     Solidity: Version{11,0} → 0x080b1000")
	fmt.Println()
	fmt.Println("  2. encodeValidator: Solidity uses raw 32-byte pubKey")
	fmt.Println("     Go SimpleValidator wraps in PublicKey{Ed25519: ...} oneof message")
	fmt.Println("     Wire format completely different")
	fmt.Println()
	fmt.Println("  3. merkleHash: Solidity [0x00]/[0x01] = uint256[1] = 32-byte prefix")
	fmt.Println("     CometBFT uses 1-byte prefix (0x00/0x01)")
	fmt.Println("     → Different merkle roots for same inputs")
	fmt.Println()
	fmt.Println("  4. hashHeader field encoding: Solidity uses raw encoding")
	fmt.Println("     CometBFT wraps in StringValue/Int64Value/BytesValue/Timestamp")
	fmt.Println("     → Every field except Version and LastBlockID is encoded differently")
	fmt.Println()
	fmt.Println("  5. hashHeader field count: Solidity conditionally skips fields (has* flags)")
	fmt.Println("     CometBFT ALWAYS hashes 14 fields (nil/empty for missing values)")
	fmt.Println("     → Different number of merkle leaves → completely different root")
}
