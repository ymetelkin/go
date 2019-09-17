package json

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type parser struct {
	r io.ByteScanner
}

//ParseObject parses new JSON object from scanner
func ParseObject(scanner io.ByteScanner) (jo Object, err error) {
	return parseObject(scanner)
}

//ParseObjectString parses new JSON object from string
func ParseObjectString(s string) (jo Object, err error) {
	scanner := strings.NewReader(s)
	return parseObject(scanner)
}

func parseObject(scanner io.ByteScanner) (jo Object, err error) {
	p := &parser{r: scanner}

	c, err := p.SkipWS()
	if err == io.EOF {
		err = errors.New("Missing JSON input")
		return
	}

	if err != nil {
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
func ParseArray(scanner io.ByteScanner) (ja Array, err error) {
	return parseArray(scanner)
}

//ParseArrayString parses new JSON object from string
func ParseArrayString(s string) (ja Array, err error) {
	scanner := strings.NewReader(s)
	return parseArray(scanner)
}

func parseArray(scanner io.ByteScanner) (ja Array, err error) {
	p := &parser{r: scanner}

	c, err := p.SkipWS()
	if err == io.EOF {
		err = errors.New("Missing JSON input")
		return
	}

	if err != nil {
		return
	}

	if c == '[' {
		ja, _, err = p.ParseArray()
	} else {
		err = fmt.Errorf("Expected '[', found '%c'", c)
	}

	return
}

func (p *parser) Parse() (jv value, params bool, err error) {
	c, err := p.SkipWS()
	if err == io.EOF {
		err = errors.New("Missing input")
		return
	}

	if err != nil {
		return
	}

	if c == '{' {
		jo, ps, err := p.ParseObject()
		if err == nil {
			params = ps
			jv = newObject(jo)
		}
	} else if c == '[' {
		ja, ps, err := p.ParseArray()
		if err == nil {
			params = ps
			jv = newArray(ja)
		}
	} else {
		err = fmt.Errorf("Expected '{' or '[', found '%c'", c)
	}

	return
}

func (p *parser) ParseObject() (jo Object, params bool, err error) {
	var ps bool

	ps, err = p.AddProperty(&jo)
	if err != nil {
		return
	}
	if ps {
		params = true
	}

	for {
		c, e := p.SkipWS()
		if e != nil {
			err = e
			return
		}

		if c == ',' {
			ps, e = p.AddProperty(&jo)
			if e != nil {
				err = e
				return
			}
			if ps {
				params = true
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

func (p *parser) AddProperty(jo *Object) (params bool, err error) {
	c, err := p.SkipWS()
	if err != nil {
		return
	}

	if c == '"' {
		var (
			name   string
			jv     value
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

			if jo.pnames == nil {
				jo.pnames = make(map[string]int)
			}
			jo.pnames[name] = pvalue
		}

		jv, ps, err = p.ParseValue()
		if err != nil {
			return
		}

		if !jv.IsEmpty() {
			if ps {
				params = true

				if pvalue == 0 && jo.pnames == nil {
					jo.pnames = make(map[string]int)
				}
				pvalue = pvalue + 2
				jo.pnames[name] = pvalue
			}

			jo.addValue(name, jv)
		}
	} else if c != '}' {
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
		c, e := p.r.ReadByte()
		if e != nil {
			err = e
			return
		}

		if c == '"' {
			c, e = p.r.ReadByte()
			if e != nil {
				err = e
				break
			}
			if c != ':' {
				err = fmt.Errorf("Expected ':', found '%c'", c)
				break
			}
			name = sb.String()
			if name == "" {
				err = errors.New("Missing property name")
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

func (p *parser) ParseValue() (jv value, params bool, err error) {
	c, err := p.SkipWS()
	if err != nil {
		return
	}

	if c == '"' {
		var sb strings.Builder

		for {
			c, e := p.r.ReadByte()
			if e != nil {
				err = e
				return
			}

			if c == '\\' {
				c, e := p.r.ReadByte()
				if e != nil {
					err = e
					return
				}

				if c == 'r' {
					sb.WriteByte('\r')
				} else if c == 'n' {
					sb.WriteByte('\n')
				} else if c == 't' {
					sb.WriteByte('\t')
				} else if c == 'b' {
					sb.WriteByte('\b')
				} else if c == 'f' {
					sb.WriteByte('\f')
				} else if c == 'a' {
					sb.WriteByte('\a')
				} else if c == 'v' {
					sb.WriteByte('\v')
				} else if c == '"' {
					sb.WriteByte('"')
				} else if c == '\\' {
					sb.WriteByte('\\')
				}
			} else if c == '"' {
				jv = newString(sb.String())
				return
			} else {
				if isParam(c) {
					params = true
				}
				sb.WriteByte(c)
			}
		}
	}

	if c == '{' {
		jo, ps, e := p.ParseObject()
		if e == nil {
			params = ps
			jv = newObject(jo)
		} else {
			err = e
		}
		return
	}

	if c == '[' {
		ja, ps, e := p.ParseArray()
		if e == nil {
			params = ps
			jv = newArray(ja)
		} else {
			err = e
		}
		return
	}

	if c == 't' {
		c, err = p.SkipString("rue")
		if err == nil {
			jv = newBool(true)
			jv.Text = "true"
		}
		p.r.UnreadByte()
	} else if c == 'f' {
		c, err = p.SkipString("alse")
		if err == nil {
			jv = newBool(false)
			jv.Text = "false"
		}
		p.r.UnreadByte()
	} else if c == 'n' {
		c, err = p.SkipString("ull")
		if err == nil {
			jv = newNull()
			jv.Text = "null"
		}
		p.r.UnreadByte()
	} else if (c >= '0' && c <= '9') || c == '-' {
		var (
			sb    strings.Builder
			float bool
		)

		sb.WriteByte(c)

		for {
			c, e := p.r.ReadByte()
			if e != nil {
				err = e
				return
			}

			if c == '.' || c == 'e' || c == 'E' || c == '-' || c == '+' {
				float = true
				sb.WriteByte(c)
			} else if c >= '0' && c <= '9' {
				sb.WriteByte(c)
			} else {
				p.r.UnreadByte()

				s := sb.String()
				if float {
					f, e := strconv.ParseFloat(s, 64)
					if e == nil {
						jv = newFloat(f)
						jv.Text = s
						return
					}
				} else {
					i, e := strconv.Atoi(s)
					if e == nil {
						jv = newInt(i)
						jv.Text = s
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
		c, e := p.SkipWS()
		if e != nil {
			err = e
			return
		}

		if c == ',' {
			ps, e = p.AddArrayValue(&ja)
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
	var jv value

	jv, params, err = p.ParseValue()
	if err == nil && !jv.IsEmpty() {
		idx := ja.addValue(jv)

		if params {
			ja.pvalues = append(ja.pvalues, idx)
		}
	}

	return
}

func (p *parser) SkipWS() (c byte, err error) {
	for {
		c, err = p.r.ReadByte()
		if err != nil || !isWS(c) {
			break
		}
	}

	return
}

func (p *parser) SkipString(s string) (c byte, err error) {
	bytes := []byte(s)
	for _, exp := range bytes {
		c, err = p.r.ReadByte()
		if err != nil {
			break
		}

		if c != exp {
			err = fmt.Errorf("Expected '%c', found '%c'", exp, c)
			break
		}
	}

	if err == nil {
		c, err = p.SkipWS()
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
	return isAlpha(c) || c == '_' || c == '-' || c == '.' || (c >= '0' && c <= '9')
}

func isParam(c byte) bool {
	return c == '$' || c == '{' || c == '?' || c == '}'
}
