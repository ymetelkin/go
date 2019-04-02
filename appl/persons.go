package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

type person struct {
	Name    string
	Code    string
	Creator string
	Rels    uniqueArray
	Types   uniqueArray
	Ids     uniqueArray
	Teams   uniqueArray
	States  uniqueArray
	Events  uniqueArray
}

type persons struct {
	Keys    map[string]int
	Persons []person
}

func (ps *persons) Parse(c Classification) {
	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			var p person

			key := fmt.Sprintf("%s_%s", o.Id, o.Value)

			if ps.Keys == nil {
				ps.Keys = make(map[string]int)
				ps.Persons = []person{}
			}

			i, ok := ps.Keys[key]
			if ok {
				p = ps.Persons[i]
			} else {
				p = person{Name: o.Value, Code: o.Id, Creator: c.System}
				p.Rels.AddString("direct")
				ps.Persons = append(ps.Persons, p)
				i = len(ps.Persons) - 1
				ps.Keys[key] = i
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

			ps.Persons[i] = p
		}
	}
}

func (ps *persons) ToJsonProperty() *json.Property {
	if ps.Keys != nil {
		ja := json.Array{}
		for _, item := range ps.Persons {
			p := item
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
		}

		return json.NewArrayProperty("persons", &ja)
	}

	return nil
}
