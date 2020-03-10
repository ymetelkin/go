package json

const (
	tObject int = iota + 1
	tArray
	tString
	tInt
	tFloat
	tBool
)

//Value interface
type Value interface {
	t() int
	string(pretty bool, level int) string
	String() (string, bool)
	Int() (int, bool)
	Float() (float64, bool)
	Bool() (bool, bool)
	Object() (Object, bool)
	Array() (Array, bool)
	Copy() Value
}

func compare(left Value, right Value) bool {
	t := left.t()
	if t != right.t() {
		if (t == tInt && right.t() == tFloat) || (t == tFloat && right.t() == tInt) {
			t = tFloat
		} else {
			return false
		}
	}

	switch t {
	case tString:
		l, _ := left.String()
		r, _ := right.String()
		if l != r {
			return false
		}
	case tInt:
		l, _ := left.Int()
		r, _ := right.Int()
		if l != r {
			return false
		}
	case tBool:
		l, _ := left.Bool()
		r, _ := right.Bool()
		if l != r {
			return false
		}
	case tFloat:
		l, _ := left.Float()
		r, _ := right.Float()
		if l != r {
			return false
		}
	case tObject:
		l, _ := left.Object()
		r, _ := right.Object()
		return l.Equals(&r)
	case tArray:
		l, _ := left.Array()
		r, _ := right.Array()
		return l.Equals(&r)
	}

	return true
}
