package task

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/service"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type UpdateAssetGateways struct{}

func (task UpdateAssetGateways) Name() string {
	return "update_asset_gateways"
}

func (task UpdateAssetGateways) Start() {
	timeInterval := conf.Get().Server.CronTimeAssetGateways
	taskName := task.Name()

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateAssetGateways) DoTask() error {
	assetGateways, err := document.AssetGateways{}.GetAllAssetGateways()
	if err != nil {
		return err
	}

	assetService := service.Get(service.Asset).(*service.AssetsService)
	if err := assetService.UpdateAssetGateway(assetGateways); err != nil {
		return err
	}

	return nil
}
