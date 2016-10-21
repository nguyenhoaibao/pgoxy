package app

type Result struct {
	IP   string
	Port string
}

type Crawler interface {
	Crawl(feed *Feed) ([]*Result, error)
}

func Crawl(crawler Crawler, feed *Feed) {
	crawler.Crawl(feed)
}
