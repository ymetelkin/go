package json

import (
	"errors"
	"fmt"
	"strings"
)

//Object represents JSON object
type Object struct {
	names      []string
	pnames     map[string]ParameterizedString
	Properties map[string]value
}

//AddProperty adds property to parent object
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

//AddString adds string property to parent object
func (jo *Object) AddString(name string, value string) error {
	return jo.addValue(name, newString(value))
}

//AddInt adds int property to parent object
func (jo *Object) AddInt(name string, value int) error {
	return jo.addValue(name, newInt(value))
}

//AddFloat adds float property to parent object
func (jo *Object) AddFloat(name string, value float64) error {
	return jo.addValue(name, newFloat(value))
}

//AddBool adds bool property to parent object
func (jo *Object) AddBool(name string, value bool) error {
	return jo.addValue(name, newBool(value))
}

//AddObject adds JSON object property to parent object
func (jo *Object) AddObject(name string, value Object) error {
	return jo.addValue(name, newObject(value))
}

//AddArray adds JSON array property to parent object
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
		return fmt.Errorf("Field [%s] does not exist", name)
	}

	jo.Properties[name] = value

	return nil
}

//SetInt sets int value of named property
func (jo *Object) SetInt(name string, value int) error {
	return jo.setValue(name, newInt(value))
}

//SetFloat sets float value of named property
func (jo *Object) SetFloat(name string, value float64) error {
	return jo.setValue(name, newFloat(value))
}

//SetBool sets int value of named property
func (jo *Object) SetBool(name string, value bool) error {
	return jo.setValue(name, newBool(value))
}

//SetString sets string value of named property
func (jo *Object) SetString(name string, value string) error {
	return jo.setValue(name, newString(value))
}

//SetObject sets JSON object value of named property
func (jo *Object) SetObject(name string, value Object) error {
	return jo.setValue(name, newObject(value))
}

//SetArray sets JSON array value of named property
func (jo *Object) SetArray(name string, value Array) error {
	return jo.setValue(name, newArray(value))
}

//Remove removes named property
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
	}
	return value{}, fmt.Errorf("Field [%s] does not exist", name)
}

//GetString returns string value of named property
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

//GetInt returns string int of named property
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

//GetFloat returns float value of named property
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

//GetBool returns bool value of named property
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

//GetObject returns JSON object value of named property
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

//GetArray returns JSON array value of named property
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

//IsEmpty checks for properties presense
func (jo *Object) IsEmpty() bool {
	if jo.Properties == nil || len(jo.Properties) == 0 {
		return true
	}

	return false
}

//Names returns all field names
func (jo *Object) Names() []string {
	return jo.names
}

//ToString returns pretty serialization
func (jo *Object) String() string {
	return jo.toString(true, 0)
}

//InlineString returns condensed serialization
func (jo *Object) InlineString() string {
	return jo.toString(false, 0)
}

func (jo *Object) toString(pretty bool, level int) string {
	if jo.Properties == nil || len(jo.Properties) == 0 {
		return "{}"
	}

	var sb strings.Builder

	sb.WriteRune(runeLeftCurly)
	if pretty {
		sb.WriteRune(runeCR)
		sb.WriteRune(runeLF)
	}

	next := level + 1

	for index, name := range jo.names {
		jv, err := jo.getValue(name)
		if err == nil {
			if index > 0 {
				sb.WriteRune(runeComma)

				if pretty {
					sb.WriteRune(runeCR)
					sb.WriteRune(runeLF)
				}
			}

			if pretty {
				i := 0
				for i <= level {
					sb.WriteRune(runeSpace)
					sb.WriteRune(runeSpace)
					i++
				}
			}

			sb.WriteRune(runeQuote)
			sb.WriteString(name)
			sb.WriteRune(runeQuote)
			sb.WriteRune(runeColon)
			if pretty {
				sb.WriteRune(runeSpace)
			}
			s := jv.ToString(pretty, next)
			sb.WriteString(s)
		}
	}

	if pretty {
		sb.WriteRune(runeCR)
		sb.WriteRune(runeLF)
		i := 0
		for i < level {
			sb.WriteRune(runeSpace)
			sb.WriteRune(runeSpace)
			i++
		}
	}
	sb.WriteRune(runeRightCurly)

	return sb.String()
}
