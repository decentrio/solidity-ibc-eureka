package main

import (
	"context"
	"crypto/ed25519"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"0x5ea000000/ecip-gnark/signature/eddsa"
	"0x5ea000000/ecip-gnark/utils"

	"filippo.io/edwards25519"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/consensys/gnark-crypto/ecc"
	groth16 "github.com/consensys/gnark/backend/groth16"
	groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/cosmos/ibc-go/v10/modules/core/02-client/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	tendermintContract "prover/bindings/SP1ICS07Tendermint"
	tendermintClient "prover/client"
	"prover/keys"
	"prover/runner"
)

const (
	flagOnlyOnce       = "only-once"
	flagProofType      = "proof-type"
	flagOutput         = "output"
	flagOutputPath     = "output-path"
	flagTrustLevel     = "trust-level"
	flagTrustingPeriod = "trusting-period"
	flagTrustedBlock   = "trusted-block"
	flagMembership     = "membership"
	flagTargetBlock    = "target-block"
	flagProvingKey     = "proving-key"
)

type Fp25519 = emulated.Curve25519Fp
type Fr25519 = emulated.Curve25519Fr

type PreHashCircuit[Base, Scalars emulated.FieldParams] struct {
	Sig eddsa.Signature[Base, Scalars] `gnark:",public"`
	// Msg  emulated.Element[Scalars]      `gnark:",public"`
	Hash emulated.Element[Scalars]      `gnark:",public"`
	Pub  eddsa.PublicKey[Base, Scalars] `gnark:",public"`
}

func (c *PreHashCircuit[Base, Scalars]) Define(api frontend.API) error {

	//A, _ := new(big.Int).SetString("0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec", 0)
	//D, _ := new(big.Int).SetString("0x52036cee2b6ffe738cc740797779e89800700a4d4141d8ab75eb4dca135978a3", 0)
	//Gx, _ := new(big.Int).SetString("0x216936d3cd6e53fec0a4e231fdd6dc5c692cc7609525a7b2c9562d608f25d51a", 0)
	//Gy, _ := new(big.Int).SetString("0x6666666666666666666666666666666666666666666666666666666666666658", 0)

	config := eddsa.Config{
		Hasher:  nil,
		FromWei: false,
	}

	err := eddsa.Verify[Base, Scalars](api, c.Sig, c.Hash, c.Pub, config)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	zLogger, _ := zap.NewProduction(zap.AddStacktrace(zap.DPanicLevel))
	defer zLogger.Sync()

	logger := zLogger.Sugar()

	rootCmd := &cobra.Command{
		Use:   "prover [command]",
		Short: "command for prover",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(
		Start(zLogger),
		Genesis(zLogger),
		Fixtures(zLogger),
	)

	err := rootCmd.Execute()
	if err != nil {
		logger.Fatal(err)
	}
}

func Start(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO:
			return nil
		},
	}
	cmd.Flags().Bool(flagOnlyOnce, false, "run only once")
	// cmd.Flags().String(flagConfigPath, ".prover/config.yaml", "the path to your prover priv and pub key")
	return cmd
}

