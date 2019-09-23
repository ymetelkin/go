package appl

import (
	"strconv"
	"strings"

	"github.com/ymetelkin/go/xml"
)

func (doc *Document) parseAdministrativeMetadata(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	var tss, pss, dcs, ins uniqueStrings

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "Provider":
			doc.parseProvider(nd)
		case "Creator":
			doc.Creator = nd.Text
		case "Source":
			doc.parseSource(nd)
		case "Contributor":
			if doc.Contributor == "" {
				doc.Contributor = nd.Text
			}
		case "SourceMaterial":
			doc.parseSourceMaterial(nd)
		case "WorkflowStatus":
			doc.WorkflowStatus = nd.Text
		case "TransmissionSource":
			tss.Append(nd.Text)
		case "ProductSource":
			pss.Append(nd.Text)
		case "ItemContentType":
			doc.parseItemContentType(nd)
		case "Workgroup":
			doc.Workgroup = nd.Text
		case "DistributionChannel":
			dcs.Append(nd.Text)
		case "ContentElement":
			doc.EditorialRole = nd.Text
		case "Fixture":
			doc.Fixture = &CodeName{
				Code: nd.Attribute("Id"),
				Name: nd.Text,
			}
		case "Reach":
			if nd.Text != "" && !strings.EqualFold(nd.Text, "UNKNOWN") {
				doc.Signals = append(doc.Signals, nd.Text)
			}
		case "InPackage":
			if nd.Text != "" {
				toks := strings.Split(nd.Text, " ")
				for _, tok := range toks {
					ins.Append(tok)
				}
			}
		case "ConsumerReady":
			if nd.Text != "" && strings.EqualFold(nd.Text, "TRUE") {
				doc.Signals = append(doc.Signals, "newscontent")
			}
		case "Signal":
			if nd.Text != "" {
				doc.Signals = append(doc.Signals, nd.Text)
			}
		case "Rating":
			doc.parseRating(nd)
		}
	}

	if !tss.IsEmpty() {
		doc.TransmissionSources = tss.Values()
	}

	if !pss.IsEmpty() {
		doc.ProductSources = pss.Values()
	}

	if !dcs.IsEmpty() {
		doc.DistributionChannels = dcs.Values()
	}

	if !ins.IsEmpty() {
		doc.InPackages = ins.Values()
	}
}

func (doc *Document) parseProvider(nd xml.Node) {
	var provider Provider

	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "Id":
				provider.Code = v
			case "Type":
				provider.Type = v
			case "SubType":
				provider.Subtype = v
			}
		}
	}
	provider.Name = nd.Text
	doc.Provider = &provider
}

func (doc *Document) parseSource(nd xml.Node) {
	var source Source

	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "Id":
				source.Code = v
			case "City":
				source.City = v
			case "CountryArea":
				source.CountryArea = v
			case "Country":
				source.Country = v
			case "County":
				source.County = v
			case "Url":
				source.URL = v
			case "Type":
				source.Type = v
			case "SubType":
				source.Subtype = v
			}
		}
	}

	source.Name = nd.Text
	doc.Sources = append(doc.Sources, source)
}

func (doc *Document) parseSourceMaterial(nd xml.Node) {
	var (
		sm  SourceMaterial
		url string
	)

	if nd.Nodes != nil {
		for _, n := range nd.Nodes {
			switch n.Name {
			case "Type":
				sm.Type = n.Text
			case "Url":
				url = n.Text
			case "PermissionGranted":
				sm.PermissionGranted = n.Text
			}
		}
	}

	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "Id":
				sm.Code = v
			case "Name":
				sm.Name = v
			}
		}
	}

	if strings.EqualFold(sm.Name, "alternate") && url != "" && doc.CanonicalLink == "" {
		doc.CanonicalLink = url
		return
	}
	doc.SourceMaterials = append(doc.SourceMaterials, sm)
}

func (doc *Document) parseItemContentType(nd xml.Node) {
	ict := ItemContentType{
		Name: nd.Text,
	}
	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "Id":
				ict.Code = v
			case "System":
				ict.Creator = v
			}
		}
	}
	doc.ItemContentType = &ict
}

func (doc *Document) parseRating(nd xml.Node) {
	if nd.Attributes != nil {
		var r Rating

		for k, v := range nd.Attributes {
			switch k {
			case "Value":
				r.Value, _ = strconv.Atoi(v)
			case "ScaleMin":
				r.ScaleMin, _ = strconv.Atoi(v)
			case "ScaleMax":
				r.ScaleMax, _ = strconv.Atoi(v)
			case "ScaleUnit":
				r.ScaleUnit = v
			case "Raters":
				r.Value, _ = strconv.Atoi(v)
			case "RaterType":
				r.RaterType = v
			case "Creator":
				r.Creator = v
			}
		}

		doc.Ratings = append(doc.Ratings, r)
	}
}
