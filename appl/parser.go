package appl

import (
	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type mediaType string

const (
	mediaTypeText        mediaType = "text"
	mediaTypePhoto       mediaType = "photo"
	mediaTypeVideo       mediaType = "video"
	mediaTypeAudio       mediaType = "audio"
	mediaTypeGraphic     mediaType = "graphic"
	mediaTypeComplexData mediaType = "complexdata"
	mediaTypeUnknown     mediaType = ""
)

type pubStatus string

const (
	pubStatusUsable   pubStatus = "usable"
	pubStatusWithheld pubStatus = "withheld"
	pubStatusCanceled pubStatus = "canceled"
	pubStatusUnknown  pubStatus = ""
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
	MediaType              mediaType
	CompositionType        string
	FriendlyKey            string
	EditorialID            string
	Function               string
	Title                  string
	Headline               string
	PubStatus              pubStatus
	FirstCreatedYear       int
	Signals                uniqueArray
	Namelines              []json.Object
	Fixture                bool
}

//XMLToJSON converts APPL XML to APPL JSON
func XMLToJSON(s string) (json.Object, error) {
	doc, err := parseXML(s)
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

	err = doc.ParseDescriptiveMetadata(&jo)
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

	doc.SetReferenceID(&jo)
	doc.SetHeadline(&jo)

	return jo, nil
}

func parseXML(s string) (document, error) {
	root, err := xml.New(s)
	if err != nil {
		return document{}, err
	}

	//fmt.Println(root.ToString())

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
