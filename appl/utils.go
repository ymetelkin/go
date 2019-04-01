package appl

import (
	"fmt"

	"github.com/ymetelkin/go/json"
)

type uniqueArray struct {
	keys   map[string]bool
	values json.Array
}

func (ua *uniqueArray) AddString(s string) {
	if s == "" {
		return
	}

	if ua.keys == nil {
		ua.keys = make(map[string]bool)
		ua.keys[s] = true
	} else {
		_, ok := ua.keys[s]
		if ok {
			return
		}
	}

	ua.values.AddString(s)
}

func (ua *uniqueArray) AddKeyValue(kn string, kv string, vn string, vv string) {
	if kv == "" || vv == "" {
		return
	}

	key := fmt.Sprintf("%s_%s", kv, vv)
	if ua.keys == nil {
		ua.keys = make(map[string]bool)
		ua.keys[key] = true
	} else {
		_, ok := ua.keys[key]
		if ok {
			return
		}
	}

	jo := json.Object{}
	jo.AddString(kn, kv)
	jo.AddString(vn, vv)
	ua.values.AddObject(&jo)
}

func (ua *uniqueArray) IsEmpty() bool {
	return ua.values.Length() == 0
}

func (ua *uniqueArray) ToJsonProperty(field string) *json.Property {
	if ua.values.Length() == 0 {
		return nil
	}
	return json.NewArrayProperty(field, &ua.values)
}