func Genesis(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genesis",
		Short: "genesis",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := godotenv.Load()
			if err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}

			trustedBlock, err := cmd.Flags().GetInt64(flagTrustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted block: %w", err)
			}
			trustingPeriod, err := cmd.Flags().GetUint32(flagTrustingPeriod)
			if err != nil {
				return fmt.Errorf("failed to get trusting period: %w", err)
			}
			trustLevel, err := cmd.Flags().GetString(flagTrustLevel)
			if err != nil {
				return fmt.Errorf("failed to get trust level from flag: %w", err)
			}
			proofType, err := cmd.Flags().GetString(flagProofType)
			if err != nil {
				return fmt.Errorf("failed to get proof type from flag: %w", err)
			}

			genesis, err := tendermintClient.GetGenesis(trustedBlock, trustingPeriod, trustLevel, proofType)
			if err != nil {
				return fmt.Errorf("failed to get genesis: %w", err)
			}

			outputType, err := cmd.Flags().GetString(flagOutput)
			if err != nil {
				return fmt.Errorf("failed to get output type from flag: %w", err)
			}

			data, err := json.Marshal(genesis)
			if err != nil {
				return fmt.Errorf("failed to marshal genesis state: %w", err)
			}

			switch outputType {
			case "json":
				fmt.Println(string(data))
			case "file":
				outputDir, err := cmd.Flags().GetString(flagOutputPath)
				if err != nil {
					return fmt.Errorf("failed to get output path from flag: %w", err)
				}

				if err := os.WriteFile(outputDir, data, 0644); err != nil {
					return fmt.Errorf("failed to write genesis state to file: %w", err)
				}
			default:
				return fmt.Errorf("unsupported output type: %s, supported types are: json, file", outputType)
			}

			return nil
		},
	}
	cmd.Flags().String(flagProofType, "groth16", "the type of proof to use (groth16, plonk)")
	cmd.Flags().Int64(flagTrustedBlock, 0, "the trusted block height, if <height> is 0 then catch latest block")
	cmd.Flags().String(flagOutput, "json", "the output structure for the genesis state (json, file)")
	cmd.Flags().String(flagOutputPath, "./data/genesis.json", "the path to the output file for the genesis state")
	cmd.Flags().String(flagTrustLevel, "2/3", "the trust level for the genesis state (e.g., 2/3)")
	cmd.Flags().Uint32(flagTrustingPeriod, 0, "the trusting period for the genesis state")
	return cmd
}

