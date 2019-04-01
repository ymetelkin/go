package appl

import (
	"fmt"
	"testing"
)

func TestIdentification(t *testing.T) {
	s := `
<Publication>
	<Identification>
  	</Identification> 
</Publication>
`
	_, err := XmlToJson(s)
	if err == nil {
		t.Error("Must throw")
	} else {
		fmt.Printf("%s\n", err.Error())
	}

	s = `
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
	pub, _ := NewXml(s)
	aj := ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}

	if string(aj.MediaType) != "photo" {
		t.Error("[type:photo] is expected")
	}
	if aj.Language.Field == "" {
		t.Error("[language:en] is expected")
	}
	if aj.ReferenceId.Field == "" {
		t.Error("[referenceid] is expected")
	}

	jo, err := aj.ToJson()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s\n", jo.ToString())
}

func TestIdentificationValidation(t *testing.T) {
	s := `
<Publication>
	<Identification>
  	</Identification> 
</Publication>
`
	_, err := XmlToJson(s)
	if err == nil {
		t.Error("Must throw")
	} else {
		fmt.Printf("%s\n", err.Error())
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
	pub, _ := NewXml(s)
	aj := ApplJson{Xml: pub}

	err := pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}

	if s, _ := aj.ReferenceId.Value.GetString(); s != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
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
	pub, _ = NewXml(s)
	aj = ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}
	if s, _ := aj.ReferenceId.Value.GetString(); s != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
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
	pub, _ = NewXml(s)
	aj = ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}
	if s, _ := aj.ReferenceId.Value.GetString(); s != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
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
	pub, _ = NewXml(s)
	aj = ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}
	if s, _ := aj.ReferenceId.Value.GetString(); s != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
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
	pub, _ = NewXml(s)
	aj = ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}
	if s, _ := aj.ReferenceId.Value.GetString(); s != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
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
	pub, _ = NewXml(s)
	aj = ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}
	if s, _ := aj.ReferenceId.Value.GetString(); s != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
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
	pub, _ = NewXml(s)
	aj = ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}
	if s, _ := aj.ReferenceId.Value.GetString(); s != "xyz" {
		t.Error("[referenceid:xyz] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
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
	pub, _ = NewXml(s)
	aj = ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		t.Error(err.Error())
	}
	if s, _ := aj.ReferenceId.Value.GetString(); s != "00000000000000000000000000000001" {
		t.Error("[referenceid:00000000000000000000000000000001] is expected")
	} else {
		fmt.Println(aj.ReferenceId)
	}
}
