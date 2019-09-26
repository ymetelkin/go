package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/xml"
)

type renditionParser struct {
	pm, pp, pt bool
}

func (doc *Document) parsePublicationComponent(node xml.Node, parser *renditionParser) {
	if node.Nodes == nil || node.Attributes == nil {
		return
	}

	var role, mediatype string

	for k, v := range node.Attributes {
		switch k {
		case "Role":
			role = v
		case "MediaType":
			mediatype = strings.ToLower(v)
		}
	}

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "TextContentItem":
			doc.parseTextContentItem(nd, role)
		case "PhotoContentItem":
			doc.parsePhotoContentItem(nd, role, mediatype, parser)
		case "VideoContentItem":
			doc.parseVideoContentItem(nd, role, mediatype)
		case "AudioContentItem":
			doc.parseAudioContentItem(nd, role, mediatype)
		case "PhotoCollectionContentItem":
			doc.parsePhotoCollectionContentItem(nd, role, mediatype)
		case "GraphicContentItem":
			doc.parseGraphicContentItem(nd, role, mediatype)
		case "WebPartContentItem":
			doc.parseWebPartContentItem(nd, role, mediatype)
		case "ComplexDataContentItem":
			doc.parseComplexDataContentItem(nd, role, mediatype)
		}
	}
}

