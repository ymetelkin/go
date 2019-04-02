package appl

import (
	"fmt"
	"testing"
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
		<SubjectClassification System="Edgil" Authority="AP Alert Category">
			<Occurrence Id="Municipal" />
		</SubjectClassification>
		<SubjectClassification System="Edgil" Authority="AP Alert Category">
			<Occurrence Id="Travel" />
		</SubjectClassification>
		<SubjectClassification System="Edgil" Authority="AP Alert Category">
			<Occurrence Id="Legal" />
		</SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Category Code">
			<Occurrence Id="a" Value="a"></Occurrence>
		</SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Subject">
			<Occurrence Id="f25af2d07e4e100484f5df092526b43e" Value="General news" ActualMatch="true"></Occurrence>
		</SubjectClassification>
		<SubjectClassification System="Editorial" Authority="AP Category Code">
			<Occurrence Id="n" Value="n"></Occurrence>
		</SubjectClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="3183" System="Teragram" Authority="AP Party">
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="3183" System="Teragram" Authority="AP Party">
		</EntityClassification> 
		<EntityClassification SystemVersion="1" AuthorityVersion="2119" System="Teragram" Authority="AP Organization">
			<Occurrence Count="8" Id="d4b82a78857310048a38ff2260dd383e" Confidence="1.000000" Value="United States Senate" ParentId="86b5cdb87dac10048932ba7fa5283c3e" ActualMatch="true" />
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
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="4b0ce2c82f2f41fdb1b0df183a5f0e43" Value="Broadcast" ActualMatch="true">
				<Property Id="B6F34A252AF74F0EBCD885E6AC1057BE" Name="AudienceType" Value="AUDPLATFORM"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="f5b16ea8760d10048047e6e7a0f4673e" Value="State" ActualMatch="true">
				<Property Id="317C913CF4AA4C5AB9DB610C250B8810" Name="AudienceType" Value="AUDSCOPE"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="f43adc08760d10048040e6e7a0f4673e" Value="National" ActualMatch="true">
				<Property Id="317C913CF4AA4C5AB9DB610C250B8810" Name="AudienceType" Value="AUDSCOPE"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="b8099e4881d610048a11df092526b43e" Value="Alabama" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="a42dc0a082af100482a7df092526b43e" Value="Connecticut" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="bcadd4f882af100482c9df092526b43e" Value="Delaware" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="2c6a186082b010048379df092526b43e" Value="Illinois" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="1608ba1082b310048433df092526b43e" Value="Iowa" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="687e74a082af1004823adf092526b43e" Value="Arkansas" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="661e48387d5b10048291c076b8e3055c" Value="United States" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
			</Occurrence>
		</AudienceClassification>
	</DescriptiveMetadata>
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.DescriptiveMetadata.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Descriptions == nil {
		t.Error("[descriptions] is expected")
	}

	if pub.DescriptiveMetadata.DateLineLocation.City == "" {
		t.Error("[datelinelocation] is expected")
	}

	if doc.Generators == nil {
		t.Error("[generators] is expected")
	}

	if doc.Categories == nil {
		t.Error("[categories] is expected")
	}

	if doc.SuppCategories == nil {
		t.Error("[suppcategories] is expected")
	}

	if doc.AlertCategories == nil {
		t.Error("[alertcategories] is expected")
	}

	if doc.Subjects == nil {
		t.Error("[subjects] is expected")
	}

	if doc.Organizations == nil {
		t.Error("[organizations] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}

func TestPersons(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
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
		<EntityClassification SystemVersion="1" AuthorityVersion="3183" System="Teragram"
			Authority="AP Party">
			<Occurrence Count="1" Id="7e4a5813691a4bdb8606638cc6d9d392" Value="Joe Webb">
				<Property Id="384682bd7b494bab97766d2ab7912388" Name="PartyType" Value="PROFESSIONAL_ATHLETE" ParentId="c474b8387e4e1004846ddf092526b43e"/>
				<Property Id="c474b8387e4e1004846ddf092526b43e" Name="PartyType" Value="SPORTS_FIGURE" 	ParentId="d188b8b8886b100481accb8225d5863e"/>
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
				<Property Id="76ddd855689d4e82a0033359bcbe5262" Name="Team" Value="Minnesota Vikings" Permission="Basic"/>
				<Property Name="extid" Value="FBN.24175" Permission="Basic"/>
				<Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Joe Webb"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="2132" System="Teragram" Authority="AP Party">
			<Occurrence Count="1" Id="b51dcec68af346999700ffe2ebaf25bd" Value="Haley Barbour">
				<Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN" ParentId="d188b8b8886b100481accb8225d5863e"/>
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
				<Property Id="b3dfffa882c4100486c3df092526b43e" Name="AssociatedState" Value="Mississippi" Permission="Premium"/>
				<Position  Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Haley Barbour"/>
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
	</DescriptiveMetadata>
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.DescriptiveMetadata.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Persons == nil {
		t.Error("[persons] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}

func TestCompanies(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
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
	</DescriptiveMetadata>
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.DescriptiveMetadata.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Companies == nil {
		t.Error("[companies] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())

}
