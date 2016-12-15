package appengine

import (
	"net/http"
	"github.com/mhaligowski/paperboy-crawler"
)

func init() {
	http.HandleFunc("/handle", crawler.HandleRequest)
}

