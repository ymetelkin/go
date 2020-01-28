package json

import (
	"strings"
)

//Array JSON array
type Array struct {
	Values []Value
	params []int
}

//NewArray constructs Array
func NewArray(vs []Value) Array {
	return Array{
		Values: vs,
	}
}

//NewStringArray constructs string Array
func NewStringArray(vs []string) (ja Array) {
	if len(vs) == 0 {
		return
	}
	ja.Values = make([]Value, len(vs))
	for i, v := range vs {
		ja.Values[i] = String(v)
	}
	return
}

//NewIntArray constructs int Array
func NewIntArray(vs []int) (ja Array) {
	if len(vs) == 0 {
		return
	}
	ja.Values = make([]Value, len(vs))
	for i, v := range vs {
		ja.Values[i] = Int(v)
	}
	return
}

//NewFloatArray constructs float64 Array
func NewFloatArray(vs []float64) (ja Array) {
	if len(vs) == 0 {
		return
	}
	ja.Values = make([]Value, len(vs))
	for i, v := range vs {
		ja.Values[i] = Float(v)
	}
	return
}

//NewObjectArray constructs Object Array
func NewObjectArray(vs []Object) (ja Array) {
	if len(vs) == 0 {
		return
	}
	ja.Values = make([]Value, len(vs))
	for i, v := range vs {
		ja.Values[i] = O(v)
	}
	return
}

//Add adds value to array
func (ja *Array) Add(v Value) {
	ja.Values = append(ja.Values, v)
}

//GetStrings gets string values
func (ja *Array) GetStrings() ([]string, bool) {
	vs := make([]string, len(ja.Values))
	for i, jv := range ja.Values {
		v, ok := jv.String()
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
		v, ok := jv.Int()
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
		v, ok := jv.Float()
		if !ok {
			return vs, false
		}
		vs[i] = v
	}
	return vs, true
}

//GetObjects gets Object values
func (ja *Array) GetObjects() ([]Object, bool) {
	vs := make([]Object, len(ja.Values))
	for i, jv := range ja.Values {
		v, ok := jv.Object()
		if !ok {
			return vs, false
		}
		vs[i] = v
	}
	return vs, true
}

//Equals compares two arrays
func (ja *Array) Equals(other *Array) bool {
	if ja.Values == nil || other.Values == nil || len(ja.Values) != len(other.Values) {
		return false
	}

	for _, lv := range ja.Values {
		var ok bool
		for _, rv := range other.Values {
			ok = compare(lv, rv)
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
