package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/xml"
)

type subjectParser struct {
	Alerts          uniqueStrings
	Categories      uniqueCodeNames
	SuppCategories  uniqueCodeNames
	AlertCategories uniqueStrings
	Subjects        []Subject
	FixtureName     string
	keys            map[string]int
}

type entityParser struct {
	Persons       []Person
	Organizations subjectParser
	Companies     []Company
	Places        []Place
	Events        []Event
	pkeys         map[string]int
	ckeys         map[string]int
	lkeys         map[string]int
	ekeys         map[string]bool
}

type audienceParser struct {
	Audiences []CodeNameTitle
	keys      map[string]bool
}

func (doc *Document) parseDescriptiveMetadata(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	var (
		desc      uniqueStrings
		gens      uniqueCodeNames
		subjects  subjectParser
		entities  entityParser
		audiences audienceParser
		services  uniqueCodeNames
	)

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
			parseAudenceClassification(nd, &audiences)
		case "SalesClassification":
			parseSalesClassification(nd, &services)
		case "Comment":
			if nd.Text != "" {
				services.Append("_apservice", nd.Text)
			}
		case "ThirdPartyMeta":
			doc.parseThirdParty(nd)
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
	doc.Organizations = entities.Organizations.Subjects
	doc.Companies = entities.Companies
	doc.Places = entities.Places
	doc.Events = entities.Events
	doc.Audiences = audiences.Audiences
	doc.Services = services.Values()

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

	*/
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

	doc.DateLineLocation = &loc

	return
}

func (doc *Document) parseSubjectClassification(node xml.Node, gens *uniqueCodeNames, subjects *subjectParser) {
	if node.Attributes == nil {
		return
	}

	auth, authv := getAuthorities(node)
	gens.Append(authv, auth)

	if node.Nodes == nil {
		return
	}

	key := strings.ToLower(auth)

	if key == "ap subject" {
		subjects.parseSubject(node)
	} else if key == "ap category code" {
		for _, nd := range node.Nodes {
			code, name := getOccurrenceCodeName(nd)
			subjects.Categories.Append(code, name)
			if subjects.FixtureName == "" {
				subjects.FixtureName = name
			}
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
		nd := node.Node("Occurrence")
		code := nd.Attribute("Value")
		if code != "" {
			test, err := strconv.Atoi(code)
			if err == nil && test >= 900 {
				doc.Fixture = &CodeName{
					Code: code,
					Name: subjects.FixtureName,
				}
			}
		}
	}
}

func parseEntityClassification(node xml.Node, gens *uniqueCodeNames, entities *entityParser) {
	if node.Attributes == nil {
		return
	}

	auth, authv := getAuthorities(node)
	gens.Append(auth, authv)

	if node.Nodes == nil {
		return
	}

	key := strings.ToLower(auth)

	if key == "ap party" {
		entities.parsePerson(node)
	} else if key == "ap organization" {
		entities.Organizations.parseSubject(node)
	} else if key == "ap company" {
		entities.parseCompany(node)
	} else if key == "ap geography" || key == "ap country" || key == "ap region" {
		entities.parsePlace(node)
	} else if key == "ap event" {
		entities.parseEvent(node)
	}
}

func parseAudenceClassification(node xml.Node, audiences *audienceParser) {
	if node.Nodes == nil || node.Attributes == nil {
		return
	}

	var auth, system string

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
					if audiences.keys == nil {
						audiences.keys = make(map[string]bool)
					}
					_, ok := audiences.keys[id]
					if ok {
						continue
					} else {
						audiences.keys[id] = true
					}

					cnt := CodeNameTitle{
						Code: id,
						Name: value,
					}

					p := o.Node("Property")
					cnt.Title = p.Attribute("Value")
					audiences.Audiences = append(audiences.Audiences, cnt)
				}
			}
		}
	}
}

func parseSalesClassification(node xml.Node, services *uniqueCodeNames) {
	if node.Nodes == nil || node.Attributes == nil {
		return
	}

	for _, nd := range node.Nodes {
		code, name := getOccurrenceCodeName(nd)
		if code != "" && name != "" {
			services.Append(code, name)
		}
	}
}

