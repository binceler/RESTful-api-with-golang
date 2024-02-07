package main

import "net/http"

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is about!"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to index page!"))
	w.WriteHeader(http.StatusOK)
}

func main() {
	/*http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Merhaba Mars!"))
	})*/
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":9999", nil)
}
