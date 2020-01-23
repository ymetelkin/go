package json

import (
	"fmt"
	"strings"
)

//Array represents JSON array
type Array struct {
	Values []Value
	params []int
}

//NewStringArray creates new string array
func NewStringArray(values []string) Array {
	var ja Array
	ja.AddStrings(values)
	return ja
}

//NewIntArray creates new int array
func NewIntArray(values []int) Array {
	var ja Array
	ja.AddInts(values)
	return ja
}

//NewFloatArray creates new float64 array
func NewFloatArray(values []float64) Array {
	var ja Array
	ja.AddFloats(values)
	return ja
}

//NewBoolArray creates new bool array
func NewBoolArray(values []bool) Array {
	var ja Array
	ja.AddBools(values)
	return ja
}

//NewObjectArray creates new Object array
func NewObjectArray(values []Object) Array {
	var ja Array
	ja.AddObjects(values)
	return ja
}

//NewArrayArray creates new Array array
func NewArrayArray(values []Array) Array {
	var ja Array
	ja.AddArrays(values)
	return ja
}

//AddString adds string to JSON array
func (ja *Array) AddString(value string) {
	ja.Values = append(ja.Values, NewString(value))
}

//AddInt adds int to JSON array
func (ja *Array) AddInt(value int) {
	ja.Values = append(ja.Values, NewInt(value))
}

//AddFloat adds float to JSON array
func (ja *Array) AddFloat(value float64) {
	ja.Values = append(ja.Values, NewFloat(value))
}

//AddBool adds bool to JSON array
func (ja *Array) AddBool(value bool) {
	ja.Values = append(ja.Values, NewBool(value))
}

//AddObject adds JSON object to JSON array
func (ja *Array) AddObject(value Object) {
	ja.Values = append(ja.Values, NewObject(value))
}

//AddArray adds JSON array to JSON array
func (ja *Array) AddArray(value Array) {
	ja.Values = append(ja.Values, NewArray(value))
}

//AddStrings adds strings to JSON array
func (ja *Array) AddStrings(values []string) {
	if len(values) == 0 {
		return
	}
	buff := make([]Value, len(values))
	for i, value := range values {
		buff[i] = NewString(value)
	}
	if len(ja.Values) == 0 {
		ja.Values = buff
	} else {
		ja.Values = append(ja.Values, buff...)
	}
}

//AddInts adds ints to JSON array
func (ja *Array) AddInts(values []int) {
	if len(values) == 0 {
		return
	}
	buff := make([]Value, len(values))
	for i, value := range values {
		buff[i] = NewInt(value)
	}
	if len(ja.Values) == 0 {
		ja.Values = buff
	} else {
		ja.Values = append(ja.Values, buff...)
	}
}

//AddFloats adds floats to JSON array
func (ja *Array) AddFloats(values []float64) {
	if len(values) == 0 {
		return
	}
	buff := make([]Value, len(values))
	for i, value := range values {
		buff[i] = NewFloat(value)
	}
	if len(ja.Values) == 0 {
		ja.Values = buff
	} else {
		ja.Values = append(ja.Values, buff...)
	}
}

//AddBools adds bools to JSON array
func (ja *Array) AddBools(values []bool) {
	if len(values) == 0 {
		return
	}
	buff := make([]Value, len(values))
	for i, value := range values {
		buff[i] = NewBool(value)
	}
	if len(ja.Values) == 0 {
		ja.Values = buff
	} else {
		ja.Values = append(ja.Values, buff...)
	}
}

//AddObjects adds JSON objects to JSON array
func (ja *Array) AddObjects(values []Object) {
	if len(values) == 0 {
		return
	}
	buff := make([]Value, len(values))
	for i, value := range values {
		buff[i] = NewObject(value)
	}
	if len(ja.Values) == 0 {
		ja.Values = buff
	} else {
		ja.Values = append(ja.Values, buff...)
	}
}

//AddArrays adds JSON arrays to JSON array
func (ja *Array) AddArrays(values []Array) {
	if len(values) == 0 {
		return
	}
	buff := make([]Value, len(values))
	for i, value := range values {
		buff[i] = NewArray(value)
	}
	if len(ja.Values) == 0 {
		ja.Values = buff
	} else {
		ja.Values = append(ja.Values, buff...)
	}
}

//Set sets value at specified index
func (ja *Array) Set(index int, value Value) (err error) {
	if len(ja.Values) > index {
		err = fmt.Errorf("Position [%d] does not exist", index)
	} else {
		ja.Values[index] = value
	}
	return
}

//SetInt adds int JSON array at specified position
func (ja *Array) SetInt(index int, value int) error {
	return ja.Set(index, NewInt(value))
}

//SetFloat sets float in JSON array at specified position
func (ja *Array) SetFloat(index int, value float64) error {
	return ja.Set(index, NewFloat(value))
}

//SetBool sets bool in JSON array at specified position
func (ja *Array) SetBool(index int, value bool) error {
	return ja.Set(index, NewBool(value))
}

//SetString sets string in JSON array at specified position
func (ja *Array) SetString(index int, value string) error {
	return ja.Set(index, NewString(value))
}

//SetObject sets JSON object in JSON array at specified position
func (ja *Array) SetObject(index int, value Object) error {
	return ja.Set(index, NewObject(value))
}

