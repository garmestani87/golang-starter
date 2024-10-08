package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cancel", handleRequest)

	server := &http.Server{
		Addr:         ":8093",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handleRequest(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		fmt.Println("request is cancelled !")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("response ....")
		return
	}
}
