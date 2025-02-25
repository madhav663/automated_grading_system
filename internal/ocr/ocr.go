package ocr

import (
	"github.com/otiai10/gosseract/v2"
	"log"
)

// OCRProcess simulates the OCR process and returns the text content
func OCRProcess(reportID string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	// Simulate processing logic (normally, you'd use the reportID to get the file)
	err := client.SetImage(reportID)
	if err != nil {
		log.Println("Error setting image:", err)
		return "", err
	}

	text, err := client.Text()
	if err != nil {
		log.Println("Error extracting text:", err)
		return "", err
	}
	return text, nil
}