//SetArray sets JSON array in JSON array at specified position
func (ja *Array) SetArray(index int, value Array) error {
	return ja.Set(index, NewArray(value))
}

//GetString gets string from JSON array at specified position
func (ja *Array) GetString(index int) (v string, ok bool) {
	if len(ja.Values) > index {
		v, ok = ja.Values[index].String()
	}
	return
}

//GetInt gets int from JSON array at specified position
func (ja *Array) GetInt(index int) (v int, ok bool) {
	if len(ja.Values) > index {
		v, ok = ja.Values[index].Int()
	}
	return
}

//GetFloat gets bool from JSON array at specified position
func (ja *Array) GetFloat(index int) (v float64, ok bool) {
	if len(ja.Values) > index {
		v, ok = ja.Values[index].Float()
	}
	return
}

//GetBool gets bool from JSON array at specified position
func (ja *Array) GetBool(index int) (v bool, ok bool) {
	if len(ja.Values) > index {
		v, ok = ja.Values[index].Bool()
	}
	return
}

//GetObject gets JSON object from JSON array at specified position
func (ja *Array) GetObject(index int) (v Object, ok bool) {
	if len(ja.Values) > index {
		v, ok = ja.Values[index].Object()
	}
	return
}

//GetArray gets JSON array from JSON array at specified position
func (ja *Array) GetArray(index int) (v Array, ok bool) {
	if len(ja.Values) > index {
		v, ok = ja.Values[index].Array()
	}
	return
}

//GetStrings converts all values to strings
func (ja *Array) GetStrings() (vs []string, ok bool) {
	if len(ja.Values) > 0 {
		vs = make([]string, len(ja.Values))
		for i, jv := range ja.Values {
			v, k := jv.String()
			if !k {
				return
			}
			vs[i] = v
		}
		ok = true
	}
	return
}

//GetInts converts all values to ints
func (ja *Array) GetInts() (vs []int, ok bool) {
	if len(ja.Values) > 0 {
		vs = make([]int, len(ja.Values))
		for i, jv := range ja.Values {
			v, k := jv.Int()
			if !k {
				return
			}
			vs[i] = v
		}
		ok = true
	}
	return
}

//GetFloats converts all values to floats
func (ja *Array) GetFloats() (vs []float64, ok bool) {
	if len(ja.Values) > 0 {
		vs = make([]float64, len(ja.Values))
		for i, jv := range ja.Values {
			v, k := jv.Float()
			if !k {
				return
			}
			vs[i] = v
		}
		ok = true
	}
	return
}

//GetBools converts all values to bools
func (ja *Array) GetBools() (vs []bool, ok bool) {
	if len(ja.Values) > 0 {
		vs = make([]bool, len(ja.Values))
		for i, jv := range ja.Values {
			v, k := jv.Bool()
			if !k {
				return
			}
			vs[i] = v
		}
		ok = true
	}
	return
}

//GetObjects converts all values to JSON objects
func (ja *Array) GetObjects() (vs []Object, ok bool) {
	if len(ja.Values) > 0 {
		vs = make([]Object, len(ja.Values))
		for i, jv := range ja.Values {
			v, k := jv.Object()
			if !k {
				return
			}
			vs[i] = v
		}
		ok = true
	}
	return
}

//GetArrays converts all values to JSON arrays
func (ja *Array) GetArrays() (vs []Array, ok bool) {
	if len(ja.Values) > 0 {
		vs = make([]Array, len(ja.Values))
		for i, jv := range ja.Values {
			v, k := jv.Array()
			if !k {
				return
			}
			vs[i] = v
		}
		ok = true
	}
	return
}

//Remove removes element from JSON array at specified position
func (ja *Array) Remove(index int) (ok bool) {
	if len(ja.Values) > index {
		values := ja.Values
		ja.Values = append(values[:index], values[index+1:]...)
		ok = true
	}
	return
}

//Matches compares array to another array
func (ja *Array) Matches(other *Array) (match bool, s string) {
	if ja.Values == nil {
		s = "Left is nil"
		return
	}
	if other.Values == nil {
		s = "Right is nil"
		return
	}
	if len(ja.Values) != len(other.Values) {
		s = fmt.Sprintf("Size mismatch: [ %v ] vs [ %v ]", len(ja.Values), len(other.Values))
		return
	}

	for i, lv := range ja.Values {
		var ok bool
		for _, rv := range other.Values {
			ok, s = lv.Matches(&rv)
			if ok {
				break
			}
		}
		if !ok {
			s = fmt.Sprintf("Array mismatch: [ %d ] [ %v ]", i, lv)
			return
		}
	}

	match = true
	return
}

//String transforms JSON array to pretty string
func (ja *Array) String() string {
	return ja.string(true, 0)
}

//InlineString transforms JSON array into inline string
func (ja *Array) InlineString() string {
	return ja.string(false, 0)
}

func (ja *Array) string(pretty bool, level int) string {
	if len(ja.Values) == 0 {
		return "[]"
	}

	var sb strings.Builder

	sb.WriteByte('[')
	if pretty {
		sb.WriteByte('\r')
		sb.WriteByte('\n')
	}

	next := level + 1

	for i, jv := range ja.Values {
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

		s := jv.string(pretty, next)
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
	sb.WriteByte(']')

	return sb.String()
}
