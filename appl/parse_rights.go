package appl

func (rights *RightsMetadata) parse(aj *ApplJson) error {
	getUsageRights(aj)

	return nil
}

func getUsageRights(aj *ApplJson) {
	urs := aj.Xml.RightsMetadata.UsageRights
	if urs == nil || len(urs) == 0 {
		return
	}

	usagerights := []ApplUsageRights{}

	for _, ur := range urs {
		empty := true
		usageright := ApplUsageRights{}

		if ur.UsageType != "" {
			usageright.UsageType = ur.UsageType
			empty = false
		}

		if ur.RightsHolder != "" {
			usageright.RightsHolder = ur.RightsHolder
			empty = false
		}

		if ur.StartDate != "" {
			usageright.StartDate = ur.StartDate
			empty = false
		}

		if ur.EndDate != "" {
			usageright.EndDate = ur.EndDate
			empty = false
		}

		if ur.Geography != nil {
			for _, g := range ur.Geography {
				usageright.Geography.Add(g)
				empty = false
			}
		}

		if ur.Limitations != nil {
			for _, lim := range ur.Limitations {
				usageright.Limitations.Add(lim)
				empty = false
			}
		}

		if ur.Group != nil && len(ur.Group) > 0 {
			groups := []ApplGroup{}
			for _, g := range ur.Group {
				group := ApplGroup{}
				add := false
				if g.Value != "" {
					group.Name = g.Value
					add = true
				}
				if g.Id != "" {
					group.Code = g.Id
					add = true
				}
				if g.Type != "" {
					group.Type = g.Type
					add = true
				}

				if add {
					groups = append(groups, group)
				}
			}

			if len(groups) > 0 {
				usageright.Groups = groups
				empty = false
			}
		}

		if !empty {
			usagerights = append(usagerights, usageright)
		}
	}

	aj.UsageRights = usagerights
}
