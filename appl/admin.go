package appl

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (doc *document) ParseAdministrativeMetadata(jo *json.Object) error {
	if doc.AdministrativeMetadata == nil {
		return errors.New("AdministrativeMetadata is missing")
	}

	var (
		summary, s1, s2    bool
		tss, pss, dcs, ins uniqueArray
		srcs, sms, rts     []Node
		ict, fx            json.Object
		cntr               string
	)

	s1 = !doc.Signals.IsEmpty()

	for _, nd := range doc.NewsLines.Nodes {
		switch nd.Name {
		case "Provider":
			getProvider(nd, &jo)
		case "Source":
			if srcs == nil {
				srcs = []Node{nd}
			} else {
				srcs = append(srcs, nd)
			}
		case "SourceMaterial":
			if sms == nil {
				sms = []Node{nd}
			} else {
				sms = append(sms, nd)
			}
		case "TransmissionSource":
			if nd.Text != "" {
				if tss == nil {
					tss = uniqueArray{}
				}
				tss.AddString(nd.Text)
			}
		case "ProductSource":
			if nd.Text != "" {
				if pss == nil {
					pss = uniqueArray{}
				}
				pss.AddString(nd.Text)
			}
		case "DistributionChannel":
			if nd.Text != "" {
				if dcs == nil {
					dcs = uniqueArray{}
				}
				dcs.AddString(nd.Text)
			}
		case "InPackage":
			if nd.Text != "" {
				if ins == nil {
					ins = uniqueArray{}
				}
				tokens := strings.Split(nd.Text, " ")
				for _, token := range tokens {
					ins.AddString(token)
				}
			}
		case "ItemContentType":
			ict := json.Object{}
			if nd.Attributes != nil {
				for _, a := range nd.Attributes {
					switch a.Name {
					case "Id":
						if a.Value != "" {
							ict.AddString("code", a.Value)
						}
					case "System":
						if a.Value != "" {
							ict.AddString("creator", a.Value)
						}
					}
				}
			}

			if nd.Text != "" {
				ict.AddString("name", nd.Text)
			}
		case "Fixture":
			fx := json.Object{}
			id := nd.GetAttribute("Id")
			if id != "" {
				fx.AddString("code", id)
			}
			if nd.Text != "" {
				fx.AddString("name", nd.Text)
			}
		case "Rating":
			if rts == nil {
				rts = []Node{nd}
			} else {
				rts = append(rts, nd)
			}
		case "Reach":
			if nd.Text != "" && !strings.EqualFold(nd.Text, "UNKNOWN") {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString(nd.Text)
				s2 = true
			}
		case "ConsumerReady":
			if nd.Text != "" && strings.EqualFold(nd.Text, "TRUE") {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString("newscontent")
				s2 = true
			}
		case "Signal":
			if nd.Text != "" {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString("newscontent")
				s2 = true
			}
		case "Contributor":
			if cntr == "" {
				cntr = nd.Text
			}
		}
	}

	getSources(srcs, &jo)
	getSourceMaterials(sms, &jo)

	jo.AddProperty(tss.ToJsonProperty("transmissionsources"))
	jo.AddProperty(pss.ToJsonProperty("productsources"))

	if ict != nil && !ict.IsEmpty() {
		jo.AddObject("itemcontenttype", ict)
	}

	jo.AddProperty(dcs.ToJsonProperty("distributionchannels"))

	if fx != nil && !fx.IsEmpty() {
		jo.AddObject("fixture", fx)
	}

	jo.AddProperty(ins.ToJsonProperty("inpackages"))

	if cntr != "" {
		jo.AddString("contributor", cntr)
	}

	getRatings(rts, &jo)

	if s2 {
		if s1 {
			jo.SetArray("signals", doc.Signals.ToJsonArray())
		} else {
			jo.AddProperty(doc.Signals.ToJsonProperty("signals"))
		}
	}
}

func getProvider(nd Node, jo *json.Object) {
	provider := json.Object{}

	if nd.Attributes != nil {
		for _, a := range nd.Attributes {
			switch a.Name {
			case "Id":
				if a.Value != "" {
					provider.AddString("code", a.Value)
				}
			case "Type":
				if a.Value != "" {
					provider.AddString("type", a.Value)
				}
			case "SubType":
				if a.Value != "" {
					provider.AddString("subtype", a.Value)
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

func getSources(srcs []Node, jo *json.Object) {
	if srcs == nil || len(srcs) == 0 {
		return
	}

	sources := json.Array{}

	for _, src := range srcs {
		source := json.Object{}

		if src.Attributes != nil {
			for _, a := range src.Attributes {
				switch a.Name {
				case "City":
					if a.Value != "" {
						source.AddString("city", a.Value)
					}
				case "Country":
					if a.Value != "" {
						source.AddString("country", a.Value)
					}
				case "Id":
					if a.Value != "" {
						source.AddString("code", a.Value)
					}
				case "Url":
					if a.Value != "" {
						source.AddString("url", a.Value)
					}
				case "Type":
					if a.Value != "" {
						source.AddString("type", a.Value)
					}
				case "SubType":
					if a.Value != "" {
						source.AddString("subtype", a.Value)
					}
				}
			}

			if src.Text != "" {
				source.AddString("name", src.Text)
			}
		}

		if !source.IsEmpty() {
			sources.AddObject(source)
		}
	}

	if len(sources) > 0 {
		jo.AddArray("sources", sources)
	}
}

func getSourceMaterials(srcs []Node, jo *json.Object) {
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
					pm = n.Text
				}
			}
		}

		if src.Attributes != nil {
			for _, a := range src.Attributes {
				switch a.Name {
				case "Id":
					id = a.Value
				case "Name":
					name = a.Value
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

func getRatings(rts []Node, jo *json.Object) {
	if rts != nil {
		ratings := json.Array{}

		for _, r := range rts {
			if r.Attributes != nil {
				var (
					rate, min, max, raters int
					unit, rt               string
				)

				for _, a := range r.Attributes {
					switch a.Name {
					case "Value":
						if a.Value != "" {
							i, err := strconv.Atoi(nd.Text)
							if err == nil {
								rate = i
							}
						}
					case "ScaleMin":
						if a.Value != "" {
							i, err := strconv.Atoi(nd.Text)
							if err == nil {
								min = i
							}
						}
					case "ScaleMax":
						if a.Value != "" {
							i, err := strconv.Atoi(nd.Text)
							if err == nil {
								max = i
							}
						}
					case "ScaleUnit":
						unit = a.Value
					case "Raters":
						if a.Value != "" {
							i, err := strconv.Atoi(nd.Text)
							if err == nil {
								raters = i
							}
						}
					case "RaterType":
						rt = a.Value
					}
				}

				if rate > 0 && min > 0 && max > 0 && unit != "" {
					rating := json.Object{}
					rating.AddInt("rating", rate)
					rating.AddInt("scalemin", min)
					rating.AddInt("scalemax", max)
					rating.AddString("scaleunit", unit)
					if raters > 0 {
						rating.AddInt("raters", raters)
					}
					if rt != "" {
						rating.AddString("ratertype", rt)
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
