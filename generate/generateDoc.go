// marketdata/generate/pdf.go

package generate

import (
	"fmt"

	"github.com/imrandil/the_real_world/models"
	"github.com/jung-kurt/gofpdf"
)

// GeneratePDF generates a PDF file containing the market data records
func GeneratePDF(items []*models.MarketData) error {
	// Initialize PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "", 12)

	// Add content to the PDF
	pdf.Cell(40, 10, "Market Data Records")
	pdf.Ln(-1) // Move to the next line

	// Add retrieved data to the PDF
	for _, item := range items {
		row := fmt.Sprintf("State: %s, District: %s, Market: %s, Commodity: %s, Variety: %s, Arrival Date: %s, Min Price: %d, Max Price: %d, Modal Price: %d",
			item.State, item.District, item.Market, item.Commodity, item.Variety, item.ArrivalDate, item.MinPrice, item.MaxPrice, item.ModalPrice)
		pdf.CellFormat(0, 10, row, "", 1, "L", false, 0, "")
	}

	// Save the PDF file
	err := pdf.OutputFileAndClose("market_data.pdf")
	if err != nil {
		return err
	}

	return nil
}
