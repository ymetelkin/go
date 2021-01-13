package v2

import (
	"fmt"
	"strings"
)

//Array is JSON array
type Array struct {
	Values []Value
	text   string
	params map[int][]Parameter
}

//ParseArray parses JSON array
func ParseArray(data []byte) (Array, error) {
	return parseArray(data, false)
}

//ParseArrayWithParameters parses parameterized JSON array
func ParseArrayWithParameters(data []byte) (Array, error) {
	return parseArray(data, true)
}

func parseArray(data []byte, parameterized bool) (Array, error) {
	p := newParser(data)
	err := p.SkipWS()
	if err != nil {
		return Array{}, err
	}
	if p.Byte != '[' {
		return Array{}, fmt.Errorf("parsing JSON array; expect '[', found '%s'", string(p.Byte))
	}
	return p.ParseArray(parameterized)
}

//NewArray constructs Array
func NewArray(vs ...Value) *Array {
	return &Array{
		Values: vs,
	}
}

//NewStringArray constructs string Array
func NewStringArray(vs []string) *Array {
	var ja Array
	if len(vs) > 0 {
		ja.Values = make([]Value, len(vs))
		for i, v := range vs {
			ja.Values[i] = &stringValue{
				value: v,
			}
		}
	}
	return &ja
}

//NewIntArray constructs int Array
func NewIntArray(vs []int) *Array {
	var ja Array
	if len(vs) > 0 {
		ja.Values = make([]Value, len(vs))
		for i, v := range vs {
			ja.Values[i] = &intValue{
				value: v,
			}
		}
	}
	return &ja
}

//NewFloatArray constructs float64 Array
func NewFloatArray(vs []float64) *Array {
	var ja Array
	if len(vs) > 0 {
		ja.Values = make([]Value, len(vs))
		for i, v := range vs {
			ja.Values[i] = &floatValue{
				value: v,
			}
		}
	}
	return &ja
}

//NewObjectArray constructs Object Array
func NewObjectArray(vs []*Object) *Array {
	var ja Array
	if len(vs) > 0 {
		ja.Values = make([]Value, len(vs))
		for i, v := range vs {
			ja.Values[i] = v
		}
	}
	return &ja
}

//AddString adds string value
func (ja *Array) AddString(v string) {
	ja.Values = append(ja.Values, &stringValue{
		value: v,
	})
}

//AddInt adds int value
func (ja *Array) AddInt(v int) {
	ja.Values = append(ja.Values, &intValue{
		value: v,
	})
}

//AddFloat adds float value
func (ja *Array) AddFloat(v float64) {
	ja.Values = append(ja.Values, &floatValue{
		value: v,
	})
}

//AddObject adds Object value
func (ja *Array) AddObject(v *Object) {
	ja.Values = append(ja.Values, v)
}

//GetStrings gets string values
func (ja *Array) GetStrings() ([]string, bool) {
	vs := make([]string, len(ja.Values))
	for i, jv := range ja.Values {
		v, ok := StringValue(jv)
		if !ok {
			return vs, false
		}
		vs[i] = v
	}
	return vs, true
}

//GetInts gets int values
func (ja *Array) GetInts() ([]int, bool) {
	vs := make([]int, len(ja.Values))
	for i, jv := range ja.Values {
		v, ok := IntValue(jv)
		if !ok {
			return vs, false
		}
		vs[i] = v
	}
	return vs, true
}

//GetFloats gets float64 values
func (ja *Array) GetFloats() ([]float64, bool) {
	vs := make([]float64, len(ja.Values))
	for i, jv := range ja.Values {
		v, ok := FloatValue(jv)
		if !ok {
			return vs, false
		}
		vs[i] = v
	}
	return vs, true
}

//GetObjects gets Object values
func (ja *Array) GetObjects() ([]*Object, bool) {
	vs := make([]*Object, len(ja.Values))
	for i, jv := range ja.Values {
		v, ok := ObjectValue(jv)
		if !ok {
			return vs, false
		}
		vs[i] = v
	}
	return vs, true
}

//Copy copies Object
func (ja *Array) Copy() *Array {
	if ja == nil {
		return nil
	}
	if len(ja.Values) == 0 {
		return &Array{}
	}

	var copy Array
	for _, jv := range ja.Values {
		v, ok := copyValue(jv)
		if ok {
			copy.Values = append(copy.Values, v)
		}
	}

	if len(ja.params) > 0 {
		copy.params = make(map[int][]Parameter)
		for i, v := range ja.params {
			copy.params[i] = v
		}
	}

	return &copy
}

//Equals compares two arrays
func (ja *Array) Equals(other *Array) bool {
	if ja.Values == nil || other.Values == nil || len(ja.Values) != len(other.Values) {
		return false
	}

	for _, lv := range ja.Values {
		var ok bool
		for _, rv := range other.Values {
			ok = compareValues(lv, rv)
			if ok {
				break
			}
		}
		if !ok {
			return false
		}
	}

	return true
}

//SetParameters replaces parameter placeholders with values
func (ja *Array) SetParameters(params *Object) Array {
	var set Array

	for i, v := range ja.Values {
		pms, _ := ja.params[i]
		v = setValueParameters(v, pms, params)
		if v == nil {
			continue
		}
		switch v.Type() {
		case StringType:
			if v.Value() == "" {
				continue
			}
		case ObjectType:
			o, ok := v.(*Object)
			if !ok || len(o.Properties) == 0 {
				continue
			}
		case ArrayType:
			a, ok := v.(*Array)
			if !ok || len(a.Values) == 0 {
				continue
			}
		}
		set.Values = append(set.Values, v)
	}

	return set
}

//GetParameters retrieves paramaters from Array
func (ja *Array) GetParameters() []Parameter {
	var (
		params []Parameter
	)

	for i, v := range ja.Values {
		switch v.Type() {
		case StringType:
			pms, ok := ja.params[i]
			if ok && len(pms) > 0 {
				params = append(params, pms...)
			}
		case ObjectType:
			o, ok := v.(*Object)
			if ok {
				params = append(params, o.GetParameters()...)
			}
		case ArrayType:
			a, ok := v.(*Array)
			if ok {
				params = append(params, a.GetParameters()...)
			}
		}
	}

	return params
}

//Value Value interface
func (ja *Array) Value() interface{} {
	return ja.Values
}

//Type Value interface
func (ja *Array) Type() ValueType {
	return ArrayType
}

//String Value interface
func (ja *Array) String() string {
	if ja.text == "" {
		sz := len(ja.Values)
		if sz == 0 {
			ja.text = "[]"
		} else {
			values := make([]string, sz)
			for i, v := range ja.Values {
				values[i] = v.String()
			}
			ja.text = fmt.Sprintf("[%s]", strings.Join(values, ","))
		}
	}
	return ja.text
}

//used when a first byte is '['
func (p *byteParser) ParseArray(parameterized bool) (Array, error) {
	var (
		params map[int][]Parameter
		values []Value
	)
	for {
		idx := p.Index
		v, pms, err := p.ParseValue(parameterized)
		if err != nil {
			if p.Byte == ']' {
				break
			}
			return Array{}, fmt.Errorf("parsing array at %d: %s", idx, err.Error())
		}

		if len(pms) > 0 {
			if params == nil {
				params = make(map[int][]Parameter)
			}
			params[len(values)] = pms
		}

		if p.Byte == ',' {
			values = append(values, v)
			continue
		}

		if p.Byte == ']' {
			values = append(values, v)
			break
		}

	}
	return Array{
		Values: values,
		params: params,
	}, p.SkipWS()
}
