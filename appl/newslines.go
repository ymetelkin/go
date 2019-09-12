package appl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func (doc *document) ParseNewsLines(jo *json.Object) error {
	if doc.NewsLines == nil || doc.NewsLines.Nodes == nil || len(doc.NewsLines.Nodes) == 0 {
		return errors.New("NewsLines is missing")
	}

	var (
		summary     bool
		overs, keys uniqueArray
		bys, byos   []xml.Node
		cr          string
	)

	jo.AddString("headline", "")

	for _, nd := range doc.NewsLines.Nodes {
		switch nd.Name {
		case "Title":
			if nd.Text != "" {
				doc.Title = nd.Text
				jo.AddString("title", nd.Text)
			}
		case "HeadLine":
			if nd.Text != "" {
				doc.Headline = nd.Text
				jo.SetString("headline", nd.Text)
			}
		case "ExtendedHeadLine":
			if nd.Text != "" {
				jo.AddString("headline_extended", nd.Text)
			}
		case "BodySubHeader":
			if nd.Text != "" && !summary {
				jo.AddString("summary", nd.Text)
				summary = true
			}
		case "DateLine":
			if nd.Text != "" {
				jo.AddString("dateline", nd.Text)
			}
		case "CreditLine":
			if nd.Text != "" {
				jo.AddString("creditline", nd.Text)
			}
			id := nd.Attribute("Id")
			if id != "" {
				jo.AddString("creditlineid", id)
			}
		case "RightsLine":
			if nd.Text != "" {
				jo.AddString("rightsline", nd.Text)
			}
		case "CopyrightLine":
			cr = nd.Text
		case "SeriesLine":
			if nd.Text != "" {
				jo.AddString("seriesline", nd.Text)
			}
		case "OutCue":
			if nd.Text != "" {
				jo.AddString("outcue", nd.Text)
			}
		case "LocationLine":
			if nd.Text != "" {
				jo.AddString("locationline", nd.Text)
			}
		case "OverLine":
			overs.AddString(nd.Text)
		case "KeywordLine":
			keys.AddString(nd.Text)
		case "ByLineOriginal":
			if byos == nil {
				byos = []xml.Node{nd}
			} else {
				byos = append(byos, nd)
			}
		case "ByLine":
			if bys == nil {
				bys = []xml.Node{nd}
			} else {
				bys = append(bys, nd)
			}
		case "NameLine":
			if nd.Text != "" {
				person := json.Object{}
				person.AddString("name", nd.Text)
				if strings.EqualFold(nd.Attribute("Parametric"), "PERSON_FEATURED") {
					rel := json.Array{}
					rel.AddString("personfeatured")
					person.AddArray("rel", rel)
				}
				person.AddString("creator", "Editorial")
				if doc.Namelines == nil {
					doc.Namelines = []json.Object{person}
				} else {
					doc.Namelines = append(doc.Namelines, person)
				}
			}
		}
	}

	jo.AddProperty(overs.ToJSONProperty("overlines"))
	jo.AddProperty(keys.ToJSONProperty("keywordlines"))

	getBylines(bys, byos, jo)

	doc.SetCopyright(cr, jo)

	return nil
}

func (doc *document) SetHeadline(jo *json.Object) {
	var headline string

	t := doc.MediaType

	if t == mediaTypeText || t == mediaTypeComplexData {
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
			s := getFirstTenWords(jo)
			if s != "" {
				headline = s
			}
		}
	} else if t == mediaTypeVideo && (doc.Function == "" || !strings.EqualFold(doc.Function, "APTNLibrary")) {
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
			s := getFirstTenWords(jo)
			if s != "" {
				headline = s
			}
		}
	} else if t == mediaTypePhoto || t == mediaTypeGraphic || (t == mediaTypeVideo && strings.EqualFold(doc.Function, "APTNLibrary")) {
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
			headline = doc.Headline
		} else {
			s := getFirstTenWords(jo)
			if s != "" {
				headline = s
			}
		}
	} else if t == mediaTypeAudio {
		if doc.Headline != "" {
			return
		} else if doc.Title != "" {
			headline = doc.Title
		}
	}

	if headline != "" {
		doc.Headline = headline
		jo.SetString("headline", headline)
	}
}

