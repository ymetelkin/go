package json

import (
	"errors"
	"fmt"
	"strings"
)

//SetParams sets parameters
func (jo *Object) SetParams(params map[string]Value, emptycheck map[string][]string) (modified bool) {
	if jo.IsEmpty() || len(jo.params) == 0{
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
		test, ok := jo.params[jp.Field]
		if !ok {
			continue
		}

		if test == 1 || test == 3 {
			_, txt, ok := setTextParams(jp.Field, params)
			if ok {
				modified = true
				if txt == "" {
					remove = append(remove, jp.Field)
					continue
				} else {
					names[jp.Field] = txt
				}
			}
		}

		if test > 1 {
			switch jp.Value.Type {
			case TypeString:
				s, ok := jp.Value.String()
				if !ok {
					continue
				}
				v, txt, ok := setTextParams(s, params)
				if ok {
					modified = true
					if txt == "" {
						if v.Type == TypeNull {
							remove = append(remove, jp.Field)
						} else {
							jo.Set(jp.Field, v)
						}
					} else {
						jo.Set(jp.Field, NewString(txt))
					}
				}
			case TypeObject:
				child, ok := jp.Value.Object()
				if !ok {
					continue
				}
				ok = child.SetParams(params, emptycheck)
				if ok {
					modified = true
					if isEmpty(jp.Field, &child, emptycheck) {
						remove = append(remove, jp.Field)
					} else {
						jo.SetObject(jp.Field, child)
					}
				}
			case TypeArray:
				ja, ok := jp.Value.Array()
				if !ok {
					continue
				}
				modified = ja.SetParams(params, emptycheck)
				if modified {
					if len(ja.Values) == 0 {
						remove = append(remove, jp.Field)
					} else {
						jo.SetArray(jp.Field, ja)
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
					Field: update,
					Value: jp.Value,
				}
				delete(jo.fields, name)
				break
			}
		}
	}

	jo.fields = nil

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

				switch jv.Type {
				case TypeString:
					s, ok := jv.String()
					if !ok {
						break
					}

					v, txt, ok := setTextParams(s, params)
					if ok {
						modified = true
						if txt == "" {
							if v.Type > 0 && v.Type != TypeNull {
								values = append(values, v)
							}
						} else {
							values = append(values, Value{
								Type: TypeString,
								data: txt,
							})
						}
					}
				case TypeObject:
					jo, ok := jv.Object()
					if !ok {
						continue
					}
					ok = jo.SetParams(params, emptycheck)
					if ok {
						modified = true
						if !jo.IsEmpty() {
							values = append(values, Value{
								Type: TypeObject,
								data: jo,
							})
						}
					}
				case TypeArray:
					a, ok := jv.Array()
					if !ok {
						continue
					}
					modified = a.SetParams(params, emptycheck)
					if modified {
						if len(a.Values) > 0 {
							values = append(values, Value{
								Type: TypeArray,
								data: a,
							})
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
				sb.WriteString(fmt.Sprintf("%v", v.data))
			} else if def != "" {
				jv = NewString(def)
				sb.WriteString(def)
			}
		} else {
			sb.WriteByte(c)
			multi = true
		}
	}

	if modified && (multi || jv.Type == TypeString) {
		text = sb.String()
	}

	return
}

func isEmpty(name string, jo *Object, emptycheck map[string][]string) bool {
	if jo.IsEmpty() {
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
	if jo.IsEmpty() {
		return
	}

	params = make(map[string]Value)

	for _, jp := range jo.Properties {
		getParams(jp.Field, params)

		switch jp.Value.Type {
		case TypeString:
			s, ok := jp.Value.String()
			if ok {
				getParams(s, params)
			}
		case TypeObject:
			o, ok := jp.Value.Object()
			if ok {
				temp := o.GetParams()
				if len(temp) > 0 {
					for k, v := range temp {
						params[k] = v
					}
				}
			}
		case TypeArray:
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
		switch jv.Type {
		case TypeString:
			s, ok := jv.String()
			if ok {
				getParams(s, params)
			}
		case TypeObject:
			o, ok := jv.Object()
			if ok {
				temp := o.GetParams()
				if len(temp) > 0 {
					for k, v := range temp {
						params[k] = v
					}
				}
			}
		case TypeArray:
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
				params[name] = NewString(def)
			}
		}
	}
}
