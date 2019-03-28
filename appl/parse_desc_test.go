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
	aj := ApplJson{Xml: pub}

	err := pub.DescriptiveMetadata.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}

	if aj.Descriptions.IsEmpty() {
		t.Error("[descriptions] is expected")
	}

	if pub.DescriptiveMetadata.DateLineLocation.City == "" {
		t.Error("[datelinelocation] is expected")
	}

	if aj.Generators == nil || len(aj.Generators) == 0 {
		t.Error("[generators] is expected")
	}

	if aj.Categories == nil || len(aj.Categories) == 0 {
		t.Error("[categories] is expected")
	}

	if aj.SuppCategories == nil || len(aj.SuppCategories) == 0 {
		t.Error("[suppcategories] is expected")
	}

	if aj.AlertCategories.IsEmpty() {
		t.Error("[alertcategories] is expected")
	}

	if aj.Subjects == nil || len(aj.Subjects) == 0 {
		t.Error("[subjects] is expected")
	}

	if aj.Organizations == nil || len(aj.Organizations) == 0 {
		t.Error("[organizations] is expected")
	}

	jo, err := aj.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}
