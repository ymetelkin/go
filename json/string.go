package json

import (
	"strconv"
	"strings"
)

type vString struct {
	v string
}

//String constructs string value
func String(v string) Value {
	return vString{v: v}
}

func (s vString) t() int {
	return tString
}

func (s vString) string(pretty bool, level int) string {
	if s.v == "" {
		return "\"\""
	}

	var (
		bs = []byte(s.v)
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

func (s vString) String() (string, bool) {
	return s.v, true
}

func (s vString) Int() (int, bool) {
	i, e := strconv.Atoi(s.v)
	return i, e == nil
}

func (s vString) Float() (float64, bool) {
	f, e := strconv.ParseFloat(s.v, 64)
	return f, e == nil
}

func (s vString) Bool() (bool, bool) {
	if s.v == "true" {
		return true, true
	}
	if s.v == "false" {
		return false, true
	}
	return false, false
}

func (s vString) Object() (Object, bool) {
	return Object{}, false
}

func (s vString) Array() (Array, bool) {
	return Array{}, false
}
