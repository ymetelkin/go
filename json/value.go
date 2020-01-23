package json

import (
	"fmt"
	"strconv"
	"strings"
)

//Value types
const (
	TypeNull int = iota
	TypeObject
	TypeArray
	TypeString
	TypeInt
	TypeFloat
	TypeBool
)

//Value represents JSON value
type Value struct {
	Type int
	data interface{}
	text string
}

//NewObject inits Object value
func NewObject(data Object) Value {
	return Value{
		Type: TypeObject,
		data: data,
	}
}

//NewArray inits Array value
func NewArray(data Array) Value {
	return Value{
		Type: TypeArray,
		data: data,
	}
}

//NewString inits string value
func NewString(data string) Value {
	return Value{
		Type: TypeString,
		data: data,
	}
}

//NewInt inits int value
func NewInt(data int) Value {
	return Value{
		Type: TypeInt,
		data: data,
	}
}

//NewFloat inits float64 value
func NewFloat(data float64) Value {
	return Value{
		Type: TypeFloat,
		data: data,
	}
}

//NewBool inits bool value
func NewBool(data bool) Value {
	return Value{
		Type: TypeBool,
		data: data,
	}
}

//Int gets data as int
func (v *Value) Int() (data int, ok bool) {
	switch v.Type {
	case TypeInt:
		data, ok = v.data.(int)
	case TypeFloat:
		f, k := v.data.(float64)
		if k {
			data = int(f)
			ok = true
		}
	case TypeString:
		s, k := v.data.(string)
		if k {
			if strings.Contains(s, ".") {
				f, err := strconv.ParseFloat(s, 64)
				if err == nil {
					data = int(f)
					ok = true
				}
			} else {
				i, err := strconv.ParseInt(s, 0, 64)
				if err == nil {
					data = int(i)
					ok = true
				}
			}
		}
	}
	return
}

//Float gets data as float64
func (v *Value) Float() (data float64, ok bool) {
	switch v.Type {
	case TypeFloat:
		data, ok = v.data.(float64)
	case TypeInt:
		i, k := v.data.(int)
		if k {
			data = float64(i)
			ok = true
		}
	case TypeString:
		s, k := v.data.(string)
		if k {
			f, err := strconv.ParseFloat(s, 64)
			ok = err == nil
			data = f
		}
	}
	return
}

//String gets data as string
func (v *Value) String() (data string, ok bool) {
	switch v.Type {
	case TypeString:
		data, ok = v.data.(string)
	case TypeInt:
		i, k := v.data.(int)
		if k {
			data = strconv.Itoa(i)
			ok = true
		}
	case TypeFloat:
		f, k := v.data.(float64)
		if k {
			data = strconv.FormatFloat(f, 'f', -1, 64)
			ok = true
		}
	case TypeBool:
		b, k := v.data.(bool)
		if k {
			data = strconv.FormatBool(b)
			ok = true
		}
	case TypeObject:
		jo, k := v.data.(Object)
		if k {
			data = jo.string(false, 0)
			ok = true
		}
	case TypeArray:
		ja, k := v.data.(Array)
		if k {
			data = ja.string(false, 0)
			ok = true
		}
	case TypeNull:
		data = "null"
		ok = true
	}
	return
}

//Bool get data as bool
func (v *Value) Bool() (data bool, ok bool) {
	switch v.Type {
	case TypeBool:
		data, ok = v.data.(bool)
	case TypeString:
		s, k := v.data.(string)
		if k {
			b, err := strconv.ParseBool(s)
			ok = err == nil
			data = b
		}
	}
	return
}

//Object gets data as
func (v *Value) Object() (data Object, ok bool) {
	if v.Type == TypeObject {
		data, ok = v.data.(Object)
	}
	return
}

//Array gets data as array
func (v *Value) Array() (data Array, ok bool) {
	if v.Type == TypeArray {
		data, ok = v.data.(Array)
	}
	return
}

//Matches compares value to another value
func (v *Value) Matches(other *Value) (match bool, s string) {
	tp := v.Type
	if tp != other.Type {
		if (tp == TypeInt && other.Type == TypeFloat) || (tp == TypeFloat && other.Type == TypeInt) {
			tp = TypeFloat
		} else {
			s = fmt.Sprintf("Type mismatch: [ %v ] vs [ %v ]", v.Type, other.Type)
			return
		}
	}

	switch tp {
	case TypeString:
		l, _ := v.String()
		r, _ := other.String()
		if l != r {
			s = fmt.Sprintf("String mismatch: [ %v ] vs [ %v ]", l, r)
			return
		}
	case TypeInt:
		l, _ := v.Int()
		r, _ := other.Int()
		if l != r {
			s = fmt.Sprintf("Integer mismatch: [ %v ] vs [ %v ]", l, r)
			return
		}
	case TypeBool:
		l, _ := v.Bool()
		r, _ := other.Bool()
		if l != r {
			s = fmt.Sprintf("Boolean mismatch: [ %v ] vs [ %v ]", l, r)
			return
		}
	case TypeFloat:
		l, _ := v.Float()
		r, _ := other.Float()
		if l != r {
			s = fmt.Sprintf("Float mismatch: [ %v ] vs [ %v ]", l, r)
			return
		}
	case TypeObject:
		l, _ := v.Object()
		r, _ := other.Object()
		return l.Matches(&r)
	case TypeArray:
		l, _ := v.Array()
		r, _ := other.Array()
		return l.Matches(&r)
	}

	match = true
	return
}

func (v *Value) string(pretty bool, level int) string {
	if v.data == nil {
		return "null"
	}

	if v.text != "" {
		return v.text
	}

	switch v.Type {
	case TypeString:
		s, ok := v.data.(string)
		if ok {
			var (
				bs = []byte(s)
				sb strings.Builder
			)

			sb.WriteByte('"')

			for _, c := range bs {
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
	case TypeInt, TypeFloat, TypeBool:
		s, ok := v.String()
		if ok {
			return s
		}
	case TypeObject:
		jo, ok := v.data.(Object)
		if ok {
			return jo.string(pretty, level)
		}
	case TypeArray:
		ja, ok := v.data.(Array)
		if ok {
			return ja.string(pretty, level)
		}
	}

	return "null"
}
