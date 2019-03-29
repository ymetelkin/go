package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (desc *DescriptiveMetadata) parse(aj *ApplJson) error {
	getDescriptions(aj)
	getClassification(aj)

	return nil
}

func getDescriptions(aj *ApplJson) {
	descs := aj.Xml.DescriptiveMetadata.Description
	if descs != nil {
		descriptions := UniqueStrings{}
		for _, desc := range descs {
			descriptions.Add(desc)
		}
		aj.Descriptions = descriptions.ToJsonProperty("descriptions")
	}
}

func getClassification(aj *ApplJson) {
	generators := []ApplGenerator{}
	generators_keys := make(map[string]bool)
	categories := make(map[string]string)
	suppcategories := make(map[string]string)
	subjects := make(map[string]ApplSubject)
	orgs := make(map[string]ApplSubject)

	classification := aj.Xml.DescriptiveMetadata.SubjectClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			authority := strings.ToLower(c.Authority)

			if authority != "" && c.AuthorityVersion != "" {
				key := fmt.Sprintf("%s%s", c.Authority, c.AuthorityVersion)
				_, ok := generators_keys[key]
				if !ok {
					generators_keys[key] = true
					generator := ApplGenerator{Name: c.Authority, Version: c.AuthorityVersion}
					generators = append(generators, generator)
				}
			}

			if c.Occurrence != nil {
				if authority == "ap subject" {
					getSubjects(c, subjects)
				} else if authority == "ap category code" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
							categories[o.Id] = o.Value
						}
					}
				} else if authority == "ap supplemental category code" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
							suppcategories[o.Id] = o.Value
						}
					}
				} else if authority == "ap alert category" {
					for _, o := range c.Occurrence {
						if o.Id != "" {
							aj.AlertCategories.Add(o.Id)
						}
					}
				} else if aj.Fixture == nil && authority == "ap audio cut number code" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
							i, err := strconv.Atoi(o.Id)
							if err == nil && i >= 900 {
								fixture := json.JsonObject{}
								fixture.AddInt("code", i)
								fixture.AddString("name", o.Value)
								aj.Fixture = &json.JsonProperty{Field: "fixture", Value: &json.JsonObjectValue{Value: fixture}}
								break
							}
						}
					}
				}
			}

		}
	}

	classification = aj.Xml.DescriptiveMetadata.EntityClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			authority := strings.ToLower(c.Authority)
			if authority != "" && c.AuthorityVersion != "" {
				key := fmt.Sprintf("%s%s", c.Authority, c.AuthorityVersion)
				_, ok := generators_keys[key]
				if !ok {
					generators_keys[key] = true
					generator := ApplGenerator{Name: c.Authority, Version: c.AuthorityVersion}
					generators = append(generators, generator)
				}
			}

			if c.Occurrence != nil {
				if authority == "ap party" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
						}
					}
				} else if authority == "ap organization" {
					getSubjects(c, orgs)
				} else if authority == "ap company" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
						}
					}
				} else if authority == "ap geography" || authority == "ap country" || authority == "ap region" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
						}
					}
				} else if authority == "ap event" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
						}
					}
				}
			}
		}
	}

	aj.Categories = categories
	aj.SuppCategories = suppcategories

	if len(subjects) > 0 {
		list := []ApplSubject{}
		for _, subject := range subjects {
			list = append(list, subject)
		}
		aj.Subjects = list
	}

	if len(orgs) > 0 {
		list := []ApplSubject{}
		for _, subject := range orgs {
			list = append(list, subject)
		}
		aj.Organizations = list
	}

	if len(generators) > 0 {
		aj.Generators = generators
	}
}

func getSubjects(c Classification, subjects map[string]ApplSubject) {
	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			key := fmt.Sprintf("%s%s", o.Id, o.Value)
			subject, ok := subjects[key]
			if !ok {
				subject = ApplSubject{Name: o.Value, Code: o.Id}
			}
			if subject.Creator == "" || strings.EqualFold(c.System, "Editorial") {
				subject.Creator = c.System
			}
			if strings.EqualFold(c.System, "RTE") {
				subject.Rels.Add("inferred")
			} else if o.ActualMatch {
				subject.Rels.Add("direct")
			} else {
				subject.Rels.Add("ancestor")
			}
			if o.ParentId != "" {
				subject.ParentIds.Add(o.ParentId)
			}
			if o.TopParent {
				subject.TopParent = true
			}

			subjects[key] = subject
		}
	}
}
