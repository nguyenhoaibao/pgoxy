package crawler

import (
	"log"

	"github.com/nguyenhoaibao/pgoxy/app"
	"github.com/nguyenhoaibao/pgoxy/net"
)

type htmlCrawler struct{}

func init() {
	var crawler htmlCrawler
	app.Register("html", crawler)
}

func (c htmlCrawler) Crawl(feed *app.Feed) ([]*app.Result, error) {
	html, err := net.GetHTML(feed.Url)
	if err != nil {
		log.Fatalln(err)
	}

	c.parse(feed, html)

	return nil, nil
}

func (c htmlCrawler) parse(feed *app.Feed, html string) ([]*app.Result, error) {
	return nil, nil
}
