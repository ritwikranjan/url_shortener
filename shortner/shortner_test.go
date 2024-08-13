package shortner_test

import (
	"testing"
	"url_shortener/db"
	"url_shortener/shortner"
)

func TestShorten(t *testing.T) {
	database := db.NewDataBase()
	lUri := "http://example.com"
	shortened := shortner.Shorten(database, lUri)
	if shortened == "" {
		t.Errorf("shortened uri should not be empty")
	}
	ok := database.Exists(shortened)
	if !ok {
		t.Errorf("shortened uri should exist")
	}
	savedUri, err := database.Get(shortened)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if savedUri != lUri {
		t.Errorf("unexpected uri: %s", lUri)
	}
}
