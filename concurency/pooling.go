package main

import (
	"fmt"
	"sync"
)

type Connection struct {
	url  string
	db   string
	port int
}

func main() {
	connectionPoll := sync.Pool{
		New: func() any {
			return &Connection{
				url:  "192.168.41.253",
				db:   "postgres",
				port: 5432,
			}
		},
	}
	
	conn := connectionPoll.Get().(*Connection)
	fmt.Printf("%+v \n", *conn)
	
	//return to poll
	connectionPoll.Put(conn)
}
