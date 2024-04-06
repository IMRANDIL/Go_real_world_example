package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/imrandil/the_real_world/data"
)

// LoadDataFromJSON loads market data from a JSON file
func LoadDataFromJSON(filePath string) ([]data.MarketData, error) {
	// Read JSON file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a slice of MarketData
	var data []data.MarketData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	return data, nil
}
