package data

import (
	"time"
)

type MarketData struct {
	State                string    `json:"state"`
	District             string    `json:"district"`
	Market               string    `json:"market"`
	Commodity            string    `json:"commodity"`
	Variety              string    `json:"variety"`
	ArrivalDate          time.Time `json:"arrival_date"`
	ArrivalDateFormatted string    `json:"arrival_date_formatted"` // New field for formatted date
	MinPrice             int       `json:"min_price"`
	MaxPrice             int       `json:"max_price"`
	ModalPrice           int       `json:"modal_price"`
}

// func ProcessCSV() {
// 	// Open the CSV file
// 	csvFile, err := os.Open("csv")
// 	if err != nil {
// 		log.Fatalf("Error opening CSV file: %s", err)
// 	}
// 	defer csvFile.Close()

// 	// Read CSV file
// 	reader := csv.NewReader(csvFile)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatalf("Error reading CSV file: %s", err)
// 	}

// 	// Create a slice of MarketData
// 	var data []MarketData

// 	// Convert CSV records to struct
// 	for i, record := range records {
// 		// Skip header row
// 		if i == 0 {
// 			continue
// 		}

// 		// Adjust the CSV parsing code to parse the arrival date as time.Time
// 		arrivalDate, _ := time.Parse("01/02/2006", record[5]) // Assuming the date format is MM/DD/YYYY
// 		minPrice, _ := strconv.Atoi(record[6])                // Convert min_price string to int
// 		maxPrice, _ := strconv.Atoi(record[7])                // Convert max_price string to int
// 		modalPrice, _ := strconv.Atoi(record[8])              // Convert modal_price string to int

// 		item := MarketData{
// 			State:                record[0],
// 			District:             record[1],
// 			Market:               record[2],
// 			Commodity:            record[3],
// 			Variety:              record[4],
// 			ArrivalDate:          arrivalDate,
// 			ArrivalDateFormatted: arrivalDate.Format("2006-01-02 15:04:05 IST"), // Format the arrival date
// 			MinPrice:             minPrice,
// 			MaxPrice:             maxPrice,
// 			ModalPrice:           modalPrice,
// 		}
// 		data = append(data, item)
// 	}

// 	// Create JSON file
// 	jsonFile, err := os.Create("output.json")
// 	if err != nil {
// 		log.Fatalf("Error creating JSON file: %s", err)
// 	}
// 	defer jsonFile.Close()

// 	// Encode data to JSON and write to file
// 	encoder := json.NewEncoder(jsonFile)
// 	err = encoder.Encode(data)
// 	if err != nil {
// 		log.Fatalf("Error encoding JSON: %s", err)
// 	}

// 	log.Println("Data successfully converted to JSON and stored in output.json")
// }