func (doc *Document) parseThirdParty(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	var tp ThirdParty

	if node.Attributes != nil {
		for k, v := range node.Attributes {
			switch k {
			case "System":
				tp.Creator = v
			case "Vocabulary":
				tp.Vocabulary = v
			case "VocabularyOwner":
				tp.VocabularyOwner = v
			}
		}
	}

	o := node.Node("Occurrence")
	if o.Attributes != nil {
		for k, v := range o.Attributes {
			switch k {
			case "Id":
				tp.Creator = v
			case "Value":
				tp.Name = v
			}
		}
	}

	doc.ThirdParties = append(doc.ThirdParties, tp)
}

func (parser *subjectParser) parseSubject(node xml.Node) {
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
				var subject Subject

				key := fmt.Sprintf("%s_%s", code, name)

				if parser.keys == nil {
					parser.keys = make(map[string]int)
				}

				i, ok := parser.keys[key]
				if ok {
					subject = parser.Subjects[i]
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
				subject.ParentIDs = subject.ids.Values()

				if !ok {
					parser.Subjects = append(parser.Subjects, subject)
					parser.keys[key] = len(parser.Subjects) - 1
				}
			}
		}
	}
}

func (parser *entityParser) parsePerson(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	system := node.Attribute("System")

	for _, nd := range node.Nodes {
		code, name := getOccurrenceCodeName(nd)
		if code != "" && name != "" {
			var person Person

			key := fmt.Sprintf("%s_%s", code, name)

			if parser.pkeys == nil {
				parser.pkeys = make(map[string]int)
			}

			i, ok := parser.pkeys[key]
			if ok {
				person = parser.Persons[i]
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
						var id, nm, va string
						for k, v := range p.Attributes {
							switch k {
							case "Id":
								id = v
							case "Name":
								nm = v
							case "Value":
								va = v
							}
						}

						if nm != "" && va != "" {
							key := strings.ToLower(nm)
							if key == "partytype" {
								if strings.EqualFold(va, "PERSON_FEATURED") {
									person.rels.Append(va)
									if person.Creator == "" {
										person.Creator = "Editorial"
									}
								} else {
									person.types.Append(va)
								}
							} else if key == "team" && id != "" {
								person.teams.Append(id, va)
							} else if key == "associatedevent" && id != "" {
								person.events.Append(id, va)
							} else if key == "associatedstate" && id != "" {
								person.states.Append(id, va)
							} else if key == "extid" {
								person.ids.Append(va)
							}
						}
					}
				}
			}

			person.Rels = person.rels.Values()
			person.Types = person.types.Values()
			person.IDs = person.ids.Values()
			person.Teams = person.teams.Values()
			person.States = person.states.Values()
			person.Events = person.events.Values()

			if !ok {
				parser.Persons = append(parser.Persons, person)
				parser.pkeys[key] = len(parser.Persons) - 1
			}
		}
	}
}

