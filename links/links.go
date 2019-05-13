package links

import (
	"time"
)

//Link struct
type Link struct {
	ID      string
	Seq     int
	Href    string
	Updated UpdateHistory
}

//UpdateHistory struct
type UpdateHistory struct {
	ID   string
	Unix int64
}

//DateTime returns formatted update time
func (upd *UpdateHistory) DateTime() string {
	return time.Unix(upd.Unix, 0).UTC().Format("2006-01-02T15:04:05.000")
}

//NewUpdateHistory constructs new update history
func NewUpdateHistory(id string) UpdateHistory {
	return UpdateHistory{ID: id, Unix: time.Now().Unix()}
}

//NewLink constructs new link
func NewLink(id string, seq int, by string) Link {
	return Link{ID: id, Seq: seq, Updated: NewUpdateHistory(by)}
}
