package appl

import "strings"

func (pub *Publication) parsePubComponents(doc *document) error {
	if pub.PublicationComponent == nil {
		return nil
	}

	for _, pc := range pub.PublicationComponent {
		if pc.Role != "" && pc.MediaType != "" {
			role := strings.ToLower(pc.Role)
			media := strings.ToLower(pc.MediaType)

			if media == "text" && pc.TextContentItem.Body.Xml != "" {
				pc.TextContentItem.parse(role, doc)
			}
		}
	}

	return nil
}
