// handlers/petrol-pump.go
package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"net/url"
	"map-service-go/models"
	"map-service-go/utils"
)

func DiscoverPlaceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	location := r.URL.Query().Get("location")
	escapedQuery := url.QueryEscape(r.URL.Query().Get("q"))

	searchQuery := models.PlaceRequestQuery{
		Location: location,
		Query: escapedQuery,
	}
	log.Println("searchQuery: %s",searchQuery)

	petrolPumpResponse, err := utils.SearchPlace(searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petrolPumpResponse)

}
