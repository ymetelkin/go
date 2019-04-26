package links

import "github.com/ymetelkin/go/es"

//LambdaLinkRequest lambda link CRUD request
type LambdaLinkRequest struct {
	CollectionID string `json:"doc_id"`
	LinkID       string `json:"link_id"`
	UserID       string `json:"user_id"`
}

//LambdaResponse lamnda response
type LambdaResponse struct {
	IsSuccess bool   `json:"success"`
	Text      string `json:"response"`
	Links     []link `json:"links"`
}

type link struct {
	Seq      int    `json:"seq"`
	ID       string `json:"doc_id"`
	Type     string `json:"media_type"`
	Date     string `json:"arrivaldatetime"`
	Headline string `json:"headline"`
	Linked   audit  `json:"linked"`
}

type audit struct {
	ID        string `json:"by"`
	Timestamp string `json:"ts"`
}

//NewLambdaResponse constructs LambdaResponse
func NewLambdaResponse(success bool, text string, col Collection, docs []es.ApplDocument) LambdaResponse {
	res := LambdaResponse{IsSuccess: success, Text: text}

	var (
		size  int
		links []link
		d     bool
	)

	if col.Links != nil {
		size = len(col.Links)
	}
	if size == 0 {
		return res
	}

	if docs != nil {
		d = len(docs) > 0
	}

	links = make([]link, size)

	for i, l := range col.Links {
		lnk := link{
			Seq:    l.Seq,
			ID:     l.ID,
			Linked: audit{ID: l.Updated.ID, Timestamp: l.Updated.DateTime()},
		}

		if d {
			doc := docs[i]
			headline, _ := doc.Body.GetString("headline")
			lnk.Type = doc.Type
			lnk.Date = doc.Date
			lnk.Headline = headline
		}

		links[i] = lnk
	}

	res.Links = links
	return res
}
