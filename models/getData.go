package models

import (
	"database/sql"
	"runtime"
	"sync"
)

// GetTableCountAndItems retrieves the count of the table and 100 items from the table
func GetTableCountAndItems(db *sql.DB) (int, []*MarketData, error) {
	// Query to get the count of the table
	countQuery := "SELECT COUNT(*) FROM market_data"

	// Execute the query to get the count of the table asynchronously
	var count int
	countErrChan := make(chan error, 1)
	go func() {
		err := db.QueryRow(countQuery).Scan(&count)
		countErrChan <- err
	}()

	// Query to get 100 items from the table
	itemsQuery := "SELECT state, district, market, commodity, variety, arrival_date, arrival_date_formatted, min_price, max_price, modal_price FROM market_data order by max_price desc"

	// Execute the query to get 100 items from the table
	rows, err := db.Query(itemsQuery)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	// Slice to store the retrieved items
	var items []*MarketData
	var mu sync.Mutex // Mutex to synchronize access to the items slice
	CPUCounts := runtime.NumCPU()
	// Use a semaphore to limit the number of concurrent goroutines
	semaphore := make(chan struct{}, CPUCounts) // Adjust the value as needed

	// Iterate through the rows and scan each row into a MarketData struct asynchronously
	for rows.Next() {
		semaphore <- struct{}{} // Acquire semaphore
		go func() {
			defer func() {
				<-semaphore // Release semaphore
			}()

			item := &MarketData{} // Initialize as a pointer to MarketData
			err := rows.Scan(&item.State, &item.District, &item.Market, &item.Commodity, &item.Variety, &item.ArrivalDate, &item.ArrivalDateFormatted, &item.MinPrice, &item.MaxPrice, &item.ModalPrice)
			if err != nil {
				countErrChan <- err
				return
			}
			mu.Lock()
			defer mu.Unlock()
			items = append(items, item)
		}()
	}

	// Wait for the count query to finish
	err = <-countErrChan
	if err != nil {
		return 0, nil, err
	}

	// Return the count of the table and the retrieved items
	return count, items, nil
}
