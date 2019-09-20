package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type subjectParser struct {
	Alerts          uniqueStrings
	Categories      uniqueCodeNames
	SuppCategories  uniqueCodeNames
	AlertCategories uniqueStrings
	Subjects        []Subject
	keys            map[string]int
}

type entityParser struct {
	Persons       []Person
	Organizations []Subject
	Companies     []Company
	Places        []Place
	Events		[]Event	
	okeys         map[string]int
	pkeys         map[string]int
	ckeys         map[string]int
	lkeys         map[string]int
	ekeys         map[string]bool
}

func (doc *Document) parseDescriptiveMetadata(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	var (
		desc     uniqueStrings
		gens     uniqueCodeNames
		subjects subjectParser
		entities entityParser
		tpm      json.Array
	)

	//geo := true

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "Description":
			desc.Append(nd.Text)
		case "DateLineLocation":
			doc.parseDatelineLocation(nd)
		case "SubjectClassification":
			doc.parseSubjectClassification(nd, &gens, &subjects)
		case "EntityClassification":
			parseEntityClassification(nd, &gens, &entities)
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

	if !desc.IsEmpty() {
		doc.Descriptions = desc.Values()
	}
	if !gens.IsEmpty() {
		doc.Generators = gens.Values()
	}
	if !subjects.Categories.IsEmpty() {
		doc.Categories = subjects.Categories.Values()
	}
	if !subjects.SuppCategories.IsEmpty() {
		doc.SuppCategories = subjects.SuppCategories.Values()
	}
	if !subjects.Alerts.IsEmpty() {
		doc.AlertCategories = subjects.Alerts.Values()
	}
	doc.Subjects = subjects.Subjects
	doc.Persons = entities.Persons
	doc.Organizations = entities.Organizations
	doc.Companies = entites.Companies
	doc.Events = entities.Events

	/*
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
	*/
}

func (doc *Document) parseSubjectClassification(node xml.Node, gens *uniqueCodeNames, subjects *subjectParser) {
	if node.Attributes == nil {
		return
	}

	auth, authv := getAuthorities(node)
	gens.Append(auth, authv)

	if node.Nodes == nil {
		return
	}

	key := strings.ToLower(auth)

	if key == "ap subject" {
		parseSubjects(node, &subjects.Subjects, &subjects.keys)
	} else if key == "ap category code" {
		for _, nd := range node.Nodes {
			code, name := getOccurrenceCodeName(nd)
			subjects.Categories.Append(code, name)
		}
	} else if key == "ap supplemental category code" {
		for _, nd := range node.Nodes {
			code, name := getOccurrenceCodeName(nd)
			subjects.SuppCategories.Append(code, name)
		}
	} else if key == "ap alert category" {
		for _, nd := range node.Nodes {
			code, _ := getOccurrenceCodeName(nd)
			subjects.Alerts.Append(code)
		}
	} else if doc.Fixture == nil && key == "ap audio cut number code" {
		for _, nd := range node.Nodes {
			code, name := getOccurrenceCodeName(nd)
			if code != "" && name != "" {
				test, err := strconv.Atoi(nd.Text)
				if err == nil && test >= 900 {
					doc.Fixture = &CodeName{
						Code: code,
						Name: name,
					}
				}
			}
		}
	}
}

