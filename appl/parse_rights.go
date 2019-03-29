package appl

import "github.com/ymetelkin/go/json"

func (rights *RightsMetadata) parse(aj *ApplJson) error {
	getUsageRights(aj)

	return nil
}

func getUsageRights(aj *ApplJson) {
	urs := aj.Xml.RightsMetadata.UsageRights
	if urs == nil || len(urs) == 0 {
		return
	}

	usagerights := json.JsonArray{}

	for _, ur := range urs {
		usageright := json.JsonObject{}
		if ur.UsageType != "" {
			usageright.AddString("usagetype", ur.UsageType)
		}
		if ur.Geography != nil {
			geography := UniqueStrings{}
			for _, g := range ur.Geography {
				geography.Add(g)
			}
			usageright.AddProperty(geography.ToJsonProperty("geography"))
		}
		if ur.RightsHolder != "" {
			usageright.AddString("rightsholder", ur.RightsHolder)
		}
		if ur.Limitations != nil {
			limitations := UniqueStrings{}
			for _, lim := range ur.Limitations {
				limitations.Add(lim)
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
			groups := json.JsonArray{}
			for _, g := range ur.Group {
				group := json.JsonObject{}
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
		aj.UsageRights = &json.JsonProperty{Field: "usagerights", Value: &json.JsonArrayValue{Value: usagerights}}
	}
}
