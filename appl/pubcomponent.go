package appl

import (
	"strconv"
	"strings"

	"github.com/ymetelkin/go/xml"
)

func (doc *Document) parsePublicationComponent(node xml.Node) {
	if node.Nodes == nil || node.Attributes == nil {
		return
	}

	var role, mediatype string
	for k, v := range node.Attributes {
		switch k {
		case "Role":
			role = v
		case "MediaType":
			mediatype = v
		}
	}

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "TextContentItem":
			doc.parseTextContentItem(nd, role, mediatype)
		case "PhotoContentItem":
		case "PhotoCollectionContentItem":
		case "GraphicContentItem":
		case "VideoContentItem":
		case "WebPartContentItem":
		case "AudioContentItem":
		case "ComplexDataContentItem":
		}
	}
}

func (doc *Document) parseTextContentItem(node xml.Node, role string, mediatype string) {
	switch role {
	case "Main":
		if doc.Story != nil {
			return
		}
	case "Caption":
		if doc.Caption != nil {
			return
		}
	case "Script":
		if doc.Script != nil {
			return
		}
	case "Shotlist":
		if doc.Shotlist != nil {
			return
		}
	case "PublishableEditorNotes":
		if doc.PublishableEditorNotes != nil {
			return
		}
	default:
		return
	}

	nd := node.Node("DataContent")
	nd = nd.Node("nitf")
	if nd.Nodes == nil {
		return
	}

	bc := bodyContentNode(nd)
	if bc == nil {
		return
	}

	var (
		sb   strings.Builder
		text Text
	)

	for _, block := range bc.Nodes {
		if block.Name == "block" {
			if block.Nodes != nil {
				var (
					txt strings.Builder
					add bool
				)
				for _, p := range block.Nodes {
					s, b := p.InlineString()
					if b {
						add = true
					}
					txt.WriteString(s)
				}

				if add {
					sb.WriteString(txt.String())
				}
			}

			if block.Text != "" {
				sb.WriteString(block.Text)
			}
		}
	}

	text.Body = sb.String()
	if text.Body == "" {
		return
	}

	nd = node.Node("Characteristics")
	nd = nd.Node("Words")
	text.Words, _ = strconv.Atoi(nd.Text)

	switch role {
	case "Main":
		doc.Story = &text
	case "Caption":
		doc.Caption = &text
	case "Script":
		doc.Script = &text
	case "Shotlist":
		doc.Shotlist = &text
	case "PublishableEditorNotes":
		doc.PublishableEditorNotes = &text
	}
}

func bodyContentNode(node xml.Node) *xml.Node {
	if node.Nodes != nil {
		for _, nd := range node.Nodes {
			if nd.Name == "body.content" {
				return &nd
			}

			test := bodyContentNode(nd)
			if test.Nodes != nil {
				return test
			}
		}
	}
	return nil
}
