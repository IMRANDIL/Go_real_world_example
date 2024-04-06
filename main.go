package main

import (
	"log"

	"github.com/imrandil/the_real_world/db"
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

	data, err := utils.LoadDataFromJSON("output.json")
	if err != nil {
		log.Fatal(err)
	}

}