func Fixtures(logger *zap.Logger) *cobra.Command {
	fixturesCmd := &cobra.Command{
		Use:   "fixtures",
		Short: "fixtures",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	fixturesCmd.AddCommand(
		UpdateClientCmd(logger),
		MembershipCmd(logger),
		Misbehaviour(logger),
	)

	return fixturesCmd
}

// UpdateClientCmd creates a command to update the client
func UpdateClientCmd(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-client",
		Short: "update client",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := godotenv.Load()
			if err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}

			// Read RPC endpoint from environment variable
			tendermintRpcEndpoint := os.Getenv("TENDERMINT_RPC_URL")
			if tendermintRpcEndpoint == "" {
				return fmt.Errorf("TENDERMINT_RPC_URL environment variable is required in .env file")
			}
			tendermintRpcClient, err := rpchttp.New(tendermintRpcEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create RPC client: %w", err)
			}
			ethRpcEndpoint := os.Getenv("ETH_RPC_URL")
			if ethRpcEndpoint == "" {
				return fmt.Errorf("ETH_RPC_URL environment variable is required in .env file")
			}

			hexAddress := os.Getenv("CONTRACT_ADDRESS")
			if hexAddress == "" {
				return fmt.Errorf("CONTRACT_ADDRESS environment variable is required in .env file")
			}

			privKey := os.Getenv("PRIVATE_KEY")
			if privKey == "" {
				return fmt.Errorf("PRIVATE_KEY environment variable is required in .env file")
			}
			privateKey, err := keys.RestoreKey(privKey)
			if err != nil {
				return fmt.Errorf("failed to restore private key: %w", err)
			}

			chainIdEth := os.Getenv("CHAIN_ID")
			if chainIdEth == "" {
				return fmt.Errorf("CHAIN_ID environment variable is required in .env file")
			}

			chainIdInt := big.NewInt(0)
			chainIdInt, ok := chainIdInt.SetString(chainIdEth, 10)
			if !ok {
				return fmt.Errorf("invalid chain id: %v", err)
			}

			ethClient, err := ethclient.Dial(ethRpcEndpoint)
			if err != nil {
				return fmt.Errorf("failed to create Ethereum client: %w", err)
			}

			trustedBlock, err := cmd.Flags().GetInt64(flagTrustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted block: %w", err)
			}
			trustingPeriod, err := cmd.Flags().GetUint32(flagTrustingPeriod)
			if err != nil {
				return fmt.Errorf("failed to get trusting period: %w", err)
			}
			trustLevel, err := cmd.Flags().GetString(flagTrustLevel)
			if err != nil {
				return fmt.Errorf("failed to get trust level from flag: %w", err)
			}
			proofType, err := cmd.Flags().GetString(flagProofType)
			if err != nil {
				return fmt.Errorf("failed to get proof type from flag: %w", err)
			}
			genesis, err := tendermintClient.GetGenesis(trustedBlock, trustingPeriod, trustLevel, proofType)
			if err != nil {
				return fmt.Errorf("failed to get genesis: %w", err)
			}

			trustedClientState := genesis.TrustedClientState
			trustedConsensusState := genesis.TrustedConsensusState

			targetBlock, err := cmd.Flags().GetInt64(flagTargetBlock)
			if err != nil {
				return fmt.Errorf("failed to get target block: %w", err)
			}

			trustedLightBlock, err := tendermintClient.GetLightBlock(tendermintRpcClient, trustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted light block: %w", err)
			}
			targetLightBlock, err := tendermintClient.GetLightBlock(tendermintRpcClient, int64(targetBlock))
			if err != nil {
				return fmt.Errorf("failed to get target light block: %w", err)
			}

			chainId := trustedLightBlock.SignedHeader.Header.ChainID
			revision := types.ParseChainID(chainId)
			header := tendermintClient.Header{
				SignedHeader: targetLightBlock.SignedHeader,
				ValidatorSet: targetLightBlock.ValSet,
				TrustedHeight: types.Height{
					RevisionNumber: revision,
					RevisionHeight: uint64(trustedLightBlock.SignedHeader.Header.Height),
				},
				TrustedValidators: trustedLightBlock.ValSet,
			}

			currentTime := time.Now().Unix()

			untrustedHeaderCommit := header.SignedHeader.Commit
			if untrustedHeaderCommit == nil {
				return fmt.Errorf("untrusted header commit is nil")
			}

			untrustedHeaderSigs := untrustedHeaderCommit.Signatures

			for i, sig := range untrustedHeaderSigs {
				trustedValidator := header.TrustedValidators.Validators[i]

				pub := trustedValidator.PubKey.Bytes()
				sigData := sig.Signature
				if len(sigData) != 64 {
					return fmt.Errorf("invalid signature length: %d", len(sigData))
				}

				voteData := untrustedHeaderCommit.VoteSignBytes(chainId, int32(i))
				R := sigData[:32]
				S, err := edwards25519.NewScalar().SetCanonicalBytes(sigData[32:])
				if err != nil {
					panic("ed25519: invalid signature")
				}

				hasher := sha512.New()
				// Precompute H = SHA512(R || A || msg)
				A := pub

				hasher.Reset()
				hasher.Write(R)
				hasher.Write(A)
				hasher.Write(voteData)
				sum := hasher.Sum(nil)

				H, err := edwards25519.NewScalar().SetUniformBytes(sum)
				if err != nil {
					panic("ed25519: internal error: setting scalar failed")
				}

				if !ed25519.Verify(pub, voteData, sigData) {
					panic("failed to verify signature outside circuit")
				}

				aX, aY, _ := utils.DecompressPoint(A)
				rX, rY, _ := utils.DecompressPoint(R)
				h := utils.ScalarToBigInt(H)
				s := utils.ScalarToBigInt(S)

				var circuit, assignment PreHashCircuit[Fp25519, Fr25519]

				_r1cs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
				// TODO: import pk, vk

				assignment.Sig = eddsa.Signature[Fp25519, Fr25519]{
					R: sw_emulated.AffinePoint[Fp25519]{
						X: emulated.ValueOf[Fp25519](rX),
						Y: emulated.ValueOf[Fp25519](rY),
					},
					S: emulated.ValueOf[Fr25519](s),
				}

				assignment.Hash = emulated.ValueOf[Fr25519](h)

				assignment.Pub = eddsa.PublicKey[Fp25519, Fr25519]{
					A: sw_emulated.AffinePoint[Fp25519]{
						X: emulated.ValueOf[Fp25519](aX),
						Y: emulated.ValueOf[Fp25519](aY),
					},
				}

				var pk groth16.ProvingKey
				pkPath, err := cmd.Flags().GetString(flagProvingKey)
				if err != nil {
					return fmt.Errorf("failed to get proving key path from flag: %w", err)
				}
				f, err := os.Open(pkPath)
				if err != nil {
					return fmt.Errorf("failed to open proving key file: %w", err)
				}
				defer f.Close()
				_, err = pk.ReadFrom(f)
				if err != nil {
					return fmt.Errorf("failed to read proving key: %w", err)
				}
				// witness
				witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
				if err != nil {
					panic(err)
				}
				witness.Public()

				proof, err := groth16.Prove(_r1cs, pk, witness)
				if err != nil {
					panic(err)
				}

				tendermintAddr := common.HexToAddress(hexAddress)
				ics07Tendermint, err := tendermintContract.NewContractSP1ICS07Tendermint(
					tendermintAddr,
					ethClient,
				)
				if err != nil {
					return fmt.Errorf("failed to create ICS07 Tendermint contract: %w", err)
				}

				auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIdInt)
				if err != nil {
					return fmt.Errorf("failed to create auth transactor: %w", err)
				}

				publicKey, err := keys.PublicKey(privateKey)
				if err != nil {
					log.Fatal(err)
				}

				fromAddress := crypto.PubkeyToAddress(*publicKey)
				nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
				if err != nil {
					log.Fatal(err)
				}
				gasPrice, err := ethClient.SuggestGasPrice(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				auth.Nonce = big.NewInt(int64(nonce))
				auth.Value = big.NewInt(0)     // in wei
				auth.GasLimit = uint64(300000) // in units
				auth.GasPrice = gasPrice

				proposedHeader := tendermintContract.IICS07TendermintMsgsHeader{
					SignedHeader: tendermintClient.ParseSignedHeader(targetLightBlock.SignedHeader),
					ValidatorSet: tendermintClient.ParseValidatorSet(targetLightBlock.ValSet),
					TrustedHeight: tendermintContract.IICS02ClientMsgsHeight{
						RevisionNumber: revision,
						RevisionHeight: uint64(trustedLightBlock.SignedHeader.Header.Height),
					},
					TrustedNextValidatorSet: tendermintClient.ParseValidatorSet(targetLightBlock.NextValSet),
				}

				proofBigInt, commitmentPoks, commitments, err := ProofToBigInts(proof)
				if err != nil {
					return fmt.Errorf("failed to convert proof to big ints: %w", err)
				}
				msg := tendermintContract.IUpdateClientMsgsMsgUpdateClient{
					ClientState:           trustedClientState,
					TrustedConsensusState: trustedConsensusState,
					ProposedHeader:        proposedHeader,
					Time:                  big.NewInt(currentTime),
					Proof:                 proofBigInt,
					CommitmentPok:         commitmentPoks,
					Commitments:           commitments,
				}

				tx, err := ics07Tendermint.UpdateClientMsg(auth, msg)
				if err != nil {
					return fmt.Errorf("failed to verify membership: %w", err)
				}
				fmt.Println("tx hash: ", tx.Hash().Hex())

			}

			return nil
		},
	}

	cmd.Flags().String(flagProofType, "groth16", "the type of proof to use (groth16, plonk)")
	cmd.Flags().Int64(flagTrustedBlock, 0, "the trusted block height, if <height> is 0 then catch latest block")
	cmd.Flags().String(flagOutput, "json", "the output structure for the genesis state (json, file)")
	cmd.Flags().String(flagOutputPath, "./data/genesis.json", "the path to the output file for the genesis state")
	cmd.Flags().String(flagTrustLevel, "2/3", "the trust level for the genesis state (e.g., 2/3)")
	cmd.Flags().Uint32(flagTrustingPeriod, 0, "the trusting period for the genesis state")
	cmd.Flags().Int64(flagTargetBlock, 0, "the target block height for update client")
	cmd.Flags().Bool(flagMembership, true, "verify membership/non-membership proof")
	cmd.Flags().String(flagProvingKey, "./pk.bin", "provking for the circuit")
	return cmd
}

