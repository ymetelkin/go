package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (nl *NewsLines) parse(doc *document) error {
	getHeadline(doc)
	getCopyrightNotice(doc)
	getBylines(doc)
	getPerson(doc)

	if nl.OverLine != nil {
		overlines := uniqueArray{}
		for _, ol := range nl.OverLine {
			overlines.AddString(ol)
		}
		doc.OverLines = overlines.ToJsonProperty("overlines")
	}

	if nl.KeywordLine != nil {
		keywordlines := uniqueArray{}
		for _, kw := range nl.KeywordLine {
			keywordlines.AddString(kw)
		}
		doc.KeywordLines = keywordlines.ToJsonProperty("keywordlines")
	}

	return nil
}

func getHeadline(doc *document) {
	var headline string

	t := doc.MediaType

	if t == MEDIATYPE_TEXT || t == MEDIATYPE_AUDIO {
		if doc.Xml.NewsLines.HeadLine != "" {
			headline = doc.Xml.NewsLines.HeadLine
		} else if doc.Xml.NewsLines.Title != "" {
			headline = doc.Xml.NewsLines.Title
		}
	} else if t == MEDIATYPE_VIDEO && (doc.Xml.PublicationManagement.Function == "" || !strings.EqualFold(doc.Xml.PublicationManagement.Function, "APTNLibrary")) {
		if doc.Xml.NewsLines.HeadLine != "" {
			headline = doc.Xml.NewsLines.HeadLine
		} else if doc.Xml.NewsLines.Title != "" {
			headline = doc.Xml.NewsLines.Title
		}
	} else if t == MEDIATYPE_PHOTO || t == MEDIATYPE_GRAPHIC || (t == MEDIATYPE_VIDEO && strings.EqualFold(doc.Xml.PublicationManagement.Function, "APTNLibrary")) {
		if doc.Xml.NewsLines.Title != "" {
			headline = doc.Xml.NewsLines.Title
		} else if doc.Xml.NewsLines.HeadLine != "" {
			headline = doc.Xml.NewsLines.HeadLine
		}
	} else if t == MEDIATYPE_AUDIO {
		if doc.Xml.NewsLines.HeadLine != "" {
			headline = doc.Xml.NewsLines.HeadLine
		} else if doc.Xml.NewsLines.Title != "" {
			headline = doc.Xml.NewsLines.Title
		}
	}

	doc.Headline = json.NewStringProperty("headline", headline)
}

func getCopyrightNotice(doc *document) {
	nl := doc.Xml.NewsLines
	var copyrightnotice string
	if nl.CopyrightLine != "" {
		copyrightnotice = nl.CopyrightLine
	} else if doc.FirstCreatedYear > 0 && doc.Xml.RightsMetadata.Copyright.Holder != "" {
		copyrightnotice = fmt.Sprintf("Copyright %d %s. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.", doc.FirstCreatedYear, doc.Xml.RightsMetadata.Copyright.Holder)
	}
	doc.CopyrightNotice = json.NewStringProperty("copyrightnotice", copyrightnotice)
}

func getBylines(doc *document) {
	nl := doc.Xml.NewsLines
	if nl.ByLine == nil || len(nl.ByLine) == 0 {
		return
	}

	bylines := json.Array{}
	edits := json.Array{}

	if nl.ByLineOriginal != nil || len(nl.ByLineOriginal) > 0 {
		for _, blo := range nl.ByLineOriginal {
			byline := json.Object{}
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
					producer := json.Object{}
					if bl.Id != "" {
						producer.AddString("code", bl.Id)
					}
					producer.AddString("name", bl.Value)
					doc.Producer = json.NewObjectProperty("producer", &producer)
				} else if strings.EqualFold(bl.Parametric, "PHOTOGRAPHER") && doc.Photographer == nil {
					photographer := json.Object{}
					if bl.Id != "" {
						photographer.AddString("code", bl.Id)
					}
					photographer.AddString("name", bl.Value)
					if bl.Title != "" {
						photographer.AddString("title", bl.Title)
					}
					doc.Photographer = json.NewObjectProperty("photographer", &photographer)

				} else if strings.EqualFold(bl.Parametric, "CAPTIONWRITER") && doc.CaptionWriter == nil {
					captionwriter := json.Object{}
					if bl.Id != "" {
						captionwriter.AddString("code", bl.Id)
					}
					captionwriter.AddString("name", bl.Value)
					if bl.Title != "" {
						captionwriter.AddString("title", bl.Title)
					}
					doc.CaptionWriter = json.NewObjectProperty("captionwriter", &captionwriter)
				} else if strings.EqualFold(bl.Parametric, "EDITEDBY") {
					edit := json.Object{}
					edit.AddString("name", bl.Value)
					edits.AddObject(&edit)
				} else {
					byline := json.Object{}
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
		doc.Bylines = json.NewArrayProperty("bylines", &bylines)
	}

	if edits.Length() > 0 {
		doc.Edits = json.NewArrayProperty("edits", &edits)
	}
}

func getPerson(doc *document) {
	nl := doc.Xml.NewsLines

	if nl.NameLine != nil && len(nl.NameLine) > 0 {
		persons := json.Array{}
		for _, name := range nl.NameLine {
			person := json.Object{}
			person.AddString("name", name.Value)
			if strings.EqualFold(name.Parametric, "PERSON_FEATURED") {
				rel := json.Array{}
				rel.AddString("personfeatured")
				person.AddArray("rel", &rel)
			}
			person.AddString("creator", "Editorial")
			persons.AddObject(&person)
		}
		doc.Person = json.NewArrayProperty("person", &persons)
	}
}
