package appl

import (
	"fmt"
	"testing"
)

func TestAdmin(t *testing.T) {
	s := `
<Publication>
	<PublicationManagement>		
		<RecordType>Change</RecordType>
		<FilingType>Text</FilingType>
		<FirstCreated UserAccount="APGBL" UserAccountSystem="APADS" UserName="APGBL\dcampbell" Year="2012" Month="3" Day="12" Time="20:54:44"/>
		<Status>Usable</Status>
		<ExplicitWarning>1</ExplicitWarning>
		<IsDigitized>false</IsDigitized>
	 </PublicationManagement>  
	<AdministrativeMetadata>
		<Provider Id="AP" Type="AP">AP</Provider>
		<Creator>AP</Creator>
		<Source Type="AP">AP</Source>
		<Contributor>NTB</Contributor>
		<SourceMaterial Name="alternate">
			<Type>Website</Type>
			<Url>http://www.apnewsarchive.com/2017/Ext-Sum-September-1st-test2/id-050b1f04ddfe47acb12c3937fe5ebe4b</Url>
		</SourceMaterial>
		<SourceMaterial Name="gb" Id="Glen Cooper Glen@fatmedia.co.uk">
			<Type>Email</Type>
			<PermissionGranted>Glen Cooper Senior Consultant Fat Media Ltd 12 Spring Garden Street Lancaster Lancashire LA1 1RQ 01524 590430</PermissionGranted>
		</SourceMaterial>
		<TransmissionSource>EapText</TransmissionSource>
		<TransmissionSource>NotMemberFeed</TransmissionSource>
		<TransmissionSource>ElvisLives</TransmissionSource>
		<TransmissionSource>JagRoars</TransmissionSource>
		<ProductSource>EuroPhotos</ProductSource>
		<ProductSource>GermanOnline</ProductSource>
		<ProductSource>AsiaPhotos</ProductSource>
		<ProductSource>GermanPhotos</ProductSource>
		<ItemContentType System="Editorial" Id="b8db75de62a043a8bc5649b226b606dd">Spot Development</ItemContentType>
		<ContentElement>FullStory</ContentElement>
		<DistributionChannel>Hosted Online</DistributionChannel>
		<DistributionChannel>Web File Delivery</DistributionChannel>
		<DistributionChannel>Web File Delivery</DistributionChannel>
		<Fixture Id="A60AC5A7AC024994B9102A73EFAB934A">Right Now</Fixture>
		<Reach Scheme="AP">HALO</Reach>
		<Reach Scheme="AP">ap_subject:General</Reach>
		<Reach Scheme="UNKNOWN">UNKNOWN</Reach>
		<Signal>YM</Signal>
		<InPackage Scheme="APText">tophd inthd alhd cthd dehd ilhd iahd arhd</InPackage>
		<InPackage Scheme="APText">asbizhd</InPackage>
		<InPackage Scheme="APText">BIZHD APFNHD</InPackage>
		<ConsumerReady>TRUE</ConsumerReady>
		<Rating Value="3" ScaleMin="1" ScaleMax="6" ScaleUnit="int" Raters="1" RaterType="Editorial" />
		<Rating Value="2" ScaleMin="1" ScaleMax="6" ScaleUnit="int" Raters="1" RaterType="Editorial" />
		<Property Name="EAI:SUBMISSIONPRIORITY"></Property>
		<Property Name="EAI:SLUGWORDCOUNT"></Property>
		<Property Name="EAI:ELVIS_CALLBACK_URL"></Property>
		<Property Name="EAI:ELVIS_WORKFLOW_ID"></Property>
	</AdministrativeMetadata>
</Publication>`
	pub, _ := NewXml(s)
	doc := document{Xml: pub}

	err := pub.PublicationManagement.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	err = pub.AdministrativeMetadata.parse(&doc)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.Provider == nil {
		t.Error("[provider] is expected")
	}

	if doc.Sources == nil {
		t.Error("[sources] is expected")
	}

	if doc.CanonicalLink == nil {
		t.Error("[canonicallink] is expected")
	}

	if doc.SourceMaterials == nil {
		t.Error("[sourcematerials] is expected")
	}

	if doc.TransmissionSources == nil {
		t.Error("[transmissionsources] is expected")
	}

	if doc.ProductSources == nil {
		t.Error("[productsources] is expected")
	}

	if doc.ItemContentType == nil {
		t.Error("[itemcontenttype] is expected")
	}

	if doc.DistributionChannels == nil {
		t.Error("[distributionchannels] is expected")
	}

	if doc.Fixture == nil {
		t.Error("[fixture] is expected")
	}

	if doc.Signals.IsEmpty() {
		t.Error("[signals] is expected")
	}

	if doc.InPackages == nil {
		t.Error("[inpackages] is expected")
	}

	jo, err := doc.ToJson()
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", jo.ToString())
}
