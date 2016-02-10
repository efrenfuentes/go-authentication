package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type hello struct {
	Message string `json:"message"`
}

func HelloIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	message := hello{"Hello, World!"}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))
}

func HelloName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	name := urlParams["name"]

	HelloMessage := "Hello, " + name
	message := hello{HelloMessage}

	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))
}
