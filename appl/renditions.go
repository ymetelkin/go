package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type renditions struct {
	Renditions    map[string]json.Object
	Counts        map[string]int
	NonRenditions []json.Property
}

func (rnds *renditions) GetRendition(title string, role string, mt MediaType, nd xml.Node, chars xml.Node) json.Object {
	jo := json.Object{}
	jo.AddString("title", title)
	jo.AddString("rel", role)
	jo.AddString("type", string(mt))

	if nd.Attributes != nil {
		for _, a := range nd.Attributes {
			addAttribute(a, &jo)
		}
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

	var (
		exit int
		tape bool
	)

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
			v := n.GetAttribute("Name")
			if v != "" && strings.HasPrefix(strings.ToLower(v), "broadcastformat") {
				runes := []rune(v)
				jo.AddString("broadcastformat", string(runes[15:]))

				exit++
				if exit == 2 {
					break
				}
			}
		case "ForeignKeys":
			if nd.Name == "VideoContentItem" && !tape {
				system := n.GetAttribute("System")
				if strings.EqualFold(system, "Tape") && n.Nodes != nil {
					for _, k := range n.Nodes {
						if k.Name == "Keys" && k.Attributes != nil {
							var (
								ok bool
								id string
							)
							for _, a := range k.Attributes {
								switch a.Name {
								case "Field":
									ok = strings.EqualFold(a.Value, "Number")
								case "Id":
									id = a.Value
								}
							}
							if ok && id != "" {
								jo.AddString("tapenumber", id)
								tape = true
								break
							}
						}
					}
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

func (rnds *renditions) AddRendition(name string, jo json.Object, multiple bool) {
	if rnds.Renditions == nil {
		rnds.Renditions = make(map[string]json.Object)
	}

	if multiple {
		if rnds.Counts == nil {
			rnds.Counts = make(map[string]int)
		}

		i, ok := rnds.Counts[name]
		if ok {
			i++
		} else {
			i = 1
		}
		rnds.Counts[name] = i

		name = fmt.Sprintf("%s%d", name, i)
	} else {
		_, ok := rnds.Renditions[name]
		if ok {
			return
		}
	}

	rnds.Renditions[name] = jo
}

func (rnds *renditions) AddNonRenditions(jo *json.Object) {
	if rnds.NonRenditions != nil {
		for _, jp := range rnds.NonRenditions {
			jo.AddProperty(jp)
		}
	}
}

func (rnds *renditions) AddRenditions(jo *json.Object) {
	if rnds.Renditions != nil {
		var (
			ja  json.Array
			add bool
		)

		for _, r := range rnds.Renditions {
			ja.AddObject(r)
			add = true
		}

		if add {
			jo.AddArray("renditions", ja)
		}
	}
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

func getBinaryName(nd xml.Node, ext string, dims bool) (string, string) {
	var name, key strings.Builder

	if ext == "" {
		ext = nd.GetAttribute("FileExtension")
	}

	if ext != "" {
		name.WriteString(strings.ToUpper(ext))
		key.WriteString(strings.ToLower(ext))
	}

	if dims && nd.Nodes != nil {
		var h, w string
		for _, n := range nd.Nodes {
			switch n.Name {
			case "Width":
				w = n.Text
			case "Height":
				h = n.Text
			}
		}

		if h != "" && w != "" {
			name.WriteString(" ")
			name.WriteString(w)
			name.WriteString("x")
			name.WriteString(h)

			key.WriteString(w)
			key.WriteString("x")
			key.WriteString(h)
		}
	}

	return name.String(), key.String()
}
