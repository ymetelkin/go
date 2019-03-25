package appl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

const (
	DIGIT_ZERO rune = 48
)

type PubStatus string

const (
	PUBSTATUS_USABLE   PubStatus = "usable"
	PUBSTATUS_WITHHELD PubStatus = "withheld"
	PUBSTATUS_CANCELED PubStatus = "canceled"
	PUBSTATUS_UNKNOWN  PubStatus = ""
)

type FirstCreated struct {
	Year  int    `xml:"Year,attr"`
	Month int    `xml:"Month,attr"`
	Day   int    `xml:"Day,attr"`
	Time  string `xml:"Time,attr"`
}

type TimeRestrictions struct {
	Restrictions []TimeRestriction `xml:"TimeRestriction"`
}

type TimeRestriction struct {
	System  string `xml:"System,attr"`
	Zone    string `xml:"Zone,attr"`
	Include bool   `xml:"Include,attr"`
}

type Editorial struct {
	Type string
}

type AssociatedWith struct {
	Value           string `xml:",chardata"`
	CompositionType string `xml:"CompositionType,attr"`
}

type XmlPublicationManagement struct {
	RecordType              string
	FilingType              string
	ChangeEvent             string
	ItemKey                 string
	ArrivalDateTime         string
	FirstCreated            FirstCreated
	LastModifiedDateTime    string
	Status                  string
	ReleaseDateTime         string
	AssociatedWith          []AssociatedWith
	RefersTo                []string
	Instructions            []string `xml:"Instruction"`
	SpecialInstructions     string
	EditorialTypes          []Editorial `xml:"Editorial"`
	EditorialId             string
	ItemStartDateTime       string
	ItemStartDateTimeActual string
	SearchDateTime          string
	ItemEndDateTime         string
	ItemExpireDateTime      string
	ExplicitWarning         string
	Function                string
	IsDigitized             string
	TimeRestrictions        TimeRestrictions

	pubStatus PubStatus
}

func (management *XmlPublicationManagement) ToJson(jo *json.JsonObject) error {
	if management.RecordType == "" {
		return errors.New("[PublicationManagement.RecordType] is missing")
	}
	if management.FilingType == "" {
		return errors.New("[PublicationManagement.FilingType] is missing")
	}

	ps, err := management.getPubStatus()
	if err != nil {
		return err
	}

	fc, err := management.FirstCreated.getFirstCreatedDate()
	if err != nil {
		return err
	}

	management.getSignals(jo)

	jo.AddString("recordtype", management.RecordType)
	jo.AddString("filingtype", management.FilingType)

	if management.ChangeEvent != "" {
		jo.AddString("changeevent", management.ChangeEvent)
	}
	if management.ItemKey != "" {
		jo.AddString("itemkey", management.ItemKey)
	}
	if management.ArrivalDateTime != "" {
		jo.AddString("arrivaldatetime", management.ArrivalDateTime+"Z")
	}

	jo.AddString("firstcreated", fc)

	if management.LastModifiedDateTime != "" {
		jo.AddString("lastmodifieddatetime", management.LastModifiedDateTime+"Z")
	}

	jo.AddString("pubstatus", string(ps))

	if management.ReleaseDateTime != "" {
		jo.AddString("releasedatetime", management.ReleaseDateTime+"Z")

	}

	management.getEditorialTypes(jo)
	management.getAssociatedWith(jo)

	if len(management.RefersTo) > 0 {
		jo.AddString("refersto", management.RefersTo[0])
	}

	management.getOutingInstructions(jo)

	if management.SpecialInstructions != "" {
		jo.AddString("specialinstructions", management.SpecialInstructions)
	}

	if management.ItemStartDateTime != "" {
		jo.AddString("editorialid", management.ItemStartDateTime)
	}
	if management.LastModifiedDateTime != "" {
		jo.AddString("itemstartdatetime", management.LastModifiedDateTime+"Z")
	}
	if management.ItemStartDateTimeActual != "" {
		jo.AddString("itemstartdatetimeactual", management.ItemStartDateTimeActual+"Z")
	}
	if management.ItemExpireDateTime != "" {
		jo.AddString("itemexpiredatetime", management.ItemExpireDateTime+"Z")
	}
	if management.SearchDateTime != "" {
		jo.AddString("searchdatetime", management.SearchDateTime+"Z")
	}
	if management.ItemEndDateTime != "" {
		jo.AddString("itemenddatetime", management.ItemEndDateTime+"Z")
	}
	if management.Function != "" {
		jo.AddString("function", management.Function)
	}

	management.getTimeRestrictions(jo)

	return nil
}

