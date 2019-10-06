package appl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ymetelkin/go/xml"
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
		<ForeignKeys System="Story">
			<Keys Field="ID" Id="Qtst147558" />
		</ForeignKeys>
		<ForeignKeys System="mosPayload">
			<Keys Field="StoryNumber" Id="Qtst147558" />
			<Keys Field="MOSObjSlugs" Id="08-05T1359_gvw 1359-1458" />
			<Keys Field="MOSItemDurations" Id="20" />
			<Keys Field="StoryFormat" Id="NR-VO" />
			<Keys Field="approvidermetadata" Id="Video_Heartbeat_efd58abf-5d4a-488b-b685-1c5e08242756" />
		</ForeignKeys>
		<ForeignKeys System="VideoHubG2">
			<Keys Field="service" Id="apservicecode:videoheartbeat" />
			<Keys Field="service" Id="apservicecode:1" />
			<Keys Field="service" Id="apservicecode:International" />
			<Keys Field="approvidermetadata" Id="Video_Heartbeat_efd58abf-5d4a-488b-b685-1c5e08242756" />
			<Keys Field="newsitemguid" Id="Qtst147558-text" />
		</ForeignKeys>
		<FilingCountry>United States</FilingCountry>
		<FilingSubject>General</FilingSubject>
		<BreakingNews>Breaking</BreakingNews>
	</FilingMetadata>
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)

	for _, nd := range xml.Nodes {
		doc.parseFilingMetadata(nd)
	}

	if len(doc.Filings) != 2 {
		t.Error("Invalid Filings")
	}
	if doc.Filings[0].ForeignKeys[0].Field != "storyid" {
		t.Error("Invalid Filings[0].ForeignKeys[0].Field")
	}
}

func temp() {
	name := "f"
	s := `ArrivalDateTime        *time.Time
	Cycle                  string
	TransmissionReference  string
	TransmissionFilename   string
	TransmissionContent    string
	ServiceLevelDesignator string
	Selector               string
	Format                 string
	Source                 string
	Category               string
	Routing                string
	Slugline               string
	OriginalMediaID        string
	ImportFolder           string
	ImportWarnings         string
	LibraryTwinCheck       string
	LibraryRequestID       string
	SpecialFieldAttn       string
	Feedline               string
	LibraryRequestLogin    string
	Products               string
	Priorityline           string
	ForeignKeys            []ForeignKey
	Country                string
	Region                 string
	Subject                string
	Topic                  string
	OnlineCode             string
	DistributionScope      string
	BreakingNews           string
	Style                  string
	Junkline               string`
	for _, tok := range strings.Split(s, "\n") {
		toks := strings.Split(strings.TrimSpace(tok), " ")
		f := strings.TrimSpace(toks[0])
		jo := strings.ToLower(f)
		fmt.Printf("if %s.%s != \"\" {\n\tjo.AddString(\"%s\", %s.%s)\n}\n", name, f, jo, name, f)
	}
}
