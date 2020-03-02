package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

// Item is a group of properties
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item = []Item{}

func main() {
	fmt.Println("Welcome to Gcache")
	r := mux.NewRouter()
	r.HandleFunc("/", welcomeFunc)
	r.HandleFunc("/{id}", findCacheByID).Methods("GET")
	r.HandleFunc("/add", addToCache).Methods("POST")

	http.ListenAndServe(":8081", r)
}

func welcomeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GCache")
}

func findCacheByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
}

func addToCache(w http.ResponseWriter, r *http.Request) {
	// Get request data as JSON from JSON body
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)

	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(items)
}
