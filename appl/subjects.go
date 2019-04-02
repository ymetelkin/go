package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
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

func (sbjs *subjects) Parse(c Classification) {
	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			var sbj subject

			key := fmt.Sprintf("%s_%s", o.Id, o.Value)

			if sbjs.Keys == nil {
				sbjs.Keys = make(map[string]int)
				sbjs.Subjects = []subject{}
			}

			i, ok := sbjs.Keys[key]
			if ok {
				sbj = sbjs.Subjects[i]
			} else {
				sbj = subject{Name: o.Value, Code: o.Id, Creator: c.System}
				sbjs.Subjects = append(sbjs.Subjects, sbj)
				i = len(sbjs.Subjects) - 1
				sbjs.Keys[key] = i
			}

			if sbj.Creator == "" || strings.EqualFold(c.System, "Editorial") {
				sbj.Creator = c.System
			}

			setRels(c, o, &sbj.Rels)

			sbj.ParentIds.AddString(o.ParentId)
			sbj.TopParent = o.TopParent

			sbjs.Subjects[i] = sbj
		}
	}
}

func (sbjs *subjects) ToJsonProperty(field string) *json.Property {
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
				subject.AddProperty(sbj.Rels.ToJsonProperty("rels"))
			}
			if !sbj.ParentIds.IsEmpty() {
				subject.AddProperty(sbj.ParentIds.ToJsonProperty("parentids"))
			}
			if sbj.TopParent {
				subject.AddBool("topparent", true)
			}
			ja.AddObject(&subject)
		}
		return json.NewArrayProperty(field, &ja)
	}

	return nil
}
