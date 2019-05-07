package json

import (
	"errors"
	"fmt"
	"strings"
)

//Parameter represents JSON parameters
type Parameter struct {
	Name       string
	Default    string
	StartIndex int
	EndIndex   int
}

//ParameterizedString represents string with JSON parameter placeholders
type ParameterizedString struct {
	Value           string
	IsParameterized bool
	Parameters      []Parameter
}

//IsOneParameter checks if the ParameterizedString is entirely one parameter
func (ps *ParameterizedString) IsOneParameter() bool {
	if !ps.IsParameterized || len(ps.Parameters) != 1 {
		return false
	}

	p := ps.Parameters[0]
	return p.StartIndex == 0 && p.EndIndex == len(ps.Value)
}

//AddWithParameters adds ParameterizedString property to JSON object
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

//ParseJSONObjectWithParameters merges JSON with parameters
func ParseJSONObjectWithParameters(s string) (Object, error) {
	jo, err := parseJSONObject(s, true)
	if err != nil {
		return Object{}, err
	}

	if !jo.IsEmpty() {
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
		if err == nil && !params.IsEmpty() {
			props = params.Properties
			jo.Remove("params")
		}
	}

	remove := []string{}

	for name, jv := range jo.Properties {
		if jv.Type == jsonParams {
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
		} else if jv.Type == jsonObject {
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
		} else if jv.Type == jsonArray {
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
		if jv.Type == jsonParams {
			modified = true

			ps, err := jv.GetParameterizedString()
			if err == nil {
				add := setValueParameters(ps.Value, ps, props)
				if !add.IsEmpty() {
					values = append(values, add)
				}
			}
		} else if jv.Type == jsonObject {
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
		} else if jv.Type == jsonArray {
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
				if jv.Type == jsonString || jv.Type == jsonInt || jv.Type == jsonFloat || jv.Type == jsonBool {
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
	r := runeNull

	var params []Parameter

	for index < size {
		r = runes[index]

		if r == runeQuote {
			end := index

			index++
			r, index = skipWhitespace(runes, size, index)
			if r == runeColon {
				value := string(runes[start:end])
				parameterized := params != nil
				ps := ParameterizedString{Value: value, IsParameterized: parameterized, Parameters: params}
				return ps, index, nil
			}
			return ParameterizedString{}, index, fmt.Errorf("Expected ':', found '%c'", r)
		}
		pstart, pdef, params = parseParameters(runes, index, r, start, pstart, pdef, params)
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

		if r == runeBackslash {
			index++
			if index < size {
				test := runes[index]
				if test == runeR {
					sb.WriteRune(runeCR)
				} else if test == runeN {
					sb.WriteRune(runeLF)
				} else if test == runeT {
					sb.WriteRune(runeHT)
				} else if test == runeB {
					sb.WriteRune(runeBS)
				} else if test == runeF {
					sb.WriteRune(runeFF)
				} else if test == runeA {
					sb.WriteRune(runeBL)
				} else if test == runeV {
					sb.WriteRune(runeVT)
				} else if test == runeQuote {
					sb.WriteRune(runeQuote)
				}
			}
		} else if r == runeQuote {
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
	if r == runeDollar {
		pstart = index + 1
		pdef = 0
	} else if r == runeLeftCurly {
		if pstart != index {
			pstart = 0
		} else {
			pstart = index + 1
		}
	} else if r == runeQuestion {
		if pstart > 1 {
			pdef = index + 1
		}
	} else if r == runeRightCurly {
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
