package json

type vBool struct {
	v bool
}

//Bool constructs bool value
func Bool(v bool) Value {
	return vBool{v: v}
}

func (b vBool) t() int {
	return tBool
}

func (b vBool) string(pretty bool, level int) string {
	if b.v {
		return "true"
	}
	return "false"
}

func (b vBool) Bool() (bool, bool) {
	return b.v, true
}

func (b vBool) Int() (int, bool) {
	if b.v {
		return 1, true
	}
	return 0, true
}

func (b vBool) Float() (float64, bool) {
	if b.v {
		return 1, true
	}
	return 0, true
}

func (b vBool) String() (string, bool) {
	if b.v {
		return "true", true
	}
	return "false", true
}

func (b vBool) Object() (Object, bool) {
	return Object{}, false
}

func (b vBool) Array() (Array, bool) {
	return Array{}, false
}

func (b vBool) Copy() Value {
	return Bool(b.v)
}
