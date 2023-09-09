package handlers

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
	"github.com/jung-kurt/gofpdf"
)

func JsonToPdf(fileName string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "N", 14)
	pageWidth := 210.0  // Width of A4 page in mm
	pageHeight := 297.0 // Height of A4 page in mm

	data, err := os.ReadFile(fileName)

	errorhandlers.HandleError(err)
	fileDataJson := JsonSchema{}
	err = json.Unmarshal(data, &fileDataJson)
	errorhandlers.HandleError(err)

	// Initialize min and max coordinates
	minX := pageWidth
	maxX := 0.0
	minY := pageHeight
	maxY := 0.0

	// Loop through the JSON data to find the min and max coordinates
	for _, item := range fileDataJson.List {
		bboxCoords := strings.Split(item.Box, ",")
		x, _ := strconv.ParseFloat(bboxCoords[0], 64)
		y, _ := strconv.ParseFloat(bboxCoords[1], 64)

		// Update min and max coordinates
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	maxX = maxX + 500
	maxY = maxY + 500

	for _, item := range fileDataJson.List {
		bboxCoords := strings.Split(item.Box, ",")
		x, _ := strconv.ParseFloat(bboxCoords[0], 64)
		y, _ := strconv.ParseFloat(bboxCoords[1], 64)

		// Calculate the interpolated position of the text relative to the page
		interpolatedX := pageWidth * ((x - minX) / (maxX - minX))
		interpolatedY := pageHeight * ((y - minY) / (maxY - minY))

		// Add text to the PDF
		pdf.Text(interpolatedX, interpolatedY, item.Text)

	}

	pdf.OutputFileAndClose("output.pdf")
}
