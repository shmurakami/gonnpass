package gonnpass

import (
	"errors"
	"net/http"
	"fmt"
	"bytes"
	"io"
	"encoding/json"
	"net/url"
	"strconv"
)

var Order = map[string]int{
	"update":  1,
	"date":    2,
	"created": 3,
}

type Option struct {
	Id        int
	Keyword   string
	Month     string
	Date      string
	Name      string
	Owner     string
	Group     Group
	GroupFlag int
	Offset    int
	Limit     int
	OrderFlag string
	Order     int
}

type Group struct {
	Id   int
	Name string
}

func Search(option Option) (Response, error) {
	query, err := parseOption(option)
	if err != nil {
		return Response{}, err
	}

	u := fmt.Sprintf("%s?%s", Url, query)
	res, err := http.Get(u)
	if err != nil {
		return Response{}, err
	}
	if res.StatusCode != 200 {
		return Response{}, errors.New("http status code is not success")
	}

	var resp bytes.Buffer
	bout := res.Body
	res.Body = nil
	if bout != nil {
		_, err := io.Copy(&resp, bout)
		if err != nil {
			return Response{}, errors.New("failed to read response bout")
		}
		bout.Close()
	}
	//body := resp.String()
	var response Response
	err = json.Unmarshal(resp.Bytes(), &response)
	if err != nil {
		fmt.Println(err.Error())
		return Response{}, err
	}
	return response, nil
}

func parseOption(o Option) (string, error) {
	q := make(url.Values)

	// TODO have mapping of option field and url key name
	if o.Id != 0 {
		q.Add("event_id", strconv.Itoa(o.Id))
	}
	if o.Keyword != "" {
		q.Add("keyword", o.Keyword)
	}
	if o.Month != "" {
		q.Add("ym", o.Month)
	}
	if o.Date != "" {
		q.Add("ymd", o.Date)
	}
	if o.Name != "" {
		q.Add("nickname", o.Name)
	}
	if o.Owner != "" {
		q.Add("owner_nickname", o.Owner)
	}
	if o.Group.Id != 0 {
		q.Add("series_id", strconv.Itoa(o.Group.Id))
	}
	if o.Offset != 0 {
		q.Add("start", strconv.Itoa(o.Offset))
	}
	if o.Limit != 0 {
		q.Add("count", strconv.Itoa(o.Limit))
	}
	if o.Order != 0 {
		q.Add("order", strconv.Itoa(o.Order))
	}

	return q.Encode(), nil
}
