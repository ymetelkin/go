package appl

import (
	"encoding/xml"
	"strings"

	"github.com/ymetelkin/go/json"
)

type body struct {
	Blocks []block `xml:"block"`
}

type block struct {
	Html string `xml:",innerxml"`
}

func (tci *TextContentItem) parse(role string, doc *document) {
	ss := []string{}
	decoder := xml.NewDecoder(strings.NewReader(tci.Body.Xml))
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "body.content" {
				var b body
				decoder.DecodeElement(&b, &se)
				if b.Blocks != nil {
					for _, b := range b.Blocks {
						ss = append(ss, b.Html)
					}
				}
				break
			}
		}
	}

	if len(ss) > 0 {
		nitf := strings.Join(ss, "")
		nitf = makePrettyString(nitf)
		jo := json.Object{}
		jo.AddString("nitf", nitf)
		if tci.Words > 0 {
			jo.AddInt("words", tci.Words)
		}

		if doc.Texts == nil {
			doc.Texts = make(map[string]*json.Object)
		}

		doc.Texts[role] = &jo
	}
}
