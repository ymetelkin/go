package xml

import (
	"fmt"
	"testing"
	"time"
)

func TestXml(t *testing.T) {
	s := `	
	<Publication xmlns="http://ap.org/schemas/03/2005/appl" Version="5.3.2">
	<Identification>
	  <ItemId>0d7d337e70684e039fb8b761d5989e3c</ItemId>
	  <RecordId>0999a4574d224f9cb435eef55a5ee7e6</RecordId>
	  <CompositeId>f034a956c7f64f76a73c273b8722e743</CompositeId>
	  <CompositionType>StandardText</CompositionType>
	  <MediaType>Text</MediaType>
	  <Priority>4</Priority>
	  <EditorialPriority>r</EditorialPriority>
	  <DefaultLanguage>de</DefaultLanguage>
	  <RecordSequenceNumber>0</RecordSequenceNumber>
	  <FriendlyKey>2855929986412794</FriendlyKey>
	</Identification>
	<PublicationManagement>
	  <RecordType>Change</RecordType>
	  <FilingType>Text</FilingType>
	  <IsDistributionReady>true</IsDistributionReady>
	  <ArrivalDateTime>2019-09-26T18:22:16</ArrivalDateTime>
	  <FirstCreated Day="26" Month="9" Time="22:21:47" Year="2019" />
	  <LastModifiedDateTime>2019-09-26T22:21:47</LastModifiedDateTime>
	  <Status>Usable</Status>
	  <ReleaseDateTime>2019-09-26T18:22:09</ReleaseDateTime>
	  <ItemStartDateTime>2019-09-26T18:22:09</ItemStartDateTime>
	  <ItemStartDateTimeActual>2019-09-26T18:22:09</ItemStartDateTimeActual>
	  <ItemExpireDateTime Scope="Item">2019-10-01T18:22:09</ItemExpireDateTime>
	</PublicationManagement>
	<NewsLines>
	  <Title>Youtube-Szene feiert in Berlin - mit Frank Elstner</Title>
	  <HeadLine>Youtube-Szene feiert in Berlin - mit Frank Elstner</HeadLine>
	  <OriginalHeadLine>Youtube-Szene feiert in Berlin - mit Frank Elstner</OriginalHeadLine>
	  <ByLine>bockc</ByLine>
	  <DateLine>Berlin (dpa) - </DateLine>
	  <CopyrightLine>(c) 2019 dpa Deutsche Presse Agentur GmbH</CopyrightLine>
	  <KeywordLine>Youtube-Szene feiert in Berlin - mit Frank Elstner</KeywordLine>
	  <KeywordLine>Internet</KeywordLine>
	</NewsLines>
	<AdministrativeMetadata>
	  <Provider Id="DPA9999" Type="ThirdParty">Deutsche Presse Agentur</Provider>
	  <Source City="Hamburg" Country="DEU" Id="DPA9999" Type="ThirdParty" Url="http://www.dpa.de">Deutsche Presse Agentur</Source>
	  <TransmissionSource>NotMemberFeed</TransmissionSource>
	  <TransmissionSource>Ingestion Manager</TransmissionSource>
	  <TransmissionSource>JagRoars</TransmissionSource>
	  <TransmissionSource>IngestionManager</TransmissionSource>
	  <TransmissionSource>Monarch</TransmissionSource>
	  <TransmissionSource>JagRoars</TransmissionSource>
	  <ProductSource>Basisdienst Inland</ProductSource>
	  <ProductSource>DataFeatures</ProductSource>
	  <ConsumerReady>TRUE</ConsumerReady>
	</AdministrativeMetadata>
	<RightsMetadata>
	  <Copyright Date="2019" Holder="Deutsche Presse-Agentur GmbH" />
	</RightsMetadata>
	<DescriptiveMetadata>
	  <DateLineLocation>
		<LatitudeDD>0.000000000</LatitudeDD>
		<LongitudeDD>0.000000000</LongitudeDD>
	  </DateLineLocation>
	  <ThirdPartyMeta Vocabulary="IngestionManagerMetadata" VocabularyOwner="cv.ap.org">
		<Occurrence Id="FeedID/5145" Value="5145" />
		<Occurrence Id="FeedName/DPANewsMessage" Value="DPA NewsMessage" />
		<Occurrence Id="TraceID/4981b31f5b03482d969be64b2e24deb4.0.0" Value="4981b31f5b03482d969be64b2e24deb4.0.0" />
	  </ThirdPartyMeta>
	</DescriptiveMetadata>
	<FilingMetadata>
	  <Id>bbeb1f08450f4b7fabd34d757fb64f0c</Id>
	  <ArrivalDateTime>2019-09-26T18:22:16</ArrivalDateTime>
	  <Cycle>BC</Cycle>
	  <TransmissionContent>All</TransmissionContent>
	  <ServiceLevelDesignator>dpa</ServiceLevelDesignator>
	  <Source>dpain</Source>
	  <SlugLine>Medien/Berlin/Deutschland</SlugLine>
	  <Products>
		<Product>101589</Product>
		<Product>101588</Product>
		<Product>101587</Product>
		<Product>101586</Product>
		<Product>101585</Product>
		<Product>101583</Product>
		<Product>101582</Product>
		<Product>101581</Product>
		<Product>101580</Product>
		<Product>101579</Product>
		<Product>101578</Product>
		<Product>101577</Product>
		<Product>101576</Product>
		<Product>101574</Product>
		<Product>101572</Product>
		<Product>101571</Product>
		<Product>101564</Product>
		<Product>101547</Product>
		<Product>101546</Product>
		<Product>47102</Product>
		<Product>47101</Product>
		<Product>47100</Product>
		<Product>46954</Product>
		<Product>46952</Product>
		<Product>46950</Product>
		<Product>46949</Product>
		<Product>46946</Product>
		<Product>46941</Product>
		<Product>46549</Product>
		<Product>46547</Product>
		<Product>46546</Product>
		<Product>46545</Product>
		<Product>46399</Product>
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
		<Product>45956</Product>
		<Product>45486</Product>
		<Product>45147</Product>
		<Product>44698</Product>
		<Product>44653</Product>
		<Product>43848</Product>
		<Product>3</Product>
		<Product>1</Product>
	  </Products>
	</FilingMetadata>
	<PublicationComponent MediaType="Text" Role="Main">
	  <TextContentItem ArrivedInFilingId="bbeb1f08450f4b7fabd34d757fb64f0c" Id="e68258256b064de98c2ef85c097aea4a">
		<DataContent>
		  <nitf>
			<body.content>
			  <block>
				<p>Berlin (dpa) - Von Fernsehgrößen wie Frank Elstner und Kai Pflaume über Die Fantastischen Vier bis zu Internetgrößen wie Ischtar Isik: In Berlin wurde am Donnerstagabend der Youtube Goldene Kamera Digital Award verliehen. Bei dem Internet-Preis gibt es acht Kategorien. Die erste Trophäe ging an Joseph DeChangeman in der Sparte Comedy &amp; Entertainment. </p>
				<p>Die Show im Kraftwerk wurde von Linda Zervakis («Tagesschau») und dem Comedian Daniele Rizzo moderiert. Zwei von Internet-Nutzern gekürte Preisträger standen schon fest: Wincent Weiss («Best Music Act») und Nintendo-Fachmann Domtendo («Let's Play &amp; Gaming»). Andere Gewinner wurden von einer Jury gekürt.</p>
				<p>Mit dem Digital Award würdigen die Funke Mediengruppe und Youtube die «originellsten und bewegendsten Internetproduktionen». Frank Elstner war mit seinem Youtube-Kanal («Wetten, das war's..?») in einer für ihn ungewöhnlichen Kategorie nominiert - als Newcomer. </p>
				<p />
				<p>#Notizblock</p>
				<p>##Redaktionelle Hinweise</p>
				<p>-Die Veranstaltung läuft noch. Sie erhalten bis 2200 einen Überblick und bis 2330 eine Zusammenfassung als Autorenbericht.</p>
				<p>##Internet</p>
				<p>-Über den Award</p>
				<p>##Orte</p>
				<p>-[Show im Kraftwerk](Köpenicker Str. 70, 10179 Berlin, Deutschland)</p>
				<p>Die folgenden Informationen sind nicht zur Veröffentlichung bestimmt</p>
				<p>##Ansprechpartner</p>
				<p>-Goldene Kamera, Presse, Jutta Rottmann, +49 162 234 66 18, &lt;presse@goldenekamera.de&gt;</p>
				<p>##Kontakte</p>
				<p>-Autorin: Caroline Bock (Berlin), +49 30 285231240, &lt;bock.caroline@dpa.com&gt;</p>
				<p>-Redaktion: Christof Bock (Berlin), +49 30 2852 32292, &lt;panorama@dpa.com&gt;</p>
				<p>-Foto: Newsdesk, +49 30 285231515, &lt;foto@dpa.com&gt;</p>
				<p>dpa ca yybb n1 bok</p>
			  </block>
			</body.content>
		  </nitf>
		</DataContent>
		<Characteristics FileExtension="xml" Format="IIM" MimeType="text/xml" SizeInBytes="1568">
		  <Words>195</Words>
		</Characteristics>
	  </TextContentItem>
	</PublicationComponent>
  </Publication>`

	//s = `<Identification><HeadLine>Vikings, Sage Rosenfels agree to 2-year contract</HeadLine></Identification>`

	start := time.Now()
	nd, err := ParseString(s)
	ts := time.Since(start)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", nd.InlineString())
	fmt.Printf("Duration: %v\n", ts)
}
