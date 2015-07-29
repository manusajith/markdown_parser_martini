package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	m := martini.Classic()

	m.Post("/markdown", func(req *http.Request) []byte {
		body := req.FormValue("content")
		unsafe := blackfriday.MarkdownBasic([]byte(body))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		return html
	})

	m.Run()
}
