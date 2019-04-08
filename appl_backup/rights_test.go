package appl

import (
	"fmt"
	"testing"
)

func TestRights(t *testing.T) {
	s := `
<Publication>
	<RightsMetadata>
		<Copyright Holder="Copyright The Associated Press" Date="2011" />
			<UsageRights>
			<UsageType>MobileNewsNetworkDistribution</UsageType>
		</UsageRights>
		<UsageRights>
			<UsageType>MarketplaceDistribution</UsageType>
			<Geography>none</Geography>
			<RightsHolder>DAVENPORT QUAD CITY TIMES</RightsHolder>
			<Limitations>none</Limitations>
			<Group Type="Corporate" Id="gs40274">Lee Enterprises</Group>    
		</UsageRights>
		<UsageRights>
			<UsageType>TimeRestriction</UsageType>
			<StartDate>2011-05-02T03:00:00+00:00</StartDate>
			<EndDate>2011-05-04T03:00:00+00:00</EndDate>
		</UsageRights>
	</RightsMetadata>
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.RightsMetadata.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.UsageRights.IsEmtpy() {
		t.Error("[usagerights] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}
