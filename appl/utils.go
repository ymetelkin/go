package appl

import (
	"fmt"
	"strings"

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
	} else {
		_, ok := ua.keys[s]
		if ok {
			return
		}
	}

	ua.keys[s] = true
	ua.values.AddString(s)
}

func (ua *uniqueArray) AddKeyValue(kn string, kv string, vn string, vv string) {
	if kv == "" || vv == "" {
		return
	}

	key := fmt.Sprintf("%s_%s", kv, vv)
	if ua.keys == nil {
		ua.keys = make(map[string]bool)
	} else {
		_, ok := ua.keys[key]
		if ok {
			return
		}
	}

	ua.keys[key] = true

	jo := json.Object{}
	jo.AddString(kn, kv)
	jo.AddString(vn, vv)
	ua.values.AddObject(&jo)
}

func (ua *uniqueArray) AddObject(key string, jo *json.Object) {
	if jo.IsEmpty() {
		return
	}

	if ua.keys == nil {
		ua.keys = make(map[string]bool)
	} else {
		_, ok := ua.keys[key]
		if ok {
			return
		}
	}

	ua.keys[key] = true
	ua.values.AddObject(jo)
}

func (ua *uniqueArray) IsEmpty() bool {
	return ua.values.Length() == 0
}

func (ua *uniqueArray) ToJsonProperty(field string) json.Property {
	if ua.values.Length() == 0 {
		return json.Property{}
	}
	return json.NewArrayProperty(field, &ua.values)
}

func setRels(c Classification, o Occurrence, rels *uniqueArray) {
	if strings.EqualFold(c.System, "RTE") {
		rels.AddString("inferred")
	} else if strings.EqualFold(o.ActualMatch, "true") {
		rels.AddString("direct")
	} else if strings.EqualFold(o.ActualMatch, "false") {
		rels.AddString("ancestor")
	}
}

func getGeoProperty(lat float64, long float64) json.Property {
	if lat == 0 || long == 0 {
		return json.Property{}
	}

	coordinates := json.Array{}
	coordinates.AddFloat(long)
	coordinates.AddFloat(lat)
	geometry := json.Object{}
	geometry.AddString("type", "Point")
	geometry.AddArray("coordinates", &coordinates)
	return json.NewObjectProperty("geometry_geojson", &geometry)
}

func makePrettyString(s string) string {
	if s == "" {
		return s
	}

	s = strings.Trim(s, " ")
	if s == "" {
		return s
	}

	runes := []rune(s)
	size := len(runes)
	t := make([]rune, size)
	i := 0
	w := 0

	for i < size {
		r := runes[i]
		if r > 13 || r < 7 {
			t[w] = r
			w++
		}
		i++
	}
	if w == 0 {
		return s
	}

	return string(t[0:w])
}
