package main

import (
    "net/http"
    "fmt"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "OK")
}

func init() {
    http.HandleFunc("/ping", handlePing)
}

