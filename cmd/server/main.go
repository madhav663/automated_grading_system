package main

import (
	
	"fmt"
	"log"
	"net/http"
	"github.com/madhav663/automated_grading_system/internal/models"
	"github.com/madhav663/automated_grading_system/internal/ocr"
	"github.com/madhav663/automated_grading_system/internal/grading"
	"github.com/madhav663/automated_grading_system/internal/feedback"
	"github.com/madhav663/automated_grading_system/pkg/utils"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	var reportRequest models.ReportRequest
	if err := utils.ValidateRequestBody(w, r, &reportRequest); err != nil {
		return
	}

	ocrResult, err := ocr.OCRProcess(reportRequest.Report.ID)
	if err != nil {
		utils.HandleError(w, err, "Error processing the image", http.StatusInternalServerError)
		return
	}

	phishingResult, err := grading.CheckPhishing(ocrResult)
	if err != nil {
		utils.HandleError(w, err, "Error checking phishing", http.StatusInternalServerError)
		return
	}

	similarityResult, err := grading.CheckSimilarity(ocrResult, "Sample comparison text")
	if err != nil {
		utils.HandleError(w, err, "Error checking similarity", http.StatusInternalServerError)
		return
	}

	feedbackResult, err := feedback.GenerateFeedback(ocrResult)
	if err != nil {
		utils.HandleError(w, err, "Error generating feedback", http.StatusInternalServerError)
		return
	}

	response := models.PhishingCheckResult{
		IsPhishing: phishingResult.IsPhishing,
		Message:    fmt.Sprintf("Similarity Score: %.2f, Feedback: %s", similarityResult.SimilarityScore, feedbackResult.Feedback),
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func RegisterRoutes() {
	http.HandleFunc("/upload", UploadFile)
}

func main() {
	RegisterRoutes()
	log.Println("Server is starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