func (doc *Document) parseTextContentItem(node xml.Node, role string) {
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

func (doc *Document) parsePhotoContentItem(node xml.Node, role string, mediatype string, parser *renditionParser) {
	r := Rendition{
		Rel:       role,
		MediaType: mediatype,
	}

	switch doc.MediaType {
	case "photo":
		switch role {
		case "Main":
			if parser.pm {
				return
			}
			parser.pm = true
			r.Title = "Full Resolution (JPG)"
		case "Preview":
			if parser.pp {
				return
			}
			parser.pp = true
			r.Title = "Preview (JPG)"
		case "Thumbnail":
			if parser.pt {
				return
			}
			parser.pt = true
			r.Title = "Thumbnail (JPG)"
		default:
			return
		}
		r.parse(node)
	case "video":
		switch role {
		case "Preview":
			r.parse(node)
			r.setExtTitle("Preview")
		case "Thumbnail":
			r.parse(node)
			r.setDimensionsTitle(role)
		default:
			return
		}
	case "graphic", "complexdata":
		switch role {
		case "Preview":
			if parser.pp {
				return
			}
			parser.pp = true
			r.Title = "Preview (JPG)"
		case "Thumbnail":
			if parser.pt {
				return
			}
			parser.pt = true
			r.Title = "Thumbnail (JPG)"
		default:
			return
		}
		r.parse(node)
	}
	doc.Renditions = append(doc.Renditions, r)
}

func (doc *Document) parseVideoContentItem(node xml.Node, role string, mediatype string) {
	if doc.MediaType != "video" {
		return
	}

	r := Rendition{
		Rel:       role,
		MediaType: mediatype,
	}

	switch role {
	case "PhysicalMain":
		r.parse(node)
		r.Title = r.PhysicalType
	case "Main":
		r.parse(node)
		if strings.EqualFold(r.FileExtension, "TXT") {
			return
		}
		r.setDimensionsTitle("Full Resolution")
	case "Preview":
		r.parse(node)
		r.setExtTitle("Preview")
	default:
		return
	}
	doc.Renditions = append(doc.Renditions, r)
}

func (doc *Document) parseAudioContentItem(node xml.Node, role string, mediatype string) {
	if mediatype != "audio" || role != "Main" {
		return
	}

	r := Rendition{
		Rel:       role,
		MediaType: mediatype,
	}

	r.parse(node)
	r.setExtTitle("Full Resolution")
	doc.Renditions = append(doc.Renditions, r)
}

func (doc *Document) parseGraphicContentItem(node xml.Node, role string, mediatype string) {
	if role != "Main" {
		return
	}

	r := Rendition{
		Rel:       role,
		MediaType: mediatype,
	}

	switch doc.MediaType {
	case "graphic":
		r.parse(node)
		r.setDimensionsTitle("Full Resolution")
		doc.Renditions = append(doc.Renditions, r)

		nd := node.Node("RelatedBinaries")
		if nd.Attribute("Name") == "MatteFileName" {
			matte := Rendition{
				Rel:       role,
				MediaType: mediatype,
			}
			matte.parse(node)
			matte.setDimensionsTitle("Full Resolution Matte")
			doc.Renditions = append(doc.Renditions, matte)
		}
	case "complexdata":
		nd := node.Node("Presentations")
		if nd.Nodes != nil {
			nd = nd.Node("Presentation")
			if nd.Nodes != nil {
				nd = nd.Node("Characteristics")
				r.Title = nd.Attribute("FrameLocation")
				if r.Title != "" {
					r.parse(node)
					doc.Parts = append(doc.Parts, r)
				}
			}
		}
	}
}

func (doc *Document) parseComplexDataContentItem(node xml.Node, role string, mediatype string) {
	if mediatype != "complexdata" || role != "Main" {
		return
	}

	r := Rendition{
		Rel:       role,
		MediaType: mediatype,
	}

	r.parse(node)
	r.setExtTitle("Full Resolution")
	doc.Renditions = append(doc.Renditions, r)
}

func (doc *Document) parseWebPartContentItem(node xml.Node, role string, mediatype string) {
	if doc.MediaType != "video" {
		return
	}

	r := Rendition{
		Rel:       role,
		MediaType: mediatype,
		Title:     "Web",
	}
	r.parse(node)
	r.ForeignKeys = parseForeignKeys(node)

	doc.Renditions = append(doc.Renditions, r)
}

func (doc *Document) parsePhotoCollectionContentItem(node xml.Node, role string, mediatype string) {
	if doc.MediaType != "video" || role != "Thumbnail" || node.Nodes == nil {
		return
	}

	const (
		TokenDot rune = 46
		Token0   rune = 48
	)

	var (
		ext, sep string
		href     bool
		w, h     int
		zeros    []rune
		times    []int
	)

	url := node.Attribute("BaseFileName")
	if url != "" {
		runes := []rune(url)
		i := len(runes) - 3

		for i >= 0 {
			r := runes[i]
			if href {
				if r == Token0 {
					if zeros == nil {
						zeros = []rune{r}
					} else {
						zeros = append(zeros, r)
					}
				} else {
					sep = string(r)
					url = string(runes[0:i])
					break
				}
			} else if r == TokenDot {
				ext = string(runes[i+1:])
				href = true
			}
			i--
		}
	}

	if sep == "" {
		sep = "/"
	}

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "File":
			v := nd.Attribute("TimeOffSetMilliseconds")
			tm, _ := strconv.Atoi(v)
			times = append(times, tm)
		case "Characteristics":
			if nd.Nodes != nil {
				for _, n := range nd.Nodes {
					switch nd.Name {
					case "Width":
						w, _ = strconv.Atoi(n.Text)
					case "Height":
						h, _ = strconv.Atoi(n.Text)
					}
				}
			}
		}
	}

	if times != nil && len(times) > 0 {
		var (
			last = len(times) - 1
			end  int
		)

		for i, ms := range times {
			shot := PhotoShot{
				Sequence:  i + 1,
				Width:     w,
				Height:    h,
				StartTime: ms,
			}

			if href {
				s := strconv.Itoa(i)
				z := string(zeros[0 : len(zeros)-len(s)])
				s = fmt.Sprintf("%s%s%s%s.%s", url, sep, z, s, ext)
				shot.Href = s
			}

			if i < last {
				end = times[i+1]
			} else {
				end = 0
			}
			shot.EndTime = end

			doc.Shots = append(doc.Shots, shot)
		}
	}
}

func bodyContentNode(node xml.Node) *xml.Node {
	if node.Nodes != nil {
		for _, nd := range node.Nodes {
			switch nd.Name {
			case "body.content":
				return &nd
			case "body":
				return bodyContentNode(nd)
			default:
				continue
			}
		}
	}
	return nil
}

