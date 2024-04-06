// model/market_data.go

package models

import (
	"database/sql"
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

// InsertMarketData inserts market data into the database
func InsertMarketData(db *sql.DB, data []MarketData) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Prepare the insert statement
	stmt, err := tx.Prepare("INSERT INTO market_data (state, district, market, commodity, variety, arrival_date, arrival_date_formatted, min_price, max_price, modal_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert each record
	for _, d := range data {
		_, err := stmt.Exec(d.State, d.District, d.Market, d.Commodity, d.Variety, d.ArrivalDate, d.ArrivalDateFormatted, d.MinPrice, d.MaxPrice, d.ModalPrice)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
