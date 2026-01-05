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
}
