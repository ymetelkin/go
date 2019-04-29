package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test(t *testing.T) {
	s := `<Publication Version="5.3.1" xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<Identification>
	  <ItemId>2fface573aba95e21aceb60b258ee22b</ItemId>
	  <RecordId>ba2bdc5bf3484163891f561baab8d38e</RecordId>
	  <CompositeId>cb11fdecc2824311a6f746dcc360c82d</CompositeId>
	  <CompositionType>StandardLibraryVideo</CompositionType>
	  <MediaType>Video</MediaType>
	  <DefaultLanguage>en-us</DefaultLanguage>
	  <RecordSequenceNumber>0</RecordSequenceNumber>
	  <FriendlyKey>1343442690</FriendlyKey>
	</Identification>
	<PublicationManagement>
	  <RecordType>Version</RecordType>
	  <FilingType>OnlineVideo</FilingType>
	  <ItemKey>AP|Video|???</ItemKey>
	  <IsDistributionReady>false</IsDistributionReady>
	  <ArrivalDateTime>2019-04-16T19:14:05</ArrivalDateTime>
	  <FirstCreated Year="2019" Month="4" Day="16" Time="19:14:05"></FirstCreated>
	  <LastModifiedDateTime>2019-01-07T16:09:43</LastModifiedDateTime>
	  <Status>Usable</Status>
	  <ReleaseDateTime>2019-04-18T19:14:05</ReleaseDateTime>
	  <EditorialId>Qapus093798</EditorialId>
	  <ItemStartDateTime>2019-04-16T19:14:05</ItemStartDateTime>
	  <ItemStartDateTimeActual>2019-04-16T19:14:05</ItemStartDateTimeActual>
	  <ExplicitWarning>0</ExplicitWarning>
	  <Function>APTNLibrary</Function>
	  <IsDigitized>true</IsDigitized>
	</PublicationManagement>
	<NewsLines>
	  <Title>asdf</Title>
	  <DateLine>sdaf - sadf</DateLine>
	  <RightsLine>AP Clients Only</RightsLine>
	  <LocationLine>sdaf</LocationLine>
	</NewsLines>
	<AdministrativeMetadata>
	  <Provider Type="AP">AP</Provider>
	  <Source Type="AP">ASSOCIATED PRESS</Source>
	  <WorkflowStatus>MetadataPending</WorkflowStatus>
	  <TransmissionSource>APTNArchive</TransmissionSource>
	  <TransmissionSource>NotMemberFeed</TransmissionSource>
	  <TransmissionSource>JagRoars</TransmissionSource>
	  <Workgroup>APTN-Archive</Workgroup>
	  <DistributionChannel>Archive</DistributionChannel>
	  <ConsumerReady>TRUE</ConsumerReady>
	  <Signal>singlesource</Signal>
	  <Signal>whitelisted</Signal>
	  <Property Name="EAI:SUBMISSIONPRIORITY"></Property>
	</AdministrativeMetadata>
	<RightsMetadata>
	  <Copyright Holder="AP" Date="2019"></Copyright>
	</RightsMetadata>
	<DescriptiveMetadata>
	  <DateLineLocation>
		<LatitudeDD>0.000000000</LatitudeDD>
		<LongitudeDD>0.000000000</LongitudeDD>
	  </DateLineLocation>
	  <SalesClassification System="Editorial" Authority="AP Sales Code">
		<Occurrence Id="EBF652509112100480DDA55C96277D3E" Value="Basic"></Occurrence>
	  </SalesClassification>
	</DescriptiveMetadata>
	<FilingMetadata>
	  <Id>2fface573aba95e21aceb60b258ee22b</Id>
	  <ArrivalDateTime>2019-04-16T19:14:05</ArrivalDateTime>
	  <TransmissionContent>All</TransmissionContent>
	  <SlugLine>Qapus093798-CR-VO-apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5</SlugLine>
	  <Products>
		<Product>46170</Product>
		<Product>46157</Product>
		<Product>46140</Product>
		<Product>45956</Product>
		<Product>45579</Product>
		<Product>45486</Product>
		<Product>45476</Product>
		<Product>44461</Product>
		<Product>43848</Product>
		<Product>35535</Product>
		<Product>8976</Product>
		<Product>1180</Product>
		<Product>6</Product>
		<Product>1</Product>
	  </Products>
	  <ForeignKeys System="Story">
		<Keys Id="Qapus093798" Field="ID"></Keys>
	  </ForeignKeys>
	  <ForeignKeys System="MOS">
		<Keys Id="vAP.Neuron.DigView" Field="mosID"></Keys>
		<Keys Id="Qapus093798-apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5" Field="objID"></Keys>
	  </ForeignKeys>
	  <ForeignKeys System="mosPayload">
		<Keys Id="VIDEO" Field="objType"></Keys>
		<Keys Id="17000" Field="duration"></Keys>
		<Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
		<Keys Id="AP" Field="Provider"></Keys>
		<Keys Id="EapVideo" Field="TransmissionSource"></Keys>
		<Keys Id="news-us" Field="ProductSource"></Keys>
		<Keys Id="2fface573aba95e21aceb60b258ee22b" Field="ItemId"></Keys>
		<Keys Id="ba2bdc5bf3484163891f561baab8d38e" Field="RecordId"></Keys>
		<Keys Id="0" Field="RecordSequenceNumber"></Keys>
		<Keys Id="2019-04-16T19:14:05" Field="DateCreated"></Keys>
		<Keys Id="2019-01-07T16:09:43" Field="LastModifiedDateTime"></Keys>
		<Keys Id="1" Field="Approved"></Keys>
		<Keys Id="JGAINES" Field="Creator"></Keys>
		<Keys Id="0" Field="MediaTime"></Keys>
		<Keys Id="RREGAN" Field="ModBy"></Keys>
		<Keys Id="20190107160943" Field="ModTime"></Keys>
		<Keys Id="20" Field="MOSItemDurations"></Keys>
		<Keys Id="apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5" Field="MOSObjSlugs"></Keys>
		<Keys Id="New Row 25 Rob Regan 3-3" Field="MOSSlugs"></Keys>
		<Keys Id="Authoritative answer: Host not found" Field="MOSStatus"></Keys>
		<Keys Id="JGAINES" Field="Owner"></Keys>
		<Keys Id="0" Field="pubApproved"></Keys>
		<Keys Id="AP Clients Only" Field="Restrictions"></Keys>
		<Keys Id="0" Field="SourceMediaTime"></Keys>
		<Keys Id="JGAINES" Field="SourceModBy"></Keys>
		<Keys Id="20150109220247" Field="SourceModTime"></Keys>
		<Keys Id="0" Field="SourceTextTime"></Keys>
		<Keys Id="6" Field="TextTime"></Keys>
		<Keys Id="0" Field="LegalBlock"></Keys>
		<Keys Id=" RREGAN " Field="SystemApprovedBy"></Keys>
		<Keys Id="ap" Field="apslatebackgroundimage"></Keys>
		<Keys Id="Qapus093798" Field="APStoryNumber"></Keys>
		<Keys Id="16:9" Field="AspectRatio"></Keys>
		<Keys Id="English/Natsound" Field="AudioDescription"></Keys>
		<Keys Id="sdaf - sadf" Field="Dateline"></Keys>
		<Keys Id="AP Clients Only" Field="DomesticRestrictions"></Keys>
		<Keys Id="asdf" Field="Headline"></Keys>
		<Keys Id="AP Clients Only" Field="InternationalRestrictions"></Keys>
		<Keys Id="2fface573aba95e21aceb60b258ee22b" Field="ItemIdArchive"></Keys>
		<Keys Id="9b658757b45044cb9fe4e60bec0da18c" Field="ItemIdOnline"></Keys>
		<Keys Id="0" Field="Kill"></Keys>
		<Keys Id="English/Natsound" Field="LanguageType"></Keys>
		<Keys Id="sdaf" Field="Location"></Keys>
		<Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
		<Keys Id="PUBLISHED-EDITS/Latest News/apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5" Field="originalFileName"></Keys>
		<Keys Id="original" Field="ProducedAspectRatio"></Keys>
		<Keys Id="AP Video US" Field="ProductsUS"></Keys>
		<Keys Id="sadf" Field="Date"></Keys>
		<Keys Id="ASSOCIATED PRESS" Field="Source"></Keys>
		<Keys Id="CR-VO" Field="StoryFormat"></Keys>
		<Keys Id="Qapus093798" Field="StoryNumber"></Keys>
		<Keys Id="apus" Field="StoryNumberPrefix"></Keys>
		<Keys Id="asdf" Field="StorySummary"></Keys>
		<Keys Id="news-us" Field="StoryWorkgroup"></Keys>
		<Keys Id="News" Field="VHCategoryCode"></Keys>
		<Keys Id="3" Field="ENPSItemType"></Keys>
		<Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
	  </ForeignKeys>
	</FilingMetadata>
	<FilingMetadata>
	  <Id>2fface573aba95e21aceb60b258ee22b</Id>
	  <ArrivalDateTime>2019-04-16T19:32:44</ArrivalDateTime>
	  <TransmissionContent>All</TransmissionContent>
	  <SlugLine>Qapus093798-CR-VO-apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5</SlugLine>
	  <Products>
		<Product>46170</Product>
		<Product>46157</Product>
		<Product>46140</Product>
		<Product>45956</Product>
		<Product>45579</Product>
		<Product>45486</Product>
		<Product>45476</Product>
		<Product>44461</Product>
		<Product>44316</Product>
		<Product>43848</Product>
		<Product>35535</Product>
		<Product>8976</Product>
		<Product>1180</Product>
		<Product>6</Product>
		<Product>1</Product>
	  </Products>
	  <ForeignKeys System="Story">
		<Keys Id="Qapus093798" Field="ID"></Keys>
	  </ForeignKeys>
	  <ForeignKeys System="MOS">
		<Keys Id="vAP.Neuron.DigView" Field="mosID"></Keys>
		<Keys Id="Qapus093798-apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5" Field="objID"></Keys>
	  </ForeignKeys>
	  <ForeignKeys System="mosPayload">
		<Keys Id="VIDEO" Field="objType"></Keys>
		<Keys Id="17000" Field="duration"></Keys>
		<Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
		<Keys Id="AP" Field="Provider"></Keys>
		<Keys Id="EapVideo" Field="TransmissionSource"></Keys>
		<Keys Id="news-us" Field="ProductSource"></Keys>
		<Keys Id="2fface573aba95e21aceb60b258ee22b" Field="ItemId"></Keys>
		<Keys Id="ba2bdc5bf3484163891f561baab8d38e" Field="RecordId"></Keys>
		<Keys Id="0" Field="RecordSequenceNumber"></Keys>
		<Keys Id="2019-04-16T19:14:05" Field="DateCreated"></Keys>
		<Keys Id="2019-01-07T16:09:43" Field="LastModifiedDateTime"></Keys>
		<Keys Id="1" Field="Approved"></Keys>
		<Keys Id="JGAINES" Field="Creator"></Keys>
		<Keys Id="0" Field="MediaTime"></Keys>
		<Keys Id="RREGAN" Field="ModBy"></Keys>
		<Keys Id="20190107160943" Field="ModTime"></Keys>
		<Keys Id="20" Field="MOSItemDurations"></Keys>
		<Keys Id="apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5" Field="MOSObjSlugs"></Keys>
		<Keys Id="New Row 25 Rob Regan 3-3" Field="MOSSlugs"></Keys>
		<Keys Id="Authoritative answer: Host not found" Field="MOSStatus"></Keys>
		<Keys Id="JGAINES" Field="Owner"></Keys>
		<Keys Id="0" Field="pubApproved"></Keys>
		<Keys Id="AP Clients Only" Field="Restrictions"></Keys>
		<Keys Id="0" Field="SourceMediaTime"></Keys>
		<Keys Id="JGAINES" Field="SourceModBy"></Keys>
		<Keys Id="20150109220247" Field="SourceModTime"></Keys>
		<Keys Id="0" Field="SourceTextTime"></Keys>
		<Keys Id="6" Field="TextTime"></Keys>
		<Keys Id="0" Field="LegalBlock"></Keys>
		<Keys Id=" RREGAN " Field="SystemApprovedBy"></Keys>
		<Keys Id="ap" Field="apslatebackgroundimage"></Keys>
		<Keys Id="Qapus093798" Field="APStoryNumber"></Keys>
		<Keys Id="16:9" Field="AspectRatio"></Keys>
		<Keys Id="English/Natsound" Field="AudioDescription"></Keys>
		<Keys Id="sdaf - sadf" Field="Dateline"></Keys>
		<Keys Id="AP Clients Only" Field="DomesticRestrictions"></Keys>
		<Keys Id="asdf" Field="Headline"></Keys>
		<Keys Id="AP Clients Only" Field="InternationalRestrictions"></Keys>
		<Keys Id="2fface573aba95e21aceb60b258ee22b" Field="ItemIdArchive"></Keys>
		<Keys Id="9b658757b45044cb9fe4e60bec0da18c" Field="ItemIdOnline"></Keys>
		<Keys Id="0" Field="Kill"></Keys>
		<Keys Id="English/Natsound" Field="LanguageType"></Keys>
		<Keys Id="sdaf" Field="Location"></Keys>
		<Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
		<Keys Id="PUBLISHED-EDITS/Latest News/apus115118_af9e23d1bf2d73ec3e1275d8cd29afa5" Field="originalFileName"></Keys>
		<Keys Id="original" Field="ProducedAspectRatio"></Keys>
		<Keys Id="AP Video US" Field="ProductsUS"></Keys>
		<Keys Id="sadf" Field="Date"></Keys>
		<Keys Id="ASSOCIATED PRESS" Field="Source"></Keys>
		<Keys Id="CR-VO" Field="StoryFormat"></Keys>
		<Keys Id="Qapus093798" Field="StoryNumber"></Keys>
		<Keys Id="apus" Field="StoryNumberPrefix"></Keys>
		<Keys Id="asdf" Field="StorySummary"></Keys>
		<Keys Id="news-us" Field="StoryWorkgroup"></Keys>
		<Keys Id="News" Field="VHCategoryCode"></Keys>
		<Keys Id="3" Field="ENPSItemType"></Keys>
		<Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
		<Keys Id="false" Field="APWorkflowTranscode"></Keys>
	  </ForeignKeys>
	</FilingMetadata>
	<PublicationComponent Role="Script" MediaType="Text">
	  <TextContentItem Id="a7a088213e5e44cda1c0422f22ae68ce" ArrivedInFilingId="2fface573aba95e21aceb60b258ee22b">
		<DataContent>
		  <nitf>
			<body.content>
			  <block>
				<p>CAPTION GOES HERE (DATE) </p>
				<p>RESTRICTION SUMMARY: </p>
				<p>CLIENTS NOTE: VIDEO ONLY - SHOTLIST AND STORYLINE TO FOLLOW AS SOON AS POSSIBLE</p>
				<p>SHOTLIST:</p>
				<p>1.</p>
				<p>STORYLINE: </p>
			  </block>
			</body.content>
		  </nitf>
		</DataContent>
		<Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml" SizeInBytes="147">
		  <Words>22</Words>
		</Characteristics>
	  </TextContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Thumbnail" MediaType="Photo">
	  <PhotoCollectionContentItem Id="bac32d3b09094ec3b079e03b4fc1ca9b" ArrivedInFilingId="2fface573aba95e21aceb60b258ee22b" BinaryPath="VideoImages" BaseFileName="http://ap.core.videos.qa.s3.amazonaws.com/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_THUMBNAIL_index/0000.jpg" PrimaryFileName="http://ap.core.videos.qa.s3.amazonaws.com/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_THUMBNAIL_index/0000.jpg">
		<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="VideoArchive" Sequence="1">http://ap.core.videos.qa.s3.amazonaws.com/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_THUMBNAIL_index/0000.jpg</BinaryLocation>
		<File TimeOffSetMilliseconds="1"></File>
		<File TimeOffSetMilliseconds="6200"></File>
		<File TimeOffSetMilliseconds="16200"></File>
		<File TimeOffSetMilliseconds="26200"></File>
		<File TimeOffSetMilliseconds="36200"></File>
		<File TimeOffSetMilliseconds="41280"></File>
		<File TimeOffSetMilliseconds="51280"></File>
		<File TimeOffSetMilliseconds="61280"></File>
		<File TimeOffSetMilliseconds="63960"></File>
		<File TimeOffSetMilliseconds="73960"></File>
		<File TimeOffSetMilliseconds="83960"></File>
		<File TimeOffSetMilliseconds="93960"></File>
		<File TimeOffSetMilliseconds="97920"></File>
		<File TimeOffSetMilliseconds="107920"></File>
		<File TimeOffSetMilliseconds="117920"></File>
		<File TimeOffSetMilliseconds="120200"></File>
		<File TimeOffSetMilliseconds="130200"></File>
		<File TimeOffSetMilliseconds="135760"></File>
		<Characteristics MimeType="a" Format="CSV" Encoding="zip" FileExtension="jpg">
		  <Width>192</Width>
		  <Height>108</Height>
		</Characteristics>
	  </PhotoCollectionContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Main" MediaType="Video">
	  <VideoContentItem Id="35bbcb5fbf964cf9971c038536e981ef" Href="http://ap.core.videos.qa.s3.amazonaws.com/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_1080i50-h264-4m-screener.mp4" ArrivedInFilingId="2fface573aba95e21aceb60b258ee22b" BinaryPath="VideoArchive">
		<Characteristics MimeType="video/mp4" Format="MPEG" FileExtension="mp4">
		  <AverageBitRate>1500000</AverageBitRate>
		  <TotalDuration>140680</TotalDuration>
		  <Width>1920</Width>
		  <Height>1080</Height>
		  <VideoCoder>H.264</VideoCoder>
		  <FrameRate>25.000000</FrameRate>
		  <AspectRatio>16:9</AspectRatio>
		  <SampleRate>1500000</SampleRate>
		  <OriginalMediaId>2fface573aba95e21aceb60b258ee22b</OriginalMediaId>
		  <LogicalAspectRatio>16:9</LogicalAspectRatio>
		  <ProducedAspectRatio>original</ProducedAspectRatio>
		</Characteristics>
	  </VideoContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Main" MediaType="Video">
	  <VideoContentItem Id="f547e218e65e4135977a71de943a5e26" Href="http://ap.core.videos.qa.s3.amazonaws.com/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_1080i60-h264-4m-screener.mp4" ArrivedInFilingId="2fface573aba95e21aceb60b258ee22b" BinaryPath="VideoArchive">
		<Characteristics MimeType="video/mp4" Format="MPEG" FileExtension="mp4">
		  <AverageBitRate>1500000</AverageBitRate>
		  <TotalDuration>140680</TotalDuration>
		  <Width>1920</Width>
		  <Height>1080</Height>
		  <VideoCoder>H.264</VideoCoder>
		  <FrameRate>29.970000</FrameRate>
		  <AspectRatio>16:9</AspectRatio>
		  <SampleRate>1500000</SampleRate>
		  <OriginalMediaId>2fface573aba95e21aceb60b258ee22b</OriginalMediaId>
		  <LogicalAspectRatio>16:9</LogicalAspectRatio>
		  <ProducedAspectRatio>original</ProducedAspectRatio>
		</Characteristics>
	  </VideoContentItem>
	</PublicationComponent>
	<PublicationComponent Role="Preview" MediaType="Video">
	  <VideoContentItem Id="10e1189b60014ce38591a60a4250aecc" Href="http://ap.core.videos.qa/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_360p25-h264-1500k.mp4" ArrivedInFilingId="2fface573aba95e21aceb60b258ee22b" BinaryPath="VideoArchive">
		<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="VideoArchive" Sequence="1">http://ap.core.videos.qa/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_360p25-h264-1500k.mp4</BinaryLocation>
		<Characteristics MimeType="video/mpeg" Format="MPEG" FileExtension="mp4">
		  <AverageBitRate>1500000</AverageBitRate>
		  <TotalDuration>140680</TotalDuration>
		  <Width>576</Width>
		  <Height>324</Height>
		  <VideoCoder>H.264</VideoCoder>
		  <FrameRate>15.000000</FrameRate>
		  <AspectRatio>16:9</AspectRatio>
		  <SampleRate>1500000</SampleRate>
		  <LogicalAspectRatio>16:9</LogicalAspectRatio>
		  <ProducedAspectRatio>original</ProducedAspectRatio>
		</Characteristics>
	  </VideoContentItem>
	</PublicationComponent>
	<PublicationComponent Role="PhysicalMain" MediaType="Video">
	  <VideoContentItem Id="724f04fea1544abc83416dbbcd76cb21" Href="http://ap.core.videos.qa.s3.amazonaws.com/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_1080i50-xdcam-50m.mxf" ArrivedInFilingId="2fface573aba95e21aceb60b258ee22b" BinaryPath="VideoArchive">
		<Property Name="OriginalVideo:True"></Property>
		<Property Name="BroadcastFormat:HD-XDCAM422-1080i50"></Property>
		<BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://qa-aparchivebackups.s3.amazonaws.com/original/2fface573aba95e21aceb60b258ee22b/2fface573aba95e21aceb60b258ee22b_1080i50ESSENCE.mxf</BinaryLocation>
		<Characteristics MimeType="application/mxf" Format="MPEG" FileExtension="mxf">
		  <AverageBitRate>50000000</AverageBitRate>
		  <TotalDuration>140680</TotalDuration>
		  <Width>1920</Width>
		  <Height>1080</Height>
		  <VideoCoder>MPEG</VideoCoder>
		  <FrameRate>25.000000</FrameRate>
		  <AspectRatio>16:9</AspectRatio>
		  <SampleRate>50000000</SampleRate>
		  <PhysicalType>Legacy Videotape</PhysicalType>
		  <LogicalAspectRatio>16:9</LogicalAspectRatio>
		  <ProducedAspectRatio>original</ProducedAspectRatio>
		</Characteristics>
	  </VideoContentItem>
	</PublicationComponent>
  </Publication>`

	req := events.APIGatewayProxyRequest{Body: s}
	res, err := execute(req)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(res.Body)
}
