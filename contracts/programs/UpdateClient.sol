// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import { IUpdateClientMsgs } from "../light-clients/msgs/IUpdateClientMsgs.sol";
import { IICS07TendermintMsgs } from "../light-clients/msgs/IICS07TendermintMsgs.sol";
import { IICS02ClientMsgs } from "../msgs/IICS02ClientMsgs.sol";
import { IUpdateClient } from "../interfaces/IUpdateClient.sol";
import { HeightCmp } from "../utils/HeightCmp.sol";

import { Strings } from "@openzeppelin-contracts/utils/Strings.sol";

contract UpdateClient is IUpdateClient {
    string clientId = "07-tendermint-0";

    error MismatchedRevisionHeight(
        uint64 expected,
        uint64 actual
    );
    error InvalidHeaderHeight(
        uint64 height
    );
    error ValSetHashMismatch(
        bytes32 expected,
        bytes32 actual
    );
    error HeaderChainIdMismatch(
        string expected,
        string actual
    );
    error FailedToVerifyHeader(
        string reason
    );

    function updateClient(
        IICS07TendermintMsgs.ClientState memory clientState,
        IICS07TendermintMsgs.ConsensusState memory trustedConsensusState,
        IUpdateClientMsgs.Header memory proposedHeader,
        uint128 time
    ) view external {
        IUpdateClientMsgs.ChainId memory chainId = getChainId(clientState.chainId);
        IUpdateClientMsgs.Options memory options = IUpdateClientMsgs.Options({
            trustThreshold: clientState.trustLevel,
            trustingPeriod: clientState.trustingPeriod,
            clockDrift: 15
        });

        IUpdateClientMsgs.ClientConsensusStatePath memory path = IUpdateClientMsgs.ClientConsensusStatePath({
            clientId: clientId,
            revisionNumber: proposedHeader.trustedHeight.revisionNumber,
            revisionHeight: proposedHeader.trustedHeight.revisionHeight
        });

        verifyHeader(
            proposedHeader,
            clientId,
            chainId,
            options,
            time,
            path,
            trustedConsensusState
        );
    }

    function verifyHeader(
        IUpdateClientMsgs.Header memory proposedHeader,
        string memory clientId,
        IUpdateClientMsgs.ChainId memory chainId,
        IUpdateClientMsgs.Options memory options,
        uint128 time,
        IUpdateClientMsgs.ClientConsensusStatePath memory path,
        IICS07TendermintMsgs.ConsensusState memory trustedConsensusState
    ) internal pure {
        // Checks that the header fields are valid.
        validateBasic(proposedHeader);

        // The tendermint-light-client crate though works on heights that are assumed
        // to have the same revision number. We ensure this here.
        verifyChainIdVersion(chainId, proposedHeader);

        // Delegate to tendermint-light-client, which contains the required checks
        // of the new header against the trusted consensus state.
        {
            bytes32 nextValSetHash = hashValSet(proposedHeader.trustedNextValidatorSet);
            if (nextValSetHash != trustedConsensusState.nextValidatorsHash) {
                revert FailedToVerifyHeader(
                    "trusted next validator set hash does not match hash stored on chain"
                );
            }

            IUpdateClientMsgs.TrustedBlockState memory trustedState = IUpdateClientMsgs.TrustedBlockState({
                chainId: chainId,
                headerTime: time,
                height: proposedHeader.trustedHeight.revisionHeight,
                nextValidators: proposedHeader.trustedNextValidatorSet,
                nextValidatorsHash: nextValSetHash
            });

            IUpdateClientMsgs.UnTrustedBlockState memory untrustedState = IUpdateClientMsgs.UnTrustedBlockState({
                signedHeader: proposedHeader.signedHeader,
                validators: proposedHeader.validatorSet
            });

            verifyUpdateHeader(
                untrustedState,
                trustedState,
                options,
                time
            );
        }


    }

    function verifyUpdateHeader(
        IUpdateClientMsgs.UnTrustedBlockState memory untrustedState,
        IUpdateClientMsgs.TrustedBlockState memory trustedState,
        IUpdateClientMsgs.Options memory options,
        uint128 time
    ) internal pure {
        verifyValSets(untrustedState);
        verifyAgainstTrusted(untrustedState, trustedState, options.trustingPeriod, time);
        /// Check that the untrusted header is from past.
        uint128 drifted = time + options.clockDrift;
        require(untrustedState.signedHeader.header.time < drifted, "invalid block: header is from the future");
        verifyCommitAgainstTrusted(untrustedState, trustedState, options);
    }

    function verifyValSets(
        IUpdateClientMsgs.UnTrustedBlockState memory untrustedState
    ) internal pure {
        // Ensure the header validator hashes match the given validators
        bytes32 valSetHash = hashValSet(untrustedState.validators);
        if (valSetHash != untrustedState.signedHeader.header.validatorsHash) {
            revert("invalid block: validator set hash mismatch");
        }

        // Ensure the header matches the commit
        // TODO:

        // Additional implementation specific validation
        validateCommit(untrustedState.signedHeader, untrustedState.validators);
    }

    function verifyAgainstTrusted(
        IUpdateClientMsgs.UnTrustedBlockState memory untrustedState,
        IUpdateClientMsgs.TrustedBlockState memory trustedState,
        uint64 trustingPeriod,
        uint128 time
    ) internal pure {
        // Ensure the latest trusted header hasn't expired
        if (time < trustedState.headerTime || time - trustedState.headerTime > trustingPeriod) {
            revert("invalid block: untrusted state is outside of trusting period");
        }

        // Check that the untrusted block is more recent than the trusted state
        require(untrustedState.signedHeader.header.time > trustedState.headerTime, "invalid block: non monotonic bft time");

        // Check that the chain-id of the untrusted block matches that of the trusted state
        require(keccak256(abi.encodePacked(untrustedState.signedHeader.header.chainId)) == keccak256(abi.encodePacked(trustedState.chainId.id)), "invalid block: chain-id mismatch");

        uint64 trustedNextHeight = trustedState.height + 1;

        if (untrustedState.signedHeader.header.height == trustedNextHeight) {
            // If the untrusted block is the very next block after the trusted block,
            // check that their (next) validator sets hashes match.
            require(untrustedState.signedHeader.header.validatorsHash == trustedState.nextValidatorsHash, "invalid block: next validator set hash mismatch");
        } else {
            // Otherwise, ensure that the untrusted block has a greater height than
            // the trusted block.
            require(untrustedState.signedHeader.header.height > trustedNextHeight, "invalid block: non increasing height");
        }
    }

    /// Verify that a) there is enough overlap between the validator sets of the
    /// trusted and untrusted blocks and b) more than 2/3 of the validators
    /// correctly committed the block.
    function verifyCommitAgainstTrusted(
        IUpdateClientMsgs.UnTrustedBlockState memory untrustedState,
        IUpdateClientMsgs.TrustedBlockState memory trustedState,
        IUpdateClientMsgs.Options memory options
    ) internal pure {
        // If the trusted validator set has changed we need to check if there’s
        // overlap between the old trusted set and the new untrested header in
        // addition to checking if the new set correctly signed the header.
        uint64 trustedNextHeight = trustedState.height + 1;
        bool needBoth = untrustedState.signedHeader.header.height != trustedNextHeight;

        if (needBoth) {
            // TODO: check_enough_trust_and_signers
        } else {
            /// Check that there is enough signers overlap between the given, untrusted
            /// validator set and the untrusted signed header.
            IICS07TendermintMsgs.TrustThreshold memory trustThreshold = IICS07TendermintMsgs.TrustThreshold({
                numerator: 2,
                denominator: 3
            });

            // TODO: check enough power
        }
    }

    function validateCommit(
        IUpdateClientMsgs.SignedHeader memory signedHeader,
        IUpdateClientMsgs.ValidatorSet memory validators
    ) internal pure {
        IUpdateClientMsgs.CommitSig[] memory commitSigs = signedHeader.commit.commitSigs;

        bool hasPresentSignature = false;
        for (uint256 i = 0; i < commitSigs.length; i++) {
            if (commitSigs[i].flag != IUpdateClientMsgs.CommitSigFlag.BLOCK_ID_FLAG_ABSENT) {
                hasPresentSignature = true;
            }
        }

        if (!hasPresentSignature) {
            revert("invalid commit: no present signatures");
        }

        if (commitSigs.length != validators.validators.length) {
            revert("invalid commit: number of signatures does not match number of validators");
        }

        for (uint256 i = 0; i < commitSigs.length; i++) {
            IUpdateClientMsgs.CommitSig memory sig = commitSigs[i];
            bytes memory validatorAddress;

            if (sig.flag == IUpdateClientMsgs.CommitSigFlag.BLOCK_ID_FLAG_ABSENT) {
                continue;
            } else {
                validatorAddress = sig.data.validatorAddress;
            }

            bool found = false;
            for (uint256 j = 0; j < validators.validators.length; j++) {
                if (keccak256(abi.encodePacked(validators.validators[j].valAddress)) == keccak256(abi.encodePacked(validatorAddress))) {
                    found = true;
                    break;
                }
            }
            if (!found) {
                revert("invalid commit: faulty signer");
            }
        }
    }

    function verifyChainIdVersion(
        IUpdateClientMsgs.ChainId memory chainId,
        IUpdateClientMsgs.Header memory header
    ) internal pure {
        IUpdateClientMsgs.ChainId memory headerChainId = getChainId(header.signedHeader.header.chainId);
        if (chainId.revisionNumber != headerChainId.revisionNumber) {
            revert HeaderChainIdMismatch(
                header.signedHeader.header.chainId,
                chainId.id
            );
        }
    }

    function validateBasic(
        IUpdateClientMsgs.Header memory header
    ) internal pure {
        IUpdateClientMsgs.ChainId memory chainId = getChainId(header.signedHeader.header.chainId);
        if (chainId.revisionNumber != header.trustedHeight.revisionNumber) {
            revert MismatchedRevisionHeight(
                chainId.revisionNumber,
                header.trustedHeight.revisionNumber
            );
        }

        IICS02ClientMsgs.Height memory height = IICS02ClientMsgs.Height({
            revisionNumber: chainId.revisionNumber,
            revisionHeight: header.signedHeader.header.height
        });

        // We need to ensure that the trusted height (representing the
        // height of the header already on chain for which this client update is
        // based on) must be smaller than height of the new header that we're
        // installing.
        if (HeightCmp.ge(header.trustedHeight, height)) {
            revert InvalidHeaderHeight(
                height.revisionHeight
            );
        }

        bytes32 valSetHash = hashValSet(header.validatorSet);
        if (valSetHash != header.signedHeader.header.validatorsHash) {
            revert ValSetHashMismatch(
                header.signedHeader.header.validatorsHash,
                valSetHash
            );
        }
    }


    function hashValSet(
        IUpdateClientMsgs.ValidatorSet memory valset
    ) pure internal returns (bytes32) {
        bytes[] memory validatorBytes = new bytes[](valset.validators.length);
        for (uint256 i = 0; i < valset.validators.length; i++) {
            bytes memory validatorHash = encodeToVec(
                IUpdateClientMsgs.SimpleValidator({
                    pubKey: valset.validators[i].pubKey,
                    votingPower: valset.validators[i].votingPower
                })
            );
            validatorBytes[i] = validatorHash;

        }

        bytes32 root = merkleHash(validatorBytes);
        return root;
    }

    function merkleHash(
        bytes[] memory bytesArray
    ) internal pure returns (bytes32) {
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

    function getChainId(string memory id) pure internal returns (IUpdateClientMsgs.ChainId memory) {
        // TODO: validate id
        bytes memory chainIdBytes = bytes(id);

        // Find the last occurrence of '-'
        int256 lastDashIndex = -1;
        for (uint256 i = chainIdBytes.length-1; i >= 0; i--) {
            if (chainIdBytes[i] == 0x2D) { // '-' character
                lastDashIndex = int256(i);
                break;
            }
        }

        if (lastDashIndex == -1) {
            if (1 <= chainIdBytes.length && chainIdBytes.length < 64) {
                return IUpdateClientMsgs.ChainId({
                    id: id, 
                    revisionNumber: 0
                });
            } else {
                revert("Invalid chain ID length");
            }
        }

        uint256 dashIndex = uint256(lastDashIndex);
        // Extract revision number string
        bytes memory revisionBytes = new bytes(chainIdBytes.length - dashIndex - 1);
        for (uint256 i = 0; i < revisionBytes.length; i++) {
            revisionBytes[i] = chainIdBytes[dashIndex + 1 + i];
        }

        // Validates the revision number not to start with leading zeros, like "01".
        // Zero is the only allowed revision number with leading zero.
        if (revisionBytes.length == 0 || (revisionBytes[0] == 0x30 && revisionBytes.length > 1)) {
            if (1 <= chainIdBytes.length && chainIdBytes.length < 64) {
                return IUpdateClientMsgs.ChainId({
                    id: id, 
                    revisionNumber: 0
                });
            } else {
                revert("Invalid chain ID length");
            }
        }

        (bool success, uint256 parsedRevisionNumber) = Strings.tryParseUint(string(revisionBytes));
        if (! success || parsedRevisionNumber > type(uint64).max) {
            if (1 <= chainIdBytes.length && chainIdBytes.length < 64) {
                return IUpdateClientMsgs.ChainId({
                    id: id, 
                    revisionNumber: 0
                });
            } else {
                revert("Invalid chain ID length");
            }
        }

        uint64 revisionNumber = uint64(parsedRevisionNumber);
        
        // Extract chain name
        bytes memory chainNameBytes = new bytes(dashIndex);
        for (uint256 i = 0; i < dashIndex; i++) {
            chainNameBytes[i] = chainIdBytes[i];
        }

        uint256 min = 1;
        // Prefix must be at most `max_id_length - 21` characters long since the
        // longest identifier we can construct is `{prefix}-{u64::MAX}` which
        // extends prefix by 21 characters.
        uint256 max = 43;
        if (chainNameBytes.length <= min || chainNameBytes.length >= max) {
            revert("Invalid chain prefix length");
        }

        return IUpdateClientMsgs.ChainId({
            id: id, 
            revisionNumber: revisionNumber
        });
    }

    function encodeVarint(uint256 value) internal pure returns (bytes memory) {
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

    function encodeToVec(
        IUpdateClientMsgs.SimpleValidator memory validator
    ) internal pure returns (bytes memory) {
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

}