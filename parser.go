package crawler

import (
	"bytes"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"
)

func parseFeedFromBytes(data []byte) (*WebFeed, error) {
	reader := bytes.NewReader(data)

	return parseFeed(reader)
}

func parseFeed(r io.Reader) (*WebFeed, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel

	return parse(d)
}

func parse(d *xml.Decoder) (*WebFeed, error) {
	result := &WebFeed{}

	e := d.Decode(result)

	if e != nil {
		return nil, e
	}

	return result, nil
}
