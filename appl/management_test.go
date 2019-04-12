package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestManagement(t *testing.T) {
	s := `
<Publication>
	<PublicationManagement>
		<RecordType>Change</RecordType>
		<FilingType>Text</FilingType>
		<IsDistributionReady>true</IsDistributionReady>
		<ArrivalDateTime>2012-03-12T20:54:44</ArrivalDateTime>
		<FirstCreated Year="2019" Month="4" Day="12" Time="17:13:00"></FirstCreated>
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
	doc, _ := parseXML(s)
	jo := json.Object{}

	err := doc.ParsePublicationManagement(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", jo.ToString())

	if string(doc.PubStatus) != "usable" {
		t.Error("[pubstatus:usable] is expected")
	}
	if doc.FirstCreatedYear == 0 {
		t.Error("[firstcreated.year] is expected")
	}
	if _, err := jo.GetString("firstcreated"); err != nil {
		t.Error("[firstcreated] is expected")
	}
	if _, err := jo.GetString("refersto"); err != nil {
		t.Error("[refersto] is expected")
	}
	if _, err := jo.GetString("embargoed"); err != nil {
		t.Error("[embargoed] is expected")
	}
	if _, err := jo.GetArray("signals"); err != nil {
		t.Error("[signals] is expected")
	}
	if _, err := jo.GetArray("outinginstructions"); err != nil {
		t.Error("[outinginstructions] is expected")
	}
	if _, err := jo.GetArray("editorialtypes"); err != nil {
		t.Error("[editorialtypes] is expected")
	}
	if v, _ := jo.GetBool("newspowerdrivetimeatlantic"); !v {
		t.Error("[newspowerdrivetimeatlantic] is expected")
	}
	if _, err := jo.GetArray("associations"); err != nil {
		t.Error("[associations] is expected")
	}
}
