package appl

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func (doc *document) ParsePhotoComponent(pc pubcomponent, parent *json.Object) error {
	nd := pc.Node.GetNode("DataContent")
	nd = nd.GetNode("nitf")
	if nd.Nodes == nil {
		return errors.New("[nitf] block is missing")
	}

	nd = getBodyContent(nd)
	if nd.Nodes == nil {
		return errors.New("[body.content] block is missing")
	}

	var sb strings.Builder

	for _, b := range nd.Nodes {
		if b.Name == "block" {
			s := b.ToInlineString()
			sb.WriteString(s)
		}
	}

	nitf := sb.String()
	if nitf == "" {
		return errors.New("[block] blocks are missing or empty")
	}

	jo := json.Object{}
	jo.AddString("nitf", nitf)

	nd = pc.Node.GetNode("Characteristics")
	nd = nd.GetNode("Words")
	i, err := strconv.Atoi(nd.Text)
	if err == nil {
		jo.AddInt("words", i)
	}

	parent.AddObject(pc.Role, jo)

	return nil
}

func getComponentMeta(title string, role string, mt MediaType, nd xml.Node, chars xml.Node) json.Object {
	jo := json.Object{}
	jo.AddString("title", title)
	jo.AddString("rel", role)
	jo.AddString("type", string(mt))

	if nd.Attributes != nil {
		for _, a := range nd.Attributes {
			addAttribute(a, &jo)
		}
	}

	if nd.Name == "VideoContentItem" {

	}

	if chars.Nodes != nil {
		n := chars.GetNode("Scenes")
		n = n.GetNode("Scene")
		if n.Text != "" {
			jo.AddString("scene", n.Text)
		}
		id := n.GetAttribute("Id")
		if id != "" {
			jo.AddString("sceneid", id)
		}
	}

	exit := 0
	for _, n := range nd.Nodes {
		switch n.Name {
		case "Presentations":
			p := n.GetNode("Presentation")
			if p.Attributes != nil || p.Nodes != nil {
				system := p.GetAttribute("System")
				if system != "" {
					jo.AddString("presentationsystem", system)
				}

				ch := p.GetNode("Characteristics")
				if ch.Attributes != nil {
					for _, a := range ch.Attributes {
						switch a.Name {
						case "Frame":
							if a.Value != "" {
								jo.AddString("presentationframe", a.Value)
							}
						case "FrameLocation":
							if a.Value != "" {
								jo.AddString("presentationframelocation", a.Value)
							}
						}
					}
				}
			}
			exit++
			if exit == 2 {
				break
			}
		case "Property":
			name := n.GetAttribute("Name")
			if name != "" && strings.HasPrefix(strings.ToLower(name), "broadcastformat") {
				runes := []rune(name)
				jo.AddString("broadcastformat", string(runes[15:]))

				exit++
				if exit == 2 {
					break
				}
			}
		}
	}

	if chars.Attributes != nil {
		for _, a := range chars.Attributes {
			addAttribute(a, &jo)
		}
	}

	if chars.Nodes != nil {
		for _, n := range chars.Nodes {
			addNode(n, &jo)
		}
	}

	return jo
}

func addAttribute(a xml.Attribute, jo *json.Object) {
	if a.Value != "" {
		name := strings.ToLower(a.Name)
		switch name {
		case "id":
			jo.AddString("code", strings.ToLower(a.Value))
		case "mediatype":
			jo.AddString("type", a.Value)
		case "sizeinbytes":
			i, err := strconv.Atoi(a.Value)
			if err == nil && i > 0 {
				jo.AddInt("sizeinbytes", i)
			}
		default:
			jo.AddString(name, a.Value)
		}
	}
}

func addNode(nd xml.Node, jo *json.Object) {
	if nd.Text != "" {
		name := strings.ToLower(nd.Name)
		switch name {
		case "framerate":
			f, err := strconv.ParseFloat(nd.Text, 64)
			if err == nil && f > 0 {
				jo.AddFloat(name, f)
			}
		case "totalduration", "resolution", "width", "height":
			i, err := strconv.Atoi(nd.Text)
			if err == nil && i > 0 {
				jo.AddInt(name, i)
			}
		default:
			jo.AddString(name, nd.Text)
		}
	}
}
