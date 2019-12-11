package json

import (
	"fmt"
	"strings"
)

//Array is JSON array structure
type Array struct {
	Values  []value
	pvalues []int
}

func (ja *Array) addValue(jv value) int {
	if len(ja.Values) == 0 {
		ja.Values = []value{jv}
		return 1
	}
	ja.Values = append(ja.Values, jv)
	return len(ja.Values) - 1
}

func (ja *Array) addValues(jvs []value) int {
	if jvs != nil {
		if ja.Values == nil {
			ja.Values = jvs
		} else {
			ja.Values = append(ja.Values, jvs...)
		}
	}

	if ja.Values == nil {
		return 0
	}

	return len(ja.Values) - 1
}

//AddString adds string to JSON array
func (ja *Array) AddString(value string) int {
	return ja.addValue(newString(value))
}

//AddInt adds int to JSON array
func (ja *Array) AddInt(value int) int {
	return ja.addValue(newInt(value))
}

//AddFloat adds float to JSON array
func (ja *Array) AddFloat(value float64) int {
	return ja.addValue(newFloat(value))
}

//AddBool adds bool to JSON array
func (ja *Array) AddBool(value bool) int {
	return ja.addValue(newBool(value))
}

//AddObject adds JSON object to JSON array
func (ja *Array) AddObject(value Object) int {
	return ja.addValue(newObject(value))
}

//AddArray adds JSON array to JSON array
func (ja *Array) AddArray(value Array) int {
	return ja.addValue(newArray(value))
}

//AddStrings adds strings to JSON array
func (ja *Array) AddStrings(values []string) int {
	return ja.addValues(newStrings(values))
}

//AddInts adds ints to JSON array
func (ja *Array) AddInts(values []int) int {
	return ja.addValues(newInts(values))
}

//AddFloats adds floats to JSON array
func (ja *Array) AddFloats(values []float64) int {
	return ja.addValues(newFloats(values))
}

//AddBools adds bools to JSON array
func (ja *Array) AddBools(values []bool) int {
	return ja.addValues(newBools(values))
}

//AddObjects adds JSON objects to JSON array
func (ja *Array) AddObjects(values []Object) int {
	return ja.addValues(newObjects(values))
}

//AddArrays adds JSON arrays to JSON array
func (ja *Array) AddArrays(values []Array) int {
	return ja.addValues(newArrays(values))
}

func (ja *Array) setValue(index int, value value) error {
	if ja.Values == nil || len(ja.Values) <= index {
		return fmt.Errorf("Position [%d] does not exist", index)
	}
	ja.Values[index] = value
	return nil
}

//SetInt adds int JSON array at specified position
func (ja *Array) SetInt(index int, value int) error {
	return ja.setValue(index, newInt(value))
}

//SetFloat sets float in JSON array at specified position
func (ja *Array) SetFloat(index int, value float64) error {
	return ja.setValue(index, newFloat(value))
}

//SetBool sets bool in JSON array at specified position
func (ja *Array) SetBool(index int, value bool) error {
	return ja.setValue(index, newBool(value))
}

//SetString sets string in JSON array at specified position
func (ja *Array) SetString(index int, value string) error {
	return ja.setValue(index, newString(value))
}

//SetObject sets JSON object in JSON array at specified position
func (ja *Array) SetObject(index int, value Object) error {
	return ja.setValue(index, newObject(value))
}

//SetArray sets JSON array in JSON array at specified position
func (ja *Array) SetArray(index int, value Array) error {
	return ja.setValue(index, newArray(value))
}

//Remove removes element from JSON array at specified position
func (ja *Array) Remove(index int) error {
	if ja.Values == nil || len(ja.Values) <= index {
		return fmt.Errorf("Position [%d] does not exist", index)
	}

	values := ja.Values
	ja.Values = append(values[:index], values[index+1:]...)

	return nil
}

func (ja *Array) getValue(index int) (*value, bool) {
	if ja.Values == nil || len(ja.Values) <= index {
		return nil, false
	}

	return &ja.Values[index], true
}

//GetString gets string from JSON array at specified position
func (ja *Array) GetString(index int) (s string, ok bool) {
	v, k := ja.getValue(index)
	if k {
		s, ok = v.GetString()
	}
	return
}

