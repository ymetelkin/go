package xml

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

//Node represents XML node with name, optional attributes, text and children nodes
type Node struct {
	Name       string
	Attributes []Attribute
	Nodes      []Node
	Text       string
	parent     *Node
}

//Attribute XML node attribute
type Attribute struct {
	Name  string
	Value string
}

//Parse creates a new node from a io.ByteScanner
func Parse(scanner io.ByteScanner) (Node, error) {
	xp := &xParser{
		r: scanner,
	}

	return xp.Parse()
}

//ParseString creates a new node from a string
func ParseString(xml string) (Node, error) {
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
		for _, a := range nd.Attributes {
			if a.Name == name {
				return a.Value
			}
		}
	}
	return ""
}

//Matches compares two nodes
func (nd *Node) Matches(other *Node) (match bool, s string) {
	if nd == nil {
		s = "Left is nil"
		return
	}
	if other == nil {
		s = "Right is nil"
		return
	}

	if nd.Name != other.Name {
		s = fmt.Sprintf("Name mismatch: [ %s ] vs [ %s ]", nd.Name, other.Name)
		return
	}

	if nd.Text != other.Text {
		s = fmt.Sprintf("Text mismatch [ %s ]: %s", nd.Name, nd.Text)
		return
	}

	var lsize, rsize int
	if nd.Attributes != nil {
		lsize = len(nd.Attributes)
	}
	if other.Attributes != nil {
		rsize = len(other.Attributes)
	}
	if lsize != rsize {
		s = fmt.Sprintf("Attribute count mismatch: [ %d ] vs [ %d ]", lsize, rsize)
		return
	}
	if lsize > 0 {
		for i, la := range nd.Attributes {
			ra := other.Attributes[i]
			if la.Name != ra.Name {
				s = fmt.Sprintf("Attribute names mismatch: [ %s ] vs [ %s ]", la.Name, ra.Name)
				return
			}
			if la.Value != ra.Value {
				s = fmt.Sprintf("Attribute mismatch: [ %s ] vs [ %s ]", la.Value, ra.Value)
				return
			}
		}
	}

	lsize = 0
	if nd.Nodes != nil {
		lsize = len(nd.Nodes)
	}
	rsize = 0
	if other.Nodes != nil {
		rsize = len(other.Nodes)
	}
	if lsize != rsize {
		s = fmt.Sprintf("Node count mismatch: [ %d ] vs [ %d ]", lsize, rsize)
		return
	}
	if lsize > 0 {
		for i, ln := range nd.Nodes {
			rn := other.Nodes[i]
			match, s = ln.Matches(&rn)
			if !match {
				return
			}
		}
	}

	match = true
	return
}

//String method serializes Node into pretty XML string
func (nd *Node) String() string {
	return nd.toString(0)
}

//InlineString method serializes Node into condenced XML string
func (nd *Node) InlineString() string {
	var sb strings.Builder

	sb.WriteByte('<')
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
			for _, a := range nd.Attributes {
				sb.WriteByte(' ')
				sb.WriteString(a.Name)
				sb.WriteByte('=')
				sb.WriteByte('"')
				sb.WriteString(a.Value)
				sb.WriteByte('"')
			}
		}
		sb.WriteByte('>')
		if nd.Nodes != nil {
			for _, n := range nd.Nodes {
				s := n.InlineString()
				sb.WriteString(s)
			}
		}
		if nd.Text != "" {
			sb.WriteString(strings.TrimSpace(nd.Text))
		}
		sb.WriteByte('<')
		sb.WriteByte('/')
		sb.WriteString(nd.Name)
		sb.WriteByte('>')
	}

	return sb.String()
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
			for _, a := range nd.Attributes {
				sb.WriteByte(' ')
				sb.WriteString(a.Name)
				sb.WriteByte('=')
				sb.WriteByte('"')
				sb.WriteString(a.Value)
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
