package main

import (
    "net/http"
    "fmt"
    "encoding/json"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "OK")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    type Input struct {
        FeedId string `json:"feed_id"`
        FeedUrl string `json:"feed_url"`
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
        return
    }

    var inputs []Input
    err := json.NewDecoder(r.Body).Decode(&inputs)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    for _, input := range inputs {
        fmt.Fprintf(w, "%s", input.FeedUrl)
    }
}

func init() {
    http.HandleFunc("/ping", handlePing)
    http.HandleFunc("/handle", handleRequest)
}

