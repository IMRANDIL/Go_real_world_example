// models/util.go

package models

import (
	"database/sql"
	"time"
)

// GetTableCountAndItems retrieves the count of the table and 100 items from the table
func GetTableCountAndItems(db *sql.DB) (int, []MarketData, error) {
	// Query to get the count of the table
	countQuery := "SELECT COUNT(*) FROM market_data"

	// Execute the query to get the count of the table
	var count int
	err := db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return 0, nil, err
	}

	// Query to get 100 items from the table
	itemsQuery := "SELECT state, district, market, commodity, variety, arrival_date, arrival_date_formatted, min_price, max_price, modal_price FROM market_data LIMIT 100"

	// Execute the query to get 100 items from the table
	rows, err := db.Query(itemsQuery)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	// Slice to store the retrieved items
	var items []MarketData

	// Iterate through the rows and scan each row into a MarketData struct
	for rows.Next() {
		var (
			state                string
			district             string
			market               string
			commodity            string
			variety              string
			arrivalDate          time.Time
			arrivalDateFormatted string
			minPrice             int
			maxPrice             int
			modalPrice           int
		)
		err := rows.Scan(&state, &district, &market, &commodity, &variety, &arrivalDate, &arrivalDateFormatted, &minPrice, &maxPrice, &modalPrice)
		if err != nil {
			return 0, nil, err
		}
		// Create a MarketData instance
		item := MarketData{
			State:                state,
			District:             district,
			Market:               market,
			Commodity:            commodity,
			Variety:              variety,
			ArrivalDate:          arrivalDate,
			ArrivalDateFormatted: arrivalDateFormatted,
			MinPrice:             minPrice,
			MaxPrice:             maxPrice,
			ModalPrice:           modalPrice,
		}
		items = append(items, item)
	}

	// Check for any errors during row iteration
	if err := rows.Err(); err != nil {
		return 0, nil, err
	}

	// Return the count of the table and the retrieved items
	return count, items, nil
}