// MembershipCmd creates a command to manage membership
func MembershipCmd(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "membership",
		Short: "membership <key_path> <is_base64> <membership_type>",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := godotenv.Load()
			if err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}

			// Read RPC endpoint from environment variable
			tendermintRpcEndpoint := os.Getenv("TENDERMINT_RPC_URL")
			if tendermintRpcEndpoint == "" {
				return fmt.Errorf("TENDERMINT_RPC_URL environment variable is required in .env file")
			}
			tendermintRpcClient, err := rpchttp.New(tendermintRpcEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create RPC client: %w", err)
			}
			ethRpcEndpoint := os.Getenv("ETH_RPC_URL")
			if ethRpcEndpoint == "" {
				return fmt.Errorf("ETH_RPC_URL environment variable is required in .env file")
			}

			hexAddress := os.Getenv("CONTRACT_ADDRESS")
			if hexAddress == "" {
				return fmt.Errorf("CONTRACT_ADDRESS environment variable is required in .env file")
			}

			privKey := os.Getenv("PRIVATE_KEY")
			if privKey == "" {
				return fmt.Errorf("PRIVATE_KEY environment variable is required in .env file")
			}
			privateKey, err := keys.RestoreKey(privKey)
			if err != nil {
				return fmt.Errorf("failed to restore private key: %w", err)
			}

			chainIdEth := os.Getenv("CHAIN_ID")
			if chainIdEth == "" {
				return fmt.Errorf("CHAIN_ID environment variable is required in .env file")
			}

			chainIdInt := big.NewInt(0)
			chainIdInt, ok := chainIdInt.SetString(chainIdEth, 10)
			if !ok {
				return fmt.Errorf("invalid chain id: %v", err)
			}

			ethClient, err := ethclient.Dial(ethRpcEndpoint)
			if err != nil {
				return fmt.Errorf("failed to create Ethereum client: %w", err)
			}

			trustedBlock, err := cmd.Flags().GetInt64(flagTrustedBlock)
			if err != nil {
				return fmt.Errorf("failed to get trusted block: %w", err)
			}
			trustingPeriod, err := cmd.Flags().GetUint32(flagTrustingPeriod)
			if err != nil {
				return fmt.Errorf("failed to get trusting period: %w", err)
			}
			trustLevel, err := cmd.Flags().GetString(flagTrustLevel)
			if err != nil {
				return fmt.Errorf("failed to get trust level from flag: %w", err)
			}
			proofType, err := cmd.Flags().GetString(flagProofType)
			if err != nil {
				return fmt.Errorf("failed to get proof type from flag: %w", err)
			}
			genesis, err := tendermintClient.GetGenesis(trustedBlock, trustingPeriod, trustLevel, proofType)
			if err != nil {
				return fmt.Errorf("failed to get genesis: %w", err)
			}

			tendermintAddr := common.HexToAddress(hexAddress)
			ics07Tendermint, err := tendermintContract.NewContractSP1ICS07Tendermint(
				tendermintAddr,
				ethClient,
			)
			if err != nil {
				return fmt.Errorf("failed to create ICS07 Tendermint contract: %w", err)
			}

			isMembership, err := cmd.Flags().GetBool(flagMembership)
			if err != nil {
				return fmt.Errorf("failed to get membership flag: %w", err)
			}

			publicKey, err := keys.PublicKey(privateKey)
			if err != nil {
				log.Fatal(err)
			}

			fromAddress := crypto.PubkeyToAddress(*publicKey)
			nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}
			gasPrice, err := ethClient.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIdInt)
			if err != nil {
				return fmt.Errorf("failed to create auth transactor: %w", err)
			}
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)     // in wei
			auth.GasLimit = uint64(300000) // in units
			auth.GasPrice = gasPrice

			kvPairs, proofs, err := runner.RunMembership(tendermintRpcClient, args[0], trustedBlock, args[1] == "true")
			if err != nil {
				return err
			}

			membershipType, err := strconv.Atoi(args[2])
			if err != nil {
				return fmt.Errorf("invalid membership type: %w", err)
			}
			if isMembership {
				msg := tendermintContract.ILightClientMsgsMsgVerifyMembership{
					Height:                tendermintContract.IICS02ClientMsgsHeight(genesis.TrustedClientState.LatestHeight),
					KvPairs:               kvPairs,
					MerkleProofs:          proofs,
					AppHash:               genesis.TrustedConsensusState.Root,
					TrustedConsensusState: tendermintContract.IICS07TendermintMsgsConsensusState(genesis.TrustedConsensusState),
					MembershipType:        uint8(membershipType),
				}

				tx, err := ics07Tendermint.VerifyMembership(auth, msg)
				if err != nil {
					return fmt.Errorf("failed to verify membership: %w", err)
				}
				fmt.Println("tx hash: ", tx.Hash().Hex())
			} else {
				msg := tendermintContract.ILightClientMsgsMsgVerifyNonMembership{
					Height:                tendermintContract.IICS02ClientMsgsHeight(genesis.TrustedClientState.LatestHeight),
					KvPairs:               kvPairs,
					MerkleProofs:          proofs,
					AppHash:               genesis.TrustedConsensusState.Root,
					TrustedConsensusState: tendermintContract.IICS07TendermintMsgsConsensusState(genesis.TrustedConsensusState),
					MembershipType:        uint8(membershipType),
				}

				ABI, err := tendermintContract.ContractSP1ICS07TendermintMetaData.GetAbi() //ContractMetadata is the contract whose bindings you have created
				if err != nil {
					panic(fmt.Errorf("could not get contract abi: %v", err))
				}

				tx, err := ics07Tendermint.VerifyNonMembership(auth, msg)
				if err != nil {
					return fmt.Errorf("failed to verify non-membership: %w", err)
				}
				fmt.Println("tx hash: ", tx.Hash().Hex())
			}
			return nil
		},
	}
	cmd.Flags().String(flagProofType, "groth16", "the type of proof to use (groth16, plonk)")
	cmd.Flags().Int64(flagTrustedBlock, 0, "the trusted block height, if <height> is 0 then catch latest block")
	cmd.Flags().String(flagOutput, "json", "the output structure for the genesis state (json, file)")
	cmd.Flags().String(flagOutputPath, "./data/genesis.json", "the path to the output file for the genesis state")
	cmd.Flags().String(flagTrustLevel, "2/3", "the trust level for the genesis state (e.g., 2/3)")
	cmd.Flags().Uint32(flagTrustingPeriod, 0, "the trusting period for the genesis state")
	cmd.Flags().Bool(flagMembership, true, "verify membership/non-membership proof")
	return cmd
}

