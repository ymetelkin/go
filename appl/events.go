package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

type event struct {
	Name    string
	Code    string
	Creator string
}

type events struct {
	Keys   map[string]bool
	Events json.Array
}

func (es *events) Parse(c Classification) {
	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			key := fmt.Sprintf("%s_%s", o.Id, o.Value)

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

			epix := strings.EqualFold("epix", c.System)
			extid := ""
			extidsource := ""
			eventproperties := json.Object{}

			for _, prop := range o.Property {
				if prop.Name != "" && prop.Value != "" {
					name := strings.ToLower(prop.Name)
					value := strings.ToLower(prop.Value)
					if name == "extid" {
						extid = value
					} else if name == "extidsource" {
						extidsource = value
						if !epix {
							epix = value == "nfl" || value == "sportradar"
						}
					} else {
						eventproperties.AddString(name, prop.Value)
					}
				}
			}

			e := json.Object{}

			if epix {
				externaleventids := json.Array{}

				id1 := json.Object{}
				id1.AddString("code", o.Id)
				id1.AddString("creator", "sportradar")
				id1.AddString("creatorcode", "sportradar:"+o.Id)
				externaleventids.AddObject(&id1)

				if extid != "" && extidsource != "" {
					id2 := json.Object{}
					id2.AddString("code", extid)
					id2.AddString("creator", extidsource)
					id2.AddString("creatorcode", fmt.Sprintf("%s:%s", extidsource, extid))
					externaleventids.AddObject(&id2)
				}

				e.AddString("name", o.Value)
				e.AddString("creator", "ePix")
				e.AddArray("externaleventids", &externaleventids)

			} else {
				e.AddString("code", o.Id)
				e.AddString("name", o.Value)
				if c.System != "" {
					e.AddString("creator", c.System)
				}
			}

			if !eventproperties.IsEmpty() {
				e.AddObject("eventproperties", &eventproperties)
			}

			if !e.IsEmpty() {
				es.Events.AddObject(&e)
			}
		}
	}
}

func (es *events) ToJsonProperty() *json.Property {
	if es.Keys != nil {
		return json.NewArrayProperty("events", &es.Events)
	}

	return nil
}
