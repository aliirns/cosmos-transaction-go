package main

import (
	"context"
	"encoding/hex"
	"fmt"

	//"fmt"
	"log"

	pylonsApp "github.com/Pylons-tech/pylons/app"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"
)

func SetPrefixes(accountAddressPrefix string) {
	// Set prefixes
	accountPubKeyPrefix := accountAddressPrefix + "pub"
	validatorAddressPrefix := accountAddressPrefix + "valoper"
	validatorPubKeyPrefix := accountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := accountAddressPrefix + "valcons"
	consNodePubKeyPrefix := accountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(accountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}

func dialGrpc(endpoint string) (*grpc.ClientConn, error) {
	grpcConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())

	if err != nil {
		log.Printf("grpc.Dial: %v", err)
		return &grpc.ClientConn{}, err
	}

	return grpcConn, nil
}

//Get Account and Account Sequence
func getAccount(address string) (authtypes.AccountI, error) {

	grpcConn, err := dialGrpc(grpcURL)
	if err != nil {
		return nil, err
	}

	defer grpcConn.Close()

	authClient := authtypes.NewQueryClient(grpcConn)
	authRes, err := authClient.Account(context.Background(), &authtypes.QueryAccountRequest{Address: address})
	if err != nil {
		return nil, err
	}

	var account authtypes.AccountI
	if err := pylonsApp.DefaultConfig().InterfaceRegistry.UnpackAny(authRes.Account, &account); err != nil {
		return nil, err
	}

	return account, nil
}

//Single MSG BroadCast
func TxPylons(accountAddress string, privateKeyHex string, grpcURL string, msg sdk.Msg, chainID string) (*txtypes.BroadcastTxResponse, error) {
	grpcConn, err := dialGrpc(grpcURL)
	if err != nil {
		return nil, err
	}

	defer grpcConn.Close()

	encfg := pylonsApp.DefaultConfig()
	txBuilder := encfg.TxConfig.NewTxBuilder()
	theaccount := accountAddress

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	//creating transaction
	txBuilder.SetGasLimit(uint64(200000))
	err = txBuilder.SetMsgs([]sdk.Msg{msg}...)
	if err != nil {
		return nil, err
	}

	//export private key --unsafe --unarmored-hex
	keyBytes, _ := hex.DecodeString(privateKeyHex)
	key := secp256k1.PrivKey{Key: keyBytes}

	//get Account and Sequence
	account, err := getAccount(theaccount)
	if err != nil {
		return nil, err
	}

	//signing Transactions
	sigV2 := signing.SignatureV2{
		PubKey: key.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       chainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	sigV2, err = tx.SignWithPrivKey(
		encfg.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		&key,
		encfg.TxConfig,
		account.GetSequence(),
	)

	err = txBuilder.SetSignatures(sigV2)
	txBytes, err := encfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	//Broadcasting transactions
	txClient := txtypes.NewServiceClient(grpcConn)
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)

	return grpcRes, err
}

//Multiple MSGS Brodcasr
func TxsPylons(accountAddress string, privateKeyHex string, grpcURL string, msgs []sdk.Msg, chainID string) (*txtypes.BroadcastTxResponse, error) {

	grpcConn, err := dialGrpc(grpcURL)
	if err != nil {
		return nil, err
	}

	defer grpcConn.Close()

	//Configuration
	encfg := pylonsApp.DefaultConfig()
	txBuilder := encfg.TxConfig.NewTxBuilder()
	theaccount := accountAddress

	//Validate Message
	// msg := types.MsgCreateAccount{
	// 	Creator:  theaccount,
	// 	Username: "the usrename",
	// 	Token:    "", ReferralAddress: ""}
	for _, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	//Set Estimated Gas Limit
	//txBuilder.SetGasLimit(20)

	//creating transaction
	err = txBuilder.SetMsgs(msgs...)
	if err != nil {
		return nil, err
	}

	//export private key --unsafe --unarmored-hex
	keyBytes, _ := hex.DecodeString(privateKeyHex)
	key := secp256k1.PrivKey{Key: keyBytes}

	//get Account and Sequence
	account, err := getAccount(theaccount)
	if err != nil {
		return nil, err
	}

	//signing Transactions
	sigV2 := signing.SignatureV2{
		PubKey: key.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       chainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	sigV2, err = tx.SignWithPrivKey(
		encfg.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		&key,
		encfg.TxConfig,
		account.GetSequence(),
	)

	err = txBuilder.SetSignatures(sigV2)
	txBytes, err := encfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	//Broadcasting transactions
	txClient := txtypes.NewServiceClient(grpcConn)

	simres, eres := txClient.Simulate(
		context.Background(),
		&txtypes.SimulateRequest{
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	fmt.Println(simres, eres)

	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_ASYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)

	return grpcRes, err
}
