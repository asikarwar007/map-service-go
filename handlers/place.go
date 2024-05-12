// handlers/petrol-pump.go
package handlers

import (
	"encoding/json"
	"map-service-go/models"
	"map-service-go/utils"
	"net/http"
)

func DiscoverPlaceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	location := r.URL.Query().Get("location")
	energyType := r.URL.Query().Get("energy-type")
	language := r.URL.Query().Get("lang")
	if language == "" {
		language = "en"
	}
	// category := r.URL.Query().Get("category")

	searchQuery := models.PlaceRequestQuery{
		Location: location,
		Category: getCategory(energyType),
		Language: language,
	}

	petrolPumpResponse, err := utils.SearchPlace(searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petrolPumpResponse)

}
func getCategory(energyType string) string {
	if energyType == "gas" {
		return "700-7600-0116"
	}
	if energyType == "ev" {
		return "700-7600-0322"
	}
	return ""
}
