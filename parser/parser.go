package parser

import (
    "encoding/xml"
    "io"
)

func ParseFeed(r io.Reader) (feed *Feed, err error) {
    d := xml.NewDecoder(r)

    return parse(d)
}

func parse(d *xml.Decoder) (feed *Feed, err error) {
    result := &Feed{}

    e := d.Decode(result)

    if e != nil {
        return nil, e
    }

    return result, nil
}

