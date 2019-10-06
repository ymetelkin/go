package json

import (
	"fmt"
	"strings"
)

//SetParams sets parameters
func (jo *Object) SetParams(props map[string]value) (modified bool) {
	if jo.IsEmpty() {
		return
	}

	if props == nil || len(props) == 0 {
		params, err := jo.GetObject("params")
		if err == nil && !params.IsEmpty() {
			props = params.Properties
			jo.Remove("params")
			modified = true
		}
	}

	if jo.pnames == nil || len(jo.pnames) == 0 {
		return false
	}

	if props == nil {
		props = make(map[string]value)
	}

	var (
		remove []string
		names  = make(map[string]string)
	)

	for name, jv := range jo.Properties {
		test, ok := jo.pnames[name]
		if !ok {
			continue
		}

		if test == 1 || test == 3 {
			_, txt, ok := setTextParams(name, props)
			if ok {
				modified = true
				if txt == "" {
					remove = append(remove, name)
					continue
				} else {
					names[name] = txt
				}
			}
		}

		if test > 1 {
			if jv.Type == jsonString {
				s, err := jv.GetString()
				if err != nil {
					continue
				}
				v, txt, ok := setTextParams(s, props)
				if ok {
					modified = true
					if txt == "" {
						if v.Type == 0 || v.Type == jsonNull {
							remove = append(remove, name)
						} else {
							jo.setValue(name, v)
						}
					} else {
						jo.setValue(name, newString(txt))
					}
				}
			} else if jv.Type == jsonObject {
				child, err := jv.GetObject()
				if err != nil {
					continue
				}
				ok = child.SetParams(props)
				if ok {
					modified = true
					if child.IsEmpty() {
						remove = append(remove, name)
					} else {
						jo.SetObject(name, child)
					}
				}
			} else if jv.Type == jsonArray {
				ja, err := jv.GetArray()
				if err != nil {
					continue
				}
				modified = ja.SetParams(props)
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

	for name, update := range names {
		for i, n := range jo.names {
			if n == name {
				jv, ok := jo.Properties[name]
				if ok {
					jo.names[i] = update
					jo.Properties[update] = jv
					delete(jo.Properties, name)
				}

				break
			}
		}
	}

	jo.pnames = nil

	return modified
}

//SetParams set array parameters
func (ja *Array) SetParams(props map[string]value) (modified bool) {
	if ja.pvalues == nil || len(ja.pvalues) == 0 {
		return
	}

	var values []value

	for i, jv := range ja.Values {
		var add bool

		for _, idx := range ja.pvalues {
			if i == idx {
				add = true

				s, err := jv.GetString()
				if err != nil {
					break
				}

				v, txt, ok := setTextParams(s, props)
				if ok {
					modified = true
					if txt == "" {
						if v.Type > 0 && v.Type != jsonNull {
							values = append(values, v)
						}
					} else {
						values = append(values, newString(txt))
					}
				}
			}
		}

		if !add {
			values = append(values, jv)
		}
	}

	ja.Values = values

	return
}

func setTextParams(s string, params map[string]value) (jv value, text string, modified bool) {
	p := &parser{
		r: newReader(s),
	}

	var (
		sb    strings.Builder
		multi bool
	)

	for {
		c, e := p.r.ReadByte()
		if e != nil {
			break
		}

		if c == '$' {
			name, def, e := p.ParseParam()
			if e != nil || name == "" {
				return
			}

			modified = true

			v, ok := params[name]
			if ok {
				jv = v
				sb.WriteString(fmt.Sprintf("%v", v.Value))
			} else if def != "" {
				jv = newString(def)
				sb.WriteString(def)
			}
		} else {
			sb.WriteByte(c)
			multi = true
		}
	}

	if modified && (multi || jv.Type == jsonString) {
		text = sb.String()
	}

	return
}

func (p *parser) ParseParam() (name string, def string, err error) {
	c, err := p.r.ReadByte()
	if err != nil || c != '{' {
		return
	}

	var (
		sb strings.Builder
		df bool
	)

	for {
		c, e := p.r.ReadByte()
		if e != nil {
			err = e
			break
		}

		if c == '?' {
			name = sb.String()
			sb.Reset()
			df = true
			continue
		}

		if c == '}' {
			if df {
				def = sb.String()
			} else {
				name = sb.String()
			}
			break
		}

		if df || isProperty(c) {
			sb.WriteByte(c)
		}
	}

	return
}
