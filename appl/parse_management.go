package appl

import (
	"errors"
	"fmt"
	"strings"
)

const (
	DIGIT_ZERO rune = 48
)

func (pm *PublicationManagement) parse(aj *ApplJson) error {
	if pm.RecordType == "" {
		return errors.New("[PublicationManagement.RecordType] is missing")
	}
	if pm.FilingType == "" {
		return errors.New("[PublicationManagement.FilingType] is missing")
	}

	err := getPubStatus(aj)
	if err != nil {
		return err
	}

	err = getFirstCreatedDate(aj)
	if err != nil {
		return err
	}

	getManagementSignals(aj)
	getOutingInstructions(aj)
	getEditorialTypes(aj)
	getAssociatedWith(aj)
	getTimeRestrictions(aj)

	if len(pm.RefersTo) > 0 {
		aj.RefersTo = pm.RefersTo[0]
	}

	return nil
}

func getPubStatus(aj *ApplJson) error {
	status := aj.Xml.PublicationManagement.Status
	if strings.EqualFold(status, "Usable") || strings.EqualFold(status, "Embargoed") || strings.EqualFold(status, "Unknown") {
		aj.PubStatus = PUBSTATUS_USABLE
	} else if strings.EqualFold(status, "Withheld") {
		aj.PubStatus = PUBSTATUS_WITHHELD
	} else if strings.EqualFold(status, "Canceled") {
		aj.PubStatus = PUBSTATUS_CANCELED
	} else {
		e := fmt.Sprintf("Invalid pub status [%s]", status)
		return errors.New(e)
	}

	return nil
}

func getFirstCreatedDate(aj *ApplJson) error {
	fc := aj.Xml.PublicationManagement.FirstCreated

	if fc.Year <= 0 {
		e := fmt.Sprintf("Invalid year [%d]", fc.Year)
		return errors.New(e)
	}

	aj.FirstCreatedYear = fc.Year

	month := ""
	if fc.Month == 0 {
		aj.FirstCreated = fmt.Sprintf("%d", fc.Year)
		return nil
	} else {
		zero := ""
		if fc.Month < 10 {
			zero = "0"
		}
		month = fmt.Sprintf("%s%d", zero, fc.Month)
	}

	day := ""
	if fc.Day == 0 {
		aj.FirstCreated = fmt.Sprintf("%d-%s", fc.Year, month)
		return nil
	} else {
		zero := ""
		if fc.Day < 10 {
			zero = "0"
		}
		day = fmt.Sprintf("%s%d", zero, fc.Day)
	}

	if fc.Time == "" {
		aj.FirstCreated = fmt.Sprintf("%d-%s-%s", fc.Year, month, day)
	} else {
		aj.FirstCreated = fmt.Sprintf("%d-%s-%sT%sZ", fc.Year, month, day, fc.Time)
	}

	return nil
}

func getManagementSignals(aj *ApplJson) {
	pm := aj.Xml.PublicationManagement

	if pm.ExplicitWarning == "1" {
		aj.Signals.Add("explicitcontent")
	} else if strings.EqualFold(pm.ExplicitWarning, "NUDITY") {
		aj.Signals.Add("NUDITY")
	} else if strings.EqualFold(pm.ExplicitWarning, "OBSCENITY") {
		aj.Signals.Add("OBSCENITY")
	} else if strings.EqualFold(pm.ExplicitWarning, "GRAPHIC CONTENT") {
		aj.Signals.Add("GRAPHICCONTENT")
	}

	if strings.EqualFold(pm.IsDigitized, "false") {
		aj.Signals.Add("isnotdigitized")
	}
}

func getOutingInstructions(aj *ApplJson) {
	pm := aj.Xml.PublicationManagement

	if pm.Instruction == nil || len(pm.Instruction) == 0 {
		return
	}

	for _, instruction := range pm.Instruction {
		aj.OutingInstructions.Add(instruction)
	}
}

func getEditorialTypes(aj *ApplJson) {
	pm := aj.Xml.PublicationManagement

	if pm.Editorial == nil || len(pm.Editorial) == 0 {
		return
	}

	embargoed := pm.ReleaseDateTime != ""

	for _, editorialtype := range pm.Editorial {
		aj.EditorialTypes.Add(editorialtype.Type)

		if embargoed {
			if strings.EqualFold(editorialtype.Type, "Advance") || strings.EqualFold(editorialtype.Type, "HoldForRelease") {
				embargoed = false
				aj.Embargoed = pm.ReleaseDateTime + "Z"
			}
		}
	}
}

func getTimeRestrictions(aj *ApplJson) {
	pm := aj.Xml.PublicationManagement
	if pm.TimeRestrictions.TimeRestriction == nil || len(pm.TimeRestrictions.TimeRestriction) == 0 {
		return
	}

	timeRestrictions := make(map[string]bool)
	for _, tr := range pm.TimeRestrictions.TimeRestriction {
		if tr.System != "" && tr.Zone != "" {
			name := fmt.Sprintf("%s%s", tr.System, tr.Zone)
			timeRestrictions[name] = tr.Include
		}
	}

	aj.TimeRestrictions = timeRestrictions
}

func getAssociatedWith(aj *ApplJson) {
	/*
	   -test the value of //AssociatedWith, if its all zeros, do not convert; otherwise, each //AssociatedWith is converted to an object $.associations[i];
	   -each object $.associations[i] has five name/value pairs, $.associations[i].type, $.associations[i].itemid, $.associations[i].representationtype, $.associations[i].associationrank and $associations[i].typerank;
	   --retrieve the value from AssociatedWith/@CompositionType, use “Appendix III: CompositionType/Type Lookup Table” to derive the value for $.associations.association{i}.type;
	   --test the value of AssociatedWith, if it’s not all zeros, load as is to  $.associations[i].itemid;
	   --hardcode “partial” for $.associations[i].representationtype;
	   --load the sequence number of the AssociatedWith node (a number starting at 1) to $.associations[i].associationrank as a number;
	   --load the sequence number of the AssociatedWith node by @CompositionType (a number starting at 1) to $.associations[i].typerank as a number; note that CompositionType may be absent OR ‘StandardIngestedContent’ (which does not output a type) and any such AssociatedWith nodes should be ranked on their own.
	*/
	pm := aj.Xml.PublicationManagement
	associations := []ApplAssociation{}

	rank := 0
	types := make(map[string]int)

	for _, aw := range pm.AssociatedWith {
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

		association := ApplAssociation{}

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
			association.Type = t
		}

		association.ItemId = aw.Value

		rank++
		association.Rank = rank

		typerank, ok := types[t]
		if ok {
			typerank++
		} else {
			typerank = 1
		}
		types[t] = typerank
		association.TypeRank = typerank

		associations = append(associations, association)
	}

	aj.Associations = associations
}
