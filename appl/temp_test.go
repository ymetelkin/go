package appl

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	s := `
	<Publication Version="5.3.1" xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<Identification>
	  <ItemId>e97307804d6e9ac088e81102b4ed00c0</ItemId>
	  <RecordId>523cbf74d1bbda76c9e3e4fb309b7824</RecordId>
	  <CompositeId>8679e104a2064922bcf38e2af724023b</CompositeId>
	  <CompositionType>StandardIngestedContent</CompositionType>
	  <MediaType>Photo</MediaType>
	  <Priority>4</Priority>
	  <EditorialPriority>r</EditorialPriority>
	  <DefaultLanguage>en-us</DefaultLanguage>
	  <RecordSequenceNumber>0</RecordSequenceNumber>
	  <FriendlyKey>56640036459174</FriendlyKey>
	</Identification>
	<PublicationManagement>
	  <RecordType>Change</RecordType>
	  <FilingType>PrintPhoto</FilingType>
	  <IsDistributionReady>true</IsDistributionReady>
	  <ArrivalDateTime>2019-04-12T17:14:13</ArrivalDateTime>
	  <FirstCreated Year="2019" Month="4" Day="12" Time="17:13:00"></FirstCreated>
	  <LastModifiedDateTime>2019-04-12T17:13:00</LastModifiedDateTime>
	  <Status>Usable</Status>
	  <ReleaseDateTime>2019-04-12T17:14:13</ReleaseDateTime>
	  <EditorialId>urn:publicid:prnewswire.com:20190412:LA17991:s1</EditorialId>
	  <ItemStartDateTime>2019-04-12T17:14:13</ItemStartDateTime>
	  <ItemStartDateTimeActual>2019-04-12T17:14:13</ItemStartDateTimeActual>
	  <ItemExpireDateTime Scope="Item">2019-04-19T17:14:13</ItemExpireDateTime>
	</PublicationManagement>
	<NewsLines>
	  <Title>Jacksons-CW Logo</Title>
	  <HeadLine>Jacksons-CW Logo</HeadLine>
	  <OriginalHeadLine>Jacksons-CW Logo</OriginalHeadLine>
	  <BodySubHeader>www.jacksonswash.com (PRNewsfoto/Jacksons Car Wash)</BodySubHeader>
	  <CopyrightLine>PR Newswire</CopyrightLine>
	</NewsLines>
	<AdministrativeMetadata>
	  <Provider Id="z00011" Type="ThirdParty">PR Newswire</Provider>
	  <Source City="New York" Country="USA" CountryArea="NY" Id="z00011" Url="http://www.prnewswire.com" Type="ThirdParty">PR Newswire</Source>
	  <TransmissionSource>NotMemberFeed</TransmissionSource>
	  <TransmissionSource>IngestionManager</TransmissionSource>
	  <TransmissionSource>Monarch</TransmissionSource>
	  <TransmissionSource>JagRoars</TransmissionSource>
	  <ProductSource>DataFeatures</ProductSource>
	  <ItemContentType>Press Release</ItemContentType>
	  <ConsumerReady>FALSE</ConsumerReady>
	</AdministrativeMetadata>
	<RightsMetadata>
	  <Copyright Holder="PR Newswire" Date="2019"></Copyright>
	</RightsMetadata>
	<DescriptiveMetadata>
	  <SubjectClassification System="Editorial" Authority="AP Category Code">
		<Occurrence Id="v" Value="v"></Occurrence>
	  </SubjectClassification>
	  <ThirdPartyMeta Vocabulary="IngestionManagerMetadata" VocabularyOwner="cv.ap.org">
		<Occurrence Id="EntryId/urn:publicid:prnewswire.com:20190412:LA17991:s1" Value="urn:publicid:prnewswire.com:20190412:LA17991:s1"></Occurrence>
	  </ThirdPartyMeta>
	  <ThirdPartyMeta Vocabulary="IngestionManagerMetadata" VocabularyOwner="cv.ap.org">
		<Occurrence Id="ManagementId/urn:publicid:prnewswire.com:2617991:s1" Value="urn:publicid:prnewswire.com:2617991:s1"></Occurrence>
	  </ThirdPartyMeta>
	  <ThirdPartyMeta Vocabulary="IngestionManagerMetadata" VocabularyOwner="cv.ap.org">
		<Occurrence Id="FeedID/8004" Value="8004"></Occurrence>
	  </ThirdPartyMeta>
	  <ThirdPartyMeta Vocabulary="IngestionManagerMetadata" VocabularyOwner="cv.ap.org">
		<Occurrence Id="FeedName/PRNewswireIATOMPreviewFeed" Value="PR Newswire IATOM Preview Feed"></Occurrence>
	  </ThirdPartyMeta>
	</DescriptiveMetadata>
	<FilingMetadata>
	  <Id>48edfda62a5b4758894360c85967a529</Id>
	  <ArrivalDateTime>2019-04-12T17:14:13</ArrivalDateTime>
	  <TransmissionContent>All</TransmissionContent>
	  <Source>prniatom</Source>
	  <Category>v</Category>
	  <SlugLine>Jacksons-CW Logo</SlugLine>
	  <Products>
		<Product>101536</Product>
		<Product>46199</Product>
		<Product>46198</Product>
		<Product>46174</Product>
		<Product>46160</Product>
		<Product>45955</Product>
		<Product>45486</Product>
		<Product>45148</Product>
		<Product>44696</Product>
		<Product>44695</Product>
		<Product>44653</Product>
		<Product>44569</Product>
		<Product>43848</Product>
		<Product>37583</Product>
		<Product>10003</Product>
		<Product>2</Product>
		<Product>1</Product>
	  </Products>
	  <ForeignKeys System="Member">
		<Keys Id="urn:publicid:prnewswire.com:20190412:LA17991:s1" Field="EntryId" IncludeInHash="Record"></Keys>
		<Keys Id="urn:publicid:prnewswire.com:2617991:s1" Field="ManagementId" IncludeInHash="Item"></Keys>
	  </ForeignKeys>
	  <DistributionScope>Local</DistributionScope>
	</FilingMetadata>
	<PublicationComponent Role="Caption" MediaType="Text">
	  <TextContentItem Id="0f8ee54813974cdfae16d237f2fa97dd" ArrivedInFilingId="48edfda62a5b4758894360c85967a529">
		<DataContent>
		  <nitf>
			<body.content>
			  <block>www.jacksonswash.com (PRNewsfoto/Jacksons Car Wash)</block>
			</body.content>
		  </nitf>
		</DataContent>
		<Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml" SizeInBytes="0"></Characteristics>
	  </TextContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Main" MediaType="Photo">
	  <PhotoContentItem Id="ca14fa5b8aef4806817088c4f65d8c49" Href="http://mrs.appl.qa.s3.amazonaws.com/e97307804d6e9ac088e81102b4ed00c0/components/versions/secured/main-v00.jpg" ArrivedInFilingId="48edfda62a5b4758894360c85967a529" BinaryPath="None">
		<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/e97307804d6e9ac088e81102b4ed00c0/components/versions/secured/main-v00.jpg</BinaryLocation>
		<Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="17101" OriginalFileName="Jacksons_CW_Logo.jpg">
		  <Width>400</Width>
		  <Height>162</Height>
		</Characteristics>
	  </PhotoContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Preview" MediaType="Photo">
	  <PhotoContentItem Id="4746a6033a524f049b2d74469a1ad487" Href="http://mrs.appl.qa.s3.amazonaws.com/e97307804d6e9ac088e81102b4ed00c0/components/versions/preview-v00.jpg" BinaryPath="None">
		<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/e97307804d6e9ac088e81102b4ed00c0/components/versions/preview-v00.jpg</BinaryLocation>
		<Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="12610" OriginalFileName="Jacksons_CW_Logo.jpg">
		  <Width>400</Width>
		  <Height>162</Height>
		</Characteristics>
	  </PhotoContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Thumbnail" MediaType="Photo">
	  <PhotoContentItem Id="c93a0ce294b4436d998b95da96bb72fb" Href="http://mrs.appl.qa.s3.amazonaws.com/e97307804d6e9ac088e81102b4ed00c0/components/versions/thumbnail-v00.jpg" BinaryPath="None">
		<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/e97307804d6e9ac088e81102b4ed00c0/components/versions/thumbnail-v00.jpg</BinaryLocation>
		<Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="2159" OriginalFileName="Jacksons_CW_Logo.jpg">
		  <Width>100</Width>
		  <Height>40</Height>
		</Characteristics>
	  </PhotoContentItem>
	</PublicationComponent>
  </Publication>`
	jo, _ := XMLToJSON(s)

	fmt.Printf("%s\n", jo.ToString())

}
