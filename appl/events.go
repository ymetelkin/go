package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type events struct {
	Keys   map[string]bool
	Events json.Array
}

func (es *events) Parse(nd xml.Node) {
	if nd.Nodes == nil {
		return
	}

	system := nd.GetAttribute("System")

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
}

func (es *events) ToJsonProperty() json.Property {
	if es.Keys != nil {
		return json.NewArrayProperty("events", es.Events)
	}

	return json.Property{}
}
