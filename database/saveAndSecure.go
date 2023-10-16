package database

import (
	_type "ddp/type"
	"encoding/json"
	"os"
)

// SaveData Function that handle for saving data into database
func SaveData(items *_type.Items) (int, error) {
	var tempHold []_type.ItemsJson
	_, saveInfoError := os.Stat("savedata.shikudb")
	if os.IsNotExist(saveInfoError) {
		itemInJson := _type.ItemsJson{ItemNames: items.ItemNames, ItemStock: items.ItemStock, ItemPrices: items.ItemPrices}
		tempHold = append(tempHold, itemInJson)
		jsonData, jsonMError := json.Marshal(tempHold)
		if jsonMError != nil {
			return 1, jsonMError
		}
		if writeError := os.WriteFile("savedata.shikudb", jsonData, os.ModePerm); writeError != nil {
			return 1, writeError
		}
		return 0, nil
	}
	saveFile, readSaveError := os.ReadFile("savedata.shikudb")
	if readSaveError != nil {
		return 1, readSaveError
	}
	if jsonUMError := json.Unmarshal(saveFile, &tempHold); jsonUMError != nil {
		return 1, jsonUMError
	}
	itemInJson := _type.ItemsJson{ItemNames: items.ItemNames, ItemStock: items.ItemStock, ItemPrices: items.ItemPrices}
	tempHold = append(tempHold, itemInJson)
	jsonNewData, jsonMError := json.Marshal(tempHold)
	if jsonMError != nil {
		return 1, jsonMError
	}
	if writeError := os.WriteFile("savedata.shikudb", jsonNewData, os.ModePerm); writeError != nil {
		return 1, writeError
	}
	return 0, nil
}
