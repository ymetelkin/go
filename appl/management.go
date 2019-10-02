package appl

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ymetelkin/go/xml"
)

const (
	digit0 rune = 48
)

func (doc *Document) parsePublicationManagement(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	var (
		embargo              bool
		types, outs, signals uniqueStrings
		asswith              []xml.Node
	)

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "RecordType":
			doc.RecordType = nd.Text
		case "FilingType":
			doc.FilingType = nd.Text
		case "ChangeEvent":
			doc.ChangeEvent = nd.Text
		case "ItemKey":
			doc.ItemKey = nd.Text
		case "ArrivalDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ArrivalDateTime = &ts
				}
			}
		case "FirstCreated":
			doc.Created = parseFirstCreated(nd)
			if doc.Created != nil && doc.Created.Year > 0 {
				doc.Copyright = &Copyright{
					Year: doc.Created.Year,
				}
			}
		case "LastModifiedDateTime":
			doc.Modified = parseLastModified(nd)
		case "Status":
			status := getPubStatus(nd.Text)
			if status != "" {
				doc.Status = status
			}
		case "ReleaseDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ReleaseDateTime = &ts
				}
			}
		case "AssociatedWith":
			asswith = append(asswith, nd)
		case "RefersTo":
			if doc.RefersTo == "" {
				doc.RefersTo = nd.Text
			}
		case "Instruction":
			outs.Append(nd.Text)
		case "SpecialInstructions":
			doc.SpecialInstructions = nd.Text
		case "Editorial":
			if nd.Nodes != nil {
				for _, n := range nd.Nodes {
					if n.Name == "Type" && n.Text != "" {
						types.Append(n.Text)

						if !embargo {
							if strings.EqualFold(n.Text, "Advance") || strings.EqualFold(n.Text, "HoldForRelease") {
								embargo = true
							}
						}
					}
				}
			}
		case "EditorialId":
			doc.EditorialID = nd.Text
		case "ItemStartDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemStartDateTime = &ts
				}
			}
		case "ItemStartDateTimeActual":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemStartDateTimeActual = &ts
				}
			}
		case "ItemExpireDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemExpireDateTime = &ts
				}
			}
		case "SearchDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.SearchDateTime = &ts
				}
			}
		case "ItemEndDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemEndDateTime = &ts
				}
			}
		case "Function":
			doc.Function = nd.Text
		case "TimeRestrictions":
			if nd.Nodes != nil {
				for _, n := range nd.Nodes {
					tr := parseTimeRestriction(n)
					doc.TimeRestrictions = append(doc.TimeRestrictions, tr)
				}
			}
		case "ExplicitWarning":
			if nd.Text == "1" {
				signals.Append("explicitcontent")
			} else if strings.EqualFold(nd.Text, "NUDITY") {
				signals.Append("NUDITY")
			} else if strings.EqualFold(nd.Text, "OBSCENITY") {
				signals.Append("OBSCENITY")
			} else if strings.EqualFold(nd.Text, "GRAPHIC CONTENT") {
				signals.Append("GRAPHICCONTENT")
			}
		case "IsDigitized":
			if strings.EqualFold(nd.Text, "false") {
				signals.Append("isnotdigitized")
			}
		}
	}

	if embargo && doc.ReleaseDateTime != nil {
		doc.Embargoed = doc.ReleaseDateTime
	}

	if !types.IsEmpty() {
		doc.EditorialTypes = types.Values()
	}

	if !outs.IsEmpty() {
		doc.Outs = outs.Values()
	}

	if !signals.IsEmpty() {
		doc.Signals = signals.Values()
	}

	if len(asswith) > 0 {
		doc.Associations = parseAssociations(asswith)
	}

	if doc.ItemEndDateTime != nil && (doc.ItemExpireDateTime == nil || (*doc.ItemEndDateTime).Before(*doc.ItemExpireDateTime)) {
		doc.ItemExpireDateTime = doc.ItemEndDateTime
	}
}

func getPubStatus(s string) string {
	if strings.EqualFold(s, "Usable") || strings.EqualFold(s, "Embargoed") || strings.EqualFold(s, "Unknown") {
		return "usable"
	} else if strings.EqualFold(s, "Withheld") {
		return "withheld"
	} else if strings.EqualFold(s, "Canceled") {
		return "canceled"
	}
	return ""
}

func parseFirstCreated(nd xml.Node) *FirstCreated {
	if nd.Attributes == nil {
		return nil
	}

	var fc FirstCreated

	for _, a := range nd.Attributes {
		switch a.Name {
		case "Year":
			year, err := strconv.Atoi(a.Value)
			if err == nil && year > 0 {
				fc.Year = year
			}
		case "Month":
			month, err := strconv.Atoi(a.Value)
			if err == nil && month > 0 {
				fc.Month = month
			}
		case "Day":
			day, err := strconv.Atoi(a.Value)
			if err == nil && day > 0 {
				fc.Day = day
			}
		case "Time":
			ts, err := parseTime(a.Value)
			if err == nil {
				fc.Time = a.Value
				fc.Hour = ts.Hour()
				fc.Minute = ts.Minute()
				fc.Second = ts.Second()
			}
		}
	}

	if fc.Year > 0 {
		ts := time.Date(fc.Year, time.Month(fc.Month), fc.Day, fc.Hour, fc.Minute, fc.Second, 0, time.UTC)
		fc.Date = &ts
	}

	ua := parseUserAccount(nd)
	fc.User = &ua

	return &fc
}

