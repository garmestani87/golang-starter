package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	channel := make(chan Post, 10)
	wg := sync.WaitGroup{}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go callApi(i, channel, &wg)
	}

	wg.Wait()
	close(channel)

	for p := range channel {
		fmt.Printf("%+v \n\n\n\n", p)
	}

}

func callApi(id int, channel chan<- Post, wg *sync.WaitGroup) {

	defer wg.Done()

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Fatalln("unexpexted error occured !!!")
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln("unexpexted error occured !!!")
	}

	result, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatalln("unexpexted error occured !!!")
	}

	p := &Post{}
	err_unmarshal := json.Unmarshal(result, p)

	if err_unmarshal != nil {
		log.Fatalln("unexpexted error occured !!!")
	}

	channel <- *p

}
