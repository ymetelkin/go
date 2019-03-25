package appl

import (
	"encoding/xml"

	"github.com/ymetelkin/go/json"
)

func XmlToJson(s string) (*json.JsonObject, error) {
	publication := XmlPublication{}
	err := xml.Unmarshal([]byte(s), &publication)
	if err != nil {
		return nil, err
	}

	jo := json.JsonObject{}
	err = publication.ToJson(&jo)
	if err != nil {
		return nil, err
	}

	return &jo, nil
}
