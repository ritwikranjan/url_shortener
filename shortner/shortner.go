package shortner

import (
	"math/rand"
	"url_shortener/db"
)

type Shortener func(database db.UriDatabase, url string) (shortened string)

const (
	allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	urlLength    = 6
)

func Shorten(database db.UriDatabase, url string) (shortened string) {
	for {
		shortened = generateRandomString()
		if !database.Exists(shortened) {
			database.Save(shortened, url)
			return
		}
	}
}

func generateRandomString() (randStr string) {
	for i := 0; i < urlLength; i++ {
		randStr += string(allowedChars[rand.Int31()%int32(len(allowedChars))])
	}
	return
}
