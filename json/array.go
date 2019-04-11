package json

import (
	"errors"
	"fmt"
	"strings"
)

//Array is json array structure
type Array struct {
	Values []value
}

func (ja *Array) addValue(jv value) int {
	if ja.Values == nil {
		ja.Values = []value{jv}
	} else {
		ja.Values = append(ja.Values, jv)
	}

	return len(ja.Values) - 1
}

func (ja *Array) addValues(jvs []value) int {
	if jvs != nil {
		if ja.Values == nil {
			ja.Values = jvs
		} else {
			ja.Values = append(ja.Values, jvs...)
		}
	}

	if ja.Values == nil {
		return 0
	}

	return len(ja.Values) - 1
}

//AddString adds string to json array
func (ja *Array) AddString(value string) int {
	return ja.addValue(newString(value))
}

//AddInt adds int to json array
func (ja *Array) AddInt(value int) int {
	return ja.addValue(newInt(value))
}

//AddFloat adds float to json array
func (ja *Array) AddFloat(value float64) int {
	return ja.addValue(newFloat(value))
}

//AddBool adds bool to json array
func (ja *Array) AddBool(value bool) int {
	return ja.addValue(newBool(value))
}

//AddObject adds json object to json array
func (ja *Array) AddObject(value Object) int {
	return ja.addValue(newObject(value))
}

//AddArray adds json array to json array
func (ja *Array) AddArray(value Array) int {
	return ja.addValue(newArray(value))
}

//AddStrings adds strings to json array
func (ja *Array) AddStrings(values []string) int {
	return ja.addValues(newStrings(values))
}

//AddInts adds ints to json array
func (ja *Array) AddInts(values []int) int {
	return ja.addValues(newInts(values))
}

//AddFloats adds floats to json array
func (ja *Array) AddFloats(values []float64) int {
	return ja.addValues(newFloats(values))
}

//AddBools adds bools to json array
func (ja *Array) AddBools(values []bool) int {
	return ja.addValues(newBools(values))
}

//AddObjects adds json objects to json array
func (ja *Array) AddObjects(values []Object) int {
	return ja.addValues(newObjects(values))
}

//AddArrays adds json arrays to json array
func (ja *Array) AddArrays(values []Array) int {
	return ja.addValues(newArrays(values))
}

func (ja *Array) setValue(index int, value value) error {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return errors.New(err)
	}
	ja.Values[index] = value
	return nil
}

//SetInt adds int json array at specified position
func (ja *Array) SetInt(index int, value int) error {
	return ja.setValue(index, newInt(value))
}

//SetFloat sets float in json array at specified position
func (ja *Array) SetFloat(index int, value float64) error {
	return ja.setValue(index, newFloat(value))
}

//SetBool sets bool in json array at specified position
func (ja *Array) SetBool(index int, value bool) error {
	return ja.setValue(index, newBool(value))
}

//SetString sets string in json array at specified position
func (ja *Array) SetString(index int, value string) error {
	return ja.setValue(index, newString(value))
}

//SetObject sets json object in json array at specified position
func (ja *Array) SetObject(index int, value Object) error {
	return ja.setValue(index, newObject(value))
}

//SetArray sets json array in json array at specified position
func (ja *Array) SetArray(index int, value Array) error {
	return ja.setValue(index, newArray(value))
}

//Remove removes element from json array at specified position
func (ja *Array) Remove(index int) error {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return errors.New(err)
	}

	values := ja.Values
	ja.Values = append(values[:index], values[index+1:]...)

	return nil
}

func (ja *Array) getValue(index int) (*value, error) {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return nil, errors.New(err)
	}

	return &ja.Values[index], nil
}

//GetString gets string from json array at specified position
func (ja *Array) GetString(index int) (string, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return "", err
	}

	s, err := jv.GetString()
	if err != nil {
		return "", err
	}

	return s, nil
}

//GetInt gets int from json array at specified position
func (ja *Array) GetInt(index int) (int, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return 0, err
	}

	i, err := jv.GetInt()
	if err != nil {
		return 0, err
	}

	return i, nil
}

//GetFloat gets bool from json array at specified position
func (ja *Array) GetFloat(index int) (float64, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return 0, err
	}

	f, err := jv.GetFloat()
	if err != nil {
		return 0, err
	}

	return f, nil
}

//GetBool gets bool from json array at specified position
func (ja *Array) GetBool(index int) (bool, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return false, err
	}

	b, err := jv.GetBool()
	if err != nil {
		return false, err
	}

	return b, nil
}

//GetObject gets json object from json array at specified position
func (ja *Array) GetObject(index int) (Object, error) {
	jv, err := ja.getValue(index)
	if err == nil {
		jo, err := jv.GetObject()
		if err == nil {
			return jo, nil
		}
	}

	return Object{}, err
}

//GetArray gets json array from json array at specified position
func (ja *Array) GetArray(index int) (Array, error) {
	jv, err := ja.getValue(index)
	if err == nil {
		a, err := jv.GetArray()
		if err == nil {
			return a, nil
		}
	}

	return Array{}, err
}

//IsEmpty checks if json array has any elements
func (ja *Array) IsEmpty() bool {
	if ja.Values == nil || len(ja.Values) == 0 {
		return true
	}

	return false
}

//Length returns number of elements in json array
func (ja *Array) Length() int {
	if ja.Values == nil {
		return 0
	}

	return len(ja.Values)
}

//ToString transforms json array to pretty string
func (ja *Array) ToString() string {
	return ja.toString(true, 0)
}

//ToInlineString transforms json array into inline string
func (ja *Array) ToInlineString() string {
	return ja.toString(false, 0)
}

func (ja *Array) toString(pretty bool, level int) string {
	if ja.Values == nil || len(ja.Values) == 0 {
		return "[]"
	}

	var sb strings.Builder

	sb.WriteRune(TOKEN_LEFT_SQUARE)
	if pretty {
		sb.WriteRune(TOKEN_CR)
		sb.WriteRune(TOKEN_LF)
	}

	next := level + 1

	for index, jv := range ja.Values {
		if index > 0 {
			sb.WriteRune(TOKEN_COMMA)

			if pretty {
				sb.WriteRune(TOKEN_CR)
				sb.WriteRune(TOKEN_LF)
			}
		}

		if pretty {
			i := 0
			for i <= level {
				sb.WriteRune(TOKEN_SPACE)
				sb.WriteRune(TOKEN_SPACE)
				i++
			}
		}

		s := jv.ToString(pretty, next)
		sb.WriteString(s)
	}

	if pretty {
		sb.WriteRune(TOKEN_CR)
		sb.WriteRune(TOKEN_LF)
		i := 0
		for i < level {
			sb.WriteRune(TOKEN_SPACE)
			sb.WriteRune(TOKEN_SPACE)
			i++
		}
	}
	sb.WriteRune(TOKEN_RIGHT_SQUARE)

	return sb.String()
}
