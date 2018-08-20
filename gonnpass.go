package gonnpass

import "errors"

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
	// parse option to query string
	// access
	// get response
	// parse response
	// unmarshal json to bytes[]
	// return
	return nil, errors.New("")
}
