package appl

import (
	"strings"

	"github.com/ymetelkin/go/json"
)

func (admin *AdministrativeMetadata) parse(doc *document) error {
	getProvider(doc)
	getSources(doc)
	getSourceMaterials(doc)
	getTransmissionSources(doc)
	getProductSources(doc)
	getItemContentType(doc)
	getDistributionChannels(doc)
	getFixture(doc)
	getAdminSignals(doc)
	getInPackages(doc)

	if admin.Contributor != nil && len(admin.Contributor) > 0 {
		doc.Contributor = json.NewStringProperty("contributor", admin.Contributor[0])
	}

	return nil
}

func getProvider(doc *document) {
	p := doc.Xml.AdministrativeMetadata.Provider

	provider := json.Object{}
	if p.Id != "" {
		provider.AddString("code", p.Id)
	}
	if p.Type != "" {
		provider.AddString("type", p.Type)
	}
	if p.SubType != "" {
		provider.AddString("subtype", p.SubType)
	}
	if p.Value != "" {
		provider.AddString("name", p.Value)
	}

	if !provider.IsEmpty() {
		doc.Provider = json.NewObjectProperty("provider", &provider)
	}
}

func getSources(doc *document) {
	srcs := doc.Xml.AdministrativeMetadata.Source
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sources := json.Array{}

	for _, src := range srcs {
		source := json.Object{}

		if src.City != "" {
			source.AddString("city", src.City)
		}
		if src.Country != "" {
			source.AddString("country", src.Country)
		}
		if src.Id != "" {
			source.AddString("code", src.Id)
		}
		if src.Url != "" {
			source.AddString("url", src.Url)
		}
		if src.Type != "" {
			source.AddString("type", src.Type)
		}
		if src.SubType != "" {
			source.AddString("subtype", src.SubType)
		}
		if src.Value != "" {
			source.AddString("name", src.Value)
		}

		sources.AddObject(&source)
	}

	doc.Sources = json.NewArrayProperty("sources", &sources)
}

func getSourceMaterials(doc *document) {
	srcs := doc.Xml.AdministrativeMetadata.SourceMaterial
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sourcematerials := json.Array{}
	for _, src := range srcs {
		name := src.Name
		if strings.EqualFold(name, "alternate") {
			if doc.CanonicalLink == nil && src.Url != "" {
				doc.CanonicalLink = json.NewStringProperty("canonicallink", src.Url)
			}
		} else {
			sourcematerial := json.Object{}
			if name != "" {
				sourcematerial.AddString("name", name)
			}
			if src.Id != "" {
				sourcematerial.AddString("code", src.Id)
			}
			if src.Type != "" {
				sourcematerial.AddString("type", src.Type)
			}
			if src.PermissionGranted != "" {
				sourcematerial.AddString("permissiongranted", src.PermissionGranted)
			}

			sourcematerials.AddObject(&sourcematerial)
		}
	}

	if sourcematerials.Length() > 0 {
		doc.SourceMaterials = json.NewArrayProperty("sourcematerials", &sourcematerials)
	}
}

func getTransmissionSources(doc *document) {
	tss := doc.Xml.AdministrativeMetadata.TransmissionSource
	if tss != nil {
		transmissionsources := uniqueArray{}
		for _, ts := range tss {
			transmissionsources.AddString(ts)
		}
		doc.TransmissionSources = transmissionsources.ToJsonProperty("transmissionsources")
	}
}

func getProductSources(doc *document) {
	pss := doc.Xml.AdministrativeMetadata.ProductSource
	if pss != nil {
		productsources := uniqueArray{}
		for _, ps := range pss {
			productsources.AddString(ps)
		}
		doc.ProductSources = productsources.ToJsonProperty("productsources")
	}
}

func getItemContentType(doc *document) {
	ict := doc.Xml.AdministrativeMetadata.ItemContentType
	itemcontenttype := json.Object{}
	if ict.System != "" {
		itemcontenttype.AddString("creator", ict.System)
	}
	if ict.Id != "" {
		itemcontenttype.AddString("code", ict.Id)
	}
	if ict.Value != "" {
		itemcontenttype.AddString("name", ict.Value)
	}
	doc.ItemContentType = json.NewObjectProperty("itemcontenttype", &itemcontenttype)
}

func getDistributionChannels(doc *document) {
	dcs := doc.Xml.AdministrativeMetadata.DistributionChannel
	if dcs != nil {
		distributionchannels := uniqueArray{}
		for _, dc := range dcs {
			distributionchannels.AddString(dc)
		}
		doc.DistributionChannels = distributionchannels.ToJsonProperty("distributionchannels")
	}
}

func getFixture(doc *document) {
	f := doc.Xml.AdministrativeMetadata.Fixture
	fixture := json.Object{}
	if f.Id != "" {
		fixture.AddString("code", f.Id)
	}
	if f.Value != "" {
		fixture.AddString("name", f.Value)
	}

	if !fixture.IsEmpty() {
		doc.Fixture = json.NewObjectProperty("fixture", &fixture)
	}
}

func getAdminSignals(doc *document) {
	admin := doc.Xml.AdministrativeMetadata

	if admin.Reach != nil {
		for _, reach := range admin.Reach {
			if !strings.EqualFold(reach, "UNKNOWN") {
				doc.Signals.AddString(reach)
			}
		}
	}

	if strings.EqualFold(admin.ConsumerReady, "TRUE") {
		doc.Signals.AddString("newscontent")
	}

	if admin.Signal != nil {
		for _, signal := range admin.Signal {
			doc.Signals.AddString(signal)
		}
	}
}

func getInPackages(doc *document) {
	ips := doc.Xml.AdministrativeMetadata.InPackage
	if ips != nil {
		inpackages := uniqueArray{}

		for _, ip := range ips {
			tokens := strings.Split(ip, " ")
			for _, token := range tokens {
				inpackages.AddString(token)
			}
		}

		doc.InPackages = inpackages.ToJsonProperty("inpackages")
	}
}

func getRatings(doc *document) {
	rs := doc.Xml.AdministrativeMetadata.Rating
	if rs != nil {
		ratings := json.Array{}

		for _, r := range rs {
			if r.Value > 0 && r.ScaleMin > 0 && r.ScaleMax > 0 && r.ScaleUnit != "" {
				rating := json.Object{}
				rating.AddInt("rating", r.Value)
				rating.AddInt("scalemin", r.ScaleMin)
				rating.AddInt("scalemax", r.ScaleMax)
				rating.AddString("scaleunit", r.ScaleUnit)
				if r.Raters > 0 {
					rating.AddInt("raters", r.Raters)
				}
				if r.RaterType != "" {
					rating.AddString("ratertype", r.RaterType)
				}
				ratings.AddObject(&rating)
			}
		}

		if ratings.Length() > 0 {
			doc.Ratings = json.NewArrayProperty("ratings", &ratings)
		}
	}
}
