/* Author: DeedWark <github.com/DeedWark>
 * @2023
 * v1
 */
package httpex

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientOpt struct {
	Timeout time.Duration
}

var (
	copt ClientOpt
)

type Header *http.Header

// MAKE REQUEST -> http.Response
func Request(method string, url string, data string, headers http.Header) (*http.Response, error) {
	if method == "" {
		method = "GET"
	}

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
		return nil, err
	}

	req.Header = headers

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Response(request *http.Response) ([]byte, error) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
