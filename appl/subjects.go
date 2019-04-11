package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type subject struct {
	Name      string
	Code      string
	Creator   string
	Rels      uniqueArray
	ParentIds uniqueArray
	TopParent bool
}

type subjects struct {
	Keys     map[string]int
	Subjects []subject
}

func (sbjs *subjects) Parse(nd xml.Node) {
	if nd.Nodes == nil {
		return
	}

	system := nd.GetAttribute("System")

	for _, n := range nd.Nodes {
		if n.Name == "Occurrence" && n.Attributes != nil {
			var (
				code, name, match, pid string
				tp                     bool
			)

			for _, a := range n.Attributes {
				switch a.Name {
				case "Id":
					code = a.Value
				case "Value":
					name = a.Value
				case "ActualMatch":
					match = a.Value
				case "ParentID":
					pid = a.Value
				case "TopParent":
					tp = a.Value == "true"
				}
			}

			if code != "" && name != "" {
				var sbj subject

				key := fmt.Sprintf("%s_%s", code, name)

				if sbjs.Keys == nil {
					sbjs.Keys = make(map[string]int)
					sbjs.Subjects = []subject{}
				}

				i, ok := sbjs.Keys[key]
				if ok {
					sbj = sbjs.Subjects[i]
				} else {
					sbj = subject{Name: name, Code: code, Creator: system}
					sbjs.Subjects = append(sbjs.Subjects, sbj)
					i = len(sbjs.Subjects) - 1
					sbjs.Keys[key] = i
				}

				if sbj.Creator == "" || strings.EqualFold(system, "Editorial") {
					sbj.Creator = system
				}

				setRels(system, match, &sbj.Rels)

				sbj.ParentIds.AddString(pid)
				sbj.TopParent = tp

				sbjs.Subjects[i] = sbj
			}
		}
	}
}

func (sbjs *subjects) ToJSONProperty(field string) json.Property {
	if sbjs.Keys != nil {
		ja := json.Array{}
		for _, item := range sbjs.Subjects {
			sbj := item
			subject := json.Object{}
			subject.AddString("name", sbj.Name)
			subject.AddString("scheme", "http://cv.ap.org/id/")
			subject.AddString("code", sbj.Code)
			if sbj.Creator != "" {
				subject.AddString("creator", sbj.Creator)
			}
			if !sbj.Rels.IsEmpty() {
				subject.AddProperty(sbj.Rels.ToJSONProperty("rels"))
			}
			if !sbj.ParentIds.IsEmpty() {
				subject.AddProperty(sbj.ParentIds.ToJSONProperty("parentids"))
			}
			if sbj.TopParent {
				subject.AddBool("topparent", true)
			}
			ja.AddObject(subject)
		}
		return json.NewArrayProperty(field, ja)
	}

	return json.Property{}
}
