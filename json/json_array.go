package json

import (
	"errors"
	"fmt"
	"strings"
)

type JsonArray struct {
	Values []JsonValue
}

func (ja *JsonArray) Copy() *JsonArray {
	if ja == nil {
		return nil
	}

	copy := JsonArray{}

	if ja.Values != nil {
		for _, jv := range ja.Values {
			copy.AddValue(&jv)
		}
	}

	return &copy
}

func (ja *JsonArray) AddValue(value *JsonValue) int {
	if ja.Values == nil {
		ja.Values = []JsonValue{*value}
	} else {
		ja.Values = append(ja.Values, *value)
	}

	return len(ja.Values) - 1
}

func (ja *JsonArray) AddString(value string) int {
	return ja.AddValue(&JsonValue{Value: value, Type: STRING})
}

func (ja *JsonArray) AddInt(value int) int {
	return ja.AddValue(&JsonValue{Value: value, Type: NUMBER})
}

func (ja *JsonArray) AddFloat(value float64) int {
	return ja.AddValue(&JsonValue{Value: value, Type: NUMBER})
}

func (ja *JsonArray) AddBoolean(value bool) int {
	return ja.AddValue(&JsonValue{Value: value, Type: BOOLEAN})
}

func (ja *JsonArray) AddObject(value *JsonObject) int {
	return ja.AddValue(&JsonValue{Value: *value, Type: OBJECT})
}

func (ja *JsonArray) AddArray(value *JsonArray) int {
	return ja.AddValue(&JsonValue{Value: *value, Type: ARRAY})
}

func (ja *JsonArray) Add(value interface{}) (int, error) {
	jv, err := newJsonValue(value)
	if err != nil {
		return len(ja.Values) - 1, err
	}
	return ja.AddValue(jv), nil
}

func (ja *JsonArray) SetValue(index int, value *JsonValue) error {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return errors.New(err)
	}
	ja.Values[index] = *value
	return nil
}

func (ja *JsonArray) Set(index int, value interface{}) error {
	jv, err := newJsonValue(value)
	if err != nil {
		return err
	}
	return ja.SetValue(index, jv)
}

func (ja *JsonArray) Remove(index int) error {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return errors.New(err)
	}

	values := ja.Values
	ja.Values = append(values[:index], values[index+1:]...)

	return nil
}

func (ja *JsonArray) GetValue(index int) (*JsonValue, error) {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return nil, errors.New(err)
	}

	return &ja.Values[index], nil
}

func (ja *JsonArray) IsEmpty() bool {
	if ja.Values == nil || len(ja.Values) == 0 {
		return true
	}

	return false
}

func (ja *JsonArray) Length() int {
	if ja.Values == nil {
		return 0
	}

	return len(ja.Values)
}

func (ja *JsonArray) ToString() string {
	return ja.toString(true, 0)
}

func (ja *JsonArray) ToInlineString() string {
	return ja.toString(false, 0)
}

func (ja *JsonArray) toString(pretty bool, level int) string {
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

		s := jv.toString(pretty, next)
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
