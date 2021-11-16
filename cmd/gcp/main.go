// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/morzhanov/go-serverless/internal"
)

var s = internal.NewService()

func Handle(w http.ResponseWriter, r *http.Request) {
	var d internal.Event
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		b, _ := json.Marshal(&d)
		s.Handle(b)
	}
	fmt.Fprint(w, html.EscapeString(d.Name))
}
