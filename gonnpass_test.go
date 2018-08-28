package gonnpass

import (
	"testing"
	"time"
	"fmt"
)

func TestParseOption(t *testing.T) {
	o := Option{
		Id:      1,
		Keyword: "keyword",
		Date:    "20180101",
		Name:    "name",
		Owner:   "owner",
		Group:   Group{Id: 1, Name: "group"},
		Order:   Order["created"],
	}

	q, _ := parseOption(o)
	expect := "event_id=1&keyword=keyword&nickname=name&order=3&owner_nickname=owner&series_id=1&ymd=20180101"
	if q != expect {
		t.Fatalf("not match\n%s\n%s", q, expect)
	}

}

func TestFilterPassed(t *testing.T) {
	r := Response{
		Events: []ResponseEvent{
			ResponseEvent{
				EventId:   1,
				StartedAt: "2018-09-01T10:00:00+09:00",
				EndedAt:   "2018-09-01T12:00:00+09:00",
			},
			ResponseEvent{
				EventId:   2,
				StartedAt: "2018-09-02T10:00:00+09:00",
				EndedAt:   "2018-09-02T12:00:00+09:00",
			},
			ResponseEvent{
				EventId:   3,
				StartedAt: "2018-09-02T13:00:00+09:00",
				EndedAt:   "2018-09-02T15:00:00+09:00",
			},
		},
	}

	now = time.Date(2018, time.September, 2, 12, 30, 0, 0, time.FixedZone("JST", 60 * 60 * 9)).Unix()

	e := filterPassed(r)

	if len(e.Events) != 1 {
		fmt.Println(len(e.Events))
		t.Fatal("failed to filter")
	}
}
