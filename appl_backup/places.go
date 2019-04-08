package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

type place struct {
	Name         string
	Code         string
	Creator      string
	Rels         uniqueArray
	ParentIds    uniqueArray
	TopParent    bool
	LocationType json.Property
	Geo          json.Property
}

type places struct {
	Keys   map[string]int
	Places []place
}

func (ps *places) Parse(c Classification) {
	if c.Occurrence == nil {
		return
	}

	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			var p place

			key := fmt.Sprintf("%s_%s", o.Id, o.Value)

			if ps.Keys == nil {
				ps.Keys = make(map[string]int)
				ps.Places = []place{}
			}

			i, ok := ps.Keys[key]
			if ok {
				p = ps.Places[i]
			} else {
				p = place{Name: o.Value, Code: o.Id, Creator: c.System}
				ps.Places = append(ps.Places, p)
				i = len(ps.Places) - 1
				ps.Keys[key] = i
			}

			if p.Creator == "" || strings.EqualFold(c.System, "Editorial") {
				p.Creator = c.System
			}

			setRels(c, o, &p.Rels)

			p.ParentIds.AddString(o.ParentId)
			p.TopParent = o.TopParent

			var (
				lat  float64
				long float64
			)

			for _, prop := range o.Property {
				if prop.Name != "" && prop.Value != "" {
					name := strings.ToLower(prop.Name)
					if name == "locationtype" && prop.Id != "" && p.LocationType.IsEmtpy() {
						jo := json.Object{}
						jo.AddString("code", prop.Id)
						jo.AddString("name", prop.Value)
						p.LocationType = json.NewObjectProperty("locationtype", &jo)
					} else if name == "centroidlatitude" && lat == 0 {
						f, err := strconv.ParseFloat(prop.Value, 64)
						if err == nil {
							lat = f
						}
					} else if name == "centroidlongitude" && long == 0 {
						f, err := strconv.ParseFloat(prop.Value, 64)
						if err == nil {
							long = f
						}
					}
				}
			}

			p.Geo = getGeoProperty(lat, long)

			ps.Places[i] = p
		}
	}
}

func (ps *places) ToJsonProperty() json.Property {
	if ps.Keys != nil {
		ja := json.Array{}
		for _, item := range ps.Places {
			p := item
			place := json.Object{}
			place.AddString("name", p.Name)
			place.AddString("scheme", "http://cv.ap.org/id/")
			place.AddString("code", p.Code)
			if p.Creator != "" {
				place.AddString("creator", p.Creator)
			}
			if !p.Rels.IsEmpty() {
				place.AddProperty(p.Rels.ToJsonProperty("rels"))
			}
			if !p.ParentIds.IsEmpty() {
				place.AddProperty(p.ParentIds.ToJsonProperty("parentids"))
			}
			if p.TopParent {
				place.AddBool("topparent", true)
			}
			place.AddProperty(p.LocationType)
			place.AddProperty(p.Geo)
			ja.AddObject(&place)
		}
		return json.NewArrayProperty("places", &ja)
	}

	return json.Property{}
}
