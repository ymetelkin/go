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
			doc.Contributor = nd.Text
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
			doc.ConsumerReady = nd.Text
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
		for _, a := range nd.Attributes {
			switch a.Name {
			case "Id":
				provider.Code = a.Value
			case "Type":
				provider.Type = a.Value
			case "SubType":
				provider.Subtype = a.Value
			}
		}
	}
	provider.Name = nd.Text
	doc.Provider = &provider
}

func (doc *Document) parseSource(nd xml.Node) {
	var source Source

	if nd.Attributes != nil {
		for _, a := range nd.Attributes {
			switch a.Name {
			case "Id":
				source.Code = a.Value
			case "City":
				source.City = a.Value
			case "CountryArea":
				source.CountryArea = a.Value
			case "Country":
				source.Country = a.Value
			case "County":
				source.County = a.Value
			case "Url":
				source.URL = a.Value
			case "Type":
				source.Type = a.Value
			case "SubType":
				source.Subtype = a.Value
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
		for _, a := range nd.Attributes {
			switch a.Name {
			case "Id":
				sm.Code = a.Value
			case "Name":
				sm.Name = a.Value
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
		for _, a := range nd.Attributes {
			switch a.Name {
			case "Id":
				ict.Code = a.Value
			case "System":
				ict.Creator = a.Value
			}
		}
	}
	doc.ItemContentType = &ict
}

func (doc *Document) parseRating(nd xml.Node) {
	if nd.Attributes != nil {
		var r Rating

		for _, a := range nd.Attributes {
			switch a.Name {
			case "Value":
				r.Value, _ = strconv.Atoi(a.Value)
			case "ScaleMin":
				r.ScaleMin, _ = strconv.Atoi(a.Value)
			case "ScaleMax":
				r.ScaleMax, _ = strconv.Atoi(a.Value)
			case "ScaleUnit":
				r.ScaleUnit = a.Value
			case "Raters":
				r.Raters, _ = strconv.Atoi(a.Value)
			case "RaterType":
				r.RaterType = a.Value
			case "Creator":
				r.Creator = a.Value
			}
		}

		doc.Ratings = append(doc.Ratings, r)
	}
}

func (doc *Document) parseNewsContent() {
	if doc.ConsumerReady == "" || strings.EqualFold(doc.ConsumerReady, "UNKNOWN") {
		if doc.Filings != nil {
			f := doc.Filings[0]
			if f.Category == "V" || f.Category == "v" {
				return
			}
			if f.Category == "r" && strings.HasPrefix(f.Selector, "apr") {
				return
			}
			if f.Category == "t" && strings.HasPrefix(f.Selector, "1tv") {
				return
			}
		}
		if doc.SuppCategories != nil {
			supp := doc.SuppCategories[0]
			if supp.Code == "V" || supp.Code == "v" {
				return
			}
		}
		if doc.ItemEndDateTime != nil && doc.ItemContentType.Name == "Advisory" {
			return
		}
		if doc.EditorialTypes != nil {
			for _, et := range doc.EditorialTypes {
				if et == "Advisory" || et == "Disregard" || et == "Elimination" || et == "Withhold" {
					return
				}
				if (et == "Correction" || et == "Add") && doc.MediaType != "text" {
					return
				}
			}
		}
		if doc.Signals != nil {
			for _, signal := range doc.Signals {
				if strings.EqualFold(signal, "TEST") {
					return
				}
			}
		}
	} else if !strings.EqualFold(doc.ConsumerReady, "TRUE") {
		return
	}

	doc.Signals = append(doc.Signals, "newscontent")
}
