// model/market_data.go

package models

import (
	"database/sql"

	"github.com/imrandil/the_real_world/data"
)

// InsertMarketDataBulk inserts market data into the database in bulk
func InsertMarketDataBulk(db *sql.DB, data []data.MarketData) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Prepare the bulk insert statement
	stmt, err := tx.Prepare("INSERT INTO market_data (state, district, market, commodity, variety, arrival_date, arrival_date_formatted, min_price, max_price, modal_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Batch size for bulk insert
	batchSize := 1000
	batch := make([]data.MarketData, 0, batchSize)

	// Insert records in batches
	for i, d := range data {
		batch = append(batch, d)
		if len(batch) == batchSize || i == len(data)-1 {
			// Execute bulk insert
			err := insertBatch(stmt, batch)
			if err != nil {
				return err
			}
			// Reset batch slice
			batch = batch[:0]
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// insertBatch inserts a batch of records into the database
// insertBatch inserts a batch of records into the database
func insertBatch(db *sql.DB, stmt *sql.Stmt, batch []data.MarketData) error {
	// Start a transaction for the batch
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert each record in the batch
	for _, d := range batch {
		_, err := tx.Stmt(stmt).Exec(d.State, d.District, d.Market, d.Commodity, d.Variety, d.ArrivalDate, d.ArrivalDateFormatted, d.MinPrice, d.MaxPrice, d.ModalPrice)
		if err != nil {
			return err
		}
	}

	// Commit the transaction for the batch
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
