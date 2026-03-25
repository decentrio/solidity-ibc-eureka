pragma solidity ^0.8.0;

import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { IICS02ClientMsgs } from "../msgs/IICS02ClientMsgs.sol";
library Encode {
    function encodeVarint(uint256 value) public pure returns (bytes memory) {
        if (value < 128) {
            return abi.encodePacked(uint8(value));
        }
        
        bytes memory result;
        while (value >= 128) {
            result = abi.encodePacked(result, uint8((value & 0x7F) | 0x80));
            value >>= 7;
        }
        result = abi.encodePacked(result, uint8(value));
        return result;
    }

    function encodeString(string memory value) public pure returns (bytes memory) {
        bytes memory valueBytes = bytes(value);
        uint256 length = valueBytes.length;
        
        // Encode length as varint
        bytes memory lengthBytes = encodeVarint(length);

        // Concatenate length and value
        return abi.encodePacked(lengthBytes, valueBytes);
    }

    /// @notice Encodes a nanosecond timestamp as a protobuf google.protobuf.Timestamp.
    /// @dev Omits fields with zero value per proto3 rules.
    function encodeTimestamp(uint128 nanos) public pure returns (bytes memory) {
        uint128 secs = nanos / 1_000_000_000;
        uint128 ns = nanos % 1_000_000_000;
        bytes memory encoded = new bytes(0);

        // Field 1: seconds (tag = 1, wire type = 0 for varint)
        if (secs > 0) {
            encoded = abi.encodePacked(encoded, uint8(0x08)); // tag: (1 << 3) | 0
            encoded = abi.encodePacked(encoded, encodeVarint(uint256(secs)));
        }
        // Field 2: nanos (tag = 2, wire type = 0 for varint)
        if (ns > 0) {
            encoded = abi.encodePacked(encoded, uint8(0x10)); // tag: (2 << 3) | 0
            encoded = abi.encodePacked(encoded, encodeVarint(uint256(ns)));
        }

        return encoded;
    }

    function encodeValidator(
        IICS07TendermintMsgs.SimpleValidator memory validator
    ) public pure returns (bytes memory) {
        bytes memory encoded = new bytes(0);
        
        // Encode field 1: pub_key (tag = 1, wire type = 2 for length-delimited)
        // Tag byte: (field_number << 3) | wire_type = (1 << 3) | 2 = 0x0A
        encoded = abi.encodePacked(encoded, uint8(0x0A)); // tag
        encoded = abi.encodePacked(encoded, uint8(32));   // length (32 bytes)
        encoded = abi.encodePacked(encoded, validator.pubKey); // data
        
        // Encode field 2: voting_power (tag = 2, wire type = 0 for varint)
        // Tag byte: (field_number << 3) | wire_type = (2 << 3) | 0 = 0x10
        encoded = abi.encodePacked(encoded, uint8(0x10)); // tag
        encoded = abi.encodePacked(encoded, encodeVarint(uint256(validator.votingPower))); // varint-encoded value

        return encoded;
    }

    function encodeVersion(IICS07TendermintMsgs.Version memory version) public pure returns (bytes memory) {
        bytes memory encoded = new bytes(0);
        
        // Field 1: blockVersion (tag = 1, wire type = 0 for varint)
        encoded = abi.encodePacked(encoded, uint8(0x08)); // tag: (1 << 3) | 0
        encoded = abi.encodePacked(encoded, encodeVarint(uint256(version.blockVersion)));
        
        // Field 2: appVersion (tag = 2, wire type = 0 for varint)
        encoded = abi.encodePacked(encoded, uint8(0x10)); // tag: (2 << 3) | 0
        encoded = abi.encodePacked(encoded, encodeVarint(uint256(version.appVersion)));
        
        return encoded;
    }

    function encodeBlockId(IICS07TendermintMsgs.BlockId memory blockId) public pure returns (bytes memory) {
        bytes memory encoded = new bytes(0);
        
        // Field 1: hashData (tag = 1, wire type = 2 for bytes)
        encoded = abi.encodePacked(encoded, uint8(0x0A)); // tag: (1 << 3) | 2
        encoded = abi.encodePacked(encoded, uint8(32)); // 32 bytes length
        encoded = abi.encodePacked(encoded, blockId.hashData);
        
        // Field 2: partSetHeader (tag = 2, wire type = 2 for message)
        bytes memory partSetHeaderEncoded = encodePartSetHeader(blockId.partSetHeader);
        encoded = abi.encodePacked(encoded, uint8(0x12)); // tag: (2 << 3) | 2
        encoded = abi.encodePacked(encoded, encodeVarint(partSetHeaderEncoded.length));
        encoded = abi.encodePacked(encoded, partSetHeaderEncoded);
        
        return encoded;
    }

    function encodePartSetHeader(IICS07TendermintMsgs.PartSetHeader memory partSetHeader) public pure returns (bytes memory) {
        bytes memory encoded = new bytes(0);
        
        // Field 1: total (tag = 1, wire type = 0 for varint)
        encoded = abi.encodePacked(encoded, uint8(0x08)); // tag: (1 << 3) | 0
        encoded = abi.encodePacked(encoded, encodeVarint(uint256(partSetHeader.total)));
        
        // Field 2: hashData (tag = 2, wire type = 2 for bytes)
        encoded = abi.encodePacked(encoded, uint8(0x12)); // tag: (2 << 3) | 2
        encoded = abi.encodePacked(encoded, uint8(32)); // 32 bytes length
        encoded = abi.encodePacked(encoded, partSetHeader.hashData);
        
        return encoded;
    }

    /// @notice Encodes a CanonicalVote as protobuf bytes (equivalent to CometBFT's VoteSignBytes).
    /// @param commit The block commit.
    /// @param chainId The chain ID string.
    /// @param valIdx The validator index into commitSigs.
    /// @return The protobuf-encoded canonical vote bytes.
    function voteSignBytes(
        IICS07TendermintMsgs.BlockCommit memory commit,
        string memory chainId,
        uint32 valIdx
    ) public pure returns (bytes memory) {
        IICS07TendermintMsgs.CommitSig memory commitSig = commit.commitSigs[valIdx];

        bool useCommitBlockId = commitSig.flag == IICS07TendermintMsgs.CommitSigFlag.BLOCK_ID_FLAG_COMMIT;
        bytes memory encodedBlockId = useCommitBlockId ? encodeBlockId(commit.blockId) : new bytes(0);
        bytes memory encodedTimestamp = encodeTimestamp(commitSig.data.timestamp);
        bytes memory chainIdBytes = bytes(chainId);

        bytes memory encoded = new bytes(0);

        // Field 1: type = PrecommitType (2), varint, tag 0x08
        encoded = abi.encodePacked(encoded, uint8(0x08), uint8(0x02));

        // Field 2: height, sfixed64, tag 0x11 (omit if zero)
        if (commit.height > 0) {
            encoded = abi.encodePacked(encoded, uint8(0x11));
            encoded = abi.encodePacked(encoded, encodeVarint(uint256(commit.height)));
        }

        // Field 3: round, sfixed64, tag 0x19 (omit if zero)
        if (commit.round > 0) {
            encoded = abi.encodePacked(encoded, uint8(0x19));
            encoded = abi.encodePacked(encoded, encodeVarint(uint256(commit.round)));
        }

        // Field 4: block_id, length-delimited, tag 0x22
        if (encodedBlockId.length > 0) {
            encoded = abi.encodePacked(encoded, uint8(0x22));
            encoded = abi.encodePacked(encoded, encodeVarint(encodedBlockId.length));
            encoded = abi.encodePacked(encoded, encodedBlockId);
        }

        // Field 5: timestamp, length-delimited, tag 0x2a
        if (encodedTimestamp.length > 0) {
            encoded = abi.encodePacked(encoded, uint8(0x2A));
            encoded = abi.encodePacked(encoded, encodeVarint(encodedTimestamp.length));
            encoded = abi.encodePacked(encoded, encodedTimestamp);
        }

        // Field 6: chain_id, length-delimited string, tag 0x32
        if (chainIdBytes.length > 0) {
            encoded = abi.encodePacked(encoded, uint8(0x32));
            encoded = abi.encodePacked(encoded, encodeString(chainId));
        }

        return encoded;
    }
}
