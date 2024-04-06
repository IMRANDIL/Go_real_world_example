package main

import (
	"fmt"
	"log"
	"time"

	"github.com/imrandil/the_real_world/db"
	"github.com/imrandil/the_real_world/models"
	_ "github.com/lib/pq"
)

func main() {
	serviceURI := "postgres://avnadmin:AVNS_j1zoNT6FNEgqRyqK0Eg@pg-192a0722-aliimranadil2-cf20.a.aivencloud.com:18547/practicedb?sslmode=require"
	start := time.Now()
	db, err := db.ConnectDB(serviceURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	count, items, err := models.GetTableCountAndItems(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count of the table", count)
	for _, item := range items {
		fmt.Printf("%+v\n", *item) // Use %+v to print struct field names with values
	}
	// fmt.Printf("data %+v\n", items) // Use %+v to print struct field names with values
	// fmt.Println("Data inserted successfully")
	fmt.Println(time.Since(start))
	fmt.Println("nubmerofrows", len(items))

}
