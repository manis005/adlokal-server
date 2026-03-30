package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ad server running 🚀")
}

func impressionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Impression:", r.URL.Query())
	w.WriteHeader(http.StatusOK)
}

func clickHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Click:", r.URL.Query())

	redirectURL := r.URL.Query().Get("url")
	if redirectURL == "" {
		redirectURL = "https://google.com"
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/impression", impressionHandler)
	http.HandleFunc("/click", clickHandler)

	port := "8080"
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}