package controller

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shinecloudnet/explorer/backend/types"
	"github.com/shinecloudnet/explorer/backend/utils"
	"github.com/shinecloudnet/explorer/backend/vo"
)

// mux.Router registrars
func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlocks,
		registerQueryRecentBlocks,
		registerQueryValidatorSet,
		registerQueryBlockInfoByBlock,
		registerQueryBlockLatestHeight,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//type Block struct {
//	*service.BlockService
//}
//
//var block = Block{
//	service.Get(service.Block).(*service.BlockService),
//}

// @Summary latest
// @Description get block latestheight
// @Tags block
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.LatestHeightRespond	"success"
// @Router /api/block/latestheight [get]
func registerQueryBlockLatestHeight(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryBlockLatestHeight, "GET", func(request vo.IrisReq) interface{} {
		return block.QueryLatestHeight()
	})

	return nil
}

// @Summary list
// @Description get blocks
// @Tags block
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(10)
// @Success 200 {object} vo.BlockForListRespond	"success"
// @Router /api/blocks [get]
func registerQueryBlocks(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlocks, "GET", func(request vo.IrisReq) interface{} {
		block.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 100))
		result := block.QueryList(page, size)
		return result
	})

	return nil
}

// @Summary blocks recent
// @Description get recent blocks
// @Tags block
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.BlockInfoVoRespond	"success"
// @Router /api/blocks/recent [get]
func registerQueryRecentBlocks(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryRecentBlocks, "GET", func(request vo.IrisReq) interface{} {
		return block.QueryRecent()
	})

	return nil
}

// @Summary block validatorset
// @Description get  block validatorset
// @Tags block
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(10)
// @Param   height   path   int true    "block height"
// @Success 200 {object} vo.ValidatorSet	"success"
// @Router /api/block/validatorset/{height} [get]
func registerQueryValidatorSet(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlockValidatorSet, "GET", func(request vo.IrisReq) interface{} {
		block.SetTid(request.TraceId)
		height, err := strconv.ParseInt(Var(request, "height"), 10, 0)
		if err != nil || height < 1 {
			panic(types.CodeInValidParam)
		}

		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), DefaultPageNum))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), DefaultPageSize))
		result := block.GetValidatorSet(height, page, size)

		return result
	})
	return nil
}

// @Summary detail
// @Description get block info
// @Tags block
// @Accept  json
// @Produce  json
// @Param   height   path   int true    "block height"
// @Success 200 {object} vo.BlockInfo	"success"
// @Router /api/block/blockinfo/{height} [get]
func registerQueryBlockInfoByBlock(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlockInfo, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		height, err := strconv.ParseInt(Var(request, "height"), 10, 0)
		if err != nil || height < 1 {
			panic(types.CodeInValidParam)
		}
		return block.QueryBlockInfo(height)
	})
	return nil
}
