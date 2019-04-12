package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (rnds *renditions) ParseVideoComponent(pc pubcomponent) int64 {
	chars := pc.Node.GetNode("Characteristics")
	if chars.Nodes == nil {
		return 0
	}

	var duration int64
	n := chars.GetNode("TotalDuration")
	if n.Text != "" {
		duration, _ = strconv.ParseInt(n.Text, 0, 64)
	}

	key := strings.ToLower(pc.Role)
	switch key {
	case "main":
		ext := pc.Node.GetAttribute("FileExtension")

		if ext != "" {
			ext = strings.ToUpper(ext)
		}

		if ext != "TXT" {
			var file string

			ofn := pc.Node.GetAttribute("OriginalFileName")
			if ofn != "" {
				toks := strings.Split(ofn, "_")
				tmp := toks[len(toks)-1]
				toks = strings.Split(tmp, ".")
				file = toks[0]
			}

			n, k := getBinaryName(chars, "", true)
			name := fmt.Sprintf("main%s%s", k, file)
			title := fmt.Sprintf("Full Resolution (%s)", n)
			jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
			rnds.AddRendition(name, jo, true)
		}
	case "physicalmain":
		pt := chars.GetNode("PhysicalType")
		if pt.Text != "" {
			name := strings.ReplaceAll(pt.Text, "-", "")
			name = strings.ReplaceAll(name, " ", "")
			name = strings.ToLower(name)
			jo := rnds.GetRendition(pt.Text, pc.Role, pc.MediaType, pc.Node, chars)
			rnds.AddRendition(name, jo, true)
		}
	case "preview":
		n, k := getBinaryName(chars, "", true)
		name := "preview" + k
		title := fmt.Sprintf("Preview (%s)", n)
		jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
		rnds.AddRendition(name, jo, true)
	case "thumbnail":
		n, k := getBinaryName(chars, "", true)
		name := "thumbnail" + k
		title := fmt.Sprintf("Thumbnail (%s)", n)
		jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
		rnds.AddRendition(name, jo, true)
	}

	return duration
}

func (rnds *renditions) ParseWebPartomponent(pc pubcomponent, mt mediaType) {
	if mt != mediaTypeVideo || !strings.EqualFold(pc.Role, "main") {
		return
	}

	chars := pc.Node.GetNode("Characteristics")
	if chars.Nodes == nil {
		return
	}

	jo := rnds.GetRendition("Web", pc.Role, pc.MediaType, pc.Node, chars)

	n := pc.Node.GetNode("ForeignKeys")
	fks := getForeignKeys(n)
	if fks != nil {
		ja := json.Array{}
		for _, fk := range fks {
			k := json.Object{}
			k.AddString(fk.Field, fk.Value)
			ja.AddObject(k)
		}
		jo.AddArray("foreignkeys", ja)
	}

	rnds.AddRendition("mainweb", jo, true)
}
