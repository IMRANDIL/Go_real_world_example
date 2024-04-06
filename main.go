package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/imrandil/the_real_world/db"
	"github.com/imrandil/the_real_world/generate"
	"github.com/imrandil/the_real_world/models"
	_ "github.com/lib/pq"
)

func fanIn(done <-chan int, channels ...<-chan models.QueryResult) <-chan models.QueryResult {
	var wg sync.WaitGroup
	fannedInStream := make(chan models.QueryResult)

	transfer := func(c <-chan models.QueryResult) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream

}

func main() {
	serviceURI := "postgres://avnadmin:AVNS_j1zoNT6FNEgqRyqK0Eg@pg-192a0722-aliimranadil2-cf20.a.aivencloud.com:18547/practicedb?sslmode=require"
	start := time.Now()
	db, err := db.ConnectDB(serviceURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	done := make(chan int)

	defer close(done)

	//process csv file
	//data.ProcessCSV()

	//drop the table if exists..
	// err = models.DropTable(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create the table if it doesn't exist
	// err = models.CreateTable(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, err := utils.LoadDataFromJSON("output.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//now insert the data into db

	// err = models.InsertMarketDataBulk(db, data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Calculate the number of rows each goroutine should process

	CPUCount := runtime.NumCPU()
	count, err := models.GetCountOfTable(db)
	if err != nil {
		log.Fatal(err)
	}

	rowsPerRoutine := count / CPUCount
	remainingRows := count % CPUCount // Distribute the remaining rows among the goroutines

	startRow := 0
	dataRetrievalChannels := make([]<-chan models.QueryResult, CPUCount)

	// Spawn goroutines to retrieve data from the database
	for i := 0; i < CPUCount; i++ {
		// Calculate the end row for this goroutine
		endRow := startRow + rowsPerRoutine
		if i < remainingRows {
			endRow++ // Distribute remaining rows among the first few goroutines
		}

		// Retrieve data from the specified range of rows
		dataRetrievalChannels[i] = models.GetTableCountAndItems(db, startRow, endRow, done)

		// Update start row for the next goroutine
		startRow = endRow
	}

	// Fan in results from all goroutines
	// Collect items from all goroutines
	var allItems []*models.MarketData
	for result := range fanIn(done, dataRetrievalChannels...) {
		if result.Err != nil {
			fmt.Println("Error:", result.Err)
			continue
		}
		allItems = append(allItems, result.Items...)
		// items := result.Items

		// for _, item := range items {
		// 	fmt.Printf("%+v\n", *item)
		// }
		// fmt.Println("length of items", len(items))
	}

	// Generate PDF with all retrieved items
	err = generate.GeneratePDF(allItems)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("data %+v\n", items) // Use %+v to print struct field names with values
	// fmt.Println("Data inserted successfully")
	fmt.Println(time.Since(start))
	fmt.Println("count row", count)
	fmt.Println("PDF generated successfully")
	// fmt.Println("nubmerofrows", len(items))

}
