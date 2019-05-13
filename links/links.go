package links

import (
	"time"
)

<<<<<<< HEAD
//Link struct
type Link struct {
	ID      string
	Seq     int
	Href    string
	Updated UpdateHistory
=======
type doc struct {
	ID   string `json:"id"`
	Href string `json:"href"`
>>>>>>> 5f47947789048c5e033d95409fb25ea7dbbfa033
}

type link struct {
	doc
	Seq     int   `json:"seq"`
	Updated audit `json:"updated"`
}

type audit struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"ts"`
	DateTime  string `json:"dt"`
}

func newAudit(id string) audit {
	now := time.Now()
	ts := now.Unix()
	dt := now.UTC().Format("2006-01-02T15:04:05.000")

	return audit{
		ID:        id,
		Timestamp: ts,
		DateTime:  dt,
	}
}

func newLink(id string, seq int, href string, audit audit) link {
	return link{
		doc: doc{
			ID:   id,
			Href: href,
		},
		Seq:     seq,
		Updated: audit,
	}
}
