package main

import (
	"flag"
	"github.com/shmurakami/gonnpass"
	"fmt"
	"log"
	"encoding/json"
	"strings"
)

/*
TODO use go-flags instead https://godoc.org/github.com/jessevdk/go-flags
TODO color
TODO loading dialog
TODO save group

./gonnpass groups [get|set|rm]
groups
  save id: name
  get id: name
 */

var (
	idFlag      = flag.Int("i", 0, "Event ID")
	keywordFlag = flag.String("k", "", "Search keyword")
	monthFlag   = flag.String("m", "", "Search by month e.g. 201801")
	dateFlag    = flag.String("d", "", "Search by date e.g. 20180101")
	nameFlag    = flag.String("n", "", "Search by attendee nickname")
	ownerFlag   = flag.String("o", "", "Search by owner nickname")
	groupFlag   = flag.Int("g", 0, "Search by group ID. Run './connpass groups get' to show groups")
	offsetFlag  = flag.Int("offset", 0, "Offset")
	limitFlag   = flag.Int("limit", 20, "Limit")
	orderFlag   = flag.String("order", "created", "created|date|updated. Default: created")
	formatFlag  = flag.String("format", "json", "json|parsed. Default: json")
	passedFlag  = flag.Bool("p", false, "Show passed event")
)

func main() {
	flag.Parse()

	option := gonnpass.Option{
		Id:        *idFlag,
		Keyword:   *keywordFlag,
		Month:     *monthFlag,
		Date:      *dateFlag,
		Name:      *nameFlag,
		Owner:     *ownerFlag,
		Group:     gonnpass.Group{},
		GroupFlag: *groupFlag,
		Offset:    *offsetFlag,
		Limit:     *limitFlag,
		Order:     gonnpass.Order["created"],
		OrderFlag: *orderFlag,
		Passed:    *passedFlag,
	}
	option = normalizeOption(option)

	response, err := gonnpass.Search(option)
	if err != nil {
		log.Fatal("failed to request to connpass api")
	}

	err = output(response, *formatFlag)
	if err != nil {
		log.Fatal(err)
	}
}

func normalizeOption(option gonnpass.Option) gonnpass.Option {
	// support keyword_or ?

	// date is prior than month
	if option.Date != "" {
		option.Month = ""
	}

	if option.GroupFlag != 0 {
		// no need to set Name
		option.Group = gonnpass.Group{Id: option.GroupFlag}
	}

	if order := gonnpass.Order[option.OrderFlag]; order != 0 {
		option.Order = order
	}

	return option
}

func output(r gonnpass.Response, f string) error {
	if f == "json" {
		b, err := json.Marshal(r)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}

	outputParsed(r)
	return nil
}

func outputParsed(r gonnpass.Response) {
	var sb strings.Builder

	for _, e := range r.Events {
		sb.WriteString(fmt.Sprintf("Event ID:	 %d\n", e.EventId))
		sb.WriteString(fmt.Sprintf("Title:		 %s\n", e.Title))
		sb.WriteString(fmt.Sprintf("Summary:	 %s\n", e.Catch))
		sb.WriteString(fmt.Sprintf("Address:	 %s | %s\n", e.Place, e.Address))
		sb.WriteString(fmt.Sprintf("Date:		 %s\n", e.StartedAt))
		sb.WriteString(fmt.Sprintf("Attendees:	 %d/%d (%d)\n", e.Accepted, e.Limit, e.Waiting))
		if e.Series.Id != 0 {
			s := e.Series
			sb.WriteString(fmt.Sprintf("Group:		 %sb (%d)\n", s.Title, s.Id))
		}
		if e.HashTag != "" {
			sb.WriteString(fmt.Sprintf("hash tag:	 #%s\n", e.HashTag))
		}
		sb.WriteString(fmt.Sprintf("\n"))
	}
	fmt.Println(sb.String())
}
