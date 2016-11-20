package main

import (
    "net/http"
    "fmt"
    "io/ioutil"

    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"

    "parser"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "OK")
}

type input struct {
    FeedId string
    FeedUrl string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
        return
    }

    input, err := parseInput(r); if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    body, err := fetchFeed(r, input.FeedUrl); if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = parser.ParseFeedFromBytes(body); if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func fetchFeed(r *http.Request, feedUrl string) ([]byte, error) {
    ctx := appengine.NewContext(r)
    client := urlfetch.Client(ctx)
    response, err := client.Get(feedUrl); if err != nil {
        return nil, err
    }

    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body); if err != nil {
        return nil, err
    }

    return body, nil
}

func parseInput(r *http.Request) (i input, err error) {
    feed_id := r.FormValue("feed_id")
    feed_url := r.FormValue("feed_url")

    return input{feed_id, feed_url}, nil
}

func init() {
    http.HandleFunc("/ping", handlePing)
    http.HandleFunc("/handle", handleRequest)
}

