package xml

import (
	"fmt"
	"testing"
	"time"
)

func TestXml(t *testing.T) {
	s := `	
  <Publication xmlns="http://ap.org/schemas/03/2005/appl" Version="5.3.2">
  <YM name="Back & White" />
  <PublicationComponent MediaType="Text" Role="Main">
    <TextContentItem ArrivedInFilingId="8e7c84199d734f7183ff0b47198f4304" Id="ebafd746085148928e851f6c67411978">
      <DataContent>
        <nitf>
          <body.content>
            <block>
              <hl2>Google wins case over reach of EU 'right to be forgotten'</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="EU--Europe-Google" name="MetaLabel" />
                <media-reference data-location="EU--Europe-Google" name="MetaKey" />
                <media-reference data-location="4e2be78d6ad34d34a87ab15cc1806eb5" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>US home prices rise at slowest pace in 7 years</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="US--Home Prices" name="MetaLabel" />
                <media-reference data-location="US--Home Prices" name="MetaKey" />
                <media-reference data-location="071f2594ba2c40cbaf820b7441b442d6" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>Volkswagen bosses charged in Germany over diesel scandal</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="EU--Germany-Volkswagen" name="MetaLabel" />
                <media-reference data-location="EU--Germany-Volkswagen" name="MetaKey" />
                <media-reference data-location="48fe9cf4337a4cc3a6b3366e2f47bc70" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>Feds: Engineer manipulated diesel emissions at Fiat Chrysler</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="APFN-US--Fiat Chrysler-Diesel" name="MetaLabel" />
                <media-reference data-location="APFN-US--Fiat Chrysler-Diesel" name="MetaKey" />
                <media-reference data-location="bd21ecaf14b14e149132928757d961aa" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>Weak consumer confidence pulls stocks lower</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="Financial Markets" name="MetaLabel" />
                <media-reference data-location="Financial Markets" name="MetaKey" />
                <media-reference data-location="138544dc9c4a4438910af152cc875fb5" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>US consumer confidence drops as economic uncertainties rise</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="APFN-US--Consumer Confidence" name="MetaLabel" />
                <media-reference data-location="APFN-US--Consumer Confidence" name="MetaKey" />
                <media-reference data-location="da1176f1645648568faf413db348af38" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>Top UK court: Johnson's suspension of Parliament was illegal</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="EU--Brexit-Supreme Court" name="MetaLabel" />
                <media-reference data-location="EU--Brexit-Supreme Court" name="MetaKey" />
                <media-reference data-location="6c86bdb652ce4066996817aec04cd51a" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>Investigator says FAA training inspectors weren't qualified</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="US--Boeing Plane-Training" name="MetaLabel" />
                <media-reference data-location="US--Boeing Plane-Training" name="MetaKey" />
                <media-reference data-location="df95532694ed484fbfba7bcf090902ce" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>Fox apologizes for 'disgraceful' comment about Thunberg</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="APFN-US--Media-Fox-Thunberg" name="MetaLabel" />
                <media-reference data-location="APFN-US--Media-Fox-Thunberg" name="MetaKey" />
                <media-reference data-location="b7aaa301943846d08d72eafd00d54b15" name="NormalizedLink" />
              </media>
            </block>
            <block>
              <hl2>Thomas Cook repatriation ramps up as bosses slammed over pay</hl2>
              <media media-type="text">
                <media-metadata name="LinkType" value="Item" />
                <media-metadata name="CompositionType" value="StandardText" />
                <media-reference data-location="EU--Britain-Thomas Cook" name="MetaLabel" />
                <media-reference data-location="EU--Britain-Thomas Cook" name="MetaKey" />
                <media-reference data-location="9418da2924e54c03bfefdb1506d71281" name="NormalizedLink" />
              </media>
            </block>
          </body.content>
        </nitf>
      </DataContent>
      <Characteristics FileExtension="xml" Format="IIM" MimeType="text/xml" SizeInBytes="1329">
        <Words>168</Words>
      </Characteristics>
    </TextContentItem>
  </PublicationComponent>
  <PublicationComponent MediaType="Text" Role="Main">
    <TextContentItem ArrivedInFilingId="8e7c84199d734f7183ff0b47198f4304" Id="644a1d9b3192469499d63c0304c6a778">
      <DataContent>
        <anpa>¥´9004¤ ¥AP-Top-Business-Headlines¤,0168¤
%headline(¥AP Top Business News at 1:38 p.m. EDT¤¤%)
%strytype(ContentType:HeadlinePackage; ContentElement:Other;%)
%endtag(%)
¥Google wins case over reach of EU 'right to be forgotten'
¥US home prices rise at slowest pace in 7 years
¥Volkswagen bosses charged in Germany over diesel scandal
¥Feds: Engineer manipulated diesel emissions at Fiat Chrysler
¥Weak consumer confidence pulls stocks lower
¥US consumer confidence drops as economic uncertainties rise
¥Top UK court: Johnson's suspension of Parliament was illegal
¥Investigator says FAA training inspectors weren't qualified
¥Fox apologizes for 'disgraceful' comment about Thunberg
¥Thomas Cook repatriation ramps up as bosses slammed over pay
</anpa>
      </DataContent>
      <Characteristics FileExtension="txt" Format="ANPA1312" MimeType="text/plain" SizeInBytes="744">
        <Words>101</Words>
      </Characteristics>
    </TextContentItem>
  </PublicationComponent>
</Publication>`

	//s = `<Identification><HeadLine>Vikings, Sage Rosenfels agree to 2-year contract</HeadLine></Identification>`

	start := time.Now()
	nd, err := ParseString(s)
	ts := time.Since(start)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", nd.String())
	fmt.Printf("Duration: %v\n", ts)
}
