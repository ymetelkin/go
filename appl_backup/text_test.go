package appl

import (
	"fmt"
	"testing"
)

func TestTexts(t *testing.T) {
	s := `
<Publication>
	<PublicationComponent Role="Caption" MediaType="Text">
		<TextContentItem Id="a4235d49caeb43388058613008a3e3e3" ArrivedInFilingId="802bc28ebcfe402890dfa9e98a82a086">
			<DataContent>
				<nitf xmlns="http://ap.org/schemas/03/2005/appl">
					<body.content>
						<block>
							<p>This undated photo courtesy of Sarah Dorio shows a tattoo on the arm of chef Hugh Acheson.</p>
							<p>Acheson has four tattoos. Once considered the province of sailors, bikers, ex-cons and, of course, college hipsters, tattoos have become standard attire in professional kitchens.  (AP Photo/Sarah Dorio)</p>
						</block>
						<block>
							<p>YM</p>
							<p>SV</p>
						</block>
					</body.content>
				</nitf>
			</DataContent>
			<Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml" SizeInBytes="293">
				<Words>47</Words>
			</Characteristics>
		</TextContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Script" MediaType="Text">
		<TextContentItem Id="cb1d185b245b448db822e38119ba8326" ArrivedInFilingId="bcab0148458e4de599115bba3c290c46">
			<DataContent>
				<nitf>
					<body.content>
						<block>
							<p>Nathan Horton’s second goal of the game keyed Boston’s three-goal third period as the Bruins beat the Winnipeg Jets 5-3 on Tuesday night (10 January).</p>
							<p>Andrew Ladd, former Bruin Blake Wheeler, and Eric Fehr scored for the Jets, who got two assists apiece from defensemen Zach Bogosian and Tobias Enstrom and ended a four-game road trip 1-3.</p>
							<p>SUGGESTED VOICE OVER: </p>
						</block>
					</body.content>
				</nitf>
			</DataContent>
			<Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml">
				<Words>236</Words>
			</Characteristics>
		</TextContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Shotlist" MediaType="Text">
		<TextContentItem Id="c43443cc048c4ad3bbaa7607e5670a92" ArrivedInFilingId="bcab0148458e4de599115bba3c290c46">
			<DataContent>
				<nitf>
					<body.content>
						<block>
							<p>TD Garden, Boston, Massachusetts, USA. 10 January 2012.</p>
							<p>First Period</p>
							<p>1. 00:00 Boston bench</p>
							<p>2. 00:05 Andrew Ladd scores for Winnipeg - Jets 1-0</p>
						</block>
					</body.content>
				</nitf>
			</DataContent>
			<Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml">
				<Words>104</Words>
			</Characteristics>
		</TextContentItem>
	</PublicationComponent>
	<PublicationComponent Role="PublishableEditorNotes" MediaType="Text">
		<TextContentItem Id="1a8f9e8c881b4e84983175ba4efa1294" ArrivedInFilingId="2b691b5b397547dba93bc6f82933d4ed">
			<DataContent>
				<nitf>
					<body.content>
						<block>
							<p>This AP Member Exchange was shared by the Fremont Tribune.</p>
						</block>
					</body.content>
				</nitf>
			</DataContent>
			<Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml" SizeInBytes="58">
				<Words>10</Words>
			</Characteristics>
		</TextContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Main" MediaType="Text">
		<TextContentItem Id="ad60ff12f12e42de94488a9b50d9a0ff" ArrivedInFilingId="c03e698999834086a1cabdeab7d5163f">
			<DataContent>
				<nitf>
					<body.content>
						<block>
							<p>LONDON (AP) — He talked nuclear threats with Russia’s president and gave an iPod to the queen.</p>
							<p>"I think if you pulled quotes from 10 years ago, 20 years ago, 30 years ago, from previous news reports, you might find similar contentions that America was on decline," Obama said. “And somehow it hasn’t worked out that way."</p>
						</block>
					</body.content>
				</nitf>
			</DataContent>
			<Characteristics MimeType="text/xml" Format="IIM" FileExtension="xml" SizeInBytes="110">
				<Words>1012</Words>
			</Characteristics>
		</TextContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Main" MediaType="Text">
		<TextContentItem Id="2b29780765ae46ebbf9d143427c0e037" ArrivedInFilingId="c03e698999834086a1cabdeab7d5163f">
			<DataContent>
				<anpa>¥´1151¤ ¥AP-EU-Obama,1466¤
					%headline(¥Obama seems everywhere as he takes global stage¤%)
					%xhl(Obama takes the world stage _ sober, joking, relaxed, seemingly everywhere on summit eve%)
					I think if you pulled quotes from 10 years ago, 20 years ago, 30 years ago, from previous news reports, you might find similar contentions that America was on decline,’’ Obama said. And somehow it hasn’t worked out that way.’’
				</anpa>
			</DataContent>
			<Characteristics MimeType="text/plain" Format="ANPA1312" FileExtension="txt" SizeInBytes="5960">
				<Words>1012</Words>
			</Characteristics>
		</TextContentItem>
	</PublicationComponent>
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.parsePubComponents(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Texts == nil {
		t.Error("[texts] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}
