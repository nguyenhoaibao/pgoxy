package crawler_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/nguyenhoaibao/pgoxy/app"
	"github.com/nguyenhoaibao/pgoxy/crawler"
)

var feeds = []*app.Feed{
	&app.Feed{
		Name: "proxylist.hidemyass.com",
		Url:  "http://proxylist.hidemyass.com",
		Type: "html",
	},
}

func TestGetDocument(t *testing.T) {
	_, err := crawler.GetDocument("")
	if err == nil {
		t.Fatal("Should receive error when url is empty")
	}

	server := mockServer("")
	defer server.Close()

	_, err = crawler.GetDocument(server.URL)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawl(t *testing.T) {
	for _, feed := range feeds {
		name := feed.Name
		htmlContent, err := ioutil.ReadFile("testdata" + string(filepath.Separator) + name + ".html")
		if err != nil {
			t.Fatal(err)
		}

		server := mockServer(string(htmlContent))
		defer server.Close()

		d, err := crawler.GetDocument(server.URL)
		if err != nil {
			t.Fatal(err)
		}

		c := crawler.NewHtmlCrawler(name, d)
		results, err := c.Crawl()
		if err != nil {
			t.Fatal(err)
		}
		if len(results) <= 0 {
			t.Fatalf("Cannot get any results")
		}
	}
}
