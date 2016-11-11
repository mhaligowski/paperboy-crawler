package parser

import "time"

type Entry struct {
    title string
}

type Feed struct {
    id string
    title string
    updated time.Time

    entries []Entry
}
