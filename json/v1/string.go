package v1

import (
	"strconv"
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

	return strconv.Quote(s.v)
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

func (s vString) Copy() Value {
	return String(s.v)
}