func (management *XmlPublicationManagement) getPubStatus() (PubStatus, error) {
	if management.pubStatus == "" {
		left := management.Status
		if strings.EqualFold(left, "Usable") || strings.EqualFold(left, "Embargoed") || strings.EqualFold(left, "Unknown") {
			management.pubStatus = PUBSTATUS_USABLE
		} else if strings.EqualFold(left, "Withheld") {
			management.pubStatus = PUBSTATUS_WITHHELD
		} else if strings.EqualFold(left, "Canceled") {
			management.pubStatus = PUBSTATUS_CANCELED
		} else {
			e := fmt.Sprintf("Invalid pub status [%s]", management.Status)
			return PUBSTATUS_UNKNOWN, errors.New(e)
		}
	}

	return management.pubStatus, nil
}

func (fc *FirstCreated) getFirstCreatedDate() (string, error) {
	if fc.Year <= 0 {
		e := fmt.Sprintf("Invalid year [%d]", fc.Year)
		return "", errors.New(e)
	}

	month := ""
	if fc.Month == 0 {
		s := fmt.Sprintf("%d", fc.Year)
		return s, nil
	} else {
		zero := ""
		if fc.Month < 10 {
			zero = "0"
		}
		month = fmt.Sprintf("%s%d", zero, fc.Month)
	}

	day := ""
	if fc.Day == 0 {
		s := fmt.Sprintf("%d-%s", fc.Year, month)
		return s, nil
	} else {
		zero := ""
		if fc.Day < 10 {
			zero = "0"
		}
		day = fmt.Sprintf("%s%d", zero, fc.Day)
	}

	if fc.Time == "" {
		s := fmt.Sprintf("%d-%s-%s", fc.Year, month, day)
		return s, nil
	} else {
		s := fmt.Sprintf("%d-%s-%sT%sZ", fc.Year, month, day, fc.Time)
		return s, nil
	}
}

func (management *XmlPublicationManagement) getSignals(jo *json.JsonObject) {
	signals := json.JsonArray{}

	if management.ExplicitWarning == "1" {
		signals.AddString("explicitcontent")
	} else if strings.EqualFold(management.ExplicitWarning, "NUDITY") {
		signals.AddString("NUDITY")
	} else if strings.EqualFold(management.ExplicitWarning, "OBSCENITY") {
		signals.AddString("OBSCENITY")
	} else if strings.EqualFold(management.ExplicitWarning, "GRAPHIC CONTENT") {
		signals.AddString("GRAPHICCONTENT")
	}

	if strings.EqualFold(management.IsDigitized, "false") {
		signals.AddString("isnotdigitized")
	}

	jo.AddArray("signals", &signals)
}

func (management *XmlPublicationManagement) getOutingInstructions(jo *json.JsonObject) {
	if len(management.Instructions) == 0 {
		return
	}

	unique := make(map[string]bool)
	for _, oi := range management.Instructions {
		unique[oi] = true
	}

	outinginstructions := json.JsonArray{}
	for key, _ := range unique {
		outinginstructions.Add(key)
	}

	jo.AddArray("outinginstructions", &outinginstructions)
}

func (management *XmlPublicationManagement) getEditorialTypes(jo *json.JsonObject) {
	if len(management.EditorialTypes) == 0 {
		return
	}

	embargoed := management.ReleaseDateTime != ""

	unique := make(map[string]bool)
	for _, et := range management.EditorialTypes {
		unique[et.Type] = true

		if embargoed {
			if strings.EqualFold(et.Type, "Advance") || strings.EqualFold(et.Type, "Advance") {
				embargoed = false
				jo.AddString("embargoed", management.ReleaseDateTime+"Z")
			}
		}
	}

	editorialtypes := json.JsonArray{}
	for key, _ := range unique {
		editorialtypes.Add(key)
	}

	jo.AddArray("editorialtypes", &editorialtypes)
}

