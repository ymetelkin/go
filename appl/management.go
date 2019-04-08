package appl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

const (
	DIGIT_ZERO rune = 48
)

func (doc *document) ParsePublicationManagement(jo *json.Object) error {
	if doc.PublicationManagement == nil {
		return errors.New("PublicationManagement is missing")
	}

	var (
		embargo, ref bool
		rdt          string
		types, outs  uniqueArray
		ass          []Node
	)

	for _, nd := range doc.PublicationManagement.Nodes {
		switch nd.Name {
		case "RecordType":
			jo.AddString("recordtype", nd.Text)
		case "FilingType":
			jo.AddString("filingtype", nd.Text)
		case "ChangeEvent":
			if nd.Text != "" {
				jo.AddString("changeevent", nd.Text)
			}
		case "ArrivalDateTime":
			if nd.Text != "" {
				jo.AddString("arrivaldatetime", nd.Text+"Z")
			}
		case "FirstCreated":
			date, year, err := getFirstCreatedDate(nd)
			if err == nil {
				doc.FirstCreatedYear = year
				jo.SetString("firstcreated", date)
			} else {
				return err
			}
		case "LastModifiedDateTime":
			if nd.Text != "" {
				jo.AddString("lastmodifieddatetime", nd.Text+"Z")
			}
		case "Status":
			ps, err := getPubStatus(nd.Text)
			if err == nil {
				doc.PubStatus = ps
				jo.SetString("pubstatus", string(ps))
			} else {
				return err
			}
		case "Instruction":
			if nd.Text != "" {
				if outs == nil {
					outs = uniqueArray{}
				}
				outs.AddString(nd.Text)
			}
		case "Editorial":
			n := nd.GetNode("Type")
			s := n.Text
			if s != "" {
				if types == nil {
					types = uniqueArray{}
				}
				types.AddString(s)

				if !embargo {
					if strings.EqualFold(s, "Advance") || strings.EqualFold(s, "HoldForRelease") {
						embargo = true
					}
				}
			}
		case "AssociatedWith":
			if ass == nil {
				ass = []Node{nd}
			} else {
				ass = append(ass, nd)
			}
		case "ReleaseDateTime":
			if nd.Text != "" {
				rdct = nd.Text + "Z"
				jo.AddString("releasedatetime", rdt)
			}
		case "SpecialInstructions":
			if nd.Test != "" {
				jo.AddString("specialinstructions", nd.Text)
			}
		case "ItemStartDateTime":
			if nd.Test != "" {
				jo.AddString("editorialid", nd.Text)
			}
		case "LastModifiedDateTime":
			if nd.Test != "" {
				jo.AddString("itemstartdatetime", nd.Text+"Z")
			}
		case "ItemStartDateTimeActual":
			if nd.Test != "" {
				jo.AddString("itemstartdatetimeactual", nd.Text+"Z")
			}
		case "ItemExpireDateTime":
			if nd.Test != "" {
				jo.AddString("itemexpiredatetime", nd.Text+"Z")
			}
		case "SearchDateTime":
			if nd.Test != "" {
				jo.AddString("searchdatetime", nd.Text+"Z")
			}
		case "ItemEndDateTime":
			if nd.Test != "" {
				jo.AddString("itemenddatetime", nd.Text+"Z")
			}
		case "Function":
			if nd.Test != "" {
				doc.Function = nd.Text
				jo.AddString("function", nd.Text)
			}
		case "TimeRestrictions":
			if nd.Nodes != nil {
				for _, n := range nd.Nodes {
					f, v := getTimeRestriction(n)
					if f != "" {
						jo.AddBool(f, v)
					}
				}
			}
		case "RefersTo":
			if !refersto && nd.Text != "" {
				jo.AddString("refersto", nd.Text)
				refersto = true
			}
		case "ExplicitWarning":
			if nd.Text == "1" {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString("explicitcontent")
			} else if strings.EqualFold(nd.Text, "NUDITY") {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString("NUDITY")
			} else if strings.EqualFold(nd.Text, "OBSCENITY") {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString("OBSCENITY")
			} else if strings.EqualFold(nd.Text, "GRAPHIC CONTENT") {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString("GRAPHICCONTENT")
			}
		case "IsDigitized":
			if strings.EqualFold(nd.Text, "false") {
				if doc.Signals == nil {
					doc.Signals = uniqueArray{}
				}
				doc.Signals.AddString("isnotdigitized")
			}
		}
	}

	if embargo && rdt != "" {
		jo.AddString("embargoed", rdt)
	}

	if outs != nil {
		jo.AddProperty(outs.ToJsonProperty("outinginstructions"))
	}

	if types != nil {
		jo.AddProperty(types.ToJsonProperty("editorialtypes"))
	}

	jo.AddProperty(doc.Signals.ToJsonProperty("signals"))

	getAssociatedWith(ass, &jo)
}

