package app

import "log"

var crawlers = make(map[string]Crawler)

func Run() {
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatalln(err)
	}

	for _, feed := range feeds {
		crawler, exists := crawlers[feed.Type]

		if !exists {
			log.Fatalln("Type %s does not exist", feed.Type)
		}

		Crawl(crawler, feed)
	}
}

func Register(name string, crawler Crawler) {
	if _, exists := crawlers[name]; exists {
		log.Fatalf("%s already exists", name)
	}

	crawlers[name] = crawler
}
