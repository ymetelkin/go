package xml

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type xParser struct {
	r io.ByteScanner
}

func (xp *xParser) Parse() (nd Node, err error) {
	c, e := xp.SkipWS()
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

	c, e = xp.r.ReadByte()
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
		c    byte
		sb   strings.Builder
		exit bool
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
		c, err = xp.r.ReadByte()
		if err != nil {
			break
		}

		if isWS(c) {
			nd.Name = sb.String()
			break
		} else if c == '>' {
			nd.Name = sb.String()
			exit = true
			break
		} else if !isAlpha(c) {
			err = fmt.Errorf("Expected alpha character, found '%c'", c)
			return
		} else {
			sb.WriteByte(c)
		}
	}

	if exit {
		return
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

	attrs := make(map[string]string)

	sb.Reset()
	sb.WriteByte(c)

	for {
		c, err = xp.r.ReadByte()
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

				attrs[name] = value

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

	if len(attrs) > 0 {
		nd.Attributes = attrs
	}

	return
}

func (xp *xParser) EndNode(nd *Node) (err error) {
	var (
		c     byte
		nodes []Node
		sb    strings.Builder
	)

	c, err = xp.SkipWS()
	if err != nil {
		return
	}

	for {
		if c == '<' {
			c, err = xp.r.ReadByte()
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

			n.parent = nd
			nodes = append(nodes, n)

			c, err = xp.SkipWS()
			if err != nil {
				break
			}
		} else {
			sb.WriteByte(c)

			c, err = xp.r.ReadByte()
			if err != nil {
				break
			}
		}
	}

	nd.Text = sb.String()

	if len(nodes) > 0 {
		nd.Nodes = nodes
	}

	return
}

func (xp *xParser) SkipWS() (c byte, err error) {
	for {
		c, err = xp.r.ReadByte()
		if err != nil || !isWS(c) {
			break
		}
	}

	return
}

func (xp *xParser) SkipString(s string) (c byte, err error) {
	bytes := []byte(s)
	for _, exp := range bytes {
		c, err = xp.r.ReadByte()
		if err != nil {
			break
		}

		if c != exp {
			err = fmt.Errorf("Expected '%c', found '%c'", exp, c)
			break
		}
	}

	if err == nil {
		c, err = xp.SkipWS()
	}

	return
}

func (xp *xParser) ReadUntil(end byte) (s string, err error) {
	var sb strings.Builder

	for {
		c, err := xp.r.ReadByte()
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
		c, err = xp.r.ReadByte()
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

func isAttribute(c byte) bool {
	return isAlpha(c) || c == ':'
}
