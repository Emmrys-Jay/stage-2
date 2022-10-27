package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Username string `json:"slackUsername"`
	Backend  bool   `json:"backend"`
	Age      int    `json:"age"`
	Bio      string `json:"bio"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bio", response)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func response(w http.ResponseWriter, r *http.Request) {
	response := Response{
		"Emmrys",
		true,
		20,
		"Student who loves learning and is currently exploring backend engineering",
	}

	result, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintln(w, "Error: couldn't marshal json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintln(w, string(result))
}
