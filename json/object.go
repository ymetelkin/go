package json

import (
	"strings"
)

//Object JSON object
type Object struct {
	Properties []Property
	fields     map[string]int
	params     map[string]int
}

//New constucts Object
func New(fields ...Property) (jo Object) {
	if len(fields) == 0 {
		return
	}
	for _, f := range fields {
		jo.Add(f.Name, f.Value)
	}
	return
}

//Add adds field
func (jo *Object) Add(field string, value Value) bool {
	if value == nil || value.t() == 0 {
		return false
	}
	jp := Property{
		Name:  field,
		Value: value,
	}

	if jo.fields == nil {
		jo.fields = map[string]int{field: 0}
		jo.Properties = []Property{jp}
		return true
	}

	_, ok := jo.fields[field]
	if ok {
		return false
	}

	jo.fields[field] = len(jo.Properties)
	jo.Properties = append(jo.Properties, jp)
	return true
}

//Remove adds field
func (jo *Object) Remove(field string) (ok bool) {
	sz := len(jo.fields)
	if sz == 0 {
		return
	}

	if sz == 1 {
		_, ok = jo.fields[field]
		if ok {
			jo.fields = nil
			jo.Properties = nil
		}
		return
	}

	var (
		fs     = make([]Property, sz-1)
		fields = make(map[string]int)
		i      int
	)

	for _, jp := range jo.Properties {
		if jp.Name == field {
			ok = true
			continue
		}
		fs[i] = jp
		fields[jp.Name] = i
		i++
	}

	jo.Properties = fs
	jo.fields = fields
	return
}

//Set adds new property. If field exists, overwrites its value.
func (jo *Object) Set(field string, value Value) {
	if len(jo.fields) > 0 {
		if i, ok := jo.fields[field]; ok {
			jo.Properties[i] = Property{
				Name:  field,
				Value: value,
			}
			return
		}
	}
	jo.Add(field, value)
}

//GetValue value
func (jo *Object) GetValue(field string) (v Value, ok bool) {
	if len(jo.fields) == 0 {
		return
	}

	var i int
	i, ok = jo.fields[field]
	if !ok {
		return
	}

	v = jo.Properties[i].Value
	return
}

//GetString gets string value
func (jo *Object) GetString(field string) (v string, ok bool) {
	var jv Value
	jv, ok = jo.GetValue(field)
	if !ok {
		return
	}

	return jv.String()
}

//GetStrings gets string values
func (jo *Object) GetStrings(field string) (vs []string, ok bool) {
	var ja Array
	ja, ok = jo.GetArray(field)
	if !ok {
		return
	}

	return ja.GetStrings()
}

//GetInt gets int value
func (jo *Object) GetInt(field string) (v int, ok bool) {
	var jv Value
	jv, ok = jo.GetValue(field)
	if !ok {
		return
	}

	return jv.Int()
}

//GetInts gets int values
func (jo *Object) GetInts(field string) (vs []int, ok bool) {
	var ja Array
	ja, ok = jo.GetArray(field)
	if !ok {
		return
	}

	return ja.GetInts()
}

//GetFloat gets float64 value
func (jo *Object) GetFloat(field string) (v float64, ok bool) {
	var jv Value
	jv, ok = jo.GetValue(field)
	if !ok {
		return
	}

	return jv.Float()
}

//GetFloats gets float64 values
func (jo *Object) GetFloats(field string) (vs []float64, ok bool) {
	var ja Array
	ja, ok = jo.GetArray(field)
	if !ok {
		return
	}

	return ja.GetFloats()
}

//GetBool gets bool value
func (jo *Object) GetBool(field string) (v bool, ok bool) {
	var jv Value
	jv, ok = jo.GetValue(field)
	if !ok {
		return
	}

	return jv.Bool()
}

//GetObject gets Object value
func (jo *Object) GetObject(field string) (v Object, ok bool) {
	var jv Value
	jv, ok = jo.GetValue(field)
	if !ok {
		return
	}

	return jv.Object()
}

//GetObjects gets Object values
func (jo *Object) GetObjects(field string) (vs []Object, ok bool) {
	var ja Array
	ja, ok = jo.GetArray(field)
	if !ok {
		return
	}

	return ja.GetObjects()
}

//GetArray gets Object value
func (jo *Object) GetArray(field string) (v Array, ok bool) {
	var jv Value
	jv, ok = jo.GetValue(field)
	if !ok {
		return
	}

	return jv.Array()
}

//Map returns object properties as map[string]Value
func (jo *Object) Map() (props map[string]Value) {
	if len(jo.Properties) == 0 {
		return
	}

	props = make(map[string]Value)
	for _, jp := range jo.Properties {
		props[jp.Name] = jp.Value
	}

	return
}

//Equals compares two JSON objects
func (jo *Object) Equals(other *Object) bool {
	if jo == nil || other == nil || len(jo.Properties) != len(other.Properties) {
		return false
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
			return false
		}
	}

	for k := range other.fields {
		if _, ok := jo.fields[k]; !ok {
			return false
		}
	}

	for _, v := range values {
		l := v[0]
		r := v[1]
		if !compare(l, r) {
			return false
		}
	}

	return true
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
	if len(jo.Properties) == 0 {
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
		sb.WriteString(jp.Name)
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
