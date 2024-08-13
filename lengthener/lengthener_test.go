package lengthener_test

import (
	"testing"
	"url_shortener/db"
	"url_shortener/lengthener"
)

func TestLengthener(t *testing.T) {
	database := db.NewDataBase()
	shortened := "short"
	uri := "uri"
	database.Save(shortened, uri)
	lUri, err := lengthener.Lengthen(database, shortened)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if lUri != uri {
		t.Errorf("unexpected uri: %s", uri)
	}
}
