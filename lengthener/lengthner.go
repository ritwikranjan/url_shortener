package lengthener

import (
	"url_shortener/db"
)

type Lengthener func(database db.UriDatabase, shortened string) (uri string, err error)

func Lengthen(database db.UriDatabase, shortened string) (uri string, err error) {
	uri, err = database.Get(shortened)
	return
}
