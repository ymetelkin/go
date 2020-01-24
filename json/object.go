package json

import (
	"fmt"
	"strings"
)

//Object represents JSON object
type Object struct {
	Properties []Property
	fields     map[string]int
	size       int
	params     map[string]int
}

//New creates new Object with properties
func New(props ...Property) (jo Object) {
	if len(props) == 0 {
		return
	}
	for _, jp := range props {
		jo.Add(jp.Field, jp.Value)
	}
	return
}

//Add adds new property. Returns false if field exists.
func (jo *Object) Add(field string, value Value) bool {
	if jo.size == 0 {
		jo.Properties = []Property{Property{
			Field: field,
			Value: value,
		}}
		jo.fields = map[string]int{
			field: 0,
		}
		jo.size = 1
	} else {
		if _, ok := jo.fields[field]; ok {
			return false
		}

		jo.Properties = append(jo.Properties, Property{
			Field: field,
			Value: value,
		})
		jo.fields[field] = jo.size
		jo.size++
	}
	return true
}

//AddString adds string property to parent object
func (jo *Object) AddString(field string, value string) bool {
	return jo.Add(field, NewString(value))
}

//AddInt adds int property to parent object
func (jo *Object) AddInt(field string, value int) bool {
	return jo.Add(field, NewInt(value))
}

//AddFloat adds float property to parent object
func (jo *Object) AddFloat(field string, value float64) bool {
	return jo.Add(field, NewFloat(value))
}

//AddBool adds bool property to parent object
func (jo *Object) AddBool(field string, value bool) bool {
	return jo.Add(field, NewBool(value))
}

//AddObject adds JSON object property to parent object
func (jo *Object) AddObject(field string, value Object) bool {
	return jo.Add(field, NewObject(value))
}

//AddArray adds JSON array property to parent object
func (jo *Object) AddArray(field string, value Array) bool {
	return jo.Add(field, NewArray(value))
}

//AddStringArray adds string array property to parent object
func (jo *Object) AddStringArray(field string, values []string) bool {
	if len(values) == 0 {
		return false
	}
	return jo.Add(field, NewArray(NewStringArray(values)))
}

//AddIntArray adds int array property to parent object
func (jo *Object) AddIntArray(field string, values []int) bool {
	if len(values) == 0 {
		return false
	}
	return jo.Add(field, NewArray(NewIntArray(values)))
}

//AddFloatArray adds float64 array property to parent object
func (jo *Object) AddFloatArray(field string, values []float64) bool {
	if len(values) == 0 {
		return false
	}
	return jo.Add(field, NewArray(NewFloatArray(values)))
}

//AddBoolArray adds bool array property to parent object
func (jo *Object) AddBoolArray(field string, values []bool) bool {
	if len(values) == 0 {
		return false
	}
	return jo.Add(field, NewArray(NewBoolArray(values)))
}

//AddObjectArray adds Object array property to parent object
func (jo *Object) AddObjectArray(field string, values []Object) bool {
	if len(values) == 0 {
		return false
	}
	return jo.Add(field, NewArray(NewObjectArray(values)))
}

//AddArrayArray adds Array array property to parent object
func (jo *Object) AddArrayArray(field string, values []Array) bool {
	if len(values) == 0 {
		return false
	}
	return jo.Add(field, NewArray(NewArrayArray(values)))
}

//Set adds new property. If field exists, overwrites its value.
func (jo *Object) Set(field string, value Value) {
	if jo.size > 0 {
		if i, ok := jo.fields[field]; ok {
			jo.Properties[i] = Property{
				Field: field,
				Value: value,
			}
			return
		}
	}

	jo.Add(field, value)
}

//SetInt sets int value of named property
func (jo *Object) SetInt(field string, value int) {
	jo.Set(field, NewInt(value))
}

//SetFloat sets float value of named property
func (jo *Object) SetFloat(field string, value float64) {
	jo.Set(field, NewFloat(value))
}

//SetBool sets int value of named property
func (jo *Object) SetBool(field string, value bool) {
	jo.Set(field, NewBool(value))
}

//SetString sets string value of named property
func (jo *Object) SetString(field string, value string) {
	jo.Set(field, NewString(value))
}

//SetObject sets JSON object value of named property
func (jo *Object) SetObject(field string, value Object) {
	jo.Set(field, NewObject(value))
}

