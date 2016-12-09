package crawler

import (
	"golang.org/x/net/context"
	"encoding/xml"
	"io/ioutil"
	"time"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type Entry struct {
	Id      string    `xml:"id" json:"id"`
	Title   string    `xml:"title" json:"title"`
	Updated time.Time `xml:"updated" json:"updated"`
	Content string    `xml:"content" json:"content"`
	Summary string    `xml:"summary" json:"summary"`
}

type Feed struct {
	XMLName xml.Name `xml:"feed" json:"-"`

	Id      string    `xml:"id" json:"id"`
	Title   string    `xml:"title" json:"title"`
	Updated time.Time `xml:"updated" json:"updated"`

	Entries []Entry   `xml:"entry" json:"entries"`
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

func getFeed(ctx context.Context, url string) (*Feed, error) {
	content, err := fetchFeed(ctx, url)

	if err != nil {
		log.Errorf(ctx, "could not fetch feed %s: %v", url, err)
		return nil, err
	}

	feed, err := parseFeedFromBytes(content)
	if err != nil {
		log.Errorf(ctx, "could not parse feed %s, %v", url, err)
		return nil, err
	}

	return feed, nil
}
