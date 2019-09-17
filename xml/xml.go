package xml

import (
	"errors"
	"io"
	"strings"
)

//Node represents XML node with name, optional attributes, text and children nodes
type Node struct {
	Name       string
	Attributes map[string]string
	Nodes      []Node
	Text       string
	parent     *Node
}

//New creates a new node from a io.ByteScanner
func New(scanner io.ByteScanner) (Node, error) {
	xp := &xParser{
		r: scanner,
	}

	return xp.Parse()
}

//NewFromString creates a new node from a string
func NewFromString(xml string) (Node, error) {
	if xml == "" {
		return Node{}, errors.New("Missing input")
	}

	xp := &xParser{
		r: strings.NewReader(xml),
	}

	return xp.Parse()
}

//Node method finds first child node
func (nd *Node) Node(name string) Node {
	if nd.Nodes != nil {
		for _, n := range nd.Nodes {
			if n.Name == name {
				return n
			}
		}
	}
	return Node{}
}

//Attribute method finds attribute by name
func (nd *Node) Attribute(name string) string {
	if nd.Attributes != nil {
		v, ok := nd.Attributes[name]
		if ok {
			return v
		}
	}
	return ""
}

//String method serializes Node into pretty XML string
func (nd *Node) String() string {
	return nd.toString(0)
}

//InlineString method serializes Node into condenced XML string
func (nd *Node) InlineString() (string, bool) {
	var (
		sb strings.Builder
		f  bool
	)

	sb.WriteByte('<')
	if nd.Name == "!" {
		sb.WriteByte('!')
		sb.WriteByte('-')
		sb.WriteByte('-')
		sb.WriteString(nd.Text)
		sb.WriteByte('-')
		sb.WriteByte('-')
		sb.WriteByte('>')
		f = true
	} else {
		sb.WriteString(nd.Name)
		if nd.Attributes != nil {
			for k, v := range nd.Attributes {
				sb.WriteByte(' ')
				sb.WriteString(k)
				sb.WriteByte('=')
				sb.WriteByte('"')
				sb.WriteString(v)
				sb.WriteByte('"')
			}
		}
		sb.WriteByte('>')
		if nd.Nodes != nil {
			for _, n := range nd.Nodes {
				s, t := n.InlineString()
				sb.WriteString(s)
				if t {
					f = true
				}
			}
		}
		if nd.Text != "" {
			sb.WriteString(nd.Text)
			f = true
		}
		sb.WriteByte('<')
		sb.WriteByte('/')
		sb.WriteString(nd.Name)
		sb.WriteByte('>')
	}

	return sb.String(), f
}

func (nd *Node) toString(level int) string {
	var sb strings.Builder
	if level > 0 {
		i := 0
		for i <= level {
			sb.WriteByte(' ')
			i++
		}
	}

	sb.WriteString("<")
	if nd.Name == "!" {
		sb.WriteByte('!')
		sb.WriteByte('-')
		sb.WriteByte('-')
		sb.WriteString(nd.Text)
		sb.WriteByte('-')
		sb.WriteByte('-')
		sb.WriteByte('>')
	} else {
		sb.WriteString(nd.Name)
		if nd.Attributes != nil {
			for k, v := range nd.Attributes {
				sb.WriteByte(' ')
				sb.WriteString(k)
				sb.WriteByte('=')
				sb.WriteByte('"')
				sb.WriteString(v)
				sb.WriteByte('"')
			}
		}
		sb.WriteString(">")
		if nd.Nodes != nil {
			for _, n := range nd.Nodes {
				sb.WriteByte('\n')
				sb.WriteString(n.toString(level + 1))
			}
			sb.WriteByte('\n')
			if level > 0 {
				i := 0
				for i <= level {
					sb.WriteByte(' ')
					i++
				}
			}
		}
		if nd.Text != "" {
			sb.WriteString(nd.Text)
		}
		sb.WriteByte('<')
		sb.WriteByte('/')
		sb.WriteString(nd.Name)
		sb.WriteByte('>')
	}

	return sb.String()
}
