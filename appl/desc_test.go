package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
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
	doc, _ := parseXML([]byte(s))
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if ja, err := jo.GetArray("descriptions"); err != nil || ja.IsEmpty() {
		t.Error("[descriptions] is expected")
	}

	if ja, err := jo.GetObject("datelinelocation"); err != nil || ja.IsEmpty() {
		t.Error("[datelinelocation] is expected")
	}

	if ja, err := jo.GetArray("generators"); err != nil || ja.IsEmpty() {
		t.Error("[generators] is expected")
	}

	if ja, err := jo.GetArray("categories"); err != nil || ja.IsEmpty() {
		t.Error("[categories] is expected")
	}

	if ja, err := jo.GetArray("suppcategories"); err != nil || ja.IsEmpty() {
		t.Error("[suppcategories] is expected")
	}

	if ja, err := jo.GetArray("alertcategories"); err != nil || ja.IsEmpty() {
		t.Error("[alertcategories] is expected")
	}

	if ja, err := jo.GetArray("subjects"); err != nil || ja.IsEmpty() {
		t.Error("[subjects] is expected")
	}

	if ja, err := jo.GetArray("organizations"); err != nil || ja.IsEmpty() {
		t.Error("[organizations] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}

func TestServices(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
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
	</DescriptiveMetadata>
</Publication>`
	doc, _ := parseXML([]byte(s))
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if ja, err := jo.GetArray("services"); err != nil || ja.IsEmpty() {
		t.Error("[services] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}

func TestThirdParties(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
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
	doc, _ := parseXML([]byte(s))
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	ja, err := jo.GetArray("thirdpartymeta")
	if err != nil || ja.Length() == 0 {
		t.Error("[thirdpartymeta] is expected")
	}

	tpm, _ := ja.GetObject(0)
	if s, err := tpm.GetString("code"); err != nil || s == "" {
		t.Error("[thirdpartymeta.code] is expected")
	}
	if s, err := tpm.GetString("name"); err != nil || s == "" {
		t.Error("[thirdpartymeta.name] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}

func TestAudences(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="9add4649b53b4702ba7d9de5d4fa607a" Value="Online" ActualMatch="true">
				<Property Id="B6F34A252AF74F0EBCD885E6AC1057BE" Name="AudienceType" Value="AUDPLATFORM" />
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="82c6a4c46fa0446090a7acaf93159e4c" Value="Print" ActualMatch="true">
				<Property Id="B6F34A252AF74F0EBCD885E6AC1057BE" Name="AudienceType" Value="AUDPLATFORM" />
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="4b0ce2c82f2f41fdb1b0df183a5f0e43" Value="Broadcast" ActualMatch="true">
				<Property Id="B6F34A252AF74F0EBCD885E6AC1057BE" Name="AudienceType" Value="AUDPLATFORM" />
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="f5b16ea8760d10048047e6e7a0f4673e" Value="State" ActualMatch="true">
				<Property Id="317C913CF4AA4C5AB9DB610C250B8810" Name="AudienceType" Value="AUDSCOPE" />
			</Occurrence>
		</AudienceClassification>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="A0EED68882C6100487CDDF092526B43E" Value="New Jersey" ActualMatch="true">
				<Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY" />
			</Occurrence>
		</AudienceClassification>
	</DescriptiveMetadata>
	<FilingMetadata>
		<Id>71c947a7587e4a468fd55d3e125e02c2</Id>
		<ArrivalDateTime>2011-07-12T17:33:51</ArrivalDateTime>
		<Cycle>BC</Cycle>
		<Format>bx</Format>
		<Source>ny---</Source>
		<Category>n</Category>
	</FilingMetadata>
	<FilingMetadata>
		<Id>71c947a7587e4a468fd55d3e125e02c2</Id>
		<ArrivalDateTime>2011-07-12T17:33:51</ArrivalDateTime>
		<Cycle>BC</Cycle>
		<Format>bx</Format>
		<Source>nyc---</Source>
		<Category>n</Category>
    </FilingMetadata>
</Publication>`
	doc, _ := parseXML([]byte(s))
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if ja, err := jo.GetArray("audiences"); err != nil || ja.IsEmpty() {
		t.Error("[audiences] is expected")
	}

	fmt.Printf("%s\n", jo.String())

	s = `
<Publication>
	<DescriptiveMetadata>
		<AudienceClassification System="Editorial" Authority="AP Audience">
			<Occurrence Id="9add4649b53b4702ba7d9de5d4fa607a" Value="Online" ActualMatch="true">
				<Property Id="B6F34A252AF74F0EBCD885E6AC1057BE" Name="AudienceType" Value="AUDPLATFORM" />
			</Occurrence>
		</AudienceClassification>
	</DescriptiveMetadata>
	<FilingMetadata>
		<Id>71c947a7587e4a468fd55d3e125e02c2</Id>
		<ArrivalDateTime>2011-07-12T17:33:51</ArrivalDateTime>
		<Cycle>BC</Cycle>
		<Format>bx</Format>
		<Source>ny---</Source>
		<Category>n</Category>
	</FilingMetadata>
	<FilingMetadata>
		<Id>71c947a7587e4a468fd55d3e125e02c2</Id>
		<ArrivalDateTime>2011-07-12T17:33:51</ArrivalDateTime>
		<Cycle>BC</Cycle>
		<Format>bx</Format>
		<Source>nyc---</Source>
		<Category>n</Category>
    </FilingMetadata>
</Publication>`
	doc, _ = parseXML([]byte(s))
	jo = json.Object{}

	err = doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if ja, err := jo.GetArray("audiences"); err != nil || ja.IsEmpty() {
		t.Error("[audiences] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}
