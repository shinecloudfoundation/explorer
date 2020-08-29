package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/utils"
)

func GetAssetTokens() (result []AssetTokens) {
	url := fmt.Sprintf(UrlAssetTokens, conf.Get().Hub.LcdUrl)
	var response IssuedTokenListResponse
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get AssetTokens error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("Unmarshal AssetTokens error", logger.String("err", err.Error()))
		return result
	}
	for _, token := range response.Result {
		result = append(result, AssetTokens{
			BaseToken{
				Id:              token.Symbol,
				Family:          "",
				Source:          "native",
				Gateway:         "",
				Symbol:          token.Symbol,
				Name:            token.Name,
				Decimal:         token.Decimals,
				CanonicalSymbol: token.Symbol,
				MinUnitAlias:    "",
				InitialSupply:   token.TotalSupply,
				MaxSupply:       token.TotalSupply,
				Mintable:        token.Mintable,
				Owner:           token.Owner,
			},
		})
	}
	return result
}

func GetAssetGateways() (result []AssetGateways) {
	url := fmt.Sprintf(UrlAssetGateways, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get GetAssetGateways error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal GetAssetGateways error", logger.String("err", err.Error()))
		return result
	}
	return result
}
