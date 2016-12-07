package crawler

import (
	"bytes"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"
)

func parseFeedFromBytes(data []byte) (*Feed, error) {
	reader := bytes.NewReader(data)

	return parseFeed(reader)
}

func parseFeed(r io.Reader) (*Feed, error) {
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
