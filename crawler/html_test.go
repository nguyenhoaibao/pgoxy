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
		Type: "html",
		Name: "sslproxies.org",
		Url:  "https://www.sslproxies.org/",
	},
}

func TestGetDocumentWithUrlEmpty(t *testing.T) {
	_, err := crawler.GetDocument("")
	if err == nil {
		t.Fatal("Should receive error when url is empty")
	}
}

func TestGetDocument(t *testing.T) {
	server := mockServer("")
	defer server.Close()

	_, err := crawler.GetDocument(server.URL)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetParserConfigWithNameEmpty(t *testing.T) {
	_, err := crawler.GetParserConfig("")
	if err == nil {
		t.Fatal("Should receive error when name is empty")
	}
}

func TestGetParserConfigWithNameInvalid(t *testing.T) {
	_, err := crawler.GetParserConfig("invalid file")
	if err == nil {
		t.Fatal("Should receive error when name is invalid")
	}
}

func TestGetParserConfig(t *testing.T) {
	_, err := crawler.GetParserConfig(feeds[0].Name)
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
		data, err := c.Crawl()
		if err != nil {
			t.Fatal(err)
		}
		if len(data) <= 0 {
			t.Fatal("Results should be greater than 0")
		}
	}
}
