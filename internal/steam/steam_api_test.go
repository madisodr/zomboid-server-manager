package steam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestParseACF(t *testing.T) {
	const testFilename = "./test-data/appworkshop_108600.acf"

	expectedACF := ACF{
		AppID:           "108600",
		SizeOnDisk:      "7953164623",
		NeedsUpdate:     "0",
		NeedsDownload:   "0",
		TimeLastUpdated: "0",
		TimeLastAppRan:  "0",
		LastBuildID:     "0",
		WorkshopItemsInstalled: map[string]WorkshopItem{
			"909647363": {
				Size:        "111006",
				TimeUpdated: "1538452157",
				Manifest:    "3386588971063334798",
			},
			"1510950729": {
				Size:        "164665858",
				TimeUpdated: "1648791612",
				Manifest:    "4082144191834930601",
			},
			"1516836158": {
				Size:        "12732881",
				TimeUpdated: "1660964843",
				Manifest:    "8461494127040368662",
			},
		},
	}

	acf, err := ParseACF(testFilename)
	assert.Nil(t, err, "ParseACF returned an error")
	assert.Equal(t, expectedACF, acf, "ParseACF did not return the expected ACF")
}
