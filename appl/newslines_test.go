package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
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
	doc, _ := parseXml(s)
	jo := json.Object{}

	err := doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	err = doc.ParseNewsLines(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Headline == "" {
		t.Error("[headline] is expected")
	}

	fmt.Printf("%s\n", jo.ToString())
}
