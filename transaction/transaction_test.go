package transaction

import (
	"fmt"
	"testing"

	pylonsApp "github.com/Pylons-tech/pylons/app"
	"github.com/aliirns/cosmos-transaction/pylons"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	grpcURL  = "127.0.0.1:9090"
	address  = "pylo1clzj28ysxvfy420gafu7f73lvafv4l5yjj77cf"
	privKey  = "091d3c2ec85b818f0d517fa6c8f832cb6c69d296a4a95f0674879950d6fa6fb8"
	Sequence = 162
	chainID  = "pylons-testnet-1"
)

func TestGetAccount(t *testing.T) {
	t.Parallel()
	res, err := getAccount(address, grpcURL)

	if err != nil {
		t.FailNow()
	}

	Address := res.GetAddress().String()
	if address != Address {
		t.Errorf("got %v \n expected %v \n", Address, address)
	}

	seq := res.GetSequence()
	if seq != Sequence {
		t.Errorf("got %v \n expected %v\n", seq, Sequence)
	}

}

func TestCosmosTx(t *testing.T) {
	//SEND COINS
	coins, err := sdk.ParseCoinsNormalized("1upylon")
	if err != nil {
		t.FailNow()
	}

	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		t.FailNow()
	}
	receiver, err := sdk.AccAddressFromBech32("pylo1v9gr07lcfmpmy8c8f2paasjuwfg7leyvrue56g")
	if err != nil {
		t.FailNow()
	}

	msg := types.NewMsgSend(sdk.AccAddress(addr), sdk.AccAddress(receiver), coins)
	config := pylonsApp.DefaultConfig()
	res, err := CosmosTx(address, privKey, grpcURL, msg, chainID, config)
	if err != nil {
		t.Errorf("transaction failed")
		fmt.Println(err)

	}

	Code := res.TxResponse.Code
	if Code != 0 {
		t.Errorf("go %v \n expected %v \n", Code, 0)
	}

}

func TestCosmosTxs(t *testing.T) {
	//SEND COINS
	coins, err := sdk.ParseCoinsNormalized("1upylon")
	if err != nil {
		t.FailNow()
	}

	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		t.FailNow()
	}
	receiver, err := sdk.AccAddressFromBech32("pylo1v9gr07lcfmpmy8c8f2paasjuwfg7leyvrue56g")
	if err != nil {
		t.FailNow()
	}

	//mulitple transactions can be sent
	pylonsMsg := pylons.CreateComplexRecipeEasel("cb86", "cb131E", address)
	cosmosBankMsg := types.NewMsgSend(sdk.AccAddress(addr), sdk.AccAddress(receiver), coins)
	msg := []sdk.Msg{cosmosBankMsg, &pylonsMsg}

	//msg = append(msg, )
	config := pylonsApp.DefaultConfig()
	res, err := CosmosTxs(address, privKey, grpcURL, msg, chainID, config)
	if err != nil {
		t.Errorf("transaction failed")
		fmt.Println(err)

	}

	Code := res.TxResponse.Code
	if Code != 0 {
		t.Errorf("go %v \n expected %v \n", Code, 0)
	}

}
