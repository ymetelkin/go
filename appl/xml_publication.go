package appl

import (
	"github.com/ymetelkin/go/json"
)

type XmlPublication struct {
	Identification        XmlIdentification
	PublicationManagement XmlPublicationManagement
}

func (publication *XmlPublication) ToJson(jo *json.JsonObject) error {
	jo.AddString("representationversion", "1.0")
	jo.AddString("representationtype", "full")

	err := publication.Identification.ToJson(jo)
	if err != nil {
		return err
	}

	err = publication.PublicationManagement.ToJson(jo)
	if err != nil {
		return err
	}

	err = publication.getReferenceId(jo)
	if err != nil {
		return err
	}

	return nil
}

func (publication *XmlPublication) getReferenceId(jo *json.JsonObject) error {
	if publication.Identification.referenceid != "" {
		return nil
	}

	return nil
}
