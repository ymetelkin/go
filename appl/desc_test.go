package appl

import (
	"testing"

	"github.com/ymetelkin/go/xml"
)

func TestDescriptions(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
		<Description>abc</Description>
		<Description>xyz</Description>
		<Description>abc</Description>
		<DateLineLocation>
			<City>Terre Haute</City>
			<CountryArea>IN</CountryArea>
			<CountryAreaName>Indiana</CountryAreaName>
			<Country>USA</Country>
			<CountryName>United States</CountryName>
			<LatitudeDD>39.4667000000</LatitudeDD>
			<LongitudeDD>-87.4139100000</LongitudeDD>	
		</DateLineLocation>
		<SubjectClassification SystemVersion="1" AuthorityVersion="3163" System="Teragram" Authority="AP Subject"></SubjectClassification>
		<SubjectClassification SystemVersion="1" AuthorityVersion="3164" System="Teragram"  Authority="AP Subject"></SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Supplemental Category Code">
			<Occurrence Id="SOC" Value="Soccer" />
		</SubjectClassification> 
		<SubjectClassification System="Edgil" Authority="AP Alert Category">
			<Occurrence Id="Political" />
		</SubjectClassification>
		<SubjectClassification System="Edgil" Authority="AP Alert Category">
			<Occurrence Id="Business" />
		</SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Subject">
			<Occurrence Id="f25af2d07e4e100484f5df092526b43e" Value="General news" ActualMatch="true"></Occurrence>
		</SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Category Code">
			<Occurrence Id="n" Value="n"></Occurrence>
		</SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Category Code">
			<Occurrence Id="a" Value="Hourly Newscast" /> 
		</SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Audio Cut Number Code"> 
			<Occurrence Value="900"/>
		</SubjectClassification>
		<SubjectClassification SystemVersion="1" AuthorityVersion="2810" System="Teragram" Authority="AP Subject">
			<Occurrence Count="4" Id="127b34388b6710048b56bd945080b03e" Confidence="1.000000" Value="Desktop and enterprise application software" ParentId="92b8a0d5e6fce547b72d3573bef61c8a" ActualMatch="true" Permission="Premium" />
		</SubjectClassification>
		<SubjectClassification SystemVersion="1" AuthorityVersion="2810" System="Teragram" Authority="AP Subject">
			<Occurrence Count="10" Id="455ef2b87df7100483d8df092526b43e" Confidence="1.000000" Value="Technology" ActualMatch="true" TopParent="true" />
		</SubjectClassification>
		<SubjectClassification SystemVersion="1" AuthorityVersion="2810" System="Teragram" Authority="AP Subject">
			<Occurrence Id="29dfdd288b6b10048ff4bd945080b03e" Value="Computing and information technology" ParentId="455ef2b87df7100483d8df092526b43e" ActualMatch="false" />
		</SubjectClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="3183" System="Teragram" Authority="AP Party">
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="3183" System="Teragram" Authority="AP Party">
		</EntityClassification> 
		<EntityClassification SystemVersion="1" AuthorityVersion="2375" System="Teragram" Authority="AP Party">
			<Occurrence Count="5" Id="b1d550d087874a0393ebfa85dab5ea0a" Value="Barack Obama">
				<Property Id="b2897c9372e741beb39ac1e67c14835f" Name="PartyType" Value="PERSON_FEATURED" ParentId="d188b8b8886b100481accb8225d5863e" />
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON" />
				<Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN" ParentId="d188b8b8886b100481accb8225d5863e" />
				<Position Value="Publication/NewsLines/HeadLine" Phrase="Barack Obama" />
				<Position Value="Publication/NewsLines/NameLine" Phrase="Barack Obama" />
				<Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Barack Obama" />
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="2375" System="Teragram" Authority="AP Party">
			<Occurrence Count="1" Value="M. Spencer Green">
				<Property Id="111a147611e548de93ad20a387d49200" Name="PartyType" Value="PHOTOGRAPHER" />
				<Position Value="Publication/NewsLines/ByLine" Phrase="M. Spencer Green" />
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="3183" System="Teragram" Authority="AP Party">
			<Occurrence Count="1" Id="7e4a5813691a4bdb8606638cc6d9d392" Value="Joe Webb">
				<Property Id="384682bd7b494bab97766d2ab7912388" Name="PartyType"
					Value="PROFESSIONAL_ATHLETE" ParentId="c474b8387e4e1004846ddf092526b43e"/>
				<Property Id="c474b8387e4e1004846ddf092526b43e" Name="PartyType" Value="SPORTS_FIGURE"
					ParentId="d188b8b8886b100481accb8225d5863e"/>
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
				<Property Id="76ddd855689d4e82a0033359bcbe5262" Name="Team" Value="Minnesota Vikings"
					Permission="Basic"/>
				<Property Name="extid" Value="FBN.24175" Permission="Basic"/>
				<Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Joe Webb"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="2132" System="Teragram" Authority="AP Party">
			<Occurrence Count="1" Id="b51dcec68af346999700ffe2ebaf25bd" Value="Haley Barbour">
				<Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN"
					ParentId="d188b8b8886b100481accb8225d5863e"/>
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
				<Property Id="b3dfffa882c4100486c3df092526b43e" Name="AssociatedState" Value="Mississippi"
					Permission="Premium"/>
				<Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Haley Barbour"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="1388" System="Teragram" Authority="AP Party">
            <Occurrence Count="1" Id="f78c840a607747f2a0c247c301b8cfcc" Value="Apolo Anton Ohno">
                <Property Id="a7b366852a2f4a708eb4e269c5beddab" Name="PartyType" Value="OLYMPIC_ATHLETE" ParentId="c474b8387e4e1004846ddf092526b43e"/>
                <Property Id="c474b8387e4e1004846ddf092526b43e" Name="PartyType" Value="SPORTS_FIGURE" ParentId="d188b8b8886b100481accb8225d5863e"/>
                <Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
                <Property Id="08a0a00882c810048942df092526b43e" Name="AssociatedState" Value="Washington" Permission="Basic"/>
                <Property Id="c1dff44882c710048903df092526b43e" Name="AssociatedState" Value="Utah" Permission="Basic"/>
                <Property Id="e3710475c4f242c5bea0272faf63cc2a" Name="AssociatedEvent" Value="2010 Vancouver Olympic Games" Permission="Basic"/>
                <Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Apolo Anton Ohno"/>
            </Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="2119" System="Teragram" Authority="AP Organization">
			<Occurrence Count="8" Id="d4b82a78857310048a38ff2260dd383e" Confidence="1.000000" Value="United States Senate" ParentId="86b5cdb87dac10048932ba7fa5283c3e" ActualMatch="true" />
		</EntityClassification>
		<EntityClassification System="PhraseFinder" Authority="AP Company">
			<Occurrence Id="ARMH" Value="ARM Holdings PLC" ParentId="C82299A88B2F1004859FBD945080B03E"/>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="3451" System="Teragram" Authority="AP Company">
			<Occurrence Count="1" Id="223003ce9d3d49659b41fdc0f56d1095" Value="ARM Holdings PLC">
				<Property Id="0ebbe896bf8a431aae2ffebbd8d5aaf9" Name="APIndustry" Value="Electronic parts manufacturing" ParentId="c82299a88b2f1004859fbd945080b03e" Permission="Basic"/>
				<Property Id="1e906e5ef5b64f5e835ab3a0c9a9be34" Name="Ticker" Value="ARM" ParentId="7365ee2ea4b5744fbfd55c229ca84715" Permission="Basic"/>
				<Property Id="59c6f47cdb704c77a25aeb07d60a7018" Name="Ticker" Value="ARMH" ParentId="b77b130b4cdf5a40951c2cbae9109e13" Permission="Basic"/>
				<Property Id="7365ee2ea4b5744fbfd55c229ca84715" Name="Exchange" Value="LSE" Permission="Basic"/>
				<Property Id="b77b130b4cdf5a40951c2cbae9109e13" Name="Exchange" Value="NASDAQ" Permission="Basic"/>
				<Property Id="e4d1fce6280e51d4b2378553d459ecfd" Name="Instrument" Value="LSE:ARM" Permission="Basic"/>
				<Property Id="728e24108a365baf09e00137fdd0269f" Name="Instrument" Value="NASDAQ:ARMH" Permission="Basic"/>
				<Property Id="c82299a88b2f1004859fbd945080b03e" Name="APIndustry" Value="Industrial products and services" ParentId="511dd198857510048d0fff2260dd383e" Permission="Basic"/>
				<Property Id="511dd198857510048d0fff2260dd383e" Name="APSubject" Value="Industries" ParentId="c8e409f8858510048872ff2260dd383e" Permission="Basic"/>
				<Property Id="c8e409f8858510048872ff2260dd383e" Name="APSubject" Value="Business" ParentId="5dd09e387dd310048b26df092526b43e" Permission="Basic"/>
				<Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="ARM Holdings PLC"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="277" System="Teragram" Authority="AP Company">
			<Occurrence Count="1" Id="223003ce9d3d49659b41fdc0f56d1095" Value="ARM HOLDINGS PLC">
				<Property Name="PrimaryTicker" Value="arm"/>
				<Property Name="extid" Value="c000047573"/>
				<Property Name="NAICS" Value="334413"/>
				<Property Id="45B96ADE4CEB4C3499262629BE8B711F" Name="Exchange" Value="XLON"/>
				<Property Id="0EBBE896BF8A431AAE2FFEBBD8D5AAF9" Name="APIndustry" Value="Electronic parts manufacturing"/>
				<Position  Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="ARM Holdings PLC"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="2038" System="Teragram" Authority="AP Geography">
			<Occurrence Count="32" Id="661f3a687d5b100482a4c076b8e3055c" Confidence="1.000000" Value="Venezuela" ParentId="661908787d5b10048204c076b8e3055c" ActualMatch="true">
				<Property Id="01f56e0e654841eca2e69bf2cbcc0526" Name="LocationType" Value="Nation" Permission="Premium" />
				<Property Name="CentroidLatitude" Value="8.0000000000" />
				<Property Name="CentroidLongitude" Value="-66.0000000000" />
			</Occurrence>
		</EntityClassification>
		<EntityClassification System="PhraseFinder" Authority="AP Country">
			<Occurrence Id="USA" Value="United States" />
		</EntityClassification>
		<EntityClassification Authority="AP Event">
			<Occurrence Id="57405" Value="Seattle Seahawks at San Francisco 49ers 11/26/2017">
				<Property Name="extid" Value="2017112608"/>
				<Property Name="extidsource" Value="NFL"/>
				<Property Name="ScheduledDateTime" Value="2017-11-26T21:05:00.000Z"/>
				<Property Name="CompetitorName1" Value="Seattle Seahawks"/>
				<Property Name="CompetitorAbbrv1" Value="SEA"/>
				<Property Name="CompetitorQualifier1" Value="Away"/>
				<Property Name="CompetitorName2" Value="San Francisco 49ers"/>
				<Property Name="CompetitorAbbrv2" Value="SF"/>
				<Property Name="CompetitorQualifier2" Value="Home"/>
				<Property Name="SeasonName" Value="NFL 2017 REG"/>
				<Property Name="TournamentName" Value="Seattle Seahawks at San Francisco 49ers 11/26/2017"/>
				<Property Name="VenueName" Value="Levi's Stadium"/>
				<Property Name="VenueCapacity" Value="68500"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification Authority="AP Event">
			<Occurrence Id="sr:match:12792850" Value="WC Qualification, UEFA">
				<Property Name="extidsource" Value="Sportradar"/>
				<Property Name="ScheduledDateTime" Value="2017-11-14T19:45:00.000Z"/>
				<Property Name="CompetitorName1" Value="Republic of Ireland"/>
				<Property Name="CompetitorAbbrv1" Value="IRL"/>
				<Property Name="CompetitorQualifier1" Value="home"/>
				<Property Name="CompetitorName2" Value="Denmark"/>
				<Property Name="CompetitorAbbrv2" Value="DEN"/>
				<Property Name="CompetitorQualifier2" Value="away"/>
				<Property Name="CompetitorCountryName2" Value="Denmark"/>
				<Property Name="CompetitorCountryCode2" Value="DNK"/>
				<Property Name="SeasonName" Value="WC Qualification 2018"/>
				<Property Name="SeasonStartDateTime" Value="2015-07-25T04:00:00.000Z"/>
				<Property Name="SeasonEndDateTime" Value="2017-11-19T05:00:00.000Z"/>
				<Property Name="TournamentName" Value="WC Qualification, UEFA"/>
				<Property Name="TournamentCategoryName" Value="International"/>
				<Property Name="TournamentSportName" Value="Soccer"/>
				<Property Name="RefereeName" Value="Marciniak, Szymon"/>
				<Property Name="RefereeNationality" Value="Poland"/>
				<Property Name="VenueName" Value="Aviva Stadium"/>
				<Property Name="VenueCapacity" Value="51700"/>
				<Property Name="VenueCity" Value="Dublin"/>
				<Property Name="VenueCountryName" Value="Ireland"/>
				<Property Name="VenueCountryCode" Value="IRL"/>
			</Occurrence>
		</EntityClassification>
				<EntityClassification System="Editorial" Authority="AP Event">
			<Occurrence Id="7f3c4a6d701a43fbbea586a70397154b" Value="Week 13"/>
		</EntityClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="9add4649b53b4702ba7d9de5d4fa607a" Value="Online" ActualMatch="true">
				<Property Id="B6F34A252AF74F0EBCD885E6AC1057BE" Name="AudienceType" Value="AUDPLATFORM"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="82c6a4c46fa0446090a7acaf93159e4c" Value="Print" ActualMatch="true">
				<Property Id="B6F34A252AF74F0EBCD885E6AC1057BE" Name="AudienceType" Value="AUDPLATFORM"></Property>
			</Occurrence>
		</AudienceClassification>
		<SalesClassification System="Editorial" Authority="AP Sales Code">
			<Occurrence Id="E8509C508A0F10048A48D56C852D093E" Value="Basic" />
		</SalesClassification>
		<SalesClassification System="Editorial" Authority="AP Sales Code">
			<Occurrence Id="F61F0E271160409AA80E05DDB4CF7C19" Value="Standard" />
		</SalesClassification>
		<SalesClassification System="Editorial" Authority="AP Sales Code">
			<Occurrence Id="E8509C508A0F10048A48D56C852D093E" Value="Basic" />
		</SalesClassification>
		<SalesClassification System="Editorial" Authority="AP Sales Code">
			<Occurrence Id="F61F0E271160409AA80E05DDB4CF7C19" Value="Standard" />
		</SalesClassification>
		<SalesClassification System="Editorial" Authority="AP Sales Code">
			<Occurrence Id="E8509C508A0F10048A48D56C852D093E" Value="Basic" />
		</SalesClassification>
		<Comment>Select</Comment>
		<Comment>Plus</Comment>
		<Comment>Strange</Comment>
		<ThirdPartyMeta Vocabulary="ExtendedDeskInfo" VocabularyOwner="http://cv.ap.org">
			<Occurrence Id="IPTCFormat" Value="false" />
		</ThirdPartyMeta>
		<ThirdPartyMeta Vocabulary="BWRegionKeywords" VocabularyOwner="http://businesswire.">
			<Occurrence Value="North America" />
		</ThirdPartyMeta>
		<ThirdPartyMeta Vocabulary="BWIndustryKeywords" VocabularyOwner="http://businesswire.">
			<Occurrence Value="Health" />
		</ThirdPartyMeta>
		<ThirdPartyMeta Vocabulary="IngestionManagerMeta" VocabularyOwner="cv.ap.org">
			<Occurrence Id="FeedID" Value="2940" />
		</ThirdPartyMeta>
		<ThirdPartyMeta Vocabulary="IngestionManagerMeta" VocabularyOwner="cv.ap.org">
			<Occurrence Id="FeedName" Value="Businesswire NewsML Feed" />
		</ThirdPartyMeta>
		<ThirdPartyMeta Vocabulary="apsubject" VocabularyOwner="cv.ap.org">
			<Occurrence Id="http://cv.ap.org/apsubject/F25AF2D07E4E100484F5DF092526B43E" Value="General News" />
		</ThirdPartyMeta>
	</DescriptiveMetadata>
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)

	doc.parseDescriptiveMetadata(xml.Node("DescriptiveMetadata"))

	if len(doc.Descriptions) != 2 {
		t.Error("Invalid Descriptions")
	}
	if doc.Descriptions[1] != "xyz" {
		t.Error("Invalid Descriptions[1]")
	}
	if doc.DateLineLocation.City != "Terre Haute" {
		t.Error("Invalid DateLineLocation.City")
	}
	if len(doc.Generators) != 5 {
		t.Error("Invalid Generators")
	}
	if doc.Generators[1].Code != "3164" {
		t.Error("Invalid Generators[1].Code")
	}
	if doc.Categories[1].Name != "Hourly Newscast" {
		t.Error("Invalid Categories[1].Name")
	}
	if doc.SuppCategories[0].Code != "SOC" {
		t.Error("Invalid SuppCategories[0].Code")
	}
	if doc.AlertCategories[1] != "Business" {
		t.Error("Invalid AlertCategories[1]")
	}
	if len(doc.Subjects) != 4 {
		t.Error("Invalid Subjects")
	}
	if doc.Subjects[2].Name != "Technology" {
		t.Error("Invalid Subjects[2].Name")
	}
	if doc.Fixture.Code != "900" {
		t.Error("Invalid Fixture.Code")
	}
}
