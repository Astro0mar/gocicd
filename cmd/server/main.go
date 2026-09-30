package main

import (
    "log"
    "net/http"

    "github.com/example/go-cicd-demo/internal/handler"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler.Hello)
    mux.HandleFunc("/health", handler.Health)

    log.Println("server running on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}
