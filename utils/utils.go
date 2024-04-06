package utils

import (
	"encoding/json"
	"io/ioutil"
)

// MarketData represents the structure of the table
type MarketData struct {
	State                string `json:"state"`
	District             string `json:"district"`
	Market               string `json:"market"`
	Commodity            string `json:"commodity"`
	Variety              string `json:"variety"`
	ArrivalDate          string `json:"arrival_date"`
	ArrivalDateFormatted string `json:"arrival_date_formatted"`
	MinPrice             int    `json:"min_price"`
	MaxPrice             int    `json:"max_price"`
	ModalPrice           int    `json:"modal_price"`
}

// LoadDataFromJSON loads market data from a JSON file
func LoadDataFromJSON(filePath string) ([]MarketData, error) {
	// Read JSON file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a slice of MarketData
	var data []MarketData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	return data, nil
}
