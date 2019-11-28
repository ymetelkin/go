package json

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	jsonObject int = iota
	jsonArray
	jsonString
	jsonInt
	jsonFloat
	jsonBool
	jsonNull
)

type value struct {
	Value interface{}
	Type  int
	Text  string
}

func (jv *value) IsEmpty() bool {
	return jv.Value == nil
}

func newInt(i int) value {
	return value{Value: i, Type: jsonInt}
}

func newFloat(f float64) value {
	return value{Value: f, Type: jsonFloat}
}

func newBool(b bool) value {
	return value{Value: b, Type: jsonBool}
}

func newString(s string) value {
	return value{Value: s, Type: jsonString}
}

func newObject(o Object) value {
	return value{Value: o, Type: jsonObject}
}

func newArray(a Array) value {
	return value{Value: a, Type: jsonArray}
}

func newNull() value {
	return value{Value: nil, Type: jsonNull}
}

func newInts(vs []int) []value {
	if len(vs) == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonInt}
	}
	return values
}

func newFloats(vs []float64) []value {
	if len(vs) == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonFloat}
	}
	return values
}

func newBools(vs []bool) []value {
	if len(vs) == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonBool}
	}
	return values
}

func newStrings(vs []string) []value {
	if len(vs) == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonString}
	}
	return values
}

func newObjects(vs []Object) []value {
	if len(vs) == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonObject}
	}
	return values
}

func newArrays(vs []Array) []value {
	if len(vs) == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonArray}
	}
	return values
}

func (jv *value) GetInt() (v int, ok bool) {
	if jv.Type == jsonInt {
		v, ok = jv.Value.(int)
		if !ok {
			u, k := jv.Value.(uint)
			if k {
				v = int(u)
				ok = true
			}
		}
	} else if jv.Type == jsonFloat {
		f, k := jv.Value.(float64)
		if k {
			v = int(f)
			ok = true
		}
	} else if jv.Type == jsonString {
		s, k := jv.Value.(string)
		if k {
			if strings.Contains(s, ".") {
				f, err := strconv.ParseFloat(s, 64)
				if err == nil {
					v = int(f)
					ok = true
				}
			} else {
				i, err := strconv.ParseInt(s, 0, 64)
				if err == nil {
					v = int(i)
					ok = true
				}
			}
		}
	}
	return
}

func (jv *value) GetFloat() (v float64, ok bool) {
	if jv.Type == jsonFloat {
		v, ok = jv.Value.(float64)
	} else if jv.Type == jsonInt {
		i, k := jv.Value.(int)
		if k {
			v = float64(i)
			ok = true
		}
	} else if jv.Type == jsonString {
		s, k := jv.Value.(string)
		if k {
			f, err := strconv.ParseFloat(s, 64)
			ok = err == nil
			v = f
		}
	}

	return
}

func (jv *value) GetString() (v string, ok bool) {
	if jv.Type == jsonString {
		v, ok = jv.Value.(string)
	} else {
		v = jv.String(true, 0)
		ok = true
	}
	return
}

func (jv *value) GetBool() (v bool, ok bool) {
	if jv.Type == jsonBool {
		v, ok = jv.Value.(bool)
	} else if jv.Type == jsonString {
		s, ok := jv.Value.(string)
		if ok {
			b, err := strconv.ParseBool(s)
			ok = err == nil
			v = b
		}
	}
	return
}

func (jv *value) GetObject() (v Object, ok bool) {
	if jv.Type == jsonObject {
		v, ok = jv.Value.(Object)
	}
	return
}

func (jv *value) GetArray() (v Array, ok bool) {
	if jv.Type == jsonArray {
		v, ok = jv.Value.(Array)
	}
	return
}

func (jv *value) Matches(other *value) (match bool, s string) {
	tp := jv.Type
	if tp != other.Type {
		if (tp == jsonInt && other.Type == jsonFloat) || (tp == jsonFloat && other.Type == jsonInt) {
			tp = jsonFloat
		} else {
			s = fmt.Sprintf("Type mismatch: [ %v ] vs [ %v ]", jv.Type, other.Type)
			return
		}
	}

	switch tp {
	case jsonString:
		lv, _ := jv.GetString()
		rv, _ := other.GetString()
		if lv != rv {
			s = fmt.Sprintf("String mismatch: [ %v ] vs [ %v ]", lv, rv)
			return
		}
	case jsonInt:
		lv, _ := jv.GetInt()
		rv, _ := other.GetInt()
		if lv != rv {
			s = fmt.Sprintf("Integer mismatch: [ %v ] vs [ %v ]", lv, rv)
			return
		}
	case jsonBool:
		lv, _ := jv.GetBool()
		rv, _ := other.GetBool()
		if lv != rv {
			s = fmt.Sprintf("Boolean mismatch: [ %v ] vs [ %v ]", lv, rv)
			return
		}
	case jsonFloat:
		lv, _ := jv.GetFloat()
		rv, _ := other.GetFloat()
		if lv != rv {
			s = fmt.Sprintf("Float mismatch: [ %v ] vs [ %v ]", lv, rv)
			return
		}
	case jsonObject:
		lv, _ := jv.GetObject()
		rv, _ := other.GetObject()
		return lv.Matches(&rv)
	case jsonArray:
		lv, _ := jv.GetArray()
		rv, _ := other.GetArray()
		return lv.Matches(&rv)
	}

	match = true
	return
}

func (jv *value) String(pretty bool, level int) string {
	if jv.Value == nil {
		return "null"
	}

	if jv.Text != "" {
		return jv.Text
	}

	switch jv.Type {
	case jsonString:
		s, ok := jv.Value.(string)
		if ok {
			bytes := []byte(s)
			var sb strings.Builder
			sb.WriteByte('"')

			for _, c := range bytes {
				switch c {
				case '"':
					sb.WriteByte('\\')
					sb.WriteByte('"')
				case '\\':
					sb.WriteByte('\\')
					sb.WriteByte('\\')
				case '\r':
					sb.WriteByte('\\')
					sb.WriteByte('r')
				case '\n':
					sb.WriteByte('\\')
					sb.WriteByte('n')
				case '\t':
					sb.WriteByte('\\')
					sb.WriteByte('t')
				case '\b':
					sb.WriteByte('\\')
					sb.WriteByte('b')
				case '\f':
					sb.WriteByte('\\')
					sb.WriteByte('f')
				case '\v':
					sb.WriteByte(' ')
				default:
					sb.WriteByte(c)
				}
			}

			sb.WriteByte('"')
			return sb.String()
		}
	case jsonInt:
		i, ok := jv.Value.(int)
		if ok {
			return strconv.Itoa(i)
		}
	case jsonFloat:
		f, ok := jv.Value.(float64)
		if ok {
			return strconv.FormatFloat(f, 'f', -1, 64)
		}
	case jsonBool:
		b, ok := jv.Value.(bool)
		if ok {
			return strconv.FormatBool(b)
		}
	case jsonObject:
		jo, ok := jv.Value.(Object)
		if ok {
			return jo.toString(pretty, level)
		}
	case jsonArray:
		ja, ok := jv.Value.(Array)
		if ok {
			return ja.toString(pretty, level)
		}
	}

	return "null"
}
