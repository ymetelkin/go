package appl

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	s := `
  <Publication Version="5.3.2" xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <Identification>
    <ItemId>40b5b5ffcb844458be58d43c4cb28731</ItemId>
    <RecordId>a48c594d08c1461b8f6d203584386b23</RecordId>
    <CompositeId>69c9016823eb4fa6a14a11c9bda0ebc5</CompositeId>
    <CompositionType>StandardText</CompositionType>
    <MediaType>Text</MediaType>
    <Priority>4</Priority>
    <EditorialPriority>r</EditorialPriority>
    <DefaultLanguage>en-us</DefaultLanguage>
    <RecordSequenceNumber>0</RecordSequenceNumber>
    <FriendlyKey>359094028847</FriendlyKey>
  </Identification>
  <PublicationManagement>
    <RecordType>Change</RecordType>
    <FilingType>Text</FilingType>
    <ItemKey>Text|0va-n|vab|AP|VA--Virginia Pronunciation Guide|DBIQQDTO0</ItemKey>
    <IsDistributionReady>true</IsDistributionReady>
    <ArrivalDateTime>2019-04-16T13:08:05</ArrivalDateTime>
    <FirstCreated Year="2019" Month="4" Day="16" Time="09:57:11"></FirstCreated>
    <Status>Usable</Status>
    <ItemStartDateTime>2019-04-16T13:08:05</ItemStartDateTime>
    <ItemStartDateTimeActual>2019-04-16T13:08:05</ItemStartDateTimeActual>
    <Destination Include="true">
      <Target>Alert</Target>
      <Target>DistributionManager</Target>
      <Target>Edge</Target>
      <Target>MainFullText</Target>
      <Target>FASTElvis</Target>
    </Destination>
    <TimeRestrictions>
      <TimeRestriction System="NewsPowerMorning" Zone="Eastern" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerMorning" Zone="Central" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerMorning" Zone="Mountain" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerMorning" Zone="Pacific" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerMorning" Zone="Alaska" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerMorning" Zone="Hawaii" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerMorning" Zone="Arizona" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerDriveTime" Zone="Eastern" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerDriveTime" Zone="Central" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerDriveTime" Zone="Mountain" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerDriveTime" Zone="Pacific" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerDriveTime" Zone="Alaska" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerDriveTime" Zone="Hawaii" Include="true"></TimeRestriction>
      <TimeRestriction System="NewsPowerDriveTime" Zone="Arizona" Include="true"></TimeRestriction>
    </TimeRestrictions>
  </PublicationManagement>
  <NewsLines>
    <Title>VA--Virginia Pronunciation Guide</Title>
    <HeadLine>Here's a guide to pronunciations in Virginia news</HeadLine>
    <OriginalHeadLine>Here's a guide to pronunciations in Virginia news</OriginalHeadLine>
    <CopyrightLine>Copyright 2019 The Associated Press. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.</CopyrightLine>
    <KeywordLine>VA--Virginia Pronunciation Guide</KeywordLine>
  </NewsLines>
  <AdministrativeMetadata>
    <Provider Id="AP" Type="AP">AP</Provider>
    <Creator>AP</Creator>
    <TransmissionSource>EapText</TransmissionSource>
    <TransmissionSource>NotMemberFeed</TransmissionSource>
    <TransmissionSource>JagRoars</TransmissionSource>
    <ConsumerReady>TRUE</ConsumerReady>
    <Property Name="EAI:SUBMISSIONPRIORITY"></Property>
    <Property Name="EAI:SLUGWORDCOUNT"></Property>
  </AdministrativeMetadata>
  <RightsMetadata>
    <Copyright Holder="AP" Date="2019"></Copyright>
  </RightsMetadata>
  <DescriptiveMetadata>
    <SubjectClassification System="Editorial" Authority="AP Category Code">
      <Occurrence Id="n" Value="n"></Occurrence>
    </SubjectClassification>
  </DescriptiveMetadata>
  <FilingMetadata>
    <Id>18e4b2fd8a1d4d05b4b1f77e4b78f059</Id>
    <ArrivalDateTime>2019-04-16T13:08:05</ArrivalDateTime>
    <Cycle>AP</Cycle>
    <TransmissionReference>v0252</TransmissionReference>
    <TransmissionContent>All</TransmissionContent>
    <ServiceLevelDesignator>v</ServiceLevelDesignator>
    <Selector>0va-n</Selector>
    <Format>bx</Format>
    <Source>vab</Source>
    <Category>n</Category>
    <Routing Type="Passcode" Expanded="false" Outed="false">0va-n</Routing>
    <SlugLine>AP-VA--Virginia Pronunciation Guide</SlugLine>
    <Products>
      <Product>101547</Product>
      <Product>101546</Product>
      <Product>100289</Product>
      <Product>100263</Product>
      <Product>100160</Product>
      <Product>100087</Product>
      <Product>100077</Product>
      <Product>46394</Product>
      <Product>46393</Product>
      <Product>46392</Product>
      <Product>46391</Product>
      <Product>46385</Product>
      <Product>46384</Product>
      <Product>46383</Product>
      <Product>46379</Product>
      <Product>46376</Product>
      <Product>46365</Product>
      <Product>46362</Product>
      <Product>46313</Product>
      <Product>46206</Product>
      <Product>46205</Product>
      <Product>46204</Product>
      <Product>46203</Product>
      <Product>46202</Product>
      <Product>46201</Product>
      <Product>46200</Product>
      <Product>46197</Product>
      <Product>46196</Product>
      <Product>46181</Product>
      <Product>46170</Product>
      <Product>46168</Product>
      <Product>46159</Product>
      <Product>46059</Product>
      <Product>45956</Product>
      <Product>45486</Product>
      <Product>45147</Product>
      <Product>44698</Product>
      <Product>43848</Product>
      <Product>43618</Product>
      <Product>41757</Product>
      <Product>39535</Product>
      <Product>39534</Product>
      <Product>39533</Product>
      <Product>39532</Product>
      <Product>39531</Product>
      <Product>39530</Product>
      <Product>39529</Product>
      <Product>39528</Product>
      <Product>33044</Product>
      <Product>3</Product>
      <Product>1</Product>
    </Products>
    <ForeignKeys System="Desk">
      <Keys Id="DBIQQDTO0" Field="AccessionNumber"></Keys>
    </ForeignKeys>
  </FilingMetadata>
  <PublicationComponent Role="Main" MediaType="Text">
    <TextContentItem Id="3d73597d6eaf445482ea45516d593aa2" ArrivedInFilingId="18e4b2fd8a1d4d05b4b1f77e4b78f059">
      <DataContent>
        <nitf>
          <body.content>
            <block>
              <p>PLACES:</p>
            </block>
          </body.content>
        </nitf>
      </DataContent>
      <Characteristics MimeType="text/xml" Format="IIM" FileExtension="xml" SizeInBytes="2516">
        <Words>310</Words>
      </Characteristics>
    </TextContentItem>
  </PublicationComponent>
</Publication>
  `
	jo, _ := XMLToJSON([]byte(s))

	fmt.Printf("%s\n", jo.String())

}
