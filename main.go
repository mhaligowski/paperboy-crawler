package crawler

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
	"context"
)

type input struct {
	FeedId  string
	FeedUrl string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if r.Method != http.MethodPost {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}

	input, err := parseInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feed, err := getFeed(ctx, input.FeedUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newEntries, err := writeEntries(ctx, feed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	feed.Entries = newEntries

	// write entries to the queue
	w.WriteHeader(http.StatusOK)
}

func parseInput(r *http.Request) (input, error) {
	feed_id := r.FormValue("feed_id")
	feed_url := r.FormValue("feed_url")

	return input{feed_id, feed_url}, nil
}

func init() {
	http.HandleFunc("/handle", handleRequest)
}

