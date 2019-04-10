package appl

import (
	"strings"

	"github.com/ymetelkin/go/json"
)

type photos struct {
	Renditions map[string]json.Object
	Counts     map[string]int
}

func (phts *photos) ParsePhotoComponent(pc pubcomponent) error {
	key := strings.ToLower(pc.Role)
	chars := pc.Node.GetNode("Characteristics")

	if phts.Renditions == nil {
		phts.Renditions = make(map[string]json.Object)
	}

	switch key {
	case "preview":
		jo := getRenditionMeta(pc.Role+" (JPG)", pc.Role, MEDIATYPE_PHOTO, pc.Node, chars)
		_, ok := phts.Renditions[key]
		if !ok {
			phts.Renditions[key] = jo
		}
	}

	return nil
}

func (phts *photos) AddRenditions(ja *json.Array) {
	if phts.Renditions != nil {
		for _, jo := range phts.Renditions {
			ja.AddObject(jo)
		}
	}
}
