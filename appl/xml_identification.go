package appl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

/*
type MediaType string

const (
	MEDIATYPE_TEXT          MediaType = "text"
	MEDIATYPE_PHOTO         MediaType = "photo"
	MEDIATYPE_VIDEO         MediaType = "video"
	MEDIATYPE_AUDIO         MediaType = "audio"
	MEDIATYPE_GRAPHIC       MediaType = "graphic"
	MEDIATYPE_COMPLEXT_DATA MediaType = "complexdata"
	MEDIATYPE_UNKNOWN       MediaType = ""
)
*/

type XmlIdentification struct {
	ItemId               string
	RecordId             string
	CompositeId          string
	CompositionType      string
	MediaType            string
	Priority             int
	EditorialPriority    string
	DefaultLanguage      string
	RecordSequenceNumber int
	FriendlyKey          string

	mediaType   MediaType
	referenceid string
}

func (identification *XmlIdentification) GetMediaType() (MediaType, error) {
	if identification.mediaType == "" {
		left := identification.MediaType
		if strings.EqualFold(left, "text") {
			identification.mediaType = MEDIATYPE_TEXT
		} else if strings.EqualFold(left, "photo") {
			identification.mediaType = MEDIATYPE_PHOTO
		} else if strings.EqualFold(left, "video") {
			identification.mediaType = MEDIATYPE_VIDEO
		} else if strings.EqualFold(left, "audio") {
			identification.mediaType = MEDIATYPE_AUDIO
		} else if strings.EqualFold(left, "graphic") {
			identification.mediaType = MEDIATYPE_GRAPHIC
		} else if strings.EqualFold(left, "complexdata") {
			identification.mediaType = MEDIATYPE_COMPLEXT_DATA
		} else {
			e := fmt.Sprintf("Invalid media type [%s]", identification.MediaType)
			return MEDIATYPE_UNKNOWN, errors.New(e)
		}
	}

	return identification.mediaType, nil
}

func (identification *XmlIdentification) ToJson(jo *json.JsonObject) error {
	if identification.ItemId == "" {
		return errors.New("[Identification.ItemId] is missing")
	}
	if identification.RecordId == "" {
		return errors.New("[Identification.RecordId] is missing")
	}
	if identification.CompositeId == "" {
		return errors.New("[Identification.CompositeId] is missing")
	}
	if identification.CompositionType == "" {
		return errors.New("[Identification.CompositionType] is missing")
	}

	t, err := identification.GetMediaType()
	if err != nil {
		return err
	}

	jo.AddString("itemid", identification.ItemId)
	jo.AddString("recordid", identification.RecordId)
	jo.AddString("compositeid", identification.CompositeId)
	jo.AddString("compositiontype", identification.CompositionType)
	jo.AddString("type", string(t))

	if identification.Priority > 0 {
		jo.AddInt("priority", identification.Priority)
	}
	if identification.EditorialPriority != "" {
		jo.AddString("editorialpriority", identification.EditorialPriority)
	}
	if len(identification.DefaultLanguage) >= 2 {
		language := string([]rune(identification.DefaultLanguage)[0:2])
		jo.AddString("language", language)
	}
	if identification.RecordSequenceNumber > 0 {
		jo.AddInt("recordsequencenumber", identification.RecordSequenceNumber)
	}
	if identification.FriendlyKey != "" {
		jo.AddString("friendlykey", identification.FriendlyKey)
	}

	if t == MEDIATYPE_PHOTO || t == MEDIATYPE_GRAPHIC {
		if identification.FriendlyKey == "" {
			identification.referenceid = identification.ItemId
		} else {
			identification.referenceid = identification.FriendlyKey
		}
	}
	jo.AddString("referenceid", identification.referenceid)

	return nil
}
