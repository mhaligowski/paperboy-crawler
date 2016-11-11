package parser

import (
    "encoding/xml"
    "io"
)


func ParseFeed(r io.Reader) *Feed {
    d := xml.NewDecoder(r)

    return parse(d)
}

func parse(d *xml.Decoder) *Feed {
    return new(Feed)
}
