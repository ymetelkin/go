package appl

import (
	"testing"

	"github.com/ymetelkin/go/xml"
)

func TestManagement(t *testing.T) {
	s := `
<Publication>
	<PublicationManagement>
		<RecordType>Change</RecordType>
		<FilingType>Text</FilingType>
		<IsDistributionReady>true</IsDistributionReady>
		<ArrivalDateTime>2012-03-12T20:54:44</ArrivalDateTime>
		<FirstCreated UserAccount="APGBL" UserAccountSystem="APADS" UserName="APGBL\wweissert" Year="2019" Month="4" Day="29" Time="00:07:36" ToolVersion="ELVIS 1.24.4.3" UserLocation="Austin, TX" UserWorkgroup="USA Central"></FirstCreated>
		<LastModifiedDateTime UserName="APGBL\dzelio" UserAccount="APGBL" UserAccountSystem="APADS"
			>2012-03-12T20:54:37</LastModifiedDateTime>
		<ReleaseDateTime>2012-03-12T20:54:44</ReleaseDateTime>
		<Status>Usable</Status>
		<SpecialInstructions>Eds: APNewsNow. Will be updated.</SpecialInstructions>
		<Instruction Type="Outing">INTERNET</Instruction>
		<Instruction Type="Outing">MOBILE</Instruction>
		<Instruction Type="Outing">INTERNET</Instruction>
		<RefersTo>YM</RefersTo>
		<Editorial AddNum="0" LeadNum="2">
			<Type>Advance</Type>
			<Type>Lead</Type>
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

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)

	doc.parsePublicationManagement(xml.Node("PublicationManagement"))

	if doc.ArrivalDateTime.Year() != 2012 {
		t.Error("Invalid ArrivalDateTime")
	}
	if doc.Created.Year != 2019 {
		t.Error("Invalid Created.Year")
	}
	if doc.Copyright.Year != 2019 {
		t.Error("Invalid Copyright.Year")
	}
	if doc.Created.User.Name != "APGBL\\wweissert" {
		t.Error("Invalid Created.User.Name")
	}
	if doc.Associations[0].ItemID != "00000000000000000000000000000010" {
		t.Error("Invalid Associations[0].ItemID")
	}
	if doc.Associations[4].TypeRank != 2 {
		t.Error("Invalid Associations[4].TypeRank")
	}
}
