package appl

import (
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (desc *DescriptiveMetadata) parse(doc *document) error {
	getDescriptions(doc)
	getDatelineLocation(doc)
	getClassification(doc)
	getThirdParty(doc)

	return nil
}

func getDescriptions(doc *document) {
	descs := doc.Xml.DescriptiveMetadata.Description
	if descs != nil {
		descriptions := uniqueArray{}
		for _, desc := range descs {
			descriptions.AddString(desc)
		}
		doc.Descriptions = descriptions.ToJsonProperty("descriptions")
	}
}

func getDatelineLocation(doc *document) {
	dll := doc.Xml.DescriptiveMetadata.DateLineLocation

	datelinelocation := json.Object{}
	if dll.City != "" {
		datelinelocation.AddString("city", dll.City)
	}
	if dll.CountryArea != "" {
		datelinelocation.AddString("countryareacode", dll.CountryArea)
	}
	if dll.CountryAreaName != "" {
		datelinelocation.AddString("countryareaname", dll.CountryAreaName)
	}
	if dll.Country != "" {
		datelinelocation.AddString("countrycode", dll.Country)
	}
	if dll.CountryName != "" {
		datelinelocation.AddString("countryname", dll.CountryName)
	}

	datelinelocation.AddProperty(getGeoProperty(dll.LatitudeDD, dll.LongitudeDD))

	if !datelinelocation.IsEmpty() {
		doc.DatelineLocation = json.NewObjectProperty("datelinelocation", &datelinelocation)
	}
}

func getClassification(doc *document) {
	generators := uniqueArray{}
	categories := uniqueArray{}
	suppcategories := uniqueArray{}
	alerts := uniqueArray{}
	sbjs := subjects{}
	orgs := subjects{}
	persons := persons{}
	companies := companies{}
	places := places{}
	events := events{}
	services := uniqueArray{}

	classification := doc.Xml.DescriptiveMetadata.SubjectClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			generators.AddKeyValue("name", c.Authority, "version", c.AuthorityVersion)

			authority := strings.ToLower(c.Authority)

			if c.Occurrence != nil {
				if authority == "ap subject" {
					sbjs.Parse(c)
				} else if authority == "ap category code" {
					for _, o := range c.Occurrence {
						categories.AddKeyValue("code", o.Id, "name", o.Value)
					}
				} else if authority == "ap supplemental category code" {
					for _, o := range c.Occurrence {
						suppcategories.AddKeyValue("code", o.Id, "name", o.Value)
					}
				} else if authority == "ap alert category" {
					for _, o := range c.Occurrence {
						alerts.AddString(o.Id)
					}
				} else if doc.Fixture == nil && authority == "ap audio cut number code" {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
							i, err := strconv.Atoi(o.Id)
							if err == nil && i >= 900 {
								fixture := json.Object{}
								fixture.AddInt("code", i)
								fixture.AddString("name", o.Value)
								doc.Fixture = json.NewObjectProperty("fixture", &fixture)
								break
							}
						}
					}
				}
			}

		}
	}

	classification = doc.Xml.DescriptiveMetadata.EntityClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			generators.AddKeyValue("name", c.Authority, "version", c.AuthorityVersion)

			authority := strings.ToLower(c.Authority)

			if c.Occurrence != nil {
				if authority == "ap party" {
					persons.Parse(c)
				} else if authority == "ap organization" {
					orgs.Parse(c)
				} else if authority == "ap company" {
					companies.Parse(c)
				} else if authority == "ap geography" || authority == "ap country" || authority == "ap region" {
					places.Parse(c)
				} else if authority == "ap event" {
					events.Parse(c)
				}
			}
		}
	}

	getAudences(doc)

	classification = doc.Xml.DescriptiveMetadata.SalesClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			if c.Occurrence != nil {
				for _, o := range c.Occurrence {
					if o.Id != "" && o.Value != "" {
						service := json.Object{}
						service.AddString("code", o.Id)
						service.AddString("name", o.Value)
						services.AddObject(o.Id, &service)
					}
				}
			}
		}
	}

	comments := doc.Xml.DescriptiveMetadata.Comment
	if comments != nil && len(comments) > 0 {
		for _, c := range comments {
			if c != "" {
				service := json.Object{}
				service.AddString("apservice", c)
				services.AddObject(c, &service)
			}
		}
	}

	doc.Generators = generators.ToJsonProperty("generators")
	doc.Categories = categories.ToJsonProperty("categories")
	doc.SuppCategories = suppcategories.ToJsonProperty("suppcategories")
	doc.AlertCategories = alerts.ToJsonProperty("alertcategories")

	doc.Subjects = sbjs.ToJsonProperty("subjects")
	doc.Persons = persons.ToJsonProperty()
	doc.Organizations = orgs.ToJsonProperty("organizations")
	doc.Companies = companies.ToJsonProperty()
	doc.Places = places.ToJsonProperty()
	doc.Events = events.ToJsonProperty()
	doc.Services = services.ToJsonProperty("services")
}

func getThirdParty(doc *document) {
	tpms := doc.Xml.DescriptiveMetadata.ThirdPartyMeta
	if tpms != nil && len(tpms) > 0 {
		thirdpartymeta := json.Array{}

		for _, tpm := range tpms {
			jo := json.Object{}
			if tpm.System != "" {
				jo.AddString("creator", tpm.System)
			}
			if tpm.Vocabulary != "" {
				jo.AddString("vocabulary", tpm.Vocabulary)
			}
			if tpm.VocabularyOwner != "" {
				jo.AddString("vocabularyowner", tpm.VocabularyOwner)
			}
			if tpm.Occurrence != nil && len(tpm.Occurrence) > 0 {
				o := tpm.Occurrence[0]
				if o.Id != "" {
					jo.AddString("code", o.Id)
				}
				if o.Value != "" {
					jo.AddString("name", o.Value)
				}
			}

			if !jo.IsEmpty() {
				thirdpartymeta.AddObject(&jo)
			}
		}

		if !thirdpartymeta.IsEmpty() {
			doc.ThirdPartyMeta = json.NewArrayProperty("thirdpartymeta", &thirdpartymeta)
		}
	}
}

func getAudences(doc *document) {
	geo := false
	audiences := uniqueArray{}

	classification := doc.Xml.DescriptiveMetadata.AudienceClassification
	if classification != nil && len(classification) > 0 {
		for _, c := range classification {
			if strings.EqualFold(c.Authority, "AP Audience") && strings.EqualFold(c.System, "Editorial") {
				if c.Occurrence != nil {
					for _, o := range c.Occurrence {
						if o.Id != "" && o.Value != "" {
							key := o.Id
							audience := json.Object{}
							audience.AddString("code", o.Id)
							audience.AddString("name", o.Value)

							if o.Property != nil && len(o.Property) > 0 {
								prop := o.Property[0]
								if prop.Value != "" {
									if strings.EqualFold(prop.Value, "AUDGEOGRAPHY") {
										geo = true
									}
									audience.AddString("type", prop.Value)
								}
							}

							audiences.AddObject(key, &audience)
						}
					}
				}
			}
		}
	}

	if !geo && doc.Xml.FilingMetadata != nil {
		for _, f := range doc.Xml.FilingMetadata {
			if strings.EqualFold(f.Category, "n") {
				state := getState(f.Source)
				if state != nil {
					audiences.AddObject(state.Code, state.ToJson())
				}
			}
		}
	}

	if !audiences.IsEmpty() {
		doc.Audiences = audiences.ToJsonProperty("audiences")
	}
}
