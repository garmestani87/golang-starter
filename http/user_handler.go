package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserHandler struct{}

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

var users = map[string]User{}

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
		getUserByName(rw, req, req.URL.Query().Get("name"))
	case req.Method == http.MethodGet && len(req.URL.Query().Get("name")) == 0:
		getUsers(rw, req)
	case req.Method == http.MethodPost:
		addUser(rw, req)
	default:
		{
			rw.Write([]byte("unhandled request !"))
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func addUser(rw http.ResponseWriter, req *http.Request) {

	key := req.Header.Get("api-key")
	if key != "jafar" {
		rw.Write([]byte("api key is invalid !"))
	}

	user := User{}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		rw.Write([]byte(err.Error()))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	users[user.Name] = user
	rw.Write([]byte("user created !"))
	rw.WriteHeader(http.StatusOK)
}

func getUsers(rw http.ResponseWriter, req *http.Request) {
	
	rw.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(rw).Encode(users)
	if err != nil {
		rw.Write([]byte(err.Error()))
		rw.WriteHeader(http.StatusInternalServerError)
	}

}

func getUserByName(rw http.ResponseWriter, req *http.Request, name string) {
	
	rw.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(rw).Encode(users[name])
	if err != nil {
		rw.Write([]byte(err.Error()))
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
