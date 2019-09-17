package appl

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func (doc *document) ParseDescriptiveMetadata(jo *json.Object) error {
	if doc.DescriptiveMetadata == nil || doc.DescriptiveMetadata.Nodes == nil || len(doc.DescriptiveMetadata.Nodes) == 0 {
		return errors.New("DescriptiveMetadata is missing")
	}

	var (
		desc, auds, gens, cats uniqueArray
		sups, alts, svcs       uniqueArray
		sbjs, orgs             subjects
		prns                   persons
		cmps                   companies
		plcs                   places
		evts                   events
		dll, fix               json.Object
		tpm                    json.Array
	)

	geo := true

	for _, nd := range doc.DescriptiveMetadata.Nodes {
		switch nd.Name {
		case "Description":
			desc.AddString(nd.Text)
		case "DateLineLocation":
			dll = getDatelineLocation(nd)
		case "SubjectClassification":
			fix = doc.ParseSubjectClassification(nd, &gens, &sbjs, &cats, &sups, &alts)
		case "EntityClassification":
			parseEntityClassification(nd, &gens, &orgs, &prns, &cmps, &plcs, &evts)
		case "AudienceClassification":
			test := getAudences(nd, &auds)
			if !test {
				geo = false
			}
		case "SalesClassification":
			parseSalesClassification(nd, &svcs)
		case "Comment":
			if nd.Text != "" {
				svc := json.Object{}
				svc.AddString("apservice", nd.Text)
				svcs.AddObject(nd.Text, svc)
			}
		case "ThirdPartyMeta":
			getThirdParty(nd, &tpm)
		}
	}

	if !fix.IsEmpty() {
		jo.AddObject("fixture", fix)
	}

	jo.AddProperty(desc.JSONProperty("descriptions"))

	if !dll.IsEmpty() {
		jo.AddObject("datelinelocation", dll)
	}

	jo.AddProperty(gens.JSONProperty("generators"))
	jo.AddProperty(cats.JSONProperty("categories"))
	jo.AddProperty(sups.JSONProperty("suppcategories"))
	jo.AddProperty(alts.JSONProperty("alertcategories"))
	jo.AddProperty(sbjs.JSONProperty("subjects"))
	jo.AddProperty(prns.JSONProperty(doc.Namelines))
	jo.AddProperty(orgs.JSONProperty("organizations"))
	jo.AddProperty(cmps.JSONProperty())
	jo.AddProperty(plcs.JSONProperty())
	jo.AddProperty(evts.JSONProperty())

	if geo && doc.Filings != nil {
		for _, f := range doc.Filings {
			if strings.EqualFold(f.Category, "n") {
				state := getState(f.Source)
				if state != nil {
					auds.AddObject(state.Code, state.ToJSON())
				}
			}
		}
	}
	jo.AddProperty(auds.JSONProperty("audiences"))

	jo.AddProperty(svcs.JSONProperty("services"))

	if !tpm.IsEmpty() {
		jo.AddArray("thirdpartymeta", tpm)
	}

	return nil
}

func (doc *document) ParseSubjectClassification(nd xml.Node, gens *uniqueArray, sbjs *subjects, cats *uniqueArray, sups *uniqueArray, alts *uniqueArray) json.Object {
	var fix json.Object

	if nd.Attributes != nil {
		auth, av := getAuthorities(nd)

		gens.AddKeyValue("name", auth, "version", av)

		if nd.Nodes != nil {
			key := strings.ToLower(auth)

			if key == "ap subject" {
				sbjs.Parse(nd)
			} else if key == "ap category code" {
				for _, n := range nd.Nodes {
					code, name := getOccurrenceCodeName(n)
					cats.AddKeyValue("code", code, "name", name)
				}
			} else if key == "ap supplemental category code" {
				for _, n := range nd.Nodes {
					code, name := getOccurrenceCodeName(n)
					sups.AddKeyValue("code", code, "name", name)
				}
			} else if key == "ap alert category" {
				for _, n := range nd.Nodes {
					code, _ := getOccurrenceCodeName(n)
					alts.AddString(code)
				}
			} else if !doc.Fixture && key == "ap audio cut number code" {
				for _, n := range nd.Nodes {
					code, name := getOccurrenceCodeName(n)
					if code != "" && name != "" {
						fix = json.Object{}
						fix.AddString("code", code)
						fix.AddString("name", name)
						doc.Fixture = true
						break
					}
				}
			}
		}
	}

	return fix
}

