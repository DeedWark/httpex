package main

import (
	"bytes"
	"encoding/json"

	"github.com/yosssi/gohtml"
)

// FmtJSON: From bytes to JSON (string)
func FmtJSON(body []byte) (string, error) {
	var fmtJSON bytes.Buffer
	if err := json.Indent(&fmtJSON, body, "", "    "); err != nil {
		return "", err
	}
	return fmtJSON.String(), nil
}

// FmtHTML: From bytes to HTML (string)
func FmtHTML(body []byte) string {
	return gohtml.Format(string(body))
}
