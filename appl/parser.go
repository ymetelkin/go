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
	Filings                []xml.Node
	ItemID                 string
	MediaType              MediaType
	CompositionType        string
	FriendlyKey            string
	EditorialID            string
	Function               string
	Title                  string
	Headline               string
	CopyrightNotice        string
	SlugLine               string
	ForeignKey             string
	PubStatus              PubStatus
	FirstCreatedYear       int
	Signals                uniqueArray
}

func XmlToJson(s string) (*json.Object, error) {
	doc, err := parseXml(s)
	if err != nil {
		return nil, err
	}

	jo := json.Object{}
	jo.AddString("representationversion", "1.0")
	jo.AddString("representationtype", "full")

	err = doc.ParseIdentification(&jo)
	if err != nil {
		return nil, err
	}

	err = doc.ParsePublicationManagement(&jo)
	if err != nil {
		return nil, err
	}

	err = doc.ParseNewsLines(&jo)
	if err != nil {
		return nil, err
	}

	err = doc.ParseAdministrativeMetadata(&jo)
	if err != nil {
		return nil, err
	}

	err = doc.ParseRightsMetadata(&jo)
	if err != nil {
		return nil, err
	}

	doc.SetReferenceId(&jo)
	doc.SetHeadline(&jo)

	return jo, nil
}

func parseXml(s string) (document, error) {
	root, err := xml.New(s)
	if err != nil {
		return nil, err
	}

	doc := document{}
	filings := []xml.Node{}

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
			filings = append(filings, nd)
		}
	}

	doc.Filings = filings

	return doc, nil
}

func (doc *document) ToJson() (*json.Object, error) {

	/*


		rm := doc.Xml.RightsMetadata

		if rm.Copyright.Holder != "" {
			jo.AddString("copyrightholder", rm.Copyright.Holder)
		}

		if rm.Copyright.Date > 0 {
			jo.AddInt("copyrightdate", rm.Copyright.Date)
		}

		jo.AddProperty(doc.UsageRights)

		//desc := doc.Xml.DescriptiveMetadata

		jo.AddProperty(doc.Descriptions)
		jo.AddProperty(doc.DatelineLocation)
		jo.AddProperty(doc.Generators)
		jo.AddProperty(doc.Categories)
		jo.AddProperty(doc.SuppCategories)
		jo.AddProperty(doc.AlertCategories)
		jo.AddProperty(doc.Subjects)
		jo.AddProperty(doc.Persons)
		jo.AddProperty(doc.Organizations)
		jo.AddProperty(doc.Companies)
		jo.AddProperty(doc.Places)
		jo.AddProperty(doc.Events)
		jo.AddProperty(doc.Audiences)
		jo.AddProperty(doc.Services)
		jo.AddProperty(doc.ThirdPartyMeta)

		jo.AddProperty(doc.Filings)

		if doc.Texts != nil {
			for k, v := range doc.Texts {
				jo.AddObject(k, v)
			}
		}
	*/

	return nil, nil
}
