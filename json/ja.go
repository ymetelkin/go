package json

type vArray struct {
	v Array
}

//A constructs Array value
func A(v Array) Value {
	return vArray{v: v}
}

//Strings constructs string Array value
func Strings(vs []string) Value {
	if len(vs) == 0 {
		return nil
	}
	return vArray{v: NewStringArray(vs)}
}

//Ints constructs int Array value
func Ints(vs []int) Value {
	if len(vs) == 0 {
		return nil
	}
	return vArray{v: NewIntArray(vs)}
}

//Floats constructs float64 Array value
func Floats(vs []float64) Value {
	if len(vs) == 0 {
		return nil
	}
	return vArray{v: NewFloatArray(vs)}
}

//Objects constructs Object Array value
func Objects(vs []Object) Value {
	if len(vs) == 0 {
		return nil
	}
	return vArray{v: NewObjectArray(vs)}
}

func (ja vArray) t() int {
	return tArray
}

func (ja vArray) string(pretty bool, level int) string {
	return ja.v.string(pretty, level)
}

func (ja vArray) String() (string, bool) {
	return ja.v.String(), true
}

func (ja vArray) Int() (int, bool) {
	return 0, false
}

func (ja vArray) Float() (float64, bool) {
	return 0, false
}

func (ja vArray) Bool() (bool, bool) {
	return false, false
}

func (ja vArray) Object() (Object, bool) {
	return Object{}, false
}

func (ja vArray) Array() (Array, bool) {
	return ja.v, true
}
