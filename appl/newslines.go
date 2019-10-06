package appl

import (
	"strings"

	"github.com/ymetelkin/go/xml"
)

func (doc *Document) parseNewsLines(node xml.Node) {
	var (
		overlines, keywordlines  uniqueStrings
		bylines, bylinesoriginal []xml.Node
	)

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "Title":
			doc.Title = nd.Text
		case "HeadLine":
			doc.Headline = nd.Text
		case "BodySubHeader":
			if doc.Summary == "" {
				doc.Summary = nd.Text
			}
		case "ExtendedHeadLine":
			doc.ExtendedHeadline = nd.Text
		case "ByLine":
			bylines = append(bylines, nd)
		case "OverLine":
			overlines.Append(nd.Text)
		case "DateLine":
			doc.Dateline = nd.Text
		case "CreditLine":
			doc.Creditline = &CodeName{
				Code: nd.Attribute("Id"),
				Name: nd.Text,
			}
		case "CopyrightLine":
			if doc.Copyright == nil {
				doc.Copyright = &Copyright{}
			}
			doc.Copyright.Notice = nd.Text
		case "RightsLine":
			doc.Rightsline = nd.Text
		case "SeriesLine":
			doc.Seriesline = nd.Text
		case "KeywordLine":
			keywordlines.Append(nd.Text)
		case "OutCue":
			doc.OutCue = nd.Text
		case "NameLine":
			if nd.Text != "" {
				person := Person{
					Name:       nd.Text,
					IsNameline: true,
					IsFeatured: strings.EqualFold(nd.Attribute("Parametric"), "PERSON_FEATURED"),
				}
				doc.Persons = append(doc.Persons, person)
			}
		case "LocationLine":
			doc.Locationline = nd.Text
		case "ByLineOriginal":
			bylinesoriginal = append(bylinesoriginal, nd)
		}
	}

	if !overlines.IsEmpty() {
		doc.Overlines = overlines.Values()
	}

	if !keywordlines.IsEmpty() {
		doc.Keywordlines = keywordlines.Values()
	}

	doc.setBylines(bylines, bylinesoriginal)
}

func (doc *Document) setHeadline() {
	var headline string

	if doc.MediaType == "text" || doc.MediaType == "complexdata" {
		if doc.Headline != "" {
			return
		} else if doc.Title != "" {
			headline = doc.Title
		} else if doc.Filings != nil {
			for _, f := range doc.Filings {
				if (f.Category == "l" || f.Category == "s") && f.Slugline != "" {
					headline = f.Slugline
					break
				}
			}
		} else {
			headline = doc.getFirstTenWords()
			if headline == "" {
				return
			}
		}
	} else if doc.MediaType == "video" && (doc.Function == "" || !strings.EqualFold(doc.Function, "APTNLibrary")) {
		if doc.Headline != "" {
			return
		} else if doc.Title != "" {
			headline = doc.Title
		} else if doc.Filings != nil {
			for _, f := range doc.Filings {
				if f.Slugline != "" {
					headline = f.Slugline
					break
				}
			}
		} else {
			headline = doc.getFirstTenWords()
			if headline == "" {
				return
			}
		}
	} else if doc.MediaType == "photo" || doc.MediaType == "graphic" || (doc.MediaType == "video" && strings.EqualFold(doc.Function, "APTNLibrary")) {
		if doc.Title != "" {
			headline = doc.Title
		} else if doc.Filings != nil {
			for _, f := range doc.Filings {
				if f.Slugline != "" {
					headline = f.Slugline
					break
				}
			}
		} else if doc.Headline != "" {
			return
		} else {
			headline = doc.getFirstTenWords()
			if headline == "" {
				return
			}
		}
	} else if doc.MediaType == "audio" {
		if doc.Headline != "" {
			return
		} else if doc.Title != "" {
			headline = doc.Title
		}
	}

	if headline != "" {
		doc.Headline = headline
	}
}

func (doc *Document) setBylines(bylines []xml.Node, bylinesoriginal []xml.Node) {
	if bylines == nil || len(bylines) == 0 {
		return
	}

	if bylinesoriginal != nil || len(bylinesoriginal) > 0 {
		for _, bl := range bylinesoriginal {
			if bl.Text != "" {
				byline := Byline{
					By: bl.Text,
				}

				title := bl.Attribute("Title")
				if title != "" {
					byline.Title = title
				} else {
					for _, bl := range bylines {
						title = bl.Attribute("Title")
						if title != "" {
							byline.Title = title
							break
						}
					}
				}
				doc.Bylines = append(doc.Bylines, byline)
			}
		}
	} else {
		for _, bl := range bylines {
			if bl.Text != "" {
				id, title, pm := getBylineAttributes(bl)

				if strings.EqualFold(title, "EditedBy") && doc.Producer == nil {
					doc.Producer = &CodeName{
						Code: id,
						Name: bl.Text,
					}
				} else if strings.EqualFold(pm, "PHOTOGRAPHER") && doc.Photographer == nil {
					doc.Photographer = &CodeNameTitle{
						Code:  id,
						Name:  bl.Text,
						Title: title,
					}
				} else if strings.EqualFold(pm, "CAPTIONWRITER") && doc.Captionwriter == nil {
					doc.Captionwriter = &CodeNameTitle{
						Code:  id,
						Name:  bl.Text,
						Title: title,
					}
				} else if strings.EqualFold(pm, "EDITEDBY") {
					doc.Edits = append(doc.Edits, bl.Text)
				} else {
					byline := Byline{
						Code:       id,
						By:         bl.Text,
						Title:      title,
						Parametric: pm,
					}
					doc.Bylines = append(doc.Bylines, byline)
				}
			}
		}
	}
}

func getBylineAttributes(nd xml.Node) (id string, title string, pm string) {
	if nd.Attributes != nil {
		for _, a := range nd.Attributes {
			switch a.Name {
			case "Id":
				id = a.Value
			case "Title":
				title = a.Value
			case "Parametric":
				pm = a.Value
			}
		}
	}
	return
}

func (doc *Document) getFirstTenWords() string {
	var s string

	if doc.Story != nil && doc.Story.Body != "" {
		s = doc.Story.Body
	} else if doc.Caption != nil && doc.Caption.Body != "" {
		s = doc.Caption.Body
	}

	if s != "" {
		toks := strings.Split(s, " ")
		if len(toks) > 10 {
			s = strings.Join(toks[0:10], " ")
		}
	}

	return s
}
