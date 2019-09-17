package xml

import (
	"fmt"
	"testing"
	"time"
)

func TestXml(t *testing.T) {
	s := `	
  <?xml version="1.0" encoding="utf-8" standalone = "no" ?>
  <!--comments must be ignored-->
	<Publication Version="5.3.0" xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <Identification>
    <ItemId>dd587867190648f3a8de18f25339fca8</ItemId>
    <RecordId>dd587867190648f3a8de18f25339fca8</RecordId>
    <CompositeId>00000000000000000000000000000000</CompositeId>
    <CompositionType>StandardPrintPhoto</CompositionType>
    <MediaType>Photo</MediaType>
    <Priority>5</Priority>
    <EditorialPriority>d</EditorialPriority>
    <DefaultLanguage>en-us</DefaultLanguage>
    <RecordSequenceNumber>0</RecordSequenceNumber>
    <FriendlyKey>19120542327253</FriendlyKey>
    <HeadLine>Vikings, Sage Rosenfels agree to 2-year contract</HeadLine>
  </Identification>
  <script>
   <![CDATA[
      <message> Welcome to TutorialsPoint </message>
   ]] >
</script >
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
