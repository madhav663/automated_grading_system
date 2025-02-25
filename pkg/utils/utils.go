package utils

import (
	"log"
	"os"
	"encoding/json"
	"net/http"
)


func Logger(message string, err error) {
	if err != nil {
		log.Printf("[ERROR] %s: %v", message, err)
	} else {
		log.Printf("[INFO] %s", message)
	}
}

func ReadEnvVar(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}


func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}


func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, map[string]string{"error": message})
}


func HandleError(w http.ResponseWriter, err error, errorMessage string, statusCode int) {
	if err != nil {
		Logger(errorMessage, err)
		RespondWithError(w, statusCode, errorMessage)
	}
}


func ValidateRequestBody(w http.ResponseWriter, r *http.Request, requestBody interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
		HandleError(w, err, "Invalid request body", http.StatusBadRequest)
		return err
	}
	return nil
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request) (*os.File, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		HandleError(w, err, "Unable to parse form", http.StatusBadRequest)
		return nil, err
	}

	
	file, _, err := r.FormFile("file")
	if err != nil {
		HandleError(w, err, "Unable to retrieve file", http.StatusBadRequest)
		return nil, err
	}
	defer file.Close()

	
	tempFile, err := os.CreateTemp("uploads", "report-*.png")
	if err != nil {
		HandleError(w, err, "Unable to create temporary file", http.StatusInternalServerError)
		return nil, err
	}
	defer tempFile.Close()


	_, err = tempFile.ReadFrom(file)
	if err != nil {
		HandleError(w, err, "Unable to save file", http.StatusInternalServerError)
		return nil, err
	}

	return tempFile, nil
}
