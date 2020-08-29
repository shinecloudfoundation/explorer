package vo

import (
	"fmt"

	"github.com/shinecloudnet/explorer/backend/utils"
)

type AccountVo struct {
	Amount          utils.Coins `json:"amount"`
	WithdrawAddress string      `json:"withdrawAddress"`
	Address         string      `json:"address"`
	Deposits        utils.Coin  `json:"deposits"`
	IsProfiler      bool        `json:"isProfiler"`
	Moniker         string      `json:"moniker"`
	Status          string      `json:"status"`
	OperatorAddress string      `json:"operator_address"`
}

func (a AccountVo) String() string {
	return fmt.Sprintf(`
		Amount          :%v
		WithdrawAddress :%v
		Address         :%v
		Deposits        :%v
		IsProfiler      :%v
		Moniker         :%v
		Status          :%v
		OperatorAddress :%v
		`, a.Amount, a.WithdrawAddress, a.Address, a.Deposits, a.IsProfiler, a.Moniker, a.Status, a.OperatorAddress)
}

type ValAccVo struct {
	AccountVo
	ValProfile
}

type AccountInfo struct {
	Rank     int          `json:"rank"`
	Address  string       `json:"address"`
	Balance  []utils.Coin `json:"balance"`
	Percent  float64      `json:"percent"`
	UpdateAt int64        `json:"update_at"`
}

func (a AccountInfo) String() string {
	return fmt.Sprintf(`
		Rank     :%v
		Address  :%v
		Balance  :%v
		Percent  :%v
		UpdateAt :%v
		`, a.Rank, a.Address, a.Balance, a.Percent, a.UpdateAt)
}

type AccountDelegationsVo struct {
	Address string     `json:"address"`
	Moniker string     `json:"moniker"`
	Amount  utils.Coin `json:"amount"`
	Shares  string     `json:"shares"`
	Height  int64     `json:"height,string"`
}

type AccountUnbondingDelegationsVo struct {
	Address string     `json:"address"`
	Moniker string     `json:"moniker"`
	Amount  utils.Coin `json:"amount"`
	Height  int64     `json:"height,string"`
	EndTime string     `json:"end_time"`
}

type AccountsInfoRespond []AccountInfo
type AccountDelegationsRespond []*AccountDelegationsVo
func (sort AccountDelegationsRespond) Len() int {
	return len(sort)
}
func (sort AccountDelegationsRespond) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}
func (sort AccountDelegationsRespond) Less(i, j int) bool {
	return sort[i].Height > sort[j].Height
}
type AccountUnbondingDelegationsRespond []*AccountUnbondingDelegationsVo
func (sort AccountUnbondingDelegationsRespond) Len() int {
	return len(sort)
}
func (sort AccountUnbondingDelegationsRespond) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}
func (sort AccountUnbondingDelegationsRespond) Less(i, j int) bool {
	return sort[i].Height > sort[j].Height
}


type AccountRewardsVo struct {
	TotalRewards       utils.CoinsAsStr     `json:"total_rewards"`
	DelagationsRewards []DelagationsRewards `json:"delagations_rewards"`
	CommissionRewards  utils.CoinsAsStr     `json:"commission_rewards"`
}

type DelagationsRewards struct {
	Address string           `json:"address"`
	Moniker string           `json:"moniker"`
	Amount  utils.CoinsAsStr `json:"amount"`
}
