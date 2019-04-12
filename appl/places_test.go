package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestPlaces(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
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
	</DescriptiveMetadata>
</Publication>`
	doc, _ := parseXML(s)
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if _, err := jo.GetArray("places"); err != nil {
		t.Error("[places] is expected")
	}

	fmt.Printf("%s\n", jo.ToString())
}
