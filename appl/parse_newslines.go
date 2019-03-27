package appl

import (
	"fmt"
	"strings"
)

func (nl *NewsLines) parse(aj *ApplJson) error {
	getHeadline(aj)
	getCopyrightNotice(aj)
	getBylines(aj)
	getPersons(aj)

	if nl.KeywordLine != nil {
		for _, kw := range nl.KeywordLine {
			aj.KeywordLines.Add(kw)
		}
	}

	return nil
}

func getHeadline(aj *ApplJson) {
	headline := ""

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

	aj.Headline = headline
}

func getCopyrightNotice(aj *ApplJson) {
	nl := aj.Xml.NewsLines
	if nl.CopyrightLine != "" {
		aj.CopyrightNotice = nl.CopyrightLine
	} else if aj.FirstCreatedYear > 0 && aj.Xml.RightsMetadata.Copyright.Holder != "" {
		aj.CopyrightNotice = fmt.Sprintf("Copyright %d %s. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.", aj.FirstCreatedYear, aj.Xml.RightsMetadata.Copyright.Holder)
	}
}

func getBylines(aj *ApplJson) {
	nl := aj.Xml.NewsLines
	if nl.ByLine == nil || len(nl.ByLine) == 0 {
		return
	}

	bylines := []ApplByline{}

	if nl.ByLineOriginal != nil || len(nl.ByLineOriginal) > 0 {
		for _, blo := range nl.ByLineOriginal {
			if blo.Value != "" {
				byline := ApplByline{Name: blo.Value}
				if blo.Title != "" {
					byline.Title = blo.Title
				} else {
					for _, bl := range nl.ByLine {
						if bl.Title != "" {
							byline.Title = bl.Title
							break
						}
					}
				}
				bylines = append(bylines, byline)
			}
		}
	} else {
		for _, bl := range nl.ByLine {
			if bl.Value != "" {
				if strings.EqualFold(bl.Title, "EditedBy") {
					producer := ApplByline{Name: bl.Value}
					if bl.Id != "" {
						producer.Code = bl.Id
					}
					aj.Producer = producer
				} else if strings.EqualFold(bl.Parametric, "PHOTOGRAPHER") && aj.Photographer.Name == "" {
					photographer := ApplByline{Name: bl.Value}
					if bl.Id != "" {
						photographer.Code = bl.Id
					}
					if bl.Title != "" {
						photographer.Title = bl.Title
					}
					aj.Photographer = photographer
				} else if strings.EqualFold(bl.Parametric, "CAPTIONWRITER") && aj.CaptionWriter.Name == "" {
					captionwriter := ApplByline{Name: bl.Value}
					if bl.Id != "" {
						captionwriter.Code = bl.Id
					}
					if bl.Title != "" {
						captionwriter.Title = bl.Title
					}
					aj.CaptionWriter = captionwriter
				} else if strings.EqualFold(bl.Parametric, "EDITEDBY") && aj.Editor.Name == "" {
					editor := ApplByline{Name: bl.Value}
					if bl.Id != "" {
						editor.Code = bl.Id
					}
					if bl.Title != "" {
						editor.Title = bl.Title
					}
					aj.Editor = editor
				} else {
					byline := ApplByline{Name: bl.Value}
					if bl.Id != "" {
						byline.Code = bl.Id
					}
					if bl.Title != "" {
						byline.Title = bl.Title
					}
					if bl.Parametric != "" {
						byline.Parametric = bl.Parametric
					}
					bylines = append(bylines, byline)
				}
			}
		}
	}

	if len(bylines) > 0 {
		aj.Bylines = bylines
	}
}

func getPersons(aj *ApplJson) {
	nl := aj.Xml.NewsLines

	if nl.NameLine != nil && len(nl.NameLine) > 0 {
		persons := []ApplPerson{}
		for _, name := range nl.NameLine {
			person := ApplPerson{Name: name.Value, IsFeatured: strings.EqualFold(name.Parametric, "PERSON_FEATURED")}
			persons = append(persons, person)
		}
		aj.Persons = persons
	}
}