func (management *XmlPublicationManagement) getTimeRestrictions(jo *json.JsonObject) {
	for _, tr := range management.TimeRestrictions.Restrictions {
		if tr.System != "" && tr.Zone != "" {
			name := fmt.Sprintf("%s%s", tr.System, tr.Zone)
			jo.AddBoolean(strings.ToLower(name), tr.Include)
		}
	}
}

func (management *XmlPublicationManagement) getAssociatedWith(jo *json.JsonObject) {
	/*
	   -test the value of //AssociatedWith, if its all zeros, do not convert; otherwise, each //AssociatedWith is converted to an object $.associations[i];
	   -each object $.associations[i] has five name/value pairs, $.associations[i].type, $.associations[i].itemid, $.associations[i].representationtype, $.associations[i].associationrank and $associations[i].typerank;
	   --retrieve the value from AssociatedWith/@CompositionType, use “Appendix III: CompositionType/Type Lookup Table” to derive the value for $.associations.association{i}.type;
	   --test the value of AssociatedWith, if it’s not all zeros, load as is to  $.associations[i].itemid;
	   --hardcode “partial” for $.associations[i].representationtype;
	   --load the sequence number of the AssociatedWith node (a number starting at 1) to $.associations[i].associationrank as a number;
	   --load the sequence number of the AssociatedWith node by @CompositionType (a number starting at 1) to $.associations[i].typerank as a number; note that CompositionType may be absent OR ‘StandardIngestedContent’ (which does not output a type) and any such AssociatedWith nodes should be ranked on their own.
	*/

	associations := json.JsonArray{}

	rank := 0
	types := make(map[string]int)

	for _, aw := range management.AssociatedWith {
		runes := []rune(aw.Value)
		empty := true
		for _, r := range runes {
			if r != DIGIT_ZERO {
				empty = false
				break
			}
		}

		if empty {
			continue
		}

		association := json.JsonObject{}

		t := ""

		if strings.EqualFold(aw.CompositionType, "StandardText") {
			t = "text"
		} else if strings.EqualFold(aw.CompositionType, "StandardPrintPhoto") {
			t = "photo"
		} else if strings.EqualFold(aw.CompositionType, "StandardOnlinePhoto") {
			t = "photo"
		} else if strings.EqualFold(aw.CompositionType, "StandardPrintGraphic") {
			t = "graphic"
		} else if strings.EqualFold(aw.CompositionType, "StandardOnlineGraphic") {
			t = "graphic"
		} else if strings.EqualFold(aw.CompositionType, "StandardBroadcastVideo") {
			t = "video"
		} else if strings.EqualFold(aw.CompositionType, "StandardOnlineVideo") {
			t = "video"
		} else if strings.EqualFold(aw.CompositionType, "StandardBroadcastAudio") {
			t = "audio"
		} else if strings.EqualFold(aw.CompositionType, "StandardOnlineAudio") {
			t = "audio"
		} else if strings.EqualFold(aw.CompositionType, "StandardLibraryVideo") {
			t = "video"
		} else if strings.EqualFold(aw.CompositionType, "StandardInteractive") {
			t = "complexdata"
		} else if strings.EqualFold(aw.CompositionType, "StandardBroadcastGraphic") {
			t = "graphic"
		} else if strings.EqualFold(aw.CompositionType, "StandardBroadcastPhoto") {
			t = "photo"
		} else if strings.EqualFold(aw.CompositionType, "StandardIngestedContent") {
			t = "notype"
		}

		if t == "" {
			continue
		}

		if t != "notype" {
			association.AddString("type", t)
		}

		association.AddString("itemid", aw.Value)
		association.AddString("representationtype", "partial")

		rank++
		association.AddInt("associationrank", rank)

		typerank, ok := types[t]
		if ok {
			typerank++
		} else {
			typerank = 1
		}
		types[t] = typerank
		association.AddInt("typerank", typerank)

		associations.AddObject(&association)
	}

	if rank > 0 {
		jo.AddArray("associations", &associations)
	}
}
