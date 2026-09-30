package handler

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
    Status  string `json:"status"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(Response{
        Message: "Hello from Go CI/CD!",
        Status:  "success",
    })
}

func Health(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(Response{
        Message: "Service is healthy",
        Status:  "ok",
    })
}
