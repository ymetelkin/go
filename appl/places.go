package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type place struct {
	Name         string
	Code         string
	Creator      string
	Rels         uniqueStrings
	ParentIds    uniqueStrings
	TopParent    string
	LocationType json.Property
	Geo          json.Property
}

type places struct {
	Keys   map[string]int
	Places []place
}

func (ps *places) Parse(nd xml.Node) {
	if nd.Nodes == nil {
		return
	}

	system := nd.Attribute("System")

	for _, n := range nd.Nodes {
		var code, name, /*match,*/ pid, tp string

		if n.Name == "Occurrence" && n.Attributes != nil {
			for k, v := range n.Attributes {
				switch k {
				case "Id":
					code = v
				case "Value":
					name = v
				case "ActualMatch":
					//match = v
				case "ParentId":
					pid = v
				case "TopParent":
					tp = v
				}
			}
		}

		if code != "" && name != "" {
			var plc place

			key := fmt.Sprintf("%s_%s", code, name)

			if ps.Keys == nil {
				ps.Keys = make(map[string]int)
				ps.Places = []place{}
			}

			i, ok := ps.Keys[key]
			if ok {
				plc = ps.Places[i]
			} else {
				plc = place{Name: name, Code: code, Creator: system}
				ps.Places = append(ps.Places, plc)
				i = len(ps.Places) - 1
				ps.Keys[key] = i
			}

			if plc.Creator == "" || strings.EqualFold(system, "Editorial") {
				plc.Creator = system
			}

			//setRels(system, match, &plc.Rels)

			plc.ParentIds.Append(pid)
			plc.TopParent = tp

			var (
				lat  float64
				long float64
			)

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
							if key == "locationtype" && id != "" && plc.LocationType.IsEmtpy() {
								jo := json.Object{}
								jo.AddString("code", id)
								jo.AddString("name", v)
								plc.LocationType = json.NewObjectProperty("locationtype", jo)
							} else if key == "centroidlatitude" && lat == 0 {
								f, err := strconv.ParseFloat(v, 64)
								if err == nil {
									lat = f
								}
							} else if key == "centroidlongitude" && long == 0 {
								f, err := strconv.ParseFloat(v, 64)
								if err == nil {
									long = f
								}
							}
						}
					}
				}
			}

			plc.Geo = getGeoProperty(lat, long)

			ps.Places[i] = plc
		}
	}
}

func (ps *places) JSONProperty() json.Property {
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
			/*
				if !p.Rels.IsEmpty() {
					place.AddProperty(p.Rels.JSONProperty("rels"))
				}
				if !p.ParentIds.IsEmpty() {
					place.AddProperty(p.ParentIds.JSONProperty("parentids"))
				}
			*/
			if p.TopParent == "true" {
				place.AddBool("topparent", true)
			} else if p.TopParent == "false" {
				place.AddBool("topparent", false)
			}
			place.AddProperty(p.LocationType)
			place.AddProperty(p.Geo)
			ja.AddObject(place)
		}
		return json.NewArrayProperty("places", ja)
	}

	return json.Property{}
}
