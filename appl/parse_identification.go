package appl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (id *Identification) parse(aj *ApplJson) error {
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

	err := getMediaType(aj)
	if err != nil {
		return err
	}

	if len(id.DefaultLanguage) >= 2 {
		language := string([]rune(id.DefaultLanguage)[0:2])
		aj.Language = &json.JsonProperty{Field: "language", Value: &json.JsonStringValue{Value: language}}
	}

	getReferenceId(aj)

	return nil
}

func getMediaType(aj *ApplJson) error {
	s := aj.Xml.Identification.MediaType

	if strings.EqualFold(s, "text") {
		aj.MediaType = MEDIATYPE_TEXT
	} else if strings.EqualFold(s, "photo") {
		aj.MediaType = MEDIATYPE_PHOTO
	} else if strings.EqualFold(s, "video") {
		aj.MediaType = MEDIATYPE_VIDEO
	} else if strings.EqualFold(s, "audio") {
		aj.MediaType = MEDIATYPE_AUDIO
	} else if strings.EqualFold(s, "graphic") {
		aj.MediaType = MEDIATYPE_GRAPHIC
	} else if strings.EqualFold(s, "complexdata") {
		aj.MediaType = MEDIATYPE_COMPLEXT_DATA
	} else {
		e := fmt.Sprintf("Invalid media type [%s]", s)
		return errors.New(e)
	}

	return nil
}

func getReferenceId(aj *ApplJson) {
	ref := aj.Xml.Identification.ItemId

	if (aj.MediaType == MEDIATYPE_PHOTO || aj.MediaType == MEDIATYPE_GRAPHIC) && aj.Xml.Identification.FriendlyKey != "" {
		ref = aj.Xml.Identification.FriendlyKey
	} else if aj.MediaType == MEDIATYPE_AUDIO && aj.Xml.PublicationManagement.EditorialId != "" {
		ref = aj.Xml.PublicationManagement.EditorialId
	} else if aj.MediaType == MEDIATYPE_COMPLEXT_DATA && aj.Xml.NewsLines.Title != "" {
		ref = aj.Xml.NewsLines.Title
	} else if aj.MediaType == MEDIATYPE_TEXT {
		if aj.Xml.NewsLines.Title != "" {
			ref = aj.Xml.NewsLines.Title
		} else if aj.Filings != nil && len(aj.Filings) > 0 {
			for _, f := range aj.Filings {
				if f.Xml.SlugLine != "" {
					ref = f.Xml.SlugLine
					break
				}
			}
		}
	} else if aj.MediaType == MEDIATYPE_VIDEO {
		if strings.EqualFold(aj.Xml.Identification.CompositionType, "StandardBroadcastVideo") {
			if aj.Xml.PublicationManagement.EditorialId != "" {
				ref = aj.Xml.PublicationManagement.EditorialId
			}
		} else {
			if aj.Filings != nil && len(aj.Filings) > 0 {
				for _, f := range aj.Filings {
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

	aj.ReferenceId = &json.JsonProperty{Field: "referenceid", Value: &json.JsonStringValue{Value: ref}}
}
