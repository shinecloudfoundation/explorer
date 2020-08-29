package task

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/lcd"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type UpdateGovParams struct{}

func (task UpdateGovParams) Name() string {
	return "update_gov_params"
}
func (task UpdateGovParams) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeGovParams

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateGovParams) DoTask() error {
	curModuleKv, err := lcd.GetAllGovModuleParam()
	if err != nil {
		return err
	}

	err = document.GovParams{}.UpdateCurrentModuleParamValue(curModuleKv)
	if err != nil {
		return err
	}

	return nil
}