//GetInt gets int from JSON array at specified position
func (ja *Array) GetInt(index int) (i int, ok bool) {
	v, k := ja.getValue(index)
	if k {
		i, ok = v.GetInt()
	}
	return
}

//GetFloat gets bool from JSON array at specified position
func (ja *Array) GetFloat(index int) (f float64, ok bool) {
	v, k := ja.getValue(index)
	if k {
		f, ok = v.GetFloat()
	}
	return
}

//GetBool gets bool from JSON array at specified position
func (ja *Array) GetBool(index int) (b bool, ok bool) {
	v, k := ja.getValue(index)
	if k {
		b, ok = v.GetBool()
	}
	return
}

//GetObject gets JSON object from JSON array at specified position
func (ja *Array) GetObject(index int) (jo Object, ok bool) {
	v, k := ja.getValue(index)
	if k {
		jo, ok = v.GetObject()
	}
	return
}

//GetArray gets JSON array from JSON array at specified position
func (ja *Array) GetArray(index int) (a Array, ok bool) {
	v, k := ja.getValue(index)
	if k {
		a, ok = v.GetArray()
	}
	return
}

//GetStrings converts all values to strings
func (ja *Array) GetStrings() (ss []string, ok bool) {
	size := ja.Length()
	if size == 0 {
		return
	}

	ss = make([]string, size)
	for i, v := range ja.Values {
		val, k := v.GetString()
		if !k {
			return
		}
		ss[i] = val
	}

	ok = true
	return
}

//GetInts converts all values to ints
func (ja *Array) GetInts() (is []int, ok bool) {
	size := ja.Length()
	if size == 0 {
		return
	}

	is = make([]int, size)
	for i, v := range ja.Values {
		val, k := v.GetInt()
		if !k {
			return
		}
		is[i] = val
	}

	ok = true
	return
}

//GetFloats converts all values to floats
func (ja *Array) GetFloats() (fs []float64, ok bool) {
	size := ja.Length()
	if size == 0 {
		return
	}

	fs = make([]float64, size)
	for i, v := range ja.Values {
		val, k := v.GetFloat()
		if !k {
			return
		}
		fs[i] = val
	}

	ok = true
	return
}

//GetBools converts all values to bools
func (ja *Array) GetBools() (bs []bool, ok bool) {
	size := ja.Length()
	if size == 0 {
		return
	}

	bs = make([]bool, size)
	for i, v := range ja.Values {
		val, k := v.GetBool()
		if !k {
			return
		}
		bs[i] = val
	}

	ok = true
	return
}

//GetObjects converts all values to JSON objects
func (ja *Array) GetObjects() (jos []Object, ok bool) {
	size := ja.Length()
	if size == 0 {
		return
	}

	jos = make([]Object, size)
	for i, v := range ja.Values {
		val, k := v.GetObject()
		if !k {
			return
		}
		jos[i] = val
	}

	ok = true
	return
}

//GetArrays converts all values to JSON arrays
func (ja *Array) GetArrays() (jas []Array, ok bool) {
	size := ja.Length()
	if size == 0 {
		return
	}

	jas = make([]Array, size)
	for i, v := range ja.Values {
		val, k := v.GetArray()
		if !k {
			return
		}
		jas[i] = val
	}

	ok = true
	return
}

//IsEmpty checks if JSON array has any elements
func (ja *Array) IsEmpty() bool {
	if ja.Values == nil || len(ja.Values) == 0 {
		return true
	}

	return false
}

//Length returns number of elements in JSON array
func (ja *Array) Length() int {
	if ja.Values == nil {
		return 0
	}

	return len(ja.Values)
}

//Matches compares two arrays
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
	return ja.toString(true, 0)
}

//InlineString transforms JSON array into inline string
func (ja *Array) InlineString() string {
	return ja.toString(false, 0)
}

func (ja *Array) toString(pretty bool, level int) string {
	if ja.Values == nil || len(ja.Values) == 0 {
		return "[]"
	}

	var sb strings.Builder

	sb.WriteByte('[')
	if pretty {
		sb.WriteByte('\r')
		sb.WriteByte('\n')
	}

	next := level + 1

	for index, jv := range ja.Values {
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

		s := jv.String(pretty, next)
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
