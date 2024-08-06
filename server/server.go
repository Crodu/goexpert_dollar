package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Exchange struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

// StartServer starts the server.
func StartServer() {
	fmt.Println("Starting server...")

	_, err := SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database setup successful.")

	http.HandleFunc("/cotacao", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server up and running on port 8080.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		http.NotFound(w, r)
		return
	}
	var exchange Exchange
	err := GetExchange(&exchange)
	if err != nil {
		fmt.Fprintf(w, "Error getting exchange rate: %v", err)
		return
	}
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchange)

	db, cancel, err := DatabaseConnection()
	defer cancel()

	if err != nil {
		fmt.Fprintf(w, "Error connecting to database: %v", err)
		return
	}
	InsertData(db, exchange.Usdbrl)

}
