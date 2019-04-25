package es

import (
	"io/ioutil"
	"strings"

	"github.com/ymetelkin/go/json"
)

//ApplDocument struct
type ApplDocument struct {
	ID   string
	Type string
	Date string
	Body json.Object
}

//ApplService struct to aggregate all APPL related ES functions
type ApplService struct {
	ec client
}

//NewApplService constructs ApplService
func NewApplService(clusterURL string) (ApplService, error) {
	ec, err := newClient(clusterURL)
	if err != nil {
		return ApplService{}, err
	}
	return ApplService{ec: ec}, nil
}

//GetDocuments get docs in order of IDs
func (appl *ApplService) GetDocuments(ids []string) ([]ApplDocument, error) {
	var (
		size int
		sb   strings.Builder
	)

	if ids != nil {
		size = len(ids)
	}

	if size == 0 {
		return []ApplDocument{}, nil
	}

	sb.WriteString(`{"query":{"bool":{"filter":{"terms":{"itemid":[`)
	for i, id := range ids {
		if i > 0 {
			sb.WriteByte(44) //comma
		}
		sb.WriteByte(34) //quote
		sb.WriteString(id)
		sb.WriteByte(34) //quote
	}
	sb.WriteString("]}}}}}")

	query := sb.String()

	sr := newSearchRequest("appl")
	sr.SetQuery(query)
	sr.SetSize(size)

	cr, err := appl.ec.Search(sr)
	if err != nil {
		return nil, err
	}

	defer cr.Close()

	bytes, err := ioutil.ReadAll(cr)
	if err != nil {
		return nil, err
	}

	s := string(bytes)

	jo, err := json.ParseJSONObject(s)
	if err != nil {
		return nil, err
	}

	jo, err = jo.GetObject("hits")
	if err != nil {
		return nil, err
	}

	ja, err := jo.GetArray("hits")
	if err != nil {
		return nil, err
	}

	hits, err := ja.GetObjects()
	if err != nil {
		return nil, err
	}

	docs := make([]ApplDocument, len(ids))

	for _, hit := range hits {
		src, err := hit.GetObject("_source")
		if err != nil {
			return nil, err
		}

		iid, err := src.GetString("itemid")
		if err != nil {
			return nil, err
		}

		tp, err := src.GetString("type")
		if err != nil {
			return nil, err
		}

		dt, err := src.GetString("arrivaldatetime")
		if err != nil {
			return nil, err
		}

		doc := ApplDocument{ID: iid, Type: tp, Date: dt, Body: src}

		for i, id := range ids {
			if id == iid {
				docs[i] = doc
			}
		}
	}

	return docs, nil
}
