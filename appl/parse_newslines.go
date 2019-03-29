package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (nl *NewsLines) parse(aj *ApplJson) error {
	getHeadline(aj)
	getCopyrightNotice(aj)
	getBylines(aj)
	getPersons(aj)

	if nl.OverLine != nil {
		overlines := UniqueStrings{}
		for _, ol := range nl.OverLine {
			overlines.Add(ol)
		}
		aj.OverLines = overlines.ToJsonProperty("overlines")
	}

	if nl.KeywordLine != nil {
		keywordlines := UniqueStrings{}
		for _, kw := range nl.KeywordLine {
			keywordlines.Add(kw)
		}
		aj.KeywordLines = keywordlines.ToJsonProperty("keywordlines")
	}

	return nil
}

func getHeadline(aj *ApplJson) {
	var headline string

	t := aj.MediaType

	if t == MEDIATYPE_TEXT || t == MEDIATYPE_AUDIO {
		if aj.Xml.NewsLines.HeadLine != "" {
			headline = aj.Xml.NewsLines.HeadLine
		} else if aj.Xml.NewsLines.Title != "" {
			headline = aj.Xml.NewsLines.Title
		}
	} else if t == MEDIATYPE_VIDEO && (aj.Xml.PublicationManagement.Function == "" || !strings.EqualFold(aj.Xml.PublicationManagement.Function, "APTNLibrary")) {
		if aj.Xml.NewsLines.HeadLine != "" {
			headline = aj.Xml.NewsLines.HeadLine
		} else if aj.Xml.NewsLines.Title != "" {
			headline = aj.Xml.NewsLines.Title
		}
	} else if t == MEDIATYPE_PHOTO || t == MEDIATYPE_GRAPHIC || (t == MEDIATYPE_VIDEO && strings.EqualFold(aj.Xml.PublicationManagement.Function, "APTNLibrary")) {
		if aj.Xml.NewsLines.Title != "" {
			headline = aj.Xml.NewsLines.Title
		} else if aj.Xml.NewsLines.HeadLine != "" {
			headline = aj.Xml.NewsLines.HeadLine
		}
	} else if t == MEDIATYPE_AUDIO {
		if aj.Xml.NewsLines.HeadLine != "" {
			headline = aj.Xml.NewsLines.HeadLine
		} else if aj.Xml.NewsLines.Title != "" {
			headline = aj.Xml.NewsLines.Title
		}
	}

	aj.Headline = &json.JsonProperty{Field: "headline", Value: &json.JsonStringValue{Value: headline}}
}

func getCopyrightNotice(aj *ApplJson) {
	nl := aj.Xml.NewsLines
	var copyrightnotice string
	if nl.CopyrightLine != "" {
		copyrightnotice = nl.CopyrightLine
	} else if aj.FirstCreatedYear > 0 && aj.Xml.RightsMetadata.Copyright.Holder != "" {
		copyrightnotice = fmt.Sprintf("Copyright %d %s. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.", aj.FirstCreatedYear, aj.Xml.RightsMetadata.Copyright.Holder)
	}
	aj.CopyrightNotice = &json.JsonProperty{Field: "copyrightnotice", Value: &json.JsonStringValue{Value: copyrightnotice}}
}

func getBylines(aj *ApplJson) {
	nl := aj.Xml.NewsLines
	if nl.ByLine == nil || len(nl.ByLine) == 0 {
		return
	}

	bylines := json.JsonArray{}
	edits := json.JsonArray{}

	if nl.ByLineOriginal != nil || len(nl.ByLineOriginal) > 0 {
		for _, blo := range nl.ByLineOriginal {
			byline := json.JsonObject{}
			byline.AddString("by", blo.Value)
			if blo.Title != "" {
				byline.AddString("title", blo.Title)
			} else {
				for _, bl := range nl.ByLine {
					if bl.Title != "" {
						byline.AddString("title", bl.Title)
						break
					}
				}
			}
			bylines.AddObject(&byline)
		}
	} else {
		for _, bl := range nl.ByLine {
			if bl.Value != "" {
				if strings.EqualFold(bl.Title, "EditedBy") {
					producer := json.JsonObject{}
					if bl.Id != "" {
						producer.AddString("code", bl.Id)
					}
					producer.AddString("name", bl.Value)
					aj.Producer = &json.JsonProperty{Field: "producer", Value: &json.JsonObjectValue{Value: producer}}
				} else if strings.EqualFold(bl.Parametric, "PHOTOGRAPHER") && aj.Photographer == nil {
					photographer := json.JsonObject{}
					if bl.Id != "" {
						photographer.AddString("code", bl.Id)
					}
					photographer.AddString("name", bl.Value)
					if bl.Title != "" {
						photographer.AddString("title", bl.Title)
					}
					aj.Photographer = &json.JsonProperty{Field: "photographer", Value: &json.JsonObjectValue{Value: photographer}}

				} else if strings.EqualFold(bl.Parametric, "CAPTIONWRITER") && aj.CaptionWriter == nil {
					captionwriter := json.JsonObject{}
					if bl.Id != "" {
						captionwriter.AddString("code", bl.Id)
					}
					captionwriter.AddString("name", bl.Value)
					if bl.Title != "" {
						captionwriter.AddString("title", bl.Title)
					}
					aj.CaptionWriter = &json.JsonProperty{Field: "captionwriter", Value: &json.JsonObjectValue{Value: captionwriter}}
				} else if strings.EqualFold(bl.Parametric, "EDITEDBY") {
					edit := json.JsonObject{}
					edit.AddString("name", bl.Value)
					edits.AddObject(&edit)
				} else {
					byline := json.JsonObject{}
					if bl.Id != "" {
						byline.AddString("code", bl.Id)
					}
					byline.AddString("by", bl.Value)
					if bl.Title != "" {
						byline.AddString("title", bl.Title)
					}
					if bl.Parametric != "" {
						byline.AddString("parametric", bl.Parametric)
					}
					bylines.AddObject(&byline)
				}
			}
		}
	}

	if bylines.Length() > 0 {
		aj.Bylines = &json.JsonProperty{Field: "bylines", Value: &json.JsonArrayValue{Value: bylines}}
	}

	if edits.Length() > 0 {
		aj.Edits = &json.JsonProperty{Field: "edits", Value: &json.JsonArrayValue{Value: edits}}
	}
}

func getPersons(aj *ApplJson) {
	nl := aj.Xml.NewsLines

	if nl.NameLine != nil && len(nl.NameLine) > 0 {
		persons := json.JsonArray{}
		for _, name := range nl.NameLine {
			person := json.JsonObject{}
			person.AddString("name", name.Value)
			if strings.EqualFold(name.Parametric, "PERSON_FEATURED") {
				rel := json.JsonArray{}
				rel.AddString("personfeatured")
				person.AddArray("rel", &rel)
			}
			person.AddString("creator", "Editorial")
			persons.AddObject(&person)
		}
		aj.Persons = &json.JsonProperty{Field: "person", Value: &json.JsonArrayValue{Value: persons}}
	}
}
