package parser

import (
    "golang.org/x/net/html/charset"
    "encoding/xml"
    "bytes"
    "io"
)

func ParseFeedFromBytes(data []byte) (feed *Feed, err error) {
    reader := bytes.NewReader(data)

    return ParseFeed(reader)
}

func ParseFeed(r io.Reader) (feed *Feed, err error) {
    d := xml.NewDecoder(r)
    d.CharsetReader = charset.NewReaderLabel

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

