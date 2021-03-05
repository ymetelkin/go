package v1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type parser struct {
	buf  []byte
	i    int
	last int
}

func newParser(bs []byte) *parser {
	return &parser{
		buf:  bs,
		last: len(bs) - 1,
	}
}

func (p *parser) Read() (b byte, ok bool) {
	if p.i > p.last {
		return
	}
	b = p.buf[p.i]
	p.i++
	ok = true
	return
}

func (p *parser) Move(i int) (ok bool) {
	test := p.i + i
	if test < 0 || test > p.last {
		return
	}
	p.i = test
	ok = true
	return
}

//ParseObject parses new JSON object from bytes
func ParseObject(bs []byte) (jo Object, err error) {
	p := newParser(bs)

	c, ok := p.SkipWS()
	if !ok {
		err = errors.New("Missing JSON input")
		return
	}

	if c == '{' {
		jo, _, err = p.ParseObject()
	} else {
		err = fmt.Errorf("Expected '{', found '%c'", c)
	}

	return
}

//ParseArray parses new JSON object from scanner
func ParseArray(bs []byte) (ja Array, err error) {
	p := newParser(bs)

	c, ok := p.SkipWS()
	if !ok {
		err = errors.New("Missing JSON input")
		return
	}

	if c == '[' {
		ja, _, err = p.ParseArray()
	} else {
		err = fmt.Errorf("Expected '[', found '%c'", c)
	}

	return
}

func (p *parser) Parse() (jv Value, params bool, err error) {
	c, ok := p.SkipWS()
	if !ok {
		err = errors.New("Missing input")
		return
	}

	if c == '{' {
		jo, ps, err := p.ParseObject()
		if err == nil {
			params = ps
			jv = O(jo)
		}
	} else if c == '[' {
		ja, ps, err := p.ParseArray()
		if err == nil {
			params = ps
			jv = A(ja)
		}
	} else {
		err = fmt.Errorf("Expected '{' or '[', found '%c'", c)
	}

	return
}

func (p *parser) ParseObject() (jo Object, params bool, err error) {
	var ps, eoo bool

	ps, eoo, err = p.AddProperty(&jo)
	if err != nil || eoo {
		return
	}
	if ps {
		params = true
	}

	for {
		c, ok := p.SkipWS()
		if !ok {
			err = errors.New("Expected '}' or ',', found EOF")
			return
		}

		if c == ',' {
			ps, eoo, err = p.AddProperty(&jo)
			if err != nil {
				return
			}
			if ps {
				params = true
			}
			if eoo {
				break
			}
		} else if c == '}' {
			break
		} else {
			err = fmt.Errorf("Expected '}', found '%c'", c)
			break
		}
	}

	return
}

func (p *parser) AddProperty(jo *Object) (params bool, eoo bool, err error) {
	c, ok := p.SkipWS()
	if !ok {
		err = errors.New("Expected '}' or field name, found EOF")
		return
	}

	if c == '"' {
		var (
			name   string
			jv     Value
			ps     bool
			pvalue int
		)

		name, ps, err = p.ParsePropertyName()
		if err != nil {
			return
		}
		if ps {
			params = true
			pvalue = 1

			if len(jo.params) == 0 {
				jo.params = make(map[string]int)
			}
			jo.params[name] = pvalue
		}

		jv, ps, err = p.ParseValue()
		if err != nil {
			return
		}

		if jv != nil && jv.t() > 0 {
			if ps {
				params = true

				if pvalue == 0 && len(jo.params) == 0 {
					jo.params = make(map[string]int)
				}
				pvalue = pvalue + 2
				jo.params[name] = pvalue
			}

			jo.Add(name, jv)
		}
	} else if c == '}' {
		eoo = true
	} else {
		err = fmt.Errorf("Expected '}', found '%c'", c)
	}

	return
}

func (p *parser) ParsePropertyName() (name string, params bool, err error) {
	var (
		sb  strings.Builder
		pos int
	)

	for {
		c, ok := p.Read()
		if !ok {
			err = errors.New("Expected field name, found EOF")
			return
		}

		if c == '"' {
			//c, ok = p.Read()
			c, ok := p.SkipWS()
			if !ok {
				err = errors.New("Expected ':', found EOF")
				return
			}

			if c != ':' {
				err = fmt.Errorf("Expected ':', found '%c'", c)
				break
			}
			name = sb.String()
			if name == "" {
				err = errors.New("Missing field name")
			}
			break
		} else {
			if isParam(c) {
				params = true
			} else {
				if pos == 0 {
					if !isAlpha(c) && c != '_' {
						err = fmt.Errorf("Invalid first character of property name: '%c'", c)
						break
					}
				} else {
					if !isProperty(c) {
						err = fmt.Errorf("Invalid character in property name: '%c'", c)
						break
					}
				}
			}

			sb.WriteByte(c)
			pos++
		}
	}

	return
}

