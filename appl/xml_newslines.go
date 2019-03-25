package appl

import (
	"fmt"

	"github.com/ymetelkin/go/json"
)

type CreditLine struct {
	Id    string `xml:"Id,attr"`
	Value string `xml:",chardata"`
}

type Copyright struct {
	Holder string `xml:"Holder,attr"`
}

type XmlNewsLines struct {
	Title            string
	HeadLine         string
	OriginalHeadLine string
	BodySubHeader    []string
	ExtendedHeadLine string
	ByLine           string
	ByLineOriginal   string
	OverLine         []string
	DateLine         string
	CreditLine       CreditLine
	CopyrightLine    string
	Copyright        Copyright
	RightsLine       string
	SeriesLine       string
	KeywordLine      []string
	OutCue           string
	NameLine         string
	LocationLine     string
}

func (newslines *XmlNewsLines) ToJson(jo *json.JsonObject, year int) error {
	if newslines.Title != "" {
		jo.AddString("title", newslines.Title)
	}
	if newslines.HeadLine != "" {
		jo.AddString("headline", newslines.HeadLine)
	}
	if newslines.OriginalHeadLine != "" {
		jo.AddString("changeevent", newslines.OriginalHeadLine)
	}
	if len(newslines.BodySubHeader) > 0 {
		jo.AddString("summary", newslines.BodySubHeader[0])
	}
	if newslines.ExtendedHeadLine != "" {
		jo.AddString("headline_extended", newslines.ExtendedHeadLine)
	}
	if newslines.ByLine != "" {
		jo.AddString("changeevent", newslines.ByLine)
	}
	if newslines.ByLineOriginal != "" {
		jo.AddString("changeevent", newslines.ByLineOriginal)
	}
	if len(newslines.OverLine) > 0 {
		overlines := json.JsonArray{}
		for _, s := range newslines.OverLine {
			overlines.AddString(s)
		}
		jo.AddArray("overlines", &overlines)
	}
	if newslines.DateLine != "" {
		jo.AddString("dateline", newslines.DateLine)
	}
	if newslines.CreditLine.Value != "" {
		jo.AddString("creditline", newslines.CreditLine.Value)
	}
	if newslines.CreditLine.Id != "" {
		jo.AddString("creditlineid", newslines.CreditLine.Id)
	}
	if newslines.CopyrightLine != "" {
		jo.AddString("copyrightnotice", newslines.CopyrightLine)
	} else if year > 0 && newslines.Copyright.Holder != "" {
		s := fmt.Sprintf("Copyright %d %s. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.", year, newslines.Copyright.Holder)
		jo.AddString("copyrightnotice", s)
	}
	if newslines.RightsLine != "" {
		jo.AddString("rightsline", newslines.RightsLine)
	}
	if newslines.SeriesLine != "" {
		jo.AddString("seriesline", newslines.SeriesLine)
	}
	if len(newslines.KeywordLine) > 0 {
		keywordlines := json.JsonArray{}
		unique := make(map[string]bool)

		for _, s := range newslines.KeywordLine {
			_, ok := unique[s]
			if !ok {
				keywordlines.Add(s)
				unique[s] = true
			}
		}

		jo.AddArray("keywordlines", &keywordlines)
	}
	if newslines.OutCue != "" {
		jo.AddString("outcue", newslines.OutCue)
	}
	if newslines.NameLine != "" {
		jo.AddString("changeevent", newslines.NameLine)
	}
	if newslines.LocationLine != "" {
		jo.AddString("changeevent", newslines.LocationLine)
	}

	return nil
}
