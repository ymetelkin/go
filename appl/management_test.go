package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func TestManagement(t *testing.T) {
	input := `
<Publication>
	<PublicationManagement>
		<RecordType>Change</RecordType>
		<FilingType>Text</FilingType>
		<IsDistributionReady>true</IsDistributionReady>
		<ArrivalDateTime>2012-03-12T20:54:44</ArrivalDateTime>
		<FirstCreated UserAccount="APGBL" UserAccountSystem="APADS" UserName="APGBL\wweissert" Year="2019" Month="4" Day="29" Time="13:07:36" ToolVersion="ELVIS 1.24.4.3" UserLocation="Austin, TX" UserWorkgroup="USA Central"></FirstCreated>
		<LastModifiedDateTime UserName="APGBL\dzelio" UserAccount="APGBL" UserAccountSystem="APADS"
			>2012-03-12T20:54:37</LastModifiedDateTime>
		<ReleaseDateTime>2012-03-12T20:54:44</ReleaseDateTime>
		<Status>Usable</Status>
		<SpecialInstructions>Eds: APNewsNow. Will be updated.</SpecialInstructions>
		<Instruction Type="Outing">INTERNET</Instruction>
		<Instruction Type="Outing">MOBILE</Instruction>
		<Instruction Type="Outing">INTERNET</Instruction>
		<RefersTo>YM</RefersTo>
		<Editorial LeadNum="0" AddNum="0">
			<Type>Advance</Type>
		</Editorial>
		<Editorial LeadNum="0" AddNum="0">
			<Type>YM</Type>
		</Editorial>
		<AssociatedWith LinkType="Item" CompositionType="StandardText" MetaLabel="x" MetaKey="a">00000000000000000000000000000000</AssociatedWith>
		<AssociatedWith LinkType="Item" CompositionType="StandardText" MetaLabel="y" MetaKey="b">00000000000000000000000000000010</AssociatedWith>
		<AssociatedWith LinkType="Item" CompositionType="StandardText" MetaLabel="x" MetaKey="a">00000000000000000000000000000020</AssociatedWith>
		<AssociatedWith LinkType="Item" CompositionType="StandardPrintPhoto" MetaLabel="y" MetaKey="b">00000000000000000000000000000030</AssociatedWith>
		<AssociatedWith LinkType="Item" CompositionType="StandardBroadcastPhoto" MetaLabel="x" MetaKey="a">00000000000000000000000000000040</AssociatedWith>
    	<AssociatedWith LinkType="Item" CompositionType="StandardIngestedContent" MetaLabel="y" MetaKey="b">00000000000000000000000000000050</AssociatedWith>
		<ItemStartDateTime>2012-03-12T20:54:44</ItemStartDateTime>
		<ItemStartDateTimeActual>2012-03-12T20:54:44</ItemStartDateTimeActual>
		<ExplicitWarning>1</ExplicitWarning>
		<IsDigitized>false</IsDigitized>
		<Destination Include="true">
			<Target>Edge</Target>
			<Target>MainFullText</Target>
			<Target>DistributionManager</Target>
			<Target>Alert</Target>
			<Target>Productizer</Target>
			<Target>Profiler</Target>
			<Target>Classification</Target>
			<Target>FASTElvis</Target>
		</Destination>
		<TimeRestrictions>
			<TimeRestriction System="NewsPowerDriveTime" Zone="Atlantic" Include="true"/>
			<TimeRestriction System="NewsPowerDriveTime" Zone="Eastern" Include="true"/>
			<TimeRestriction System="NewsPowerDriveTime" Zone="Central" Include="true"/>
			<TimeRestriction System="NewsPowerDriveTime" Zone="Mountain" Include="true"/>
		</TimeRestrictions>
	 </PublicationManagement>   
</Publication>
`

	expected := `
	{
		"recordtype": "Change",
		"filingtype": "Text",
		"arrivaldatetime": "2012-03-12T20:54:44Z",
		"firstcreated": "2019-04-29T13:07:36Z",
		"firstcreator": {
		  "username": "APGBL\\wweissert",
		  "useraccount": "APGBL",
		  "useraccountsystem": "APADS",
		  "toolversion": "ELVIS 1.24.4.3",
		  "userworkgroup": "USA Central",
		  "userlocation": "Austin, TX"
		},
		"lastmodifieddatetime": "2012-03-12T20:54:37Z",
		"lastmodifier": {
		  "username": "APGBL\\dzelio",
		  "useraccount": "APGBL",
		  "useraccountsystem": "APADS"
		},
		"releasedatetime": "2012-03-12T20:54:44Z",
		"pubstatus": "usable",
		"specialinstructions": "Eds:APNewsNow.Willbeupdated.",
		"refersto": "YM",
		"associations": [
		  {
			"type": "text",
			"itemid": "00000000000000000000000000000010",
			"representationtype": "partial",
			"associationrank": 1,
			"typerank": 1
		  },
		  {
			"type": "text",
			"itemid": "00000000000000000000000000000020",
			"representationtype": "partial",
			"associationrank": 2,
			"typerank": 2
		  },
		  {
			"type": "photo",
			"itemid": "00000000000000000000000000000030",
			"representationtype": "partial",
			"associationrank": 3,
			"typerank": 1
		  },
		  {
			"type": "photo",
			"itemid": "00000000000000000000000000000040",
			"representationtype": "partial",
			"associationrank": 4,
			"typerank": 2
		  },
		  {
			"itemid": "00000000000000000000000000000050",
			"representationtype": "partial",
			"associationrank": 5,
			"typerank": 1
		  }
		],
		"itemstartdatetime": "2012-03-12T20:54:44Z",
		"itemstartdatetimeactual": "2012-03-12T20:54:44Z",
		"embargoed": "2012-03-12T20:54:44Z",
		"editorialtypes": [
		  "Advance",
		  "YM"
		],
		"outinginstructions": [
		  "INTERNET",
		  "MOBILE"
		],
		"signals": [
		  "explicitcontent",
		  "isnotdigitized"
		],
		"newspowerdrivetimeatlantic": true,
		"newspowerdrivetimeeastern": true,
		"newspowerdrivetimecentral": true,
		"newspowerdrivetimemountain": true
	  }`

	xml, err := xml.ParseString(input)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)
	doc.XML = &xml
	doc.JSON = new(json.Object)

	doc.parsePublicationManagement(xml.Node("PublicationManagement"))
	//fmt.Println(doc.JSON.String())

	test, _ := json.ParseObjectString(expected)
	left := doc.JSON.InlineString()
	right := test.InlineString()
	if left != right {
		t.Error("Failed PublicationManagement")
		fmt.Println(left)
		fmt.Println(right)
	}
}
