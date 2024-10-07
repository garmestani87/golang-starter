package main

import (
	"log"
	"net/http"
	"time"
)

type CustomHandler struct {
}

func main() {
	server := &http.Server{
		Addr:         ":8091",
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
		Handler:      &CustomHandler{},
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (ch *CustomHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("this is custom handler !!!"))
}
