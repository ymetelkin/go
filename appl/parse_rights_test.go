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
	aj := ApplJson{Xml: pub}

	err := pub.RightsMetadata.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}

	if aj.UsageRights == nil {
		t.Error("[usagerights] is expected")
	}

	jo, err := aj.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}
