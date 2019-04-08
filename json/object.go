package json

import (
	"errors"
	"fmt"
	"strings"
)

type Object struct {
	names      []string
	pnames     map[string]ParameterizedString
	Properties map[string]value
}

func (jo *Object) AddProperty(jp Property) error {
	if jp.IsEmtpy() {
		return errors.New("Missing property")
	}
	return jo.addValue(jp.Field, jp.Value)
}

func (jo *Object) addValue(name string, jv value) error {
	name = strings.Trim(name, " ")
	if name == "" {
		return errors.New("Missing field name")
	}

	if jo.Properties == nil {
		jo.Properties = make(map[string]value)
	}

	_, ok := jo.Properties[name]
	if ok {
		return errors.New("Field already exists: " + name)
	}

	jo.Properties[name] = jv

	if jo.names == nil {
		jo.names = []string{name}
	} else {
		jo.names = append(jo.names, name)
	}

	return nil
}

func (jo *Object) AddString(name string, value string) error {
	return jo.addValue(name, newString(value))
}

func (jo *Object) AddInt(name string, value int) error {
	return jo.addValue(name, newInt(value))
}

func (jo *Object) AddFloat(name string, value float64) error {
	return jo.addValue(name, newFloat(value))
}

func (jo *Object) AddBool(name string, value bool) error {
	return jo.addValue(name, newBool(value))
}

func (jo *Object) AddObject(name string, value Object) error {
	return jo.addValue(name, newObject(value))
}

func (jo *Object) AddArray(name string, value Array) error {
	return jo.addValue(name, newArray(value))
}

func (jo *Object) setValue(name string, value value) error {
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

func (jo *Object) SetInt(name string, value int) error {
	return jo.setValue(name, newInt(value))
}

func (jo *Object) SetFloat(name string, value float64) error {
	return jo.setValue(name, newFloat(value))
}

func (jo *Object) SetBool(name string, value bool) error {
	return jo.setValue(name, newBool(value))
}

func (jo *Object) SetString(name string, value string) error {
	return jo.setValue(name, newString(value))
}

func (jo *Object) SetObject(name string, value Object) error {
	return jo.setValue(name, newObject(value))
}

func (jo *Object) SetArray(name string, value Array) error {
	return jo.setValue(name, newArray(value))
}

func (jo *Object) Remove(name string) error {
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

func (jo *Object) getValue(name string) (value, error) {
	name = strings.Trim(name, " ")
	if name == "" {
		return value{}, errors.New("Missing field name")
	}

	if jo.Properties == nil {
		return value{}, nil
	}

	jv, ok := jo.Properties[name]
	if ok {
		return jv, nil
	} else {
		err := fmt.Sprintf("Field [%s] does not exist", name)
		return value{}, errors.New(err)
	}
}

func (jo *Object) GetString(name string) (string, error) {
	jv, err := jo.getValue(name)
	if err != nil {
		return "", err
	}

	s, err := jv.GetString()
	if err != nil {
		return "", err
	}

	return s, nil
}

func (jo *Object) GetInt(name string) (int, error) {
	jv, err := jo.getValue(name)
	if err != nil {
		return 0, err
	}

	i, err := jv.GetInt()
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (jo *Object) GetFloat(name string) (float64, error) {
	jv, err := jo.getValue(name)
	if err != nil {
		return 0, err
	}

	f, err := jv.GetFloat()
	if err != nil {
		return 0, err
	}

	return f, nil
}

func (jo *Object) GetBool(name string) (bool, error) {
	jv, err := jo.getValue(name)
	if err != nil {
		return false, err
	}

	b, err := jv.GetBool()
	if err != nil {
		return false, err
	}

	return b, nil
}

func (jo *Object) GetObject(name string) (Object, error) {
	jv, err := jo.getValue(name)
	if err == nil {
		o, err := jv.GetObject()
		if err == nil {
			return o, nil
		}
	}

	return Object{}, err
}

func (jo *Object) GetArray(name string) (Array, error) {
	jv, err := jo.getValue(name)
	if err == nil {
		ja, err := jv.GetArray()
		if err == nil {
			return ja, nil
		}
	}

	return Array{}, err
}

func (jo *Object) IsEmpty() bool {
	if jo.Properties == nil || len(jo.Properties) == 0 {
		return true
	}

	return false
}

func (jo *Object) ToString() string {
	return jo.toString(true, 0)
}

func (jo *Object) ToInlineString() string {
	return jo.toString(false, 0)
}

func (jo *Object) toString(pretty bool, level int) string {
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
		jv, err := jo.getValue(name)
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
			s := jv.ToString(pretty, next)
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
