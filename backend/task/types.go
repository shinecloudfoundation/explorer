package task

import (
	"fmt"
	"time"

	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/utils"
	"github.com/robfig/cron/v3"
)

var (
	engine = Engine{
		task: []TimerTask{},
	}

	taskControlModel document.TaskControl
	tcService        TaskControlService

	cstZone = time.FixedZone("CST", 8*3600)
)

func init() {
	engine.AppendTask(UpdateValidator{})
	engine.AppendTask(UpdateGovParams{})
	engine.AppendTask(UpdateValidatorIcons{})
	engine.AppendTask(UpdateAssetTokens{})
	//engine.AppendTask(UpdateAssetGateways{})
	engine.AppendTask(ValidatorStaticInfo{})
	engine.AppendTask(UpdateProposalVoters{})
}

type TimerTask interface {
	Start()
	Name() string
	DoTask() error
}

type Engine struct {
	task []TimerTask
}

func (e *Engine) Start() {
	if len(e.task) == 0 {
		return
	}
	for _, t := range e.task {
		var taskId = fmt.Sprintf("%s[%s]", t.Name(), utils.FmtTime(time.Now(), utils.DateFmtYYYYMMDD))
		logger.Info("timerTask begin to work", logger.String("taskId", taskId))
		go t.Start()
	}
}

func (e *Engine) AppendTask(task TimerTask) {
	e.task = append(e.task, task)
}

func Start() {
	engine.Start()

	// tasks manager by cron job
	c := cron.New()
	c.Start()

	txNumTask := TxNumGroupByDayTask{}
	txNumTask.init()
	c.AddFunc("01 0 * * *", func() {
		txNumTask.Start()
		new(StaticRewardsTask).Start()
		new(StaticValidatorTask).Start()
	})
}
