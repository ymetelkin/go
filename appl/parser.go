package appl

import (
	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type MediaType string

const (
	MEDIATYPE_TEXT          MediaType = "text"
	MEDIATYPE_PHOTO         MediaType = "photo"
	MEDIATYPE_VIDEO         MediaType = "video"
	MEDIATYPE_AUDIO         MediaType = "audio"
	MEDIATYPE_GRAPHIC       MediaType = "graphic"
	MEDIATYPE_COMPLEXT_DATA MediaType = "complexdata"
	MEDIATYPE_UNKNOWN       MediaType = ""
)

type PubStatus string

const (
	PUBSTATUS_USABLE   PubStatus = "usable"
	PUBSTATUS_WITHHELD PubStatus = "withheld"
	PUBSTATUS_CANCELED PubStatus = "canceled"
	PUBSTATUS_UNKNOWN  PubStatus = ""
)

type document struct {
	Identification         xml.Node
	PublicationManagement  xml.Node
	NewsLines              xml.Node
	AdministrativeMetadata xml.Node
	RightsMetadata         xml.Node
	DescriptiveMetadata    xml.Node
	Filings                []filing
	PublicationComponents  []pubcomponent
	ItemID                 string
	MediaType              MediaType
	CompositionType        string
	FriendlyKey            string
	EditorialID            string
	Function               string
	Title                  string
	Headline               string
	PubStatus              PubStatus
	FirstCreatedYear       int
	Signals                uniqueArray
	Fixture                bool
}

func XmlToJson(s string) (json.Object, error) {
	doc, err := parseXml(s)
	if err != nil {
		return json.Object{}, err
	}

	jo := json.Object{}
	jo.AddString("representationversion", "1.0")
	jo.AddString("representationtype", "full")

	err = doc.ParseIdentification(&jo)
	if err != nil {
		return json.Object{}, err
	}

	err = doc.ParsePublicationManagement(&jo)
	if err != nil {
		return json.Object{}, err
	}

	err = doc.ParseNewsLines(&jo)
	if err != nil {
		return json.Object{}, err
	}

	err = doc.ParseAdministrativeMetadata(&jo)
	if err != nil {
		return json.Object{}, err
	}

	err = doc.ParseRightsMetadata(&jo)
	if err != nil {
		return json.Object{}, err
	}

	if doc.Filings != nil {
		filings := json.Array{}
		for _, f := range doc.Filings {
			filings.AddObject(f.JSON)
		}
		jo.AddArray("filings", filings)
	}

	doc.ParsePublicationComponents(&jo)

	doc.SetReferenceId(&jo)
	doc.SetHeadline(&jo)

	return jo, nil
}

func parseXml(s string) (document, error) {
	root, err := xml.New(s)
	if err != nil {
		return document{}, err
	}

	var (
		fs  []filing
		pcs []pubcomponent
		doc document
	)

	for _, nd := range root.Nodes {
		switch nd.Name {
		case "Identification":
			doc.Identification = nd
		case "PublicationManagement":
			doc.PublicationManagement = nd
		case "NewsLines":
			doc.NewsLines = nd
		case "AdministrativeMetadata":
			doc.AdministrativeMetadata = nd
		case "RightsMetadata":
			doc.RightsMetadata = nd
		case "DescriptiveMetadata":
			doc.DescriptiveMetadata = nd
		case "FilingMetadata":
			f := parseFiling(nd)
			if fs == nil {
				fs = []filing{f}
			} else {
				fs = append(fs, f)
			}
		case "PublicationComponent":
			pc := parsePublicationComponent(nd)
			if pcs == nil {
				pcs = []pubcomponent{pc}
			} else {
				pcs = append(pcs, pc)
			}
		}
	}

	doc.Filings = fs
	doc.PublicationComponents = pcs

	return doc, nil
}
