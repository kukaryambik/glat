package unjson

import (
	"bytes"
	"encoding/json"
)

// BaseJSON is struct for Gitlab API response
type BaseJSON struct {
	ID   int
	Name string
}

// Pretty - Prettyfy Json
func Pretty(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

// Get from json
func Get(input string) BaseJSON {
	var out BaseJSON
	json.Unmarshal([]byte(input), &out)
	return out
}
