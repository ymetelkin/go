package appl

import (
	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type pubcomponent struct {
	Role      string
	MediaType MediaType
	Node      xml.Node
}

func (doc *document) ParsePublicationComponents(jo *json.Object) {
	var (
		txts texts
		phts photos
	)

	if doc.PublicationComponents != nil {
		for _, pc := range doc.PublicationComponents {
			switch pc.Node.Name {
			case "TextContentItem":
				txts.ParseTextComponent(pc)
			case "PhotoContentItem":
				phts.ParsePhotoComponent(pc)
			case "GraphicContentItem":
			case "VideoContentItem":
			case "AudioContentItem":
			case "ComplexDataContentItem":
			}
		}

		ja := json.Array{}

		txts.AddProperties(jo)
		phts.AddRenditions(&ja)

		if !ja.IsEmpty() {
			jo.AddArray("renditions", ja)
		}
	}
}

func parsePublicationComponent(nd xml.Node) pubcomponent {
	if nd.Nodes == nil || len(nd.Nodes) == 0 || nd.Attributes == nil {
		return pubcomponent{}
	}

	var (
		role string
		mt   MediaType
	)

	for _, a := range nd.Attributes {
		switch a.Name {
		case "Role":
			role = a.Value
		case "MediaType":
			mt, _ = getMediaType(a.Value)
		}
	}

	if role == "" || mt == MEDIATYPE_UNKNOWN {
		return pubcomponent{}
	}

	return pubcomponent{Role: role, MediaType: mt, Node: nd.Nodes[0]}
}
