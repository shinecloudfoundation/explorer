package controller

import "github.com/shinecloudnet/explorer/backend/service"

var (
	account      service.AccountService
	block        service.BlockService
	common       service.CommonService
	gov          service.ProposalService
	govparams    service.GovParamsService
	stake        service.ValidatorService
	tx           service.TxService
	tokenstats   service.TokenStatsService
	bondedtokens service.BondedTokensService
	asset        service.AssetsService
	htlc         service.HtlcService
)
