package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

type audiences struct {
	Keys      map[string]bool
	Audiences json.Array
}

func (as *audiences) Parse(c Classification) {
	geo := false

	if strings.EqualFold(c.Authority, "AP Audience") && strings.EqualFold(c.System, "Editorial") {
		if c.Occurrence != nil {
			for _, o := range c.Occurrence {
				if o.Id != "" && o.Value != "" {
					key := fmt.Sprintf("%s_%s", o.Id, o.Value)

					if as.Keys == nil {
						as.Keys = make(map[string]bool)
						as.Audiences = json.Array{}
					}

					_, ok := as.Keys[key]
					if ok {
						continue
					} else {
						as.Keys[key] = true
					}

					audience := json.Object{}
					audience.AddString("code", o.Id)
					audience.AddString("name", o.Value)

					if o.Property != nil && len(o.Property) > 0 {
						prop := o.Property[0]
						if prop.Value != "" {
							if strings.EqualFold(prop.Value, "AUDGEOGRAPHY") {
								geo = true
							}
							audience.AddString("type", prop.Value)
						}
					}

					as.Audiences.AddObject(&audience)
				}
			}
		}
	}

	if !geo {

	}
}

func (as *audiences) ToJsonProperty() *json.Property {
	if as.Keys != nil {
		return json.NewArrayProperty("audiences", &as.Audiences)
	}

	return nil
}
