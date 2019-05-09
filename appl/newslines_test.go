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
		<NameLine Parametric="PERSON_FEATURED">Magdalena Neuner</NameLine>
	</NewsLines>  
	<DescriptiveMetadata>
		<EntityClassification SystemVersion="1" AuthorityVersion="2375" System="Teragram" Authority="AP Party">
			<Occurrence Count="1" Value="M. Spencer Green">
				<Property Id="111a147611e548de93ad20a387d49200" Name="PartyType" Value="PHOTOGRAPHER" />
				<Position Value="Publication/NewsLines/ByLine" Phrase="M. Spencer Green" />
			</Occurrence>
		</EntityClassification>
	</DescriptiveMetadata>
</Publication>`
	doc, _ := parseXML([]byte(s))
	jo := json.Object{}

	err := doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	err = doc.ParseNewsLines(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	err = doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Headline == "" {
		t.Error("[headline] is expected")
	}

	if _, err := jo.GetArray("persons"); err != nil {
		t.Error("[persons] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}
