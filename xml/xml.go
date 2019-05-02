package xml

import (
	"bytes"
	"encoding/xml"
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

//Attribute represents XML attribute with name and value
type Attribute struct {
	Name  string
	Value string
}

//New creates a new node from a string
func New(b []byte) (Node, error) {
	decoder := xml.NewDecoder(bytes.NewReader(b))

	var parent *Node
	var root *Node

	for {
		t, err := decoder.Token()
		if err != nil && err != io.EOF {
			return Node{}, err
		}
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			nd := Node{Name: se.Name.Local, parent: parent}
			parent = &nd
			if root == nil {
				root = &nd
			}
			if se.Attr != nil {
				size := len(se.Attr)
				if size > 0 {
					attributes := make([]Attribute, size)
					for i, attr := range se.Attr {
						attributes[i] = Attribute{Name: attr.Name.Local, Value: attr.Value}
					}
					nd.Attributes = attributes
				}
			}
		case xml.CharData:
			if parent != nil {
				if parent.Text == "" {
					parent.Text = getText(se)
				} else {
					parent.Text += getText(se)
				}
			}
		case xml.Comment:
			if parent != nil {
				tc := Node{Name: "!", Text: string(se), parent: parent}
				if parent.Nodes == nil {
					parent.Nodes = []Node{tc}
				} else {
					parent.Nodes = append(parent.Nodes, tc)
				}
			}
		case xml.EndElement:
			if se.Name.Local == parent.Name && parent.parent != nil {
				if parent.parent.Text != "" {
					s, _ := parent.ToInlineString()
					parent.parent.Text += s
				} else if parent.parent.Nodes == nil {
					parent.parent.Nodes = []Node{*parent}
				} else {
					parent.parent.Nodes = append(parent.parent.Nodes, *parent)
				}
				parent = parent.parent
			}
		}
	}

	return *root, nil
}

//GetNode method finds child node and returns it if found
func (nd *Node) GetNode(name string) Node {
	if nd.Nodes != nil {
		for _, n := range nd.Nodes {
			if n.Name == name {
				return n
			}
		}
	}
	return Node{}
}

//GetAttribute method finds attribute and returns its value if found
func (nd *Node) GetAttribute(name string) string {
	if nd.Attributes != nil {
		for _, a := range nd.Attributes {
			if a.Name == name {
				return a.Value
			}
		}
	}
	return ""
}

//ToString method serializes Node into pretty XML string
func (nd *Node) ToString() string {
	return nd.toString(0)
}

//ToInlineString method serializes Node into condenced XML string
func (nd *Node) ToInlineString() (string, bool) {
	var (
		sb strings.Builder
		f  bool
	)

	sb.WriteString("<")
	if nd.Name == "!" {
		sb.WriteString("!--")
		sb.WriteString(nd.Text)
		sb.WriteString("-->")
		f = true
	} else {
		sb.WriteString(nd.Name)
		if nd.Attributes != nil {
			for _, a := range nd.Attributes {
				sb.WriteString(" ")
				sb.WriteString(a.Name)
				sb.WriteString("=\"")
				sb.WriteString(a.Value)
				sb.WriteString("\"")
			}
		}
		sb.WriteString(">")
		if nd.Nodes != nil {
			for _, n := range nd.Nodes {
				s, t := n.ToInlineString()
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
		sb.WriteString("</")
		sb.WriteString(nd.Name)
		sb.WriteString(">")
	}

	return sb.String(), f
}

func (nd *Node) toString(level int) string {
	var sb strings.Builder
	if level > 0 {
		i := 0
		for i <= level {
			sb.WriteString("  ")
			i++
		}
	}

	sb.WriteString("<")
	if nd.Name == "!" {
		sb.WriteString("!--")
		sb.WriteString(nd.Text)
		sb.WriteString("-->")
	} else {
		sb.WriteString(nd.Name)
		if nd.Attributes != nil {
			for _, a := range nd.Attributes {
				sb.WriteString(" ")
				sb.WriteString(a.Name)
				sb.WriteString("=\"")
				sb.WriteString(a.Value)
				sb.WriteString("\"")
			}
		}
		sb.WriteString(">")
		if nd.Nodes != nil {
			for _, n := range nd.Nodes {
				sb.WriteString("\n")
				sb.WriteString(n.toString(level + 1))
			}
			sb.WriteString("\n")
			if level > 0 {
				i := 0
				for i <= level {
					sb.WriteString("  ")
					i++
				}
			}
		}
		if nd.Text != "" {
			sb.WriteString(nd.Text)
		}
		sb.WriteString("</")
		sb.WriteString(nd.Name)
		sb.WriteString(">")
	}

	return sb.String()
}

func getText(bytes []byte) string {
	var ok bool
	start := -1

	for i, b := range bytes {
		if b > 13 {
			if start == -1 {
				start = i
			}
			if b != 32 {
				ok = true
				break
			}
		}
	}

	if ok {
		return string(bytes[start:])
	}

	return ""
}
