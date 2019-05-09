package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestSubjects(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
		<Description>A</Description>
		<Description>B</Description>
		<Description>A</Description>
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
	</DescriptiveMetadata>
</Publication>`
	doc, _ := parseXML([]byte(s))
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if _, err := jo.GetArray("descriptions"); err != nil {
		t.Error("[descriptions] is expected")
	}

	if _, err := jo.GetObject("datelinelocation"); err != nil {
		t.Error("[datelinelocation] is expected")
	}

	if _, err := jo.GetArray("generators"); err != nil {
		t.Error("[generators] is expected")
	}

	if _, err := jo.GetArray("categories"); err != nil {
		t.Error("[categories] is expected")
	}

	if _, err := jo.GetArray("suppcategories"); err != nil {
		t.Error("[suppcategories] is expected")
	}

	if _, err := jo.GetArray("alertcategories"); err != nil {
		t.Error("[alertcategories] is expected")
	}

	if _, err := jo.GetArray("subjects"); err != nil {
		t.Error("[subjects] is expected")
	}

	if _, err := jo.GetArray("organizations"); err != nil {
		t.Error("[organizations] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}
