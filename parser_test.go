package parser

import (
    "testing"
    "strings"
)

func TestInitial(t *testing.T) {
    r := strings.NewReader("test")
    v := ParseFeed(r)

    if len(v) != 0 {
        t.Error("Expected empty slice, got ", len(v))
    }
}

