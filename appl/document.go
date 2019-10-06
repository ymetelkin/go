package appl

import (
	"io"
	"strings"
	"time"

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
	EditorialID             string
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
	Producer                *CodeName
	Photographer            *CodeNameTitle
	Captionwriter           *CodeNameTitle
	Edits                   []string
	Dateline                string
	Creditline              *CodeName
	Copyright               *Copyright
	Rightsline              string
	Seriesline              string
	OutCue                  string
	Locationline            string
	Provider                *Provider
	Creator                 string
	Sources                 []Source
	Contributor             string
	SourceMaterials         []SourceMaterial
	CanonicalLink           string
	WorkflowStatus          string
	TransmissionSources     []string
	ProductSources          []string
	ItemContentType         *ItemContentType
	Workgroup               string
	DistributionChannels    []string
	InPackages              []string
	ConsumerReady           string
	EditorialRole           string
	Fixture                 *CodeName
	Ratings                 []Rating
	UsageRights             []UsageRights
	Descriptions            []string
	DateLineLocation        *Location
	Generators              []CodeName
	Categories              []CodeName
	SuppCategories          []CodeName
	AlertCategories         []string
	Subjects                []Subject
	Persons                 []Person
	Organizations           []Subject
	Companies               []Company
	Places                  []Place
	Events                  []Event
	Audiences               []CodeNameTitle
	Services                []CodeName
	Perceptions             []Perception
	ThirdParties            []ThirdParty
	Filings                 []Filing
	Story                   *Text
	Caption                 *Text
	Script                  *Text
	Shotlist                *Text
	PublishableEditorNotes  *Text
	Renditions              []Rendition
	Parts                   []Rendition
	Shots                   []PhotoShot

	XML *xml.Node
	//JSON *json.Object
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
	Time   string
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

//CodeName struct
type CodeName struct {
	Code string
	Name string
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

//Provider struct
type Provider struct {
	Code    string
	Type    string
	Subtype string
	Name    string
}

//Source struct
type Source struct {
	Code        string
	City        string
	Country     string
	County      string
	CountryArea string
	URL         string
	Type        string
	Subtype     string
	Name        string
}

//SourceMaterial struct
type SourceMaterial struct {
	Code              string
	Type              string
	PermissionGranted string
	Name              string
}

//ItemContentType struct
type ItemContentType struct {
	Code    string
	Creator string
	Name    string
}

//Rating struct
type Rating struct {
	Value     int
	ScaleMin  int
	ScaleMax  int
	ScaleUnit string
	Raters    int
	RaterType string
	Creator   string
}

//UsageRights struct
type UsageRights struct {
	UsageType    string
	Geography    []string
	RightsHolder string
	Limitations  []string
	StartDate    *time.Time
	EndDate      *time.Time
	Groups       []CodeNameTitle
}

//Location struct
type Location struct {
	City            string
	CountryAreaCode string
	CountryAreaName string
	CountryCode     string
	CountryName     string
	Geo             *Geo
}

//Geo struct
type Geo struct {
	Longitude float64
	Latitude  float64
}

//Subject struct
type Subject struct {
	Code      string
	Name      string
	Creator   string
	Rels      []string
	ParentIDs []string
	TopParent *bool
	rels      uniqueStrings
	ids       uniqueStrings
}

//Person struct
type Person struct {
	Code       string
	Name       string
	Creator    string
	IsNameline bool
	IsFeatured bool
	Rels       []string
	Types      []string
	IDs        []string
	Teams      []CodeName
	States     []CodeName
	Events     []CodeName
	rels       uniqueStrings
	types      uniqueStrings
	ids        uniqueStrings
	teams      uniqueCodeNames
	states     uniqueCodeNames
	events     uniqueCodeNames
}

//Company struct
type Company struct {
	Code       string
	Name       string
	Creator    string
	Rels       []string
	Symbols    []string
	Industries []CodeName
	rels       uniqueStrings
	symbols    uniqueStrings
	tickers    uniqueCodeNames
	industries uniqueCodeNames
	exchanges  map[string]string
}

//Place struct
type Place struct {
	Code         string
	Name         string
	Creator      string
	Rels         []string
	ParentIDs    []string
	TopParent    *bool
	LocationType *CodeName
	Geo          *Geo
	rels         uniqueStrings
	ids          uniqueStrings
}

//Event struct
type Event struct {
	Code        string
	Name        string
	Creator     string
	ExternalIDs []CodeName
	Properties  []CodeName
	props       uniqueCodeNames
}

//Perception struct
type Perception struct {
	Code    string
	Name    string
	Creator string
	Rel     string
}

//ThirdParty struct
type ThirdParty struct {
	Code            string
	Name            string
	Creator         string
	Vocabulary      string
	VocabularyOwner string
}

//Filing APPL filing
type Filing struct {
	ID                     string
	ArrivalDateTime        *time.Time
	Cycle                  string
	TransmissionReference  string
	TransmissionFilename   string
	TransmissionContent    string
	ServiceLevelDesignator string
	Selector               string
	Format                 string
	Source                 string
	Category               string
	Routings               map[string][]string
	Slugline               string
	OriginalMediaID        string
	ImportFolder           string
	ImportWarnings         string
	LibraryTwinCheck       string
	LibraryRequestID       string
	SpecialFieldAttn       string
	Feedline               string
	LibraryRequestLogin    string
	Products               []int
	Priorityline           string
	ForeignKeys            []ForeignKey
	Countries              []string
	Regions                []string
	Subjects               []string
	Topics                 []string
	OnlineCode             string
	DistributionScope      string
	BreakingNews           string
	Style                  string
	Junkline               string
}

//ForeignKey APPL foreign key
type ForeignKey struct {
	Field string
	Value string
}

//Text struct
type Text struct {
	Body  string
	Words int
}

//Rendition struct
type Rendition struct {
	Title                     string
	Rel                       string
	Code                      string
	MediaType                 string
	TapeNumber                string
	FileExtension             string
	Scene                     string
	SceneID                   string
	BroadcastFormat           string
	PresentationSystem        string
	PresentationFrame         string
	PresentationFrameLocation string
	ByteSize                  int
	Width                     int
	Height                    int
	FrameRate                 float64
	Resolution                int
	ResolutionUnits           string
	TotalDuration             int
	PhysicalType              string
	ForeignKeys               []ForeignKey
	Attributes                map[string]string
	Characteristics           map[string]string
	//FramesTotal               int
	//AverageBitRate            float64
	//DataRate                  int
	//DurationFrameValue        int
	//Resolution                int
	//ResolutionUnits           string
	//TotalDuration             int
	//InTimeCode                int
	//InTimeFrameValue          int
	//PixelDepth                int
	//Rotation                  int
	//SampleRate                float64
	//SampleSize                int
	//Words                     int
}

//PhotoShot struct
type PhotoShot struct {
	Sequence  int
	Href      string
	Width     int
	Height    int
	StartTime int
	EndTime   int
}

//Parse create new Document from byte stream
func Parse(scanner io.ByteScanner) (doc Document, err error) {
	xml, err := xml.Parse(scanner)
	if err != nil {
		return
	}

	doc.XML = &xml

	err = doc.parse()
	if err != nil {
		return
	}

	return
}

func (doc *Document) parse() (err error) {
	var rp renditionParser

	for _, nd := range doc.XML.Nodes {
		switch nd.Name {
		case "Identification":
			doc.parseIdentification(nd)
		case "PublicationManagement":
			doc.parsePublicationManagement(nd)
		case "NewsLines":
			doc.parseNewsLines(nd)
		case "AdministrativeMetadata":
			doc.parseAdministrativeMetadata(nd)
		case "RightsMetadata":
			doc.parseRightsMetadata(nd)
		case "DescriptiveMetadata":
			doc.parseDescriptiveMetadata(nd)
		case "FilingMetadata":
			doc.parseFilingMetadata(nd)
		case "PublicationComponent":
			doc.parsePublicationComponent(nd, &rp)
		}
	}

	doc.setReferenceID()
	doc.setHeadline()
	doc.parseNewsContent()

	return
}

//ParseString create new Document from byte stream
func ParseString(s string) (Document, error) {
	scanner := strings.NewReader(s)
	return Parse(scanner)
}
