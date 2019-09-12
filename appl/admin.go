package appl

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func (doc *document) ParseAdministrativeMetadata(jo *json.Object) error {
	if doc.AdministrativeMetadata == nil || doc.AdministrativeMetadata.Nodes == nil || len(doc.AdministrativeMetadata.Nodes) == 0 {
		return errors.New("AdministrativeMetadata is missing")
	}

	var (
		s1, s2             bool
		tss, pss, dcs, ins uniqueArray
		srcs, sms, rts     []xml.Node
		ict, fx            json.Object
		cntr               string
	)

	s1 = !doc.Signals.IsEmpty()

	for _, nd := range doc.AdministrativeMetadata.Nodes {
		switch nd.Name {
		case "Provider":
			getProvider(nd, jo)
		case "Creator":
			if nd.Text != "" {
				jo.AddString("creator", nd.Text)
			}
		case "Source":
			if srcs == nil {
				srcs = []xml.Node{nd}
			} else {
				srcs = append(srcs, nd)
			}
		case "SourceMaterial":
			if sms == nil {
				sms = []xml.Node{nd}
			} else {
				sms = append(sms, nd)
			}
		case "WorkflowStatus":
			if nd.Text != "" {
				jo.AddString("workflowstatus", nd.Text)
			}
		case "TransmissionSource":
			tss.AddString(nd.Text)
		case "ProductSource":
			pss.AddString(nd.Text)
		case "DistributionChannel":
			dcs.AddString(nd.Text)
		case "InPackage":
			if nd.Text != "" {
				toks := strings.Split(nd.Text, " ")
				for _, tok := range toks {
					ins.AddString(tok)
				}
			}
		case "ItemContentType":
			ict = json.Object{}
			if nd.Attributes != nil {
				for k, v := range nd.Attributes {
					switch k {
					case "Id":
						if v != "" {
							ict.AddString("code", v)
						}
					case "System":
						if v != "" {
							ict.AddString("creator", v)
						}
					}
				}
			}

			if nd.Text != "" {
				ict.AddString("name", nd.Text)
			}
		case "Workgroup":
			if nd.Text != "" {
				jo.AddString("workgroup", nd.Text)
			}
		case "ContentElement":
			if nd.Text != "" {
				jo.AddString("editorialrole", nd.Text)
			}
		case "Fixture":
			fx = json.Object{}
			id := nd.Attribute("Id")
			if id != "" {
				fx.AddString("code", id)
			}
			if nd.Text != "" {
				fx.AddString("name", nd.Text)
			}
		case "Rating":
			if rts == nil {
				rts = []xml.Node{nd}
			} else {
				rts = append(rts, nd)
			}
		case "Reach":
			if nd.Text != "" && !strings.EqualFold(nd.Text, "UNKNOWN") {
				doc.Signals.AddString(nd.Text)
				s2 = true
			}
		case "ConsumerReady":
			if nd.Text != "" && strings.EqualFold(nd.Text, "TRUE") {
				doc.Signals.AddString("newscontent")
				s2 = true
			}
		case "Signal":
			if nd.Text != "" {
				doc.Signals.AddString("newscontent")
				s2 = true
			}
		case "Contributor":
			if cntr == "" {
				cntr = nd.Text
			}
		}
	}

	getSources(srcs, jo)
	getSourceMaterials(sms, jo)

	jo.AddProperty(tss.ToJSONProperty("transmissionsources"))
	jo.AddProperty(pss.ToJSONProperty("productsources"))

	if !ict.IsEmpty() {
		jo.AddObject("itemcontenttype", ict)
	}

	jo.AddProperty(dcs.ToJSONProperty("distributionchannels"))

	if !fx.IsEmpty() {
		doc.Fixture = true
		jo.AddObject("fixture", fx)
	}

	jo.AddProperty(ins.ToJSONProperty("inpackages"))

	if cntr != "" {
		jo.AddString("contributor", cntr)
	}

	getRatings(rts, jo)

	if s2 {
		if s1 {
			jo.SetArray("signals", doc.Signals.ToJSONArray())
		} else {
			jo.AddProperty(doc.Signals.ToJSONProperty("signals"))
		}
	}

	return nil
}

func getProvider(nd xml.Node, jo *json.Object) {
	provider := json.Object{}

	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "Id":
				if v != "" {
					provider.AddString("code", v)
				}
			case "Type":
				if v != "" {
					provider.AddString("type", v)
				}
			case "SubType":
				if v != "" {
					provider.AddString("subtype", v)
				}
			}
		}
	}

	if nd.Text != "" {
		provider.AddString("name", nd.Text)
	}

	if !provider.IsEmpty() {
		jo.AddObject("provider", provider)
	}
}

