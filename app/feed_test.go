package app_test

import (
	"testing"

	"github.com/nguyenhoaibao/pgoxy/app"
)

func TestGetFeeds(t *testing.T) {
	feeds, err := app.GetFeeds()
	if err != nil {
		t.Fatal(err)
	}

	if len(feeds) <= 0 {
		t.Fatal("Cannot load any feed")
	}
}
