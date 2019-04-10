package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestFilings(t *testing.T) {
	s := `
<Publication>
	<FilingMetadata>
		<Id>05979aab5b7f45d69a20cebe9c6dfd59</Id>
		<ArrivalDateTime>2017-09-01T13:07:04</ArrivalDateTime>
		<Cycle>AP</Cycle>
		<TransmissionReference>V8086</TransmissionReference>
		<TransmissionFilename>http://elvisqa3.rnd.local/News/Stories/QA3-TOK-2017-Sep-01-000001/QA3-TOK-2017-Sep-01-000001.docx</TransmissionFilename>
		<TransmissionContent>All</TransmissionContent>
		<ServiceLevelDesignator>V</ServiceLevelDesignator>
		<Selector>-----</Selector>
		<Format>bx</Format>
		<Source>onfiler</Source>
		<Category>a</Category>
		<Routing Type="Prolog" Expanded="false" Outed="false">cabas los caol nvbas  an v</Routing>
		<Routing Type="Passcode" Expanded="false" Outed="false">caba us50s caba- us50</Routing>
		<Routing Type="GroupSID" Expanded="true" Outed="false">KY GA NY</Routing>
		<Routing Type="GroupSID" Expanded="false" Outed="false">NAM</Routing>
		<Routing Type="GroupSID" Expanded="false" Outed="true">ONLN</Routing>
		<Routing Type="SID" Expanded="true" Outed="true">CASAV UTSSN CASFO NYNDI TXDIS</Routing>
		<SlugLine>AP-September-1st-test2</SlugLine>
		<Products>	
			<Product>3</Product>
			<Product>1</Product>
		</Products>
		<ForeignKeys System="Story">
			<Keys Id="91316" Field="ID" />
		</ForeignKeys>
		<ForeignKeys System="MOS">
			<Keys Id="osm.archive1.lon.ap.mos" Field="mosID" />
			<Keys Id="91316" Field="objID" />
		</ForeignKeys>
		<FilingCountry>United States</FilingCountry>
		<FilingSubject>General</FilingSubject>
		<FilingOnlineCode>1110</FilingOnlineCode>
		<BreakingNews>Breaking</BreakingNews>
	</FilingMetadata>
	<FilingMetadata>
		<Id>0c40c0e29ce84d538a76fd2de489f1a3</Id>
		<ArrivalDateTime>2017-09-01T13:07:04</ArrivalDateTime>
		<Cycle>AP</Cycle>
		<TransmissionReference>V9746</TransmissionReference>
		<TransmissionFilename>http://elvisqa3.rnd.local/News/Stories/QA3-TOK-2017-Sep-01-000001/QA3-TOK-2017-Sep-01-000001.docx</TransmissionFilename>
		<TransmissionContent>All</TransmissionContent>
		<ServiceLevelDesignator>V</ServiceLevelDesignator>
		<Selector>1nt--</Selector>
		<Format>bx</Format>
		<Source>btrunk</Source>
		<Category>a</Category>
		<Routing Type="Passcode" Expanded="false" Outed="false">1nt--</Routing>
		<SlugLine>AP-September 1st test2</SlugLine>
		<Products>
			<Product>3</Product>
			<Product>1</Product>
		</Products>
		<ForeignKeys System="Desk">
		<Keys Id="DB6KLMU01" Field="AccessionNumber"></Keys>
		</ForeignKeys>
		<FilingCountry>United States</FilingCountry>
		<FilingSubject>General</FilingSubject>
		<BreakingNews>Breaking</BreakingNews>
	</FilingMetadata>
</Publication>`
	doc, _ := parseXml(s)
	jo := json.Object{}

	filings := json.Array{}
	for _, f := range doc.Filings {
		filings.AddObject(f.JSON)
	}
	jo.AddArray("filings", filings)

	if _, err := jo.GetArray("filings"); err != nil {
		t.Error("[filings] is expected")
	}

	fmt.Printf("%s\n", jo.ToString())
}
