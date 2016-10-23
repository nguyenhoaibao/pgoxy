package app

import (
	"encoding/json"
	"os"
)

const dataFile = "../data/sites.json"

type Feed struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
