// models/util.go

package models

import (
	"database/sql"
	"log"
)

// QueryResult holds the result of a database query
type QueryResult struct {
	Items []*MarketData
	Err   error
}

func GetCountOfTable(db *sql.DB) (int, error) {
	// Query to get the count of the table
	countQuery := "SELECT COUNT(*) FROM market_data"
	var count int
	err := db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count, nil
}

// GetTableCountAndItems retrieves the count of the table and items from the specified range of rows
func GetTableCountAndItems(db *sql.DB, startRow, endRow int, done <-chan int) <-chan QueryResult {
	resultChan := make(chan QueryResult)

	go func() {
		defer close(resultChan)

		// // Query to get the count of the table
		// countQuery := "SELECT COUNT(*) FROM market_data"
		// var count int
		// err := db.QueryRow(countQuery).Scan(&count)
		// if err != nil {
		// 	resultChan <- QueryResult{Err: err}
		// 	return
		// }

		// Query to get items from the specified range of rows
		itemsQuery := "SELECT state, district, market, commodity, variety, arrival_date, arrival_date_formatted, min_price, max_price, modal_price FROM market_data ORDER BY max_price DESC OFFSET $1 LIMIT $2"
		rows, err := db.Query(itemsQuery, startRow, endRow-startRow)

		if err != nil {
			resultChan <- QueryResult{Err: err}
			return
		}
		defer rows.Close()

		// Slice to store the retrieved items
		var items []*MarketData

		// Iterate through the rows and scan each row into a MarketData struct
		for rows.Next() {
			item := &MarketData{} // Initialize as a pointer to MarketData
			err := rows.Scan(&item.State, &item.District, &item.Market, &item.Commodity, &item.Variety, &item.ArrivalDate, &item.ArrivalDateFormatted, &item.MinPrice, &item.MaxPrice, &item.ModalPrice)
			if err != nil {
				resultChan <- QueryResult{Err: err}
				return
			}
			items = append(items, item)
		}

		// Check for any errors during row iteration
		if err := rows.Err(); err != nil {
			resultChan <- QueryResult{Err: err}
			return
		}

		// Send the count and items over the result channel
		resultChan <- QueryResult{Items: items}
	}()

	return resultChan
}
