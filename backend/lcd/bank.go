package lcd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type TokenStats struct {
	LooseTokens  []*Coin `json:"loose_tokens"`
	BurnedTokens []*Coin `json:"burned_tokens"`
	BondedTokens []*Coin `json:"bonded_tokens"`
	TotalSupply  []*Coin `json:"total_supply"`
}

type TokenSupplyResponse struct {
	Height string          `json:"height"`
	Result []TokenSupplyVo `json:"result"`
}

type TokenSupplyVo struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func GetBankTokenStats() (TokenStats, error) {
	var result TokenStats

	url := fmt.Sprintf(UrlBankTokenStats, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	var tokenSupplyResponse TokenSupplyResponse
	if err := json.Unmarshal(resBytes, &tokenSupplyResponse); err != nil {
		return result, err
	}

	stakePool := StakePool()

	result.BondedTokens = []*Coin{
		&Coin{
			Denom:  NativeTokenDenom,
			Amount: stakePool.BondedTokens,
		},
	}

	bondedTokenAmount, err := strconv.Atoi(stakePool.BondedTokens)
	if err != nil {
		return result, err
	}
	notBondedTokenAmount, err := strconv.Atoi(stakePool.NotBondedTokens)
	if err != nil {
		return result, err
	}

	totalBalanceOfNativeTokenResult, err := GetTokenTotalSupply(NativeTokenDenom)
	if err != nil {
		return result, err
	}
	nativeTokenAmountTotalSupply, err := strconv.Atoi(totalBalanceOfNativeTokenResult.Amount)
	if err != nil {
		return result, err
	}
	looseNativeTokens := nativeTokenAmountTotalSupply - bondedTokenAmount - notBondedTokenAmount

	for _, token := range tokenSupplyResponse.Result {
		result.TotalSupply = append(result.TotalSupply, &Coin{
			Denom:  token.Denom,
			Amount: token.Amount,
		})
		if token.Denom == NativeTokenDenom {
			result.LooseTokens = append(result.LooseTokens, &Coin{
				Denom:  token.Denom,
				Amount: strconv.Itoa(looseNativeTokens),
			})
		} else {
			result.LooseTokens = append(result.LooseTokens, &Coin{
				Denom:  token.Denom,
				Amount: token.Amount,
			})
		}
	}

	return result, nil
}

func GetTokens(data []*Coin) Coin {

	for _, val := range data {
		if val.Denom == utils.CoinTypeAtto {
			return Coin{Denom: val.Denom, Amount: val.Amount}
		}
	}
	return Coin{}
}

func GetTokenStatsCirculation() (Coin, error) {
	resBytes, err := utils.Get(UrlTokenStatsCirculation)
	if err != nil {
		return Coin{}, err
	}
	return Coin{
		Amount: string(resBytes),
		Denom:  utils.CoinTypeIris,
	}, nil
}

func GetTokenStatsSupply() (Coin, error) {
	resBytes, err := utils.Get(UrlTokenStatsSupply)
	if err != nil {
		return Coin{}, err
	}
	return Coin{
		Amount: string(resBytes),
		Denom:  utils.CoinTypeIris,
	}, nil
}
func GetCommunityTax() (Coin, error) {
	url := fmt.Sprintf(UrlAccount, conf.Get().Hub.LcdUrl, CommunityTaxAddr)
	resBytes, err := utils.Get(url)
	if err != nil {
		return Coin{}, err
	}
	acc := Account01411{}
	if err := json.Unmarshal(resBytes, &acc); err != nil {
		return Coin{}, err
	}

	return GetTokens(acc.Value.Coins), nil
}

func GetTokenTotalSupply(denom string) (Coin, error) {
	url := fmt.Sprintf(UrlTotalSupplyDenom, conf.Get().Hub.LcdUrl, denom)
	resBytes, err := utils.Get(url)
	if err != nil {
		return Coin{}, err
	}
	totalSupply := TotalSupply{}
	if err := json.Unmarshal(resBytes, &totalSupply); err != nil {
		return Coin{}, err
	}

	return Coin{Denom: denom, Amount: totalSupply.Result}, nil
}

func GetBalanceOfAddr(addr string) ([]Coin, error) {
	url := fmt.Sprintf(UrlBalance, conf.Get().Hub.LcdUrl, addr)
	resBytes, err := utils.Get(url)
	if err != nil {
		return nil, err
	}
	balance := Balance{}
	if err := json.Unmarshal(resBytes, &balance); err != nil {
		return nil, err
	}

	return balance.Result, nil
}
