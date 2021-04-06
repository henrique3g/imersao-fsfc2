package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed public/*.jpg public/*.css
var staticContent embed.FS
//go:embed public/index.html
var indexHtml embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "public/index.html")
		t, _ := template.ParseFS(indexHtml, "public/index.html")
		t.Execute(w, nil)
	})
	http.Handle("/public/", http.FileServer(http.FS(staticContent)))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		print(err.Error())
	}
}
