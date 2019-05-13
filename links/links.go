package links

import (
	"time"
)

//Link struct
type Link struct {
	ID      string `json:"id"`
	Href    string `json:"href"`
	Seq     int    `json:"seq"`
	Updated UpdateHistory
}

//UpdateHistory struct
type UpdateHistory struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"ts"`
	DateTime  string `json:"dt"`
}

//NewUpdateHistory constructs new update history
func NewUpdateHistory(id string) UpdateHistory {
	now := time.Now()
	ts := now.Unix()
	dt := now.UTC().Format("2006-01-02T15:04:05.000")

	return UpdateHistory{
		ID:        id,
		Timestamp: ts,
		DateTime:  dt,
	}
}

//NewLink constructs new link
func NewLink(id string, seq int, href string, by string) Link {
	return Link{
		ID:      id,
		Seq:     seq,
		Href:    href,
		Updated: NewUpdateHistory(by),
	}
}
