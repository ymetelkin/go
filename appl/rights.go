package appl

import (
	"github.com/ymetelkin/go/json"
)

func (doc *document) ParseRightsMetadata(parent *json.Object) error {
	ja := json.Array{}

	for _, nd := range doc.RightsMetadata.Nodes {
		switch nd.Name {
		case "UsageRights":
			if nd.Nodes != nil {
				var (
					ut, rh, sd, ed string
					geo, lim       uniqueArray
					gps            json.Array
				)
				jo := json.Object{}

				for _, n := range nd.Nodes {
					switch n.Name {
					case "UsageType":
						ut = n.Text
					case "Geography":
						geo.AddString(n.Text)
					case "RightsHolder":
						rh = n.Text
					case "Limitations":
						lim.AddString(n.Text)
					case "StartDate":
						sd = n.Text
					case "EndDate":
						ed = n.Text
					case "Group":
						g := json.Object{}
						if n.Attributes != nil {
							for k, v := range n.Attributes {
								switch k {
								case "Type":
									if v != "" {
										g.AddString("type", v)
									}
								case "Id":
									if v != "" {
										g.AddString("code", v)
									}
								}
							}
						}
						if n.Text != "" {
							g.AddString("name", n.Text)
						}
						if !g.IsEmpty() {
							gps.AddObject(g)
						}
					}
				}

				if ut != "" {
					jo.AddString("usagetype", ut)
				}

				jo.AddProperty(geo.ToJSONProperty("geography"))

				if rh != "" {
					jo.AddString("rightsholder", rh)
				}

				jo.AddProperty(lim.ToJSONProperty("limitations"))

				if sd != "" {
					jo.AddString("startdate", parseIsoDate(sd))
				}

				if ed != "" {
					jo.AddString("enddate", parseIsoDate(ed))
				}

				if !gps.IsEmpty() {
					jo.AddArray("groups", gps)
				}

				if !jo.IsEmpty() {
					ja.AddObject(jo)
				}
			}
		}
	}

	if !ja.IsEmpty() {
		parent.AddArray("usagerights", ja)
	}

	return nil
}
