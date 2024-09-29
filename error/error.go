package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HttpMethod int

const (
	GET  HttpMethod = 1
	POST HttpMethod = 2
)

type Result[T any] struct {
	err      HttpError
	response *T
}

type HttpError struct {
	statusCode int
	method     HttpMethod
	error      error
	message    string
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("http error occured with status code = %d, method = %d, message = %s ", err.statusCode, err.method, err.message)
}

func (err *HttpError) Unwrap() error {
	return err.error
}

func newHttpError(statusCode int, message string, method HttpMethod) *HttpError {
	err := HttpError{
		statusCode: statusCode,
		message:    message,
		method:     method,
		error:      errors.New(""),
	}
	return &err
}

func newResult[T http.Response](response *T, err *HttpError) (res *Result[T]) {
	return &Result[T]{
		err:      *err,
		response: response,
	}
}

func main() {
	result := validateUrl("")
	if result.err.Unwrap() != nil {
		fmt.Println(result.err.Error())
		return
	}
	valid, _ := validateStatusCode(result.response.StatusCode)
	if valid {
		body, err := io.ReadAll(result.response.Body)
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
	}
	return true, nil

}

func validateUrl(url string) *Result[http.Response] {
	if url == "" {
		return newResult(nil, newHttpError(500, "url is invalid !", GET))
	}
	res, err := http.Get(url)
	if err == nil {
		return newResult(res, nil)
	}
	return newResult(nil, newHttpError(500, "unexpected error !", GET))
}
