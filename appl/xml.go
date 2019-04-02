package appl

import "encoding/xml"

type Publication struct {
	Identification         Identification
	PublicationManagement  PublicationManagement
	NewsLines              NewsLines
	AdministrativeMetadata AdministrativeMetadata
	RightsMetadata         RightsMetadata
	DescriptiveMetadata    DescriptiveMetadata
	FilingMetadata         []FilingMetadata
}

type Identification struct {
	ItemId               string
	RecordId             string
	CompositeId          string
	CompositionType      string
	MediaType            string
	Priority             int
	EditorialPriority    string
	DefaultLanguage      string
	RecordSequenceNumber int
	FriendlyKey          string
}

type PublicationManagement struct {
	RecordType              string
	FilingType              string
	ChangeEvent             string
	ItemKey                 string
	ArrivalDateTime         string
	FirstCreated            FirstCreated
	LastModifiedDateTime    string
	Status                  string
	ReleaseDateTime         string
	AssociatedWith          []AssociatedWith
	RefersTo                []string
	Instruction             []string
	SpecialInstructions     string
	Editorial               []Editorial
	EditorialId             string
	ItemStartDateTime       string
	ItemStartDateTimeActual string
	SearchDateTime          string
	ItemEndDateTime         string
	ItemExpireDateTime      string
	ExplicitWarning         string
	Function                string
	IsDigitized             string
	TimeRestrictions        TimeRestrictions
}

type NewsLines struct {
	Title            string
	HeadLine         string
	OriginalHeadLine string
	BodySubHeader    []string
	ExtendedHeadLine string
	ByLine           []ByLine
	ByLineOriginal   []ByLineOriginal
	OverLine         []string
	DateLine         string
	CreditLine       CreditLine
	CopyrightLine    string
	RightsLine       string
	SeriesLine       string
	KeywordLine      []string
	OutCue           string
	NameLine         []NameLine
	LocationLine     string
}

type AdministrativeMetadata struct {
	Provider            Provider
	Creator             string
	Source              []Source
	Contributor         []string
	SourceMaterial      []SourceMaterial
	WorkflowStatus      string
	TransmissionSource  []string
	ProductSource       []string
	ItemContentType     ItemContentType
	Workgroup           string
	DistributionChannel []string
	ContentElement      string
	Fixture             Fixture
	Reach               []string
	ConsumerReady       string
	Signal              []string
	InPackage           []string
	Rating              []Rating
}

type RightsMetadata struct {
	Copyright   Copyright
	UsageRights []UsageRights
}

type DescriptiveMetadata struct {
	Description            []string
	DateLineLocation       DateLineLocation
	SubjectClassification  []Classification `xml:"SubjectClassification"`
	EntityClassification   []Classification
	AudienceClassification []Classification
	SalesClassification    []Classification
	Comment                []Classification
	ThirdPartyMeta         []Classification
}

type FilingMetadata struct {
	SlugLine    string
	ForeignKeys ForeignKeys
}

type FirstCreated struct {
	Year  int    `xml:"Year,attr"`
	Month int    `xml:"Month,attr"`
	Day   int    `xml:"Day,attr"`
	Time  string `xml:"Time,attr"`
}

type TimeRestrictions struct {
	TimeRestriction []TimeRestriction
}

type TimeRestriction struct {
	System  string `xml:"System,attr"`
	Zone    string `xml:"Zone,attr"`
	Include bool   `xml:"Include,attr"`
}

type Editorial struct {
	Type string
}

type AssociatedWith struct {
	Value           string `xml:",chardata"`
	CompositionType string `xml:"CompositionType,attr"`
}

type CreditLine struct {
	Id    string `xml:"Id,attr"`
	Value string `xml:",chardata"`
}

type Copyright struct {
	Holder string `xml:"Holder,attr"`
	Date   int    `xml:"Date,attr"`
}

type ByLine struct {
	Value      string `xml:",chardata"`
	Id         string `xml:"Id,attr"`
	Title      string `xml:"Title,attr"`
	Parametric string `xml:"Parametric,attr"`
}

type ByLineOriginal struct {
	Value string `xml:",chardata"`
	Title string `xml:"Title,attr"`
}

type NameLine struct {
	Value      string `xml:",chardata"`
	Parametric string `xml:"Parametric,attr"`
}

type Provider struct {
	Value   string `xml:",chardata"`
	Id      string `xml:"Id,attr"`
	Type    string `xml:"Type,attr"`
	SubType string `xml:"SubType,attr"`
}

type Source struct {
	Value   string `xml:",chardata"`
	City    string `xml:"City,attr"`
	Country string `xml:"Country,attr"`
	Id      string `xml:"Id,attr"`
	Url     string `xml:"Url,attr"`
	Type    string `xml:"Type,attr"`
	SubType string `xml:"SubType,attr"`
}

type SourceMaterial struct {
	Id                string `xml:"Id,attr"`
	Name              string `xml:"Name,attr"`
	Type              string
	Url               string
	PermissionGranted string
}

type ItemContentType struct {
	Value  string `xml:",chardata"`
	Id     string `xml:"Id,attr"`
	System string `xml:"System,attr"`
}

type Fixture struct {
	Value string `xml:",chardata"`
	Id    string `xml:"Id,attr"`
}

type Rating struct {
	Value     int    `xml:"Value,attr"`
	ScaleMin  int    `xml:"ScaleMin,attr"`
	ScaleMax  int    `xml:"ScaleMax,attr"`
	ScaleUnit string `xml:"ScaleUnit,attr"`
	Raters    int    `xml:"Raters,attr"`
	RaterType string `xml:"RaterType,attr"`
}

type UsageRights struct {
	UsageType    string
	Geography    []string
	RightsHolder string
	Limitations  []string
	StartDate    string
	EndDate      string
	Group        []Group
}

type Group struct {
	Value string `xml:",chardata"`
	Id    string `xml:"Id,attr"`
	Type  string `xml:"Type,attr"`
}

type DateLineLocation struct {
	City            string
	CountryArea     string
	CountryAreaName string
	Country         string
	CountryName     string
	LatitudeDD      float64
	LongitudeDD     float64
}

type Classification struct {
	SystemVersion    string `xml:"SystemVersion,attr"`
	AuthorityVersion string `xml:"AuthorityVersion,attr"`
	System           string `xml:"System,attr"`
	Authority        string `xml:"Authority,attr"`
	Occurrence       []Occurrence
}

type Occurrence struct {
	Id          string `xml:"Id,attr"`
	Value       string `xml:"Value,attr"`
	ActualMatch string `xml:"ActualMatch,attr"`
	ParentId    string `xml:"ParentId,attr"`
	TopParent   bool   `xml:"TopParent,attr"`
	Property    []Property
}

type Property struct {
	Id       string `xml:"Id,attr"`
	Name     string `xml:"Name,attr"`
	Value    string `xml:"Value,attr"`
	ParentId string `xml:"ParentId,attr"`
}

type ForeignKeys struct {
	System string `xml:"System,attr"`
	Keys   []Keys `xml:"Keys"`
}

type Keys struct {
	Id    string `xml:"Id,attr"`
	Field string `xml:"Field,attr"`
}

func NewXml(s string) (*Publication, error) {
	publication := Publication{}
	err := xml.Unmarshal([]byte(s), &publication)
	return &publication, err
}

func (publication *Publication) ToString() (string, error) {
	bytes, err := xml.Marshal(publication)

	s := ""
	if err == nil {
		s = string(bytes)
	}

	return s, err
}
