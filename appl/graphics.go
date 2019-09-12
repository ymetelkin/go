package appl

import (
	"fmt"
	"strings"
)

func (rnds *renditions) ParseGraphicComponent(pc pubcomponent, mt mediaType) {
	key := strings.ToLower(pc.Role)
	chars := pc.Node.Node("Characteristics")

	switch key {
	case "main":
		if mt == mediaTypeComplexData {
			var title string
			nd := pc.Node.Node("Presentations")
			if nd.Nodes != nil {
				nd = nd.Node("Presentation")
				if nd.Nodes != nil {
					nd = nd.Node("Characteristics")
					title = nd.Attribute("FrameLocation")
				}
			}

			jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
			rnds.AddRendition(key, jo, false)
		} else {
			n, k := getBinaryName(chars, "", false)
			title := fmt.Sprintf("Full Resolution (%s)", n)
			jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
			rnds.AddRendition("main"+k, jo, true)

			nd := pc.Node.Node("RelatedBinaries")
			a := nd.Attribute("Name")
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
