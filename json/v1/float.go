package v1

import (
	"fmt"
)

type vFloat struct {
	v float64
}

//Float constructs float64 value
func Float(v float64) Value {
	return vFloat{v: v}
}

func (f vFloat) t() int {
	return tFloat
}

func (f vFloat) string(pretty bool, level int) string {
	return fmt.Sprintf("%v", f.v)
}

func (f vFloat) Float() (float64, bool) {
	return f.v, true
}

func (f vFloat) Int() (int, bool) {
	return int(f.v), true
}

func (f vFloat) String() (string, bool) {
	return fmt.Sprintf("%v", f.v), true
}

func (f vFloat) Bool() (bool, bool) {
	if f.v == 1 {
		return true, true
	}
	if f.v == 0 {
		return false, true
	}
	return false, false
}

func (f vFloat) Object() (Object, bool) {
	return Object{}, false
}

func (f vFloat) Array() (Array, bool) {
	return Array{}, false
}

func (f vFloat) Copy() Value {
	return Float(f.v)
}
