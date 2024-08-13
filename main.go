package main

import (
	"net/http"
	"url_shortener/application"
	"url_shortener/db"
)

func main() {
	database := db.NewDataBase()
	app := application.NewUrlShortenerApplication(database)

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		shortened := app.Shorten(url)
		w.Write([]byte(shortened))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortened := r.URL.Path[1:]
		uri, err := app.Lengthen(shortened)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, uri, http.StatusFound)
	})

	http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range database {
			w.Write([]byte(k + " : " + v + "\n"))
		}
	})

	http.ListenAndServe(":8080", nil)
}
