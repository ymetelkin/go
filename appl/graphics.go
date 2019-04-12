package appl

import (
	"fmt"
	"strings"
)

func (rnds *renditions) ParseGraphicComponent(pc pubcomponent, mt mediaType) {
	key := strings.ToLower(pc.Role)
	chars := pc.Node.GetNode("Characteristics")

	switch key {
	case "main":
		if mt == mediaTypeComplexData {
			var title string
			nd := pc.Node.GetNode("Presentations")
			if nd.Nodes != nil {
				nd = nd.GetNode("Presentation")
				if nd.Nodes != nil {
					nd = nd.GetNode("Characteristics")
					title = nd.GetAttribute("FrameLocation")
				}
			}

			jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
			rnds.AddRendition(key, jo, false)
		} else {
			n, k := getBinaryName(chars, "", false)
			title := fmt.Sprintf("Full Resolution (%s)", n)
			jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
			rnds.AddRendition("main"+k, jo, true)

			nd := pc.Node.GetNode("RelatedBinaries")
			a := nd.GetAttribute("Name")
			if strings.EqualFold(a, "MatteFileName") {
				title = fmt.Sprintf("Full Resolution Matte (%s)", n)
				jo = rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
				rnds.AddRendition("mainmatte"+k, jo, true)
			}
		}
	case "preview", "thumbnail":
		jo := rnds.GetRendition(pc.Role+" (JPG)", pc.Role, pc.MediaType, pc.Node, chars)
		rnds.AddRendition(key, jo, false)
	}
}
