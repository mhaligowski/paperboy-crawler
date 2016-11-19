package main

import (
    "net/http"
    "fmt"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "OK")
}

type Input struct {
    FeedId string
    FeedUrl string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
        return
    }

    input, err := parseInput(r)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    fmt.Fprintf(w, "%s", input.FeedUrl)
}

func parseInput(r *http.Request) (i Input, err error) {
    feed_id := r.FormValue("feed_id")
    feed_url := r.FormValue("feed_url")

    return Input{feed_id, feed_url}, nil
}

func init() {
    http.HandleFunc("/ping", handlePing)
    http.HandleFunc("/handle", handleRequest)
}

