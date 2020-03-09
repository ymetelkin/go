package json

import (
	"errors"
	"strings"
)

//SetParams sets parameters
func (jo *Object) SetParams(params map[string]Value, emptycheck map[string][]string) (modified bool) {
	if len(jo.Properties) == 0 || len(jo.params) == 0 {
		return
	}

	if len(params) == 0 {
		ps, ok := jo.GetObject("params")
		if ok {
			params = ps.Map()
			jo.Remove("params")
			modified = true
		}
	}

	if params == nil {
		params = make(map[string]Value)
	}

	var (
		remove []string
		names  = make(map[string]string)
	)

	for _, jp := range jo.Properties {
		test, ok := jo.params[jp.Name]
		if !ok {
			continue
		}

		if test == 1 || test == 3 {
			_, txt, ok := setTextParams(jp.Name, params)
			if ok {
				modified = true
				if txt == "" {
					remove = append(remove, jp.Name)
					continue
				} else {
					names[jp.Name] = txt
				}
			}
		}

		if test > 1 {
			switch jp.Value.t() {
			case tString:
				s, ok := jp.Value.String()
				if !ok {
					continue
				}
				v, txt, ok := setTextParams(s, params)
				if ok {
					modified = true
					if txt == "" {
						if v == nil || v.t() == 0 {
							remove = append(remove, jp.Name)
						} else {
							jo.Set(jp.Name, v)
						}
					} else {
						jo.Set(jp.Name, String(txt))
					}
				}
			case tObject:
				child, ok := jp.Value.Object()
				if !ok {
					continue
				}
				ok = child.SetParams(params, emptycheck)
				if ok {
					modified = true
					if isEmpty(jp.Name, &child, emptycheck) {
						remove = append(remove, jp.Name)
					} else {
						jo.Set(jp.Name, O(child))
					}
				}
			case tArray:
				ja, ok := jp.Value.Array()
				if !ok {
					continue
				}
				modified = ja.SetParams(params, emptycheck)
				if modified {
					if len(ja.Values) == 0 {
						remove = append(remove, jp.Name)
					} else {
						jo.Set(jp.Name, A(ja))
					}
				}
			}
		}
	}

	for _, name := range remove {
		jo.Remove(name)
	}

	for name, update := range names {
		for f, i := range jo.fields {
			if f == name {
				jp := jo.Properties[i]
				jo.fields[update] = i
				jo.Properties[i] = Property{
					Name:  update,
					Value: jp.Value,
				}
				delete(jo.fields, name)
				break
			}
		}
	}

	//jo.fields = nil

	return modified
}

//SetParams set array parameters
func (ja *Array) SetParams(params map[string]Value, emptycheck map[string][]string) (modified bool) {
	if len(ja.params) == 0 {
		return
	}

	if params == nil {
		params = make(map[string]Value)
	}

	var values []Value

	for i, jv := range ja.Values {
		var add bool

		for _, idx := range ja.params {
			if i == idx {
				add = true

				switch jv.t() {
				case tString:
					s, ok := jv.String()
					if !ok {
						break
					}

					v, txt, ok := setTextParams(s, params)
					if ok {
						modified = true
						if txt == "" {
							if v != nil && v.t() > 0 {
								values = append(values, v)
							}
						} else {
							values = append(values, String(txt))
						}
					}
				case tObject:
					jo, ok := jv.Object()
					if !ok {
						continue
					}
					ok = jo.SetParams(params, emptycheck)
					if ok {
						modified = true
						if len(jo.Properties) > 0 {
							values = append(values, O(jo))
						}
					}
				case tArray:
					a, ok := jv.Array()
					if !ok {
						continue
					}
					modified = a.SetParams(params, emptycheck)
					if modified {
						if len(a.Values) > 0 {
							values = append(values, A(a))
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

func setTextParams(s string, params map[string]Value) (jv Value, text string, modified bool) {
	p := newParser([]byte(s))

	var (
		sb    strings.Builder
		multi bool
	)

	for {
		c, ok := p.Read()
		if !ok {
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
				s, _ := v.String()
				sb.WriteString(s)
			} else if def != "" {
				jv = String(def)
				sb.WriteString(def)
			}
		} else {
			sb.WriteByte(c)
			multi = true
		}
	}

	if modified && (multi || (jv != nil && jv.t() == tString)) {
		text = sb.String()
	}

	return
}

func isEmpty(name string, jo *Object, emptycheck map[string][]string) bool {
	if len(jo.Properties) == 0 {
		return true
	}

	if len(emptycheck) == 0 {
		return false
	}

	for k, vs := range emptycheck {
		if k == name {
			for _, v := range vs {
				_, ok := jo.fields[v]
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
	c, ok := p.Read()
	if !ok || c != '{' {
		return
	}

	var (
		sb strings.Builder
		df bool
	)

	for {
		c, ok := p.Read()
		if !ok {
			err = errors.New("Expected parameter, found EOF")
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

//GetParams gets object parameters
func (jo *Object) GetParams() (params map[string]Value) {
	if len(jo.Properties) == 0 {
		return
	}

	params = make(map[string]Value)

	for _, jp := range jo.Properties {
		getParams(jp.Name, params)

		switch jp.Value.t() {
		case tString:
			s, ok := jp.Value.String()
			if ok {
				getParams(s, params)
			}
		case tObject:
			o, ok := jp.Value.Object()
			if ok {
				temp := o.GetParams()
				if len(temp) > 0 {
					for k, v := range temp {
						params[k] = v
					}
				}
			}
		case tArray:
			a, ok := jp.Value.Array()
			if ok {
				temp := a.GetParams()
				if len(temp) > 0 {
					for k, v := range temp {
						params[k] = v
					}
				}
			}
		}
	}
	return
}

//GetParams gets array parameters
func (ja *Array) GetParams() (params map[string]Value) {
	if len(ja.Values) == 0 {
		return
	}

	params = make(map[string]Value)

	for _, jv := range ja.Values {
		switch jv.t() {
		case tString:
			s, ok := jv.String()
			if ok {
				getParams(s, params)
			}
		case tObject:
			o, ok := jv.Object()
			if ok {
				temp := o.GetParams()
				if len(temp) > 0 {
					for k, v := range temp {
						params[k] = v
					}
				}
			}
		case tArray:
			a, ok := jv.Array()
			if ok {
				temp := a.GetParams()
				if len(temp) > 0 {
					for k, v := range temp {
						params[k] = v
					}
				}
			}
		}
	}

	return params
}

func getParams(s string, params map[string]Value) {
	p := newParser([]byte(s))

	for {
		c, ok := p.Read()
		if !ok {
			break
		}

		if c == '$' {
			name, def, e := p.ParseParam()
			if e == nil && name != "" {
				params[name] = String(def)
			}
		}
	}
}
