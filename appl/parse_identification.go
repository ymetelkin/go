package appl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (id *Identification) parse(doc *document) error {
	if id.ItemId == "" {
		return errors.New("[Identification.ItemId] is missing")
	}
	if id.RecordId == "" {
		return errors.New("[Identification.RecordId] is missing")
	}
	if id.CompositeId == "" {
		return errors.New("[Identification.CompositeId] is missing")
	}
	if id.CompositionType == "" {
		return errors.New("[Identification.CompositionType] is missing")
	}

	err := getMediaType(doc)
	if err != nil {
		return err
	}

	if len(id.DefaultLanguage) >= 2 {
		language := string([]rune(id.DefaultLanguage)[0:2])
		doc.Language = json.NewStringProperty("language", language)
	}

	getReferenceId(doc)

	return nil
}

func getMediaType(doc *document) error {
	s := doc.Xml.Identification.MediaType

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

func getReferenceId(doc *document) {
	ref := doc.Xml.Identification.ItemId

	if (doc.MediaType == MEDIATYPE_PHOTO || doc.MediaType == MEDIATYPE_GRAPHIC) && doc.Xml.Identification.FriendlyKey != "" {
		ref = doc.Xml.Identification.FriendlyKey
	} else if doc.MediaType == MEDIATYPE_AUDIO && doc.Xml.PublicationManagement.EditorialId != "" {
		ref = doc.Xml.PublicationManagement.EditorialId
	} else if doc.MediaType == MEDIATYPE_COMPLEXT_DATA && doc.Xml.NewsLines.Title != "" {
		ref = doc.Xml.NewsLines.Title
	} else if doc.MediaType == MEDIATYPE_TEXT {
		if doc.Xml.NewsLines.Title != "" {
			ref = doc.Xml.NewsLines.Title
		} else if doc.Filings.Filings != nil {
			for _, f := range doc.Filings.Filings {
				if f.SlugLine != "" {
					ref = f.SlugLine
					break
				}
			}
		}
	} else if doc.MediaType == MEDIATYPE_VIDEO {
		if strings.EqualFold(doc.Xml.Identification.CompositionType, "StandardBroadcastVideo") {
			if doc.Xml.PublicationManagement.EditorialId != "" {
				ref = doc.Xml.PublicationManagement.EditorialId
			}
		} else {
			if doc.Filings.Filings != nil {
				for _, f := range doc.Filings.Filings {
					if f.ForeignKeys != nil {
						for _, v := range f.ForeignKeys {
							ref = v
							break
						}
						break
					}
				}
			}
		}
	}

	doc.ReferenceId = json.NewStringProperty("referenceid", ref)
}
