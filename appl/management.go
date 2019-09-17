package appl

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ymetelkin/go/json"
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
		embargo, ref         bool
		rdt                  string
		types, outs, signals uniqueArray
		asswith              []xml.Node
	)

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "RecordType":
			doc.RecordType = nd.Text
			doc.JSON.AddString("recordtype", nd.Text)
		case "FilingType":
			doc.FilingType = nd.Text
			doc.JSON.AddString("filingtype", nd.Text)
		case "ChangeEvent":
			if nd.Text != "" {
				doc.ChangeEvent = nd.Text
				doc.JSON.AddString("changeevent", nd.Text)
			}
		case "ItemKey":
			if nd.Text != "" {
				doc.ItemKey = nd.Text
				doc.JSON.AddString("itemkey", nd.Text)
			}
		case "ArrivalDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ArrivalDateTime = &ts
					doc.JSON.AddString("arrivaldatetime", formatDate(ts))
				}
			}
		case "FirstCreated":
			fc, jo := newFirstCreated(nd)
			if fc.Date != nil {
				doc.Created = &fc
				doc.JSON.AddString("firstcreated", formatDate(*fc.Date))
			}
			if !jo.IsEmpty() {
				doc.JSON.AddObject("firstcreator", jo)
			}
		case "LastModifiedDateTime":
			lm, jo := newLastModified(nd)
			if lm.Date != nil {
				doc.Modified = &lm
				doc.JSON.AddString("lastmodifieddatetime", formatDate(*lm.Date))
			}
			if !jo.IsEmpty() {
				doc.JSON.AddObject("lastmodifier", jo)
			}
		case "Status":
			status := getPubStatus(nd.Text)
			if status != "" {
				doc.Status = status
				doc.JSON.AddString("pubstatus", status)
			}
		case "ReleaseDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ReleaseDateTime = &ts
					rdt = formatDate(ts)
					doc.JSON.AddString("releasedatetime", rdt)
				}
			}
		case "AssociatedWith":
			asswith = append(asswith, nd)
			doc.JSON.AddArray("associations", json.Array{})
		case "RefersTo":
			if !ref && nd.Text != "" {
				doc.RefersTo = nd.Text
				doc.JSON.AddString("refersto", nd.Text)
				ref = true
			}
		case "Instruction":
			outs.AddString(nd.Text)
		case "SpecialInstructions":
			if nd.Text != "" {
				doc.SpecialInstructions = nd.Text
				doc.JSON.AddString("specialinstructions", nd.Text)
			}
		case "Editorial":
			n := nd.Node("Type")
			s := n.Text
			if s != "" {
				types.AddString(s)

				if !embargo {
					if strings.EqualFold(s, "Advance") || strings.EqualFold(s, "HoldForRelease") {
						embargo = true
					}
				}
			}
		case "EditorialId":
			if nd.Text != "" {
				doc.EditorialID = nd.Text
				doc.JSON.AddString("editorialid", nd.Text)
			}
		case "ItemStartDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemStartDateTime = &ts
					doc.JSON.AddString("itemstartdatetime", formatDate(ts))
				}
			}
		case "ItemStartDateTimeActual":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemStartDateTimeActual = &ts
					doc.JSON.AddString("itemstartdatetimeactual", formatDate(ts))
				}
			}
		case "ItemExpireDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemExpireDateTime = &ts
					doc.JSON.AddString("itemexpiredatetime", formatDate(ts))
				}
			}
		case "SearchDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.SearchDateTime = &ts
					doc.JSON.AddString("searchdatetime", formatDate(ts))
				}
			}
		case "ItemEndDateTime":
			if nd.Text != "" {
				ts, err := parseDate(nd.Text)
				if err == nil {
					doc.ItemEndDateTime = &ts
					doc.JSON.AddString("itemenddatetime", formatDate(ts))
				}
			}
		case "Function":
			if nd.Text != "" {
				doc.Function = nd.Text
				doc.JSON.AddString("function", nd.Text)
			}
		case "TimeRestrictions":
			if nd.Nodes != nil {
				for _, n := range nd.Nodes {
					tr := newTimeRestriction(n)
					doc.TimeRestrictions = append(doc.TimeRestrictions, tr)
					if tr.ID != "" {
						doc.JSON.AddBool(tr.ID, tr.Include)
					}
				}
			}
		case "ExplicitWarning":
			if nd.Text == "1" {
				signals.AddString("explicitcontent")
			} else if strings.EqualFold(nd.Text, "NUDITY") {
				signals.AddString("NUDITY")
			} else if strings.EqualFold(nd.Text, "OBSCENITY") {
				signals.AddString("OBSCENITY")
			} else if strings.EqualFold(nd.Text, "GRAPHIC CONTENT") {
				signals.AddString("GRAPHICCONTENT")
			}
		case "IsDigitized":
			if strings.EqualFold(nd.Text, "false") {
				signals.AddString("isnotdigitized")
			}
		}
	}

	if embargo && doc.ReleaseDateTime != nil {
		doc.Embargoed = doc.ReleaseDateTime
		doc.JSON.AddString("embargoed", rdt)
	}

	if !types.IsEmpty() {
		doc.EditorialTypes = types.Values()
		doc.JSON.AddProperty(types.ToJSONProperty("editorialtypes"))
	}

	if !outs.IsEmpty() {
		doc.Outs = outs.Values()
		doc.JSON.AddProperty(outs.ToJSONProperty("outinginstructions"))
	}

	if !signals.IsEmpty() {
		doc.Signals = signals.Values()
		doc.JSON.AddProperty(signals.ToJSONProperty("signals"))
	}

	if len(asswith) > 0 {
		asses, ja := getAssociations(asswith)
		if len(asses) > 0 {
			doc.Associations = asses
			doc.JSON.SetArray("associations", ja)
		}
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

func newFirstCreated(nd xml.Node) (fc FirstCreated, jo json.Object) {
	if nd.Attributes == nil {
		return
	}

	for k, v := range nd.Attributes {
		switch k {
		case "Year":
			year, err := strconv.Atoi(v)
			if err == nil && year > 0 {
				fc.Year = year
			}
		case "Month":
			month, err := strconv.Atoi(v)
			if err == nil && month > 0 {
				fc.Month = month
			}
		case "Day":
			day, err := strconv.Atoi(v)
			if err == nil && day > 0 {
				fc.Day = day
			}
		case "Time":
			ts, err := parseTime(v)
			if err == nil {
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

	ua, jo := newUserAccount(nd)
	fc.User = &ua

	return
}

func newLastModified(nd xml.Node) (lm LastModified, jo json.Object) {
	if nd.Attributes == nil {
		return
	}

	ts, err := parseDate(nd.Text)
	if err == nil {
		lm.Date = &ts
	}

	ua, jo := newUserAccount(nd)
	lm.User = &ua

	return
}

func newUserAccount(nd xml.Node) (ua UserAccount, jo json.Object) {
	for k, v := range nd.Attributes {
		switch k {
		case "UserName":
			ua.Name = v
		case "UserAccount":
			ua.Account = v
		case "UserAccountSystem":
			ua.System = v
		case "ToolVersion":
			ua.ToolVersion = v
		case "UserWorkgroup":
			ua.Workgroup = v
		case "UserLocation":
			ua.Location = v
		}
	}

	if ua.Name != "" {
		jo.AddString("username", ua.Name)
	}
	if ua.Account != "" {
		jo.AddString("useraccount", ua.Account)
	}
	if ua.System != "" {
		jo.AddString("useraccountsystem", ua.System)
	}
	if ua.ToolVersion != "" {
		jo.AddString("toolversion", ua.ToolVersion)
	}
	if ua.Workgroup != "" {
		jo.AddString("userworkgroup", ua.Workgroup)
	}
	if ua.Location != "" {
		jo.AddString("userlocation", ua.Location)
	}

	return
}

func newTimeRestriction(nd xml.Node) (tr TimeRestriction) {
	if nd.Attributes != nil {
		var (
			system, zone string
		)
		for k, v := range nd.Attributes {
			switch k {
			case "System":
				tr.System = v
			case "Zone":
				tr.Zone = v
			case "Include":
				tr.Include = v == "true"
			}
		}

		if system != "" && zone != "" {
			tr.ID = strings.ToLower(fmt.Sprintf("%s%s", system, zone))
		}
	}

	return
}

func getAssociations(nodes []xml.Node) (asses []Association, ja json.Array) {
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

		var (
			ass Association
			jo  json.Object
		)

		if ct != "notype" {
			ass.Type = ct
			jo.AddString("type", ct)
		}

		ass.ItemID = nd.Text
		jo.AddString("itemid", nd.Text)
		jo.AddString("representationtype", "partial")

		rank++
		ass.Rank = rank
		jo.AddInt("associationrank", rank)

		typerank, ok := types[ct]
		if ok {
			typerank++
		} else {
			typerank = 1
		}
		types[ct] = typerank
		ass.TypeRank = typerank
		jo.AddInt("typerank", typerank)

		asses = append(asses, ass)
		ja.AddObject(jo)
	}

	return
}
