package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func (doc *Document) parseNewsLines(node xml.Node) {
	var (
		overlines, keywordlines  uniqueArray
		bylines, bylinesoriginal []xml.Node
	)

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "Title":
			if nd.Text != "" {
				doc.Title = nd.Text
				doc.JSON.AddString("title", nd.Text)
			}
		case "HeadLine":
			if nd.Text != "" {
				doc.Headline = nd.Text
				doc.JSON.AddString("headline", nd.Text)
			}
		case "BodySubHeader":
			if nd.Text != "" && doc.Summary == "" {
				doc.Summary = nd.Text
				doc.JSON.AddString("summary", nd.Text)
			}
		case "ExtendedHeadLine":
			if nd.Text != "" {
				doc.ExtendedHeadline = nd.Text
				doc.JSON.AddString("headline_extended", nd.Text)
			}
		case "ByLine":
			bylines = append(bylines, nd)
		case "OverLine":
			overlines.AddString(nd.Text)
		case "DateLine":
			if nd.Text != "" {
				doc.Dateline = nd.Text
				doc.JSON.AddString("dateline", nd.Text)
			}
		case "CreditLine":
			if nd.Text != "" {
				doc.Creditline = nd.Text
				doc.JSON.AddString("creditline", nd.Text)
			}
			id := nd.Attribute("Id")
			if id != "" {
				doc.JSON.AddString("creditlineid", id)
			}
		case "CopyrightLine":
			doc.Copyright = &Copyright{
				Notice: nd.Text,
			}
			doc.JSON.AddString("copyrightnotice", nd.Text)
			doc.JSON.AddString("copyrightholder", "")
			doc.JSON.AddInt("copyrightdate", 0)
		case "RightsLine":
			if nd.Text != "" {
				doc.Rightsline = nd.Text
				doc.JSON.AddString("rightsline", nd.Text)
			}
		case "SeriesLine":
			if nd.Text != "" {
				doc.Seriesline = nd.Text
				doc.JSON.AddString("seriesline", nd.Text)
			}
		case "KeywordLine":
			keywordlines.AddString(nd.Text)
		case "OutCue":
			if nd.Text != "" {
				doc.OutCue = nd.Text
				doc.JSON.AddString("outcue", nd.Text)
			}
		case "NameLine":
			if nd.Text != "" {
				person := Person{
					Name:       nd.Text,
					IsFeatured: strings.EqualFold(nd.Attribute("Parametric"), "PERSON_FEATURED"),
				}

				doc.Persons = append(doc.Persons, person)
			}
		case "LocationLine":
			if nd.Text != "" {
				doc.Locationline = nd.Text
				doc.JSON.AddString("locationline", nd.Text)
			}
		case "ByLineOriginal":
			bylinesoriginal = append(bylinesoriginal, nd)
		}
	}

	if !overlines.IsEmpty() {
		doc.Overlines = overlines.Values()
		doc.JSON.AddArray("overlines", overlines.JSONArray())
	}

	if !keywordlines.IsEmpty() {
		doc.Keywordlines = keywordlines.Values()
		doc.JSON.AddArray("keywordlines", keywordlines.JSONArray())
	}

	if doc.Persons != nil && len(doc.Persons) > 0 {
		var ja json.Array
		for _, p := range doc.Persons {
			ja.AddObject(p.json())
		}
		doc.JSON.AddArray("person", ja)
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
		doc.JSON.SetString("headline", headline)
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
					doc.Producer = &CodeNameTitle{
						Code: id,
						Name: bl.Text,
					}
					doc.JSON.AddObject("producer", doc.Producer.json())
				} else if strings.EqualFold(pm, "PHOTOGRAPHER") && doc.Photographer == nil {
					doc.Photographer = &CodeNameTitle{
						Code:  id,
						Name:  bl.Text,
						Title: title,
					}
					doc.JSON.AddObject("photographer", doc.Photographer.json())
				} else if strings.EqualFold(pm, "CAPTIONWRITER") && doc.Captionwriter == nil {
					doc.Captionwriter = &CodeNameTitle{
						Code:  id,
						Name:  bl.Text,
						Title: title,
					}
					doc.JSON.AddObject("captionwriter", doc.Captionwriter.json())
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

	if doc.Bylines != nil && len(doc.Bylines) > 0 {
		var ja json.Array
		for _, bl := range doc.Bylines {
			ja.AddObject(bl.json())
		}
		doc.JSON.AddArray("bylines", ja)
	}

	if doc.Edits != nil && len(doc.Edits) > 0 {
		var ja json.Array
		for _, e := range doc.Edits {
			var jo json.Object
			jo.AddString("name", e)
			ja.AddObject(jo)
		}
		doc.JSON.AddArray("edits", ja)
	}
}

func (cnt *CodeNameTitle) json() (jo json.Object) {
	if cnt.Code != "" {
		jo.AddString("code", cnt.Code)
	}
	jo.AddString("name", cnt.Name)
	if cnt.Title != "" {
		jo.AddString("title", cnt.Title)
	}
	return
}

func (bl *Byline) json() (jo json.Object) {
	if bl.Code != "" {
		jo.AddString("code", bl.Code)
	}
	jo.AddString("by", bl.By)
	if bl.Title != "" {
		jo.AddString("title", bl.Title)
	}
	if bl.Parametric != "" {
		jo.AddString("parametric", bl.Parametric)
	}
	return
}

func (doc *Document) setCopyright() {
	/*
		rmd := doc.RightsMetadata
		nd := rmd.Node("Copyright")
		if nd.Attributes != nil {
			for k, v := range nd.Attributes {
				switch k {
				case "Holder":
					holder = v
				case "Date":
					i, err := strconv.Atoi(v)
					if err == nil {
						year = i
					}
				}
			}
		}
	*/

	var (
		cr   Copyright
		year int
		rm   bool
	)

	if doc.Copyright != nil {
		cr = *doc.Copyright
		rm = true
	}

	if doc.Created != nil {
		year = doc.Created.Year
	}

	if cr.Notice == "" && cr.Holder != "" && year > 0 {
		cr.Notice = fmt.Sprintf("Copyright %d %s. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.", year, cr.Holder)
	}

	if cr.Notice != "" {
		if rm {
			doc.JSON.SetString("copyrightnotice", cr.Notice)
		} else {
			doc.JSON.AddString("copyrightnotice", cr.Notice)
		}
	} else if rm {
		doc.JSON.Remove("copyrightnotice")
	}

	if cr.Holder != "" {
		if rm {
			doc.JSON.SetString("copyrightholder", cr.Holder)
		} else {
			doc.JSON.AddString("copyrightholder", cr.Holder)
		}
	} else if rm {
		doc.JSON.Remove("copyrightholder")
	}

	if year > 0 {
		if rm {
			doc.JSON.SetInt("copyrightdate", year)
		} else {
			doc.JSON.AddInt("copyrightdate", year)
		}
	} else if rm {
		doc.JSON.Remove("copyrightdate")
	}
}

func getBylineAttributes(nd xml.Node) (id string, title string, pm string) {
	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "Id":
				id = v
			case "Title":
				title = v
			case "Parametric":
				pm = v
			}
		}
	}
	return
}

func (p *Person) json() (jo json.Object) {
	jo.AddString("name", p.Name)

	if p.IsFeatured {
		var ja json.Array
		ja.AddString("personfeatured")
		jo.AddArray("rel", ja)
	}

	jo.AddString("creator", "Editorial")

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
