package main

//"fmt"

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Wallet struct {
	address string
}

func (w Wallet) CreateCookbook(indexOffset int, amount int) []sdk.Msg {

	msgs := []sdk.Msg{}

	for i := indexOffset; i < indexOffset+amount; i++ {
		msgs = append(msgs, &types.MsgCreateCookbook{Creator: w.address, Id: fmt.Sprintf("cb%v", i), Name: fmt.Sprintf("testCookBook%v", i), Description: "this is a cookbook", Developer: "ali", Version: "v1.0.0", SupportEmail: "e@email.com", Enabled: true})
	}

	return msgs

}

func (w Wallet) CreateRecipes(indexOffset int, amount int, cookBookID string) []sdk.Msg {
	msgs := []sdk.Msg{}

	entries := types.EntriesList{CoinOutputs: []types.CoinOutput{}, ItemOutputs: []types.ItemOutput{}, ItemModifyOutputs: []types.ItemModifyOutput{}}
	for i := indexOffset; i < indexOffset+amount; i++ {
		msgs = append(msgs, &types.MsgCreateRecipe{Creator: w.address, CookbookId: cookBookID, Id: fmt.Sprintf("testRecipe%v%v", cookBookID, i), Name: fmt.Sprintf("testRecipeName%v", i), Description: "this is a recipe", Version: "v1.0.1", CoinInputs: []types.CoinInput{}, ItemInputs: []types.ItemInput{}, Entries: entries, Outputs: []types.WeightedOutputs{}, BlockInterval: 0, CostPerBlock: sdk.NewCoin("upylon", sdkmath.NewInt(0)), Enabled: false, ExtraInfo: "extraInfo"})
	}

	return msgs

}

func (w Wallet) ExecuteRecipes(indexOffset int, amount int, cookBookID string, recipeId string) []sdk.Msg {
	msgs := []sdk.Msg{}

	for i := indexOffset; i < indexOffset+amount; i++ {
		msgs = append(msgs, &types.MsgExecuteRecipe{Creator: w.address, CookbookId: cookBookID, RecipeId: recipeId, CoinInputsIndex: 0, ItemIds: []string{}, PaymentInfos: []types.PaymentInfo{}})
	}

	return msgs

}
