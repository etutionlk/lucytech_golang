package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	urlInput := r.FormValue("urlInput")

	fmt.Println(urlInput)
	http.Redirect(w, r, "/result/", http.StatusFound)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/result/", http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/templates/index.html")
	})

	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/result/", resultHandler)

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
