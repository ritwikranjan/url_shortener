package db

import "fmt"

type LocalDataBase map[string]string

func NewDataBase() LocalDataBase {
	return make(LocalDataBase)
}

func (db LocalDataBase) Save(shortened, uri string) {
	db[shortened] = uri
}

func (db LocalDataBase) Get(shortened string) (uri string, err error) {
	uri, ok := db[shortened]
	if !ok {
		err = fmt.Errorf("uri not found for %s", shortened)
		return "", err
	}
	return uri, nil
}

func (db LocalDataBase) Exists(shortened string) (ok bool) {
	_, ok = db[shortened]
	return
}