func Misbehaviour(logger *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "misbehaviour",
		Short: "misbehaviour",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

func ProofToBigInts(proof groth16.Proof) ([8]*big.Int, [2]*big.Int, [2]*big.Int, error) {
	var out [8]*big.Int
	var commitmentPoks [2]*big.Int
	var commitments [2]*big.Int

	p, ok := proof.(*groth16_bn254.Proof)
	if !ok {
		return out, commitmentPoks, commitments, fmt.Errorf("expected BN254 proof")
	}

	// A (G1)
	p.Ar.X.BigInt(out[0])
	p.Ar.Y.BigInt(out[1])

	// B (G2) — NOTE the order (imag, real) for Solidity
	p.Bs.X.A0.BigInt(out[2])
	p.Bs.X.A1.BigInt(out[3])
	p.Bs.Y.A0.BigInt(out[4])
	p.Bs.Y.A1.BigInt(out[5])

	// C (G1)
	p.Krs.X.BigInt(out[6])
	p.Krs.Y.BigInt(out[7])

	p.CommitmentPok.X.BigInt(commitmentPoks[0])
	p.CommitmentPok.Y.BigInt(commitmentPoks[1])

	p.Commitments[0].X.BigInt(commitments[0])
	p.Commitments[0].Y.BigInt(commitments[1])
	return out, commitmentPoks, commitments, nil
}
