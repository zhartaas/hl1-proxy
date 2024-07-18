package main

import (
	"github.com/swaggo/http-swagger" // http-swagger middleware)
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/proxy", app.proxy)
	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		httpSwagger.Handler(
			httpSwagger.URL("http://localhost:4000/swagger/doc.json"),
		).ServeHTTP(w, r)
	})
	return mux
}
