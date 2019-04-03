package appl

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (pub *Publication) parseFilings(doc *document) error {
	if pub.FilingMetadata == nil {
		return nil
	}

	fs := []filing{}
	ja := json.Array{}

	for _, fm := range pub.FilingMetadata {
		tmp := fm
		jo := json.Object{}
		f := filing{}

		if tmp.Id == "" {
			return errors.New("[FilingMetadata.Id] is missing")
		} else {
			jo.AddString("filingid", tmp.Id)
		}
		if tmp.ArrivalDateTime == "" {
			return errors.New("[FilingMetadata.ArrivalDateTime] is missing")
		} else {
			jo.AddString("filingarrivaldatetime", tmp.ArrivalDateTime+"Z")
		}
		if tmp.Cycle != "" {
			jo.AddString("cycle", tmp.Cycle)
		}
		if tmp.TransmissionReference != "" {
			jo.AddString("transmissionreference", tmp.TransmissionReference)
		}
		if tmp.TransmissionFilename != "" {
			jo.AddString("transmissionfilename", tmp.TransmissionFilename)
		}
		if tmp.TransmissionContent != "" {
			jo.AddString("transmissioncontent", tmp.TransmissionContent)
		}
		if tmp.ServiceLevelDesignator != "" {
			jo.AddString("serviceleveldesignator", tmp.ServiceLevelDesignator)
		}
		if tmp.Selector != "" {
			jo.AddString("selector", tmp.Selector)
		}
		if tmp.Format != "" {
			jo.AddString("format", tmp.Format)
		}
		if tmp.Source != "" {
			jo.AddString("filingsource", tmp.Source)
		}
		if tmp.Category != "" {
			jo.AddString("filingcategory", tmp.Category)
		}
		if tmp.Routing != nil {
			routings := json.Object{}

			for _, r := range tmp.Routing {
				if r.Value != "" && r.Type != "" {
					values := uniqueArray{}
					tokens := strings.Split(r.Value, " ")
					for _, s := range tokens {
						values.AddString(s)
					}

					expanded := ""
					if r.Expanded {
						expanded = "expanded"
					}

					out := "add"
					if r.Outed {
						out = "out"
					}

					field := fmt.Sprintf("%s%s%ss", expanded, strings.ToLower(r.Type), out)
					routings.AddProperty(values.ToJsonProperty(field))
				}
			}

			if !routings.IsEmpty() {
				jo.AddObject("routings", &routings)
			}
		}
		if tmp.SlugLine != "" {
			jo.AddString("slugline", tmp.SlugLine)
			f.SlugLine = tmp.SlugLine
		}
		if tmp.OriginalMediaId != "" {
			jo.AddString("originalmediaid", tmp.OriginalMediaId)
		}
		if tmp.ImportFolder != "" {
			jo.AddString("importfolder", tmp.ImportFolder)
		}
		if tmp.ImportWarnings != "" {
			jo.AddString("importwarnings", tmp.ImportWarnings)
		}
		if tmp.LibraryTwinCheck != "" {
			jo.AddString("librarytwincheck", tmp.LibraryTwinCheck)
		}
		if tmp.LibraryRequestId != "" {
			jo.AddString("libraryrequestid", tmp.LibraryRequestId)
		}
		if tmp.SpecialFieldAttn != "" {
			jo.AddString("specialfieldattn", tmp.SpecialFieldAttn)
		}
		if tmp.FeedLine != "" {
			jo.AddString("feedline", tmp.FeedLine)
		}
		if tmp.LibraryRequestLogin != "" {
			jo.AddString("libraryrequestlogin", tmp.LibraryRequestLogin)
		}
		if tmp.Products.Product != nil {
			products := json.Array{}
			for _, p := range tmp.Products.Product {
				products.AddInt(p)
			}
			jo.AddArray("products", &products)
		}
		if tmp.PriorityLine != "" {
			jo.AddString("priorityline", tmp.PriorityLine)
		}
		if tmp.ForeignKeys != nil {
			foreignkeys := json.Array{}
			for _, fk := range tmp.ForeignKeys {
				if fk.System != "" && fk.Keys != nil {
					for _, k := range fk.Keys {
						if k.Id != "" && k.Field != "" {
							field := fk.System + k.Field
							field = strings.ReplaceAll(field, " ", "")
							field = strings.ToLower(field)
							field = html.EscapeString(field)
							foreignkey := json.Object{}
							foreignkey.AddString(field, k.Id)
							foreignkeys.AddObject(&foreignkey)

							if f.ForeignKeys == nil {
								f.ForeignKeys = make(map[string]string)
							}
							f.ForeignKeys[field] = k.Id
						}
					}
				}
			}

			if !foreignkeys.IsEmpty() {
				jo.AddArray("foreignkeys", &foreignkeys)
			}
		}
		if tmp.FilingCountry != nil && len(tmp.FilingCountry) > 0 {
			ua := uniqueArray{}
			for _, s := range tmp.FilingCountry {
				ua.AddString(s)
			}
			jo.AddProperty(ua.ToJsonProperty("filingcountries"))
		}
		if tmp.FilingRegion != nil && len(tmp.FilingRegion) > 0 {
			ua := uniqueArray{}
			for _, s := range tmp.FilingRegion {
				ua.AddString(s)
			}
			jo.AddProperty(ua.ToJsonProperty("filingregions"))
		}

		filingsubjects := uniqueArray{}
		if tmp.FilingSubject != nil {
			for _, s := range tmp.FilingSubject {
				filingsubjects.AddString(s)
			}
		}
		if tmp.FilingSubSubject != nil {
			for _, s := range tmp.FilingSubSubject {
				filingsubjects.AddString(s)
			}
		}
		jo.AddProperty(filingsubjects.ToJsonProperty("filingsubjects"))

		if tmp.FilingTopic != nil && len(tmp.FilingTopic) > 0 {
			ua := uniqueArray{}
			for _, s := range tmp.FilingTopic {
				ua.AddString(s)
			}
			jo.AddProperty(ua.ToJsonProperty("filingtopics"))
		}
		if tmp.FilingOnlineCode != "" {
			jo.AddString("filingonlinecode", tmp.FilingOnlineCode)
		}
		if tmp.DistributionScope != "" {
			jo.AddString("distributionscope", tmp.DistributionScope)
		}
		if tmp.BreakingNews != "" {
			jo.AddString("breakingnews", tmp.BreakingNews)
		}
		if tmp.FilingStyle != "" {
			jo.AddString("filingstyle", tmp.FilingStyle)
		}
		if tmp.JunkLine != "" {
			jo.AddString("junkline", tmp.JunkLine)
		}

		ja.AddObject(&jo)
		fs = append(fs, f)
	}

	if !ja.IsEmpty() {
		doc.Filings = filings{Filings: fs, Json: json.NewArrayProperty("filings", &ja)}
	}

	return nil
}
