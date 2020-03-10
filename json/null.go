package json

type vNull struct {
}

//Null constructs null value
func Null() Value {
	return vNull{}
}

func (s vNull) t() int {
	return 0
}

func (s vNull) string(pretty bool, level int) string {
	return "null"
}

func (s vNull) String() (string, bool) {
	return "null", true
}

func (s vNull) Int() (int, bool) {
	return 0, false
}

func (s vNull) Float() (float64, bool) {
	return 0, false
}

func (s vNull) Bool() (bool, bool) {
	return false, false
}

func (s vNull) Object() (Object, bool) {
	return Object{}, false
}

func (s vNull) Array() (Array, bool) {
	return Array{}, false
}

func (s vNull) Copy() Value {
	return Null()
}
