package appl

import (
	"errors"

	"github.com/ymetelkin/go/json"
)

func (doc *document) ParseRightsMetadata(parent *json.Object) error {

	if doc.RightsMetadata.Nodes == nil {
		return errors.New("RightsMetadata is missing")
	}

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
						ut = nd.Text
					case "Geography":
						geo.AddString(nd.Text)
					case "RightsHolder":
						rh = nd.Text
					case "Limitations":
						lim.AddString(nd.Text)
					case "StartDate":
						sd = nd.Text
					case "EndDate":
						ed = nd.Text
					case "Group":
						g := json.Object{}
						if nd.Attributes != nil {
							for _, a := range nd.Attributes {
								switch a.Name {
								case "Type":
									if a.Value != "" {
										g.AddString("type", a.Value)
									}
								case "Id":
									if a.Value != "" {
										g.AddString("code", a.Value)
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

				jo.AddProperty(geo.ToJsonProperty("geography"))

				if rh != "" {
					jo.AddString("rightsholder", rh)
				}

				jo.AddProperty(lim.ToJsonProperty("limitations"))

				if sd != "" {
					jo.AddString("startdate", sd)
				}

				if ed != "" {
					jo.AddString("enddate", ed)
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
