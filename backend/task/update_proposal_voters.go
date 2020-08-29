package task

import (
	"github.com/shinecloudnet/explorer/backend/conf"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/orm/document"
	"github.com/shinecloudnet/explorer/backend/service"
	"github.com/shinecloudnet/explorer/backend/utils"
)

type UpdateProposalVoters struct{}

func (task UpdateProposalVoters) Name() string {
	return "update_proposal_voters"
}

func (task UpdateProposalVoters) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeProposalVoters

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateProposalVoters) DoTask() error {
	status := []string{document.ProposalStatusVoting}
	sorts := []string{document.Proposal_Field_VotingEndTime}
	proposals, err := document.Proposal{}.GetProposalsByStatus(status, sorts, false)
	if err != nil {
		return err
	}
	proposalService := service.ProposalService{}
	proposalService.UpdateProposalVoters(proposals)

	return nil
}
