package gonnpass

import (
	"strings"
	"encoding/json"
	"fmt"
)

// or use net/url ?
var Url = "https://connpass.com/api/v1/event/"

type Response struct {
	ResultsReturned  int             `json:"ResultsReturned"`
	Events           []ResponseEvent `json:"events"`
	ResultsStart     int             `json:"results_start"`
	ResultsAvailable int             `json:"results_available"`
}

type ResponseEvent struct {
	EventId          int                 `json:"event_id"`
	EventUrl         string              `json:"event_url"`
	EventType        string              `json:"event_type"`
	OwnerNickname    string              `json:"owner_nickname"`
	Series           ResponseEventSeries `json:"series"`
	UpdatedAt        string              `json:"updated_at"`
	Lat              string              `json:"lat"`
	StartedAt        string              `json:"started_at"`
	HashTag          string              `json:"hash_tag"`
	Title            string              `json:"title"`
	Lon              string              `json:"lon"`
	Waiting          int                 `json:"waiting"`
	Limit            int                 `json:"limit"`
	OwnerId          int                 `json:"owner_id"`
	OwnerDisplayName string              `json:"owner_display_name"`
	Description      string              `json:"description"`
	Address          string              `json:"address"`
	Catch            string              `json:"catch"`
	Accepted         int                 `json:"accepted"`
	EndedAt          string              `json:"ended_at"`
	Place            string              `json:"place"`
}

type ResponseEventSeries struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func (r *Response) Format(f string) (string, error) {
	if f == "json" {
		b, err := json.Marshal(r)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

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
	return sb.String(), nil
}
