package appl

/*

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
	Rels    uniqueStrings
	Types   uniqueStrings
	Ids     uniqueStrings
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

	system := nd.Attribute("System")

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
				prs.Rels.Append("direct")
				ps.Persons = append(ps.Persons, prs)
				i = len(ps.Persons) - 1
				ps.Keys[key] = i
			}

			if n.Nodes != nil {
				for _, p := range n.Nodes {
					if p.Attributes != nil {
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
									prs.Rels.Append(v)
									if prs.Creator == "" {
										prs.Creator = "Editorial"
									}
								} else {
									prs.Types.Append(v)
								}
							} else if key == "team" && id != "" {
								prs.Teams.AddKeyValue("code", id, "name", v)
							} else if key == "associatedevent" && id != "" {
								prs.Events.AddKeyValue("code", id, "name", v)
							} else if key == "associatedstate" && id != "" {
								prs.States.AddKeyValue("code", id, "name", v)
							} else if key == "extid" {
								prs.Ids.Append(v)
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

func (ps *persons) JSONProperty(namelines []json.Object) json.Property {
	ja := json.Array{}
	var add bool

	if ps.Keys != nil {
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
					person.AddProperty(p.Rels.JSONProperty("rels"))
				}
				if !p.Types.IsEmpty() {
					person.AddProperty(p.Types.JSONProperty("types"))
				}
			
			if !p.Teams.IsEmpty() {
				person.AddProperty(p.Teams.JSONProperty("teams"))
			}
			if !p.States.IsEmpty() {
				person.AddProperty(p.States.JSONProperty("associatedstates"))
			}
			if !p.Events.IsEmpty() {
				person.AddProperty(p.Events.JSONProperty("associatedevents"))
			}
			
				if !p.Ids.IsEmpty() {
					person.AddProperty(p.Ids.JSONProperty("extids"))
				}
			

			ja.AddObject(person)
			add = true
		}
	}

	if namelines != nil {
		for _, jo := range namelines {
			ja.AddObject(jo)
			add = true
		}
	}

	if add {
		return json.NewArrayProperty("persons", ja)
	}

	return json.Property{}
}
*/