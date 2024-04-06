package generate

import (
	"fmt"

	"github.com/imrandil/the_real_world/models"
	"github.com/jung-kurt/gofpdf"
)

// // GeneratePDF generates a PDF file containing the market data records
// func GeneratePDF(items []*models.MarketData) error {
// 	// Initialize PDF
// 	pdf := gofpdf.New("P", "mm", "A4", "")

// 	// Set font
// 	pdf.SetFont("Arial", "", 12)

// 	// Add retrieved data to the PDF
// 	for _, item := range items {
// 		// Add a new page for each record
// 		pdf.AddPage()

// 		// Add content to the PDF
// 		pdf.CellFormat(190, 10, "Market Data Record", "", 0, "C", false, 0, "")
// 		pdf.Ln(10) // Move to the next line

// 		// Define key-value pairs
// 		data := map[string]interface{}{
// 			"State":                    item.State,
// 			"District":                 item.District,
// 			"Market":                   item.Market,
// 			"Commodity":                item.Commodity,
// 			"Variety":                  item.Variety,
// 			"Arrival Date":             item.ArrivalDate,
// 			"Arrival Date (Formatted)": item.ArrivalDateFormatted,
// 			"Min Price":                item.MinPrice,
// 			"Max Price":                item.MaxPrice,
// 			"Modal Price":              item.ModalPrice,
// 		}

// 		// Add each key-value pair to the PDF
// 		for key, value := range data {
// 			// Add key in one column and value in another column
// 			pdf.CellFormat(50, 10, key+":", "1", 0, "", false, 0, "")
// 			pdf.CellFormat(140, 10, fmt.Sprintf("%v", value), "1", 0, "", false, 0, "")
// 			pdf.Ln(10) // Move to the next line
// 		}
// 	}

// 	// Save the PDF file
// 	err := pdf.OutputFileAndClose("market_data.pdf")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// MarketDataRecord represents a single market data record
type MarketDataRecord struct {
	State                string
	District             string
	Market               string
	Commodity            string
	Variety              string
	ArrivalDate          string
	ArrivalDateFormatted string
	MinPrice             int
	MaxPrice             int
	ModalPrice           int
}

// GeneratePDF generates a PDF file containing the market data records
func GeneratePDF(items []*models.MarketData) error {
	// Initialize PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Set font
	pdf.SetFont("Arial", "", 12)

	// Add retrieved data to the PDF
	for _, item := range items {
		// Add a new page for each record
		pdf.AddPage()

		// Add content to the PDF
		pdf.CellFormat(190, 10, "Market Data Record", "", 0, "C", false, 0, "")
		pdf.Ln(10) // Move to the next line

		// Create a MarketDataRecord struct
		record := MarketDataRecord{
			State:                item.State,
			District:             item.District,
			Market:               item.Market,
			Commodity:            item.Commodity,
			Variety:              item.Variety,
			ArrivalDate:          item.ArrivalDate.Format("2006-01-02 15:04:05"),
			ArrivalDateFormatted: item.ArrivalDateFormatted,
			MinPrice:             item.MinPrice,
			MaxPrice:             item.MaxPrice,
			ModalPrice:           item.ModalPrice,
		}

		// Add each field to the PDF in the specified order
		pdf.CellFormat(50, 10, "State:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, record.State, "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "District:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, record.District, "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Market:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, record.Market, "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Commodity:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, record.Commodity, "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Variety:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, record.Variety, "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Arrival Date:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, record.ArrivalDate, "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Arrival Date (Formatted):", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, record.ArrivalDateFormatted, "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Min Price:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, fmt.Sprintf("%d", record.MinPrice), "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Max Price:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, fmt.Sprintf("%d", record.MaxPrice), "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
		pdf.CellFormat(50, 10, "Modal Price:", "1", 0, "", false, 0, "")
		pdf.CellFormat(140, 10, fmt.Sprintf("%d", record.ModalPrice), "1", 0, "", false, 0, "")
		pdf.Ln(10) // Move to the next line
	}

	// Save the PDF file
	err := pdf.OutputFileAndClose("market_data.pdf")
	if err != nil {
		return err
	}

	return nil
}