//SetArray sets JSON array value of named property
func (jo *Object) SetArray(field string, value Array) {
	jo.Set(field, NewArray(value))
}

//Remove removes property. Returns false if field doesn't exists.
func (jo *Object) Remove(field string) bool {
	if jo.size == 0 {
		return false
	}

	if i, ok := jo.fields[field]; ok {
		if jo.size == 1 {
			jo.Properties = nil
			jo.fields = nil
			jo.size = 0
			return true
		}

		switch i {
		case 0:
			jo.Properties = jo.Properties[1:]
		case jo.size - 1:
			jo.Properties = jo.Properties[:i]
		default:
			jo.Properties = append(jo.Properties[:i], jo.Properties[i+1:]...)
		}

		delete(jo.fields, field)
		jo.size--

		for i, jp := range jo.Properties {
			jo.fields[jp.Field] = i
		}
		return true
	}

	return false
}

//Get gets value. Returns false if field doesn't exist.
func (jo *Object) Get(field string) (value Value, ok bool) {
	if jo.size > 0 {
		if i, k := jo.fields[field]; k {
			value = jo.Properties[i].Value
			ok = true
		}
	}
	return
}

//GetString returns string value of named property
func (jo *Object) GetString(field string) (s string, ok bool) {
	v, k := jo.Get(field)
	if k {
		s, ok = v.String()
	}
	return
}

//GetInt returns string int of named property
func (jo *Object) GetInt(field string) (i int, ok bool) {
	v, k := jo.Get(field)
	if k {
		i, ok = v.Int()
	}
	return
}

//GetFloat returns float value of named property
func (jo *Object) GetFloat(field string) (f float64, ok bool) {
	v, k := jo.Get(field)
	if k {
		f, ok = v.Float()
	}
	return
}

//GetBool returns bool value of named property
func (jo *Object) GetBool(field string) (b bool, ok bool) {
	v, k := jo.Get(field)
	if k {
		b, ok = v.Bool()
	}
	return
}

//GetObject returns JSON object value of named property
func (jo *Object) GetObject(field string) (o Object, ok bool) {
	v, k := jo.Get(field)
	if k {
		o, ok = v.Object()
	}
	return
}

//GetArray returns JSON array value of named property
func (jo *Object) GetArray(field string) (ja Array, ok bool) {
	v, k := jo.Get(field)
	if k {
		ja, ok = v.Array()
	}
	return
}

//Map returns object properties as map[string]Value
func (jo *Object) Map() (props map[string]Value) {
	if jo.size == 0 {
		return
	}

	props = make(map[string]Value)
	for _, jp := range jo.Properties {
		props[jp.Field] = jp.Value
	}

	return
}

//IsEmpty checks for properties presense
func (jo *Object) IsEmpty() bool {
	return jo.size == 0
}

//Matches compares two JSON objects
func (jo *Object) Matches(other *Object) (match bool, s string) {
	if jo == nil {
		s = "Left is nil"
		return
	}
	if other == nil {
		s = "Right is nil"
		return
	}
	if jo.size != other.size {
		s = "Property count mismatch"
	}

	values := make(map[string][]Value)

	for k, l := range jo.fields {
		r, ok := other.fields[k]
		if ok {
			values[k] = []Value{
				jo.Properties[l].Value,
				other.Properties[r].Value,
			}
		} else {
			s = fmt.Sprintf("Extra property: %s", k)
			return
		}
	}

	for k := range other.fields {
		_, ok := jo.fields[k]
		if !ok {
			s = fmt.Sprintf("Missing property: %s", k)
			return
		}
	}

	for k, v := range values {
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

//String returns pretty serialization
func (jo *Object) String() string {
	return jo.string(true, 0)
}

//InlineString returns condensed serialization
func (jo *Object) InlineString() string {
	return jo.string(false, 0)
}

func (jo *Object) string(pretty bool, level int) string {
	if jo.size == 0 {
		return "{}"
	}

	var sb strings.Builder

	sb.WriteByte('{')
	if pretty {
		sb.WriteByte('\r')
		sb.WriteByte('\n')
	}

	next := level + 1

	for i, jp := range jo.Properties {
		if i > 0 {
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
		sb.WriteString(jp.Field)
		sb.WriteByte('"')
		sb.WriteByte(':')
		if pretty {
			sb.WriteByte(' ')
		}
		s := jp.Value.string(pretty, next)
		sb.WriteString(s)
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
