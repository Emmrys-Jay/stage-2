package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Emmrys-Jay/stage-2/parser"
	"github.com/gorilla/mux"
)

type RequestBody struct {
	OperationType string `json:"operation_type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
}

type ResponseBody struct {
	Username      string `json:"slackUsername"`
	OperationType string `json:"operation_type"`
	Result        int    `json:"result"`
}

func main() {
	// port := os.Getenv("PORT")
	router := mux.NewRouter()

	router = router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/", computeResult)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func computeResult(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		fmt.Fprintf(w, "error:  %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	operation, result := parser.ParseRequest(reqBody.OperationType, reqBody.X, reqBody.Y)

	response := ResponseBody{
		Username:      "Emmrys",
		Result:        result,
		OperationType: operation,
	}

	respJson, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Fprintf(w, "error:  %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}
