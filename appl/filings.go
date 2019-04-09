package appl

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type filing struct {
	Source      string
	Category    string
	Slugline    string
	ForeignKeys []foreignkey
	JSON        json.Object
}

type foreignkey struct {
	Field string
	Value string
}

func parseFiling(nd xml.Node) filing {
	if nd.Nodes == nil {
		return filing{}
	}

	var (
		f                  filing
		jo, rts            json.Object
		ids                json.Array
		fcs, frs, fss, fts uniqueArray
	)

	for _, n := range nd.Nodes {
		switch n.Name {
		case "Id":
			if n.Text != "" {
				jo.AddString("filingid", n.Text)
			}
		case "ArrivalDateTime":
			if n.Text != "" {
				jo.AddString("filingarrivaldatetime", n.Text)
			}
		case "Cycle":
			if n.Text != "" {
				jo.AddString("cycle", n.Text)
			}
		case "TransmissionReference":
			if n.Text != "" {
				jo.AddString("transmissionreference", n.Text)
			}
		case "TransmissionFilename":
			if n.Text != "" {
				jo.AddString("transmissionfilename", n.Text)
			}
		case "TransmissionContent":
			if n.Text != "" {
				jo.AddString("transmissioncontent", n.Text)
			}
		case "ServiceLevelDesignator":
			if n.Text != "" {
				jo.AddString("serviceleveldesignator", n.Text)
			}
		case "Selector":
			if n.Text != "" {
				jo.AddString("selector", n.Text)
			}
		case "Source":
			if n.Text != "" {
				f.Source = nd.Text
				jo.AddString("filingsource", n.Text)
			}
		case "Category":
			if n.Text != "" {
				f.Category = nd.Text
				jo.AddString("filingcategory", n.Text)
			}
		case "Routing":
			if n.Attributes != nil {
				var t, e, o string

				for _, a := range n.Attributes {
					switch a.Name {
					case "Type":
						t = a.Value
					case "Expanded":
						if a.Value == "true" {
							e = "expanded"
						}
					case "Outed":
						if a.Value == "true" {
							o = "out"
						} else {
							o = "add"
						}
					}
				}
				if nd.Text != "" && t != "" {
					ua := uniqueArray{}
					tokens := strings.Split(nd.Text, " ")
					for _, s := range tokens {
						ua.AddString(s)
					}

					field := fmt.Sprintf("%s%s%ss", e, strings.ToLower(t), o)
					rts.AddProperty(ua.ToJsonProperty(field))
				}
			}
		case "SlugLine":
			if n.Text != "" {
				f.Slugline = nd.Text
				jo.AddString("slugline", n.Text)
			}
		case "OriginalMediaId":
			if n.Text != "" {
				jo.AddString("originalmediaid", n.Text)
			}
		case "ImportFolder":
			if n.Text != "" {
				jo.AddString("importfolder", n.Text)
			}
		case "ImportWarnings":
			if n.Text != "" {
				jo.AddString("importwarnings", n.Text)
			}
		case "LibraryTwinCheck":
			if n.Text != "" {
				jo.AddString("librarytwincheck", n.Text)
			}
		case "LibraryRequestId":
			if n.Text != "" {
				jo.AddString("libraryrequestid", n.Text)
			}
		case "SpecialFieldAttn":
			if n.Text != "" {
				jo.AddString("specialfieldattn", n.Text)
			}
		case "FeedLine":
			if n.Text != "" {
				jo.AddString("feedline", n.Text)
			}
		case "LibraryRequestLogin":
			if n.Text != "" {
				jo.AddString("libraryrequestlogin", n.Text)
			}
		case "Products":
			if nd.Nodes != nil {
				for _, n := range nd.Nodes {
					i, err := strconv.Atoi(nd.Text)
					if err == nil {
						ids.AddInt(i)
					}
				}
			}
		case "PriorityLine":
			if n.Text != "" {
				jo.AddString("priorityline", n.Text)
			}
		case "ForeignKeys":
			system := nd.GetAttribute("System")
			if system != "" && nd.Nodes != nil {
				for _, n := range nd.Nodes {
					if n.Attributes != nil {
						var id, fld string
						for _, a := range n.Attributes {
							switch a.Name {
							case "Id":
								id = a.Value
							case "Field":
								fld = a.Value
							}
						}
						if id != "" && fld != "" {
							field := system + fld
							field = strings.ReplaceAll(field, " ", "")
							field = strings.ToLower(field)
							field = html.EscapeString(field)
							fk := foreignkey{Field: field, Value: id}
							if f.ForeignKeys == nil {
								f.ForeignKeys = []foreignkey{fk}
							} else {
								f.ForeignKeys = append(f.ForeignKeys, fk)
							}
						}
					}
				}
			}
		case "FilingCountry":
			if nd.Text != "" {
				fcs.AddString(nd.Text)
			}
		case "FilingRegion":
			if nd.Text != "" {
				frs.AddString(nd.Text)
			}
		case "FilingSubject", "FilingSubSubject":
			if nd.Text != "" {
				fss.AddString(nd.Text)
			}
		case "FilingTopic":
			if nd.Text != "" {
				fts.AddString(nd.Text)
			}
		case "FilingOnlineCode":
			if n.Text != "" {
				jo.AddString("filingonlinecode", n.Text)
			}
		case "DistributionScope":
			if n.Text != "" {
				jo.AddString("distributionscope", n.Text)
			}
		case "BreakingNews":
			if n.Text != "" {
				jo.AddString("breakingnews", n.Text)
			}
		case "FilingStyle":
			if n.Text != "" {
				jo.AddString("filingstyle", n.Text)
			}
		case "JunkLine":
			if n.Text != "" {
				jo.AddString("junkline", n.Text)
			}
		}
	}

	if !ids.IsEmpty() {
		jo.AddArray("products", ids)
	}

	if f.ForeignKeys != nil {
		ja := json.Array{}
		for _, fk := range f.ForeignKeys {
			k := json.Object{}
			k.AddString(fk.Field, fk.Value)
			ja.AddObject(k)
		}
		jo.AddArray("foreignkeys", ja)
	}

	if !rts.IsEmpty() {
		jo.AddObject("routings", rts)
	}

	if !fcs.IsEmpty() {
		jo.AddProperty(fcs.ToJsonProperty("filingcountries"))
	}

	if !frs.IsEmpty() {
		jo.AddProperty(fcs.ToJsonProperty("filingregions"))
	}

	if !fss.IsEmpty() {
		jo.AddProperty(fss.ToJsonProperty("filingsubjects"))
	}

	if !fts.IsEmpty() {
		jo.AddProperty(fts.ToJsonProperty("filingtopics"))
	}

	f.JSON = jo

	return f
}
