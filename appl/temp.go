package appl

import (
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
func NewNode(s string) (Node, error) {
	decoder := xml.NewDecoder(strings.NewReader(s))

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
				parent.Text = strings.TrimSpace(string(se))
			}
		case xml.EndElement:
			if se.Name.Local == parent.Name && parent.parent != nil {
				if parent.parent.Nodes == nil {
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

//ToString method serializes Node into XML string
func (nd *Node) ToString() string {
	return nd.toString(0)
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
	if nd.Text != "" {
		sb.WriteString(nd.Text)
	}
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
	sb.WriteString("</")
	sb.WriteString(nd.Name)
	sb.WriteString(">")

	return sb.String()
}
