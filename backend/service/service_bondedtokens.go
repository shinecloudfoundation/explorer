package service

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/utils"
	"github.com/shinecloudnet/explorer/backend/vo"
)

type BondedTokensService struct {
	BaseService
}

func (service *BondedTokensService) QueryBondedTokensValidator(vtype string) (vo.BondedTokensRespond, error) {

	_, validators, err := document.Validator{}.GetValidatorListByPage(vtype, 0, 0, false, false)
	if err != nil {
		return nil, err
	}
	result := make([]*vo.BondedTokensVo, 0, len(validators))
	blackList := service.QueryBlackList()
	for _, val := range validators {
		bondedtoken := &vo.BondedTokensVo{
			Moniker:         val.Description.Moniker,
			Identity:        val.Description.Identity,
			VotingPower:     val.VotingPower,
			OperatorAddress: val.OperatorAddress,
			Icons:           val.Icons,
		}
		if item, ok := blackList[val.OperatorAddress]; ok {
			bondedtoken.Moniker = item.Moniker
			bondedtoken.Identity = item.Identity
		}
		bondedtoken.OwnerAddress = utils.Convert(conf.Get().Hub.Prefix.AccAddr, val.OperatorAddress)

		result = append(result, bondedtoken)
	}

	return result, nil
}
