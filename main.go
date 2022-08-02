package main

import (
	"fmt"
	"log"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//"fmt"

const (
	grpcURL = "127.0.0.1:9090"
	chainID = "pylons-testnet-1"
)

func main() {

	f, err := os.OpenFile("TestLogs-Execute-Recipe-2.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: %v", err)
	}

	defer f.Close()

	myaddress := "pylo1clzj28ysxvfy420gafu7f73lvafv4l5yjj77cf"
	myprivateKey := "091d3c2ec85b818f0d517fa6c8f832cb6c69d296a4a95f0674879950d6fa6fb8"
	W := Wallet{address: myaddress}
	//testedFunction := "ExecuteRecipe"
	//incrementBy := 1
	offsetCal := 0

	// for i := 1; i < 10; i += incrementBy {

	// 	t1 := time.Now()
	// 	res, _ := TxsPylons(myaddress, myprivateKey, grpcURL, W.ExecuteRecipes(offsetCal, i, "cb130", "LOUDGetCharactercb130"), chainID)
	// 	offsetCal += i

	// 	t2 := time.Now()
	// 	diff := t2.Sub(t1)

	log.SetOutput(f)
	// 	log.Println(testedFunction, i, diff, res.TxResponse.Code, res.TxResponse.TxHash, res.TxResponse.GasUsed, res.TxResponse.GasWanted, myaddress)

	// 	fmt.Println(testedFunction, i, diff, res.TxResponse.Code, res.TxResponse.TxHash, res.TxResponse.GasUsed, res.TxResponse.GasWanted, myaddress)
	// }

	Msgs := W.ExecuteRecipes(offsetCal, 2, "cb130", "LOUDGetCharactercb130")

	for i, m := range Msgs {
		go threaded(myaddress, myprivateKey, m, i)
	}

	select {}

}

func threaded(myaddress string, myprivateKey string, m sdk.Msg, i int) {
	res, err := TxPylons(myaddress, myprivateKey, grpcURL, m, chainID)
	if err != nil {

		log.Fatal(err, res)
	}
	fmt.Println("Execute Recipe", i, res.TxResponse.Code, res.TxResponse.TxHash, res.TxResponse.GasUsed, res.TxResponse.GasWanted, myaddress)
	log.Println("Execute Recipe", i, res.TxResponse.Code, res.TxResponse.TxHash, res.TxResponse.GasUsed, res.TxResponse.GasWanted, myaddress)
}
