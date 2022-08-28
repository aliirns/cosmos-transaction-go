package transaction

import (
	"context"
	"encoding/hex"

	//"fmt"
	"log"

	pylonsApp "github.com/Pylons-tech/pylons/app"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"
)

type CosmosTransaction struct {
	grpcURL        string
	accountAddress string
	privateKeyHex  string
	networkConfig  network.Config
	chainID        string
}

func (C *CosmosTransaction) SetPrefixes(accountAddressPrefix string) {
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

func (C *CosmosTransaction) dialGrpc() (*grpc.ClientConn, error) {
	grpcConn, err := grpc.Dial(
		C.grpcURL,
		grpc.WithInsecure())

	if err != nil {
		log.Printf("grpc.Dial: %v", err)
		return &grpc.ClientConn{}, err
	}

	return grpcConn, nil
}

//Get Account and Account Sequence
func (C *CosmosTransaction) GetAccount() (authtypes.AccountI, error) {

	grpcConn, err := C.dialGrpc()
	if err != nil {
		return nil, err
	}

	defer grpcConn.Close()

	authClient := authtypes.NewQueryClient(grpcConn)
	authRes, err := authClient.Account(context.Background(), &authtypes.QueryAccountRequest{Address: C.accountAddress})
	if err != nil {
		return nil, err
	}

	var account authtypes.AccountI
	if err := pylonsApp.DefaultConfig().InterfaceRegistry.UnpackAny(authRes.Account, &account); err != nil {
		return nil, err
	}

	return account, nil
}

func (C *CosmosTransaction) SignValidateTx(msg sdk.Msg, gasLimit uint64) ([]byte, error) {

	//Validate Msg
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	txBuilder := C.networkConfig.TxConfig.NewTxBuilder()

	//creating transaction
	txBuilder.SetGasLimit(gasLimit)
	err := txBuilder.SetMsgs([]sdk.Msg{msg}...)
	if err != nil {
		return nil, err
	}

	//export private key --unsafe --unarmored-hex
	keyBytes, _ := hex.DecodeString(C.privateKeyHex)
	key := secp256k1.PrivKey{Key: keyBytes}

	//get Account and Sequence
	account, err := C.GetAccount()
	if err != nil {
		return nil, err
	}

	//Signing Tx
	//Empty Signature Step 1
	sigV2 := signing.SignatureV2{
		PubKey: key.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       C.chainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	//Signature Step2
	sigV2, err = tx.SignWithPrivKey(
		C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		&key,
		C.networkConfig.TxConfig,
		account.GetSequence(),
	)
	if err != nil {
		return nil, err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	txBytes, err := C.networkConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil

}

func (C *CosmosTransaction) SignTx(msg sdk.Msg, gasLimit uint64) ([]byte, error) {

	txBuilder := C.networkConfig.TxConfig.NewTxBuilder()

	//creating transaction
	txBuilder.SetGasLimit(gasLimit)
	err := txBuilder.SetMsgs([]sdk.Msg{msg}...)
	if err != nil {
		return nil, err
	}

	//export private key --unsafe --unarmored-hex
	keyBytes, _ := hex.DecodeString(C.privateKeyHex)
	key := secp256k1.PrivKey{Key: keyBytes}

	//get Account and Sequence
	account, err := C.GetAccount()
	if err != nil {
		return nil, err
	}

	//Signing Tx
	//Empty Signature Step 1
	sigV2 := signing.SignatureV2{
		PubKey: key.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       C.chainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	//Signature Step2
	sigV2, err = tx.SignWithPrivKey(
		C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		&key,
		C.networkConfig.TxConfig,
		account.GetSequence(),
	)
	if err != nil {
		return nil, err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	txBytes, err := C.networkConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil

}

func (C *CosmosTransaction) SignValidateTxs(msgs []sdk.Msg, gasLimit uint64) ([]byte, error) {
	grpcConn, err := C.dialGrpc()
	if err != nil {
		return []byte{}, err
	}
	defer grpcConn.Close()

	//Validate Msg
	for _, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	txBuilder := C.networkConfig.TxConfig.NewTxBuilder()

	//creating transaction
	txBuilder.SetGasLimit(gasLimit)
	err = txBuilder.SetMsgs(msgs...)
	if err != nil {
		return nil, err
	}

	//export private key --unsafe --unarmored-hex
	keyBytes, _ := hex.DecodeString(C.privateKeyHex)
	key := secp256k1.PrivKey{Key: keyBytes}

	//get Account and Sequence
	account, err := C.GetAccount()
	if err != nil {
		return nil, err
	}

	//Signing Tx
	//Empty Signature Step 1
	sigV2 := signing.SignatureV2{
		PubKey: key.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       C.chainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	//Signature Step2
	sigV2, err = tx.SignWithPrivKey(
		C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		&key,
		C.networkConfig.TxConfig,
		account.GetSequence(),
	)
	if err != nil {
		return nil, err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	txBytes, err := C.networkConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil

}

func (C *CosmosTransaction) SignTxs(msgs []sdk.Msg, gasLimit uint64) ([]byte, error) {

	txBuilder := C.networkConfig.TxConfig.NewTxBuilder()

	//creating transaction
	txBuilder.SetGasLimit(gasLimit)
	err := txBuilder.SetMsgs(msgs...)
	if err != nil {
		return nil, err
	}

	//export private key --unsafe --unarmored-hex
	keyBytes, _ := hex.DecodeString(C.privateKeyHex)
	key := secp256k1.PrivKey{Key: keyBytes}

	//get Account and Sequence
	account, err := C.GetAccount()
	if err != nil {
		return nil, err
	}

	//Signing Tx
	//Empty Signature Step 1
	sigV2 := signing.SignatureV2{
		PubKey: key.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       C.chainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	//Signature Step2
	sigV2, err = tx.SignWithPrivKey(
		C.networkConfig.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		&key,
		C.networkConfig.TxConfig,
		account.GetSequence(),
	)
	if err != nil {
		return nil, err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	txBytes, err := C.networkConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil

}

func (C *CosmosTransaction) BroadcastTx(txBytes []byte) (txtypes.BroadcastTxResponse, error) {
	//Broadcasting transactions
	grpcConn, err := C.dialGrpc()
	if err != nil {
		return txtypes.BroadcastTxResponse{}, err
	}
	defer grpcConn.Close()

	txClient := txtypes.NewServiceClient(grpcConn)

	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_ASYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)

	return *grpcRes, err
}
