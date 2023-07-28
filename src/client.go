/* Author: DeedWark <github.com/DeedWark>
 * @2023
 * v1
 */

package httpex

import "net/http"

// Get: HTTP GET Method
func Get(url string, headers http.Header) ([]byte, error) {
	get, err := Request("GET", url, "", headers)
	if err != nil {
		return nil, err
	}
	return Response(get)
}

// Head: HTTP HEAD Method
func Head(url string, headers http.Header) ([]byte, error) {
	head, err := Request("HEAD", url, "", headers)
	if err != nil {
		return nil, err
	}
	return Response(head)
}

// Post: HTTP POST Method
func Post(url string, body string, headers http.Header) ([]byte, error) {
	post, err := Request("POST", url, body, headers)
	if err != nil {
		return nil, err
	}
	return Response(post)
}

func Delete(url string, body string, headers http.Header) ([]byte, error) {
	del, err := Request("DELETE", url, body, headers)
	if err != nil {
		return nil, err
	}
	return Response(del)
}
