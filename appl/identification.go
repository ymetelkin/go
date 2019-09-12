package appl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (doc *document) ParseIdentification(jo *json.Object) error {
	if doc.Identification == nil || doc.Identification.Nodes == nil || len(doc.Identification.Nodes) == 0 {
		return errors.New("Identification is missing")
	}

	for _, nd := range doc.Identification.Nodes {
		switch nd.Name {
		case "ItemId":
			doc.ItemID = nd.Text
			jo.AddString("itemid", nd.Text)
		case "RecordId":
			jo.AddString("recordid", nd.Text)
		case "CompositeId":
			jo.AddString("compositeid", nd.Text)
		case "CompositionType":
			doc.CompositionType = nd.Text
			jo.AddString("compositiontype", nd.Text)
		case "MediaType":
			mt, err := getMediaType(nd.Text)
			if err == nil {
				doc.MediaType = mt
				jo.AddString("type", string(mt))
			} else {
				return err
			}
		case "Priority":
			i, err := strconv.Atoi(nd.Text)
			if err == nil {
				jo.AddInt("priority", i)
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
		case "RecordSequenceNumber":
			i, err := strconv.Atoi(nd.Text)
			if err == nil {
				jo.AddInt("recordsequencenumber", i)
			}
		case "FriendlyKey":
			if nd.Text != "" {
				doc.FriendlyKey = nd.Text
				jo.AddString("friendlykey", nd.Text)
			}
		}
	}

	jo.AddString("referenceid", doc.ItemID)

	return nil
}

func (doc *document) SetReferenceID(jo *json.Object) {
	var ref string

	if (doc.MediaType == mediaTypePhoto || doc.MediaType == mediaTypeGraphic) && doc.FriendlyKey != "" {
		ref = doc.FriendlyKey
	} else if doc.MediaType == mediaTypeAudio && doc.EditorialID != "" {
		ref = doc.EditorialID
	} else if doc.MediaType == mediaTypeComplexData && doc.Title != "" {
		ref = doc.Title
	} else if doc.MediaType == mediaTypeText {
		if doc.Title != "" {
			ref = doc.Title
		} else if doc.Filings != nil {
			for _, f := range doc.Filings {
				if f.Slugline != "" {
					ref = f.Slugline
					break
				}
			}
		}
	} else if doc.MediaType == mediaTypeVideo {
		if doc.CompositionType == "StandardBroadcastVideo" {
			if doc.EditorialID != "" {
				ref = doc.EditorialID
			}
		} else if doc.Filings != nil {
			var exit bool
			for _, f := range doc.Filings {
				if f.ForeignKeys != nil {
					for _, fk := range f.ForeignKeys {
						if fk.Field == "storyid" && fk.Value != "" {
							ref = fk.Value
							exit = true
							break
						}
					}
				}
				if exit {
					break
				}
			}
		}
	}

	if ref != "" {
		jo.SetString("referenceid", ref)
	}
}

func getMediaType(s string) (mediaType, error) {
	var mt mediaType
	if strings.EqualFold(s, "text") {
		mt = mediaTypeText
	} else if strings.EqualFold(s, "photo") {
		mt = mediaTypePhoto
	} else if strings.EqualFold(s, "video") {
		mt = mediaTypeVideo
	} else if strings.EqualFold(s, "audio") {
		mt = mediaTypeAudio
	} else if strings.EqualFold(s, "graphic") {
		mt = mediaTypeGraphic
	} else if strings.EqualFold(s, "complexdata") {
		mt = mediaTypeComplexData
	} else {
		e := fmt.Sprintf("Invalid media type [%s]", s)
		return mediaTypeUnknown, errors.New(e)
	}

	return mt, nil
}