func (p *parser) ParseValue() (jv Value, params bool, err error) {
	c, ok := p.SkipWS()
	if !ok {
		err = errors.New("Expected field value, found EOF")
		return
	}

	if c == '"' {
		i := p.i - 1

		for {
			c, ok := p.Read()
			if !ok {
				err = errors.New("Expected field string value, found EOF")
				return
			}
			if c == '"' {
				if p.buf[p.i-2] == '\\' { //escaped quote is ignored
					continue
				}
				q := string(p.buf[i:p.i])
				uq, e := strconv.Unquote(q)
				if e != nil {
					q = strings.ReplaceAll(q, "\\/", "/") //weird escaping of /
					uq, e = strconv.Unquote(q)
					if e != nil {
						err = e
						return
					}
				}
				jv = String(uq)
				return
			} else if isParam(c) {
				params = true
			}
		}
	}

	if c == '{' {
		jo, ps, e := p.ParseObject()
		if e == nil {
			params = ps
			jv = O(jo)
		} else {
			err = e
		}
		return
	}

	if c == '[' {
		ja, ps, e := p.ParseArray()
		if e == nil {
			params = ps
			jv = A(ja)
		} else {
			err = e
		}
		return
	}

	if c == 't' {
		c, ok = p.SkipString("rue")
		if ok {
			jv = Bool(true)
		}
		p.Move(-1)
	} else if c == 'f' {
		c, ok = p.SkipString("alse")
		if ok {
			jv = Bool(false)
		}
		p.Move(-1)
	} else if c == 'n' {
		c, ok = p.SkipString("ull")
		if ok {
			jv = Null()
		}
		p.Move(-1)
	} else if (c >= '0' && c <= '9') || c == '-' {
		var (
			sb    strings.Builder
			float bool
		)

		sb.WriteByte(c)

		for {
			c, ok := p.Read()
			if !ok {
				err = errors.New("Expected digit, found EOF")
				return
			}

			if c == '.' || c == 'e' || c == 'E' || c == '-' || c == '+' {
				float = true
				sb.WriteByte(c)
			} else if c >= '0' && c <= '9' {
				sb.WriteByte(c)
			} else {
				p.Move(-1)

				s := sb.String()
				if float {
					f, e := strconv.ParseFloat(s, 64)
					if e == nil {
						jv = Float(f)
						return
					}
				} else {
					i, e := strconv.Atoi(s)
					if e == nil {
						jv = Int(i)
						return
					}
				}

				err = fmt.Errorf("Expected number, found '%s'", s)
				return
			}
		}
	}

	return
}

func (p *parser) ParseArray() (ja Array, params bool, err error) {
	var ps bool

	ps, err = p.AddArrayValue(&ja)
	if err != nil {
		return
	}
	if ps {
		params = true
	}

	for {
		c, ok := p.SkipWS()
		if !ok {
			err = errors.New("Expected ']' or ',', found EOF")
			return
		}

		if c == ',' {
			ps, e := p.AddArrayValue(&ja)
			if e != nil {
				err = e
				return
			}
			if ps {
				params = true
			}
		} else if c == ']' {
			break
		} else {
			err = fmt.Errorf("Expected ']', found '%c'", c)
			break
		}
	}

	return
}

func (p *parser) AddArrayValue(ja *Array) (params bool, err error) {
	var jv Value

	jv, params, err = p.ParseValue()
	if jv == nil {
		p.Move(-1)
		return
	}
	if err == nil && jv.t() > 0 {
		ja.Values = append(ja.Values, jv)

		if params {
			ja.params = append(ja.params, len(ja.Values)-1)
		}
	}

	return
}

func (p *parser) SkipWS() (c byte, ok bool) {
	for {
		c, ok = p.Read()
		if !ok || !isWS(c) {
			break
		}
	}

	return
}

func (p *parser) SkipString(s string) (c byte, ok bool) {
	bs := []byte(s)
	for _, exp := range bs {
		c, ok = p.Read()
		if !ok {
			break
		}

		if c != exp {
			return
		}
	}

	if ok {
		c, ok = p.SkipWS()
	}

	return
}

func isWS(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r' || c == '\f' || c == '\v' || c == '\b'
}

func isAlpha(c byte) bool {
	if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
		return false
	}
	return true
}

func isProperty(c byte) bool {
	return isAlpha(c) || c == '_' || c == '-' || c == '.' || (c >= '0' && c <= '9') || c == '%'
}

func isParam(c byte) bool {
	return c == '$' || c == '{' || c == '?' || c == '}'
}