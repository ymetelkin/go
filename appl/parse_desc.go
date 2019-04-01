package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

type subject struct {
	Name      string
	Code      string
	Creator   string
	Rels      uniqueArray
	ParentIds uniqueArray
	TopParent bool
}

type person struct {
	Name    string
	Code    string
	Creator string
	Rels    *uniqueArray
	Types   uniqueArray
	Ids     uniqueArray
	Teams   uniqueArray
	States  uniqueArray
	Events  uniqueArray
}

func (desc *DescriptiveMetadata) parse(doc *document) error {
	getDescriptions(doc)
	getDatelineLocation(doc)
	getClassification(doc)

	return nil
}

func getDescriptions(doc *document) {
	descs := doc.Xml.DescriptiveMetadata.Description
	if descs != nil {
		descriptions := uniqueArray{}
		for _, desc := range descs {
			descriptions.AddString(desc)
		}
		doc.Descriptions = descriptions.ToJsonProperty("descriptions")
	}
}

func getDatelineLocation(doc *document) {
	dll := doc.Xml.DescriptiveMetadata.DateLineLocation

	datelinelocation := json.Object{}
	if dll.City != "" {
		datelinelocation.AddString("city", dll.City)
	}
	if dll.CountryArea != "" {
		datelinelocation.AddString("countryareacode", dll.CountryArea)
	}
	if dll.CountryAreaName != "" {
		datelinelocation.AddString("countryareaname", dll.CountryAreaName)
	}
	if dll.Country != "" {
		datelinelocation.AddString("countrycode", dll.Country)
	}
	if dll.CountryName != "" {
		datelinelocation.AddString("countryname", dll.CountryName)
	}
	if dll.LatitudeDD != 0 && dll.LongitudeDD != 0 {
		coordinates := json.Array{}
		coordinates.AddFloat(dll.LongitudeDD)
		coordinates.AddFloat(dll.LatitudeDD)
		geometry := json.Object{}
		geometry.AddString("type", "Point")
		geometry.AddArray("coordinates", &coordinates)

		datelinelocation.AddObject("geometry_geojson", &geometry)
	}

	if !datelinelocation.IsEmpty() {
		doc.DatelineLocation = json.NewObjectProperty("datelinelocation", &datelinelocation)
	}
}

func getClassification(doc *document) {
	generators := uniqueArray{}
	categories := uniqueArray{}
	suppcategories := uniqueArray{}
	subjects := []subject{}
	subjectKeys := make(map[string]int)
	orgs := []subject{}
	orgKeys := make(map[string]int)
	persons := []person{}
	personKeys := make(map[string]int)
	alerts := uniqueArray{}

	classification := doc.Xml.DescriptiveMetadata.SubjectClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			generators.AddKeyValue("name", c.Authority, "version", c.AuthorityVersion)

			authority := strings.ToLower(c.Authority)

			if c.Occurrence != nil {
				if authority == "ap subject" {
					subjects = getSubjects(c, subjectKeys, subjects)
				} else if authority == "ap category code" {
					for _, o := range c.Occurrence {
						categories.AddKeyValue("code", o.Id, "name", o.Value)
					}
				} else if authority == "ap supplemental category code" {
					for _, o := range c.Occurrence {
						suppcategories.AddKeyValue("code", o.Id, "name", o.Value)
					}
				} else if authority == "ap alert category" {
					for _, o := range c.Occurrence {
						alerts.AddString(o.Id)
					}
				} else if doc.Fixture == nil && authority == "ap audio cut number code" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
							i, err := strconv.Atoi(o.Id)
							if err == nil && i >= 900 {
								fixture := json.Object{}
								fixture.AddInt("code", i)
								fixture.AddString("name", o.Value)
								doc.Fixture = json.NewObjectProperty("fixture", &fixture)
								break
							}
						}
					}
				}
			}

		}
	}

	classification = doc.Xml.DescriptiveMetadata.EntityClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			generators.AddKeyValue("name", c.Authority, "version", c.AuthorityVersion)

			authority := strings.ToLower(c.Authority)

			if c.Occurrence != nil {
				if authority == "ap party" {
					persons = getPersons(c, personKeys, persons)
				} else if authority == "ap organization" {
					orgs = getSubjects(c, orgKeys, orgs)
				} else if authority == "ap company" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
						}
					}
				} else if authority == "ap geography" || authority == "ap country" || authority == "ap region" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
						}
					}
				} else if authority == "ap event" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
						}
					}
				}
			}
		}
	}

	doc.Generators = generators.ToJsonProperty("generators")
	doc.Categories = categories.ToJsonProperty("categories")
	doc.SuppCategories = suppcategories.ToJsonProperty("suppcategories")
	doc.AlertCategories = alerts.ToJsonProperty("alertcategories")

	doc.Subjects = setSubjects("subjects", subjects)
	doc.Persons = setPersons(persons)
	doc.Organizations = setSubjects("organizations", orgs)
}

