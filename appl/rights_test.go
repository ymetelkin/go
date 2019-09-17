package appl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestRights(t *testing.T) {
	s := `
<Publication>
	<PublicationManagement>
		<RecordType>Change</RecordType>
		<FilingType>Text</FilingType>
		<IsDistributionReady>true</IsDistributionReady>
		<ArrivalDateTime>2012-03-12T20:54:44</ArrivalDateTime>
		<FirstCreated Year="2012" Month="3" Day="12" Time="20:54:44"/>
	</PublicationManagement>
	<NewsLines>
		<Title>FBN--Vikings-Free Agency</Title>
	</NewsLines>  
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
			<Geography>none</Geography>
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
	doc, _ := parseXML(strings.NewReader(s))
	jo := json.Object{}

	/*
		err := doc.ParsePublicationManagement(&jo)
		if err != nil {
			t.Error(err.Error())
		}
	*/

	err := doc.ParseNewsLines(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	err = doc.ParseRightsMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", jo.String())

	if _, err := jo.GetString("copyrightnotice"); err != nil {
		t.Error("[copyrightnotice] is expected")
	}

	if _, err := jo.GetString("copyrightholder"); err != nil {
		t.Error("[copyrightholder] is expected")
	}

	if _, err := jo.GetInt("copyrightdate"); err != nil {
		t.Error("[copyrightdate] is expected")
	}

	if _, err := jo.GetArray("usagerights"); err != nil {
		t.Error("[usagerights] is expected")
	}
}
