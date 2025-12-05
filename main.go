package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
http.HandleFunc("/url", UrlHandler)

}
func UrlHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>URL visited:</h1><p>%s</p>", r.URL.Path)
}
