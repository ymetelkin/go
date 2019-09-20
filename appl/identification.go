package appl

import (
	"strconv"
	"strings"

	"github.com/ymetelkin/go/xml"
)

func (doc *Document) parseIdentification(node xml.Node) {
	if node.Nodes == nil {
		return
	}

	for _, nd := range node.Nodes {
		switch nd.Name {
		case "ItemId":
			doc.ItemID = nd.Text
		case "RecordId":
			doc.RecordID = nd.Text
		case "CompositeId":
			doc.CompositeID = nd.Text
		case "CompositionType":
			doc.CompositionType = nd.Text
		case "MediaType":
			doc.MediaType = strings.ToLower(nd.Text)
		case "Priority":
			priority, err := strconv.Atoi(nd.Text)
			if err == nil {
				doc.Priority = priority
			}
		case "EditorialPriority":
			doc.EditorialPriority = nd.Text
		case "DefaultLanguage":
			if len(nd.Text) >= 2 {
				language := string([]rune(nd.Text)[0:2])
				doc.Language = language
			}
		case "RecordSequenceNumber":
			rsn, err := strconv.Atoi(nd.Text)
			if err == nil {
				doc.RSN = rsn
			}
		case "FriendlyKey":
			doc.FriendlyKey = nd.Text
		}
	}

	doc.ReferenceID = doc.ItemID
}

func (doc *Document) setReferenceID() {
	var ref string

	if (doc.MediaType == "photo" || doc.MediaType == "graphic") && doc.FriendlyKey != "" {
		ref = doc.FriendlyKey
	} else if doc.MediaType == "audio" && doc.EditorialID != "" {
		ref = doc.EditorialID
	} else if doc.MediaType == "complexdata" && doc.Title != "" {
		ref = doc.Title
	} else if doc.MediaType == "text" {
		if doc.Title != "" {
			ref = doc.Title
		} else if doc.Filings != nil {
			for _, f := range doc.Filings {
				if f.Slugline != "" {
					ref = f.Slugline
					break
				}
			}
		}
	} else if doc.MediaType == "video" {
		if doc.CompositionType == "StandardBroadcastVideo" {
			if doc.EditorialID != "" {
				ref = doc.EditorialID
			}
		} else if doc.Filings != nil {
			var exit bool
			for _, f := range doc.Filings {
				if f.ForeignKeys != nil {
					for _, fk := range f.ForeignKeys {
						if fk.Field == "storyid" && fk.Value != "" {
							ref = fk.Value
							exit = true
							break
						}
					}
				}
				if exit {
					break
				}
			}
		}
	}

	if ref != "" {
		doc.ReferenceID = ref
	}
}
