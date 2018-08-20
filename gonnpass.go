package gonnpass

import (
	"errors"
	"net/http"
	"fmt"
)

var Order map[string]uint = map[string]uint{
	"update":  1,
	"date":    2,
	"created": 3,
}

type Group struct {
	Id   uint
	Name string
}

func Search(option Option) ([]Response, error) {
	query, err := parseOption(option)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?%s", Url, query)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// parse option to query string
	// access
	// get response
	// parse response
	// unmarshal json to bytes[]
	// return
	return nil, errors.New("")
}

func parseOption(option Option) (string, error) {
	return "", nil
}
