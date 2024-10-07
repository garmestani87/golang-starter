package main

import (
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/go", http.RedirectHandler("https://go.dev/", 302))
	mux.HandleFunc("/dev", devHandler)

	server := &http.Server{
		Addr:         ":8090",
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
		Handler:      mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func devHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(403)
	rw.Write([]byte("I am a developer !"))
}
