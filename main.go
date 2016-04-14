package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("./app"))))

	mux.HandleFunc("/api/concerts", handleConcerts)

	// Handler wasn't matched, show not found.
	mux.HandleFunc("/", handleNotFound)
	log.Fatal(http.ListenAndServe(":80", mux))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("app/index.html")
	if err != nil {
		log.Println(err.Error())
	}

	reader := bufio.NewReader(f)

	io.Copy(w, reader)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
