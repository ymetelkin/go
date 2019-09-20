package appl

import (
	"testing"

	"github.com/ymetelkin/go/xml"
)

func TestAdmin(t *testing.T) {
	s := `
<Publication>
	<AdministrativeMetadata>
		<Provider Id="AP" Type="AP">AP</Provider>
		<Creator>AP</Creator>
		<Source City="Buenos Aires" Country="Argentina" Id="z00052" Url="http://www.dyn.com" Type="ThirdParty" SubType="NewsAgency">Agencia Diarios y Noticias (DYN)</Source>
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
		<InPackage Scheme="APText">asbizhd</InPackage>
		<InPackage Scheme="APText">BIZHD APFNHD</InPackage>
		<ConsumerReady>TRUE</ConsumerReady>
		<Rating Value="3" ScaleMin="1" ScaleMax="6" ScaleUnit="int" Raters="1" RaterType="Editorial" />
		<Rating Value="2" ScaleMin="1" ScaleMax="6" ScaleUnit="int" Raters="1" RaterType="Editorial" />		
	</AdministrativeMetadata>
</Publication>`

	xml, err := xml.ParseString(s)
	if err != nil {
		t.Error(err.Error())
	}

	doc := new(Document)

	doc.parseAdministrativeMetadata(xml.Node("AdministrativeMetadata"))

	if doc.Provider.Code != "AP" {
		t.Error("Invalid Provider")
	}
	if doc.Creator != "AP" {
		t.Error("Invalid Creator")
	}
	if doc.Sources[0].Country != "Argentina" {
		t.Error("Invalid .Sources[0].Country")
	}
	if doc.CanonicalLink != "http://www.apnewsarchive.com/2017/Ext-Sum-September-1st-test2/id-050b1f04ddfe47acb12c3937fe5ebe4b" {
		t.Error("Invalid CanonicalLink")
	}
	if doc.SourceMaterials[0].Type != "Email" {
		t.Error("Invalid SourceMaterials[1].Type")
	}
	if doc.TransmissionSources[3] != "JagRoars" {
		t.Error("Invalid TransmissionSources[3]")
	}
	if doc.ProductSources[3] != "GermanPhotos" {
		t.Error("Invalid ProductSources[3]")
	}
	if doc.DistributionChannels[1] != "Web File Delivery" {
		t.Error("Invalid DistributionChannels[1]")
	}
	if len(doc.DistributionChannels) != 2 {
		t.Error("Invalid length of DistributionChannels")
	}
	if doc.InPackages[2] != "APFNHD" {
		t.Error("Invalid InPackages[2]")
	}
	if doc.Ratings[1].Value != 2 {
		t.Error("Invalid Ratings[1].Value")
	}
	if doc.Signals[2] != "YM" {
		t.Error("Invalid Signals[2]")
	}
	if len(doc.Signals) != 4 {
		t.Error("Invalid length of Signals")
	}
}
