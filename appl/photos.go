package appl

import (
	"strings"

	"github.com/ymetelkin/go/json"
)

type photos struct {
	renditions map[string]json.Object
	counts     map[string]int
}

func (phts *photos) ParsePhotoComponent(pc pubcomponent) error {
	key := strings.ToLower(pc.Role)
	chars := pc.Node.GetNode("Characteristics")

	if phts.renditions == nil {
		phts.renditions = make(map[string]json.Object)
	}

	switch key {
	case "preview":
		jo := getRenditionMeta(pc.Role+" (JPG)", pc.Role, MEDIATYPE_PHOTO, pc.Node, chars)
		_, ok := phts.renditions[key]
		if !ok {
			phts.renditions[key] = jo
		}
	}

	return nil
}

func (phts *photos) AddRenditions(ja *json.Array) {
	if phts.renditions != nil {
		for _, jo := range phts.renditions {
			ja.AddObject(jo)
		}
	}
}
