package es

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ymetelkin/go/json"
)

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
func (appl *ApplService) GetDocuments(ids []string, fields []string) (json.Array, error) {
	var (
		size int
		sb   strings.Builder
		docs json.Array
	)

	if ids != nil {
		size = len(ids)
	}

	if size == 0 {
		return docs, nil
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
	fmt.Println(query)

	sr := newSearchRequest("appl")
	sr.SetQuery(query)
	sr.SetSize(size)

	if fields != nil && len(fields) > 0 {
		sr.SetSource(fields)
	}

	cr, err := appl.ec.Search(sr)
	if err != nil {
		return docs, err
	}

	defer cr.Close()

	bytes, err := ioutil.ReadAll(cr)
	if err != nil {
		return docs, err
	}

	s := string(bytes)

	jo, err := json.ParseJSONObject(s)
	if err != nil {
		return docs, err
	}

	jo, err = jo.GetObject("hits")
	if err != nil {
		return docs, err
	}

	ja, err := jo.GetArray("hits")
	if err != nil {
		return docs, err
	}

	hits, err := ja.GetObjects()
	if err != nil {
		return docs, err
	}

	tmp := make([]json.Object, size)

	for _, hit := range hits {
		test, err := hit.GetString("_id")
		if err != nil {
			return docs, err
		}
		src, err := hit.GetObject("_source")
		if err != nil {
			return docs, err
		}

		for i, id := range ids {
			if id == test {
				tmp[i] = src
				break
			}
		}
	}

	for _, doc := range tmp {
		docs.AddObject(doc)
	}

	return docs, nil
}