func (doc *document) SetCopyright(s string, jo *json.Object) {
	var (
		holder string
		year   int
	)

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

	if s == "" && holder != "" && doc.FirstCreatedYear > 0 {
		s = fmt.Sprintf("Copyright %d %s. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.", doc.FirstCreatedYear, holder)
	}

	if s != "" {
		jo.AddString("copyrightnotice", s)
	}
	if holder != "" {
		jo.AddString("copyrightholder", holder)
	}

	if year > 0 {
		jo.AddInt("copyrightdate", year)
	}
}

func getBylines(bys []xml.Node, byos []xml.Node, jo *json.Object) {
	if bys == nil || len(bys) == 0 {
		return
	}

	bylines := json.Array{}
	edits := json.Array{}

	if byos != nil || len(byos) > 0 {
		for _, blo := range byos {
			if blo.Text != "" {
				byline := json.Object{}
				byline.AddString("by", blo.Text)
				title := blo.Attribute("Title")
				if title != "" {
					byline.AddString("title", title)
				} else {
					for _, bl := range bys {
						title = bl.Attribute("Title")
						if title != "" {
							byline.AddString("title", title)
							break
						}
					}
				}
				bylines.AddObject(byline)
			}
		}
	} else {
		var pr, ph, cw bool

		for _, bl := range bys {
			if bl.Text != "" {
				id, title, pm := getBylineAttributes(bl)

				if strings.EqualFold(title, "EditedBy") && !pr {
					producer := json.Object{}
					if id != "" {
						producer.AddString("code", id)
					}
					producer.AddString("name", bl.Text)
					jo.AddObject("producer", producer)
					pr = true
				} else if strings.EqualFold(pm, "PHOTOGRAPHER") && !ph {
					photographer := json.Object{}
					if id != "" {
						photographer.AddString("code", id)
					}
					photographer.AddString("name", bl.Text)
					if title != "" {
						photographer.AddString("title", title)
					}
					jo.AddObject("photographer", photographer)
					ph = true
				} else if strings.EqualFold(pm, "CAPTIONWRITER") && !cw {
					captionwriter := json.Object{}
					if id != "" {
						captionwriter.AddString("code", id)
					}
					captionwriter.AddString("name", bl.Text)
					if title != "" {
						captionwriter.AddString("title", title)
					}
					jo.AddObject("captionwriter", captionwriter)
					cw = true
				} else if strings.EqualFold(pm, "EDITEDBY") {
					edit := json.Object{}
					edit.AddString("name", bl.Text)
					edits.AddObject(edit)
				} else {
					byline := json.Object{}
					if id != "" {
						byline.AddString("code", id)
					}
					byline.AddString("by", bl.Text)
					if title != "" {
						byline.AddString("title", title)
					}
					if pm != "" {
						byline.AddString("parametric", pm)
					}
					bylines.AddObject(byline)
				}
			}
		}
	}

	if bylines.Length() > 0 {
		jo.AddArray("bylines", bylines)
	}

	if edits.Length() > 0 {
		jo.AddArray("edits", edits)
	}
}

func getBylineAttributes(nd xml.Node) (string, string, string) {
	var id, title, pm string

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
	return id, title, pm
}

func getPerson(ns []xml.Node, jo *json.Object) {
	if ns == nil || len(ns) == 0 {
		return
	}

	persons := json.Array{}
	var add bool
	for _, n := range ns {
		if n.Text != "" {
			add = true
			person := json.Object{}
			person.AddString("name", n.Text)
			if strings.EqualFold(n.Attribute("Parametric"), "PERSON_FEATURED") {
				rel := json.Array{}
				rel.AddString("personfeatured")
				person.AddArray("rel", rel)
			}
			person.AddString("creator", "Editorial")
			persons.AddObject(person)
		}
	}

	if add {
		jo.AddArray("person", persons)
	}
}

func getFirstTenWords(jo *json.Object) string {
	o, err := jo.GetObject("main")
	if err == nil {
		s, _ := o.GetString("nitf")
		if s != "" {
			toks := strings.Split(s, " ")
			if len(toks) > 10 {
				return strings.Join(toks[0:10], " ")
			}

			return s
		}
	}
	return ""
}
