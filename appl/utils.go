package appl

import (
	"errors"
	"fmt"
	"html"
	"sort"
	"strings"
	"time"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type uniqueStrings struct {
	keys   map[string]bool
	values []string
}

func (us *uniqueStrings) Append(s string) {
	if s == "" {
		return
	}

	if us.keys == nil {
		us.keys = make(map[string]bool)
	} else {
		_, ok := us.keys[s]
		if ok {
			return
		}
	}

	us.keys[s] = true
	us.values = append(us.values, s)
}

func (us *uniqueStrings) IsEmpty() bool {
	return us.values == nil || len(us.values) == 0
}

func (us *uniqueStrings) Values() []string {
	return us.values
}

func (us *uniqueStrings) JSONArray() (ja json.Array) {
	if !us.IsEmpty() {
		for _, v := range us.values {
			ja.AddString(v)
		}
	}
	return
}

func (us *uniqueStrings) AppendRel(system string, match *bool) {
	if strings.EqualFold(system, "RTE") {
		us.Append("inferred")
	} else if match != nil {
		if *match {
			us.Append("direct")
		} else {
			us.Append("ancestor")
		}
	}
}

type uniqueCodeNames struct {
	keys   map[string]bool
	values []CodeName
}

func (us *uniqueCodeNames) Append(k string, v string) {
	if k == "" || v == "" {
		return
	}

	key := fmt.Sprintf("%s_%s", k, v)

	if us.keys == nil {
		us.keys = make(map[string]bool)
	} else {
		_, ok := us.keys[key]
		if ok {
			return
		}
	}

	us.keys[key] = true
	cn := CodeName{
		Code: k,
		Name: v,
	}
	us.values = append(us.values, cn)
}

func (us *uniqueCodeNames) IsEmpty() bool {
	return us.values == nil || len(us.values) == 0
}

func (us *uniqueCodeNames) Values() []CodeName {
	return us.values
}

func (us *uniqueCodeNames) JSONArray(f func(CodeName) json.Object) (ja json.Array) {
	if !us.IsEmpty() {
		for _, cm := range us.values {
			var jo json.Object
			if f == nil {
				jo = cm.json()
			} else {
				jo = f(cm)
			}
			ja.AddObject(jo)
		}
	}
	return
}

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
	ua.values.AddObject(jo)
}

func (ua *uniqueArray) AddObject(key string, jo json.Object) {
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

func (ua *uniqueArray) JSONProperty(field string) json.Property {
	if ua.values.Length() == 0 {
		return json.Property{}
	}
	return json.NewArrayProperty(field, ua.values)
}

func (ua *uniqueArray) JSONArray() json.Array {
	return ua.values
}

func (ua *uniqueArray) Values() (values []string) {
	if ua.keys == nil {
		return
	}

	for k := range ua.keys {
		values = append(values, k)
	}

	sort.Strings(values)

	return
}

func getGeoProperty(lat float64, long float64) json.Property {
	if lat == 0 && long == 0 {
		return json.Property{}
	}

	coordinates := json.Array{}
	coordinates.AddFloat(long)
	coordinates.AddFloat(lat)
	geometry := json.Object{}
	geometry.AddString("type", "Point")
	geometry.AddArray("coordinates", coordinates)
	return json.NewObjectProperty("geometry_geojson", geometry)
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

func parseDate(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, errors.New("Missing date input")
	}

	formats := []string{
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02",
		"2006-01",
		"2006",
		"02/01/06",
		"02/01/2006",
	}

	for _, format := range formats {
		ts, err := time.Parse(format, s)
		if err == nil {
			return ts.UTC(), nil
		}
	}

	return time.Time{}, errors.New("Invalid date input")
}

func properDate(s string) string {
	ts, err := parseDate(s)
	if err != nil {
		return ""
	}

	return ts.Format("2006-01-02T15:04:05Z")
}

func formatDate(ts time.Time) string {
	return ts.Format("2006-01-02T15:04:05Z")
}

func parseTime(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, errors.New("Missing time input")
	}

	formats := []string{
		"15:04:05",
		"15:04",
		"15:04:05Z",
		"15:04:05-0700",
	}

	for _, format := range formats {
		ts, err := time.Parse(format, s)
		if err == nil {
			return ts.UTC(), nil
		}
	}

	return time.Time{}, errors.New("Invalid time input")
}

/*
func parseIsoDate(s string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t, err = time.Parse("01/02/06", s)
		if err != nil {
			return s
		}
	}
	return t.Format("2006-01-02T15:04:05Z")
}
*/

func formatTime(ms int64) string {
	t := time.Unix(0, ms*int64(time.Millisecond)).UTC()
	return t.Format("15:04:05.000")
}

func getForeignKeys(nd xml.Node) []foreignkey {
	var fks []foreignkey

	system := nd.Attribute("System")
	if system != "" && nd.Nodes != nil {
		for _, k := range nd.Nodes {
			if k.Attributes != nil {
				var id, fld string
				for k, v := range k.Attributes {
					switch k {
					case "Id":
						id = v
					case "Field":
						fld = v
					}
				}
				if id != "" && fld != "" {
					field := system + fld
					field = strings.ReplaceAll(field, " ", "")
					field = strings.ToLower(field)
					field = html.EscapeString(field)
					fk := foreignkey{Field: field, Value: id}

					if fks == nil {
						fks = []foreignkey{fk}
					} else {
						fks = append(fks, fk)
					}
				}
			}
		}
	}
	return fks
}
