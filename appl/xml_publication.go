package appl

import (
	"github.com/ymetelkin/go/json"
)

type XmlPublication struct {
	Identification XmlIdentification
}

func (publication *XmlPublication) ToJson(jo *json.JsonObject) error {
	jo.AddString("representationversion", "1.0")
	jo.AddString("representationtype", "full")

	err := publication.Identification.ToJson(jo)
	if err != nil {
		return err
	}

	return nil
}
