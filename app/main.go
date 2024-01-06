package main

import (
    "encoding/json"
    "net/http"
)

type HelloWorldResponse struct {
    Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := HelloWorldResponse{Message: "Hello World"}
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/greetings", helloWorldHandler)
    http.ListenAndServe(":8080", nil)
}
