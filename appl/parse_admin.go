package appl

import (
	"strings"

	"github.com/ymetelkin/go/json"
)

func (admin *AdministrativeMetadata) parse(aj *ApplJson) error {
	getProvider(aj)
	getSources(aj)
	getSourceMaterials(aj)
	getTransmissionSources(aj)
	getProductSources(aj)
	getItemContentType(aj)
	getDistributionChannels(aj)
	getFixture(aj)
	getAdminSignals(aj)
	getInPackages(aj)

	if admin.Contributor != nil && len(admin.Contributor) > 0 {
		aj.Contributor = json.NewStringProperty("contributor", admin.Contributor[0])
	}

	return nil
}

func getProvider(aj *ApplJson) {
	p := aj.Xml.AdministrativeMetadata.Provider

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
		aj.Provider = json.NewObjectProperty("provider", &provider)
	}
}

func getSources(aj *ApplJson) {
	srcs := aj.Xml.AdministrativeMetadata.Source
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

	aj.Sources = json.NewArrayProperty("sources", &sources)
}

func getSourceMaterials(aj *ApplJson) {
	srcs := aj.Xml.AdministrativeMetadata.SourceMaterial
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sourcematerials := json.Array{}
	for _, src := range srcs {
		name := src.Name
		if strings.EqualFold(name, "alternate") {
			if aj.CanonicalLink == nil && src.Url != "" {
				aj.CanonicalLink = json.NewStringProperty("canonicallink", src.Url)
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
		aj.SourceMaterials = json.NewArrayProperty("sourcematerials", &sourcematerials)
	}
}

func getTransmissionSources(aj *ApplJson) {
	tss := aj.Xml.AdministrativeMetadata.TransmissionSource
	if tss != nil {
		transmissionsources := UniqueStrings{}
		for _, ts := range tss {
			transmissionsources.Add(ts)
		}
		aj.TransmissionSources = transmissionsources.ToJsonProperty("transmissionsources")
	}
}

func getProductSources(aj *ApplJson) {
	pss := aj.Xml.AdministrativeMetadata.ProductSource
	if pss != nil {
		productsources := UniqueStrings{}
		for _, ps := range pss {
			productsources.Add(ps)
		}
		aj.ProductSources = productsources.ToJsonProperty("productsources")
	}
}

func getItemContentType(aj *ApplJson) {
	ict := aj.Xml.AdministrativeMetadata.ItemContentType
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
	aj.ItemContentType = json.NewObjectProperty("itemcontenttype", &itemcontenttype)
}

func getDistributionChannels(aj *ApplJson) {
	dcs := aj.Xml.AdministrativeMetadata.DistributionChannel
	if dcs != nil {
		distributionchannels := UniqueStrings{}
		for _, dc := range dcs {
			distributionchannels.Add(dc)
		}
		aj.DistributionChannels = distributionchannels.ToJsonProperty("distributionchannels")
	}
}

func getFixture(aj *ApplJson) {
	f := aj.Xml.AdministrativeMetadata.Fixture
	fixture := json.Object{}
	if f.Id != "" {
		fixture.AddString("code", f.Id)
	}
	if f.Value != "" {
		fixture.AddString("name", f.Value)
	}

	if !fixture.IsEmpty() {
		aj.Fixture = json.NewObjectProperty("fixture", &fixture)
	}
}

func getAdminSignals(aj *ApplJson) {
	admin := aj.Xml.AdministrativeMetadata

	if admin.Reach != nil {
		for _, reach := range admin.Reach {
			if !strings.EqualFold(reach, "UNKNOWN") {
				aj.Signals.Add(reach)
			}
		}
	}

	if strings.EqualFold(admin.ConsumerReady, "TRUE") {
		aj.Signals.Add("newscontent")
	}

	if admin.Signal != nil {
		for _, signal := range admin.Signal {
			aj.Signals.Add(signal)
		}
	}
}

func getInPackages(aj *ApplJson) {
	ips := aj.Xml.AdministrativeMetadata.InPackage
	if ips != nil {
		inpackages := UniqueStrings{}

		for _, ip := range ips {
			tokens := strings.Split(ip, " ")
			for _, token := range tokens {
				inpackages.Add(token)
			}
		}

		aj.InPackages = inpackages.ToJsonProperty("inpackages")
	}
}

func getRatings(aj *ApplJson) {
	rs := aj.Xml.AdministrativeMetadata.Rating
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
			aj.Ratings = json.NewArrayProperty("ratings", &ratings)
		}
	}
}
