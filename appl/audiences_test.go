package appl

import (
	"fmt"
	"testing"
)

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
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.DescriptiveMetadata.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Audiences == nil {
		t.Error("[audences] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}
