package crawler

import (
	"encoding/xml"
	"time"
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

	Entries []Entry `xml:"entry"`
}
