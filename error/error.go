package main

import (
	// "errors"
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	statusCode int
	message    string
}

func main() {
	response, err := validateUrl("https://jsonplaceholder.typicode.com/posts")
	if err != nil || response == nil {
		fmt.Println(err.message)
		return
	}
	valid, _ := validateStatusCode(response.StatusCode)
	if valid {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("response : %s", string(body))
	}
}

func validateStatusCode(code int) (valid bool, err error) {
	if code != 200 {
		return false, fmt.Errorf("error occured %d", code)
		// return false, errors.New("error occured !")
	}
	return true, nil

}

func validateUrl(url string) (response *http.Response, error *HttpError) {
	if url == "" {
		return nil, newHttpError(500, "url is invalid")
	}
	res, err := http.Get(url)
	if err != nil {
		return res, nil
	}
	return nil, newHttpError(500, "unexpected error !")
}

func newHttpError(statusCode int, message string) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		message:    message,
	}
}
