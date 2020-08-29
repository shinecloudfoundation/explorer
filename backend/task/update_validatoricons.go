package task

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/service"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type UpdateValidatorIcons struct{}

func (task UpdateValidatorIcons) Name() string {
	return "update_validator_icons"
}

func (task UpdateValidatorIcons) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeValidatorIcons

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateValidatorIcons) DoTask() error {
	err := new(service.ValidatorService).UpdateValidatorIcons()
	if err != nil {
		return err
	}

	return nil
}