func (r *Rendition) parse(node xml.Node) {
	if node.Attributes != nil {
		for k, v := range node.Attributes {
			if k == "Id" {
				r.Code = strings.ToLower(v)
			} else {
				if r.Attributes == nil {
					r.Attributes = make(map[string]string)
				}
				r.Attributes[k] = v
			}
		}
	}

	var pr, bc, sc bool

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "ForeignKeys":
			if node.Name == "VideoContentItem" && r.TapeNumber == "" {
				system := nd.Attribute("System")
				if strings.EqualFold(system, "Tape") && nd.Nodes != nil {
					for _, n := range nd.Nodes {
						if n.Name == "Keys" && n.Attributes != nil {
							var (
								ok bool
								id string
							)
							for k, v := range n.Attributes {
								switch k {
								case "Field":
									ok = strings.EqualFold(v, "Number")
								case "Id":
									id = v
								}
							}
							if ok && id != "" {
								r.TapeNumber = id
								break
							}
						}
					}
				}
			}
		case "Presentations":
			if pr {
				continue
			}
			n := nd.Node("Presentation")
			r.PresentationSystem = n.Attribute("System")
			ch := n.Node("Characteristics")
			if ch.Attributes != nil {
				for k, v := range ch.Attributes {
					switch k {
					case "Frame":
						r.PresentationFrame = v
					case "FrameLocation":
						r.PresentationFrameLocation = v
					}
				}
			}
			pr = true
		case "Property":
			if bc {
				continue
			}
			name := nd.Attribute("Name")
			if name != "" && strings.HasPrefix(strings.ToLower(name), "broadcastformat") {
				r.BroadcastFormat = string(name[16:])
				bc = true
			}
		}
	}

	chars := node.Node("Characteristics")
	if chars.Attributes != nil {
		for k, v := range chars.Attributes {
			switch k {
			case "FileExtension":
				r.FileExtension = v
			case "SizeInBytes":
				r.ByteSize, _ = strconv.Atoi(v)
			default:
				if r.Characteristics == nil {
					r.Characteristics = make(map[string]string)
				}
				r.Characteristics[k] = v
			}
		}
	}

	if chars.Nodes == nil {
		return
	}

	for _, nd := range chars.Nodes {
		switch nd.Name {
		case "Height":
			r.Height, _ = strconv.Atoi(nd.Text)
		case "Width":
			r.Width, _ = strconv.Atoi(nd.Text)
		case "Resolution":
			r.Resolution, _ = strconv.Atoi(nd.Text)
		case "ResolutionUnits":
			r.ResolutionUnits = nd.Text
		case "FrameRate":
			r.FrameRate, _ = strconv.ParseFloat(nd.Text, 64)
		case "TotalDuration":
			r.TotalDuration, _ = strconv.Atoi(nd.Text)
		case "PhysicalType":
			r.PhysicalType = nd.Text
			r.Characteristics[nd.Name] = nd.Text
		case "Scenes":
			if !sc {
				n := nd.Node("Scene")
				r.Scene = n.Text
				r.SceneID = n.Attribute("Id")
				sc = true
			}
		default:
			if r.Characteristics == nil {
				r.Characteristics = make(map[string]string)
			}
			r.Characteristics[nd.Name] = nd.Text
		}
	}
}

func (r *Rendition) setExtTitle(role string) {
	if r.FileExtension != "" {
		r.Title = fmt.Sprintf("%s (%s)", role, strings.ToUpper(r.FileExtension))
	} else {
		r.Title = role
	}
}

func (r *Rendition) setDimensionsTitle(role string) {
	var sb strings.Builder
	if r.FileExtension != "" {
		sb.WriteString(strings.ToUpper(r.FileExtension))
	}
	if r.Width > 0 && r.Height > 0 {
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(r.Width))
		sb.WriteString("x")
		sb.WriteString(strconv.Itoa(r.Height))
	}
	if sb.Len() > 0 {
		r.Title = fmt.Sprintf("%s (%s)", role, sb.String())
	} else {
		r.Title = role
	}
}
