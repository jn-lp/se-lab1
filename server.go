package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response format
type indexRes struct {
	Time string `json:"time"`
}

// Handler for '/' URI
func indexHandler(w http.ResponseWriter, r *http.Request) {
// Filling response
	res := indexRes{time.Now().Format(time.RFC3339)}
// Serializing response
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
// Setting headers
	w.Header().Set("Content-Type", "application/json")
// Writing headers
	w.WriteHeader(http.StatusAccepted)
// Writing responce
	w.Write(jsonRes)
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8795", nil)
}
