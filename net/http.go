package net

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHTML(url string) (string, error) {
	if url == "" {
		return "", errors.New("Url is required")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Url %s response error code %d\n", url, resp.StatusCode)
	}

	html, err := ioutil.ReadAll(resp.Body)

	return string(html), err
}
