package db

type UriDatabase interface {
	Save(shortened, uri string)
	Get(shortened string) (uri string, err error)
	Exists(shortened string) bool
}
