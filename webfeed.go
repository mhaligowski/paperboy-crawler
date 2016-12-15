package crawler

import (
	"time"
	"encoding/xml"
)

type Entry struct {
	Id      string    `xml:"id" json:"id"`
	Title   string    `xml:"title" json:"title"`
	Updated time.Time `xml:"updated" json:"updated"`
	Content string    `xml:"content" json:"content"`
	Summary string    `xml:"summary" json:"summary"`
}

type WebFeed struct {
	XMLName xml.Name `xml:"feed" json:"-"`

	Id      string    `xml:"id" json:"id"`
	Title   string    `xml:"title" json:"title"`
	Updated time.Time `xml:"updated" json:"updated"`

	Entries []Entry   `xml:"entry" json:"entries"`
}

