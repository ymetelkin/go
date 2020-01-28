package json

type vObject struct {
	v Object
}

//O constructs Object value
func O(v Object) Value {
	return vObject{v: v}
}

func (jo vObject) t() int {
	return tObject
}

func (jo vObject) string(pretty bool, level int) string {
	return jo.v.string(pretty, level)
}

func (jo vObject) String() (string, bool) {
	return jo.v.String(), true
}

func (jo vObject) Int() (int, bool) {
	return 0, false
}

func (jo vObject) Float() (float64, bool) {
	return 0, false
}

func (jo vObject) Bool() (bool, bool) {
	return false, false
}

func (jo vObject) Object() (Object, bool) {
	return jo.v, true
}

func (jo vObject) Array() (Array, bool) {
	return Array{}, false
}
