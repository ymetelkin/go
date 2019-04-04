package json

import (
	"errors"
	"fmt"
	"strings"
)

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

func (ja *Array) AddString(value string) int {
	return ja.addValue(newString(value))
}

func (ja *Array) AddInt(value int) int {
	return ja.addValue(newInt(value))
}

func (ja *Array) AddFloat(value float64) int {
	return ja.addValue(newFloat(value))
}

func (ja *Array) AddBool(value bool) int {
	return ja.addValue(newBool(value))
}

func (ja *Array) AddObject(value *Object) int {
	return ja.addValue(newObject(value))
}

func (ja *Array) AddArray(value *Array) int {
	return ja.addValue(newArray(value))
}

func (ja *Array) setValue(index int, value value) error {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return errors.New(err)
	}
	ja.Values[index] = value
	return nil
}

func (ja *Array) SetInt(index int, value int) error {
	return ja.setValue(index, newInt(value))
}

func (ja *Array) SetFloat(index int, value float64) error {
	return ja.setValue(index, newFloat(value))
}

func (ja *Array) SetBool(index int, value bool) error {
	return ja.setValue(index, newBool(value))
}

func (ja *Array) SetString(index int, value string) error {
	return ja.setValue(index, newString(value))
}

func (ja *Array) SetObject(index int, value *Object) error {
	return ja.setValue(index, newObject(value))
}

func (ja *Array) SetArray(index int, value *Array) error {
	return ja.setValue(index, newArray(value))
}

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

func (ja *Array) GetObject(index int) (*Object, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return nil, err
	}

	obj, err := jv.GetObject()
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (ja *Array) GetArray(index int) (*Array, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return nil, err
	}

	a, err := jv.GetArray()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (ja *Array) IsEmpty() bool {
	if ja.Values == nil || len(ja.Values) == 0 {
		return true
	}

	return false
}

func (ja *Array) Length() int {
	if ja.Values == nil {
		return 0
	}

	return len(ja.Values)
}

func (ja *Array) ToString() string {
	return ja.toString(true, 0)
}

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
