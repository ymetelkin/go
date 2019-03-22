package json

import (
	"errors"
	"fmt"
	"strings"
)

type JsonObject struct {
	names      []string
	pnames     map[string]ParameterizedString
	pvalues    map[string]ParameterizedString
	Properties map[string]JsonValue
}

func (jo *JsonObject) Copy() *JsonObject {
	if jo == nil {
		return nil
	}

	copy := JsonObject{}

	if jo.Properties != nil {
		for name, value := range jo.Properties {
			copy.Add(name, value)
		}
	}

	return &copy
}

func (jo *JsonObject) Add(name string, value interface{}) error {
	name = strings.Trim(name, " ")
	if name == "" {
		return errors.New("Missing field name")
	}

	if jo.Properties == nil {
		jo.Properties = make(map[string]JsonValue)
	}

	jv, err := NewJsonValue(value)
	if err != nil {
		return err
	}
	jo.Properties[name] = *jv

	if jo.names == nil {
		jo.names = []string{name}
	} else {
		jo.names = append(jo.names, name)
	}

	return nil
}

func (jo *JsonObject) Set(name string, value interface{}) error {
	name = strings.Trim(name, " ")
	if name == "" {
		return errors.New("Missing field name")
	}

	exists := false

	if jo.Properties != nil {
		_, exists = jo.Properties[name]
	}

	if !exists {
		err := fmt.Sprintf("Field [%s] does not exist", name)
		return errors.New(err)
	}

	jv, err := NewJsonValue(value)
	if err != nil {
		return err
	}
	jo.Properties[name] = *jv

	return nil
}

func (jo *JsonObject) Remove(name string) error {
	name = strings.Trim(name, " ")
	if name == "" {
		return errors.New("Missing field name")
	}

	if jo.Properties == nil {
		return nil
	}

	delete(jo.Properties, name)

	tmp := []string{}
	for _, n := range jo.names {
		if n != name {
			tmp = append(tmp, n)
		}
	}
	jo.names = tmp

	return nil
}

func (jo *JsonObject) GetValue(name string) (*JsonValue, error) {
	name = strings.Trim(name, " ")
	if name == "" {
		return nil, errors.New("Missing field name")
	}

	if jo.Properties == nil {
		return nil, nil
	}

	value, ok := jo.Properties[name]
	if ok {
		return &value, nil
	} else {
		err := fmt.Sprintf("Field [%s] does not exist", name)
		return nil, errors.New(err)
	}
}

func (jo *JsonObject) GetString(name string) (string, error) {
	jv, err := jo.GetValue(name)
	if err != nil {
		return "", err
	}

	s, err := jv.GetString()
	if err != nil {
		return "", err
	}

	return s, nil
}

func (jo *JsonObject) GetInt(name string) (int, error) {
	jv, err := jo.GetValue(name)
	if err != nil {
		return 0, err
	}

	i, err := jv.GetInt()
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (jo *JsonObject) GetFloat(name string) (float64, error) {
	jv, err := jo.GetValue(name)
	if err != nil {
		return 0, err
	}

	f, err := jv.GetFloat()
	if err != nil {
		return 0, err
	}

	return f, nil
}

func (jo *JsonObject) GetBoolean(name string) (bool, error) {
	jv, err := jo.GetValue(name)
	if err != nil {
		return false, err
	}

	b, err := jv.GetBoolean()
	if err != nil {
		return false, err
	}

	return b, nil
}

func (jo *JsonObject) GetObject(name string) (*JsonObject, error) {
	jv, err := jo.GetValue(name)
	if err != nil {
		return nil, err
	}

	obj, err := jv.GetObject()
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (jo *JsonObject) GetArray(name string) (*JsonArray, error) {
	jv, err := jo.GetValue(name)
	if err != nil {
		return nil, err
	}

	ja, err := jv.GetArray()
	if err != nil {
		return nil, err
	}

	return ja, nil
}

func (jo *JsonObject) IsEmpty() bool {
	if jo.Properties == nil || len(jo.Properties) == 0 {
		return true
	}

	return false
}

func (jo *JsonObject) ToString(pretty bool) string {
	return jo.toString(pretty, 0)
}

func (jo *JsonObject) toString(pretty bool, level int) string {
	if jo.Properties == nil || len(jo.Properties) == 0 {
		return "{}"
	}

	var sb strings.Builder

	sb.WriteRune(TOKEN_LEFT_CURLY)
	if pretty {
		sb.WriteRune(TOKEN_CR)
		sb.WriteRune(TOKEN_LF)
	}

	next := level + 1

	for index, name := range jo.names {
		jv, err := jo.GetValue(name)
		if err == nil {
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

			sb.WriteRune(TOKEN_QUOTE)
			sb.WriteString(name)
			sb.WriteRune(TOKEN_QUOTE)
			sb.WriteRune(TOKEN_COLON)
			if pretty {
				sb.WriteRune(TOKEN_SPACE)
			}
			s := jv.toString(pretty, next)
			sb.WriteString(s)
		}
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
	sb.WriteRune(TOKEN_RIGHT_CURLY)

	return sb.String()
}
