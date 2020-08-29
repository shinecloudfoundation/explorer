package filter

import (
	"fmt"
	"time"

	"github.com/gorilla/mux"
	"github.com/shinecloudnet/explorer/backend/logger"
	"github.com/shinecloudnet/explorer/backend/types"
	"github.com/shinecloudnet/explorer/backend/vo"
)

// display user's request information,optional
type LogPreFilter struct {
}

func (LogPreFilter) Do(request *vo.IrisReq, data interface{}) (interface{}, types.BizCode) {
	start := time.Now()
	request.TraceId = fmt.Sprintf("%d", start.UnixNano())
	request.Start = start

	traceId := logger.String("traceId", request.TraceId)
	apiPath := logger.String("path", request.RequestURI)
	formValue := logger.Any("form", request.Form)
	urlValue := logger.Any("url", mux.Vars(request.Request))
	agent := logger.String("agent", request.UserAgent())
	defer func() {
		if r := recover(); r != nil {
			logger.Error("LogPostFilter failed", traceId)
		}
	}()

	logger.Info("LogPreFilter", traceId, apiPath, agent, formValue, urlValue)
	return nil, types.CodeSuccess
}

func (LogPreFilter) Match() []Condition {
	return []Condition{globalCondition}
}

func (LogPreFilter) Type() Type {
	return Pre
}

// display user's request information,optional
type LogPostFilter struct {
}

func (LogPostFilter) Do(request *vo.IrisReq, data interface{}) (interface{}, types.BizCode) {
	traceId := logger.String("traceId", request.TraceId)
	costSecond := time.Now().Unix() - request.Start.Unix()
	cost := logger.Int64("cost", costSecond)
	apiPath := logger.String("path", request.RequestURI)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("LogPostFilter failed", traceId)
		}
	}()
	logger.Info("LogPostFilter", apiPath, traceId, cost)
	if costSecond >= 3 {
		logger.Warn("LogPostFilter api too slow", apiPath, traceId, cost)
	}
	return nil, types.CodeSuccess
}

func (LogPostFilter) Match() []Condition {
	return []Condition{globalCondition}
}

func (LogPostFilter) Type() Type {
	return Pre
}
