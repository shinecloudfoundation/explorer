package task

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/service"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type UpdateValidator struct{}

func (task UpdateValidator) Name() string {
	return "update_validator"
}
func (task UpdateValidator) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeValidators

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})

}

func (task UpdateValidator) DoTask() error {
	validators, err := document.Validator{}.GetAllValidator()

	if err != nil {
		return err
	}

	validatorService := service.Get(service.Validator).(*service.ValidatorService)
	err = validatorService.UpdateValidators(validators)

	if err != nil {
		return err
	}

	return nil
}
