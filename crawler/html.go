package crawler

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type htmlCrawler struct {
	name string
	doc  *goquery.Document
}

type htmlParserSelector struct {
	Selector string `json:"selector"`
}

type htmlParserData struct {
	IP struct {
		Order  int    `json:"order"`
		Method string `json:"method"`
	} `json:"ip"`
	Port struct {
		Order  int    `json:"order"`
		Method string `json:"method"`
	} `json:"port"`
}

type htmlParser struct {
	Name   string `json:"name"`
	Parser struct {
		Container htmlParserSelector `json:"container"`
		Items     htmlParserSelector `json:"items"`
		Item      htmlParserSelector `json:"item"`
		Data      htmlParserData     `json:"data"`
	} `json:"parser"`
}

func NewHtmlCrawler(name string, d *goquery.Document) *htmlCrawler {
	return &htmlCrawler{name: name, doc: d}
}

func GetDocument(url string) (*goquery.Document, error) {
	if url == "" {
		return nil, errors.New("Url is required")
	}

	doc, err := goquery.NewDocument(url)
	return doc, err
}

func GetParserConfig(name string) (*htmlParser, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	parserFile := "sites" + string(filepath.Separator) + name + ".json"

	f, err := os.Open(parserFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var p htmlParser
	err = json.NewDecoder(f).Decode(&p)

	return &p, err
}

func (c *htmlCrawler) Parse(p *htmlParser) []*Result {
	var results []*Result

	container := c.doc.Find(p.Parser.Container.Selector)
	items := container.Find(p.Parser.Items.Selector)
	items.Each(func(_ int, s *goquery.Selection) {
		ip := strings.TrimSpace(s.Find(p.Parser.Item.Selector).Eq(p.Parser.Data.IP.Order).Text())
		port := strings.TrimSpace(s.Find(p.Parser.Item.Selector).Eq(p.Parser.Data.Port.Order).Text())

		results = append(results, &Result{IP: ip, Port: port})
	})

	return results
}

func (c *htmlCrawler) Crawl() ([]*Result, error) {
	p, err := GetParserConfig(c.name)
	if err != nil {
		return nil, err
	}

	results := c.Parse(p)

	return results, nil
}
