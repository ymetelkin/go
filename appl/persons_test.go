package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestPersons(t *testing.T) {
	s := `
<Publication>
	<DescriptiveMetadata>
		<EntityClassification SystemVersion="1" AuthorityVersion="2375" System="Teragram" Authority="AP Party">
			<Occurrence Count="5" Id="b1d550d087874a0393ebfa85dab5ea0a" Value="Barack Obama">
				<Property Id="b2897c9372e741beb39ac1e67c14835f" Name="PartyType" Value="PERSON_FEATURED" ParentId="d188b8b8886b100481accb8225d5863e" />
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON" />
				<Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN" ParentId="d188b8b8886b100481accb8225d5863e" />
				<Position Value="Publication/NewsLines/HeadLine" Phrase="Barack Obama" />
				<Position Value="Publication/NewsLines/NameLine" Phrase="Barack Obama" />
				<Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Barack Obama" />
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="2375" System="Teragram" Authority="AP Party">
			<Occurrence Count="1" Value="M. Spencer Green">
				<Property Id="111a147611e548de93ad20a387d49200" Name="PartyType" Value="PHOTOGRAPHER" />
				<Position Value="Publication/NewsLines/ByLine" Phrase="M. Spencer Green" />
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="3183" System="Teragram"
			Authority="AP Party">
			<Occurrence Count="1" Id="7e4a5813691a4bdb8606638cc6d9d392" Value="Joe Webb">
				<Property Id="384682bd7b494bab97766d2ab7912388" Name="PartyType" Value="PROFESSIONAL_ATHLETE" ParentId="c474b8387e4e1004846ddf092526b43e"/>
				<Property Id="c474b8387e4e1004846ddf092526b43e" Name="PartyType" Value="SPORTS_FIGURE" 	ParentId="d188b8b8886b100481accb8225d5863e"/>
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
				<Property Id="76ddd855689d4e82a0033359bcbe5262" Name="Team" Value="Minnesota Vikings" Permission="Basic"/>
				<Property Name="extid" Value="FBN.24175" Permission="Basic"/>
				<Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Joe Webb"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="2132" System="Teragram" Authority="AP Party">
			<Occurrence Count="1" Id="b51dcec68af346999700ffe2ebaf25bd" Value="Haley Barbour">
				<Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN" ParentId="d188b8b8886b100481accb8225d5863e"/>
				<Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
				<Property Id="b3dfffa882c4100486c3df092526b43e" Name="AssociatedState" Value="Mississippi" Permission="Premium"/>
				<Position  Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Haley Barbour"/>
			</Occurrence>
		</EntityClassification>
		<EntityClassification SystemVersion="1" AuthorityVersion="1388" System="Teragram" Authority="AP Party">
            <Occurrence Count="1" Id="f78c840a607747f2a0c247c301b8cfcc" Value="Apolo Anton Ohno">
                <Property Id="a7b366852a2f4a708eb4e269c5beddab" Name="PartyType" Value="OLYMPIC_ATHLETE" ParentId="c474b8387e4e1004846ddf092526b43e"/>
                <Property Id="c474b8387e4e1004846ddf092526b43e" Name="PartyType" Value="SPORTS_FIGURE" ParentId="d188b8b8886b100481accb8225d5863e"/>
                <Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"/>
                <Property Id="08a0a00882c810048942df092526b43e" Name="AssociatedState" Value="Washington" Permission="Basic"/>
                <Property Id="c1dff44882c710048903df092526b43e" Name="AssociatedState" Value="Utah" Permission="Basic"/>
                <Property Id="e3710475c4f242c5bea0272faf63cc2a" Name="AssociatedEvent" Value="2010 Vancouver Olympic Games" Permission="Basic"/>
                <Position Value="Publication/PublicationComponent/TextContentItem/DataContent/nitf/body.content/block/p" Phrase="Apolo Anton Ohno"/>
            </Occurrence>
        </EntityClassification>
	</DescriptiveMetadata>
</Publication>`
	doc, _ := parseXml(s)
	jo := json.Object{}

	err := doc.ParseDescriptiveMetadata(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if _, err := jo.GetArray("persons"); err != nil {
		t.Error("[persons] is expected")
	}

	fmt.Printf("%s\n", jo.ToString())
}
