package appl

import (
	"fmt"
	"strconv"

	"github.com/ymetelkin/go/xml"
)

func (doc *Document) parseRightsMetadata(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "Copyright":
			if nd.Attributes != nil {
				for k, v := range nd.Attributes {
					switch k {
					case "Holder":
						if doc.Copyright == nil {
							doc.Copyright = &Copyright{}
						}
						doc.Copyright.Holder = v
					case "Date":
						year, err := strconv.Atoi(v)
						if err == nil && year > 0 {
							if doc.Copyright == nil {
								doc.Copyright = &Copyright{}
							}
							doc.Copyright.Year = year
						}
					}
				}
			}
			if doc.Copyright != nil && doc.Copyright.Notice == "" && doc.Copyright.Holder != "" && doc.Copyright.Year > 0 {
				doc.Copyright.Notice = fmt.Sprintf("Copyright %d %s. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.", doc.Copyright.Year, doc.Copyright.Holder)
			}
		case "UsageRights":
			if nd.Nodes != nil {
				var (
					ur       UsageRights
					geo, lim uniqueStrings
				)

				for _, n := range nd.Nodes {
					switch n.Name {
					case "UsageType":
						ur.UsageType = n.Text
					case "Geography":
						geo.Append(n.Text)
					case "RightsHolder":
						ur.RightsHolder = n.Text
					case "Limitations":
						lim.Append(n.Text)
					case "StartDate":
						if n.Text != "" {
							ts, err := parseDate(n.Text)
							if err == nil {
								ur.StartDate = &ts
							}
						}
					case "EndDate":
						if n.Text != "" {
							ts, err := parseDate(n.Text)
							if err == nil {
								ur.EndDate = &ts
							}
						}
					case "Group":
						var grp CodeNameTitle
						if n.Attributes != nil {
							for k, v := range n.Attributes {
								switch k {
								case "Type":
									grp.Title = v
								case "Id":
									grp.Code = v
								}
							}
						}

						grp.Name = n.Text
						ur.Groups = append(ur.Groups, grp)
					}
				}
				if !geo.IsEmpty() {
					ur.Geography = geo.Values()
				}
				if !lim.IsEmpty() {
					ur.Limitations = lim.Values()
				}
				doc.UsageRights = append(doc.UsageRights, ur)
			}
		}
	}
}
