package appl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ymetelkin/go/xml"
)

func TestTextContentItem(t *testing.T) {
	s := `
<Publication>
  <PublicationComponent Role="Main" MediaType="Text">
    <TextContentItem Id="ad60ff12f12e42de94488a9b50d9a0ff" ArrivedInFilingId="c03e698999834086a1cabdeab7d5163f">
      <DataContent>
        <nitf>
          <body.content>
            <block>
              <p>LONDON (AP) — He talked nuclear threats with Russia’s president and gave an iPod to the queen.</p>
              <p>"I think if you pulled quotes from 10 years ago, 20 years ago, 30 years ago, from previous news reports, you might find similar contentions that America was on decline," Obama said. “And somehow it hasn’t worked out that way."</p>
              <p>
                Test href <a href="#">Test</a> should parse <a></a> as text
              </p>
            </block>
          </body.content>
        </nitf>
      </DataContent>
      <Characteristics MimeType="text/xml" Format="IIM" FileExtension="xml" SizeInBytes="110">
        <Words>666</Words>
      </Characteristics>
    </TextContentItem>
  </PublicationComponent>
  <PublicationComponent Role="Main" MediaType="Text">
    <TextContentItem Id="2b29780765ae46ebbf9d143427c0e037" ArrivedInFilingId="c03e698999834086a1cabdeab7d5163f">
      <DataContent>
        <anpa>
          ¥´1151¤ ¥AP-EU-Obama,1466¤
          %headline(¥Obama seems everywhere as he takes global stage¤%)
          %xhl(Obama takes the world stage _ sober, joking, relaxed, seemingly everywhere on summit eve%)
        </anpa>
      </DataContent>
      <Characteristics MimeType="text/plain" Format="ANPA1312" FileExtension="txt" SizeInBytes="5960">
        <Words>1012</Words>
      </Characteristics>
    </TextContentItem>
  </PublicationComponent>
  <PublicationComponent Role="Caption" MediaType="Text">
    <TextContentItem Id="a4235d49caeb43388058613008a3e3e3" ArrivedInFilingId="802bc28ebcfe402890dfa9e98a82a086">
      <DataContent>
        <nitf xmlns="http://ap.org/schemas/03/2005/appl">
          <body.content>
            <block>
              <p>This undated photo courtesy of Sarah Dorio shows a tattoo on the arm of chef Hugh Acheson. Acheson has four tattoos. Once considered the province of sailors, bikers, ex-cons and, of course, college hipsters, tattoos have become standard attire in professional kitchens.  (AP Photo/Sarah Dorio)</p>
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
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)

	for _, nd := range xml.Nodes {
		doc.parsePublicationComponent(nd, nil)
	}

	fmt.Println(doc.Story.Body)

	if !strings.HasPrefix(doc.Story.Body, "<p>LONDON (AP) — He talked nuclear threats") {
		t.Error("Invalid Story.Body")
	}
	if doc.Story.Words != 666 {
		t.Error("Invalid Story.Words")
	}
	if !strings.HasPrefix(doc.Caption.Body, "<p>This undated photo courtesy of Sarah Dorio") {
		t.Error("Invalid Caption.Body")
	}
	if doc.Caption.Words != 47 {
		t.Error("Invalid Caption.Words")
	}
	if !strings.HasPrefix(doc.Script.Body, "<p>Nathan Horton’s second goal") {
		t.Error("Invalid Script.Body")
	}
	if doc.Script.Words != 236 {
		t.Error("Invalid Script.Words")
	}
	if !strings.HasPrefix(doc.Shotlist.Body, "<p>TD Garden, Boston, Massachusetts") {
		t.Error("Invalid Shotlist.Body")
	}
	if doc.Shotlist.Words != 104 {
		t.Error("Invalid Shotlist.Words")
	}
	if doc.PublishableEditorNotes.Body != "<p>This AP Member Exchange was shared by the Fremont Tribune.</p>" {
		t.Error("Invalid PublishableEditorNotes.Body")
	}
	if doc.PublishableEditorNotes.Words != 10 {
		t.Error("Invalid PublishableEditorNotes.Words")
	}
}

func TestPhotoContentItem(t *testing.T) {
	s := `
<Publication>
  <PublicationComponent Role="Main" MediaType="Photo">
    <PhotoContentItem Id="8940d6f6e3c6493796751c4a84fae083" Href="\20080224\Main\77EE9F437837495B82FFECC14750D1B7.jpg" BinaryPath="File">
        <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="264471" Digest="180e91f5fd8494568cccd0bf4b71c7c0" OriginalFileName="CORRECTION TURKEY IRAQ.JPEG">
            <Width>1470</Width>
            <Height>1084</Height>
            <PhotoType>Horizontal</PhotoType>
        </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent Role="Thumbnail" MediaType="Photo">
    <PhotoContentItem Id="f5a4081c7b044a9facbbff946e7c5b69" Href="\20080224\Thumbnail\77EE9F437837495B82FFECC14750D1B7.jpg" BinaryPath="File">
        <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" Digest="61dd31519d2078735e9a5b301d9bdc55" OriginalFileName="CORRECTION TURKEY IRAQ.JPEG">
            <Width>128</Width>
            <Height>94</Height>
        </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent Role="Thumbnail" MediaType="Photo">
    <PhotoContentItem Id="1f42bab189c8cf0d2f0f6a706700515c" Href="\20130422\Thumbnail\4d72e25be8ac4200acfc49e458408f32_105x070.jpg" ArrivedInFilingId="2322633d39de489d809ec565d013457e" BinaryPath="File">
        <BinaryLocation To="2013-05-13T19:42:48" BinaryPath="Akamai" Sequence="0">jpg\2013\20130422\23\1f42bab189c8cf0d2f0f6a706700515c.jpg</BinaryLocation>
        <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://docs.core.prod.s3.amazonaws.com/eaacffa2cfe24e5bafdc13b29a2058e2/components/thumbnail02.jpg</BinaryLocation>
        <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="3234" OriginalFileName="0422apvus_boston_wrap_105x070.jpg">
            <Width>105</Width>
            <Height>70</Height>
        </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent Role="Thumbnail" MediaType="Photo">
      <PhotoContentItem Id="1f42faf489c8cf0d2f0f6a706700f5a3" Href="\20130422\Thumbnail\4d72e25be8ac4200acfc49e458408f32_180x135.jpg" ArrivedInFilingId="d1ffa4e116b144aebd748d028bded4ed" BinaryPath="File">
          <BinaryLocation To="2013-05-13T19:42:50" BinaryPath="Akamai" Sequence="0">jpg\2013\20130422\23\1f42faf489c8cf0d2f0f6a706700f5a3.jpg</BinaryLocation>
          <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://docs.core.prod.s3.amazonaws.com/eaacffa2cfe24e5bafdc13b29a2058e2/components/thumbnail03.jpg</BinaryLocation>
          <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="7039" OriginalFileName="0422apvus_boston_wrap_180x135.jpg">
              <Width>180</Width>
              <Height>135</Height>
          </Characteristics>
      </PhotoContentItem>
  </PublicationComponent>
  <PublicationComponent Role="Preview" MediaType="Photo">
    <PhotoContentItem Id="515507acc6be4177954f8dfb415aba35" Href="\20080224\Preview\77EE9F437837495B82FFECC14750D1B7.jpg" BinaryPath="File">
        <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" Digest="e87a50b721cbeb54d97bbe12b4c59ed5" OriginalFileName="CORRECTION TURKEY IRAQ.JPEG">
            <Width>512</Width>
            <Height>377</Height>
        </Characteristics>
    </PhotoContentItem>
  </PublicationComponent>
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)
	doc.MediaType = "photo"
	parser := &renditionParser{}
	for _, nd := range xml.Nodes {
		doc.parsePublicationComponent(nd, parser)
	}

	if len(doc.Renditions) != 3 {
		t.Error("Invalid Renditions")
	}

	doc = new(Document)
	doc.MediaType = "video"
	parser = &renditionParser{}
	for _, nd := range xml.Nodes {
		doc.parsePublicationComponent(nd, parser)
	}

	if len(doc.Renditions) != 4 {
		t.Error("Invalid Renditions")
	}
	if doc.Renditions[2].Title != "Thumbnail (JPG 180x135)" {
		t.Error("Renditions[3].Title")
	}
}

