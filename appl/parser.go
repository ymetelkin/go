package appl

import (
	"github.com/ymetelkin/go/json"
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
	Xml                  *Publication
	MediaType            MediaType
	Language             *json.Property
	ReferenceId          *json.Property
	PubStatus            PubStatus
	FirstCreated         *json.Property
	FirstCreatedYear     int
	Signals              uniqueArray
	OutingInstructions   *json.Property
	EditorialTypes       *json.Property
	Embargoed            *json.Property
	TimeRestrictions     map[string]bool
	Associations         *json.Property
	RefersTo             *json.Property
	Headline             *json.Property
	CopyrightNotice      *json.Property
	KeywordLines         *json.Property
	Bylines              *json.Property
	Producer             *json.Property
	Photographer         *json.Property
	CaptionWriter        *json.Property
	Edits                *json.Property
	OverLines            *json.Property
	Person               *json.Property
	Provider             *json.Property
	Sources              *json.Property
	Contributor          *json.Property
	CanonicalLink        *json.Property
	SourceMaterials      *json.Property
	TransmissionSources  *json.Property
	ProductSources       *json.Property
	ItemContentType      *json.Property
	DistributionChannels *json.Property
	Fixture              *json.Property
	InPackages           *json.Property
	Ratings              *json.Property
	UsageRights          *json.Property
	Descriptions         *json.Property
	DatelineLocation     *json.Property
	Generators           *json.Property
	Categories           *json.Property
	SuppCategories       *json.Property
	AlertCategories      *json.Property
	Subjects             *json.Property
	Persons              *json.Property
	Organizations        *json.Property
	Filings              []ApplFiling
}

/*
type ApplAssociation struct {
	Type     string
	ItemId   string
	Rank     int
	TypeRank int
}

type ApplByline struct {
	Name       string
	Code       string
	Title      string
	Parametric string
}


type ApplPerson struct {
	Name       string
	IsFeatured bool
}

type ApplProvider struct {
	Name    string
	Code    string
	Type    string
	Subtype string
	IsEmpty bool
}

type ApplSource struct {
	Name    string
	Code    string
	Type    string
	Subtype string
	City    string
	Country string
	Url     string
	IsEmpty bool
}

type ApplSourceMaterial struct {
	Name              string
	Code              string
	Type              string
	PermissionGranted string
	IsEmpty           bool
}


type ApplItemContentType struct {
	Name    string
	Code    string
	Creator string
	IsEmpty bool
}

type ApplFixture struct {
	Name    string
	Code    string
	IsEmpty bool
}


type ApplRating struct {
	Value     int
	ScaleMin  int
	ScaleMax  int
	ScaleUnit string
	Raters    int
	RaterType string
}

type ApplUsageRights struct {
	UsageType    string
	Geography    uniqueArray
	RightsHolder string
	Limitations  uniqueArray
	StartDate    string
	EndDate      string
	Groups       []ApplGroup
}

type ApplGroup struct {
	Name string
	Code string
	Type string
}

type ApplGenerator struct {
	Name    string
	Version string
}

type ApplSubject struct {
	Name      string
	Code      string
	Creator   string
	Rels      uniqueArray
	ParentIds uniqueArray
	TopParent bool
}
*/

type ApplFiling struct {
	Xml         *FilingMetadata
	ForeignKeys map[string]string
	Products    []int
}

func XmlToJson(s string) (*json.Object, error) {
	pub, err := NewXml(s)
	if err != nil {
		return nil, err
	}

	doc := document{Xml: pub}

	err = pub.Identification.parse(&doc)
	if err != nil {
		return nil, err
	}

	err = pub.PublicationManagement.parse(&doc)
	if err != nil {
		return nil, err
	}

	err = pub.NewsLines.parse(&doc)
	if err != nil {
		return nil, err
	}

	err = pub.AdministrativeMetadata.parse(&doc)
	if err != nil {
		return nil, err
	}

	err = pub.RightsMetadata.parse(&doc)
	if err != nil {
		return nil, err
	}

	err = pub.DescriptiveMetadata.parse(&doc)
	if err != nil {
		return nil, err
	}

	jo, err := doc.ToJson()
	if err != nil {
		return nil, err
	}

	return jo, nil
}

