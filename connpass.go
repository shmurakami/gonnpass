package gonnpass

// or use net/url ?
var Url = "https://connpass.com/api/v1/event/"

type Response struct {
	Events []ResponseEvent `json:"events"`
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