func TestVideoContentItem(t *testing.T) {
	s := `
<Publication>
  <PublicationComponent Role="PhysicalMain" MediaType="Video">
      <VideoContentItem Id="5b810c44e160455fb2c6950fdff1d376" Href="PATH_UNKNOWN" ArrivedInFilingId="5a5c571cefd14ae99a40de7a7cfa7899" BinaryPath="None">
          <Characteristics MimeType="video/quicktime" Format="Quicktime" FileExtension="mov">
              <AverageBitRate>25000000.000000</AverageBitRate>
              <TotalDuration>228000</TotalDuration>
              <Width>720</Width>
              <Height>576</Height>
              <VideoCoder>DV PAL</VideoCoder>
              <FrameRate>25.000000</FrameRate>
              <SampleRate>25000000.000000</SampleRate>
              <PhysicalType>Primary Datatape</PhysicalType>
          </Characteristics>
      </VideoContentItem>
  </PublicationComponent>
  <PublicationComponent Role="PhysicalMain" MediaType="Video">
      <VideoContentItem Id="1b4af4b154ef4deea977994442355044" Href="PATH_UNKNOWN" ArrivedInFilingId="5a5c571cefd14ae99a40de7a7cfa7899" BinaryPath="None">
          <Characteristics MimeType="video/quicktime" Format="Quicktime" FileExtension="mov">
              <AverageBitRate>25000000.000000</AverageBitRate>
              <TotalDuration>228000</TotalDuration>
              <Width>720</Width>
              <Height>576</Height>
              <VideoCoder>DV PAL</VideoCoder>
              <FrameRate>25.000000</FrameRate>
              <SampleRate>25000000.000000</SampleRate>
              <PhysicalType>Backup Datatape</PhysicalType>
          </Characteristics>
      </VideoContentItem>
  </PublicationComponent>
  <PublicationComponent Role="PhysicalMain" MediaType="Video">
      <VideoContentItem Id="e0dbb770807245aabe68113367b754ef" Href="EF02/0987" ArrivedInFilingId="5a5c571cefd14ae99a40de7a7cfa7899" BinaryPath="None">
          <ForeignKeys System="Tape">
              <Keys Id="EF02/0987" Field="Number" />
          </ForeignKeys>
          <Characteristics MimeType="Video" Format="Flash Video">
              <TotalDuration>228000</TotalDuration>
              <ManualInTimeCode>07:21:20:00</ManualInTimeCode>
              <PhysicalType>Legacy Videotape</PhysicalType>
          </Characteristics>
      </VideoContentItem>
  </PublicationComponent>
  <PublicationComponent Role="Main" MediaType="Video">
    <VideoContentItem Id="2df2bc22e61c48e28f96225fbc79f7aa" ArrivedInFilingId="3ae59dda139640dfa9b2a3cf5cfaf75b" BinaryPath="None">
      <Characteristics MimeType="text/plain" Format="IIM" FileExtension="txt" SizeInBytes="0" OriginalFileName="dummy.txt"/>
  </VideoContentItem>
  </PublicationComponent>
<PublicationComponent Role="Main" MediaType="Video">
  <VideoContentItem Id="dbef488fd4016f13150f776d76004d1d" Href="\20120715\Main\cb804fb5a166498f8a48b8e044a52d2b_x080b.wmv" ArrivedInFilingId="12e4418fdf76416983b427546dc76417" BinaryPath="File">
    <ExpireDateTime>2012-08-14T15:54:30</ExpireDateTime>
    <BinaryLocation To="2012-08-05T11:54:30" BinaryPath="Akamai" Sequence="0">wmv\2012\20120715\15\dbef488fd4016f13150f776d76004d1d.wmv</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://docs.core.prod.s3.amazonaws.com/c01188dea35745b8be1a874dc9adc804/components/secured/x080b.wmv</BinaryLocation>
    <Characteristics MimeType="video/x-ms-wmv" Format="Windows Media" FileExtension="wmv" SizeInBytes="17843088" OriginalFileName="0715dv_aurora_borealis_x080b.wmv">
        <AverageBitRate>1500.000000</AverageBitRate>
        <Width>856</Width>
        <Height>480</Height>
        <VideoCoder>Microsoft Windows Media Video</VideoCoder>
        <SampleRate>44000.000000</SampleRate>
        <ProducedAspectRatio>original</ProducedAspectRatio>
    </Characteristics>
  </VideoContentItem>
</PublicationComponent>
<PublicationComponent Role="Main" MediaType="Video">
  <VideoContentItem Id="d6eecc5cd4036f13150f776d7600343a" Href="\20120715\Main\cb804fb5a166498f8a48b8e044a52d2b_x070c.wmv" ArrivedInFilingId="4d2a172c370741cda1ed1fb97dd9caf2" BinaryPath="File">
      <ExpireDateTime>2012-08-14T15:56:40</ExpireDateTime>
      <BinaryLocation To="2012-08-05T11:56:40" BinaryPath="Akamai" Sequence="0">wmv\2012\20120715\15\d6eecc5cd4036f13150f776d7600343a.wmv</BinaryLocation>
      <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://docs.core.prod.s3.amazonaws.com/c01188dea35745b8be1a874dc9adc804/components/secured/x070c.wmv</BinaryLocation>
      <Characteristics MimeType="video/x-ms-wmv" Format="Windows Media" FileExtension="wmv" SizeInBytes="16251248" OriginalFileName="0715dv_aurora_borealis_x070c.wmv">
          <AverageBitRate>850.000000</AverageBitRate>
          <Width>856</Width>
          <Height>480</Height>
          <VideoCoder>Microsoft Windows Media Video</VideoCoder>
          <SampleRate>44000.000000</SampleRate>
          <ProducedAspectRatio>original</ProducedAspectRatio>
      </Characteristics>
    </VideoContentItem>
  </PublicationComponent>
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)
	doc.MediaType = "video"
	parser := &renditionParser{}
	for _, nd := range xml.Nodes {
		doc.parsePublicationComponent(nd, parser)
	}

	if len(doc.Renditions) != 5 {
		t.Error("Invalid Renditions")
	}
}
