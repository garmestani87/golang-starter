package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Result[T any] struct {
	items []T
}

func main() {
	
	wg := sync.WaitGroup{}
	results := Result[string]{}
	
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go callApi(fmt.Sprint("https://jsonplaceholder.typicode.com/posts/", i)  , &results, &wg)
	}
	
	wg.Wait()
	for _, v := range results.items {
		fmt.Println(v)
	}
}

func callApi(url string, res *Result[string], wg *sync.WaitGroup) {
	
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("error ocured when calling api !")
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Fatal("error ocured when parsing response !")
	}
	
	res.items = append(res.items, string(body))
}
