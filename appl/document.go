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
	RefersTo                string
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
	Title                   string
	Headline                string
	ExtendedHeadline        string
	Summary                 string
	Overlines               []string
	Keywordlines            []string
	Bylines                 []Byline
	Producer                *CodeNameTitle
	Photographer            *CodeNameTitle
	Captionwriter           *CodeNameTitle
	Edits                   []string
	Dateline                string
	Creditline              string
	Copyright               *Copyright
	Rightsline              string
	Seriesline              string
	OutCue                  string
	Persons                 []Person
	Locationline            string

	EditorialID string
	Filings     []Filing

	Story    *Text
	Caption  *Text
	Script   *Text
	Shotlist *Text

	XML  *xml.Node
	JSON *json.Object
}

//UserAccount struct
type UserAccount struct {
	Name        string
	Account     string
	System      string
	ToolVersion string
	Workgroup   string
	Location    string
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

//Byline struct
type Byline struct {
	Code       string
	By         string
	Title      string
	Parametric string
}

//CodeNameTitle struct
type CodeNameTitle struct {
	Code  string
	Name  string
	Title string
}

//Copyright struct
type Copyright struct {
	Notice string
	Holder string
	Year   int
}

//Person struct
type Person struct {
	Name       string
	IsFeatured bool
}

//Text struct
type Text struct {
	Body  string
	Words int
}

//New create new Document from byte stream
func New(scanner io.ByteScanner) (doc Document, err error) {
	xml, err := xml.Parse(scanner)
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
			doc.parsePublicationManagement(nd)
		case "NewsLines":
			doc.parseNewsLines(nd)
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

	doc.setReferenceID()
	doc.setHeadline()

	return
}
