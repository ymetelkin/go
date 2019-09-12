package appl

import (
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type renditions struct {
	Renditions    json.Array
	Counts        map[string]int
	NonRenditions []json.Property
}

func (rnds *renditions) GetRendition(title string, role string, mt mediaType, nd xml.Node, chars xml.Node) json.Object {
	jo := json.Object{}
	jo.AddString("title", title)
	jo.AddString("rel", role)
	jo.AddString("type", string(mt))

	if nd.Attributes != nil {
		for k, v := range nd.Attributes {
			addAttribute(k, v, &jo)
		}
	}

	if chars.Nodes != nil {
		n := chars.Node("Scenes")
		n = n.Node("Scene")
		if n.Text != "" {
			jo.AddString("scene", n.Text)
		}
		id := n.Attribute("Id")
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
			p := n.Node("Presentation")
			if p.Attributes != nil || p.Nodes != nil {
				system := p.Attribute("System")
				if system != "" {
					jo.AddString("presentationsystem", system)
				}

				ch := p.Node("Characteristics")
				if ch.Attributes != nil {
					for k, v := range ch.Attributes {
						switch k {
						case "Frame":
							if v != "" {
								jo.AddString("presentationframe", v)
							}
						case "FrameLocation":
							if v != "" {
								jo.AddString("presentationframelocation", v)
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
			v := n.Attribute("Name")
			if v != "" && strings.HasPrefix(strings.ToLower(v), "broadcastformat") {
				runes := []rune(v)
				jo.AddString("broadcastformat", string(runes[16:]))

				exit++
				if exit == 2 {
					break
				}
			}
		case "ForeignKeys":
			if nd.Name == "VideoContentItem" && !tape {
				system := n.Attribute("System")
				if strings.EqualFold(system, "Tape") && n.Nodes != nil {
					for _, k := range n.Nodes {
						if k.Name == "Keys" && k.Attributes != nil {
							var (
								ok bool
								id string
							)
							for k, v := range k.Attributes {
								switch k {
								case "Field":
									ok = strings.EqualFold(v, "Number")
								case "Id":
									id = v
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
		for k, v := range chars.Attributes {
			addAttribute(k, v, &jo)
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

	if multiple || !ok {
		rnds.Renditions.AddObject(jo)
	}
}

func (rnds *renditions) AddNonRenditions(jo *json.Object) {
	if rnds.NonRenditions != nil {
		for _, jp := range rnds.NonRenditions {
			jo.AddProperty(jp)
		}
	}
}

func (rnds *renditions) AddRenditions(jo *json.Object) {
	if !rnds.Renditions.IsEmpty() {
		jo.AddArray("renditions", rnds.Renditions)
	}
}

func addAttribute(k string, v string, jo *json.Object) {
	if v != "" {
		name := strings.ToLower(k)
		switch name {
		case "id":
			jo.AddString("code", strings.ToLower(v))
		case "mediatype":
			jo.AddString("type", v)
		case "sizeinbytes":
			i, err := strconv.Atoi(v)
			if err == nil && i > 0 {
				jo.AddInt("sizeinbytes", i)
			}
		default:
			jo.AddString(name, v)
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
		ext = nd.Attribute("FileExtension")
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
