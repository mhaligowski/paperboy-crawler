package crawler

import (
	"strings"
	"testing"
)

const feed = `
<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title type="text">dive into mark</title>
  <subtitle type="html">A &amp;lt;em&amp;gt;lot&amp;lt;/em&amp;gt; of effort
       went into making this effortless</subtitle>
  <updated>2005-07-31T12:29:29Z</updated>
  <id>tag:example.org,2003:3</id>
  <link rel="alternate" type="text/html" hreflang="en" href="http://example.org/" />
  <link rel="self" type="application/atom+xml" href="http://example.org/feed.atom" />
  <rights>Copyright (c) 2003, Mark Pilgrim</rights>
  <generator uri="http://www.example.com/" version="1.0">Example Toolkit</generator>
  <entry>
    <title>Atom draft-07 snapshot</title>
    <link rel="alternate" type="text/html" href="http://example.org/2005/04/02/atom" />
    <link rel="enclosure" type="audio/mpeg" length="1337" href="http://example.org/audio/ph34r_my_podcast.mp3" />
    <id>tag:example.org,2003:3.2397</id>
    <updated>2005-07-31T12:29:29Z</updated>
    <published>2003-12-13T08:29:29-04:00</published>
    <author>
      <name>Mark Pilgrim</name>
      <uri>http://example.org/</uri>
      <email>f8dy@example.com</email>
    </author>
    <contributor>
      <name>Sam Ruby</name>
    </contributor>
    <contributor>
      <name>Joe Gregorio</name>
    </contributor>
    <content type="xhtml" xml:lang="en" xml:base="http://diveintomark.org/">
      <div xmlns="http://www.w3.org/1999/xhtml">
        <p>
          <i>[Update: The Atom draft is finished.]</i>
        </p>
      </div>
    </content>
  </entry>
</feed>
`

func TestCanParseTitle(t *testing.T) {
	r := strings.NewReader(feed)
	v, e := ParseFeed(r)

	if e != nil {
		t.Error("Error while parsing", e.Error())
	}

	if v.Title != "dive into mark" {
		t.Errorf("Expected \"dive into mark\", got [[ %s ]]", v.Title)
	}
}

func TestCanParseUpdated(t *testing.T) {
	r := strings.NewReader(feed)
	v, e := ParseFeed(r)

	if e != nil {
		t.Error("Error while parsing", e.Error())
	}

	if v.Updated.UnixNano() != 1122812969000000000 {
		t.Errorf("Expected time for 1122812969000000000, got [[ %s ]]", v.Updated.UnixNano())
	}
}

func TestThrowsError(t *testing.T) {
	r := strings.NewReader("invalid feed")
	_, e := ParseFeed(r)

	if e == nil {
		t.Error("Didn't encounter error when parsing")
	}
}

func TestEntriesAreFilled(t *testing.T) {
	r := strings.NewReader(feed)
	v, _ := ParseFeed(r)

	if len(v.Entries) == 0 {
		t.Errorf("Expected number of entries to be 1, got [[ %v ]]", len(v.Entries))
	}
}
