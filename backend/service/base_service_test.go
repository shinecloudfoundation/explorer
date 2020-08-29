package service

import (
	"testing"

	"github.com/shinecloudnet/explorer/backend/utils"
)

func TestBaseService_QueryBlackList(t *testing.T) {
	service := BaseService{}

	res := service.QueryBlackList()
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}
