package appl

import (
	"github.com/ymetelkin/go/xml"
)

type pubcomponent struct {
	Role      string
	MediaType MediaType
	Node      xml.Node
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
