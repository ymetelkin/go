package json

import (
	"errors"
	"fmt"
	"strings"
)

type JsonProperty struct {
	Field string
	Value JsonValue
}

type JsonObject struct {
	names      []string
	pnames     map[string]ParameterizedString
	Properties map[string]JsonValue
}

func (jo *JsonObject) Copy() *JsonObject {
	if jo == nil {
		return nil
	}

	copy := JsonObject{}

	if jo.Properties != nil {
		for name, value := range jo.Properties {
			copy.AddValue(name, value)
		}
	}

	return &copy
}

func (jo *JsonObject) AddProperty(jp *JsonProperty) error {
	if jp == nil {
		return errors.New("Missing property")
	}
	return jo.AddValue(jp.Field, jp.Value)
}

func (jo *JsonObject) AddValue(name string, value JsonValue) error {
	name = strings.Trim(name, " ")
	if name == "" {
		return errors.New("Missing field name")
	}

	if jo.Properties == nil {
		jo.Properties = make(map[string]JsonValue)
	}

	jo.Properties[name] = value

	if jo.names == nil {
		jo.names = []string{name}
	} else {
		jo.names = append(jo.names, name)
	}

	return nil
}

func (jo *JsonObject) AddString(name string, value string) error {
	return jo.AddValue(name, &JsonStringValue{Value: value})
}

func (jo *JsonObject) AddInt(name string, value int) error {
	return jo.AddValue(name, &JsonIntValue{Value: value})
}

func (jo *JsonObject) AddFloat(name string, value float64) error {
	return jo.AddValue(name, &JsonFloatValue{Value: value})
}

func (jo *JsonObject) AddBoolean(name string, value bool) error {
	return jo.AddValue(name, &JsonBooleanValue{Value: value})
}

func (jo *JsonObject) AddObject(name string, value *JsonObject) error {
	return jo.AddValue(name, &JsonObjectValue{Value: *value})
}

func (jo *JsonObject) AddArray(name string, value *JsonArray) error {
	return jo.AddValue(name, &JsonArrayValue{Value: *value})
}

func (jo *JsonObject) SetValue(name string, value JsonValue) error {
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

	jo.Properties[name] = value

	return nil
}

func (jo *JsonObject) SetInt(name string, value int) error {
	return jo.SetValue(name, &JsonIntValue{Value: value})
}

func (jo *JsonObject) SetFloat(name string, value float64) error {
	return jo.SetValue(name, &JsonFloatValue{Value: value})
}

func (jo *JsonObject) SetBoolean(name string, value bool) error {
	return jo.SetValue(name, &JsonBooleanValue{Value: value})
}

func (jo *JsonObject) SetString(name string, value string) error {
	return jo.SetValue(name, &JsonStringValue{Value: value})
}

func (jo *JsonObject) SetObject(name string, value JsonObject) error {
	return jo.SetValue(name, &JsonObjectValue{Value: value})
}

func (jo *JsonObject) SetArray(name string, value JsonArray) error {
	return jo.SetValue(name, &JsonArrayValue{Value: value})
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

func (jo *JsonObject) GetValue(name string) (JsonValue, error) {
	name = strings.Trim(name, " ")
	if name == "" {
		return nil, errors.New("Missing field name")
	}

	if jo.Properties == nil {
		return nil, nil
	}

	value, ok := jo.Properties[name]
	if ok {
		return value, nil
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

	s, err := getString(jv)
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

	i, err := getInt(jv)
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

	f, err := getFloat(jv)
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

	b, err := getBoolean(jv)
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

	obj, err := getObject(jv)
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

	ja, err := getArray(jv)
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

func (jo *JsonObject) ToString() string {
	return jo.toString(true, 0)
}

func (jo *JsonObject) ToInlineString() string {
	return jo.toString(false, 0)
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
			s := toString(jv, pretty, next)
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
