package appl

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/xml"
)

func (doc *Document) parseFilingMetadata(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	var (
		f              Filing
		routings       map[string][]string
		cs, rs, ss, ts uniqueStrings
	)

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "Id":
			f.ID = nd.Text
		case "ArrivalDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					f.ArrivalDateTime = &ts
				}
			}
		case "Cycle":
			f.Cycle = nd.Text
		case "TransmissionReference":
			f.TransmissionReference = nd.Text
		case "TransmissionFilename":
			f.TransmissionFilename = nd.Text
		case "TransmissionContent":
			f.TransmissionContent = nd.Text
		case "ServiceLevelDesignator":
			f.ServiceLevelDesignator = nd.Text
		case "Selector":
			f.Selector = nd.Text
		case "Format":
			f.Format = nd.Text
		case "Source":
			f.Source = nd.Text
		case "Category":
			f.Category = nd.Text
		case "Routing":
			if nd.Attributes != nil {
				var t, e, o string

				for k, v := range nd.Attributes {
					switch k {
					case "Type":
						t = v
					case "Expanded":
						if v == "true" {
							e = "expanded"
						}
					case "Outed":
						if v == "true" {
							o = "out"
						} else {
							o = "add"
						}
					}
				}
				if nd.Text != "" && t != "" {
					var us uniqueStrings
					toks := strings.Split(nd.Text, " ")
					for _, s := range toks {
						us.Append(s)
					}

					field := fmt.Sprintf("%s%s%ss", e, strings.ToLower(t), o)

					if routings == nil {
						routings = make(map[string][]string)
					}
					routings[field] = us.Values()
				}
			}
		case "SlugLine":
			f.Slugline = nd.Text
		case "OriginalMediaId":
			f.OriginalMediaID = nd.Text
		case "ImportFolder":
			f.ImportFolder = nd.Text
		case "ImportWarnings":
			f.ImportWarnings = nd.Text
		case "LibraryTwinCheck":
			f.LibraryTwinCheck = nd.Text
		case "LibraryRequestId":
			f.LibraryRequestID = nd.Text
		case "SpecialFieldAttn":
			f.SpecialFieldAttn = nd.Text
		case "FeedLine":
			f.Feedline = nd.Text
		case "LibraryRequestLogin":
			f.LibraryRequestLogin = nd.Text
		case "Products":
			if nd.Nodes != nil {
				for _, p := range nd.Nodes {
					p, err := strconv.Atoi(p.Text)
					if err == nil {
						f.Products = append(f.Products, p)
					}
				}
			}
		case "PriorityLine":
			f.Priorityline = nd.Text
		case "ForeignKeys":
			fks := parseForeignKeys(nd)
			if fks != nil {
				f.ForeignKeys = append(f.ForeignKeys, fks...)
			}
		case "FilingCountry":
			cs.Append(nd.Text)
		case "FilingRegion":
			rs.Append(nd.Text)
		case "FilingSubject", "FilingSubSubject":
			ss.Append(nd.Text)
		case "FilingTopic":
			ts.Append(nd.Text)
		case "FilingOnlineCode":
			f.OnlineCode = nd.Text
		case "DistributionScope":
			f.DistributionScope = nd.Text
		case "BreakingNews":
			f.BreakingNews = nd.Text
		case "FilingStyle":
			f.Style = nd.Text
		case "JunkLine":
			f.Junkline = nd.Text
		}
	}

	f.Routings = routings

	if !cs.IsEmpty() {
		f.Countries = cs.Values()
	}
	if !rs.IsEmpty() {
		f.Regions = rs.Values()
	}
	if !ss.IsEmpty() {
		f.Subjects = cs.Values()
	}
	if !ts.IsEmpty() {
		f.Topics = ts.Values()
	}

	doc.Filings = append(doc.Filings, f)

	if strings.EqualFold(f.Category, "n") {
		state := getState(f.Source)
		if state.Code != "" {
			if doc.Audiences != nil {
				for _, au := range doc.Audiences {
					if au.Title == "AUDGEOGRAPHY" || au.Code == state.Code {
						return
					}
				}
			}
			cnt := CodeNameTitle{
				Code:  state.Code,
				Name:  state.Name,
				Title: "AUDGEOGRAPHY",
			}
			doc.Audiences = append(doc.Audiences, cnt)
		}
	}
}

func parseForeignKeys(nd xml.Node) (fks []ForeignKey) {
	system := nd.Attribute("System")
	if system != "" && nd.Nodes != nil {
		for _, k := range nd.Nodes {
			if k.Attributes != nil {
				var id, fld string
				for k, v := range k.Attributes {
					switch k {
					case "Id":
						id = v
					case "Field":
						fld = v
					}
				}
				if id != "" && fld != "" {
					field := system + fld
					field = strings.ReplaceAll(field, " ", "")
					field = strings.ToLower(field)
					field = html.EscapeString(field)
					fk := ForeignKey{
						Field: field,
						Value: id,
					}
					fks = append(fks, fk)
				}
			}
		}
	}
	return fks
}
