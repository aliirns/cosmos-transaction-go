package pylons

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//SAMPLE MESSAGES FOR THE PYLONS CHAIN

func CreateCookbook(indexOffset int, amount int, address string) []sdk.Msg {

	msgs := []sdk.Msg{}

	for i := indexOffset; i < indexOffset+amount; i++ {
		msgs = append(msgs, &types.MsgCreateCookbook{Creator: address, Id: fmt.Sprintf("cb%v", i), Name: fmt.Sprintf("testCookBook%v", i), Description: "this is a cookbook", Developer: "ali", Version: "v1.0.0", SupportEmail: "e@email.com", Enabled: true})
	}

	return msgs

}

func CreateRecipes(indexOffset int, amount int, cookBookID string, address string) []sdk.Msg {
	msgs := []sdk.Msg{}

	entries := types.EntriesList{CoinOutputs: []types.CoinOutput{}, ItemOutputs: []types.ItemOutput{}, ItemModifyOutputs: []types.ItemModifyOutput{}}
	for i := indexOffset; i < indexOffset+amount; i++ {
		msgs = append(msgs, &types.MsgCreateRecipe{Creator: address, CookbookId: cookBookID, Id: fmt.Sprintf("testRecipe%v%v", cookBookID, i), Name: fmt.Sprintf("testRecipeName%v", i), Description: "this is a recipe", Version: "v1.0.1", CoinInputs: []types.CoinInput{}, ItemInputs: []types.ItemInput{}, Entries: entries, Outputs: []types.WeightedOutputs{}, BlockInterval: 0, CostPerBlock: sdk.NewCoin("upylon", sdkmath.NewInt(0)), Enabled: false, ExtraInfo: "extraInfo"})
	}

	return msgs

}

func ExecuteRecipes(indexOffset int, amount int, cookBookID string, recipeId string, address string) []sdk.Msg {
	msgs := []sdk.Msg{}

	for i := indexOffset; i < indexOffset+amount; i++ {
		msgs = append(msgs, &types.MsgExecuteRecipe{Creator: address, CookbookId: cookBookID, RecipeId: recipeId, CoinInputsIndex: 0, ItemIds: []string{}, PaymentInfos: []types.PaymentInfo{}})
	}

	return msgs

}

func CreateComplexRecipeEasel(cookBookID string, recipeID string, address string) types.MsgCreateRecipe {
	msg := types.MsgCreateRecipe{
		Creator:       address,
		CookbookId:    cookBookID,
		Id:            recipeID,
		Name:          "Hfwfqhhrsh4stushth",
		Description:   "Hrqrahar4hrahrajtsjfsjtsjzt",
		Version:       "v1.1.1",
		CoinInputs:    []types.CoinInput{},
		ItemInputs:    []types.ItemInput{},
		BlockInterval: 0,
		CostPerBlock:  sdk.NewCoin("upylon", sdkmath.NewInt(0)),
		Enabled:       true,
		ExtraInfo:     "extrainfo",
		Entries: types.EntriesList{ItemOutputs: []types.ItemOutput{types.ItemOutput{
			Id:      "Easel_NFT",
			Doubles: []types.DoubleParam{{Key: "Residual", WeightRanges: []types.DoubleWeightRange{{Lower: sdk.MustNewDecFromStr("0.5"), Upper: sdk.MustNewDecFromStr("0.5"), Weight: 1}}}},
			Longs: []types.LongParam{
				{Key: "Quantity", WeightRanges: []types.IntWeightRange{{Lower: 1, Upper: 1, Weight: 1}}},
				{Key: "Width", WeightRanges: []types.IntWeightRange{{Lower: 1080, Upper: 1080, Weight: 1}}},
				{Key: "Height", WeightRanges: []types.IntWeightRange{{Lower: 2400, Upper: 2400, Weight: 1}}},
				{Key: "Duration", WeightRanges: []types.IntWeightRange{{Lower: 0, Upper: 0, Weight: 1}}}},
			Strings: []types.StringParam{
				{
					Key:   "Name",
					Value: "Hfwfqhhrsh4stushth",
				}, {
					Key:   "App_Type",
					Value: "Easel",
				}, {
					Key:   "Description",
					Value: "Hrqrahar4hrahrajtsjfsjtsjzt",
				}, {
					Key:   "Hashtags",
					Value: "",
				}, {
					Key:   "NFT_Format",
					Value: "Image",
				}, {
					Key:   "NFT_URL",
					Value: "https://ipfs.io/ipfs/bafkreiav5d5co4giq42ojsrypebuw4nnwrk3doycxdqskyvcmf53bb345e",
				}, {
					Key:   "Thumbnail_URL",
					Value: "",
				}, {
					Key:   "Creator",
					Value: "Wvoyw9v9uvsu9gs",
				}, {
					Key:   "cid",
					Value: "bafkreiav5d5co4giq42ojsrypebuw4nnwrk3doycxdqskyvcmf53bb345e",
				}, {
					Key:   "fileSize",
					Value: "206.92KB",
				},
			},
			TransferFee:     []sdk.Coin{sdk.NewCoin("upylon", sdkmath.NewInt(1))},
			TradePercentage: sdk.MustNewDecFromStr("0.5"),
			Quantity:        1,
			AmountMinted:    0,
			Tradeable:       true,
		}}},
		Outputs: []types.WeightedOutputs{{EntryIds: []string{"Easel_NFT"}, Weight: 1}},
	}

	return msg

}
