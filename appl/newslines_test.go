package appl

import (
	"testing"

	"github.com/ymetelkin/go/xml"
)

func TestNewslines(t *testing.T) {
	s := `
<Publication>
	<NewsLines>
		<Title>FBN--Vikings-Free Agency</Title>
		<HeadLine>Vikings, Sage Rosenfels agree to 2-year contract</HeadLine>
		<OriginalHeadLine>Vikings, Sage Rosenfels agree to 2-year contract</OriginalHeadLine>
		<ExtendedHeadLine>Agent: Vikings, Rosenfels agree in principle to 2-year contract to give team veteran backup QB</ExtendedHeadLine>
		<ByLine Title="AP Sports Writer">DAVE CAMPBELL</ByLine>
		<DateLine>EDEN PRAIRIE, Minn.</DateLine>
		<CopyrightLine>Copyright 2012 The Associated Press. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.</CopyrightLine>
		<KeywordLine>Vikings-Free Agency</KeywordLine>
		<ByLineOriginal>By DAVE CAMPBELL</ByLineOriginal>
		<NameLine Parametric="PERSON_FEATURED">Magdalena Neuner</NameLine>
	</NewsLines> 
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)

	doc.parseNewsLines(xml.Node("NewsLines"))

	if doc.Headline != "Vikings, Sage Rosenfels agree to 2-year contract" {
		t.Error("Invalid Headline")
	}
	if doc.Copyright.Notice != "Copyright 2012 The Associated Press. All rights reserved. This material may not be published, broadcast, rewritten or redistributed." {
		t.Error("Invalid Copyright.Notice")
	}
	if doc.Keywordlines[0] != "Vikings-Free Agency" {
		t.Error("Invalid Keywordlines")
	}
	if doc.Namelines[0].Name != "Magdalena Neuner" {
		t.Error("Invalid Persons.Nameline[0].Name")
	}
	if doc.Bylines[0].By != "By DAVE CAMPBELL" {
		t.Error("Invalid Bylines")
	}
}
