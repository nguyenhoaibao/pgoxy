package crawler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func mockServer(content string) *httptest.Server {
	handerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		w.WriteHeader(200)
		fmt.Fprintf(w, content)
	}

	return httptest.NewServer(http.HandlerFunc(handerFunc))
}
