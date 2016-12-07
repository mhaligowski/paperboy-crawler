package crawler

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
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

	body, err := fetchFeed(r, input.FeedUrl)
	if err != nil {
		log.Errorf(ctx, "could not fetch feed %s: %v", input.FeedUrl, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	feed, err := ParseFeedFromBytes(body)
	if err != nil {
		log.Errorf(ctx, "could not parse feed %s, %v", input.FeedUrl, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, entry := range feed.Entries {
		entry.Summary = entry.Summary[:1500]

		_, _, err := addEntryIfDoesntExist(ctx, entry)
		if err != nil {
			log.Errorf(ctx, "could not put feed entry %s from feed %s: %v", entry.Id, feed.Id, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func fetchFeed(r *http.Request, feedUrl string) ([]byte, error) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	response, err := client.Get(feedUrl)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseInput(r *http.Request) (input, error) {
	feed_id := r.FormValue("feed_id")
	feed_url := r.FormValue("feed_url")

	return input{feed_id, feed_url}, nil
}

func init() {
	http.HandleFunc("/handle", handleRequest)
}
