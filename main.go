package main

import (
	"fmt"
	"log"

	"github.com/imrandil/the_real_world/db"
	"github.com/imrandil/the_real_world/models"
	"github.com/imrandil/the_real_world/utils"
	_ "github.com/lib/pq"
)

func main() {
	serviceURI := "postgres://avnadmin:AVNS_j1zoNT6FNEgqRyqK0Eg@pg-192a0722-aliimranadil2-cf20.a.aivencloud.com:18547/practicedb?sslmode=require"

	db, err := db.ConnectDB(serviceURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//process csv file
	//data.ProcessCSV()

	//drop the table if exists..
	err = models.DropTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// Create the table if it doesn't exist
	err = models.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	data, err := utils.LoadDataFromJSON("output.json")
	if err != nil {
		log.Fatal(err)
	}

	//now insert the data into db

	err = models.InsertMarketDataBulk(db, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully")

}
