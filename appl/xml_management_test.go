package appl

import (
	"fmt"
	"testing"
)

func TestManagement(t *testing.T) {
	s := `
<Publication 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
   	Version="5.0.0.9"
	   xmlns="http://ap.org/schemas/03/2005/appl">
	<Identification>
		<ItemId>1</ItemId>
		<RecordId>2</RecordId>
		<CompositeId>3</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Photo</MediaType>
	 </Identification> 
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
	jo, err := XmlToJson(s)
	if err != nil {
		t.Error(err.Error())
	} else if s, _ := jo.GetString("recordtype"); s != "Change" {
		t.Error("Mismatch: [recordtype]")
	} else if s, _ := jo.GetString("filingtype"); s != "Text" {
		t.Error("Mismatch: [type]")
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}
}

func TestManagementValidation(t *testing.T) {
	s := `
<Publication 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
   	Version="5.0.0.9"
	   xmlns="http://ap.org/schemas/03/2005/appl">
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
}
