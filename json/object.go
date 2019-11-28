package json

import (
	"fmt"
	"strings"
)

//Object represents JSON object
type Object struct {
	names      []string
	pnames     map[string]int
	Properties map[string]value
}

func (jo *Object) addValue(name string, jv value) bool {
	if jo.Properties == nil {
		jo.Properties = make(map[string]value)
	}

	_, ok := jo.Properties[name]
	if ok {
		return false
	}

	jo.Properties[name] = jv

	if len(jo.names) == 0 {
		jo.names = []string{name}
	} else {
		jo.names = append(jo.names, name)
	}

	return true
}

//AddString adds string property to parent object
func (jo *Object) AddString(name string, value string) bool {
	return jo.addValue(name, newString(value))
}

//AddInt adds int property to parent object
func (jo *Object) AddInt(name string, value int) bool {
	return jo.addValue(name, newInt(value))
}

//AddFloat adds float property to parent object
func (jo *Object) AddFloat(name string, value float64) bool {
	return jo.addValue(name, newFloat(value))
}

//AddBool adds bool property to parent object
func (jo *Object) AddBool(name string, value bool) bool {
	return jo.addValue(name, newBool(value))
}

//AddObject adds JSON object property to parent object
func (jo *Object) AddObject(name string, value Object) bool {
	return jo.addValue(name, newObject(value))
}

//AddArray adds JSON array property to parent object
func (jo *Object) AddArray(name string, value Array) bool {
	return jo.addValue(name, newArray(value))
}

func (jo *Object) setValue(name string, value value) {
	if len(jo.Properties) > 0 {
		_, ok := jo.Properties[name]
		if ok {
			jo.Properties[name] = value
			return
		}
	}
	jo.addValue(name, value)
}

//SetInt sets int value of named property
func (jo *Object) SetInt(name string, value int) {
	jo.setValue(name, newInt(value))
}

//SetFloat sets float value of named property
func (jo *Object) SetFloat(name string, value float64) {
	jo.setValue(name, newFloat(value))
}

//SetBool sets int value of named property
func (jo *Object) SetBool(name string, value bool) {
	jo.setValue(name, newBool(value))
}

//SetString sets string value of named property
func (jo *Object) SetString(name string, value string) {
	jo.setValue(name, newString(value))
}

//SetObject sets JSON object value of named property
func (jo *Object) SetObject(name string, value Object) {
	jo.setValue(name, newObject(value))
}

//SetArray sets JSON array value of named property
func (jo *Object) SetArray(name string, value Array) {
	jo.setValue(name, newArray(value))
}

//Remove removes named property
func (jo *Object) Remove(name string) {
	if len(jo.Properties) == 0 {
		return
	}

	_, ok := jo.Properties[name]
	if !ok {
		return
	}

	delete(jo.Properties, name)

	if len(jo.Properties) == 0 {
		return
	}

	var (
		tmp = make([]string, len(jo.Properties))
		i   int
	)

	for _, n := range jo.names {
		if n != name {
			tmp[i] = n
			i++
		}
	}
	jo.names = tmp
}

func (jo *Object) getValue(name string) (jv value, ok bool) {
	if len(jo.Properties) == 0 {
		return
	}
	jv, ok = jo.Properties[name]
	return
}

//GetString returns string value of named property
func (jo *Object) GetString(name string) (s string, ok bool) {
	v, k := jo.getValue(name)
	if k {
		s, ok = v.GetString()
	}
	return
}

//GetInt returns string int of named property
func (jo *Object) GetInt(name string) (i int, ok bool) {
	v, k := jo.getValue(name)
	if k {
		i, ok = v.GetInt()
	}
	return
}

//GetFloat returns float value of named property
func (jo *Object) GetFloat(name string) (f float64, ok bool) {
	v, k := jo.getValue(name)
	if k {
		f, ok = v.GetFloat()
	}
	return
}

//GetBool returns bool value of named property
func (jo *Object) GetBool(name string) (b bool, ok bool) {
	v, k := jo.getValue(name)
	if k {
		b, ok = v.GetBool()
	}
	return
}

//GetObject returns JSON object value of named property
func (jo *Object) GetObject(name string) (o Object, ok bool) {
	v, k := jo.getValue(name)
	if k {
		o, ok = v.GetObject()
	}
	return
}

//GetArray returns JSON array value of named property
func (jo *Object) GetArray(name string) (ja Array, ok bool) {
	v, k := jo.getValue(name)
	if k {
		ja, ok = v.GetArray()
	}
	return
}

//IsEmpty checks for properties presense
func (jo *Object) IsEmpty() bool {
	return len(jo.Properties) == 0
}

//Names returns all field names
func (jo *Object) Names() []string {
	return jo.names
}

//Matches compares two objects
func (jo *Object) Matches(other *Object) (match bool, s string) {
	if jo == nil {
		s = "Left is nil"
		return
	}
	if other == nil {
		s = "Right is nil"
		return
	}

	props := make(map[string][]value)

	for k, l := range jo.Properties {
		r, ok := other.Properties[k]
		if ok {
			props[k] = []value{l, r}
		} else {
			s = fmt.Sprintf("Extra property: %s", k)
			return
		}
	}

	for k := range other.Properties {
		_, ok := jo.Properties[k]
		if !ok {
			s = fmt.Sprintf("Missing property: %s", k)
			return
		}
	}

	for k, v := range props {
		l := &v[0]
		r := &v[1]
		match, s = l.Matches(r)
		if !match {
			s = fmt.Sprintf("Mismatched property \"%s\": %s", k, s)
			return
		}
	}

	match = true
	return
}

//ToString returns pretty serialization
func (jo *Object) String() string {
	return jo.toString(true, 0)
}

//InlineString returns condensed serialization
func (jo *Object) InlineString() string {
	return jo.toString(false, 0)
}

func (jo *Object) toString(pretty bool, level int) string {
	if jo.Properties == nil || len(jo.Properties) == 0 {
		return "{}"
	}

	var sb strings.Builder

	sb.WriteByte('{')
	if pretty {
		sb.WriteByte('\r')
		sb.WriteByte('\n')
	}

	next := level + 1

	for index, name := range jo.names {
		jv, ok := jo.getValue(name)
		if ok {
			if index > 0 {
				sb.WriteByte(',')

				if pretty {
					sb.WriteByte('\r')
					sb.WriteByte('\n')
				}
			}

			if pretty {
				i := 0
				for i <= level {
					sb.WriteByte(' ')
					sb.WriteByte(' ')
					i++
				}
			}

			sb.WriteByte('"')
			sb.WriteString(name)
			sb.WriteByte('"')
			sb.WriteByte(':')
			if pretty {
				sb.WriteByte(' ')
			}
			s := jv.String(pretty, next)
			sb.WriteString(s)
		}
	}

	if pretty {
		sb.WriteByte('\r')
		sb.WriteByte('\n')
		i := 0
		for i < level {
			sb.WriteByte(' ')
			sb.WriteByte(' ')
			i++
		}
	}
	sb.WriteByte('}')

	return sb.String()
}
