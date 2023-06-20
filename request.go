package httpex

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ClientOpt struct {
	Timeout time.Duration
}

var (
	copt ClientOpt
)

// Request: Make HTTP request
func Request(method string, url string, data string) ([]byte, error) {

	// dataBody is useful in case of data like JSON
	var dataBody io.Reader
	switch data {
	case "":
		dataBody = nil
	default:
		dataBody = bytes.NewBuffer([]byte(data))
	}

	client := &http.Client{Timeout: copt.Timeout}

	req, err := http.NewRequest(method, url, dataBody)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, nil
}

// Get: HTTP Get method
func Get(url string) ([]byte, error) {
	return Request(http.MethodGet, url, "")
}

// Head: HTTP Head method
func Head(url string) ([]byte, error) {
	return Request(http.MethodHead, url, "")
}

// Post: HTTP Post method
func Post(url string, data string) ([]byte, error) {
	return Request(http.MethodPost, url, data)
}

// Put: HTTP Put method
func Put(url string, data string) ([]byte, error) {
	return Request(http.MethodPut, url, data)
}

// Patch: HTTP Patch method
func Patch(url string, data string) ([]byte, error) {
	return Request(http.MethodPatch, url, data)
}

// Delete: HTTP Delete method
func Delete(url string, data string) ([]byte, error) {
	return Request(http.MethodDelete, url, data)
}

// Connect: HTTP Connect method
func Connect(url string, data string) ([]byte, error) {
	return Request(http.MethodConnect, url, data)
}

// Options: HTTP Options method
func Options(url string, data string) ([]byte, error) {
	return Request(http.MethodOptions, url, data)
}

// Trace: HTTP Trace method
func Trace(url string, data string) ([]byte, error) {
	return Request(http.MethodTrace, url, data)
}
