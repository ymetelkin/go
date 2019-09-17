package appl

import (
	"io"
	"time"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

//Document stuct hold APPL XML, JSON and a few mportant properties
type Document struct {
	ItemID                  string
	RecordID                string
	CompositeID             string
	CompositionType         string
	MediaType               string
	Priority                int
	EditorialPriority       string
	Language                string
	RSN                     int
	FriendlyKey             string
	ReferenceID             string
	RecordType              string
	FilingType              string
	ChangeEvent             string
	ItemKey                 string
	ArrivalDateTime         *time.Time
	Created                 *FirstCreated
	Modified                *LastModified
	Status                  string
	ReleaseDateTime         *time.Time
	Associations            []Association
	Outs                    []string
	SpecialInstructions     string
	EditorialTypes          []string
	ItemStartDateTime       *time.Time
	ItemStartDateTimeActual *time.Time
	ItemExpireDateTime      *time.Time
	SearchDateTime          *time.Time
	ItemEndDateTime         *time.Time
	Embargoed               *time.Time
	Signals                 []string
	Function                string
	TimeRestrictions        []TimeRestriction
	RefersTo                string

	EditorialID string
	Title       string
	Filings     []Filing

	XML  *xml.Node
	JSON *json.Object
}

//UserAccount struct
type UserAccount struct {
	Name        string
	Account     string
	System      string
	ToolVersion string
	Location    string
	Workgroup   string
}

//FirstCreated struct
type FirstCreated struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
	Date   *time.Time
	User   *UserAccount
}

//LastModified struct
type LastModified struct {
	Date *time.Time
	User *UserAccount
}

//Association AssociatedWith stuct
type Association struct {
	ItemID   string
	Type     string
	Rank     int
	TypeRank int
}

//TimeRestriction struct
type TimeRestriction struct {
	ID      string
	System  string
	Zone    string
	Include bool
}

//New create new Document from byte stream
func New(scanner io.ByteScanner) (doc Document, err error) {
	xml, err := xml.New(scanner)
	if err != nil {
		return
	}

	doc.XML = &xml

	doc.JSON = new(json.Object)

	err = doc.parse()
	if err != nil {
		return
	}

	return
}

func (doc *Document) parse() (err error) {
	doc.JSON.AddString("representationversion", "1.0")
	doc.JSON.AddString("representationtype", "full")

	for _, nd := range doc.XML.Nodes {
		switch nd.Name {
		case "Identification":
			doc.parseIdentification(nd)
		case "PublicationManagement":
			//doc.PublicationManagement = &nd
		case "NewsLines":
			//doc.NewsLines = &nd
		case "AdministrativeMetadata":
			//doc.AdministrativeMetadata = &nd
		case "RightsMetadata":
			//doc.RightsMetadata = &nd
		case "DescriptiveMetadata":
			//doc.DescriptiveMetadata = &nd
		case "FilingMetadata":
		case "PublicationComponent":
		}
	}

	doc.SetReferenceID()

	return
}