func parseEntityClassification(node xml.Node, gens *uniqueCodeNames, entities *entityParser) {
	if node.Attributes != nil {
		auth, authv := getAuthorities(node)
		gens.Append(auth, authv)

		if node.Nodes != nil {
			key := strings.ToLower(auth)

			if key == "ap party" {
				parsePerson(node, &entities.Persons, &entities.pkeys)
			} else if key == "ap organization" {
				parseSubject(node, &entities.Companies, &entities.ckeys)
			} else if key == "ap company" {
				parseCompany(node, &entities.Organizations, &entities.okeys)
			} else if key == "ap geography" || key == "ap country" || key == "ap region" {
				parsePlace(node, &entities.Places, &entities.lkeys)
			} else if key == "ap event" {
				parseEvent(node, &entities.Events, &entities.ekeys)
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

func (doc *Document) parseDatelineLocation(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	var (
		loc Location
		geo Geo
	)

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "City":
			loc.City = nd.Text
		case "CountryArea":
			loc.CountryAreaCode = nd.Text
		case "CountryAreaName":
			loc.CountryAreaName = nd.Text
		case "Country":
			loc.CountryCode = nd.Text
		case "CountryName":
			loc.CountryName = nd.Text
		case "LatitudeDD":
			f, err := strconv.ParseFloat(nd.Text, 64)
			if err == nil {
				geo.Latitude = f
			}
		case "LongitudeDD":
			f, err := strconv.ParseFloat(nd.Text, 64)
			if err == nil {
				geo.Longitude = f
			}
		}
	}

	if geo.Longitude != 0 && geo.Latitude != 0 {
		loc.Geo = &geo
	}

	return
}

func parseSubject(node xml.Node, psubjects *[]Subject, pkeys *map[string]int) {
	if node.Nodes == nil {
		return
	}

	system := node.Attribute("System")

	for _, nd := range node.Nodes {
		if nd.Name == "Occurrence" && nd.Attributes != nil {
			var (
				code, name, pid string
				match, tp       *bool
			)

			for k, v := range nd.Attributes {
				switch k {
				case "Id":
					code = v
				case "Value":
					name = v
				case "ActualMatch":
					if v != "" {
						test := strings.EqualFold(v, "true")
						match = &test
					}
				case "ParentId":
					pid = v
				case "TopParent":
					if v != "" {
						test := strings.EqualFold(v, "true")
						tp = &test
					}
				}
			}

			if code != "" && name != "" {
				var (
					subject  Subject
					subjects []Subject
					keys     map[string]int
				)

				key := fmt.Sprintf("%s_%s", code, name)

				if pkeys == nil {
					keys = make(map[string]int)
				} else {
					keys = *pkeys
				}

				if psubjects != nil {
					subjects = *psubjects
				}

				i, ok := keys[key]
				if ok {
					subject = subjects[i]
				} else {
					subject = Subject{
						Name:      name,
						Code:      code,
						Creator:   system,
						TopParent: tp,
					}
				}

				if subject.Creator == "" || strings.EqualFold(system, "Editorial") {
					subject.Creator = "Editorial"
				}

				subject.rels.AppendRel(system, match)
				subject.ids.Append(pid)

				subject.Rels = subject.rels.Values()
				subject.ParentIds = subject.ids.Values()

				if !ok {
					subjects = append(subjects, subject)
					keys[key] = len(subjects) - 1
				}

				psubjects = &subjects
				pkeys = &keys
			}
		}
	}
}

func parsePerson(node xml.Node, ppersons *[]Person, pkeys *map[string]int) {
	if node.Nodes == nil {
		return
	}

	system := node.Attribute("System")

	for _, nd := range node.Nodes {
		code, name := getOccurrenceCodeName(nd)
		if code != "" && name != "" {
			var (
				person  Person
				persons []Person
				keys    map[string]int
			)

			key := fmt.Sprintf("%s_%s", code, name)

			if pkeys == nil {
				keys = make(map[string]int)
			} else {
				keys = *pkeys
			}

			if ppersons != nil {
				persons = *ppersons
			}

			i, ok := keys[key]
			if ok {
				person = persons[i]
			} else {
				person = Person{
					Name:    name,
					Code:    code,
					Creator: system,
				}
				person.rels.Append("direct")
			}

			if person.Creator == "" || strings.EqualFold(system, "Editorial") {
				person.Creator = "Editorial"
			}

			if nd.Nodes != nil {
				for _, p := range nd.Nodes {
					if p.Name == "Property" && p.Attributes != nil {
						var id, n, v string
						for k, v := range p.Attributes {
							switch k {
							case "Id":
								id = v
							case "Name":
								n = v
							case "Value":
								v = v
							}
						}

						if n != "" && v != "" {
							key := strings.ToLower(n)
							if key == "partytype" {
								if strings.EqualFold(v, "PERSON_FEATURED") {
									person.rels.Append(v)
									if person.Creator == "" {
										person.Creator = "Editorial"
									}
								} else {
									person.types.Append(v)
								}
							} else if key == "team" && id != "" {
								person.teams.Append(id, v)
							} else if key == "associatedevent" && id != "" {
								person.events.Append(id, v)
							} else if key == "associatedstate" && id != "" {
								person.states.Append(id, v)
							} else if key == "extid" {
								person.ids.Append(v)
							}
						}
					}
				}
			}

			person.Rels = person.rels.Values()
			person.Types = person.types.Values()
			person.Ids = person.ids.Values()
			person.Teams = person.teams.Values()
			person.States = person.states.Values()
			person.Events = person.events.Values()

			if !ok {
				persons = append(persons, person)
				keys[key] = len(persons) - 1
			}

			ppersons = &persons
			pkeys = &keys
		}
	}
}

func parseCompany(node xml.Node, pcompanies *[]Company, ckeys *map[string]int) {
	if node.Nodes == nil {
		return
	}

	system := node.Attribute("System")

	for _, nd := range node.Nodes {
		code, name := getOccurrenceCodeName(nd)
		if code != "" && name != "" {
			var (
				company  Company
				companies   []Company
				keys    map[string]int
				keys      map[string]int

			key := fmt.Sprintf("%s_%s", code, name)

			if ckeys == nil {
				keys = make(map[string]int)
			} else {
				keys = *ckeys
			}

			if pcompanies != nil {
				companies = *pcompanies
			}

			i, ok := keys[key]
			if ok {
				company = companies[i]
			} else {
				company = Company{
					Name:    name,
					Code:    code,
					Creator: system,
				}
				company.rels.Append("direct")
			}

			if company.Creator == "" || strings.EqualFold(system, "Editorial") {
				company.Creator = "Editorial"
			}

			if nd.Nodes != nil {
				for _, p := range nd.Nodes {
					if p.Name == "Property" && p.Attributes != nil {
						var id, pid, n, v string
						for k, v := range p.Attributes {
							switch k {
							case "Id":
								id = v
							case "Name":
								n = v
							case "Value":
								v = v
							case "ParentId":
								pid = v
							}
						}

						if n != "" && v != "" {
							key := strings.ToLower(n)
							if key == "apindustry" && id != "" {
								company.industries.Append(id, v)
							} else if key == "instrument" {
								company.symbols.Append(strings.ToUpper(v))
							} else if key == "primaryticker" || key == "ticker" {
								company.tickers.Append(pid, strings.ToUpper(v))
							} else if key == "exchange" {
								if company.exchanges == nil {
									company.exchanges = make(map[string]string)
								}
								company.exchanges[id] = strings.ToUpper(v)
							}
						}						
						}
				}
			}

			if company.tickers != nil && company.exchanges != nil {
				var def string
				for _, ticker := range company.tickers {
					var exchange string

					ex, ok := exchanges[ticker.Code]
					if ok {
						exchange = ex
					} else {
						if def == "" && len(company.exchanges) > 0 {
							def = company.exchanges[0]
						}
						exchange = def
					}

					if exchange != "" {
						instrument := fmt.Sprintf("%s:%s", exchange, ticker.Name)
						company.symbols.Append(instrument)
					}
				}
			}

			company.Rels = company.rels.Values()
			company.Industries = company.industries.Values()
			company.Symbols = company.symbols.Values()

			if !ok {
				companies = append(companies, company)
				keys[key] = len(companies) - 1
			}

			pcompanies = &companies
			ckeys = &keys
		}
	}
}

func parsePlace(node xml.Node, pplaces *[]Place, pkeys *map[string]int) {
	if node.Nodes == nil {
		return
	}

	system := node.Attribute("System")

	for _, nd := range node.Nodes {
		if nd.Name == "Occurrence" && nd.Attributes != nil {
			var (
				code, name, pid string
				match, tp       *bool
			)

			for k, v := range nd.Attributes {
				switch k {
				case "Id":
					code = v
				case "Value":
					name = v
				case "ActualMatch":
					if v != "" {
						test := strings.EqualFold(v, "true")
						match = &test
					}
				case "ParentId":
					pid = v
				case "TopParent":
					if v != "" {
						test := strings.EqualFold(v, "true")
						tp = &test
					}
				}
			}

			if code != "" && name != "" {
				var (
					place  Place
					places []Place
					keys     map[string]int
					keys   map[string]int
					geo    Geo

				key := fmt.Sprintf("%s_%s", code, name)

				if pkeys == nil {
					keys = make(map[string]int)
				} else {
					keys = *pkeys
				}

				if pplaces != nil {
					places = *pplaces
				}

				i, ok := keys[key]
				if ok {
					place = places[i]
				} else {
					place = Place{
						Name:      name,
						Code:      code,
						Creator:   system,
						TopParent: tp,
					}
				}

				if place.Creator == "" || strings.EqualFold(system, "Editorial") {
					place.Creator = "Editorial"
				}

				if nd.Nodes != nil {
					for _, p := range nd.Nodes {
						if p.Name == "Property" && p.Attributes != nil {
							var id, n, v string
							for k, v := range p.Attributes {
								switch k {
								case "Id":
									id = v
								case "Name":
									n = v
								case "Value":
									v = v
								}
							}

							if n != "" && v != "" {
								key := strings.ToLower(n)
								if key == "locationtype" && id != "" && place.LocationType == nil {
									place.LocationType = &CodeName{
										Code: id,
										Name: v,
									}
								} else if key == "centroidlatitude" && geo.Latitude == 0 {
									f, err := strconv.ParseFloat(v, 64)
									if err == nil {
										geo.Latitude = f
									}
								} else if key == "centroidlongitude" && geo.Longitude == 0 {
									f, err := strconv.ParseFloat(v, 64)
									if err == nil {
										geo.Longitude = f
									}
								}
							}							
							}
					}
				}

				if geo.Longitude !=0 && geo.Latitude != 0{
				if geo.Longitude != 0 && geo.Latitude != 0 {
				}

				place.rels.AppendRel(system, match)
				place.ids.Append(pid)

				place.Rels = place.rels.Values()
				place.ParentIds = place.ids.Values()

				if !ok {
					places = append(places, place)
					keys[key] = len(places) - 1
				}

				pplaces = &places
				pkeys = &keys
			}
		}
	}
}

func parseEvent(node xml.Node, pevents *[]Event, pkeys *map[string]bool) {
	if node.Nodes == nil {
		return
	}
/*
	system := node.Attribute("System")

	for _, n := range nd.Nodes {
		code, name := getOccurrenceCodeName(n)
		if code != "" && name != "" {
			key := fmt.Sprintf("%s_%s", code, name)

			if es.Keys == nil {
				es.Keys = make(map[string]bool)
				es.Events = json.Array{}
			}

			_, ok := es.Keys[key]
			if ok {
				continue
			} else {
				es.Keys[key] = true
			}

			epix := strings.EqualFold("epix", system)
			extid := ""
			extidsource := ""
			eventproperties := json.Object{}

			if n.Nodes != nil {
				for _, p := range n.Nodes {
					if p.Attributes != nil {
						var n, v string
						for k, v := range p.Attributes {
							switch k {
							case "Name":
								n = v
							case "Value":
								v = v
							}
						}

						if n != "" && v != "" {
							key := strings.ToLower(n)
							val := strings.ToLower(v)
							if key == "extid" {
								extid = val
							} else if key == "extidsource" {
								extidsource = val
								if !epix {
									epix = val == "nfl" || val == "sportradar"
								}
							} else {
								eventproperties.AddString(key, v)
							}
						}
					}
				}
			}

			e := json.Object{}

			if epix {
				externaleventids := json.Array{}

				id1 := json.Object{}
				id1.AddString("code", code)
				id1.AddString("creator", "sportradar")
				id1.AddString("creatorcode", "sportradar:"+code)
				externaleventids.AddObject(id1)

				if extid != "" && extidsource != "" {
					id2 := json.Object{}
					id2.AddString("code", extid)
					id2.AddString("creator", extidsource)
					id2.AddString("creatorcode", fmt.Sprintf("%s:%s", extidsource, extid))
					externaleventids.AddObject(id2)
				}

				e.AddString("name", name)
				e.AddString("creator", "ePix")
				e.AddArray("externaleventids", externaleventids)

			} else {
				e.AddString("code", code)
				e.AddString("name", name)
				if system != "" {
					e.AddString("creator", system)
				}
			}

			if !eventproperties.IsEmpty() {
				e.AddObject("eventproperties", eventproperties)
			}

			if !e.IsEmpty() {
				es.Events.AddObject(e)
			}
		}
	}
	*/
}

func getThirdParty(node xml.Node, ja *json.Array) {
	if node.Nodes == nil {
		return
	}

	jo := json.Object{}

	if node.Attributes != nil {
		for k, v := range node.Attributes {
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

	o := node.Node("Occurrence")
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

func getAudences(node xml.Node, ua *uniqueArray) bool {
	if node.Nodes == nil || node.Attributes == nil {
		return true
	}

	var (
		auth, system string
		geo          bool
	)

	for k, v := range node.Attributes {
		switch k {
		case "Authority":
			auth = v
		case "System":
			system = v
		}
	}

	if strings.EqualFold(auth, "AP Audience") && strings.EqualFold(system, "Editorial") && node.Nodes != nil {
		for _, o := range node.Nodes {
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

func getAuthorities(node xml.Node) (auth string, authv string) {
	for k, v := range node.Attributes {
		switch k {
		case "Authority":
			auth = v
		case "AuthorityVersion":
			authv = v
		}
	}
	return
}

func getOccurrenceCodeName(node xml.Node) (string, string) {
	var code, name string
	if node.Name == "Occurrence" && node.Attributes != nil {
		for k, v := range node.Attributes {
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
