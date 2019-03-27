package appl

import "strings"

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
		aj.Contributor = admin.Contributor[0]
	}

	return nil
}

func getProvider(aj *ApplJson) {
	p := aj.Xml.AdministrativeMetadata.Provider
	empty := true
	provider := ApplProvider{}
	if p.Value != "" {
		provider.Name = p.Value
		empty = false
	}
	if p.Id != "" {
		provider.Code = p.Id
		empty = false
	}
	if p.Type != "" {
		provider.Type = p.Type
		empty = false
	}
	if p.SubType != "" {
		provider.Subtype = p.SubType
		empty = false
	}

	provider.IsEmpty = empty

	aj.Provider = provider
}

func getSources(aj *ApplJson) {
	srcs := aj.Xml.AdministrativeMetadata.Source
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sources := []ApplSource{}

	for _, src := range srcs {
		empty := true
		source := ApplSource{}
		if src.Value != "" {
			source.Name = src.Value
			empty = false
		}
		if src.Id != "" {
			source.Code = src.Id
			empty = false
		}
		if src.Type != "" {
			source.Type = src.Type
			empty = false
		}
		if src.SubType != "" {
			source.Subtype = src.SubType
			empty = false
		}
		if src.City != "" {
			source.City = src.City
			empty = false
		}
		if src.Country != "" {
			source.Country = src.Country
			empty = false
		}
		if src.Url != "" {
			source.Url = src.Url
			empty = false
		}

		source.IsEmpty = empty

		sources = append(sources, source)
	}

	aj.Sources = sources
}

func getSourceMaterials(aj *ApplJson) {
	srcs := aj.Xml.AdministrativeMetadata.SourceMaterial
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sources := []ApplSourceMaterial{}
	link := false

	for _, src := range srcs {
		name := src.Name
		if strings.EqualFold(name, "alternate") {
			if !link && src.Url != "" {
				aj.CanonicalLink = src.Url
				link = true
			}
		} else {
			empty := true
			source := ApplSourceMaterial{}
			if src.Name != "" {
				source.Name = src.Name
				empty = false
			}
			if src.Id != "" {
				source.Code = src.Id
				empty = false
			}
			if src.Type != "" {
				source.Type = src.Type
				empty = false
			}
			if src.PermissionGranted != "" {
				source.PermissionGranted = src.PermissionGranted
				empty = false
			}

			source.IsEmpty = empty

			sources = append(sources, source)
		}
	}

	if len(sources) > 0 {
		aj.SourceMaterials = sources
	}
}

func getTransmissionSources(aj *ApplJson) {
	tss := aj.Xml.AdministrativeMetadata.TransmissionSource

	if tss == nil || len(tss) == 0 {
		return
	}

	for _, ts := range tss {
		aj.TransmissionSources.Add(ts)
	}
}

func getProductSources(aj *ApplJson) {
	pss := aj.Xml.AdministrativeMetadata.ProductSource

	if pss == nil || len(pss) == 0 {
		return
	}

	for _, ps := range pss {
		aj.ProductSources.Add(ps)
	}
}

func getItemContentType(aj *ApplJson) {
	ict := aj.Xml.AdministrativeMetadata.ItemContentType
	empty := true
	itemcontenttype := ApplItemContentType{}
	if ict.Value != "" {
		itemcontenttype.Name = ict.Value
		empty = false
	}
	if ict.Id != "" {
		itemcontenttype.Code = ict.Id
		empty = false
	}
	if ict.System != "" {
		itemcontenttype.Creator = ict.System
		empty = false
	}

	itemcontenttype.IsEmpty = empty

	aj.ItemContentType = itemcontenttype
}

func getDistributionChannels(aj *ApplJson) {
	dcs := aj.Xml.AdministrativeMetadata.DistributionChannel

	if dcs == nil || len(dcs) == 0 {
		return
	}
	for _, dc := range dcs {
		aj.DistributionChannels.Add(dc)
	}
}

func getFixture(aj *ApplJson) {
	f := aj.Xml.AdministrativeMetadata.Fixture
	empty := true
	fixture := ApplFixture{}
	if f.Value != "" {
		fixture.Name = f.Value
		empty = false
	}
	if f.Id != "" {
		fixture.Code = f.Id
		empty = false
	}

	fixture.IsEmpty = empty

	aj.Fixture = fixture
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

	if ips == nil || len(ips) == 0 {
		return
	}
	for _, ip := range ips {
		tokens := strings.Split(ip, " ")
		for _, token := range tokens {
			aj.InPackages.Add(token)
		}
	}
}
