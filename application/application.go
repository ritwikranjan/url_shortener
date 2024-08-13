package application

import (
	"url_shortener/db"
	"url_shortener/lengthener"
	"url_shortener/shortner"
)

type UrlShortenerApplication interface {
	Shorten(url string) (shortened string)
	Lengthen(shortened string) (uri string, err error)
}

type urlShortenerApplicationImpl struct {
	db db.UriDatabase
	s  shortner.Shortener
	l  lengthener.Lengthener
}

func (a *urlShortenerApplicationImpl) Shorten(url string) (shortened string) {
	return a.s(a.db, url)
}

func (a *urlShortenerApplicationImpl) Lengthen(shortened string) (uri string, err error) {
	return a.l(a.db, shortened)
}

func NewUrlShortenerApplication(database db.UriDatabase) UrlShortenerApplication {
	return &urlShortenerApplicationImpl{
		db: database,
		s:  shortner.Shorten,
		l:  lengthener.Lengthen,
	}
}
