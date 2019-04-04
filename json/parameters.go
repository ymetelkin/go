package json

import (
	"errors"
	"fmt"
	"strings"
)

type Parameter struct {
	Name       string
	Default    string
	StartIndex int
	EndIndex   int
}

type ParameterizedString struct {
	Value           string
	IsParameterized bool
	Parameters      []Parameter
}

func (ps *ParameterizedString) IsOneParameter() bool {
	if !ps.IsParameterized || len(ps.Parameters) != 1 {
		return false
	}

	p := ps.Parameters[0]
	return p.StartIndex == 0 && p.EndIndex == len(ps.Value)
}

func (jo *Object) AddWithParameters(pname ParameterizedString, jv value) error {
	name := pname.Value
	name = strings.Trim(name, " ")
	if name == "" {
		return errors.New("Missing field name")
	}

	if jo.Properties == nil {
		jo.Properties = make(map[string]value)
	}

	jo.Properties[name] = jv
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

func ParseJsonObjectWithParameters(s string) (*Object, error) {
	jo, err := parseJsonObject(s, true)
	if err != nil {
		return nil, err
	}

	if jo != nil && !jo.IsEmpty() {
		jo.setObjectParameters(nil)
	}

	return jo, nil
}

func (jo *Object) setObjectParameters(props map[string]value) bool {
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
				if update.IsEmpty() {
					jo.Remove(name)
				} else {
					jo.setValue(name, update)
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
						jo.SetObject(name, child)
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
						jo.SetArray(name, ja)
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
				if update.IsEmpty() {
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

func (ja *Array) setArrayParameters(props map[string]value) bool {
	modified := false
	values := []value{}

	for _, jv := range ja.Values {
		if jv.Type == PARAMETERIZED {
			modified = true

			ps, err := jv.GetParameterizedString()
			if err == nil {
				add := setValueParameters(ps.Value, ps, props)
				if !add.IsEmpty() {
					values = append(values, add)
				}
			}
		} else if jv.Type == OBJECT {
			jo, err := jv.GetObject()
			if err == nil {
				modified = jo.setObjectParameters(props)
				if modified {
					if !jo.IsEmpty() {
						values = append(values, newObject(jo))
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
						values = append(values, newArray(ja))
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

func setValueParameters(s string, ps ParameterizedString, params map[string]value) value {
	if ps.IsOneParameter() {
		p := ps.Parameters[0]
		jv, ok := params[p.Name]
		if ok {
			return jv
		} else if p.Default != "" {
			return newString(p.Default)
		}
	} else {
		runes := []rune(s)
		replace := make(map[string]string)

		for _, p := range ps.Parameters {
			key := string(runes[p.StartIndex:p.EndIndex])
			jv, ok := params[p.Name]
			if ok {
				if jv.Type == STRING || jv.Type == INT || jv.Type == FLOAT || jv.Type == BOOL {
					replace[key] = fmt.Sprintf("%v", jv.Value)
				}
			} else {
				replace[key] = p.Default
			}
		}

		for key, value := range replace {
			s = strings.ReplaceAll(s, key, value)
		}

		if len(s) > 0 {
			return newString(s)
		}
	}

	return value{}
}

func parsePropertyNameWithParameters(runes []rune, size int, index int) (ParameterizedString, int, error) {
	index++
	start := index
	pstart := 0
	pdef := 0
	r := TOKEN_NULL

	var params []Parameter

	for index < size {
		r = runes[index]

		if r == TOKEN_QUOTE {
			end := index

			index++
			r, index = skipWhitespace(runes, size, index)
			if r == TOKEN_COLON {
				value := string(runes[start:end])
				parameterized := params != nil
				ps := ParameterizedString{Value: value, IsParameterized: parameterized, Parameters: params}
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

	var params []Parameter
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
			ps := ParameterizedString{Value: value, IsParameterized: parameterized, Parameters: params}
			return ps, index, nil
		} else {
			pstart, pdef, params = parseParameters(runes, index, r, start, pstart, pdef, params)
			sb.WriteRune(r)
		}

		index++
	}
	return ParameterizedString{}, index, nil
}

func parseParameters(runes []rune, index int, r rune, start int, pstart int, pdef int, params []Parameter) (int, int, []Parameter) {
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
			param := Parameter{name, def, pstart - start - 2, index - start + 1}

			if params == nil {
				params = []Parameter{param}
			} else {
				params = append(params, param)
			}

			pstart = 0
			pdef = 0
		}
	}

	return pstart, pdef, params
}
