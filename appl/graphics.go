package appl

import "strings"

func (rnds *renditions) ParseGraphicComponent(pc pubcomponent, mt MediaType) {
	key := strings.ToLower(pc.Role)
	chars := pc.Node.GetNode("Characteristics")

	switch key {
	case "main":
		if mt == mediaTypeComplexData {

		} else {

		}
	case "preview", "thumbnail":
		jo := rnds.GetRendition(pc.Role+" (JPG)", pc.Role, pc.MediaType, pc.Node, chars)
		rnds.AddRendition(key, jo, false)
	}
}
