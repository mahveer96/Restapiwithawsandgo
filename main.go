package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
