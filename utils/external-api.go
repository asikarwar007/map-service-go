package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"map-service-go/models" // Adjust the import path according to your project structure
	"net/http"
	"os"
)

// FetchGenerateCaptcha calls the external captcha service and returns the response.
func SearchPlace(requestData models.PlaceRequestQuery) (*models.PlaceResponse, error) {
	apiBaseURL := os.Getenv("BROWSE_API_URL")
	apiKey := os.Getenv("apiKey")
	apiUrl := fmt.Sprintf("%s?at=%s&categories=%s&limit=5&lang=%s&apiKey=%s", apiBaseURL, requestData.Location, requestData.Category, requestData.Language, apiKey)

	resp, err := http.Get(apiUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return nil, HandleHTTPErrorNew(resp)
	}

	var placeResponse models.PlaceResponse
	if err := json.NewDecoder(resp.Body).Decode(&placeResponse); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &placeResponse, nil
}
