package xml

import (
	"fmt"
	"testing"
	"time"
)

func TestXml(t *testing.T) {
	s := `
	<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<Publication xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" Version="5.3.1">
  <Identification>
    <ItemId>7701d7a83ca9c26854541e62e844961a</ItemId>
    <RecordId>0de4053ee31bac0167762f4356122dca</RecordId>
    <CompositeId>00000000000000000000000000000000</CompositeId>
    <CompositionType>StandardOnlineVideo</CompositionType>
    <MediaType>Video</MediaType>
    <DefaultLanguage>en</DefaultLanguage>
    <RecordSequenceNumber>0</RecordSequenceNumber>
    <FriendlyKey>1764098236</FriendlyKey>
  </Identification>
  <PublicationManagement>
    <RecordType>Change</RecordType>
    <FilingType>OnlineVideo</FilingType>
    <ItemKey>13  60Sec LON Video 20190625I</ItemKey>
    <IsDistributionReady>true</IsDistributionReady>
    <ArrivalDateTime>2019-10-01T16:08:33</ArrivalDateTime>
    <FirstCreated Day="26" Month="6" Time="10:24:16" Year="2019"/>
    <LastModifiedDateTime>2019-06-26T10:24:18</LastModifiedDateTime>
    <Status>Usable</Status>
    <EditorialId>13  60Sec LON Video 20190625I</EditorialId>
    <ItemStartDateTime>2019-06-26T10:24:16</ItemStartDateTime>
    <ItemStartDateTimeActual>2019-06-26T10:24:16</ItemStartDateTimeActual>
    <Function>OnlineVideo</Function>
    <IsDigitized>true</IsDigitized>
  </PublicationManagement>
  <NewsLines>
    <Title>Test for Story on 2606</Title>
    <HeadLine>Publishing API Headline733032</HeadLine>
    <OriginalHeadLine>Test Headlines </OriginalHeadLine>
    <DateLine>Now</DateLine>
    <CopyrightLine>Copyright 2019 The Associated Press. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.</CopyrightLine>
    <RightsLine>International: AP Clients Only | US: AP Clients Only / This video may not be edited</RightsLine>
    <LocationLine>London </LocationLine>
  </NewsLines>
  <AdministrativeMetadata>
    <Provider Id="apus" Type="AP">AP</Provider>
    <Source Type="AP">ASSOCIATED PRESS</Source>
    <TransmissionSource>EapVideo</TransmissionSource>
    <TransmissionSource>NotMemberFeed</TransmissionSource>
    <TransmissionSource>ENPS</TransmissionSource>
    <TransmissionSource>QAENPSWDC02</TransmissionSource>
    <TransmissionSource>JagRoars</TransmissionSource>
    <ProductSource>news-us</ProductSource>
    <ContentElement>VOSOT</ContentElement>
    <ConsumerReady>true</ConsumerReady>
    <Signal>ConsumerReady</Signal>
  </AdministrativeMetadata>
  <RightsMetadata/>
  <DescriptiveMetadata>
    <DateLineLocation>
      <City>Now Bahār</City>
      <CountryAreaName>Zabul</CountryAreaName>
      <CountryName>Afghanistan</CountryName>
      <LatitudeDD>32.217180000</LatitudeDD>
      <LongitudeDD>67.570730000</LongitudeDD>
    </DateLineLocation>
    <Comment>AP Video US</Comment>
  </DescriptiveMetadata>
  
  <FilingMetadata>
    <Id>89b8adf46194a617c0104aff54557b90</Id>
    <ArrivalDateTime>2019-06-26T10:25:05</ArrivalDateTime>
    <TransmissionContent>All</TransmissionContent>
    <SlugLine>Automation Test for Publishing API733032</SlugLine>
    <Products>
      <Product>101551</Product>
      <Product>100913</Product>
      <Product>100890</Product>
      <Product>100547</Product>
      <Product>100530</Product>
      <Product>100483</Product>
      <Product>100414</Product>
      <Product>100362</Product>
      <Product>100358</Product>
      <Product>100357</Product>
      <Product>100288</Product>
      <Product>46872</Product>
      <Product>46855</Product>
      <Product>46325</Product>
      <Product>46170</Product>
      <Product>46157</Product>
      <Product>46140</Product>
      <Product>45956</Product>
      <Product>45544</Product>
      <Product>45486</Product>
      <Product>45150</Product>
      <Product>43848</Product>
      <Product>41987</Product>
      <Product>40437</Product>
      <Product>37895</Product>
      <Product>35529</Product>
      <Product>8976</Product>
      <Product>1180</Product>
      <Product>6</Product>
      <Product>1</Product>
    </Products>
    <ForeignKeys System="APArchive">
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="ItemId" Id="a40062c5cadda3bbe4ad971be11bfdf9"/>
    </ForeignKeys>
    <ForeignKeys System="APOnline">
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="ItemId" Id="ef77fb618e4a4827b2f71639de637aa2"/>
    </ForeignKeys>
    <ForeignKeys System="Story">
      <Keys Field="ID" Id="Q3094033"/>
    </ForeignKeys>
    <ForeignKeys System="MOS">
      <Keys Field="mosID" Id="storyfiler.qa.wdc.fox.mos"/>
      <Keys Field="ncsID" Id="QAENPSWDC02"/>
      <Keys Field="roID" Id="QAENPSWDC02;P_YNEWSUS\W;AFB2779D-A138-4C2B-91042C1A49CD8B0F"/>
      <Keys Field="storyID" Id="QAENPSWDC02;P_YNEWSUS\W\R_AFB2779D-A138-4C2B-91042C1A49CD8B0F;4D362CB3-EC7C-4033-B9A705CF98682DB3"/>
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="binaryMosID" Id="BUS.GALLERY.ORIGIN-ONE.LON.AP.MOS"/>
      <Keys Field="objID" Id="e8b92c5c-cf4d-4d55-905a-d6c70035feb7"/>
    </ForeignKeys>
    <ForeignKeys System="mosPayload">
      <Keys Field="Approved" Id="1"/>
      <Keys Field="Creator" Id="JGAINES"/>
      <Keys Field="MediaTime" Id="0"/>
      <Keys Field="ModBy" Id="SKULKARNI"/>
      <Keys Field="ModTime" Id="20190626102136"/>
      <Keys Field="MOSItemDurations" Id="60"/>
      <Keys Field="MOSObjSlugs" Id="13  60Sec LON Video 20190625I"/>
      <Keys Field="MOSSlugs" Id="Test for Story on 2606-2"/>
      <Keys Field="Owner" Id="JGAINES"/>
      <Keys Field="pubApproved" Id="0"/>
      <Keys Field="Restrictions" Id="AP Clients Only"/>
      <Keys Field="SourceMediaTime" Id="0"/>
      <Keys Field="SourceModBy" Id="JGAINES"/>
      <Keys Field="SourceModTime" Id="20150109220247"/>
      <Keys Field="SourceTextTime" Id="0"/>
      <Keys Field="StoryProducer" Id="Shriraj "/>
      <Keys Field="TextTime" Id="6"/>
      <Keys Field="LegalBlock" Id="0"/>
      <Keys Field="SystemApprovedBy" Id=" SKULKARNI "/>
      <Keys Field="apslatebackgroundimage" Id="ap"/>
      <Keys Field="APStoryNumber" Id="Q3094033"/>
      <Keys Field="AspectRatio" Id="16:9"/>
      <Keys Field="AudioDescription" Id="English/Natsound"/>
      <Keys Field="Dateline" Id="London  - Now"/>
      <Keys Field="DomesticRestrictions" Id="AP Clients Only"/>
      <Keys Field="Headline" Id="Test Headlines "/>
      <Keys Field="InternationalRestrictions" Id="AP Clients Only"/>
      <Keys Field="ItemIdArchive" Id="a40062c5cadda3bbe4ad971be11bfdf9"/>
      <Keys Field="ItemIdOnline" Id="ef77fb618e4a4827b2f71639de637aa2"/>
      <Keys Field="Kill" Id="0"/>
      <Keys Field="LanguageType" Id="English/Natsound"/>
      <Keys Field="Location" Id="London "/>
      <Keys Field="LogicalAspectRatio" Id="16:9"/>
      <Keys Field="originalFileName" Id="\\CTCISILONRND.apgix.local\ipdnas03\Transcoding\TempFiles_QA\ENPSMediaSource\13  60Sec LON Video 20190625I.mov"/>
      <Keys Field="ProducedAspectRatio" Id="original"/>
      <Keys Field="ProductsUS" Id="AP Video US"/>
      <Keys Field="Date" Id="Now"/>
      <Keys Field="Source" Id="ASSOCIATED PRESS"/>
      <Keys Field="StoryFormat" Id="CR-VOSOT"/>
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="StoryNumberPrefix" Id="apus"/>
      <Keys Field="StorySummary" Id="Test Headlines "/>
      <Keys Field="StoryWorkgroup" Id="news-us"/>
      <Keys Field="VHCategoryCode" Id="News"/>
      <Keys Field="ENPSItemType" Id="3"/>
      <Keys Field="MetadataOnlyUpdate" Id="true"/>
    </ForeignKeys>
    <ForeignKeys System="DigView">
      <Keys Field="OriginalFileName" Id="13  60Sec LON Video 20190625I"/>
    </ForeignKeys>
    <ForeignKeys System="Fox">
      <Keys Field="functionid" Id="core-9103d3bc0a953af1e266394d644abe25-apvideous"/>
    </ForeignKeys>
  </FilingMetadata>
  <FilingMetadata>
    <Id>ea71b6711fe65824d19e479159020a63</Id>
    <ArrivalDateTime>2019-06-26T10:30:48</ArrivalDateTime>
    <TransmissionContent>All</TransmissionContent>
    <SlugLine>Test for Story on 2606</SlugLine>
    <Products>
      <Product>101551</Product>
      <Product>100913</Product>
      <Product>100890</Product>
      <Product>100547</Product>
      <Product>100530</Product>
      <Product>100483</Product>
      <Product>100414</Product>
      <Product>100362</Product>
      <Product>100358</Product>
      <Product>100357</Product>
      <Product>100288</Product>
      <Product>46872</Product>
      <Product>46855</Product>
      <Product>46325</Product>
      <Product>46170</Product>
      <Product>46157</Product>
      <Product>46140</Product>
      <Product>45956</Product>
      <Product>45544</Product>
      <Product>45486</Product>
      <Product>45150</Product>
      <Product>43848</Product>
      <Product>41987</Product>
      <Product>40437</Product>
      <Product>37895</Product>
      <Product>35529</Product>
      <Product>8976</Product>
      <Product>1180</Product>
      <Product>6</Product>
      <Product>1</Product>
    </Products>
    <ForeignKeys System="APArchive">
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="ItemId" Id="a40062c5cadda3bbe4ad971be11bfdf9"/>
    </ForeignKeys>
    <ForeignKeys System="APOnline">
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="ItemId" Id="ef77fb618e4a4827b2f71639de637aa2"/>
    </ForeignKeys>
    <ForeignKeys System="Story">
      <Keys Field="ID" Id="Q3094033"/>
    </ForeignKeys>
    <ForeignKeys System="MOS">
      <Keys Field="mosID" Id="storyfiler.qa.wdc.fox.mos"/>
      <Keys Field="ncsID" Id="QAENPSWDC02"/>
      <Keys Field="roID" Id="QAENPSWDC02;P_YNEWSUS\W;AFB2779D-A138-4C2B-91042C1A49CD8B0F"/>
      <Keys Field="storyID" Id="QAENPSWDC02;P_YNEWSUS\W\R_AFB2779D-A138-4C2B-91042C1A49CD8B0F;4D362CB3-EC7C-4033-B9A705CF98682DB3"/>
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="binaryMosID" Id="BUS.GALLERY.ORIGIN-ONE.LON.AP.MOS"/>
      <Keys Field="objID" Id="e8b92c5c-cf4d-4d55-905a-d6c70035feb7"/>
    </ForeignKeys>
    <ForeignKeys System="mosPayload">
      <Keys Field="Approved" Id="1"/>
      <Keys Field="Creator" Id="JGAINES"/>
      <Keys Field="MediaTime" Id="0"/>
      <Keys Field="ModBy" Id="SKULKARNI"/>
      <Keys Field="ModTime" Id="20190626102136"/>
      <Keys Field="MOSItemDurations" Id="60"/>
      <Keys Field="MOSObjSlugs" Id="13  60Sec LON Video 20190625I"/>
      <Keys Field="MOSSlugs" Id="Test for Story on 2606-2"/>
      <Keys Field="Owner" Id="JGAINES"/>
      <Keys Field="pubApproved" Id="0"/>
      <Keys Field="Restrictions" Id="AP Clients Only"/>
      <Keys Field="SourceMediaTime" Id="0"/>
      <Keys Field="SourceModBy" Id="JGAINES"/>
      <Keys Field="SourceModTime" Id="20150109220247"/>
      <Keys Field="SourceTextTime" Id="0"/>
      <Keys Field="StoryProducer" Id="Shriraj "/>
      <Keys Field="TextTime" Id="6"/>
      <Keys Field="LegalBlock" Id="0"/>
      <Keys Field="SystemApprovedBy" Id=" SKULKARNI "/>
      <Keys Field="apslatebackgroundimage" Id="ap"/>
      <Keys Field="APStoryNumber" Id="Q3094033"/>
      <Keys Field="AspectRatio" Id="16:9"/>
      <Keys Field="AudioDescription" Id="English/Natsound"/>
      <Keys Field="Dateline" Id="London  - Now"/>
      <Keys Field="DomesticRestrictions" Id="AP Clients Only"/>
      <Keys Field="Headline" Id="Test Headlines "/>
      <Keys Field="InternationalRestrictions" Id="AP Clients Only"/>
      <Keys Field="ItemIdArchive" Id="a40062c5cadda3bbe4ad971be11bfdf9"/>
      <Keys Field="ItemIdOnline" Id="ef77fb618e4a4827b2f71639de637aa2"/>
      <Keys Field="Kill" Id="0"/>
      <Keys Field="LanguageType" Id="English/Natsound"/>
      <Keys Field="Location" Id="London "/>
      <Keys Field="LogicalAspectRatio" Id="16:9"/>
      <Keys Field="originalFileName" Id="\\CTCISILONRND.apgix.local\ipdnas03\Transcoding\TempFiles_QA\ENPSMediaSource\13  60Sec LON Video 20190625I.mov"/>
      <Keys Field="ProducedAspectRatio" Id="original"/>
      <Keys Field="ProductsUS" Id="AP Video US"/>
      <Keys Field="Date" Id="Now"/>
      <Keys Field="Source" Id="ASSOCIATED PRESS"/>
      <Keys Field="StoryFormat" Id="CR-VOSOT"/>
      <Keys Field="StoryNumber" Id="Q3094033"/>
      <Keys Field="StoryNumberPrefix" Id="apus"/>
      <Keys Field="StorySummary" Id="Test Headlines "/>
      <Keys Field="StoryWorkgroup" Id="news-us"/>
      <Keys Field="VHCategoryCode" Id="News"/>
      <Keys Field="ENPSItemType" Id="3"/>
      <Keys Field="MetadataOnlyUpdate" Id="true"/>
    </ForeignKeys>
    <ForeignKeys System="DigView">
      <Keys Field="OriginalFileName" Id="13  60Sec LON Video 20190625I"/>
    </ForeignKeys>
    <ForeignKeys System="Fox">
      <Keys Field="functionid" Id="core-9103d3bc0a953af1e266394d644abe25-apvideous"/>
    </ForeignKeys>
  </FilingMetadata>
  <PublicationComponent MediaType="Video" Role="Main">
    <VideoContentItem ArrivedInFilingId="c09e35a2aa8e450eafa38fe1171cb6ec" BinaryPath="None" Id="82e44363b5a74164804a7ab7da4a9e82">
      <Characteristics FileExtension="txt" Format="IIM" MimeType="text/plain" SizeInBytes="0">
        <TotalDuration>60000</TotalDuration>
        <ProducedAspectRatio>original</ProducedAspectRatio>
      </Characteristics>
    </VideoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Text" Role="Caption">
    <TextContentItem ArrivedInFilingId="c09e35a2aa8e450eafa38fe1171cb6ec" Id="538f016eb1c04354935a3265d36383a8">
      <DataContent>
        <nitf>
          <body.content>
            <block>
              <p>Test Headlines </p>
            </block>
          </body.content>
        </nitf>
      </DataContent>
      <Characteristics FileExtension="xml" Format="NITF" MimeType="text/xml" SizeInBytes="0">
        <Words>0</Words>
      </Characteristics>
    </TextContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Text" Role="Script">
    <TextContentItem ArrivedInFilingId="c09e35a2aa8e450eafa38fe1171cb6ec" Id="52a64a5567f044f48cad85e55c9b1469">
      <DataContent>
        <nitf>
          <body.content>
            <block>
              <p>CAPTION GOES HERE (DATE) </p>
              <p>RESTRICTION SUMMARY: </p>
              <p>CLIENTS NOTE: VIDEO ONLY - SHOTLIST AND STORYLINE TO FOLLOW AS SOON AS POSSIBLE</p>
              <p>SHOTLIST:</p>
              <p>1.</p>
              <p>STORYLINE: </p>
            </block>
          </body.content>
        </nitf>
      </DataContent>
      <Characteristics FileExtension="xml" Format="NITF" MimeType="text/xml" SizeInBytes="0">
        <Words>0</Words>
      </Characteristics>
    </TextContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Video" Role="Preview">
    <VideoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="VideoArchive" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/secured/x030n.mp4" Id="db0a35391a49434194a0bae384ef4c92">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/secured/x030n.mp4</BinaryLocation>
      <Characteristics FileExtension="mp4" Format="MPEG" MimeType="video/mpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_x030n.mp4" SizeInBytes="0">
        <AverageBitRate>1500000</AverageBitRate>
        <TotalDuration>30600</TotalDuration>
        <Width>640</Width>
        <Height>360</Height>
        <VideoCoder>H.264</VideoCoder>
        <FrameRate>25.000000</FrameRate>
        <AspectRatio>16:9</AspectRatio>
        <SampleRate>1500000</SampleRate>
        <LanguageDescription>en</LanguageDescription>
        <LogicalAspectRatio>16:9</LogicalAspectRatio>
        <ProducedAspectRatio>original</ProducedAspectRatio>
      </Characteristics>
    </VideoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Video" Role="Main">
    <VideoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="VideoArchive" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/secured/x040n.mp4" Id="fb11168cd9f348ada288dc77a19dd49b">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/secured/x040n.mp4</BinaryLocation>
      <Characteristics FileExtension="mp4" Format="MPEG" MimeType="video/mpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_x040n.mp4" SizeInBytes="0">
        <AverageBitRate>1500000</AverageBitRate>
        <TotalDuration>30600</TotalDuration>
        <Width>640</Width>
        <Height>360</Height>
        <VideoCoder>H.264</VideoCoder>
        <FrameRate>25.000000</FrameRate>
        <AspectRatio>16:9</AspectRatio>
        <SampleRate>1500000</SampleRate>
        <LanguageDescription>en</LanguageDescription>
        <LogicalAspectRatio>16:9</LogicalAspectRatio>
        <ProducedAspectRatio>original</ProducedAspectRatio>
      </Characteristics>
    </VideoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Photo" Role="Thumbnail">
    <PhotoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="URL" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail.jpg" Id="1a940b415aa44c62af9987cf5066eed2">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail.jpg</BinaryLocation>
      <Characteristics Encoding="zip" FileExtension="jpg" Format="JPEG Baseline" MimeType="image/jpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f.jpg" SizeInBytes="0">
        <Width>105</Width>
        <Height>70</Height>
      </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Photo" Role="Preview">
    <PhotoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="URL" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/preview.jpg" Id="05ee068592b74f4d9f1ec780493c6d7f">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/preview.jpg</BinaryLocation>
      <Characteristics Encoding="zip" FileExtension="jpg" Format="JPEG Baseline" MimeType="image/jpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_1920x1080.jpg" SizeInBytes="0">
        <Width>1920</Width>
        <Height>1080</Height>
      </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Photo" Role="Thumbnail">
    <PhotoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="URL" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail01.jpg" Id="19007dd00e93498fb0e576da82c769e0">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail01.jpg</BinaryLocation>
      <Characteristics Encoding="zip" FileExtension="jpg" Format="JPEG Baseline" MimeType="image/jpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_092x069.jpg" SizeInBytes="0">
        <Width>92</Width>
        <Height>69</Height>
      </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Photo" Role="Preview">
    <PhotoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="URL" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/preview01.jpg" Id="e030129ee90444b8bed70791d92f2dc5">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/preview01.jpg</BinaryLocation>
      <Characteristics Encoding="zip" FileExtension="jpg" Format="JPEG Baseline" MimeType="image/jpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_400x300.jpg" SizeInBytes="0">
        <Width>400</Width>
        <Height>300</Height>
      </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Photo" Role="Thumbnail">
    <PhotoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="URL" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail02.jpg" Id="bddf477730d14af2aa5eeecccd7ffdde">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail02.jpg</BinaryLocation>
      <Characteristics Encoding="zip" FileExtension="jpg" Format="JPEG Baseline" MimeType="image/jpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_105x070.jpg" SizeInBytes="0">
        <Width>105</Width>
        <Height>70</Height>
      </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Photo" Role="Thumbnail">
    <PhotoContentItem ArrivedInFilingId="641d30c18b60472aa20600bfaaa5ddaf" BinaryPath="URL" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail03.jpg" Id="203bcc1baf0c4edcbaa0fb86c1a475eb">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/thumbnail03.jpg</BinaryLocation>
      <Characteristics Encoding="zip" FileExtension="jpg" Format="JPEG Baseline" MimeType="image/jpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_180x135.jpg" SizeInBytes="0">
        <Width>180</Width>
        <Height>135</Height>
      </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Video" Role="Main">
    <VideoContentItem ArrivedInFilingId="d3a97c2178984c8a96364f0e35a37456" BinaryPath="VideoArchive" Href="http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/secured/x090i.mp4" Id="534387906ffc41b895bb53de8a2fc97d">
      <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
      <BinaryLocation BinaryPath="URL" Sequence="1" To="9999-12-31T23:59:59">http://mrs.appl.qa.s3.amazonaws.com/ef77fb618e4a4827b2f71639de637aa2/components/secured/x090i.mp4</BinaryLocation>
      <Characteristics FileExtension="mp4" Format="MPEG" MimeType="video/mpeg" OriginalFileName="2337811aabdd4709a2c230dd174e2d4f_x090i.mp4" SizeInBytes="0">
        <AverageBitRate>10000000</AverageBitRate>
        <TotalDuration>30600</TotalDuration>
        <Width>1920</Width>
        <Height>1080</Height>
        <VideoCoder>H.264</VideoCoder>
        <FrameRate>29.970000</FrameRate>
        <AspectRatio>16:9</AspectRatio>
        <SampleRate>10000000</SampleRate>
        <LanguageDescription>en</LanguageDescription>
        <LogicalAspectRatio>16:9</LogicalAspectRatio>
        <ProducedAspectRatio>original</ProducedAspectRatio>
      </Characteristics>
    </VideoContentItem>
  </PublicationComponent>
</Publication>
`

	//s = `<a/>`

	start := time.Now()
	nd, err := ParseString(s)
	ts := time.Since(start)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", nd.InlineString())
	fmt.Printf("Duration: %v\n", ts)
}