func getPubStatus(s string) (PubStatus, error) {
	if strings.EqualFold(s, "Usable") || strings.EqualFold(s, "Embargoed") || strings.EqualFold(s, "Unknown") {
		return PUBSTATUS_USABLE, nil
	} else if strings.EqualFold(s, "Withheld") {
		return PUBSTATUS_WITHHELD, nil
	} else if strings.EqualFold(s, "Canceled") {
		return PUBSTATUS_CANCELED, nil
	} else {
		e := fmt.Sprintf("Invalid pub status [%s]", status)
		return PUBSTATUS_UNKNOWN, errors.New(e)
	}
}

func getFirstCreatedDate(nd Node) (string, int, error) {
	if nd.Attributes == nil {
		return errors.New("FirstCreated year is missing")
	}

	var (
		year                   int
		date, month, day, time string
	)

	for _, a := range nd.Attributes {
		switch a.Name {
		case "Year":
			i, err := strconv.Atoi(a.Value)
			if err == nil && i > 0 {
				year = i
			}
		case "Month":
			i, err := strconv.Atoi(a.Value)
			if err == nil {
				zero := ""
				if i < 10 {
					zero = "0"
				}
				month = fmt.Sprintf("%s%d", zero, i)
			}
		case "Day":
			i, err := strconv.Atoi(a.Value)
			if err == nil {
				zero := ""
				if i < 10 {
					zero = "0"
				}
				day = fmt.Sprintf("%s%d", zero, i)
			}
		case "Time":
			time = nd.Text
		}
	}

	if year <= 0 {
		e := fmt.Sprintf("Invalid FirstCreated year [%d]", doc.FirstCreatedYear)
		return errors.New(e)
	}

	if month == "" {
		date = fmt.Sprintf("%d", year)
	} else {
		if day == "" {
			date = fmt.Sprintf("%d-%s", year, month)
		} else {
			if time == "" {
				date = fmt.Sprintf("%d-%s-%s", year, month, day)
			} else {
				date = fmt.Sprintf("%d-%s-%sT%sZ", year, month, day, fc.Time)
			}
		}
	}

	return date, year, nil
}

func getTimeRestriction(nd Node) (string, bool) {
	if nd.Attributes != nil {
		var (
			system, zone, include string
		)
		for _, a := range nd.Attributes {
			switch a.Name {
			case "System":
				system = a.Value
			case "Zone":
				zone = a.Value
			case "Include":
				include = a.Value
			}

			if sys != "" && zone != "" {
				s := fmt.Sprintf("%s%s", system, zone)
				return strings.ToLower(s), include == "true"
			}
		}

		return "", false
	}
}

func getAssociatedWith(ass []Node, parent *json.Object) {
	/*
	   -test the value of //AssociatedWith, if its all zeros, do not convert; otherwise, each //AssociatedWith is converted to an object $.associations[i];
	   -each object $.associations[i] has five name/value pairs, $.associations[i].type, $.associations[i].itemid, $.associations[i].representationtype, $.associations[i].associationrank and $associations[i].typerank;
	   --retrieve the value from AssociatedWith/@CompositionType, use “Appendix III: CompositionType/Type Lookup Table” to derive the value for $.associations.association{i}.type;
	   --test the value of AssociatedWith, if it’s not all zeros, load as is to  $.associations[i].itemid;
	   --hardcode “partial” for $.associations[i].representationtype;
	   --load the sequence number of the AssociatedWith node (a number starting at 1) to $.associations[i].associationrank as a number;
	   --load the sequence number of the AssociatedWith node by @CompositionType (a number starting at 1) to $.associations[i].typerank as a number; note that CompositionType may be absent OR ‘StandardIngestedContent’ (which does not output a type) and any such AssociatedWith nodes should be ranked on their own.
	*/
	ja := json.Array{}

	rank := 0
	types := make(map[string]int)

	for _, aw := range ass {
		runes := []rune(aw.Text)
		empty := true
		for _, r := range runes {
			if r != DIGIT_ZERO {
				empty = false
				break
			}
		}

		if empty || aw.Attributes == nil {
			continue
		}

		ct := aw.GetAttribute("CompositionType")
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

		jo := json.Object{}

		if ct != "notype" {
			jo.AddString("type", t)
		}

		jo.AddString("itemid", aw.Text)
		jo.AddString("representationtype", "partial")

		rank++
		jo.AddInt("associationrank", rank)

		typerank, ok := types[t]
		if ok {
			typerank++
		} else {
			typerank = 1
		}
		types[t] = typerank
		jo.AddInt("typerank", typerank)

		ja.AddObject(jo)
	}

	if ja.Length() > 0 {
		parent.AddArray("associations", ja)
	}
}
