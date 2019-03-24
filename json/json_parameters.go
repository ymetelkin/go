package json

import (
	"errors"
	"fmt"
	"strings"
)

type JsonParameter struct {
	Name       string
	Default    string
	StartIndex int
	EndIndex   int
}

type ParameterizedString struct {
	Value           string
	IsParameterized bool
	Parameters      []JsonParameter
}

func (ps ParameterizedString) IsOneParameter() bool {
	if !ps.IsParameterized || len(ps.Parameters) != 1 {
		return false
	}

	p := ps.Parameters[0]
	return p.StartIndex == 0 && p.EndIndex == len(ps.Value)
}

func (jo *JsonObject) AddWithParameters(pname ParameterizedString, value *JsonValue) error {
	name := pname.Value
	name = strings.Trim(name, " ")
	if name == "" {
		return errors.New("Missing field name")
	}

	if jo.Properties == nil {
		jo.Properties = make(map[string]JsonValue)
	}

	jv, err := NewJsonValue(*value)
	if err != nil {
		return err
	}
	jo.Properties[name] = *jv

	if jo.names == nil {
		jo.names = []string{name}
	} else {
		jo.names = append(jo.names, name)
	}

	if pname.IsParameterized {
		if jo.pnames == nil {
			jo.pnames = make(map[string]ParameterizedString)
		}

		jo.pnames[name] = pname
	}

	return nil
}

func ParseJsonObjectWithParameters(s string) (*JsonObject, error) {
	jo, err := parseJsonObject(s, true)
	if err != nil {
		return nil, err
	}

	if jo != nil && !jo.IsEmpty() {
		jo.setObjectParameters(nil)
	}

	return jo, nil
}

func (jo *JsonObject) setObjectParameters(props map[string]JsonValue) bool {
	if jo.IsEmpty() {
		return false
	}

	modified := false

	if props == nil || len(props) == 0 {
		params, err := jo.GetObject("params")
		if err == nil && params != nil && !params.IsEmpty() {
			props = params.Properties
			jo.Remove("params")
		}
	}

	remove := []string{}

	for name, jv := range jo.Properties {
		if jv.Type == PARAMETERIZED {
			modified = true

			ps, err := jv.GetParameterizedString()
			if err == nil {
				update := setValueParameters(ps.Value, ps, props)
				if update == nil {
					jo.Remove(name)
				} else {
					jo.Set(name, *update)
				}
			}
		} else if jv.Type == OBJECT {
			child, err := jv.GetObject()
			if err == nil {
				modified = child.setObjectParameters(props)
				if modified {
					if child.IsEmpty() {
						remove = append(remove, name)
					} else {
						jo.Set(name, JsonValue{Value: *child, Type: OBJECT})
					}
				}
			}
		} else if jv.Type == ARRAY {
			ja, err := jv.GetArray()
			if err == nil {
				modified = ja.setArrayParameters(props)
				if modified {
					if ja.IsEmpty() {
						remove = append(remove, name)
					} else {
						jo.Set(name, JsonValue{Value: *ja, Type: ARRAY})
					}
				}
			}
		}
	}

	for _, name := range remove {
		jo.Remove(name)
	}

	if jo.pnames != nil {
		for name, ps := range jo.pnames {
			jv, ok := jo.Properties[name]
			if ok {
				update := setValueParameters(name, ps, props)
				if update == nil {
					jo.Remove(name)
				} else {
					s, err := update.GetString()
					if err == nil {
						var idx int
						for i, n := range jo.names {
							if n == name {
								idx = i
								break
							}
						}
						jo.names[idx] = s

						jo.Properties[s] = jv
						delete(jo.Properties, name)
					} else {
						jo.Remove(name)
					}
				}
			}
		}
	}

	jo.pnames = nil

	return modified
}

func (ja *JsonArray) setArrayParameters(props map[string]JsonValue) bool {
	modified := false
	values := []JsonValue{}

	for _, jv := range ja.Values {
		if jv.Type == PARAMETERIZED {
			modified = true

			ps, err := jv.GetParameterizedString()
			if err == nil {
				add := setValueParameters(ps.Value, ps, props)
				if add != nil {
					values = append(values, *add)
				}
			}
		} else if jv.Type == OBJECT {
			jo, err := jv.GetObject()
			if err == nil {
				modified = jo.setObjectParameters(props)
				if modified {
					if !jo.IsEmpty() {
						values = append(values, JsonValue{Value: *jo, Type: OBJECT})
					}
				} else {
					values = append(values, jv)
				}
			}
		} else if jv.Type == ARRAY {
			ja, err := jv.GetArray()
			if err == nil {
				modified = ja.setArrayParameters(props)
				if modified {
					if !ja.IsEmpty() {
						values = append(values, JsonValue{Value: *ja, Type: ARRAY})
					}
				} else {
					values = append(values, jv)
				}
			}
		} else {
			values = append(values, jv)
		}
	}

	ja.Values = values

	return modified
}

