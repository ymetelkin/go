package appl

import (
	"fmt"
	"testing"
)

func TestNewslines(t *testing.T) {
	s := `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Photo</MediaType>
	</Identification>
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
	</NewsLines>  
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.Identification.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	err = pub.NewsLines.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Headline.IsEmtpy() {
		t.Error("[headline] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}
