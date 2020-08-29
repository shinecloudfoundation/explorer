package lcd_test

import (
	"encoding/json"
	"testing"

	"github.com/shinecloudnet/explorer/backend/lcd"
)

func TestGetAssetTokens(t *testing.T) {
	res := lcd.GetAssetTokens()

	bytesData, _ := json.Marshal(res)
	t.Log(string(bytesData))
}

func TestGetAssetGateways(t *testing.T) {
	res := lcd.GetAssetGateways()

	bytesData, _ := json.Marshal(res)
	t.Log(string(bytesData))
}
