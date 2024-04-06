// models/create_table.go

package models

import (
	"database/sql"
	"fmt"
)

// CreateTable creates the market_data table if it doesn't exist
func CreateTable(db *sql.DB) error {
	// SQL statement to create the table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS market_data (
		state TEXT,
		district TEXT,
		market TEXT,
		commodity TEXT,
		variety TEXT,
		arrival_date DATE,
		arrival_date_formatted TIMESTAMP,
		min_price INT,
		max_price INT,
		modal_price INT
	);
	`

	// Execute the SQL statement to create the table
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	fmt.Println("Table 'market_data' created successfully.")
	return nil
}
