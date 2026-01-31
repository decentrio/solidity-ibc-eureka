package transaction

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"operator/keys"
	"os"
	"strings"

	tendermintContract "operator/bindings/SP1ICS07Tendermint"
	updateclient "operator/bindings/UpdateClient"
	services "operator/services"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/gogoproto/proto"
	channeltypesv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Handler struct {
}

func (h *Handler) SendEthTx(ctx services.Context, msg any) error {
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

	publicKey, err := keys.PublicKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := ctx.EthClient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := ctx.EthClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainIdInt := big.NewInt(0)
	chainIdInt, ok := chainIdInt.SetString(chainIdEth, 10)
	if !ok {
		return fmt.Errorf("invalid chain id: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIdInt)
	if err != nil {
		return fmt.Errorf("failed to create auth transactor: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	hexAddress := os.Getenv("CONTRACT_ADDRESS")
	if hexAddress == "" {
		return fmt.Errorf("CONTRACT_ADDRESS environment variable is required in .env file")
	}

	tendermintAddr := common.HexToAddress(hexAddress)
	ics07Tendermint, err := tendermintContract.NewContractSP1ICS07Tendermint(
		tendermintAddr,
		ctx.EthClient(),
	)
	if err != nil {
		return fmt.Errorf("failed to create ICS07 Tendermint contract: %w", err)
	}

	switch msg := msg.(type) {
	case updateclient.IUpdateClientMsgsMsgUpdateClient:
		parsedABI, err := abi.JSON(strings.NewReader("../../abi/SP1ICS07Tendermint.json"))
		if err != nil {
			panic(err)
		}

		data, err := parsedABI.Pack(
			"transfer",
			msg,
		)
		if err != nil {
			panic(err)
		}

		_, err = ics07Tendermint.UpdateClient(auth, data)
		if err != nil {
			return fmt.Errorf("failed to verify membership: %w", err)
		}
	case tendermintContract.ILightClientMsgsMsgVerifyMembership:
		_, err := ics07Tendermint.VerifyMembership(auth, msg)
		if err != nil {
			return fmt.Errorf("failed to verify membership: %w", err)
		}
	case tendermintContract.ILightClientMsgsMsgVerifyNonMembership:
		_, err := ics07Tendermint.VerifyNonMembership(auth, msg)
		if err != nil {
			return fmt.Errorf("failed to verify membership: %w", err)
		}
	default:
		return fmt.Errorf("unsupported message type")
	}

	return nil
}

func (h *Handler) SendCosmosTx(svcCtx services.Context, msg any) error {
	protoMsg, ok := msg.(proto.Message)
	if !ok {
		return fmt.Errorf("message must be a proto.Message")
	}
	// Get the private key from environment variable
	privKeyHex := os.Getenv("COSMOS_PRIVATE_KEY")
	if privKeyHex == "" {
		return fmt.Errorf("COSMOS_PRIVATE_KEY environment variable is required in .env file")
	}

	// Decode the private key
	privKeyBytes, err := hex.DecodeString(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		return fmt.Errorf("failed to decode private key: %w", err)
	}

	privKey := secp256k1.PrivKey{Key: privKeyBytes}
	signerAddr := sdk.AccAddress(privKey.PubKey().Address())

	// Get chain configuration from environment
	chainID := os.Getenv("COSMOS_CHAIN_ID")
	if chainID == "" {
		return fmt.Errorf("COSMOS_CHAIN_ID environment variable is required in .env file")
	}

	// Get gas and fee configuration
	gasLimit := uint64(200000) // Default gas limit
	if gasStr := os.Getenv("COSMOS_GAS_LIMIT"); gasStr != "" {
		if _, err := fmt.Sscanf(gasStr, "%d", &gasLimit); err != nil {
			return fmt.Errorf("failed to parse COSMOS_GAS_LIMIT: %w", err)
		}
	}

	feeDenom := os.Getenv("COSMOS_FEE_DENOM")
	if feeDenom == "" {
		feeDenom = "stake" // Default fee denom
	}

	feeAmount := int64(1000) // Default fee amount
	if feeStr := os.Getenv("COSMOS_FEE_AMOUNT"); feeStr != "" {
		if _, err := fmt.Sscanf(feeStr, "%d", &feeAmount); err != nil {
			return fmt.Errorf("failed to parse COSMOS_FEE_AMOUNT: %w", err)
		}
	}

	// Query account info (account number and sequence) from the chain
	accountNumber, sequence, err := h.queryAccountInfo(svcCtx, signerAddr.String())
	if err != nil {
		return fmt.Errorf("failed to query account info: %w", err)
	}

	// Setup encoding config
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)
	channeltypesv2.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)
	txConfig := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	// Build the transaction
	txBuilder := txConfig.NewTxBuilder()

	// Convert the proto.Message to sdk.Msg
	sdkMsg, ok := protoMsg.(sdk.Msg)
	if !ok {
		return fmt.Errorf("message does not implement sdk.Msg interface")
	}

	if err := txBuilder.SetMsgs(sdkMsg); err != nil {
		return fmt.Errorf("failed to set messages: %w", err)
	}

	txBuilder.SetGasLimit(gasLimit)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin(feeDenom, sdkmath.NewInt(feeAmount))))

	// First, set an empty signature to populate signer info for sign bytes generation
	pubKey := privKey.PubKey()
	emptySig := sdksigning.SignatureV2{
		PubKey: pubKey,
		Data: &sdksigning.SingleSignatureData{
			SignMode:  sdksigning.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: sequence,
	}
	if err := txBuilder.SetSignatures(emptySig); err != nil {
		return fmt.Errorf("failed to set empty signature: %w", err)
	}

	// Create signer data
	signerData := authsigning.SignerData{
		Address:       signerAddr.String(),
		ChainID:       chainID,
		AccountNumber: accountNumber,
		Sequence:      sequence,
		PubKey:        pubKey,
	}

	// Get sign bytes using the adapter function
	signBytes, err := authsigning.GetSignBytesAdapter(
		context.Background(),
		txConfig.SignModeHandler(),
		sdksigning.SignMode_SIGN_MODE_DIRECT,
		signerData,
		txBuilder.GetTx(),
	)
	if err != nil {
		return fmt.Errorf("failed to get sign bytes: %w", err)
	}

	// Sign the bytes
	sigRaw, err := privKey.Sign(signBytes)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Set the actual signature
	sigV2 := sdksigning.SignatureV2{
		PubKey: pubKey,
		Data: &sdksigning.SingleSignatureData{
			SignMode:  sdksigning.SignMode_SIGN_MODE_DIRECT,
			Signature: sigRaw,
		},
		Sequence: sequence,
	}

	if err := txBuilder.SetSignatures(sigV2); err != nil {
		return fmt.Errorf("failed to set signatures: %w", err)
	}

	// Encode the transaction
	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return fmt.Errorf("failed to encode transaction: %w", err)
	}

	// Broadcast the transaction
	result, err := svcCtx.CosmosClient().BroadcastTxSync(context.Background(), txBytes)
	if err != nil {
		return fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	if result.Code != 0 {
		return fmt.Errorf("transaction failed with code %d: %s", result.Code, result.Log)
	}

	log.Printf("Transaction broadcast successfully. Hash: %s", result.Hash.String())

	return nil
}

// queryAccountInfo queries the account number and sequence for the given address
func (h *Handler) queryAccountInfo(svcCtx services.Context, address string) (uint64, uint64, error) {
	// Build the query request
	queryReq := &authtypes.QueryAccountRequest{
		Address: address,
	}

	reqBytes, err := proto.Marshal(queryReq)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to marshal query request: %w", err)
	}

	// Query path for auth account
	queryPath := "/cosmos.auth.v1beta1.Query/Account"

	// Make ABCI query
	result, err := svcCtx.CosmosClient().ABCIQuery(context.Background(), queryPath, reqBytes)
	if err != nil {
		return 0, 0, fmt.Errorf("ABCI query failed: %w", err)
	}

	if result.Response.Code != 0 {
		return 0, 0, fmt.Errorf("query failed with code %d: %s", result.Response.Code, result.Response.Log)
	}

	// Setup interface registry to decode the account
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(interfaceRegistry)

	// Decode the response
	var queryResp authtypes.QueryAccountResponse
	if err := proto.Unmarshal(result.Response.Value, &queryResp); err != nil {
		return 0, 0, fmt.Errorf("failed to unmarshal query response: %w", err)
	}

	// Unpack the account from Any type
	var account sdk.AccountI
	if err := interfaceRegistry.UnpackAny(queryResp.Account, &account); err != nil {
		return 0, 0, fmt.Errorf("failed to unpack account: %w", err)
	}

	return account.GetAccountNumber(), account.GetSequence(), nil
}
