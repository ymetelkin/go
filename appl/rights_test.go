package appl

import (
	"testing"

	"github.com/ymetelkin/go/xml"
)

func TestRights(t *testing.T) {
	s := `
<Publication>
	<RightsMetadata>
		<Copyright Holder="Copyright The Associated Press" Date="2011" />	
		<UsageRights>
			<UsageType>MobileNewsNetworkDistribution</UsageType>
			<Limitations>none</Limitations>
		</UsageRights>	
		<UsageRights>
			<UsageType>MobileNewsNetworkDistribution</UsageType>
		</UsageRights>
		<UsageRights>
			<UsageType>MarketplaceDistribution</UsageType>
			<Geography>NJ</Geography>
			<RightsHolder>DAVENPORT QUAD CITY TIMES</RightsHolder>
			<Limitations>none</Limitations>
			<Group Type="Corporate" Id="gs40274">Lee Enterprises</Group>    
		</UsageRights>
		<UsageRights>
			<UsageType>TimeRestriction</UsageType>
			<StartDate>2011-05-02T03:00:00+00:00</StartDate>
			<EndDate>2011-05-04</EndDate>
		</UsageRights>
		<UsageRights>
			<UsageType>TimeRestriction</UsageType>
			<StartDate>2011-05-02T03:10:22Z</StartDate>
			<EndDate>12/12/05</EndDate>
		</UsageRights>
	</RightsMetadata>
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)

	doc.parseRightsMetadata(xml.Node("RightsMetadata"))

	if doc.Copyright.Holder != "Copyright The Associated Press" {
		t.Error("Invalid Copyright.Holder")
	}
	if doc.Copyright.Year != 2011 {
		t.Error("Invalid Copyright.Year")
	}
	if doc.UsageRights[2].Geography[0] != "NJ" {
		t.Error("Invalid UsageRights[2].Geography[0]")
	}
	if doc.UsageRights[4].StartDate.Minute() != 10 {
		t.Error("Invalid UsageRights[4].StartDate.Minute()")
	}
	if doc.UsageRights[4].EndDate.Year() != 2005 {
		t.Error("Invalid UsageRights[4].EndDate.Year()")
	}
}
