package database

import (
	_type "ddp/type"
	"encoding/json"
	"os"
)

// LoadDatabase a simple function to load data from shiku database
func LoadDatabase() (int, []_type.ItemsJson, error) {
	var itemSLice []_type.ItemsJson
	_, err := os.Stat("savedata.shikudb")
	if os.IsNotExist(err) {
		return 2, nil, nil
	}
	readSaveData, readError := os.ReadFile("savedata.shikudb")
	if readError != nil {
		return 1, nil, readError
	}
	if jsonUMError := json.Unmarshal(readSaveData, &itemSLice); jsonUMError != nil {
		return 1, nil, jsonUMError
	}
	return 0, itemSLice, nil
}