func (parser *entityParser) parseCompany(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	system := node.Attribute("System")

	for _, nd := range node.Nodes {
		code, name := getOccurrenceCodeName(nd)
		if code != "" && name != "" {
			var company Company

			key := fmt.Sprintf("%s_%s", code, name)

			if parser.ckeys == nil {
				parser.ckeys = make(map[string]int)
			}

			i, ok := parser.ckeys[key]
			if ok {
				company = parser.Companies[i]
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
						var id, pid, nm, va string
						for k, v := range p.Attributes {
							switch k {
							case "Id":
								id = v
							case "Name":
								nm = v
							case "Value":
								va = v
							case "ParentId":
								pid = v
							}
						}

						if nm != "" && va != "" {
							key := strings.ToLower(nm)
							if key == "apindustry" && id != "" {
								company.industries.Append(id, va)
							} else if key == "instrument" {
								company.symbols.Append(strings.ToUpper(va))
							} else if key == "primaryticker" || key == "ticker" {
								company.tickers.Append(pid, strings.ToUpper(va))
							} else if key == "exchange" {
								if company.exchanges == nil {
									company.exchanges = make(map[string]string)
								}
								company.exchanges[id] = strings.ToUpper(va)
							}
						}
					}
				}
			}

			if !company.tickers.IsEmpty() && company.exchanges != nil {
				var (
					def     string
					tickers = company.tickers.Values()
				)
				for _, ticker := range tickers {
					var exchange string

					ex, ok := company.exchanges[ticker.Code]
					if ok {
						exchange = ex
					} else {
						if def == "" && len(company.exchanges) > 0 {
							for _, v := range company.exchanges {
								def = v
								break
							}

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
				parser.Companies = append(parser.Companies, company)
				parser.ckeys[key] = len(parser.Companies) - 1
			}
		}
	}
}

func (parser *entityParser) parsePlace(node xml.Node) {
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
					place Place
					geo   Geo
				)

				key := fmt.Sprintf("%s_%s", code, name)

				if parser.lkeys == nil {
					parser.lkeys = make(map[string]int)
				}

				i, ok := parser.lkeys[key]
				if ok {
					place = parser.Places[i]
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
							var id, nm, va string
							for k, v := range p.Attributes {
								switch k {
								case "Id":
									id = v
								case "Name":
									nm = v
								case "Value":
									va = v
								}
							}

							if nm != "" && va != "" {
								key := strings.ToLower(nm)
								if key == "locationtype" && id != "" && place.LocationType == nil {
									place.LocationType = &CodeName{
										Code: id,
										Name: va,
									}
								} else if key == "centroidlatitude" && geo.Latitude == 0 {
									f, err := strconv.ParseFloat(va, 64)
									if err == nil {
										geo.Latitude = f
									}
								} else if key == "centroidlongitude" && geo.Longitude == 0 {
									f, err := strconv.ParseFloat(va, 64)
									if err == nil {
										geo.Longitude = f
									}
								}
							}
						}
					}
				}

				if geo.Longitude != 0 && geo.Latitude != 0 {
					place.Geo = &geo
				}

				place.rels.AppendRel(system, match)
				place.ids.Append(pid)

				place.Rels = place.rels.Values()
				place.ParentIDs = place.ids.Values()

				if !ok {
					parser.Places = append(parser.Places, place)
					parser.lkeys[key] = len(parser.Places) - 1
				}
			}
		}
	}
}

func (parser *entityParser) parseEvent(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	system := node.Attribute("System")

	for _, nd := range node.Nodes {
		code, name := getOccurrenceCodeName(nd)
		if code != "" && name != "" {
			var event Event

			key := fmt.Sprintf("%s_%s", code, name)

			if parser.ekeys == nil {
				parser.ekeys = make(map[string]bool)
			}

			_, ok := parser.ekeys[key]
			if ok {
				continue
			} else {
				parser.ekeys[key] = true
			}

			var (
				epix               = strings.EqualFold("epix", system)
				extid, extidsource string
			)

			if nd.Nodes != nil {
				for _, p := range nd.Nodes {
					if p.Attributes != nil {
						var nm, va string
						for k, v := range p.Attributes {
							switch k {
							case "Name":
								nm = v
							case "Value":
								va = v
							}
						}

						if nm != "" && va != "" {
							key := strings.ToLower(nm)
							val := strings.ToLower(va)
							if key == "extid" {
								extid = val
							} else if key == "extidsource" {
								extidsource = val
								if !epix {
									epix = val == "nfl" || val == "sportradar"
								}
							} else {
								event.props.Append(key, va)
							}
						}
					}
				}
			}

			event.Name = name
			if epix {
				event.Creator = "ePix"
				event.ExternalIDs = []CodeName{
					CodeName{
						Code: code,
						Name: "sportradar",
					},
				}
				if extid != "" && extidsource != "" {
					cn := CodeName{
						Code: extid,
						Name: extidsource,
					}
					event.ExternalIDs = append(event.ExternalIDs, cn)
				}
			} else {
				event.Code = code
				if system != "" {
					event.Creator = system
				}
			}
			event.Properties = event.props.Values()

			parser.Events = append(parser.Events, event)
		}
	}
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
