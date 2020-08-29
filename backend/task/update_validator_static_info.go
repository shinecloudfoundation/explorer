package task

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/service"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type ValidatorStaticInfo struct {
}

func (task ValidatorStaticInfo) Name() string {
	return "update_validator_static_data"
}

func (task ValidatorStaticInfo) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeValidatorStaticInfo

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task ValidatorStaticInfo) DoTask() error {
	validatorService := service.ValidatorService{}
	return validatorService.UpdateValidatorStaticInfo()
}
