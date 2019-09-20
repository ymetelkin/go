package appl

import (
	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type pubcomponent struct {
	Role      string
	MediaType mediaType
	Node      xml.Node
}

func (doc *document) ParsePublicationComponents(jo *json.Object) {
	var (
		txts texts
		rnds renditions
	)

	if doc.PublicationComponents != nil {
		var (
			duration int64
			phcol    []pubcomponent
		)

		for _, pc := range doc.PublicationComponents {
			switch pc.Node.Name {
			case "TextContentItem":
				txts.ParseTextComponent(pc)
			case "PhotoContentItem":
				rnds.ParsePhotoComponent(pc, doc.MediaType)
			case "PhotoCollectionContentItem":
				if phcol == nil {
					phcol = []pubcomponent{pc}
				} else {
					phcol = append(phcol, pc)
				}
			case "GraphicContentItem":
				rnds.ParseGraphicComponent(pc, doc.MediaType)
			case "VideoContentItem":
				test := rnds.ParseVideoComponent(pc)
				if duration == 0 && test > 0 {
					duration = test
				}
			case "WebPartContentItem":
				rnds.ParseWebPartomponent(pc, doc.MediaType)
			case "AudioContentItem":
				rnds.ParseAudioComponent(pc)
			case "ComplexDataContentItem":
				rnds.ParseComplexDataComponent(pc)
			}
		}

		if phcol != nil {
			for _, pc := range phcol {
				rnds.ParsePhotoCollection(pc, duration)
			}
		}

		txts.AddProperties(jo)
		rnds.AddNonRenditions(jo)
		rnds.AddRenditions(jo)
	}
}

func parsePublicationComponent2(nd xml.Node) pubcomponent {
	if nd.Nodes == nil || len(nd.Nodes) == 0 || nd.Attributes == nil {
		return pubcomponent{}
	}

	var (
		role string
		mt   mediaType
	)

	for k, v := range nd.Attributes {
		switch k {
		case "Role":
			role = v
		case "MediaType":
			//mt, _ = getMediaType(v)
		}
	}

	if role == "" || mt == mediaTypeUnknown {
		return pubcomponent{}
	}

	return pubcomponent{Role: role, MediaType: mt, Node: nd.Nodes[0]}
}
