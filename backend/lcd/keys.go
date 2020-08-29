package lcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/types"
	"github.com/shinecloudnet/explorer/backend/utils"
	"github.com/shinecloudnet/explorer/backend/vo"
)

type (
	Coin struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	}

	Value struct {
		Coins      []*Coin   `json:"coins"`
		Address    string    `json:"address"`
		PublicKey  PublicKey `json:"public_key"`
		AccountNum string    `json:"account_number"`
		Sequence   string    `json:"sequence"`
	}

	PublicKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	Account01411Response struct {
		Height string       `json:"height"`
		Result Account01411 `json:"result"`
	}
	Account01411 struct {
		Type  string `json:"type"`
		Value Value  `json:"value"`
	}
	TotalSupply struct {
		Height string `json:"height"`
		Result string `json:"result"`
	}
	Balance struct {
		Height string `json:"height"`
		Result []Coin `json:"result"`
	}
)

func buildAccountVo(acc Account01411) AccountVo {

	res := AccountVo{}

	res.Address = acc.Value.Address
	res.Sequence = acc.Value.Sequence
	res.AccountNumber = acc.Value.AccountNum
	res.PublicKey.Type = acc.Value.PublicKey.Type
	res.PublicKey.Value = acc.Value.PublicKey.Value

	coinsStrArr := []string{}
	for _, v := range acc.Value.Coins {
		coinsStrArr = append(coinsStrArr, v.Amount+v.Denom)
	}
	res.Coins = coinsStrArr
	return res
}

func Account(address string) (result AccountVo, err error) {
	url := fmt.Sprintf(UrlAccount, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}
	var response Account01411Response
	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	result = buildAccountVo(response.Result)
	return result, nil
}

func Faucet(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetAccountService, conf.Get().Server.FaucetUrl)
	return utils.Forward(req, uri)
}

func GetToken(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetApplyService, conf.Get().Server.FaucetUrl)
	return utils.Forward(req, uri)
}

func GetIconsByKey(key string) (string, error) {
	url := fmt.Sprintf(UrlLookupIconsByKeySuffix, key)
	resBytes, err := utils.Get(url)
	if err != nil {
		return "", err
	}
	var picdata vo.LookupIcons
	if err := json.Unmarshal(resBytes, &picdata); err != nil {
		logger.Error("get icons error", logger.String("err", err.Error()))
		return "", err
	}

	if picdata.Status.Code != 0 {
		return "", errors.New("get icons failed")
	}

	if len(picdata.Them) == 0 {
		return "", nil
	}

	return picdata.Them[0].Pictures.Primary.Url, nil
}
