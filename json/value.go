package json

import (
	"errors"
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
	jsonParams
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
	if vs == nil {
		return nil
	}
	size := len(vs)
	if size == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonInt}
	}
	return values
}

func newFloats(vs []float64) []value {
	if vs == nil {
		return nil
	}
	size := len(vs)
	if size == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonFloat}
	}
	return values
}

func newBools(vs []bool) []value {
	if vs == nil {
		return nil
	}
	size := len(vs)
	if size == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonBool}
	}
	return values
}

func newStrings(vs []string) []value {
	if vs == nil {
		return nil
	}
	size := len(vs)
	if size == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonString}
	}
	return values
}

func newObjects(vs []Object) []value {
	if vs == nil {
		return nil
	}
	size := len(vs)
	if size == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonObject}
	}
	return values
}

func newArrays(vs []Array) []value {
	if vs == nil {
		return nil
	}
	size := len(vs)
	if size == 0 {
		return nil
	}
	values := make([]value, len(vs))
	for i, v := range vs {
		values[i] = value{Value: v, Type: jsonArray}
	}
	return values
}

func newParameterizedString(ps ParameterizedString) value {
	return value{Value: ps, Type: jsonParams}
}

func (jv *value) GetInt() (int, error) {
	if jv.Type == jsonInt {
		i, ok := jv.Value.(int)
		if ok {
			return i, nil
		}
		u, ok := jv.Value.(uint)
		if ok {
			return int(u), nil
		}
	} else if jv.Type == jsonFloat {
		f, ok := jv.Value.(float64)
		if ok {
			return int(f), nil
		}
	} else if jv.Type == jsonString {
		s, ok := jv.Value.(string)
		if ok {
			if strings.Contains(s, ".") {
				f, err := strconv.ParseFloat(s, 64)
				return int(f), err
			}
			i, err := strconv.ParseInt(s, 0, 64)
			return int(i), err
		}
		return 0, errors.New("Cannot read string value")
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return 0, errors.New(err)
}

func (jv *value) GetFloat() (float64, error) {
	if jv.Type == jsonFloat || jv.Type == jsonInt {
		f, ok := jv.Value.(float64)
		if ok {
			return f, nil
		}
		err := fmt.Sprintf("Unsupported integer type: %T", jv.Value)
		return 0, errors.New(err)
	} else if jv.Type == jsonString {
		s, ok := jv.Value.(string)
		if ok {
			f, err := strconv.ParseFloat(s, 64)
			return f, err
		}
		return 0, errors.New("Cannot read string value")
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return 0, errors.New(err)
}

func (jv *value) GetString() (string, error) {
	if jv.Type == jsonString {
		s, ok := jv.Value.(string)
		if ok {
			return s, nil
		}
		return "", errors.New("Cannot read string value")
	}
	return jv.ToString(true, 0), nil
}

func (jv *value) GetBool() (bool, error) {
	if jv.Type == jsonBool {
		b, ok := jv.Value.(bool)
		if ok {
			return b, nil
		}
		return false, errors.New("Cannot read string value")
	} else if jv.Type == jsonString {
		s, ok := jv.Value.(string)
		if ok {
			b, err := strconv.ParseBool(s)
			return b, err
		}
		return false, errors.New("Cannot read string value")
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return false, errors.New(err)
}

func (jv *value) GetObject() (Object, error) {
	if jv.Type == jsonObject {
		jo, ok := jv.Value.(Object)
		if ok {
			return jo, nil
		}
		return Object{}, errors.New("Cannot read Object value")
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return Object{}, errors.New(err)
}

func (jv *value) GetArray() (Array, error) {
	if jv.Type == jsonArray {
		ja, ok := jv.Value.(Array)
		if ok {
			return ja, nil
		}
		return Array{}, errors.New("Cannot read Array value")
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return Array{}, errors.New(err)
}

func (jv *value) GetParameterizedString() (ParameterizedString, error) {
	if jv.Type == jsonParams {
		ps, ok := jv.Value.(ParameterizedString)
		if ok {
			return ps, nil
		}
		return ParameterizedString{}, errors.New("Cannot read string value")
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return ParameterizedString{}, errors.New(err)
}

func (jv *value) ToString(pretty bool, level int) string {
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
			runes := []rune(s)
			var sb strings.Builder
			sb.WriteRune(tokenQUOTE)

			for _, r := range runes {
				switch r {
				case tokenQUOTE:
					sb.WriteRune(tokenBackslash)
					sb.WriteRune(tokenQUOTE)
				case tokenBackslash:
					sb.WriteRune(tokenBackslash)
					sb.WriteRune(tokenBackslash)
				case tokenCR:
					sb.WriteRune(tokenBackslash)
					sb.WriteRune(tokenR)
				case tokenLF:
					sb.WriteRune(tokenBackslash)
					sb.WriteRune(tokenN)
				case tokenHT:
					sb.WriteRune(tokenBackslash)
					sb.WriteRune(tokenT)
				case tokenBS:
					sb.WriteRune(tokenBackslash)
					sb.WriteRune(tokenB)
				case tokenFF:
					sb.WriteRune(tokenBackslash)
					sb.WriteRune(tokenF)
				case tokenVT:
					sb.WriteRune(tokenSPACE)
				default:
					sb.WriteRune(r)
				}
			}

			sb.WriteRune(tokenQUOTE)
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
