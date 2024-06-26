package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	err := openDB()
	if err != nil {
		log.Panic(err)
	}
	defer closeDB()
	err = setupDB()
    r := chi.NewRouter()
    r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
        temp, _ := template.New("").ParseFiles("templates/index.html")
		temp.ExecuteTemplate(w, "Base", nil)
    })
    http.ListenAndServe("localhost:3000", r)
}
