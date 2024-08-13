package application_test

import (
	"testing"
	"url_shortener/application"
	"url_shortener/db"
)

func TestUrlShortenerApplication(t *testing.T) {
	db := db.NewDataBase()
	app := application.NewUrlShortenerApplication(db)

	lUris := []string{"http://www.google.com", "http://www.bing.com", "http://www.facebook.com"}
	for _, uri := range lUris {
		shortened := app.Shorten(uri)
		lUri, err := app.Lengthen(shortened)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if lUri != uri {
			t.Errorf("unexpected uri: %s", uri)
		}
	}
}

func TestUniqueShortened(t *testing.T) {
	db := db.NewDataBase()
	app := application.NewUrlShortenerApplication(db)

	uri := "http://www.google.com"
	shortened1 := app.Shorten(uri)
	shortened2 := app.Shorten(uri)
	if shortened1 == shortened2 {
		t.Errorf("shortened urls should be unique")
	}
}
