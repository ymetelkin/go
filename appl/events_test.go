package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestEvents(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
		<EntityClassification Authority="AP Event">
			<Occurrence Id="57405" Value="Seattle Seahawks at San Francisco 49ers 11/26/2017">
				<Property Name="extid" Value="2017112608"/>
				<Property Name="extidsource" Value="NFL"/>
				<Property Name="ScheduledDateTime" Value="2017-11-26T21:05:00.000Z"/>
				<Property Name="CompetitorName1" Value="Seattle Seahawks"/>
				<Property Name="CompetitorAbbrv1" Value="SEA"/>
				<Property Name="CompetitorQualifier1" Value="Away"/>
				<Property Name="CompetitorName2" Value="San Francisco 49ers"/>
				<Property Name="CompetitorAbbrv2" Value="SF"/>
				<Property Name="CompetitorQualifier2" Value="Home"/>
				<Property Name="SeasonName" Value="NFL 2017 REG"/>
				<Property Name="TournamentName" Value="Seattle Seahawks at San Francisco 49ers 11/26/2017"/>
				<Property Name="VenueName" Value="Levi's Stadium"/>
				<Property Name="VenueCapacity" Value="68500"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification Authority="AP Event">
			<Occurrence Id="sr:match:12792850" Value="WC Qualification, UEFA">
				<Property Name="extidsource" Value="Sportradar"/>
				<Property Name="ScheduledDateTime" Value="2017-11-14T19:45:00.000Z"/>
				<Property Name="CompetitorName1" Value="Republic of Ireland"/>
				<Property Name="CompetitorAbbrv1" Value="IRL"/>
				<Property Name="CompetitorQualifier1" Value="home"/>
				<Property Name="CompetitorName2" Value="Denmark"/>
				<Property Name="CompetitorAbbrv2" Value="DEN"/>
				<Property Name="CompetitorQualifier2" Value="away"/>
				<Property Name="CompetitorCountryName2" Value="Denmark"/>
				<Property Name="CompetitorCountryCode2" Value="DNK"/>
				<Property Name="SeasonName" Value="WC Qualification 2018"/>
				<Property Name="SeasonStartDateTime" Value="2015-07-25T04:00:00.000Z"/>
				<Property Name="SeasonEndDateTime" Value="2017-11-19T05:00:00.000Z"/>
				<Property Name="TournamentName" Value="WC Qualification, UEFA"/>
				<Property Name="TournamentCategoryName" Value="International"/>
				<Property Name="TournamentSportName" Value="Soccer"/>
				<Property Name="RefereeName" Value="Marciniak, Szymon"/>
				<Property Name="RefereeNationality" Value="Poland"/>
				<Property Name="VenueName" Value="Aviva Stadium"/>
				<Property Name="VenueCapacity" Value="51700"/>
				<Property Name="VenueCity" Value="Dublin"/>
				<Property Name="VenueCountryName" Value="Ireland"/>
				<Property Name="VenueCountryCode" Value="IRL"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification System="Editorial" Authority="AP Event">
			<Occurrence Id="7f3c4a6d701a43fbbea586a70397154b" Value="Week 13"/>
		</EntityClassification>
	</DescriptiveMetadata>
</Publication>`
	doc, _ := parseXML([]byte(s))
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if _, err := jo.GetArray("events"); err != nil {
		t.Error("[events] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}
