package appl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ymetelkin/go/json"
)

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
	doc, _ := parseXML(strings.NewReader(s))
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if _, err := jo.GetArray("companies"); err != nil {
		t.Error("[companies] is expected")
	}

	fmt.Printf("%s\n", jo.String())

}
