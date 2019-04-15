package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestPhotos(t *testing.T) {
	s := `
<Publication>
	<PublicationComponent Role="Thumbnail" MediaType="Photo">
		<PhotoContentItem Id="1d89d6edbc381f28e20f6a706700e8d8" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Thumbnail\2018\20180703\23\1d89d6edbc381f28e20f6a706700e8d8.jpg" ArrivedInFilingId="ed047806d5264a1cbc89840bfd43b07e" BinaryPath="None">
			<BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
			<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/thumbnail.jpg</BinaryLocation>
			<Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="7769" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a.jpg">
				<Width>105</Width>
				<Height>70</Height>
			</Characteristics>
		</PhotoContentItem>
	</PublicationComponent>
</Publication>`
	doc, _ := parseXML(s)
	jo := json.Object{}

	doc.ParsePublicationComponents(&jo)

	if _, err := jo.GetArray("renditions"); err != nil {
		t.Error("[renditions] is expected")
	}

	fmt.Printf("%s\n", jo.ToString())
}
