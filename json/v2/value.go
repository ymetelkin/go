package v2

import (
	"errors"
)

//ValueType is value type
type ValueType int

//Value types
const (
	ObjectType ValueType = iota + 1
	ArrayType
	StringType
	IntType
	FloatType
	BoolType
	NullType
)

//Value JSON value interface
type Value interface {
	Value() interface{}
	Type() ValueType
	String() string
	//Copy() Value
}

func (p *byteParser) ParseValue(parameterized bool) (Value, []Parameter, error) {
	err := p.SkipWS()
	if err != nil {
		return nil, nil, err
	}

	switch p.Byte {
	case '"':
		return p.ParseString(parameterized)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '+':
		v, err := p.ParseNumber()
		return v, nil, err
	case '{':
		jo, err := p.ParseObject(parameterized)
		return jo, nil, err
	case '[':
		ja, err := p.ParseArray(parameterized)
		return ja, nil, err
	case 't':
		v, err := p.ParseTrue()
		return v, nil, err
	case 'f':
		v, err := p.ParseFalse()
		return v, nil, err
	case 'n':
		v, err := p.ParseNull()
		return v, nil, err
	}

	return nil, nil, errors.New("invalid JSON")
}

//ObjectValue returns underlying value as *Object
func ObjectValue(v Value) (*Object, bool) {
	if v.Type() != ObjectType {
		return nil, false
	}
	val, ok := v.(*Object)
	return val, ok
}

//ArrayValue returns underlying value as *Array
func ArrayValue(v Value) (*Array, bool) {
	if v.Type() != ArrayType {
		return nil, false
	}
	val, ok := v.(*Array)
	return val, ok
}

//StringValue returns underlying value as string
func StringValue(v Value) (string, bool) {
	if v.Type() != StringType {
		return "", false
	}
	val, ok := (v.Value()).(string)
	return val, ok
}

//IntValue returns underlying value as string
func IntValue(v Value) (int, bool) {
	if v.Type() != IntType {
		return 0, false
	}
	val, ok := (v.Value()).(int)
	return val, ok
}

//FloatValue returns underlying value as float64
func FloatValue(v Value) (float64, bool) {
	if v.Type() != FloatType {
		return 0, false
	}
	val, ok := (v.Value()).(float64)
	return val, ok
}

//BoolValue returns underlying value as string
func BoolValue(v Value) (bool, bool) {
	if v.Type() != BoolType {
		return false, false
	}
	val, ok := (v.Value()).(bool)
	return val, ok
}

//NullValue checks if underlying value is nil
func NullValue(v Value) bool {
	return v.Type() == NullType
}

func copyValue(v Value) (Value, bool) {
	switch v.Type() {
	case StringType, IntType, FloatType, BoolType:
		return v, true
	case ObjectType:
		jo, ok := ObjectValue(v)
		if !ok {
			return nil, false
		}
		copy := jo.Copy()
		return copy, true
	case ArrayType:
		ja, ok := ArrayValue(v)
		if !ok {
			return nil, false
		}
		copy := ja.Copy()
		return copy, true
	}

	return nil, false

}

func compareValues(left Value, right Value) bool {
	lt := left.Type()
	rt := right.Type()
	if lt != rt {
		if (lt == IntType && rt == FloatType) || (lt == FloatType && rt == IntType) {
			lt = FloatType
		} else {
			return false
		}
	}

	switch lt {
	case StringType:
		l, _ := StringValue(left)
		r, _ := StringValue(right)
		if l != r {
			return false
		}
	case IntType:
		l, _ := IntValue(left)
		r, _ := IntValue(right)
		if l != r {
			return false
		}
	case BoolType:
		l, _ := BoolValue(left)
		r, _ := BoolValue(right)
		if l != r {
			return false
		}
	case FloatType:
		l, _ := FloatValue(left)
		r, _ := FloatValue(right)
		if l != r {
			return false
		}
	case ObjectType:
		l, _ := ObjectValue(left)
		r, _ := ObjectValue(right)
		return l != nil && r != nil && l.Equals(r)
	case ArrayType:
		l, _ := ArrayValue(left)
		r, _ := ArrayValue(right)
		return l != nil && r != nil && l.Equals(r)
	}

	return true
}
