package lcd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/types"
	"github.com/shinecloudnet/explorer/backend/utils"
)

func Validator(address string) (result ValidatorVo, err error) {
	url := fmt.Sprintf(UrlValidator, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}
	var response ValidatorVoResponse
	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return response.Result.toValidatorVo(), nil
}

func Validators(page, size int) (result []ValidatorVo) {
	url := fmt.Sprintf(UrlValidators, conf.Get().Hub.LcdUrl, types.TypeValStatusBonded, page, size)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get Validators error", logger.String("err", err.Error()))
		return result
	}

	var response ValidatorsVoResponse
	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("Unmarshal Validators error", logger.String("err", err.Error()))
		return result
	}
	for _, val := range response.Result {
		result = append(result, val.toValidatorVo())
	}

	url = fmt.Sprintf(UrlValidators, conf.Get().Hub.LcdUrl, types.TypeValStatusUnbonding, page, size)
	resBytes, err = utils.Get(url)
	if err != nil {
		logger.Error("get Validators error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("Unmarshal Validators error", logger.String("err", err.Error()))
		return result
	}
	for _, val := range response.Result {
		result = append(result, val.toValidatorVo())
	}

	url = fmt.Sprintf(UrlValidators, conf.Get().Hub.LcdUrl, types.TypeValStatusUnbonded, page, size)
	resBytes, err = utils.Get(url)
	if err != nil {
		logger.Error("get Validators error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("Unmarshal Validators error", logger.String("err", err.Error()))
		return result
	}
	for _, val := range response.Result {
		result = append(result, val.toValidatorVo())
	}

	return result
}

func QueryWithdrawAddr(address string) (result string) {
	url := fmt.Sprintf(UrlWithdrawAddress, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	var response StringResponse
	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("Unmarshal Validator withdraw address error", logger.String("err", err.Error()))
		return result
	}
	//result = strings.Trim(response.Result, "\"")
	return response.Result
}

func GetDelegationsByDelAddr(delAddr string) (height string, delegations []DelegationVo) {
	url := fmt.Sprintf(UrlDelegationsByDelegator, conf.Get().Hub.LcdUrl, delAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}
	var response DelegationVoResponse
	if err := json.Unmarshal(resAsBytes, &response); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return response.Height, response.Result
}

func GetDelegationsByValidatorAddr(valAddr string) (delegations []DelegationVo) {

	url := fmt.Sprintf(UrlDelegationsByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}
	var response DelegationVoResponse
	if err := json.Unmarshal(resAsBytes, &response); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return response.Result
}

func GetWithdrawAddressByValidatorAcc(validatorAcc string) (string, error) {

	url := fmt.Sprintf(UrlDistributionWithdrawAddressByValidatorAcc, conf.Get().Hub.LcdUrl, validatorAcc)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return "", err
	}

	var withdrawAddr StringResponse
	if err := json.Unmarshal(resAsBytes, &withdrawAddr); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}

	return withdrawAddr.Result, nil
}

func GetDistributionRewardsByValidatorAcc(validatorAcc string) (utils.CoinsAsStr, []RewardsFromDelegations, utils.CoinsAsStr, error) {

	url := fmt.Sprintf(UrlDistributionRewardsByValidatorAcc, conf.Get().Hub.LcdUrl, validatorAcc)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return nil, nil, nil, err
	}

	var rewardsResponse DistributionRewardsResponse
	err = json.Unmarshal(resAsBytes, &rewardsResponse)
	if err != nil {
		return nil, nil, nil, err
	}

	validatorOperator := utils.Convert(conf.Get().Hub.Prefix.ValAddr, validatorAcc)
	url = fmt.Sprintf(UrlDistributionInfoOfValidatorAcc, conf.Get().Hub.LcdUrl, validatorOperator)

	resAsBytes, err = utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return nil, nil, nil, err
	}
	var distributionInfoResposne ValidatorDistributionInfoResponse
	err = json.Unmarshal(resAsBytes, &distributionInfoResposne)
	if err != nil {
		return nil, nil, nil, err
	}

	rewards := rewardsResponse.Result.Rewards
	totalReward := rewardsResponse.Result.Total
	commission := distributionInfoResposne.Result.ValCommission

	totalAmount := 0.0
	if len(totalReward) > 0 {
		totalAmount, err = strconv.ParseFloat(totalReward[0].Amount, 64)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	commissionAmount := 0.0
	if len(commission) > 0 {
		commissionAmount, err = strconv.ParseFloat(commission[0].Amount, 64)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	totalAmount += commissionAmount
	if len(totalReward) > 0 {
		totalReward[0].Amount = strconv.FormatFloat(totalAmount, 'f', -1, 64)
	}

	return commission, rewards, totalReward, nil
}

func GetJailedUntilAndMissedBlocksCountByConsensusPublicKey(publicKey string) (string, string, string, error) {

	url := fmt.Sprintf(UrlValidatorsSigningInfoByConsensuPublicKey, conf.Get().Hub.LcdUrl, publicKey)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return "", "", "", err
	}

	var valSign ValidatorSigningInfo

	err = json.Unmarshal(resAsBytes, &valSign)
	if err != nil {
		return "", "", "", err
	}

	return valSign.JailedUntil, valSign.MissedBlocksCount, valSign.StartHeight, nil
}

func GetRedelegationsByValidatorAddr(valAddr string) (redelegations []ReDelegations) {

	url := fmt.Sprintf(UrlRedelegationsByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}
	var response ReDelegationsResponse
	if err := json.Unmarshal(resAsBytes, &response); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return response.Result
}

func GetUnbondingDelegationsByValidatorAddr(valAddr string) (unbondingDelegations []UnbondingDelegation) {

	url := fmt.Sprintf(UrlUnbondingDelegationByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}
	var response UnbondingDelegationsResponse
	if err := json.Unmarshal(resAsBytes, &response); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return response.Result.toUnbondingDelegationList()
}

func GetUnbondingDelegationsByDelegatorAddr(delAddr string) (unbondingDelegations []UnbondingDelegation) {

	url := fmt.Sprintf(UrlUnbondingDelegationByDelegator, conf.Get().Hub.LcdUrl, delAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}
	var response UnbondingDelegationsResponse
	if err := json.Unmarshal(resAsBytes, &response); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return response.Result.toUnbondingDelegationList()
}

func DelegationByValidator(address string) (result []DelegationVo) {
	url := fmt.Sprintf(UrlDelegationByVal, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	var reponse DelegationVoResponse
	if err := json.Unmarshal(resBytes, &reponse); err != nil {
		logger.Error("Unmarshal Delegation error", logger.String("err", err.Error()))
		return result
	}
	return reponse.Result
}

func StakePool() (result StakePoolVo) {
	url := fmt.Sprintf(UrlStakePool, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	var response StakePoolVoResponse
	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("Unmarshal StakePool error", logger.String("err", err.Error()))
		return result
	}
	return response.Result
}

func SignInfo(consensusPubkey string) (result SignInfoVo) {
	url := fmt.Sprintf(UrlSignInfo, conf.Get().Hub.LcdUrl, consensusPubkey)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	var response SignInfoVoResponse
	if err := json.Unmarshal(resBytes, &response); err != nil {
		logger.Error("Unmarshal SignInfoVo error", logger.String("err", err.Error()))
		return result
	}
	return response.Result
}
