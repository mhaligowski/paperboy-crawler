package crawler

import (
	"encoding/json"

	"net/http"
	"net/url"

	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
	"google.golang.org/appengine/log"
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

	body, err := json.Marshal(feed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task := taskqueue.NewPOSTTask("/jobs", url.Values{})
	task.Header.Set("Content-Type", "application/json")
	task.Payload = body
	log.Debugf(ctx, "Payload: %q", body)

	taskqueue.Add(ctx, task, "StreamUpdates")

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

