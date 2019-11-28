package json

import (
	"fmt"
	"strings"
)

//SetParams sets parameters
func (jo *Object) SetParams(props map[string]value, emptycheck map[string][]string) (modified bool) {
	if jo.IsEmpty() {
		return
	}

	if props == nil || len(props) == 0 {
		params, ok := jo.GetObject("params")
		if ok && !params.IsEmpty() {
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
				s, ok := jv.GetString()
				if !ok {
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
				child, ok := jv.GetObject()
				if !ok {
					continue
				}
				ok = child.SetParams(props, emptycheck)
				if ok {
					modified = true
					if isEmpty(name, &child, emptycheck) {
						remove = append(remove, name)
					} else {
						jo.SetObject(name, child)
					}
				}
			} else if jv.Type == jsonArray {
				ja, ok := jv.GetArray()
				if !ok {
					continue
				}
				modified = ja.SetParams(props, emptycheck)
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
func (ja *Array) SetParams(props map[string]value, emptycheck map[string][]string) (modified bool) {
	if ja.pvalues == nil || len(ja.pvalues) == 0 {
		return
	}

	var values []value

	for i, jv := range ja.Values {
		var add bool

		for _, idx := range ja.pvalues {
			if i == idx {
				add = true

				if jv.Type == jsonString {
					s, ok := jv.GetString()
					if !ok {
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
				} else if jv.Type == jsonObject {
					jo, ok := jv.GetObject()
					if !ok {
						continue
					}
					ok = jo.SetParams(props, emptycheck)
					if ok {
						modified = true
						if !jo.IsEmpty() {
							values = append(values, newObject(jo))
						}
					}
				} else if jv.Type == jsonArray {
					a, ok := jv.GetArray()
					if !ok {
						continue
					}
					modified = a.SetParams(props, emptycheck)
					if modified {
						if !a.IsEmpty() {
							values = append(values, newArray(a))
						}
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

func isEmpty(name string, jo *Object, emptycheck map[string][]string) bool {
	if jo.IsEmpty() {
		return true
	}

	if emptycheck == nil {
		return false
	}

	for k, vs := range emptycheck {
		if k == name {
			for _, v := range vs {
				_, ok := jo.Properties[v]
				if !ok {
					return true
				}
			}
			break
		}
	}

	return false
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

//GetParams gets parameters
func (jo *Object) GetParams() map[string]string {
	params := make(map[string]string)

	if jo.IsEmpty() {
		return params
	}

	for name, jv := range jo.Properties {
		getParams(name, params)

		if jv.Type == jsonString {
			s, ok := jv.GetString()
			if ok {
				getParams(s, params)
			}
		} else if jv.Type == jsonObject {
			o, ok := jv.GetObject()
			if ok {
				temp := o.GetParams()
				for k, v := range temp {
					params[k] = v
				}
			}
		} else if jv.Type == jsonArray {
			a, ok := jv.GetArray()
			if ok {
				temp := a.GetParams()
				for k, v := range temp {
					params[k] = v
				}
			}
		}
	}

	return params
}

//GetParams gets parameters
func (ja *Array) GetParams() map[string]string {
	params := make(map[string]string)

	if ja.IsEmpty() {
		return params
	}

	for _, jv := range ja.Values {
		if jv.Type == jsonString {
			s, ok := jv.GetString()
			if ok {
				getParams(s, params)
			}
		} else if jv.Type == jsonObject {
			o, ok := jv.GetObject()
			if ok {
				temp := o.GetParams()
				for k, v := range temp {
					params[k] = v
				}
			}
		} else if jv.Type == jsonArray {
			a, ok := jv.GetArray()
			if ok {
				temp := a.GetParams()
				for k, v := range temp {
					params[k] = v
				}
			}
		}
	}

	return params
}

func getParams(s string, params map[string]string) {
	p := &parser{
		r: newReader(s),
	}

	for {
		c, e := p.r.ReadByte()
		if e != nil {
			break
		}

		if c == '$' {
			name, def, e := p.ParseParam()
			if e == nil && name != "" {
				params[name] = def
			}
		}
	}
}
