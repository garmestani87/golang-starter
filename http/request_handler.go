package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserHandler struct{}

type User struct {
	name     string
	lastname string
	age      int
}

func main() {

	server := &http.Server{
		Addr:         ":8090",
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
		Handler:      &UserHandler{},
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (handler *UserHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == http.MethodGet && len(req.URL.Query().Get("name")) > 0:
		getUserByName(rw, req)
	case req.Method == http.MethodGet && len(req.URL.Query().Get("name")) == 0:
		getUsers(rw, req)
	case req.Method == http.MethodPost:
		addUser(rw, req)
	default:
		{
			rw.Write([]byte("unhandled request !"))
			rw.WriteHeader(500)
		}
	}
}

func addUser(rw http.ResponseWriter, req *http.Request) {

	key := req.Header.Get("api-key")
	if key != "jafar" {
		rw.Write([]byte("api key is invalid !"))
	}

	user := User{}

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(user)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write([]byte("user created !"))
	rw.WriteHeader(200)
}

func getUsers(rw http.ResponseWriter, req *http.Request) {
	users := &[]User{
		User{name: "taghi", lastname: "rezaei", age: 10},
		User{name: "mola", lastname: "taghavi", age: 20},
	}
	err := json.NewEncoder(rw).Encode(users)
	if err != nil {
		log.Fatal(err)
	}

}

func getUserByName(rw http.ResponseWriter, req *http.Request) {
	user := &User{name: "ali", lastname: "razavi", age: 10}
	rw.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(rw).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}
