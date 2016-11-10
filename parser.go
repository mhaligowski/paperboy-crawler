package parser

import (
    "encoding/xml"
    "io"
)

type FeedItem string

func ParseFeed(r io.Reader) []FeedItem {
    d := xml.NewDecoder(r)

    return parse(d)
}

func parse(d *xml.Decoder) []FeedItem {
    return make([]FeedItem, 0, 0)
}
