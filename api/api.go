package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Embeding is a handler for embedding a text
func Embeding(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Hello World from Go! 👋"
	resp["language"] = "go"
	resp["cloud"] = "Hosted on Vercel! ▲"
	resp["github"] = "https://github.com/riccardogiorato/template-go-vercel/blob/main/api/json.go"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
}
