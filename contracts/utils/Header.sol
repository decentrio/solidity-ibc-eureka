pragma solidity ^0.8.0;

import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { Encode } from "./Encode.sol";
library Header {

    function hashValSet(
        IICS07TendermintMsgs.ValidatorSet memory valset
    ) pure public returns (bytes32) {
        bytes[] memory validatorBytes = new bytes[](valset.validators.length);
        for (uint256 i = 0; i < valset.validators.length; i++) {
            bytes memory validatorHash = Encode.encodeValidator(
                IICS07TendermintMsgs.SimpleValidator({
                    pubKey: valset.validators[i].pubKey,
                    votingPower: valset.validators[i].votingPower
                })
            );
            validatorBytes[i] = validatorHash;

        }

        bytes32 root = merkleHash(validatorBytes);
        return root;
    }

    function hashHeader(
        IICS07TendermintMsgs.BlockHeader memory header
    ) pure public returns (bytes32) {
        uint256 fieldCount = 9;
        
        if (header.hasLastBlockId) fieldCount++;
        if (header.hasLastCommitHash) fieldCount++;
        if (header.hasDataHash) fieldCount++;
        if (header.hasLastResultsHash) fieldCount++;
        if (header.hasEvidenceHash) fieldCount++;
        
        bytes[] memory headerBytes = new bytes[](fieldCount);
        uint256 index = 0;
        
        // Always present fields
        headerBytes[index++] = Encode.encodeVersion(header.version);
        headerBytes[index++] = Encode.encodeString(header.chainId);
        headerBytes[index++] = Encode.encodeVarint(uint256(header.height));
        headerBytes[index++] = Encode.encodeVarint(uint256(header.time));

        // Conditional fields
        if (header.hasLastBlockId) {
            headerBytes[index++] = Encode.encodeBlockId(header.lastBlockId);
        }
        
        if (header.hasLastCommitHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.lastCommitHash);
        }
        
        if (header.hasDataHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.dataHash);
        }
        
        // Always present fields
        headerBytes[index++] = abi.encodePacked(uint8(32), header.validatorsHash);
        headerBytes[index++] = abi.encodePacked(uint8(32), header.nextValidatorsHash);
        headerBytes[index++] = abi.encodePacked(uint8(32), header.consensusHash);
        
        // appHash is always present (no has flag in struct)
        uint256 appHashLength = header.appHash.length;
        headerBytes[index++] = abi.encodePacked(Encode.encodeVarint(appHashLength), header.appHash);
        
        // Conditional fields
        if (header.hasLastResultsHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.lastResultsHash);
        }
        
        if (header.hasEvidenceHash) {
            headerBytes[index++] = abi.encodePacked(uint8(32), header.evidenceHash);
        }
        
        // proposerAddress is always present (no has flag in struct)
        uint256 proposerAddressLength = header.proposerAddress.length;
        headerBytes[index++] = abi.encodePacked(Encode.encodeVarint(proposerAddressLength), header.proposerAddress);

        return merkleHash(headerBytes);
    }

    function merkleHash(
        bytes[] memory bytesArray
    ) public pure returns (bytes32) {
        if (bytesArray.length == 0) {
            return bytes32(0);
        }

        // tmhash(0x00 || leaf)
        // Pre and post-conditions: the hasher is in the reset state
        // before and after calling this function.
        if (bytesArray.length == 1) {
            return sha256(abi.encodePacked([0x00], bytesArray[0]));
        }

        uint256 split = nextPowerOfTwo(bytesArray.length) / 2;
        bytes32 left = merkleHash(getSlice(bytesArray, 0, split));
        bytes32 right = merkleHash(getSlice(bytesArray, split, bytesArray.length));

        return sha256(abi.encodePacked([0x01], left, right));
    }

    function getSlice(bytes[] memory bytesArray, uint256 from, uint256 to)
        public pure returns (bytes[] memory result) {
        require(from <= to && to <= bytesArray.length, "Invalid range");
        
        uint256 length = to - from;
        result = new bytes[](length);
        
        assembly {
            let src := add(add(bytesArray, 0x20), mul(from, 0x20))
            let dst := add(result, 0x20)
            let size := mul(length, 0x20)
            
            for { let i := 0 } lt(i, size) { i := add(i, 0x20) } {
                mstore(add(dst, i), mload(add(src, i)))
            }
        }
    }

    function nextPowerOfTwo(uint256 n) public pure returns (uint256) {
        if (n == 0) return 1;
        
        // Handle the case where n is already a power of 2
        if (n & (n - 1) == 0) return n;
        
        // Find the next power of 2
        uint256 power = 1;
        while (power < n) {
            power <<= 1;
        }
        return power;
    }
}