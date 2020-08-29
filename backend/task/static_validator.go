package task

import (
	"time"

	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

type StaticValidatorTask struct {
	validator document.Validator
}

func (task StaticValidatorTask) Name() string {
	return "static_validator"
}
func (task StaticValidatorTask) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeTxNumByDay

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
}

func (task StaticValidatorTask) DoTask() error {
	ops, err := task.getAllValidatorTokens()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return document.ExStaticValidator{}.Batch(ops)
}

func (task StaticValidatorTask) getAllValidatorTokens() ([]txn.Op, error) {

	validators, err := task.getValidatorFromDb()
	if err != nil {
		return nil, err
	}

	return task.saveExValidatorStaticOps(validators)
}

func (task StaticValidatorTask) getValidatorFromDb() ([]document.Validator, error) {
	return task.validator.GetAllValidator()
}

func (task StaticValidatorTask) saveExValidatorStaticOps(validators []document.Validator) ([]txn.Op, error) {
	today := utils.TruncateTime(time.Now().In(cstZone), utils.Day)
	ops := make([]txn.Op, 0, len(validators))
	now := time.Now().Unix()
	for _, addr := range validators {
		item, err := task.loadValidatorTokens(addr, today)
		item.CreateAt = now
		if err != nil {
			continue
		}
		op := txn.Op{
			C:      document.CollectionNameExValidatorStatic,
			Id:     bson.NewObjectId(),
			Insert: item,
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (task StaticValidatorTask) loadValidatorTokens(validator document.Validator, today time.Time) (document.ExStaticValidator, error) {

	item := document.ExStaticValidator{
		Id:              bson.NewObjectId(),
		OperatorAddress: validator.OperatorAddress,
		Status:          validator.GetValidatorStatus(),
		Date:            today,
		SelfBond:        validator.SelfBond,
		DelegatorShares: validator.DelegatorShares,
		Tokens:          validator.Tokens,
		Commission:      validator.Commission,
	}
	delegationFromOther := funcSubStr(item.Tokens, item.SelfBond)
	if delegationFromOther != nil {
		item.Delegations = delegationFromOther.FloatString(18)
	} else {
		item.Delegations = "0.0"
	}
	return item, nil
}
