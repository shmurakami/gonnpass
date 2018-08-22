package main

import (
	"flag"
	"github.com/shmurakami/gonnpass"
	"fmt"
	"log"
)

/*
gonnpass
just curl wrapper
https://connpass.com/about/api/

./gonnpass (get) (options)
get
  parse arguments
  -i --id event_id
  -k --keyword keyword
  -m --month ym 201808
  -d --date ymd 20180831
  -n --name nickname
  -o --owner owner_nickname
  -g --group series_id
  --offset start 1
  --order updated/date/created order 3
  --limit count 10 20?

TODO use go-flags instead https://godoc.org/github.com/jessevdk/go-flags


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
	orderFlag   = flag.String("order", "created", "created/date/updated. Default: created")
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
	}
	option = normalizeOption(option)

	response, err := gonnpass.Search(option)
	if err != nil {
		log.Fatal("failed to request to connpass api")
	}

	fmt.Println(response)
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
