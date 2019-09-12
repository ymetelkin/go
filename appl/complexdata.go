package appl

import (
	"fmt"
	"strings"
)

func (rnds *renditions) ParseComplexDataComponent(pc pubcomponent) {
	key := strings.ToLower(pc.Role)
	chars := pc.Node.Node("Characteristics")

	switch key {
	case "main":
		n, k := getBinaryName(chars, "", false)
		title := fmt.Sprintf("Full Resolution (%s)", n)
		jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
		rnds.AddRendition(k, jo, true)
	case "preview", "thumbnail":
		jo := rnds.GetRendition(pc.Role+" (JPG)", pc.Role, pc.MediaType, pc.Node, chars)
		rnds.AddRendition(key, jo, false)
	}
}
