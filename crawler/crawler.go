package crawler

import (
	"log"
	"os"
	"sync"

	"github.com/nguyenhoaibao/pgoxy/app"
)

type Result struct {
	IP   string
	Port string
}

type Crawler interface {
	Crawl() ([]*Result, error)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Run() {
	feeds, err := app.GetFeeds()
	if err != nil {
		log.Fatal(err)
	}
	if len(feeds) <= 0 {
		log.Fatal("Cannot load any feeds")
	}

	results := make(chan *Result)

	var wg sync.WaitGroup
	wg.Add(len(feeds))

	for _, feed := range feeds {
		go func(feed *app.Feed) {
			Crawl(feed, results)
			wg.Done()
		}(feed)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	WriteFile(results)
}

func Crawl(feed *app.Feed, results chan<- *Result) {
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
		data, err := c.Crawl()
		if err != nil {
			log.Println(err)
			return
		}

		for _, r := range data {
			results <- r
		}
	}
}

func WriteFile(results <-chan *Result) {
	f, _ := os.Create("data.txt")

	for result := range results {
		log.Printf("Receive %s:%s", result.IP, result.Port)
		f.WriteString(result.IP + ":" + result.Port + "\n")
	}

	f.Sync()
}