func parseEntityClassification(nd xml.Node, gens *uniqueArray, orgs *subjects, prns *persons, cmps *companies, plcs *places, evts *events) {
	if nd.Attributes != nil {
		auth, av := getAuthorities(nd)

		gens.AddKeyValue("name", auth, "version", av)

		if nd.Nodes != nil {
			key := strings.ToLower(auth)

			if key == "ap party" {
				prns.Parse(nd)
			} else if key == "ap organization" {
				orgs.Parse(nd)
			} else if key == "ap company" {
				cmps.Parse(nd)
			} else if key == "ap geography" || key == "ap country" || key == "ap region" {
				plcs.Parse(nd)
			} else if key == "ap event" {
				evts.Parse(nd)
			}
		}
	}
}

func parseSalesClassification(nd xml.Node, svcs *uniqueArray) {
	if nd.Attributes != nil {

		if nd.Nodes != nil {
			for _, n := range nd.Nodes {
				code, name := getOccurrenceCodeName(n)
				if code != "" && name != "" {
					jo := json.Object{}
					jo.AddString("code", code)
					jo.AddString("apsales", name)
					svcs.AddObject(code, jo)
				}
			}
		}
	}
}

func getDatelineLocation(nd xml.Node) json.Object {
	if nd.Nodes == nil {
		return json.Object{}
	}

	var lat, lng float64

	jo := json.Object{}

	for _, n := range nd.Nodes {
		switch n.Name {
		case "City":
			if n.Text != "" {
				jo.AddString("city", n.Text)
			}
		case "CountryArea":
			if n.Text != "" {
				jo.AddString("countryareacode", n.Text)
			}
		case "CountryAreaName":
			if n.Text != "" {
				jo.AddString("countryareaname", n.Text)
			}
		case "Country":
			if n.Text != "" {
				jo.AddString("countrycode", n.Text)
			}
		case "CountryName":
			if n.Text != "" {
				jo.AddString("countryname", n.Text)
			}
		case "LatitudeDD":
			f, err := strconv.ParseFloat(n.Text, 64)
			if err == nil {
				lat = f
			}
		case "LongitudeDD":
			f, err := strconv.ParseFloat(n.Text, 64)
			if err == nil {
				lng = f
			}
		}
	}

	jo.AddProperty(getGeoProperty(lat, lng))

	return jo
}

func getThirdParty(nd xml.Node, ja *json.Array) {
	if nd.Nodes == nil {
		return
	}

	jo := json.Object{}

	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "System":
				if v != "" {
					jo.AddString("creator", v)
				}
			case "Vocabulary":
				if v != "" {
					jo.AddString("vocabulary", v)
				}
			case "VocabularyOwner":
				if v != "" {
					jo.AddString("vocabularyowner", v)
				}
			}
		}
	}

	o := nd.Node("Occurrence")
	if o.Attributes != nil {
		for k, v := range o.Attributes {
			switch k {
			case "Id":
				if v != "" {
					jo.AddString("code", v)
				}
			case "Value":
				if v != "" {
					jo.AddString("name", v)
				}
			}
		}
	}

	if !jo.IsEmpty() {
		ja.AddObject(jo)
	}
}

func getAudences(nd xml.Node, ua *uniqueArray) bool {
	if nd.Nodes == nil || nd.Attributes == nil {
		return true
	}

	var (
		auth, system string
		geo          bool
	)

	for k, v := range nd.Attributes {
		switch k {
		case "Authority":
			auth = v
		case "System":
			system = v
		}
	}

	if strings.EqualFold(auth, "AP Audience") && strings.EqualFold(system, "Editorial") && nd.Nodes != nil {
		for _, o := range nd.Nodes {
			if o.Name == "Occurrence" {
				var id, value string

				if o.Attributes != nil {
					for k, v := range o.Attributes {
						switch k {
						case "Id":
							id = v
						case "Value":
							value = v
						}
					}
				}

				if id != "" && value != "" {
					jo := json.Object{}
					jo.AddString("code", id)
					jo.AddString("name", value)

					p := o.Node("Property")
					a := p.Attribute("Value")
					if a != "" {
						if strings.EqualFold(a, "AUDGEOGRAPHY") {
							geo = true
						}
						jo.AddString("type", a)
					}
					ua.AddObject(id, jo)
				}
			}
		}
	}

	return !geo
}

func getAuthorities(nd xml.Node) (string, string) {
	var auth, av string
	for k, v := range nd.Attributes {
		switch k {
		case "Authority":
			auth = v
		case "AuthorityVersion":
			av = v
		}
	}
	return auth, av
}

func getOccurrenceCodeName(nd xml.Node) (string, string) {
	var code, name string
	if nd.Name == "Occurrence" && nd.Attributes != nil {
		for k, v := range nd.Attributes {
			switch k {
			case "Id":
				code = v
			case "Value":
				name = v
			}
		}
	}
	return code, name
}
