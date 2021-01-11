package v1

import "strconv"

type vInt struct {
	v int
}

//Int constructs integer value
func Int(v int) Value {
	return vInt{v: v}
}

func (i vInt) t() int {
	return tInt
}

func (i vInt) string(pretty bool, level int) string {
	return strconv.Itoa(i.v)
}

func (i vInt) Int() (int, bool) {
	return i.v, true
}

func (i vInt) String() (string, bool) {
	return strconv.Itoa(i.v), true
}

func (i vInt) Float() (float64, bool) {
	return float64(i.v), true
}

func (i vInt) Bool() (bool, bool) {
	if i.v == 1 {
		return true, true
	}
	if i.v == 0 {
		return false, true
	}
	return false, false
}

func (i vInt) Object() (Object, bool) {
	return Object{}, false
}

func (i vInt) Array() (Array, bool) {
	return Array{}, false
}

func (i vInt) Copy() Value {
	return Int(i.v)
}
