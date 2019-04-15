package appl

import (
	"fmt"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestVideos(t *testing.T) {
	s := `
<Publication>
	<PublicationComponent Role="Main" MediaType="Video">
		<VideoContentItem Id="48bb100d2ddf44cab7fa9d51b375dbdb" ArrivedInFilingId="a96b5cb4ed384ac2a3a8601cef932185" BinaryPath="None">
			<Characteristics MimeType="text/plain" Format="IIM" FileExtension="txt" SizeInBytes="0" OriginalFileName="dummy.txt">
				<TotalDuration>81000</TotalDuration>
				<ProducedAspectRatio>original</ProducedAspectRatio>
			</Characteristics>
		</VideoContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Main" MediaType="Video">
		<VideoContentItem Id="c201299bbc3a1f28e20f6d7034003820" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Main\2018\20180703\23\c201299bbc3a1f28e20f6d7034003820.mp4" ArrivedInFilingId="7bc3736a63994f7a9cd6a1ec4b2b9586" BinaryPath="None">
			<BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
			<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/secured/x090i.mp4</BinaryLocation>
			<Characteristics MimeType="video/mpeg" Format="MPEG" FileExtension="mp4" SizeInBytes="105681605" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_x090i.mp4">
				<AverageBitRate>10240.000000</AverageBitRate>
				<TotalDuration>81000</TotalDuration>
				<Width>1920</Width>
				<Height>1080</Height>
				<VideoCoder>H.264</VideoCoder>
				<FrameRate>60.000000</FrameRate>
				<SampleRate>44000.000000</SampleRate>
				<ProducedAspectRatio>original</ProducedAspectRatio>
			</Characteristics>
		</VideoContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Main" MediaType="Video">
		<VideoContentItem Id="8844a9b1bc3b1f28e20f6d70340066f7" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Main\2018\20180703\23\8844a9b1bc3b1f28e20f6d70340066f7.mp4" ArrivedInFilingId="7f2d66a5f9894b829acc5682295ff225" BinaryPath="None">
			<BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
			<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/secured/x070n.mp4</BinaryLocation>
			<Characteristics MimeType="video/mpeg" Format="MPEG" FileExtension="mp4" SizeInBytes="5783002" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_x070n.mp4">
				<AverageBitRate>448.000000</AverageBitRate>
				<TotalDuration>81000</TotalDuration>
				<Width>576</Width>
				<Height>324</Height>
				<VideoCoder>H.264</VideoCoder>
				<FrameRate>30.000000</FrameRate>
				<SampleRate>44000.000000</SampleRate>
				<ProducedAspectRatio>original</ProducedAspectRatio>
			</Characteristics>
		</VideoContentItem>
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
