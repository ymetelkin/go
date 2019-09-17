package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func TestIdentification(t *testing.T) {
	input := `
	<Publication>
		<Identification>
			<ItemId>c8f59c9cb3724b04b753ef957f6d885b</ItemId>
			<RecordId>24b1f7fcf6b54d20a96a1e86a9cfd5dd</RecordId>
			<CompositeId>00000000000000000000000000000000</CompositeId>
			<CompositionType>StandardPrintPhoto</CompositionType>
			<MediaType>Photo</MediaType>
			<Priority>5</Priority>
			<EditorialPriority>d</EditorialPriority>
			<DefaultLanguage>en-us</DefaultLanguage>
			<RecordSequenceNumber>2</RecordSequenceNumber>
			<FriendlyKey>18212677756771</FriendlyKey>
		</Identification>
	</Publication>`

	expected := `
	{
		"itemid": "c8f59c9cb3724b04b753ef957f6d885b",
		"recordid": "24b1f7fcf6b54d20a96a1e86a9cfd5dd",
		"compositeid": "00000000000000000000000000000000",
		"compositiontype": "StandardPrintPhoto",
		"type": "photo",
		"priority": 5,
		"editorialpriority": "d",
		"language": "en",
		"recordsequencenumber": 2,
		"friendlykey": "18212677756771",
		"referenceid": "c8f59c9cb3724b04b753ef957f6d885b"
	  }`

	xml, err := xml.ParseString(input)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)
	doc.XML = &xml
	doc.JSON = new(json.Object)

	doc.parseIdentification(xml.Node("Identification"))
	//fmt.Println(doc.JSON.String())

	test, _ := json.ParseObjectString(expected)
	left := doc.JSON.InlineString()
	right := test.InlineString()
	if left != right {
		t.Error("Failed Identification")
		fmt.Println(left)
		fmt.Println(right)
	}
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
	x, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc := new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "xyz" {
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
	x, err = xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc = new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "00000000000000000000000000000001" {
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

	x, err = xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc = new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "xyz" {
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
	x, err = xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc = new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "00000000000000000000000000000001" {
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
	x, err = xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc = new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "xyz" {
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
	x, err = xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc = new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "00000000000000000000000000000001" {
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
	x, err = xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc = new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "xyz" {
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
	x, err = xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}
	doc = new(Document)
	doc.XML = &x
	doc.JSON = new(json.Object)

	doc.parseIdentification(x.Node("Identification"))
	doc.setReferenceID()

	if doc.ReferenceID != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	}
}
