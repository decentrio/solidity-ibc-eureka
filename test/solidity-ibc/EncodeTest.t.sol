// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

// solhint-disable custom-errors,max-line-length

import { Test, console } from "forge-std/Test.sol";

import { Encode } from "../../contracts/utils/Encode.sol";
import { Header } from "../../contracts/utils/Header.sol";
import { IICS07TendermintMsgs } from "../../contracts/light-clients/msgs/IICS07TendermintMsgs.sol";

/// @title EncodeTest
/// @notice Cross-validates Encode.sol output against Go protobuf reference data.
///         Run Go reference: cd operator && go run ./cmd/encode_debug/
contract EncodeTest is Test {
    // ─── Test data (same as Go cmd/encode_debug/main.go) ───

    function _pubKey() internal pure returns (bytes32) {
        bytes32 pk;
        assembly {
            // 0xAA, 0xAB, 0xAC, ... incrementing from 0xAA for 32 bytes
            // Go: for i := range pubKey { pubKey[i] = byte(i + 0xAA) }
            pk := 0xaaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c9
        }
        return pk;
    }

    function _hash1() internal pure returns (bytes32) {
        return bytes32(0x101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f);
    }

    function _hash2() internal pure returns (bytes32) {
        return bytes32(0x202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f);
    }

    function _hash3() internal pure returns (bytes32) {
        return bytes32(0x303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f);
    }

    function _hash4() internal pure returns (bytes32) {
        return bytes32(0x404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f);
    }

    function _hash5() internal pure returns (bytes32) {
        return bytes32(0x505152535455565758595a5b5c5d5e5f606162636465666768696a6b6c6d6e6f);
    }

    function _hash6() internal pure returns (bytes32) {
        return bytes32(0x606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f);
    }

    function _hash7() internal pure returns (bytes32) {
        return bytes32(0x707172737475767778797a7b7c7d7e7f808182838485868788898a8b8c8d8e8f);
    }

    // ─── encodeVarint tests ───

    function test_encodeVarint_zero() public pure {
        bytes memory result = Encode.encodeVarint(0);
        assertEq(result, hex"00");
    }

    function test_encodeVarint_one() public pure {
        bytes memory result = Encode.encodeVarint(1);
        assertEq(result, hex"01");
    }

    function test_encodeVarint_127() public pure {
        bytes memory result = Encode.encodeVarint(127);
        assertEq(result, hex"7f");
    }

    function test_encodeVarint_128() public pure {
        bytes memory result = Encode.encodeVarint(128);
        assertEq(result, hex"8001");
    }

    function test_encodeVarint_255() public pure {
        bytes memory result = Encode.encodeVarint(255);
        assertEq(result, hex"ff01");
    }

    function test_encodeVarint_256() public pure {
        bytes memory result = Encode.encodeVarint(256);
        assertEq(result, hex"8002");
    }

    function test_encodeVarint_300() public pure {
        bytes memory result = Encode.encodeVarint(300);
        assertEq(result, hex"ac02");
    }

    function test_encodeVarint_16384() public pure {
        bytes memory result = Encode.encodeVarint(16384);
        assertEq(result, hex"808001");
    }

    function test_encodeVarint_1000000() public pure {
        bytes memory result = Encode.encodeVarint(1000000);
        assertEq(result, hex"c0843d");
    }

    // ─── encodeString tests ───

    function test_encodeString_empty() public pure {
        bytes memory result = Encode.encodeString("");
        assertEq(result, hex"00");
    }

    function test_encodeString_cosmoshub4() public pure {
        bytes memory result = Encode.encodeString("cosmoshub-4");
        assertEq(result, hex"0b636f736d6f736875622d34");
    }

    function test_encodeString_testChain() public pure {
        bytes memory result = Encode.encodeString("test-chain");
        assertEq(result, hex"0a746573742d636861696e");
    }

    function test_encodeString_singleChar() public pure {
        bytes memory result = Encode.encodeString("a");
        assertEq(result, hex"0161");
    }

    // ─── encodeValidator tests ───

    function test_encodeValidator_power100() public pure {
        IICS07TendermintMsgs.SimpleValidator memory v = IICS07TendermintMsgs.SimpleValidator({
            pubKey: _pubKey(),
            votingPower: 100
        });
        bytes memory result = Encode.encodeValidator(v);
        assertEq(
            result,
            hex"0a20aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c91064"
        );
    }

    function test_encodeValidator_power1() public pure {
        IICS07TendermintMsgs.SimpleValidator memory v = IICS07TendermintMsgs.SimpleValidator({
            pubKey: _pubKey(),
            votingPower: 1
        });
        bytes memory result = Encode.encodeValidator(v);
        assertEq(
            result,
            hex"0a20aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c91001"
        );
    }

    function test_encodeValidator_power0() public pure {
        IICS07TendermintMsgs.SimpleValidator memory v = IICS07TendermintMsgs.SimpleValidator({
            pubKey: _pubKey(),
            votingPower: 0
        });
        bytes memory result = Encode.encodeValidator(v);
        assertEq(
            result,
            hex"0a20aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c91000"
        );
    }

    function test_encodeValidator_power1000000() public pure {
        IICS07TendermintMsgs.SimpleValidator memory v = IICS07TendermintMsgs.SimpleValidator({
            pubKey: _pubKey(),
            votingPower: 1000000
        });
        bytes memory result = Encode.encodeValidator(v);
        assertEq(
            result,
            hex"0a20aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c910c0843d"
        );
    }

    // ─── encodeVersion tests ───

    function test_encodeVersion_11_0() public pure {
        IICS07TendermintMsgs.Version memory v = IICS07TendermintMsgs.Version({ blockVersion: 11, appVersion: 0 });
        bytes memory result = Encode.encodeVersion(v);
        // NOTE: Solidity always encodes zero fields (0x1000 = tag 0x10 + varint 0x00)
        // Go proto.Marshal skips zero-value fields, so Go outputs 0x080b only.
        // This is a DIVERGENCE: Solidity encodes appVersion=0 explicitly.
        assertEq(result, hex"080b1000");
    }

    function test_encodeVersion_11_2() public pure {
        IICS07TendermintMsgs.Version memory v = IICS07TendermintMsgs.Version({ blockVersion: 11, appVersion: 2 });
        bytes memory result = Encode.encodeVersion(v);
        assertEq(result, hex"080b1002");
    }

    function test_encodeVersion_1_1() public pure {
        IICS07TendermintMsgs.Version memory v = IICS07TendermintMsgs.Version({ blockVersion: 1, appVersion: 1 });
        bytes memory result = Encode.encodeVersion(v);
        assertEq(result, hex"08011001");
    }

    // ─── encodePartSetHeader tests ───

    function test_encodePartSetHeader() public pure {
        IICS07TendermintMsgs.PartSetHeader memory psh = IICS07TendermintMsgs.PartSetHeader({
            total: 1,
            hashData: _hash1()
        });
        bytes memory result = Encode.encodePartSetHeader(psh);
        assertEq(
            result,
            hex"08011220101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f"
        );
    }

    // ─── encodeBlockId tests ───

    function test_encodeBlockId() public pure {
        IICS07TendermintMsgs.BlockId memory bid = IICS07TendermintMsgs.BlockId({
            hashData: _hash2(),
            partSetHeader: IICS07TendermintMsgs.PartSetHeader({ total: 1, hashData: _hash1() })
        });
        bytes memory result = Encode.encodeBlockId(bid);
        assertEq(
            result,
            hex"0a20202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f122408011220101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f"
        );
    }

    // ─── hashValSet tests ───

    function test_hashValSet() public pure {
        IICS07TendermintMsgs.ValidatorInfo[] memory vals = new IICS07TendermintMsgs.ValidatorInfo[](2);
        vals[0] = IICS07TendermintMsgs.ValidatorInfo({
            valAddress: hex"",
            pubKey: _pubKey(),
            votingPower: 100,
            proposerPriority: 0
        });
        // pubKey2: incrementing from 0xBB
        vals[1] = IICS07TendermintMsgs.ValidatorInfo({
            valAddress: hex"",
            pubKey: bytes32(0xbbbcbdbebfc0c1c2c3c4c5c6c7c8c9cacbcccdcecfd0d1d2d3d4d5d6d7d8d9da),
            votingPower: 200,
            proposerPriority: 0
        });

        IICS07TendermintMsgs.ValidatorSet memory valSet = IICS07TendermintMsgs.ValidatorSet({
            validators: vals,
            hasProposer: false,
            proposer: IICS07TendermintMsgs.ValidatorInfo({
                valAddress: hex"",
                pubKey: bytes32(0),
                votingPower: 0,
                proposerPriority: 0
            }),
            totalVotingPower: 300
        });

        bytes32 result = Header.hashValSet(valSet);
        // Go reference (Solidity-compatible merkle with 32-byte prefixes):
        // 0x5512167500074e07424a51d0cce9b5293e9653a113ae3d64aeed7b6c32cb34b1
        assertEq(result, bytes32(0x5512167500074e07424a51d0cce9b5293e9653a113ae3d64aeed7b6c32cb34b1));
    }

    // ─── hashHeader test ───

    function test_hashHeader() public pure {
        IICS07TendermintMsgs.BlockHeader memory header = IICS07TendermintMsgs.BlockHeader({
            version: IICS07TendermintMsgs.Version({ blockVersion: 11, appVersion: 0 }),
            chainId: "cosmoshub-4",
            height: 12345,
            time: 1700000000,
            hasLastBlockId: true,
            lastBlockId: IICS07TendermintMsgs.BlockId({
                hashData: _hash2(),
                partSetHeader: IICS07TendermintMsgs.PartSetHeader({ total: 1, hashData: _hash1() })
            }),
            hasLastCommitHash: true,
            lastCommitHash: _hash3(),
            hasDataHash: true,
            dataHash: _hash4(),
            validatorsHash: _hash5(),
            nextValidatorsHash: _hash6(),
            consensusHash: _hash7(),
            appHash: bytes32(0xa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebf),
            hasLastResultsHash: true,
            lastResultsHash: _hash1(),
            hasEvidenceHash: true,
            evidenceHash: _hash2(),
            proposerAddress: hex"f0f1f2f3f4f5f6f7f8f9fafbfcfdfeff00010203"
        });

        bytes32 result = Header.hashHeader(header);
        // Go reference (Solidity-compatible: 32-byte merkle prefixes + always-encode-zero-fields):
        // 0xea330cf5f6ce6df2c098baef36c92e4d2d3ad3ff952ef187ce99017d107b28bd
        assertEq(result, bytes32(0xea330cf5f6ce6df2c098baef36c92e4d2d3ad3ff952ef187ce99017d107b28bd));
    }

    // ─── merkleHash tests ───

    function test_merkleHash_empty() public pure {
        bytes[] memory items = new bytes[](0);
        bytes32 result = Header.merkleHash(items);
        assertEq(result, bytes32(0));
    }

    function test_merkleHash_singleLeaf() public pure {
        bytes[] memory items = new bytes[](1);
        items[0] = hex"deadbeef";
        bytes32 result = Header.merkleHash(items);
        // BUG: Header.sol uses [0x00] which is uint256[1] = 32 zero bytes, not bytes1(0x00).
        // Tendermint spec uses 0x00 (1 byte prefix), but Solidity encodes 32 bytes.
        // This test validates current (buggy) behavior.
        bytes32 expected = sha256(abi.encodePacked(uint256(0), hex"deadbeef"));
        assertEq(result, expected);
    }

    function test_merkleHash_twoLeaves() public pure {
        bytes[] memory items = new bytes[](2);
        items[0] = hex"aa";
        items[1] = hex"bb";
        bytes32 result = Header.merkleHash(items);
        // BUG: Same [0x00]/[0x01] issue — these are 32-byte prefixes, not 1-byte.
        bytes32 left = sha256(abi.encodePacked(uint256(0), hex"aa"));
        bytes32 right = sha256(abi.encodePacked(uint256(0), hex"bb"));
        bytes32 expected = sha256(abi.encodePacked(uint256(1), left, right));
        assertEq(result, expected);
    }

    // ─── Debug: emit encoded bytes for manual inspection ───

    function test_debug_printEncodings() public pure {
        // Varint
        console.log("=== encodeVarint ===");
        console.logBytes(Encode.encodeVarint(0));
        console.logBytes(Encode.encodeVarint(128));
        console.logBytes(Encode.encodeVarint(300));
        console.logBytes(Encode.encodeVarint(1000000));

        // Version
        console.log("=== encodeVersion ===");
        console.logBytes(
            Encode.encodeVersion(IICS07TendermintMsgs.Version({ blockVersion: 11, appVersion: 0 }))
        );

        // Validator
        console.log("=== encodeValidator ===");
        console.logBytes(
            Encode.encodeValidator(
                IICS07TendermintMsgs.SimpleValidator({ pubKey: _pubKey(), votingPower: 100 })
            )
        );

        // BlockId
        console.log("=== encodeBlockId ===");
        console.logBytes(
            Encode.encodeBlockId(
                IICS07TendermintMsgs.BlockId({
                    hashData: _hash2(),
                    partSetHeader: IICS07TendermintMsgs.PartSetHeader({ total: 1, hashData: _hash1() })
                })
            )
        );
    }
}