func getSources(srcs []xml.Node, jo *json.Object) {
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sources := json.Array{}

	for _, src := range srcs {
		source := json.Object{}

		if src.Attributes != nil {
			for k, v := range src.Attributes {
				switch k {
				case "City":
					if v != "" {
						source.AddString("city", v)
					}
				case "Country":
					if v != "" {
						source.AddString("country", v)
					}
				case "County":
					if v != "" {
						source.AddString("county", v)
					}
				case "CountryArea":
					if v != "" {
						source.AddString("countryarea", v)
					}
				case "Id":
					if v != "" {
						source.AddString("code", v)
					}
				case "Url":
					if v != "" {
						source.AddString("url", v)
					}
				case "Type":
					if v != "" {
						source.AddString("type", v)
					}
				case "SubType":
					if v != "" {
						source.AddString("subtype", v)
					}
				}
			}
		}

		if src.Text != "" {
			source.AddString("name", src.Text)
		}

		if !source.IsEmpty() {
			sources.AddObject(source)
		}
	}

	if !sources.IsEmpty() {
		jo.AddArray("sources", sources)
	}
}

func getSourceMaterials(srcs []xml.Node, jo *json.Object) {
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sourcematerials := json.Array{}
	var cl bool

	for _, src := range srcs {
		var id, name, t, url, pg string

		if src.Nodes != nil {
			for _, n := range src.Nodes {
				switch n.Name {
				case "Type":
					t = n.Text
				case "Url":
					url = n.Text
				case "PermissionGranted":
					pg = n.Text
				}
			}
		}

		if src.Attributes != nil {
			for k, v := range src.Attributes {
				switch k {
				case "Id":
					id = v
				case "Name":
					name = v
				}
			}
		}

		if strings.EqualFold(name, "alternate") {
			if !cl && url != "" {
				jo.AddString("canonicallink", url)
				cl = true
			}
		} else {
			sourcematerial := json.Object{}
			if name != "" {
				sourcematerial.AddString("name", name)
			}
			if id != "" {
				sourcematerial.AddString("code", id)
			}
			if t != "" {
				sourcematerial.AddString("type", t)
			}
			if pg != "" {
				sourcematerial.AddString("permissiongranted", pg)
			}

			if !sourcematerial.IsEmpty() {
				sourcematerials.AddObject(sourcematerial)
			}
		}
	}

	if sourcematerials.Length() > 0 {
		jo.AddArray("sourcematerials", sourcematerials)
	}
}

func getRatings(rts []xml.Node, jo *json.Object) {
	if rts != nil {
		ratings := json.Array{}

		for _, r := range rts {
			if r.Attributes != nil {
				var (
					rate, min, max, raters int
					unit, rt, cr           string
				)

				for k, v := range r.Attributes {
					switch k {
					case "Value":
						if v != "" {
							i, err := strconv.Atoi(v)
							if err == nil {
								rate = i
							}
						} else {
							rate = -1
						}
					case "ScaleMin":
						if v != "" {
							i, err := strconv.Atoi(v)
							if err == nil {
								min = i
							}
						} else {
							min = -1
						}
					case "ScaleMax":
						if v != "" {
							i, err := strconv.Atoi(v)
							if err == nil {
								max = i
							}
						} else {
							max = -1
						}
					case "ScaleUnit":
						unit = v
					case "Raters":
						if v != "" {
							i, err := strconv.Atoi(v)
							if err == nil {
								raters = i
							}
						} else {
							raters = -1
						}
					case "RaterType":
						rt = v
					case "Creator":
						cr = v
					}
				}

				if rate != -1 && min != -1 && max != -1 && unit != "" {
					rating := json.Object{}
					rating.AddInt("rating", rate)
					rating.AddInt("scalemin", min)
					rating.AddInt("scalemax", max)
					rating.AddString("scaleunit", unit)
					if raters != -1 {
						rating.AddInt("raters", raters)
					}
					if rt != "" {
						rating.AddString("ratertype", rt)
					}
					if cr != "" {
						rating.AddString("creator", cr)
					}
					ratings.AddObject(rating)
				}
			}
		}

		if ratings.Length() > 0 {
			jo.AddArray("ratings", ratings)
		}
	}
}
