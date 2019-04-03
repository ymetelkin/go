package appl

import "github.com/ymetelkin/go/json"

func (rights *RightsMetadata) parse(doc *document) error {
	getUsageRights(doc)

	return nil
}

func getUsageRights(doc *document) {
	urs := doc.Xml.RightsMetadata.UsageRights
	if urs == nil || len(urs) == 0 {
		return
	}

	usagerights := json.Array{}

	for _, ur := range urs {
		usageright := json.Object{}
		if ur.UsageType != "" {
			usageright.AddString("usagetype", ur.UsageType)
		}
		if ur.Geography != nil {
			geography := uniqueArray{}
			for _, g := range ur.Geography {
				geography.AddString(g)
			}
			usageright.AddProperty(geography.ToJsonProperty("geography"))
		}
		if ur.RightsHolder != "" {
			usageright.AddString("rightsholder", ur.RightsHolder)
		}
		if ur.Limitations != nil {
			limitations := uniqueArray{}
			for _, lim := range ur.Limitations {
				limitations.AddString(lim)
			}
			usageright.AddProperty(limitations.ToJsonProperty("limitations"))
		}
		if ur.StartDate != "" {
			usageright.AddString("startdate", ur.StartDate)
		}
		if ur.EndDate != "" {
			usageright.AddString("enddate", ur.EndDate)
		}
		if ur.Group != nil {
			groups := json.Array{}
			for _, g := range ur.Group {
				group := json.Object{}
				if g.Type != "" {
					group.AddString("type", g.Type)
				}
				if g.Id != "" {
					group.AddString("code", g.Id)
				}
				if g.Value != "" {
					group.AddString("name", g.Value)
				}
				if !group.IsEmpty() {
					groups.AddObject(&group)
				}
			}
			if !groups.IsEmpty() {
				usageright.AddArray("groups", &groups)
			}
		}

		if !usageright.IsEmpty() {
			usagerights.AddObject(&usageright)
		}
	}

	if !usagerights.IsEmpty() {
		doc.UsageRights = json.NewArrayProperty("usagerights", &usagerights)
	}
}
