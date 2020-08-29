package service

import (
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/utils"
	"go.uber.org/zap"
)

var (
	accountService   = &AccountService{}
	blockService     = &BlockService{}
	commonService    = &CommonService{}
	proposalService  = &ProposalService{}
	txService        = &TxService{}
	delegatorService = &DelegatorService{}
	govParamsService = &GovParamsService{}
	validatorService = &ValidatorService{}
	assetsService    = &AssetsService{}
	htlcService      = &HtlcService{}

	BlackValidatorsMap        = make(map[string]document.BlackList)
	ValidatorsDescriptionMap  = make(map[string]document.Description)
	BlackValidatorsHash       = utils.Md5Encryption([]byte("nil"))
	ValidatorsdescriptionHash = utils.Md5Encryption([]byte("nil"))
	ValidatorsHashHasNotInit  = utils.Md5Encryption([]byte("nil"))
)

const (
	_ Module = iota
	Account
	Block
	Common
	Proposal
	Tx
	Delegator
	GovParams
	Validator
	Asset
	Htlc
)

type Module int

func Get(m Module) Service {
	switch m {
	case Account:
		return accountService
	case Block:
		return blockService
	case Common:
		return commonService
	case Proposal:
		return proposalService
		// case Candidate:
		// 	return stakeService
	case Tx:
		return txService
	case Delegator:
		return delegatorService
	case GovParams:
		return govParamsService
	case Validator:
		return validatorService
	case Asset:
		return assetsService
	case Htlc:
		return htlcService
	}
	return nil
}

type Service interface {
	GetModule() Module
}

type BaseService struct {
	tid string
}

func (base *BaseService) SetTid(traceId string) {
	base.tid = traceId
}

func (base *BaseService) GetTid() string {
	return base.tid
}

func (base *BaseService) GetTraceLog() zap.Field {
	return logger.String("traceId", base.GetTid())
}

// store black list data in memory
// use redis to cache black list data in feature.
func init() {
	getBlackValidators()
	getValidatorsDescription()
}

func getBlackValidators() {
	blackListMap := document.BlackList{}.QueryBlackList()
	BlackValidatorsHash = utils.Md5Encryption(utils.MarshalJsonIgnoreErr(blackListMap))
	BlackValidatorsMap = blackListMap
}

func (b *BaseService) QueryBlackList() map[string]document.BlackList {
	if BlackValidatorsHash != ValidatorsHashHasNotInit {
		return BlackValidatorsMap
	} else {
		b.ReloadBlackValidators()
		return BlackValidatorsMap
	}
}

func (_ *BaseService) ReloadBlackValidators() {
	getBlackValidators()
}

func getValidatorsDescription() {
	descriptionMap := document.Validator{}.QueryValidatorDescription()
	ValidatorsdescriptionHash = utils.Md5Encryption(utils.MarshalJsonIgnoreErr(descriptionMap))
	ValidatorsDescriptionMap = descriptionMap
}
func (b *BaseService) QueryDescriptionList() map[string]document.Description {
	if ValidatorsdescriptionHash != ValidatorsHashHasNotInit {
		return ValidatorsDescriptionMap
	} else {
		b.ReloadValidatorsDescription()
		return ValidatorsDescriptionMap
	}
}

func (_ *BaseService) ReloadValidatorsDescription() {
	getValidatorsDescription()
}

func (_ *BaseService) BuildFTMoniker(From, To string) (Fmoniker, Tmoniker string) {
	if valaddr := utils.GetValaddr(To); valaddr != "" {
		if val, ok := ValidatorsDescriptionMap[valaddr]; ok {
			Tmoniker = val.Moniker
		}
		if one, ok := BlackValidatorsMap[valaddr]; ok {
			Tmoniker = one.Moniker
		}
	}
	if valaddr := utils.GetValaddr(From); valaddr != "" {
		if val, ok := ValidatorsDescriptionMap[valaddr]; ok {
			Fmoniker = val.Moniker
		}
		if one, ok := BlackValidatorsMap[valaddr]; ok {
			Fmoniker = one.Moniker
		}
	}
	return
}
