package appl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (doc *document) ParseIdentification(jo *json.Object) error {
	if doc.Identification == nil {
		return errors.New("Identification is missing")
	}

	for _, nd := range doc.Identification.Nodes {
		switch nd.Name {
		case "ItemId":
			doc.ItemID = nd.Text
			jo.SetString("itemid", nd.Text)
		case "RecordId":
			jo.SetString("recordid", nd.Text)
		case "CompositeId":
			jo.SetString("compositeid", nd.Text)
		case "CompositionType":
			doc.CompositionType = nd.Text
			jo.SetString("compositiontype", nd.Text)
		case "MediaType":
			mt, err := getMediaType(nd.Text)
			if err == nil {
				doc.MediaType = mt
				jo.SetString("type", string(mt))
			} else {
				return err
			}
		case "Priority":
			if nd.Text != "" {
				i, err := strconv.Atoi(nd.Text)
				if err == nil {
					jo.AddInt("priority", i)
				}
			}
		case "EditorialPriority":
			if nd.Text != "" {
				jo.AddString("editorialpriority", nd.Text)
			}
		case "DefaultLanguage":
			if len(nd.Text) >= 2 {
				language := string([]rune(nd.Text)[0:2])
				jo.AddString("language", language)
			}
			if nd.Text != "" {
				jo.AddString("editorialpriority", nd.Text)
			}
		case "RecordSequenceNumber":
			if nd.Text != "" {
				i, err := strconv.Atoi(nd.Text)
				if err == nil {
					jo.AddInt("recordsequencenumber", i)
				}
			}
		case "FriendlyKey":
			if nd.Text != "" {
				doc.FriendlyKey = nd.Text
				jo.AddString("friendlykey", nd.Text)
			}
		}
	}

	jo.AddString("referenceid", doc.ItemID)
}

func (doc *document) SetReferenceId(jo *json.Object) {
	var ref string

	if (doc.MediaType == MEDIATYPE_PHOTO || doc.MediaType == MEDIATYPE_GRAPHIC) && doc.FriendlyKey != "" {
		ref = doc.FriendlyKey
	} else if doc.MediaType == MEDIATYPE_AUDIO && doc.EditorialId != "" {
		ref = doc.EditorialId
	} else if doc.MediaType == MEDIATYPE_COMPLEXT_DATA && doc.Title != "" {
		ref = doc.Title
	} else if doc.MediaType == MEDIATYPE_TEXT {
		if doc.Title != "" {
			ref = doc.Title
		} else if doc.SlugLine != "" {
			ref = doc.SlugLine
		}
	} else if doc.MediaType == MEDIATYPE_VIDEO {
		if doc.CompositionType == "StandardBroadcastVideo" {
			if doc.EditorialId != "" {
				ref = doc.EditorialId
			}
		} else if doc.ForeignKey != "" {
			ref = doc.ForeignKey
		}
	}

	if ref != "" {
		jo.SetString("referenceid", ref)
	}
}

func getMediaType(s string) (MediaType, error) {
	if strings.EqualFold(s, "text") {
		doc.MediaType = MEDIATYPE_TEXT
	} else if strings.EqualFold(s, "photo") {
		doc.MediaType = MEDIATYPE_PHOTO
	} else if strings.EqualFold(s, "video") {
		doc.MediaType = MEDIATYPE_VIDEO
	} else if strings.EqualFold(s, "audio") {
		doc.MediaType = MEDIATYPE_AUDIO
	} else if strings.EqualFold(s, "graphic") {
		doc.MediaType = MEDIATYPE_GRAPHIC
	} else if strings.EqualFold(s, "complexdata") {
		doc.MediaType = MEDIATYPE_COMPLEXT_DATA
	} else {
		e := fmt.Sprintf("Invalid media type [%s]", s)
		return errors.New(e)
	}

	return nil
}
