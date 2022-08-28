package transaction

import (
	"fmt"
	"testing"

	pylonsApp "github.com/Pylons-tech/pylons/app"
	txtypes "github.com/aliirns/cosmos-transaction-go/txtypes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

func TestGetAccount(t *testing.T) {

	t.Parallel()
	conf := pylonsApp.DefaultConfig()
	address := ""
	CT := CosmosTransaction{grpcURL: "", accountAddress: address, privateKeyHex: "", networkConfig: conf, chainID: "pylons-testnet-3"}
	res, err := CT.GetAccount()

	if err != nil {
		t.FailNow()
	}

	Address := res.GetAddress().String()
	if address != Address {
		t.Errorf("got %v \n expected %v \n", Address, address)
	}

}

func TestCosmosTx(t *testing.T) {
	//SEND COINS
	address := ""
	conf := pylonsApp.DefaultConfig()
	CT := CosmosTransaction{grpcURL: "", accountAddress: address, privateKeyHex: "", networkConfig: conf, chainID: "pylons-testnet-3"}
	CT.SetPrefixes("pylo")

	receiver := "pylo1v9gr07lcfmpmy8c8f2paasjuwfg7leyvrue56g"

	coins, err := sdk.ParseCoinsNormalized("100000upylon")
	if err != nil {
		t.FailNow()
	}
	cosmosBankMsg, err := txtypes.NewSendTx(address, string(receiver), coins)
	if err != nil {
		t.FailNow()
	}

	txBytes, err := CT.SignValidateTx(cosmosBankMsg, 700)
	if err != nil {
		t.FailNow()
	}

	res, err := CT.BroadcastTx(txBytes)
	if err != nil {
		t.Errorf("transaction failed")
		fmt.Println(err)

	}

	Code := res.TxResponse.Code
	if Code != 0 {
		t.Errorf("go %v \n expected %v \n", Code, 0)
	}
}

func TestIBCTx(t *testing.T) {
	address := "pylonsAccount"
	conf := pylonsApp.DefaultConfig()
	CT := CosmosTransaction{grpcURL: "", accountAddress: address, privateKeyHex: "", networkConfig: conf, chainID: "pylons-testnet-3"}
	CT.SetPrefixes("pylo")

	receiver := "axelarAccount"

	coin, err := sdk.ParseCoinNormalized("1000upylon")
	if err != nil {
		t.FailNow()
	}

	H := clienttypes.NewHeight(3, 3560000)
	cosmosIBCMsg := txtypes.NewIBCTransferTx(address, receiver, coin, "channel-58", H)

	txBytes, err := CT.SignValidateTx(cosmosIBCMsg, 700)
	if err != nil {
		t.FailNow()
	}

	res, err := CT.BroadcastTx(txBytes)
	if err != nil {
		t.Errorf("transaction failed")
		fmt.Println(err)

	}

	Code := res.TxResponse.Code
	if Code != 0 {
		t.Errorf("go %v \n expected %v \n", Code, 0)
	}

}
