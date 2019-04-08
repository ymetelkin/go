package appl

import (
	"errors"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (doc *document) ParseNewsLines(jo *json.Object) error {
	if doc.NewsLines == nil {
		return errors.New("NewsLines is missing")
	}

	var (
		summary       bool
		overs, keys   uniqueArray
		bys, byos, ns []Node
	)

	for _, nd := range doc.NewsLines.Nodes {
		switch nd.Name {
		case "Title":
			if nd.Text != "" {
				doc.Title = nd.Text
				jo.AddString("title", nd.Text)
			}
		case "HeadLine":
			jo.AddString("headline", nd.Text)
			if nd.Text != "" {
				doc.HeadLine = nd.Text
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
		case "RightsLine":
			if nd.Text != "" {
				jo.AddString("rightsline", nd.Text)
			}
		case "CopyrightLine":
			if nd.Text != "" {
				doc.CopyrightNotice = nd.Text
				jo.AddString("copyrightnotice", nd.Text)
			}
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
			if nd.Text != "" {
				if overs == nil {
					overs = uniqueArray{}
				}
				overs.AddString(nd.Text)
			}
		case "KeywordLine":
			if nd.Text != "" {
				if keys == nil {
					keys = uniqueArray{}
				}
				keys.AddString(nd.Text)
			}
		case "ByLineOriginal":
			if byos == nil {
				byos = []Node{nd}
			} else {
				byos = append(byos, nd)
			}
		case "ByLine":
			if bys == nil {
				bys = []Node{nd}
			} else {
				bys = append(byos, nd)
			}
		case "NameLine":
			if ns == nil {
				ns = []Node{nd}
			} else {
				ns = append(byos, nd)
			}
		}
	}

	if overs != nil {
		jo.AddProperty(overs.ToJsonProperty("overlines"))
	}

	if keys != nil {
		jo.AddProperty(keys.ToJsonProperty("keywordlines"))
	}

	getBylines(bys, byos, &jo)
	getPerson(ns, &jo)
}

func (doc *document) SetHeadline(jo *json.Object) {
	var headline string

	t := doc.MediaType

	if t == MEDIATYPE_TEXT || t == MEDIATYPE_AUDIO {
		if doc.HeadLine != "" {
			return
		} else if doc.Title != "" {
			headline = doc.Title
		}
	} else if t == MEDIATYPE_VIDEO && (doc.Function == "" || !strings.EqualFold(doc.Function, "APTNLibrary")) {
		if doc.HeadLine != "" {
			return
		} else if doc.Title != "" {
			headline = doc.Title
		}
	} else if t == MEDIATYPE_PHOTO || t == MEDIATYPE_GRAPHIC || (t == MEDIATYPE_VIDEO && strings.EqualFold(doc.Function, "APTNLibrary")) {
		if doc.Title != "" {
			headline = doc.Title
		} else if doc.HeadLine != "" {
			headline = doc.HeadLine
		}
	} else if t == MEDIATYPE_AUDIO {
		if doc.HeadLine != "" {
			return
		} else if doc.Title != "" {
			headline = doc.Title
		}
	}

	if headline != "" {
		jo.SetString("headline", headline)
	}

	doc.Headline = json.NewStringProperty("headline", headline)
}

func getBylines(bys []Node, byos []Node, jo *json.Object) {
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
				title := blo.GetAttribute("Title")
				if title != "" {
					byline.AddString("title", title)
				} else {
					for _, bl := range bys {
						title = bl.GetAttribute("Title")
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
					edits.AddObject(&edit)
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

func getBylineAttributes(nd Node) (string, string, string) {
	var id, title, pm string

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
	return id, title, pm
}

func getPerson(ns []Node, jo *json.Object) {
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
			if strings.EqualFold(n.GetAttribute("Parametric"), "PERSON_FEATURED") {
				rel := json.Array{}
				rel.AddString("personfeatured")
				person.AddArray("rel", rel)
			}
			person.AddString("creator", "Editorial")
			persons.AddObject(person)
		}
	}

	if add {
		jo.AddArray("person", &persons)
	}
}
