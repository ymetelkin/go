package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
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

func (ps *persons) Parse(nd xml.Node) {
	if nd.Nodes == nil {
		return
	}

	system := nd.GetAttribute("System")

	for _, n := range nd.Nodes {
		code, name := getOccurrenceCodeName(n)
		if code != "" && name != "" {
			var prs person

			key := fmt.Sprintf("%s_%s", code, name)

			if ps.Keys == nil {
				ps.Keys = make(map[string]int)
				ps.Persons = []person{}
			}

			i, ok := ps.Keys[key]
			if ok {
				prs = ps.Persons[i]
			} else {
				prs = person{Name: name, Code: code, Creator: system}
				prs.Rels.AddString("direct")
				ps.Persons = append(ps.Persons, prs)
				i = len(ps.Persons) - 1
				ps.Keys[key] = i
			}

			if n.Nodes != nil {
				for _, p := range n.Nodes {
					if p.Attributes != nil {
						var id, pid, n, v string
						for _, a := range p.Attributes {
							switch a.Name {
							case "Id":
								id = a.Value
							case "Name":
								n = a.Value
							case "Value":
								v = a.Value
							case "ParentId":
								pid = a.Value
							}
						}

						if n != "" && v != "" {
							key := strings.ToLower(n)
							if key == "partytype" {
								if strings.EqualFold(v, "PERSON_FEATURED") {
									prs.Rels.AddString(v)
									if prs.Creator == "" {
										prs.Creator = "Editorial"
									}
								} else {
									prs.Types.AddString(v)
								}
							} else if key == "team" && id != "" {
								prs.Teams.AddKeyValue("code", id, "name", v)
							} else if key == "associatedevent" && id != "" {
								prs.Events.AddKeyValue("code", id, "name", v)
							} else if key == "associatedstate" && id != "" {
								prs.States.AddKeyValue("code", id, "name", v)
							} else if key == "extid" {
								prs.Ids.AddString(v)
							}
						}
					}
				}
			}

			if prs.Creator == "" || strings.EqualFold(system, "Editorial") {
				prs.Creator = system
			}

			ps.Persons[i] = prs
		}
	}
}

func (ps *persons) ToJsonProperty() json.Property {
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

			ja.AddObject(person)
		}

		return json.NewArrayProperty("persons", ja)
	}

	return json.Property{}
}
