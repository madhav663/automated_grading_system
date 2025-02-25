package models

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"bytes"
	"os"
)

type Report struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type PhishingCheckResult struct {
	IsPhishing bool   `json:"is_phishing"`
	Message    string `json:"message"`
}

type SimilarityCheckResult struct {
	SimilarityScore float64 `json:"similarity_score"`
	Message         string  `json:"message"`
}

type Feedback struct {
	Text     string `json:"text"`
	Feedback string `json:"feedback"`
}

type ReportRequest struct {
	Report Report `json:"report"`
}

type FeedbackRequest struct {
	Text string `json:"text"`
}

type OCRResult struct {
	Text string `json:"text"`
}

func CheckPhishing(text string) (*PhishingCheckResult, error) {
	apiURL := os.Getenv("OLLAMA_API_URL") + "/check_phishing"
	apiKey := os.Getenv("OLLAMA_API_KEY")

	requestData := map[string]string{"text": text}
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result PhishingCheckResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	return &result, nil
}

func CheckSimilarity(text1, text2 string) (*SimilarityCheckResult, error) {
	apiURL := os.Getenv("OLLAMA_API_URL") + "/check_similarity"
	apiKey := os.Getenv("OLLAMA_API_KEY")

	requestData := map[string]string{"text1": text1, "text2": text2}
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result SimilarityCheckResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	return &result, nil
}

func GenerateFeedback(text string) (*Feedback, error) {
	apiURL := os.Getenv("OLLAMA_API_URL") + "/generate_feedback"
	apiKey := os.Getenv("OLLAMA_API_KEY")

	requestData := map[string]string{"text": text}
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var feedback Feedback
	if err := json.NewDecoder(resp.Body).Decode(&feedback); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	return &feedback, nil
}

func OCRProcess(reportID string) (*OCRResult, error) {
	ocrText := fmt.Sprintf("Extracted text for report ID: %s", reportID)
	return &OCRResult{Text: ocrText}, nil
}
