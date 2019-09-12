package appl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestIdentification(t *testing.T) {
	s := `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Photo</MediaType>
		<Priority>5</Priority>
		<EditorialPriority>d</EditorialPriority>
		<DefaultLanguage>en-us</DefaultLanguage>
		<RecordSequenceNumber>2</RecordSequenceNumber>
		<FriendlyKey>18212677756771</FriendlyKey>
	</Identification>
</Publication>`
	doc, _ := parseXML(strings.NewReader(s))
	jo := json.Object{}

	err := doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}

	if string(doc.MediaType) != "photo" {
		t.Error("[type:photo] is expected")
	}
	if v, _ := jo.GetString("itemid"); v != "00000000000000000000000000000001" {
		t.Error("[itemid:00000000000000000000000000000001] is expected")
	}
	if v, _ := jo.GetString("language"); v != "en" {
		t.Error("[language:en] is expected")
	}

	fmt.Printf("%s\n", jo.String())
}

func TestIdentificationReferenceId(t *testing.T) {
	s := `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Photo</MediaType>
		<FriendlyKey>xyz</FriendlyKey>
	</Identification>
</Publication>`
	doc, _ := parseXML(strings.NewReader(s))
	jo := json.Object{}

	err := doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	}

	s = `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Photo</MediaType>
	</Identification>
</Publication>`
	doc, _ = parseXML(strings.NewReader(s))
	jo = json.Object{}

	err = doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	}

	s = `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>ComplexData</MediaType>
	</Identification>
	<NewsLines>
		<Title>xyz</Title>
	</NewsLines> 
</Publication>`
	doc, _ = parseXML(strings.NewReader(s))
	jo = json.Object{}

	err = doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	err = doc.ParseNewsLines(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	}

	s = `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>ComplexData</MediaType>
	</Identification>
</Publication>`
	doc, _ = parseXML(strings.NewReader(s))
	jo = json.Object{}

	err = doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	}

	s = `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Text</MediaType>
	</Identification>
	<NewsLines>
		<Title>xyz</Title>
	</NewsLines> 
</Publication>`
	doc, _ = parseXML(strings.NewReader(s))
	jo = json.Object{}

	err = doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	err = doc.ParseNewsLines(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	}

	s = `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Text</MediaType>
	</Identification>
</Publication>`
	doc, _ = parseXML(strings.NewReader(s))
	jo = json.Object{}

	err = doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	}

	s = `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardBroadcastVideo</CompositionType>
		<MediaType>Video</MediaType>
	</Identification>
	<PublicationManagement>
		<EditorialId>xyz</EditorialId>
	</PublicationManagement> 
</Publication>`
	doc, _ = parseXML(strings.NewReader(s))
	jo = json.Object{}

	err = doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	err = doc.ParsePublicationManagement(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	}

	s = `
<Publication>
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardBroadcastVideo</CompositionType>
		<MediaType>Video</MediaType>
	</Identification>
</Publication>`
	doc, _ = parseXML(strings.NewReader(s))
	jo = json.Object{}

	err = doc.ParseIdentification(&jo)
	if err != nil {
		t.Error(err.Error())
	}
	doc.SetReferenceID(&jo)

	if v, _ := jo.GetString("referenceid"); v != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	}
}
