package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	func() {
		server := &http.Server{
			Addr:         ":8089",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 20 * time.Second,
		}
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
