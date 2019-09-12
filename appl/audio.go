package appl

import (
	"fmt"
	"strings"
)

func (rnds *renditions) ParseAudioComponent(pc pubcomponent) {
	if !strings.EqualFold(pc.Role, "Main") {
		return
	}
	chars := pc.Node.Node("Characteristics")
	if chars.Nodes == nil {
		return
	}

	n, k := getBinaryName(chars, "", false)
	name := "main" + k
	title := fmt.Sprintf("Full Resolution (%s)", n)
	jo := rnds.GetRendition(title, pc.Role, pc.MediaType, pc.Node, chars)
	rnds.AddRendition(name, jo, true)
}
