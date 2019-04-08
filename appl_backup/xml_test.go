package appl

import (
	"fmt"
	"testing"
)

func TestXmlParsing(t *testing.T) {
	s := GetTestXml()
	pub, err := NewXml(s)
	if err != nil {
		t.Error(err.Error())
	} else if pub.Identification.ItemId != "00000000000000000000000000000001" {
		t.Error("Missing: [ItemId]")
	} else if pub.Identification.RecordId != "00000000000000000000000000000002" {
		t.Error("Missing: [RecordId]")
	} else if pub.Identification.CompositeId != "00000000000000000000000000000003" {
		t.Error("Missing: [CompositeId]")
	} else if pub.Identification.CompositionType != "StandardPrintPhoto" {
		t.Error("Missing: [CompositionType]")
	} else if pub.Identification.MediaType != "Photo" {
		t.Error("Missing: [MediaType]")
	} else if pub.PublicationManagement.RecordType != "Change" {
		t.Error("Missing: [RecordType]")
	} else if pub.PublicationManagement.FilingType != "Text" {
		t.Error("Missing: [FilingType]")
	} else if pub.PublicationManagement.FirstCreated.Year != 2012 {
		t.Error("Missing: [FirstCreated.Year]")
	}

	s, _ = pub.ToString()
	fmt.Println(s)
}

func GetTestXml() string {
	return `
<Publication
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:xsd="http://www.w3.org/2001/XMLSchema"
   	Version="5.0.0.9"
	   xmlns="http://ap.org/schemas/03/2005/appl">
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Photo</MediaType>
		<Priority>5</Priority>
		<EditorialPriority>d</EditorialPriority>
		<DefaultLanguage>en-us</DefaultLanguage>
		<RecordSequenceNumber>2</RecordSequenceNumber>
		<FriendlyKey>18212677756771</FriendlyKey>
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
	<AdministrativeMetadata>
		<Provider Id="AP" Type="AP">AP</Provider>
		<Creator>AP</Creator>
		<Source Type="AP">AP</Source>
		<SourceMaterial Name="alternate">
		<Type>Website</Type>
		<Url>http://www.apnewsarchive.com/2017/Ext-Sum-September-1st-test2/id-050b1f04ddfe47acb12c3937fe5ebe4b</Url>
		</SourceMaterial>
		<TransmissionSource>EapText</TransmissionSource>
		<TransmissionSource>NotMemberFeed</TransmissionSource>
		<TransmissionSource>ElvisLives</TransmissionSource>
		<TransmissionSource>JagRoars</TransmissionSource>
		<ItemContentType System="Editorial" Id="b8db75de62a043a8bc5649b226b606dd">Spot Development</ItemContentType>
		<ContentElement>FullStory</ContentElement>
		<Reach Scheme="AP">HALO</Reach>
		<InPackage Scheme="APText">tophd inthd alhd cthd dehd ilhd iahd arhd</InPackage>
		<ConsumerReady>TRUE</ConsumerReady>
		<Property Name="EAI:SUBMISSIONPRIORITY"></Property>
		<Property Name="EAI:SLUGWORDCOUNT"></Property>
		<Property Name="EAI:ELVIS_CALLBACK_URL"></Property>
		<Property Name="EAI:ELVIS_WORKFLOW_ID"></Property>
    </AdministrativeMetadata>
</Publication>`
}
