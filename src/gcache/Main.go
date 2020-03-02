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
	r.HandleFunc("/item/{id}", findCacheByID).Methods("GET")
	r.HandleFunc("/item", addToCache).Methods("POST")
	r.HandleFunc("/item/{id}", deleteCacheBydID).Methods("DELETE")
	r.HandleFunc("/items", getAllItems).Methods("GET")

	http.ListenAndServe(":8081", r)
}

func welcomeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GCache")
}

func deleteCacheBydID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var newItems []Item = []Item{}
	var deletedItem Item
	for _, item := range items {
		if item.ID != id {
			newItems = append(newItems, item)
		} else {
			deletedItem = item
		}
	}
	items = newItems
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedItem)
}

func findCacheByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
		}
	}
}

func addToCache(w http.ResponseWriter, r *http.Request) {
	// Get request data as JSON from JSON body
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)

	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(items)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
