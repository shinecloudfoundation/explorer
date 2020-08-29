package document_test

import (
	"encoding/json"
	"testing"

	"github.com/shinecloudnet/explorer/backend/orm/document"
)

func TestAsset_GetAllAssets(t *testing.T) {
	allAsset, err := document.AssetToken{}.GetAllAssets()
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range allAsset {
		t.Logf("k: %v \nv: %v\n", k, v)
	}
}

func TestAssetToken_GetAssetTokenDetail(t *testing.T) {
	allAsset, err := document.AssetToken{}.GetAssetTokenDetail("mondex.sun")
	if err != nil {
		t.Fatal(err)
	}

	bytesmsg, _ := json.Marshal(allAsset)
	t.Log(string(bytesmsg))
}
