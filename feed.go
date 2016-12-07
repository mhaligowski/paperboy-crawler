package crawler

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"time"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type Entry struct {
	Id      string    `xml:"id" datastore:"id"`
	Title   string    `xml:"title" datastore:"title"`
	Updated time.Time `xml:"updated" datastore:"updated"`
	Content string    `xml:"content" datastore:"content"`
	Summary string    `xml:"summary" datastore:"summary"`
}

type Feed struct {
	XMLName xml.Name `xml:"feed"`

	Id      string    `xml:"id"`
	Title   string    `xml:"title"`
	Updated time.Time `xml:"updated"`

	Entries []Entry   `xml:"entry"`
}

func fetchFeed(ctx context.Context, feedUrl string) ([]byte, error) {
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

func getFeed(ctx context.Context, url string) (Feed, error) {
	content, err := fetchFeed(ctx, url)

	if err != nil {
		log.Errorf(ctx, "could not fetch feed %s: %v", input.FeedUrl, err)
		return nil, err
	}

	feed, err := parseFeedFromBytes(content)
	if err != nil {
		log.Errorf(ctx, "could not parse feed %s, %v", input.FeedUrl, err)
		return nil, err
	}

	return feed, nil
}
