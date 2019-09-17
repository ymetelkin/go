package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func TestNewslines(t *testing.T) {
	input := `
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

	expected := `
	{
		"title": "FBN--Vikings-Free Agency",
		"headline": "Vikings, Sage Rosenfels agree to 2-year contract",
		"headline_extended": "Agent: Vikings, Rosenfels agree in principle to 2-year contract to give team veteran backup QB",
		"dateline": "EDEN PRAIRIE, Minn.",
		"copyrightnotice": "Copyright 2012 The Associated Press. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.",
		"keywordlines": [
		  "Vikings-Free Agency"
		],
		"person": [
		  {
			"name": "Magdalena Neuner",
			"rel": [
			  "personfeatured"
			],
			"creator": "Editorial"
		  }
		],
		"bylines": [
		  {
			"by": "By DAVE CAMPBELL",
			"title": "AP Sports Writer"
		  }
		]
	  }`

	xml, err := xml.ParseString(input)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)
	doc.XML = &xml
	doc.JSON = new(json.Object)

	doc.parseNewsLines(xml.Node("NewsLines"))
	doc.setCopyright()
	//fmt.Println(doc.JSON.String())

	test, _ := json.ParseObjectString(expected)
	left := doc.JSON.InlineString()
	right := test.InlineString()
	if left != right {
		t.Error("Failed NewsLines")
		fmt.Println(left)
		fmt.Println(right)
	}
}
