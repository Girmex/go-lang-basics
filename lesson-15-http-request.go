package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func HttpRequest() {
    http.HandleFunc("/hello", helloHandler) 
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)   
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{"message": "Hello, Go!"}
    json.NewEncoder(w).Encode(response)
}

