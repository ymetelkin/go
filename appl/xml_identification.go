package appl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

type MediaType string

const (
	TEXT          MediaType = "text"
	PHOTO         MediaType = "photo"
	VIDEO         MediaType = "video"
	AUDIO         MediaType = "audio"
	GRAPHIC       MediaType = "graphic"
	COMPLEXT_DATA MediaType = "complexdata"
	UNKNOWN       MediaType = ""
)

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

	mediaType MediaType
}

func (identification *XmlIdentification) GetMediaType() (MediaType, error) {
	if identification.mediaType == "" {
		s := strings.ToLower(identification.MediaType)
		switch s {
		case "text":
			identification.mediaType = TEXT
		case "photo":
			identification.mediaType = PHOTO
		case "video":
			identification.mediaType = VIDEO
		case "audio":
			identification.mediaType = AUDIO
		case "graphic":
			identification.mediaType = GRAPHIC
		case "complexdata":
			identification.mediaType = COMPLEXT_DATA
		default:
			e := fmt.Sprintf("Invalid media type [%s]", s)
			return UNKNOWN, errors.New(e)
		}
	}

	return identification.mediaType, nil
}

func (identification *XmlIdentification) ToJson(jo *json.JsonObject) error {
	if identification.ItemId == "" {
		return errors.New("[ItemId] is missing")
	}
	if identification.RecordId == "" {
		return errors.New("[RecordId] is missing")
	}
	if identification.CompositeId == "" {
		return errors.New("[CompositeId] is missing")
	}
	if identification.CompositionType == "" {
		return errors.New("[CompositionType] is missing")
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
		jo.AddString("friendlykey ", identification.FriendlyKey)
	}

	referenceid := ""
	if t == PHOTO || t == GRAPHIC {
		if identification.FriendlyKey == "" {
			referenceid = identification.ItemId
		} else {
			referenceid = identification.FriendlyKey
		}
	}
	jo.AddString("referenceid", referenceid)

	return nil
}
