package appl

import (
	"fmt"
	"testing"
)

func TestManagement(t *testing.T) {
	s := `
<Publication>
	<PublicationManagement>
  	</PublicationManagement> 
</Publication>
`
	_, err := XmlToJson(s)
	if err == nil {
		t.Error("Must throw")
	} else {
		fmt.Printf("%s\n", err.Error())
	}

	s = `
<Publication>
	<PublicationManagement>
		<RecordType>Change</RecordType>
		<FilingType>Text</FilingType>
		<IsDistributionReady>true</IsDistributionReady>
		<ArrivalDateTime>2012-03-12T20:54:44</ArrivalDateTime>
		<FirstCreated UserAccount="APGBL" UserAccountSystem="APADS" UserName="APGBL\dcampbell"
			Year="2012" Month="3" Day="12" Time="20:54:44"/>
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
	pub, _ := NewXml(s)
	aj := ApplJson{Xml: pub}

	err = pub.PublicationManagement.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}

	if string(aj.PubStatus) != "usable" {
		t.Error("[pubstatus:usable] is expected")
	}
	if aj.FirstCreatedYear == 0 {
		t.Error("[firstcreated.year] is expected")
	}
	if aj.FirstCreated == nil {
		t.Error("[firstcreated] is expected")
	}
	if aj.RefersTo == nil {
		t.Error("[refersto] is expected")
	}
	if aj.Embargoed == nil {
		t.Error("[embargoed] is expected")
	}
	if aj.Signals.IsEmpty() {
		t.Error("[signals] is expected")
	}
	if aj.OutingInstructions == nil {
		t.Error("[outinginstructions] is expected")
	}
	if aj.EditorialTypes == nil {
		t.Error("[editorialtypes] is expected")
	}
	if aj.TimeRestrictions == nil || len(aj.TimeRestrictions) == 0 {
		t.Error("[timerestrictions] is expected")
	}
	if aj.Associations == nil {
		t.Error("[associations] is expected")
	}

	jo, err := aj.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}
