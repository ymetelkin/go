package appl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (doc *document) ParseIdentification(jo *json.Object) error {
	if doc.Identification.Nodes == nil {
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
			if nd.Text != "" {
				jo.AddString("editorialpriority", nd.Text)
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

func (doc *document) SetReferenceId(jo *json.Object) {
	var ref string

	if (doc.MediaType == MEDIATYPE_PHOTO || doc.MediaType == MEDIATYPE_GRAPHIC) && doc.FriendlyKey != "" {
		ref = doc.FriendlyKey
	} else if doc.MediaType == MEDIATYPE_AUDIO && doc.EditorialID != "" {
		ref = doc.EditorialID
	} else if doc.MediaType == MEDIATYPE_COMPLEXT_DATA && doc.Title != "" {
		ref = doc.Title
	} else if doc.MediaType == MEDIATYPE_TEXT {
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
	} else if doc.MediaType == MEDIATYPE_VIDEO {
		if doc.CompositionType == "StandardBroadcastVideo" {
			if doc.EditorialID != "" {
				ref = doc.EditorialID
			}
		} else if doc.Filings != nil {
			var exit bool
			for _, f := range doc.Filings {
				if f.ForeignKeys != nil {
					for _, fk := range f.ForeignKeys {
						if fk.Value != "" {
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

func getMediaType(s string) (MediaType, error) {
	var mt MediaType
	if strings.EqualFold(s, "text") {
		mt = MEDIATYPE_TEXT
	} else if strings.EqualFold(s, "photo") {
		mt = MEDIATYPE_PHOTO
	} else if strings.EqualFold(s, "video") {
		mt = MEDIATYPE_VIDEO
	} else if strings.EqualFold(s, "audio") {
		mt = MEDIATYPE_AUDIO
	} else if strings.EqualFold(s, "graphic") {
		mt = MEDIATYPE_GRAPHIC
	} else if strings.EqualFold(s, "complexdata") {
		mt = MEDIATYPE_COMPLEXT_DATA
	} else {
		e := fmt.Sprintf("Invalid media type [%s]", s)
		return MEDIATYPE_UNKNOWN, errors.New(e)
	}

	return mt, nil
}
