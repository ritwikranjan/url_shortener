package db_test

import (
	"testing"
	"url_shortener/db"
)

func TestDataBase(t *testing.T) {
	database := db.NewDataBase()
	if database == nil {
		t.Errorf("database should not be nil")
	}
}

func TestSave(t *testing.T) {
	database := db.NewDataBase()
	shortened := "short"
	uri := "uri"
	database.Save(shortened, uri)
	if !database.Exists(shortened) {
		t.Errorf("shortened uri should exist")
	}
}

func TestGet(t *testing.T) {
	database := db.NewDataBase()
	shortened := "short"
	uri := "uri"
	database.Save(shortened, uri)
	lUri, err := database.Get(shortened)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if lUri != uri {
		t.Errorf("unexpected uri: %s", uri)
	}
}

func TestGetNotFound(t *testing.T) {
	database := db.NewDataBase()
	shortened := "short"
	_, err := database.Get(shortened)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExists(t *testing.T) {
	database := db.NewDataBase()
	shortened := "short"
	uri := "uri"
	database.Save(shortened, uri)
	if !database.Exists(shortened) {
		t.Errorf("shortened uri should exist")
	}
}
