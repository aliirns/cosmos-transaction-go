package main

import (
	"context"
	"encoding/hex"
	"fmt"

	//"fmt"
	"log"

	pylonsApp "github.com/Pylons-tech/pylons/app"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"
)

const (
	grpcURL = "127.0.0.1:9090"
)

func main() {
	// fmt.Println("Started")
	// SetPrefixes("pylo")

	// grpcConn, err := dialGrpc(grpcURL)
	// if err != nil {
	// 	fmt.Println("")
	// }

	// defer grpcConn.Close()
	// //Configuration
	// encfg := pylonsApp.DefaultConfig()
	// txBuilder := encfg.TxConfig.NewTxBuilder()
	// theaccount := "pylo1v9gr07lcfmpmy8c8f2paasjuwfg7leyvrue56g"

	// //create message payload
	// msg := types.MsgCreateAccount{
	// 	Creator:  theaccount,
	// 	Username: "codeerKey",
	// 	Token:    "", ReferralAddress: ""}

	// if err := msg.ValidateBasic(); err != nil {
	// 	fmt.Println("Error in Validation", err)
	// }

	// // //creating transaction
	// txBuilder.SetGasLimit(uint64(200000))
	// err = txBuilder.SetMsgs([]sdk.Msg{&msg}...)

	// // //private key --unsafe --unarmored-hex
	// keyBytes, _ := hex.DecodeString("8f4b8a359aa5b989a60a401e680ec3bc6bc72ac9bad49c9cd60e3e664ce19f6c")
	// key := secp256k1.PrivKey{Key: keyBytes}

	// account, err := getAccount(theaccount)
	// if err != nil {
	// 	fmt.Println("error in get account", account, err)
	// 	return
	// }

	// sigV2 := signing.SignatureV2{
	// 	PubKey: key.PubKey(),
	// 	Data: &signing.SingleSignatureData{
	// 		SignMode:  encfg.TxConfig.SignModeHandler().DefaultMode(),
	// 		Signature: nil,
	// 	},
	// 	Sequence: account.GetSequence(),
	// }

	// err = txBuilder.SetSignatures(sigV2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// signerData := xauthsigning.SignerData{
	// 	ChainID:       "pylons-testnet-1",
	// 	AccountNumber: account.GetAccountNumber(),
	// 	Sequence:      account.GetSequence(),
	// }

	// sigV2, err = tx.SignWithPrivKey(
	// 	encfg.TxConfig.SignModeHandler().DefaultMode(),
	// 	signerData,
	// 	txBuilder,
	// 	&key,
	// 	encfg.TxConfig,
	// 	account.GetSequence(),
	// )

	// err = txBuilder.SetSignatures(sigV2)
	// txBytes, err := encfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// txClient := txtypes.NewServiceClient(grpcConn)

	// grpcRes, err := txClient.BroadcastTx(
	// 	context.Background(),
	// 	&txtypes.BroadcastTxRequest{
	// 		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
	// 		TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
	// 	},
	// )

	// fmt.Println(grpcRes, err)
	msg := types.MsgCreateAccount{
		Creator:  "pylo1rsx89p92y36fcymuwdzxr9v0gzt5ksdp8cv8lv",
		Username: "razin",
		Token:    "", ReferralAddress: ""}
	res, err := TxPylons("pylo1rsx89p92y36fcymuwdzxr9v0gzt5ksdp8cv8lv", "7501369bef07ec31db3213e017a0ad511fe96dcc919a21517ad1478d22a3cb34", grpcURL, &msg)
	fmt.Println(res, err)

}

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
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)

	if err != nil {
		log.Printf("grpc.Dial: %v", err)
		return &grpc.ClientConn{}, err
	}

	return grpcConn, nil
}

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

//
func TxPylons(accountAddress string, privateKeyHex string, grpcURL string, msg sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	SetPrefixes("pylo")

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
		ChainID:       "pylons-testnet-1",
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