func parseLastModified(nd xml.Node) *LastModified {
	if nd.Text == "" {
		return nil
	}

	var lm LastModified

	ts, err := parseDate(nd.Text)
	if err == nil {
		lm.Date = &ts
	}

	if nd.Attributes != nil {
		ua := parseUserAccount(nd)
		lm.User = &ua
	}

	return &lm
}

func parseUserAccount(nd xml.Node) (ua UserAccount) {
	for _, a := range nd.Attributes {
		switch a.Name {
		case "UserName":
			ua.Name = a.Value
		case "UserAccount":
			ua.Account = a.Value
		case "UserAccountSystem":
			ua.System = a.Value
		case "ToolVersion":
			ua.ToolVersion = a.Value
		case "UserWorkgroup":
			ua.Workgroup = a.Value
		case "UserLocation":
			ua.Location = a.Value
		}
	}
	return
}

func parseTimeRestriction(nd xml.Node) (tr TimeRestriction) {
	if nd.Attributes != nil {
		for _, a := range nd.Attributes {
			switch a.Name {
			case "System":
				tr.System = a.Value
			case "Zone":
				tr.Zone = a.Value
			case "Include":
				tr.Include = a.Value == "true"
			}
		}

		if tr.System != "" && tr.Zone != "" {
			tr.ID = strings.ToLower(fmt.Sprintf("%s%s", tr.System, tr.Zone))
		}
	}

	return
}

func parseAssociations(nodes []xml.Node) (asses []Association) {
	/*
	   -test the value of //AssociatedWith, if its all zeros, do not convert; otherwise, each //AssociatedWith is converted to an object $.associations[i];
	   -each object $.associations[i] has five name/value pairs, $.associations[i].type, $.associations[i].itemid, $.associations[i].representationtype, $.associations[i].associationrank and $associations[i].typerank;
	   --retrieve the value from AssociatedWith/@CompositionType, use “Appendix III: CompositionType/Type Lookup Table” to derive the value for $.associations.association{i}.type;
	   --test the value of AssociatedWith, if it’s not all zeros, load as is to  $.associations[i].itemid;
	   --hardcode “partial” for $.associations[i].representationtype;
	   --load the sequence number of the AssociatedWith node (a number starting at 1) to $.associations[i].associationrank as a number;
	   --load the sequence number of the AssociatedWith node by @CompositionType (a number starting at 1) to $.associations[i].typerank as a number; note that CompositionType may be absent OR ‘StandardIngestedContent’ (which does not output a type) and any such AssociatedWith nodes should be ranked on their own.
	*/

	rank := 0
	types := make(map[string]int)

	for _, nd := range nodes {
		runes := []rune(nd.Text)
		empty := true
		for _, r := range runes {
			if r != digit0 {
				empty = false
				break
			}
		}

		if empty || nd.Attributes == nil {
			continue
		}

		ct := nd.Attribute("CompositionType")
		if strings.EqualFold(ct, "StandardText") {
			ct = "text"
		} else if strings.EqualFold(ct, "StandardPrintPhoto") {
			ct = "photo"
		} else if strings.EqualFold(ct, "StandardOnlinePhoto") {
			ct = "photo"
		} else if strings.EqualFold(ct, "StandardPrintGraphic") {
			ct = "graphic"
		} else if strings.EqualFold(ct, "StandardOnlineGraphic") {
			ct = "graphic"
		} else if strings.EqualFold(ct, "StandardBroadcastVideo") {
			ct = "video"
		} else if strings.EqualFold(ct, "StandardOnlineVideo") {
			ct = "video"
		} else if strings.EqualFold(ct, "StandardBroadcastAudio") {
			ct = "audio"
		} else if strings.EqualFold(ct, "StandardOnlineAudio") {
			ct = "audio"
		} else if strings.EqualFold(ct, "StandardLibraryVideo") {
			ct = "video"
		} else if strings.EqualFold(ct, "StandardInteractive") {
			ct = "complexdata"
		} else if strings.EqualFold(ct, "StandardBroadcastGraphic") {
			ct = "graphic"
		} else if strings.EqualFold(ct, "StandardBroadcastPhoto") {
			ct = "photo"
		} else if strings.EqualFold(ct, "StandardIngestedContent") {
			ct = "notype"
		} else {
			continue
		}

		var ass Association

		if ct != "notype" {
			ass.Type = ct
		}

		ass.ItemID = nd.Text

		rank++
		ass.Rank = rank

		typerank, ok := types[ct]
		if ok {
			typerank++
		} else {
			typerank = 1
		}
		types[ct] = typerank
		ass.TypeRank = typerank

		asses = append(asses, ass)
	}

	return
}
