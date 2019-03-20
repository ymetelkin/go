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
			copy.Add(jv)
		}
	}

	return &copy
}

func (ja *JsonArray) Add(value interface{}) (int, error) {
	jv, err := NewJsonValue(value)
	if err != nil {
		return -1, err
	}

	if ja.Values == nil {
		ja.Values = []JsonValue{*jv}
	} else {
		ja.Values = append(ja.Values, *jv)
	}

	return len(ja.Values) - 1, nil
}

func (ja *JsonArray) Set(index int, value interface{}) error {
	if ja.Values == nil || len(ja.Values) <= index {
		err := fmt.Sprintf("Position [%d] does not exist", index)
		return errors.New(err)
	}

	jv, err := NewJsonValue(value)
	if err != nil {
		return err
	}
	ja.Values[index] = *jv

	return nil
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

func (ja *JsonArray) ToString(pretty bool) string {
	return ja.toString(pretty, 0)
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
