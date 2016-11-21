package crawler

import (
	"bytes"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"
)

func ParseFeedFromBytes(data []byte) (*Feed, error) {
	reader := bytes.NewReader(data)

	return ParseFeed(reader)
}

func ParseFeed(r io.Reader) (*Feed, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel

	return parse(d)
}

func parse(d *xml.Decoder) (*Feed, error) {
	result := &Feed{}

	e := d.Decode(result)

	if e != nil {
		return nil, e
	}

	return result, nil
}
