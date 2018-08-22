package gonnpass

import "testing"

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
