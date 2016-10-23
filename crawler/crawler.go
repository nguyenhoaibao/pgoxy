package crawler

import (
	"log"

	"github.com/nguyenhoaibao/pgoxy/app"
)

type Result struct {
	IP   string
	Port string
}

type Crawler interface {
	Crawl() ([]*Result, error)
}

func Run() {
	feeds, err := app.GetFeeds()
	if err != nil {
		log.Fatalln(err)
	}

	for _, feed := range feeds {
		Crawl(feed)
	}
}

func Crawl(feed *app.Feed) {
	var c Crawler

	switch feed.Type {
	case "html":
		name := feed.Name
		url := feed.Url

		d, err := GetDocument(url)
		if err != nil {
			log.Fatal(err)
		}

		c = NewHtmlCrawler(name, d)
		c.Crawl()
	}
}
