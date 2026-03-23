package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"testing"
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

// Cross-validation test: calls Solidity functions via anvil + cast call,
// compares output with Go proto.Marshal(). No hardcoded expected values.
//
// Prerequisites: anvil, forge, cast (Foundry toolchain)
//
// Run: cd operator && go test -v -run TestCrossValidate ./cmd/encode_debug/

const (
	anvilHost    = "127.0.0.1"
	anvilPort    = "18545"
	anvilRPCURL  = "http://127.0.0.1:18545"
	anvilPrivKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // anvil default
)

// ensureFoundryPath adds ~/.foundry/bin to PATH if Foundry tools aren't found.
func ensureFoundryPath() {
	if _, err := exec.LookPath("anvil"); err == nil {
		return
	}
	home, _ := os.UserHomeDir()
	foundryBin := home + "/.foundry/bin"
	if _, err := os.Stat(foundryBin + "/anvil"); err == nil {
		os.Setenv("PATH", foundryBin+":"+os.Getenv("PATH"))
	}
}

// waitForAnvil polls until anvil is accepting TCP connections.
func waitForAnvil(t *testing.T, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", anvilHost+":"+anvilPort, 200*time.Millisecond)
		if err == nil {
			conn.Close()
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatalf("anvil did not start within %s", timeout)
}

// forgeCreate deploys a contract and returns its address.
func forgeCreate(t *testing.T, rootDir string, contract string, extraArgs ...string) string {
	t.Helper()
	args := []string{"create",
		"--rpc-url", anvilRPCURL,
		"--private-key", anvilPrivKey,
		"--broadcast",
	}
	args = append(args, extraArgs...)
	args = append(args, contract)
	cmd := exec.Command("forge", args...)
	cmd.Dir = rootDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("forge create %s failed: %v\n%s", contract, err, out)
	}
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Deployed to:") {
			addr := strings.TrimSpace(strings.TrimPrefix(line, "Deployed to:"))
			return addr
		}
	}
	t.Fatalf("could not parse deployed address from forge output:\n%s", out)
	return ""
}

// deployWrapper deploys Encode, Header libraries then EncodeWrapper, returns wrapper address.
func deployWrapper(t *testing.T) string {
	t.Helper()
	rootDir := findRepoRoot(t)

	// Deploy libraries first
	encodeAddr := forgeCreate(t, rootDir, "contracts/utils/Encode.sol:Encode")
	t.Logf("Encode library deployed at %s", encodeAddr)

	headerAddr := forgeCreate(t, rootDir, "contracts/utils/Header.sol:Header",
		"--libraries", fmt.Sprintf("contracts/utils/Encode.sol:Encode:%s", encodeAddr),
	)
	t.Logf("Header library deployed at %s", headerAddr)

	// Deploy wrapper with library links
	wrapperAddr := forgeCreate(t, rootDir, "test/solidity-ibc/EncodeWrapper.sol:EncodeWrapper",
		"--libraries", fmt.Sprintf("contracts/utils/Encode.sol:Encode:%s", encodeAddr),
		"--libraries", fmt.Sprintf("contracts/utils/Header.sol:Header:%s", headerAddr),
	)
	t.Logf("EncodeWrapper deployed at %s", wrapperAddr)
	return wrapperAddr
}

// castCall invokes a read-only function via `cast call` and returns raw hex bytes.
func castCall(t *testing.T, addr, sig string, args ...string) []byte {
	t.Helper()
	cmdArgs := []string{"call", "--rpc-url", anvilRPCURL, addr, sig}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("cast", cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("cast call %s failed: %v\n%s", sig, err, out)
	}
	raw := strings.TrimSpace(string(out))
	return decodeCastBytes(t, raw)
}

// castCallBytes32 invokes a function returning bytes32 via `cast call`.
func castCallBytes32(t *testing.T, addr, sig string, args ...string) []byte {
	t.Helper()
	cmdArgs := []string{"call", "--rpc-url", anvilRPCURL, addr, sig}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("cast", cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("cast call %s failed: %v\n%s", sig, err, out)
	}
	raw := strings.TrimSpace(string(out))
	// bytes32 returns a single 32-byte hex value
	raw = strings.TrimPrefix(raw, "0x")
	b, err := hex.DecodeString(raw)
	if err != nil {
		t.Fatalf("failed to decode bytes32 hex %q: %v", raw, err)
	}
	return b
}

