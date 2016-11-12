package parser

import (
    "encoding/xml"
    "time"
)

type Entry struct {
    Title string `xml:"title"`
}

type Feed struct {
    XMLName xml.Name `xml:"feed"`

    Id string `xml:"id"`
    Title string `xml:"title"`
    Updated time.Time `xml:"updated"`

    Entries []Entry `xml:"Entry"`
}
