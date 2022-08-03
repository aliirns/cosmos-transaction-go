package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	//"strconv"
	"sync"
	//"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// 	"os"
	// 	"os/exec"
	// 	sdk "github.com/cosmos/cosmos-sdk/types"
)

//"fmt"

const (
	grpcURL     = "127.0.0.1:9090"
	chainID     = "pylons-testnet-1"
	_KEYNAME    = 0
	_ADDRESS    = 1
	_PRIVATEKEY = 2
)

var wg sync.WaitGroup

func main() {

	C := Chef{address: "pylo1clzj28ysxvfy420gafu7f73lvafv4l5yjj77cf"}
	msg := C.CreateComplexRecipeEasel("cb130", "cb131E")

	wg.Add(1)

	go threadedLoadTest("alii", C.address, "091d3c2ec85b818f0d517fa6c8f832cb6c69d296a4a95f0674879950d6fa6fb8", &msg, 0, "Create-Recipe")

	wg.Wait()

}

//read address and keys from CSV
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

//go
func threadedLoadTest(myKey string, myaddress string, myprivateKey string, m sdk.Msg, i int, txString string) {
	defer wg.Done()
	res, err := TxPylons(myaddress, myprivateKey, grpcURL, m, chainID)

	if err != nil {
		fmt.Println("A Failure as occured", err)
	}
	log.Println(i, txString, res.TxResponse.Code, res.TxResponse.Height, myKey, res.TxResponse.TxHash)
}