// decodeCastBytes decodes `bytes` return value from cast output.
// When cast call includes the return type hint like "(bytes)", it auto-decodes
// and returns just the raw hex content (e.g. "0x080b" or empty string for empty bytes).
func decodeCastBytes(t *testing.T, raw string) []byte {
	t.Helper()
	raw = strings.TrimPrefix(raw, "0x")
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return []byte{}
	}
	b, err := hex.DecodeString(raw)
	if err != nil {
		t.Fatalf("failed to decode hex %q: %v", raw, err)
	}
	return b
}

// findRepoRoot walks up from the current directory to find the repo root (has foundry.toml).
func findRepoRoot(t *testing.T) string {
	t.Helper()
	// The test runs from operator/cmd/encode_debug/, repo root is 3 levels up
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(dir + "/foundry.toml"); err == nil {
			return dir
		}
		parent := dir[:strings.LastIndex(dir, "/")]
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatal("could not find repo root (foundry.toml)")
	return ""
}

// cdcEncodeGo mirrors CometBFT's internal cdcEncode for header field hashing.
func cdcEncodeGo(item interface{}) []byte {
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

func TestCrossValidate(t *testing.T) {
	ensureFoundryPath()

	// Check tools are available
	for _, tool := range []string{"anvil", "forge", "cast"} {
		if _, err := exec.LookPath(tool); err != nil {
			t.Skipf("%s not found in PATH, skipping cross-validation test", tool)
		}
	}

	// Start anvil on a non-default port to avoid conflicts
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	anvil := exec.CommandContext(ctx, "anvil", "--port", anvilPort, "--silent")
	if err := anvil.Start(); err != nil {
		t.Fatalf("failed to start anvil: %v", err)
	}
	defer func() {
		cancel()
		_ = anvil.Wait()
	}()

	waitForAnvil(t, 10*time.Second)
	t.Log("anvil started")

	// Deploy EncodeWrapper
	addr := deployWrapper(t)

	// ─── Shared test data ───
	var pubKey [32]byte
	for i := range pubKey {
		pubKey[i] = byte(i + 0xAA)
	}
	pubKeyHex := "0x" + hex.EncodeToString(pubKey[:])

	var pubKey2 [32]byte
	for i := range pubKey2 {
		pubKey2[i] = byte(i + 0xBB)
	}
	pubKey2Hex := "0x" + hex.EncodeToString(pubKey2[:])

	var hash1 [32]byte
	for i := range hash1 {
		hash1[i] = byte(i + 0x10)
	}
	hash1Hex := "0x" + hex.EncodeToString(hash1[:])

	var hash2 [32]byte
	for i := range hash2 {
		hash2[i] = byte(i + 0x20)
	}
	hash2Hex := "0x" + hex.EncodeToString(hash2[:])

	// ═══════════════════════════════════════════════════
	// 1. encodeVersion
	// ═══════════════════════════════════════════════════
	t.Run("encodeVersion", func(t *testing.T) {
		tests := []struct {
			block, app uint64
		}{
			{11, 0}, {11, 2}, {1, 1}, {0, 0}, {0, 5},
		}
		for _, tc := range tests {
			name := fmt.Sprintf("block_%d_app_%d", tc.block, tc.app)
			t.Run(name, func(t *testing.T) {
				// Go proto.Marshal
				v := &cmtversion.Consensus{Block: tc.block, App: tc.app}
				goBytes, err := v.Marshal()
				if err != nil {
					t.Fatal(err)
				}

				// Solidity via cast call
				solBytes := castCall(t, addr,
					"encodeVersion(uint64,uint64)(bytes)",
					fmt.Sprintf("%d", tc.block),
					fmt.Sprintf("%d", tc.app),
				)

				compareBytes(t, "encodeVersion", goBytes, solBytes)
			})
		}
	})

	// ═══════════════════════════════════════════════════
	// 2. encodePartSetHeader
	// ═══════════════════════════════════════════════════
	t.Run("encodePartSetHeader", func(t *testing.T) {
		// Go proto.Marshal
		psh := &cmttypes.PartSetHeader{Total: 1, Hash: hash1[:]}
		goBytes, _ := psh.Marshal()

		// Solidity
		solBytes := castCall(t, addr,
			"encodePartSetHeader(uint32,bytes32)(bytes)",
			"1", hash1Hex,
		)

		compareBytes(t, "encodePartSetHeader", goBytes, solBytes)
	})

	// ═══════════════════════════════════════════════════
	// 3. encodeBlockId
	// ═══════════════════════════════════════════════════
	t.Run("encodeBlockId", func(t *testing.T) {
		// Go proto.Marshal
		bid := &cmttypes.BlockID{
			Hash:          hash2[:],
			PartSetHeader: cmttypes.PartSetHeader{Total: 1, Hash: hash1[:]},
		}
		goBytes, _ := bid.Marshal()

		// Solidity
		solBytes := castCall(t, addr,
			"encodeBlockId(bytes32,uint32,bytes32)(bytes)",
			hash2Hex, "1", hash1Hex,
		)

		compareBytes(t, "encodeBlockId", goBytes, solBytes)
	})

	// ═══════════════════════════════════════════════════
	// 4. encodeValidator
	// ═══════════════════════════════════════════════════
	t.Run("encodeValidator", func(t *testing.T) {
		votingPowers := []int64{100, 1, 0, 1000000}
		for _, vp := range votingPowers {
			t.Run(fmt.Sprintf("power_%d", vp), func(t *testing.T) {
				// Go proto.Marshal
				sv := &cmttypes.SimpleValidator{
					PubKey: &cmtcrypto.PublicKey{
						Sum: &cmtcrypto.PublicKey_Ed25519{Ed25519: pubKey[:]},
					},
					VotingPower: vp,
				}
				goBytes, _ := sv.Marshal()

				// Solidity
				solBytes := castCall(t, addr,
					"encodeValidator(bytes32,uint64)(bytes)",
					pubKeyHex,
					fmt.Sprintf("%d", vp),
				)

				compareBytes(t, "encodeValidator", goBytes, solBytes)
			})
		}
	})

	// ═══════════════════════════════════════════════════
	// 5. cdcEncodeString
	// ═══════════════════════════════════════════════════
	t.Run("cdcEncodeString", func(t *testing.T) {
		tests := []string{"cosmoshub-4", "", "a", "hello-world-chain-12345"}
		for _, s := range tests {
			name := s
			if name == "" {
				name = "empty"
			}
			t.Run(name, func(t *testing.T) {
				goBytes := cdcEncodeGo(s)
				if goBytes == nil {
					goBytes = []byte{}
				}

				solBytes := castCall(t, addr,
					"cdcEncodeString(string)(bytes)",
					s,
				)

				compareBytes(t, "cdcEncodeString", goBytes, solBytes)
			})
		}
	})

	// ═══════════════════════════════════════════════════
	// 6. cdcEncodeInt64
	// ═══════════════════════════════════════════════════
	t.Run("cdcEncodeInt64", func(t *testing.T) {
		tests := []int64{12345, 0, 1, 1700000000}
		for _, v := range tests {
			t.Run(fmt.Sprintf("%d", v), func(t *testing.T) {
				goBytes := cdcEncodeGo(v)
				if goBytes == nil {
					goBytes = []byte{}
				}

				solBytes := castCall(t, addr,
					"cdcEncodeInt64(uint256)(bytes)",
					fmt.Sprintf("%d", v),
				)

				compareBytes(t, "cdcEncodeInt64", goBytes, solBytes)
			})
		}
	})

	// ═══════════════════════════════════════════════════
	// 7. cdcEncodeBytes32
	// ═══════════════════════════════════════════════════
	t.Run("cdcEncodeBytes32", func(t *testing.T) {
		tests := []struct {
			name    string
			val     [32]byte
			isZero  bool // bytes32(0) in Solidity = "no value" = nil in Go
		}{
			{"hash1", hash1, false},
			{"zeros", [32]byte{}, true},
		}
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				var goBytes []byte
				if tc.isZero {
					// bytes32(0) represents "no value" in header hashing.
					// Solidity returns empty bytes; Go cdcEncode(nil) returns nil.
					goBytes = []byte{}
				} else {
					goBytes = cdcEncodeGo(tc.val[:])
					if goBytes == nil {
						goBytes = []byte{}
					}
				}

				valHex := "0x" + hex.EncodeToString(tc.val[:])
				solBytes := castCall(t, addr,
					"cdcEncodeBytes32(bytes32)(bytes)",
					valHex,
				)

				compareBytes(t, "cdcEncodeBytes32", goBytes, solBytes)
			})
		}
	})

	// ═══════════════════════════════════════════════════
	// 8. encodeTimestamp
	// ═══════════════════════════════════════════════════
	t.Run("encodeTimestamp", func(t *testing.T) {
		tests := []uint64{1700000000, 0, 1}
		for _, secs := range tests {
			t.Run(fmt.Sprintf("%d", secs), func(t *testing.T) {
				var goBytes []byte
				if secs == 0 {
					goBytes = []byte{}
				} else {
					ts, _ := gogotypes.StdTimeMarshal(time.Unix(int64(secs), 0))
					goBytes = ts
				}

				solBytes := castCall(t, addr,
					"encodeTimestamp(uint256)(bytes)",
					fmt.Sprintf("%d", secs),
				)

				compareBytes(t, "encodeTimestamp", goBytes, solBytes)
			})
		}
	})

	// ═══════════════════════════════════════════════════
	// 9. merkleHash
	// ═══════════════════════════════════════════════════
	t.Run("merkleHash", func(t *testing.T) {
		t.Run("single_item", func(t *testing.T) {
			items := [][]byte{{0xde, 0xad, 0xbe, 0xef}}
			goHash := merkle.HashFromByteSlices(items)

			// cast call with bytes array — encode as tuple
			solHash := castCallBytes32(t, addr,
				"merkleHash(bytes[])(bytes32)",
				"[0xdeadbeef]",
			)

			compareBytes(t, "merkleHash(single)", goHash, solHash)
		})

		t.Run("two_items", func(t *testing.T) {
			items := [][]byte{{0xaa}, {0xbb}}
			goHash := merkle.HashFromByteSlices(items)

			solHash := castCallBytes32(t, addr,
				"merkleHash(bytes[])(bytes32)",
				"[0xaa,0xbb]",
			)

			compareBytes(t, "merkleHash(two)", goHash, solHash)
		})

		t.Run("three_items", func(t *testing.T) {
			items := [][]byte{{0x01, 0x02}, {0x03, 0x04}, {0x05, 0x06}}
			goHash := merkle.HashFromByteSlices(items)

			solHash := castCallBytes32(t, addr,
				"merkleHash(bytes[])(bytes32)",
				"[0x0102,0x0304,0x0506]",
			)

			compareBytes(t, "merkleHash(three)", goHash, solHash)
		})
	})

	// ═══════════════════════════════════════════════════
	// 10. hashValSet
	// ═══════════════════════════════════════════════════
	t.Run("hashValSet", func(t *testing.T) {
		// Go: use types.Validator.Bytes() then merkle.HashFromByteSlices
		pk1 := ed25519.PubKey(pubKey[:])
		val1 := types.NewValidator(pk1, 100)
		pk2 := ed25519.PubKey(pubKey2[:])
		val2 := types.NewValidator(pk2, 200)
		goHash := merkle.HashFromByteSlices([][]byte{val1.Bytes(), val2.Bytes()})

		// Solidity
		solHash := castCallBytes32(t, addr,
			"hashValSet(bytes32[],uint64[])(bytes32)",
			fmt.Sprintf("[%s,%s]", pubKeyHex, pubKey2Hex),
			"[100,200]",
		)

		compareBytes(t, "hashValSet", goHash, solHash)
	})
}

func compareBytes(t *testing.T, label string, goBytes, solBytes []byte) {
	t.Helper()
	if !bytes.Equal(goBytes, solBytes) {
		t.Errorf("%s MISMATCH\n  Go:       0x%s\n  Solidity: 0x%s",
			label,
			hex.EncodeToString(goBytes),
			hex.EncodeToString(solBytes),
		)
	} else {
		t.Logf("%s OK → 0x%s", label, hex.EncodeToString(goBytes))
	}
}