func setValueParameters(s string, ps ParameterizedString, params map[string]JsonValue) *JsonValue {
	if ps.IsOneParameter() {
		p := ps.Parameters[0]
		jv, ok := params[p.Name]
		if ok {
			return &jv
		} else if p.Default != "" {
			return &JsonValue{Value: p.Default, Type: STRING}
		}
	} else {
		runes := []rune(s)
		replace := make(map[string]string)

		for _, p := range ps.Parameters {
			key := string(runes[p.StartIndex:p.EndIndex])
			jv, ok := params[p.Name]
			if ok && (jv.Type == STRING || jv.Type == NUMBER || jv.Type == BOOLEAN) {
				replace[key] = fmt.Sprintf("%v", jv.Value)
			} else {
				replace[key] = p.Default
			}
		}

		for key, value := range replace {
			s = strings.ReplaceAll(s, key, value)
		}

		if len(s) > 0 {
			return &JsonValue{Value: s, Type: STRING}
		}
	}

	return nil
}

func parsePropertyNameWithParameters(runes []rune, size int, index int) (ParameterizedString, int, error) {
	index++
	start := index
	pstart := 0
	pdef := 0
	r := TOKEN_NULL

	var params []JsonParameter

	for index < size {
		r = runes[index]

		if r == TOKEN_QUOTE {
			end := index

			index++
			r, index = skipWhitespace(runes, size, index)
			if r == TOKEN_COLON {
				value := string(runes[start:end])
				parameterized := params != nil
				ps := ParameterizedString{value, parameterized, params}
				return ps, index, nil
			} else {
				err := fmt.Sprintf("Expected ':', found '%c'", r)
				return ParameterizedString{}, index, errors.New(err)
			}
		} else {
			pstart, pdef, params = parseParameters(runes, index, r, start, pstart, pdef, params)
		}

		index++
	}

	err := fmt.Sprintf("Expected '\"', found '%c'", r)
	return ParameterizedString{}, index, errors.New(err)
}

func parseTextValueWithParameters(runes []rune, size int, index int) (ParameterizedString, int, error) {
	start := index
	pstart := 0
	pdef := 0

	var params []JsonParameter
	var sb strings.Builder

	for index < size {
		r := runes[index]

		if r == TOKEN_BACKSLASH {
			index++
			if index < size {
				test := runes[index]
				if test == TOKEN_R {
					sb.WriteRune(TOKEN_CR)
				} else if test == TOKEN_N {
					sb.WriteRune(TOKEN_LF)
				} else if test == TOKEN_T {
					sb.WriteRune(TOKEN_HT)
				} else if test == TOKEN_B {
					sb.WriteRune(TOKEN_BS)
				} else if test == TOKEN_F {
					sb.WriteRune(TOKEN_FF)
				} else if test == TOKEN_A {
					sb.WriteRune(TOKEN_BL)
				} else if test == TOKEN_V {
					sb.WriteRune(TOKEN_VT)
				}
			}
		} else if r == TOKEN_QUOTE {
			value := string(runes[start:index])
			parameterized := params != nil
			ps := ParameterizedString{value, parameterized, params}
			return ps, index, nil
		} else {
			pstart, pdef, params = parseParameters(runes, index, r, start, pstart, pdef, params)
			sb.WriteRune(r)
		}

		index++
	}
	return ParameterizedString{}, index, nil
}

func parseParameters(runes []rune, index int, r rune, start int, pstart int, pdef int, params []JsonParameter) (int, int, []JsonParameter) {
	if r == TOKEN_DOLLAR {
		pstart = index + 1
		pdef = 0
	} else if r == TOKEN_LEFT_CURLY {
		if pstart != index {
			pstart = 0
		} else {
			pstart = index + 1
		}
	} else if r == TOKEN_QUESTION {
		if pstart > 1 {
			pdef = index + 1
		}
	} else if r == TOKEN_RIGHT_CURLY {
		if pstart > 1 {
			end := index
			def := ""
			if pdef > 3 && pdef < index-1 {
				def = string(runes[pdef:index])
				end = pdef - 1
			}
			name := string(runes[pstart:end])
			param := JsonParameter{name, def, pstart - start - 2, index - start + 1}

			if params == nil {
				params = []JsonParameter{param}
			} else {
				params = append(params, param)
			}

			pstart = 0
			pdef = 0
		}
	}

	return pstart, pdef, params
}
