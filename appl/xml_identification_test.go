package appl

import (
	"fmt"
	"testing"
)

func TestIdentification(t *testing.T) {
	s := `
<Publication 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
   	Version="5.0.0.9"
	   xmlns="http://ap.org/schemas/03/2005/appl">
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
	  <PublicationManagement>
		<RecordType>Change</RecordType>
		<FilingType>Text</FilingType>
		<FirstCreated UserAccount="APGBL" UserAccountSystem="APADS" UserName="APGBL\dcampbell"
			Year="2012" Month="3" Day="12" Time="20:54:44"/>
		<Status>Usable</Status>
	  </PublicationManagement>
</Publication>
`
	jo, err := XmlToJson(s)
	if err != nil {
		t.Error(err.Error())
	} else if s, _ := jo.GetString("itemid"); s != "00000000000000000000000000000001" {
		t.Error("Mismatch: [itemid]")
	} else if s, _ := jo.GetString("type"); s != "photo" {
		t.Error("Mismatch: [type]")
	} else if s, _ := jo.GetString("language"); s != "en" {
		t.Error("Mismatch: [language]")
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}
}

func TestIdentificationValidation(t *testing.T) {
	s := `
<Publication 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
   	Version="5.0.0.9"
	   xmlns="http://ap.org/schemas/03/2005/appl">
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
	<Publication 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
   	Version="5.0.0.9"
	   xmlns="http://ap.org/schemas/03/2005/appl">
	<Identification>
		<ItemId>00000000000000000000000000000001</ItemId>
		<RecordId>00000000000000000000000000000002</RecordId>
		<CompositeId>00000000000000000000000000000003</CompositeId>
		<CompositionType>StandardPrintPhoto</CompositionType>
		<MediaType>Picture</MediaType>
		<Priority>5</Priority>
		<EditorialPriority>d</EditorialPriority>
		<DefaultLanguage>en-us</DefaultLanguage>
		<RecordSequenceNumber>2</RecordSequenceNumber>
		<FriendlyKey>18212677756771</FriendlyKey>
  	</Identification> 
</Publication>
`

	_, err = XmlToJson(s)
	if err == nil {
		t.Error("Must throw")
	} else {
		fmt.Printf("%s\n", err.Error())
	}
}
