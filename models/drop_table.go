// models/create_table.go

package models

import (
	"database/sql"
	"fmt"
)

// CreateTable creates the market_data table if it doesn't exist
func DropTable(db *sql.DB) error {
	// SQL statement to create the table
	dropTableSQL := `
	DROP TABLE IF EXISTS market_data;
	`

	// Execute the SQL statement to create the table
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		return err
	}

	fmt.Println("Table 'market_data' droped successfully.")
	return nil
}
