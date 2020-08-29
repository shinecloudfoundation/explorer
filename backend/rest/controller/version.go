package controller

import (
	"github.com/gorilla/mux"
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/types"
	"github.com/shinecloudnet/explorer/backend/vo"
)

// @Summary version
// @Description version
// @Tags version
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string  "success"
// @Router /api/version [get]
func RegisterQueryVersion(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryApiVersion, "GET", func(request vo.IrisReq) interface{} {
		return map[string]string{"version": conf.Get().Server.ApiVersion}
	})

	return nil
}
