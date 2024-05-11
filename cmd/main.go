// cmd/main.go
package main

import (
	"log"
	"net/http"
	"map-service-go/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)
const apiKey = "zIe7U6pVVOUlvohqTmqaMJefRH5ujqbjzj4rnctZLJU" 
type SearchResult struct {
	Title       string `json:"title"`
	Vicinity    string `json:"vicinity"`
	Category    string `json:"category"`
}

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/need-energy", handlers.DiscoverPlaceHandler).Methods("GET", "OPTIONS")

	// Apply CORS headers
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(corsMiddleware)

	log.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", r)

}
func corsMiddleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}