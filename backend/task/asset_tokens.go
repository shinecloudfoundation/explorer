package task

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/service"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type UpdateAssetTokens struct{}

func (task UpdateAssetTokens) Name() string {
	return "update_asset_tokens"
}

func (task UpdateAssetTokens) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeAssetTokens

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateAssetTokens) DoTask() error {
	assetTokens, err := document.AssetToken{}.GetAllAssets()
	if err != nil {
		return err
	}

	assetService := service.Get(service.Asset).(*service.AssetsService)
	if err := assetService.UpdateAssetTokens(assetTokens); err != nil {
		return err
	}

	return nil
}
