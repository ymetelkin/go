package xml

import (
	"errors"
	"fmt"
	"io"
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

type xParser struct {
	r io.ByteScanner
}

type reader struct {
	s    string
	i    int
	size int
}

func newReader(s string) *reader {
	return &reader{
		s:    s,
		size: len(s) - 1,
	}
}

func (r *reader) ReadByte() (byte, error) {
	if r.i > r.size {
		return 0, io.EOF
	}
	b := r.s[r.i]
	r.i++
	return b, nil
}

func (r *reader) UnreadByte() error {
	if r.i <= 0 {
		return errors.New("reader.UnreadByte: at beginning of string")
	}
	r.i--
	return nil
}

//Parse parses Node from bytes
func Parse(bs []byte) (nd Node, err error) {
	p := newParser(bs)

	c, ok := p.SkipWS()
	if !ok {
		err = errors.New("Missing XML input")
		return
	}

	if c != '<' {
		err = fmt.Errorf("Expected '<', found '%c'", c)
		return
	}

	c, ok = p.Read()
	if c == '?' {
		c, e = xp.SkipString("xml ")
		if err != nil {
			return
		}

		ok, c, e := xp.Find('?')
		if !ok {
			err = fmt.Errorf("Expected '?', found '%c'", c)
			return
		}
		if e != nil {
			err = e
			return
		}

		ok, c, e = xp.Find('>')
		if !ok {
			err = fmt.Errorf("Expected '>', found '%c'", c)
			return
		}
		if e != nil {
			err = e
			return
		}

		c, e = xp.SkipWS()
		if e == io.EOF {
			err = errors.New("Missing input")
			return
		}

		if e != nil {
			err = e
			return
		}

		if c != '<' {
			err = fmt.Errorf("Expected '<', found '%c'", c)
			return
		}
	} else {
		xp.r.UnreadByte()
	}

	nd, closed, e := xp.StartNode()
	if e != nil {
		err = e
		return
	}

	if !closed {
		err = xp.EndNode(&nd)
	}

	return
}

func (xp *xParser) StartNode() (nd Node, closed bool, err error) {
	var (
		c  byte
		sb strings.Builder
	)

	c, err = xp.SkipWS()
	if err != nil {
		return
	}

	if c == '>' || c == '/' {
		err = errors.New("Missing node name")
		return
	}

	if !isAlpha(c) {
		if c == '!' {
			c, err = xp.SkipString("--")
			if err != nil {
				return
			}

			ok, c, e := xp.Find('-')
			if !ok {
				err = fmt.Errorf("Expected '-', found '%c'", c)
				return
			}
			if e != nil {
				err = e
				return
			}

			c, err = xp.SkipString("->")
			if err != nil {
				return
			}

			if c != '<' {
				err = fmt.Errorf("Expected '<', found '%c'", c)
				return
			}

			return xp.StartNode()
		}

		err = fmt.Errorf("Expected alpha character, found '%c'", c)
		return
	}

	sb.WriteByte(c)

	for {
		c, err = p.Read()
		if err != nil {
			break
		}

		if isWS(c) {
			nd.Name = sb.String()
			break
		} else if c == '>' {
			nd.Name = sb.String()
			return
		} else if c == '/' {
			c, err = xp.SkipWS()
			if err != nil {
				return
			}
			if c == '>' {
				nd.Name = sb.String()
				closed = true
				return
			}
			err = fmt.Errorf("Expected '>', found '%c'", c)
			return
		} else if !isNodeName(c) {
			err = fmt.Errorf("Expected alpha character, found '%c'", c)
			return
		} else {
			sb.WriteByte(c)
		}
	}

	c, err = xp.SkipWS()
	if err != nil {
		return
	}

	if c == '>' {
		nd.Name = sb.String()
		return
	}

	if c == '/' {
		c, err = xp.SkipWS()
		if err != nil {
			return
		}
		if c == '>' {
			nd.Name = sb.String()
			closed = true
			return
		}

		err = fmt.Errorf("Expected '>', found '%c'", c)
		return
	}

	sb.Reset()
	sb.WriteByte(c)

	for {
		c, err = p.Read()
		if err != nil {
			break
		}

		if isAttribute(c) {
			sb.WriteByte(c)
			continue
		}

		if isWS(c) {
			c, err = xp.SkipWS()
			if err != nil {
				return
			}
		}

		if c == '=' {
			name := sb.String()
			sb.Reset()

			c, err = xp.SkipWS()
			if err != nil {
				return
			}

			if c == '"' {
				value, e := xp.ReadUntil('"')
				if e != nil {
					err = e
					return
				}

				nd.AddAttribute(name, value)

				c, err = xp.SkipWS()
				if err != nil {
					return
				}

				if c == '>' {
					break
				}

				if c == '/' {
					c, err = xp.SkipWS()
					if err != nil {
						return
					}
					if c == '>' {
						closed = true
						break
					}

					err = fmt.Errorf("Expected '>', found '%c'", c)
					return
				}

				sb.Reset()
				sb.WriteByte(c)
			} else {
				err = fmt.Errorf("Expected '\"', found '%c'", c)
				return
			}
		} else {
			err = fmt.Errorf("Expected '=', found '%c'", c)
			return
		}
	}

	return
}

func (xp *xParser) EndNode(nd *Node) (err error) {
	var (
		c          byte
		nodes      []Node
		sb, spaces strings.Builder
		txt, nds   bool
	)

	c, err = xp.SkipWS()
	if err != nil {
		return
	}

	for {
		if c == '<' {
			c, err = p.Read()
			if err != nil {
				break
			}

			if c == '/' {
				c, err = xp.SkipString(nd.Name)
				if err != nil {
					break
				}

				if c != '>' {
					err = fmt.Errorf("Expected '>', found '%c'", c)
				}

				break
			}

			if c == '!' {
				c, err = xp.SkipString("[CDATA[")
				xp.r.UnreadByte()
				if err == nil {
					s, e := xp.ReadUntil(']')
					if e != nil {
						err = e
						return
					}

					ok, c, e := xp.Find(']')
					if !ok {
						err = fmt.Errorf("Expected ']', found '%c'", c)
						return
					}
					if e != nil {
						err = e
						return
					}

					ok, c, e = xp.Find('>')
					if !ok {
						err = fmt.Errorf("Expected '>', found '%c'", c)
						return
					}
					if e != nil {
						err = e
						return
					}

					sb.WriteString(strings.TrimSpace(s))
					txt = true

					c, err = xp.SkipWS()
					if err != nil {
						break
					}

					continue
				}
			}

			xp.r.UnreadByte()
			n, closed, err := xp.StartNode()
			if err != nil {
				break
			}

			if !closed {
				err = xp.EndNode(&n)
				if err != nil {
					break
				}
			}

			if txt {
				s := n.InlineString()
				sb.WriteString(s)

				c, err = p.Read()
				if err != nil {
					break
				}
			} else {
				nds = true
				n.parent = nd
				nodes = append(nodes, n)

				c, spaces, err = xp.CheckWS()
				if err != nil {
					break
				}
			}
		} else {
			if nds {
				for _, n := range nodes {
					sb.WriteString(n.InlineString())
				}
				nds = false
			}

			if spaces.Len() > 0 {
				sb.WriteString(spaces.String())
				spaces.Reset()
			}

			sb.WriteByte(c)
			txt = true

			c, err = p.Read()
			if err != nil {
				break
			}
		}
	}

	if txt {
		nd.Text = sb.String()
	}

	if nds {
		nd.Nodes = nodes
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

func (xp *xParser) CheckWS() (c byte, sb strings.Builder, err error) {
	for {
		c, err = p.Read()
		if err != nil || !isWS(c) {
			break
		}
		sb.WriteByte(c)
	}

	return
}

func (p *parser) SkipString(s string) (c byte, err error) {
	bs := []byte(s)
	for _, exp := range bs {
		c, err = p.Read()
		if err != nil {
			break
		}

		if c != exp {
			err = fmt.Errorf("Expected '%c', found '%c'", exp, c)
			break
		}
	}

	if err == nil {
		c, ok = p.SkipWS()
	}

	return
}

func (xp *xParser) ReadUntil(end byte) (s string, err error) {
	var sb strings.Builder

	for {
		c, err := p.Read()
		if err != nil {
			break
		}

		if c == end {
			s = sb.String()
			break
		}

		sb.WriteByte(c)
	}

	return
}

func (xp *xParser) Find(target byte) (ok bool, c byte, err error) {
	for {
		c, err = p.Read()
		if err != nil {
			break
		}

		if c == target {
			ok = true
			break
		}
	}

	return
}

func isWS(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func isAlpha(c byte) bool {
	if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
		return false
	}
	return true
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isNodeName(c byte) bool {
	return isAlpha(c) || c == '.' || c == '_' || c == '-' || isDigit(c)
}

func isAttribute(c byte) bool {
	return isAlpha(c) || c == ':' || c == '_' || c == '-' || isDigit(c)
}