func (doc *document) ToJson() (*json.Object, error) {
	jo := json.Object{}
	jo.AddString("representationversion", "1.0")
	jo.AddString("representationtype", "full")

	id := doc.Xml.Identification
	jo.AddString("itemid", id.ItemId)
	jo.AddString("recordid", id.RecordId)
	jo.AddString("compositeid", id.CompositeId)
	jo.AddString("compositiontype", id.CompositionType)
	jo.AddString("type", string(doc.MediaType))

	if id.Priority > 0 {
		jo.AddInt("priority", id.Priority)
	}

	if id.EditorialPriority != "" {
		jo.AddString("editorialpriority", id.EditorialPriority)
	}

	jo.AddProperty(doc.Language)

	if id.RecordSequenceNumber > 0 {
		jo.AddInt("recordsequencenumber", id.RecordSequenceNumber)
	}

	if id.FriendlyKey != "" {
		jo.AddString("friendlykey", id.FriendlyKey)
	}

	jo.AddProperty(doc.ReferenceId)

	pm := doc.Xml.PublicationManagement
	jo.AddString("recordtype", pm.RecordType)
	jo.AddString("filingtype", pm.FilingType)

	if pm.ChangeEvent != "" {
		jo.AddString("changeevent", pm.ChangeEvent)
	}

	if pm.ItemKey != "" {
		jo.AddString("itemkey", pm.ItemKey)
	}

	if pm.ArrivalDateTime != "" {
		jo.AddString("arrivaldatetime", pm.ArrivalDateTime+"Z")
	}

	jo.AddProperty(doc.FirstCreated)

	if pm.LastModifiedDateTime != "" {
		jo.AddString("lastmodifieddatetime", pm.LastModifiedDateTime+"Z")
	}

	jo.AddString("pubstatus", string(doc.PubStatus))

	if pm.ReleaseDateTime != "" {
		jo.AddString("releasedatetime", pm.ReleaseDateTime+"Z")
	}

	jo.AddProperty(doc.Embargoed)
	jo.AddProperty(doc.EditorialTypes)
	jo.AddProperty(doc.Associations)
	jo.AddProperty(doc.RefersTo)
	jo.AddProperty(doc.OutingInstructions)

	if pm.SpecialInstructions != "" {
		jo.AddString("specialinstructions", pm.SpecialInstructions)
	}

	if pm.ItemStartDateTime != "" {
		jo.AddString("editorialid", pm.ItemStartDateTime)
	}

	if pm.LastModifiedDateTime != "" {
		jo.AddString("itemstartdatetime", pm.LastModifiedDateTime+"Z")
	}

	if pm.ItemStartDateTimeActual != "" {
		jo.AddString("itemstartdatetimeactual", pm.ItemStartDateTimeActual+"Z")
	}

	if pm.ItemExpireDateTime != "" {
		jo.AddString("itemexpiredatetime", pm.ItemExpireDateTime+"Z")
	}

	if pm.SearchDateTime != "" {
		jo.AddString("searchdatetime", pm.SearchDateTime+"Z")
	}

	if pm.ItemEndDateTime != "" {
		jo.AddString("itemenddatetime", pm.ItemEndDateTime+"Z")
	}

	if pm.Function != "" {
		jo.AddString("function", pm.Function)
	}

	jo.AddProperty(doc.Signals.ToJsonProperty("signals"))

	if doc.TimeRestrictions != nil {
		for k, v := range doc.TimeRestrictions {
			jo.AddBool(k, v)
		}
	}

	nl := doc.Xml.NewsLines

	if nl.Title != "" {
		jo.AddString("title", nl.Title)
	}

	jo.AddProperty(doc.Headline)

	if nl.ExtendedHeadLine != "" {
		jo.AddString("headline_extended", nl.ExtendedHeadLine)
	}

	if len(nl.BodySubHeader) > 0 {
		jo.AddString("summary", nl.BodySubHeader[0])
	}

	jo.AddProperty(doc.Bylines)
	jo.AddProperty(doc.Producer)
	jo.AddProperty(doc.Photographer)
	jo.AddProperty(doc.CaptionWriter)
	jo.AddProperty(doc.Edits)
	jo.AddProperty(doc.Bylines)
	jo.AddProperty(doc.OverLines)

	if nl.DateLine != "" {
		jo.AddString("dateline", nl.DateLine)
	}

	if nl.CreditLine.Value != "" {
		jo.AddString("creditline", nl.CreditLine.Value)
	}

	if nl.CreditLine.Id != "" {
		jo.AddString("creditlineid", nl.CreditLine.Id)
	}

	jo.AddProperty(doc.CopyrightNotice)

	if nl.RightsLine != "" {
		jo.AddString("rightsline", nl.RightsLine)
	}

	if nl.SeriesLine != "" {
		jo.AddString("seriesline", nl.SeriesLine)
	}

	jo.AddProperty(doc.KeywordLines)

	if nl.OutCue != "" {
		jo.AddString("outcue", nl.OutCue)
	}

	jo.AddProperty(doc.Persons)

	if nl.LocationLine != "" {
		jo.AddString("locationline", nl.LocationLine)
	}

	admin := doc.Xml.AdministrativeMetadata

	jo.AddProperty(doc.Provider)

	if admin.Creator != "" {
		jo.AddString("creator", admin.Creator)
	}

	jo.AddProperty(doc.Sources)
	jo.AddProperty(doc.Contributor)
	jo.AddProperty(doc.CanonicalLink)
	jo.AddProperty(doc.SourceMaterials)

	if admin.WorkflowStatus != "" {
		jo.AddString("workflowstatus", admin.WorkflowStatus)
	}

	jo.AddProperty(doc.TransmissionSources)
	jo.AddProperty(doc.ProductSources)
	jo.AddProperty(doc.ItemContentType)

	if admin.Workgroup != "" {
		jo.AddString("workgroup", admin.Workgroup)
	}

	jo.AddProperty(doc.DistributionChannels)

	if admin.ContentElement != "" {
		jo.AddString("editorialrole", admin.ContentElement)
	}

	jo.AddProperty(doc.Fixture)
	jo.AddProperty(doc.InPackages)
	jo.AddProperty(doc.Ratings)

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
	jo.AddProperty(doc.Organizations)

	return &jo, nil
}
