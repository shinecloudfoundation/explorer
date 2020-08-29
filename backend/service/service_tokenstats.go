package service

import (
	"fmt"
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/lcd"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/types"
	"github.com/shinecloudnet/explorer/backend/utils"
	"github.com/shinecloudnet/explorer/backend/vo"
	"sync"
)

type TokenStatsService struct {
	BaseService
}

func (service *TokenStatsService) QueryTokenStats() (vo.TokenStatsVo, error) {

	var (
		tokenStatsVO     vo.TokenStatsVo
		supply           lcd.Coin
		circulation      lcd.Coin
		banktokenstats   lcd.TokenStats
		foundationBonded string
	)

	var group sync.WaitGroup
	group.Add(4)

	go func() {
		defer group.Done()
		var err error
		banktokenstats, err = lcd.GetBankTokenStats()
		if err != nil {
			logger.Error("GetBankTokenStats have error", logger.String("err", err.Error()))
		}
	}()
	go func() {
		defer group.Done()
		var err error
		supply, err = lcd.GetTokenStatsSupply()
		if err != nil {
			logger.Error("GetTokenStatsSupply have error", logger.String("err", err.Error()))
		}
	}()
	go func() {
		defer group.Done()
		var err error
		circulation, err = lcd.GetTokenStatsCirculation()
		if err != nil {
			logger.Error("GetTokenStatsCirculation have error", logger.String("err", err.Error()))
		}
	}()
	go func() {
		defer group.Done()
		if conf.Get().Hub.Prefix.AccAddr == types.MainnetAccPrefix {
			res := accountService.QueryDelegations(types.FoundationDelegatorAddr)
			amt := float64(0)
			for _, v := range res {
				amt += v.Amount.Amount
			}
			foundationBonded = utils.ParseStringFromFloat64(amt)
		} else {
			foundationBonded = "0"
		}
	}()
	group.Wait()

	//initsupply := lcd.GetTokenInitSupply()
	burnedtokens := lcd.GetTokens(banktokenstats.BurnedTokens)
	bondedtokens := lcd.GetTokens(banktokenstats.BondedTokens)

	tokenStatsVO.TotalsupplyTokens = LoadCoinVoFromLcdCoin(&supply)
	tokenStatsVO.CirculationTokens = LoadCoinVoFromLcdCoin(&circulation)
	//tokenStatsVO.InitsupplyTokens = LoadCoinVoFromLcdCoin(&initsupply)
	tokenStatsVO.DelegatedTokens = LoadCoinVoFromLcdCoin(&bondedtokens)
	tokenStatsVO.BurnedTokens = LoadCoinVoFromLcdCoin(&burnedtokens)
	if conf.Get().Hub.Prefix.AccAddr == types.MainnetAccPrefix {
		if balance, err := lcd.GetCommunityTax(); err == nil {
			tokenStatsVO.CommunityTax = LoadCoinVoFromLcdCoin(&balance)
		}
	}

	tokenStatsVO.FoundationBonded = LoadCoinVoFromLcdCoin(&lcd.Coin{
		Denom:  types.IRISUint,
		Amount: foundationBonded,
	})

	return tokenStatsVO, nil
}

func LoadCoinVoFromLcdCoin(coin *lcd.Coin) *vo.CoinVo {
	return &vo.CoinVo{
		Denom:  coin.Denom,
		Amount: coin.Amount,
	}
}

func (service *TokenStatsService) QueryTokensAccountTotal() (map[string]vo.TokenStatsSegment, error) {

	accounts, err := document.Account{}.GetAllAccount()
	if err != nil {
		logger.Error("QueryTokensAccountTotal have error", logger.String("err", err.Error()))
		return nil, err
	}

	var totalAmt = float64(0)

	for _, acc := range accounts {
		totalAmt += acc.Total.Amount
	}
	//fmt.Println("============>>:",len(accounts))

	if len(accounts) <= 5 {
		return computeSegment(accounts, totalAmt), nil
	}

	return computeSegment2(accounts, totalAmt), nil
}

func computeSegment(accounts []document.Account, totalAmt float64) map[string]vo.TokenStatsSegment {
	accList := make(map[string]vo.TokenStatsSegment)
	for index, acc := range accounts {
		rate, _ := utils.NewRatFromFloat64(acc.Total.Amount / totalAmt).Float64()
		accList[fmt.Sprint(index+1)] = vo.TokenStatsSegment{
			Percent: rate,
			TotalAmount: &vo.CoinVo{
				Denom:  acc.Total.Denom,
				Amount: utils.ParseStringFromFloat64(acc.Total.Amount),
			},
		}
	}
	return accList
}

func computeSegment2(accounts []document.Account, totalAmt float64) map[string]vo.TokenStatsSegment {
	result := make(map[string]vo.TokenStatsSegment)
	total := len(accounts)
	if total > 5 {
		result[fmt.Sprintf("%v-%v", 1, 5)] = CountNTotalAmount(0, 4, totalAmt, accounts)
		if total <= 10 {
			result[fmt.Sprintf("%v-%v", 6, total)] = CountNTotalAmount(5, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 6, 10)] = CountNTotalAmount(5, 9, totalAmt, accounts)
		if total <= 50 {
			result[fmt.Sprintf("%v-%v", 11, total)] = CountNTotalAmount(10, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 11, 50)] = CountNTotalAmount(10, 49, totalAmt, accounts)
		if total <= 100 {
			result[fmt.Sprintf("%v-%v", 51, total)] = CountNTotalAmount(50, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 51, 100)] = CountNTotalAmount(50, 99, totalAmt, accounts)
		if total <= 500 {
			result[fmt.Sprintf("%v-%v", 101, total)] = CountNTotalAmount(100, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 101, 500)] = CountNTotalAmount(100, 499, totalAmt, accounts)
		if total <= 1000 {
			result[fmt.Sprintf("%v-%v", 501, total)] = CountNTotalAmount(500, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 501, 1000)] = CountNTotalAmount(500, 999, totalAmt, accounts)
		result[fmt.Sprintf("%v-", 1001)] = CountNTotalAmount(1000, total-1, totalAmt, accounts)

	}
	return result
}

func CountNTotalAmount(start_index, end_index int, totalamt float64, account []document.Account) (result vo.TokenStatsSegment) {

	var vaTotalAmt float64
	if totalamt <= 0 {
		return
	}

	retdata := &vo.CoinVo{
		Denom: account[start_index].Total.Denom,
	}

	for pos := start_index; pos <= end_index; pos++ {
		vaTotalAmt += account[pos].Total.Amount
	}
	rate, _ := utils.NewRatFromFloat64(vaTotalAmt / totalamt).Float64()

	retdata.Amount = utils.ParseStringFromFloat64(vaTotalAmt)
	result.Percent = rate
	result.TotalAmount = retdata

	return
}
