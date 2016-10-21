package app

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const dataDir = "data"

type Feed struct {
	Type          string `json:"type"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	CrawlingPages int    `json:"crawling_pages"`
}

func GetFeeds() ([]*Feed, error) {
	d, err := os.Open(dataDir)
	if err != nil {
		return nil, err
	}

	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var feeds []*Feed

	for _, file := range files {
		if file.Mode().IsRegular() {
			fileData, err := os.Open(dataDir + string(filepath.Separator) + file.Name())
			if err != nil {
				return nil, err
			}

			var feed Feed

			err = json.NewDecoder(fileData).Decode(&feed)
			if err != nil {
				return nil, err
			}

			feeds = append(feeds, &feed)
		}
	}

	return feeds, nil
}