func getSubjects(c Classification, keys map[string]int, subjects []subject) []subject {
	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			var sbj subject

			key := fmt.Sprintf("%s_%s", o.Id, o.Value)
			i, ok := keys[key]
			if ok {
				sbj = subjects[i]
			} else {
				sbj = subject{Name: o.Value, Code: o.Id, Creator: c.System}
				subjects = append(subjects, sbj)
				i = len(subjects) - 1
				keys[key] = i
			}

			if sbj.Creator == "" || strings.EqualFold(c.System, "Editorial") {
				sbj.Creator = c.System
			}

			setRels(c, o, &sbj.Rels)

			sbj.ParentIds.AddString(o.ParentId)
			if o.TopParent {
				sbj.TopParent = true
			}

			subjects[i] = sbj
		}
	}

	return subjects
}

func setSubjects(field string, subjects []subject) *json.Property {
	if len(subjects) > 0 {
		ja := json.Array{}
		for _, sbj := range subjects {
			subject := json.Object{}
			subject.AddString("name", sbj.Name)
			subject.AddString("scheme", "http://cv.ap.org/id/")
			subject.AddString("code", sbj.Code)
			if sbj.Creator != "" {
				subject.AddString("creator", sbj.Creator)
			}
			if !sbj.Rels.IsEmpty() {
				subject.AddProperty(sbj.Rels.ToJsonProperty("rels"))
			}
			if !sbj.ParentIds.IsEmpty() {
				subject.AddProperty(sbj.ParentIds.ToJsonProperty("parentids"))
			}
			subject.AddBool("topparent", sbj.TopParent)
			ja.AddObject(&subject)
		}
		return json.NewArrayProperty(field, &ja)
	}

	return nil
}

func getPersons(c Classification, keys map[string]int, persons []person) []person {
	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			var p person

			key := fmt.Sprintf("%s_%s", o.Id, o.Value)
			i, ok := keys[key]
			if ok {
				p = persons[i]
			} else {
				p = person{Name: o.Value, Code: o.Id, Creator: c.System, Rels: &uniqueArray{}, Types: uniqueArray{}, Teams: uniqueArray{}, States: uniqueArray{}, Events: uniqueArray{}}
				persons = append(persons, p)
				i = len(persons) - 1
				keys[key] = i
			}

			for _, prop := range o.Property {
				if prop.Name != "" && prop.Value != "" {
					name := strings.ToLower(prop.Name)
					if name == "partytype" {
						if strings.EqualFold(prop.Value, "PERSON_FEATURED") {
							p.Rels.AddString(prop.Value)
							if p.Creator == "" {
								p.Creator = "Editorial"
							}
						} else {
							p.Types.AddString(prop.Value)
						}
					} else if name == "team" && prop.Id != "" {
						p.Teams.AddKeyValue("code", prop.Id, "name", prop.Value)
					} else if name == "associatedevent" && prop.Id != "" {
						p.Events.AddKeyValue("code", prop.Id, "name", prop.Value)
					} else if name == "associatedstate" && prop.Id != "" {
						p.States.AddKeyValue("code", prop.Id, "name", prop.Value)
					} else if name == "extid" {
						p.Ids.AddString(prop.Value)
					}
				}
			}

			if p.Creator == "" || strings.EqualFold(c.System, "Editorial") {
				p.Creator = c.System
			}

			setRels(c, o, p.Rels)

			persons[i] = p
		}
	}

	return persons
}

func setPersons(persons []person) *json.Property {
	if len(persons) > 0 {
		ja := json.Array{}
		for _, p := range persons {
			person := json.Object{}
			person.AddString("name", p.Name)
			person.AddString("scheme", "http://cv.ap.org/id/")
			person.AddString("code", p.Code)
			if p.Creator != "" {
				person.AddString("creator", p.Creator)
			}
			if !p.Rels.IsEmpty() {
				person.AddProperty(p.Rels.ToJsonProperty("rels"))
			}
			if !p.Types.IsEmpty() {
				person.AddProperty(p.Types.ToJsonProperty("types"))
			}
			if !p.Teams.IsEmpty() {
				person.AddProperty(p.Teams.ToJsonProperty("teams"))
			}
			if !p.States.IsEmpty() {
				person.AddProperty(p.States.ToJsonProperty("associatedstates"))
			}
			if !p.Events.IsEmpty() {
				person.AddProperty(p.Events.ToJsonProperty("associatedevents"))
			}
			if !p.Ids.IsEmpty() {
				person.AddProperty(p.Ids.ToJsonProperty("extids"))
			}
			ja.AddObject(&person)
			fmt.Printf("%p\t%p\n", &person, &p.Rels)
		}
		return json.NewArrayProperty("persons", &ja)
	}

	return nil
}

func setRels(c Classification, o Occurrence, rels *uniqueArray) {
	if strings.EqualFold(c.System, "RTE") {
		rels.AddString("inferred")
	} else if o.ActualMatch {
		rels.AddString("direct")
	} else {
		rels.AddString("ancestor")
	}
}
