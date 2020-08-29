package msgvo

import "encoding/json"

type TxMsgSetWithdrawAddress struct {
	DelegatorAddr string `json:"delegator_addr"`
	WithdrawAddr  string `json:"withdraw_addr"`
}

type TxMsgWithdrawDelegatorReward struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
}

type TxMsgWithdrawDelegatorRewards struct {
	DelegatorAddr string `json:"delegator_addr"`
}

// msg struct for validator withdraw
type TxMsgWithdrawValidatorRewards struct {
	ValidatorAddr string `json:"validator_addr"`
}

func (vo *TxMsgSetWithdrawAddress) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgWithdrawDelegatorReward) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgWithdrawDelegatorRewards) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgWithdrawValidatorRewards) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
