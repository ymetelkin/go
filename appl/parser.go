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

type ApplJson struct {
	Xml                  *Publication
	MediaType            MediaType
	Language             *json.JsonProperty
	ReferenceId          *json.JsonProperty
	PubStatus            PubStatus
	FirstCreated         *json.JsonProperty
	FirstCreatedYear     int
	Signals              UniqueStrings
	OutingInstructions   *json.JsonProperty
	EditorialTypes       *json.JsonProperty
	Embargoed            *json.JsonProperty
	TimeRestrictions     map[string]bool
	Associations         *json.JsonProperty
	RefersTo             *json.JsonProperty
	Headline             *json.JsonProperty
	CopyrightNotice      *json.JsonProperty
	KeywordLines         *json.JsonProperty
	Bylines              *json.JsonProperty
	Producer             *json.JsonProperty
	Photographer         *json.JsonProperty
	CaptionWriter        *json.JsonProperty
	Edits                *json.JsonProperty
	OverLines            *json.JsonProperty
	Persons              *json.JsonProperty
	Provider             *json.JsonProperty
	Sources              *json.JsonProperty
	Contributor          *json.JsonProperty
	CanonicalLink        *json.JsonProperty
	SourceMaterials      *json.JsonProperty
	TransmissionSources  *json.JsonProperty
	ProductSources       *json.JsonProperty
	ItemContentType      *json.JsonProperty
	DistributionChannels *json.JsonProperty
	Fixture              *json.JsonProperty
	InPackages           *json.JsonProperty
	Ratings              *json.JsonProperty
	UsageRights          *json.JsonProperty
	Descriptions         *json.JsonProperty
	Generators           []ApplGenerator
	Categories           map[string]string
	SuppCategories       map[string]string
	AlertCategories      UniqueStrings
	Subjects             []ApplSubject
	Organizations        []ApplSubject
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
*/

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
	Geography    UniqueStrings
	RightsHolder string
	Limitations  UniqueStrings
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
	Rels      UniqueStrings
	ParentIds UniqueStrings
	TopParent bool
}

type ApplFiling struct {
	Xml         *FilingMetadata
	ForeignKeys map[string]string
	Products    []int
}

func XmlToJson(s string) (*json.JsonObject, error) {
	pub, err := NewXml(s)
	if err != nil {
		return nil, err
	}

	aj := ApplJson{Xml: pub}

	err = pub.Identification.parse(&aj)
	if err != nil {
		return nil, err
	}

	err = pub.PublicationManagement.parse(&aj)
	if err != nil {
		return nil, err
	}

	err = pub.NewsLines.parse(&aj)
	if err != nil {
		return nil, err
	}

	err = pub.AdministrativeMetadata.parse(&aj)
	if err != nil {
		return nil, err
	}

	err = pub.RightsMetadata.parse(&aj)
	if err != nil {
		return nil, err
	}

	err = pub.DescriptiveMetadata.parse(&aj)
	if err != nil {
		return nil, err
	}

	jo, err := aj.ToJson()
	if err != nil {
		return nil, err
	}

	return jo, nil
}

func (aj *ApplJson) ToJson() (*json.JsonObject, error) {
	jo := json.JsonObject{}
	jo.AddString("representationversion", "1.0")
	jo.AddString("representationtype", "full")

	id := aj.Xml.Identification
	jo.AddString("itemid", id.ItemId)
	jo.AddString("recordid", id.RecordId)
	jo.AddString("compositeid", id.CompositeId)
	jo.AddString("compositiontype", id.CompositionType)
	jo.AddString("type", string(aj.MediaType))

	if id.Priority > 0 {
		jo.AddInt("priority", id.Priority)
	}

	if id.EditorialPriority != "" {
		jo.AddString("editorialpriority", id.EditorialPriority)
	}

	jo.AddProperty(aj.Language)

	if id.RecordSequenceNumber > 0 {
		jo.AddInt("recordsequencenumber", id.RecordSequenceNumber)
	}

	if id.FriendlyKey != "" {
		jo.AddString("friendlykey", id.FriendlyKey)
	}

	jo.AddProperty(aj.ReferenceId)

	pm := aj.Xml.PublicationManagement
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

	jo.AddProperty(aj.FirstCreated)

	if pm.LastModifiedDateTime != "" {
		jo.AddString("lastmodifieddatetime", pm.LastModifiedDateTime+"Z")
	}

	jo.AddString("pubstatus", string(aj.PubStatus))

	if pm.ReleaseDateTime != "" {
		jo.AddString("releasedatetime", pm.ReleaseDateTime+"Z")
	}

	jo.AddProperty(aj.Embargoed)
	jo.AddProperty(aj.EditorialTypes)
	jo.AddProperty(aj.Associations)
	jo.AddProperty(aj.RefersTo)
	jo.AddProperty(aj.OutingInstructions)

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

	jo.AddProperty(aj.Signals.ToJsonProperty("signals"))

	if aj.TimeRestrictions != nil {
		for k, v := range aj.TimeRestrictions {
			jo.AddBoolean(k, v)
		}
	}

	nl := aj.Xml.NewsLines

	if nl.Title != "" {
		jo.AddString("title", nl.Title)
	}

	jo.AddProperty(aj.Headline)

	if nl.ExtendedHeadLine != "" {
		jo.AddString("headline_extended", nl.ExtendedHeadLine)
	}

	if len(nl.BodySubHeader) > 0 {
		jo.AddString("summary", nl.BodySubHeader[0])
	}

	jo.AddProperty(aj.Bylines)
	jo.AddProperty(aj.Producer)
	jo.AddProperty(aj.Photographer)
	jo.AddProperty(aj.CaptionWriter)
	jo.AddProperty(aj.Edits)
	jo.AddProperty(aj.Bylines)
	jo.AddProperty(aj.OverLines)

	if nl.DateLine != "" {
		jo.AddString("dateline", nl.DateLine)
	}

	if nl.CreditLine.Value != "" {
		jo.AddString("creditline", nl.CreditLine.Value)
	}

	if nl.CreditLine.Id != "" {
		jo.AddString("creditlineid", nl.CreditLine.Id)
	}

	jo.AddProperty(aj.CopyrightNotice)

	if nl.RightsLine != "" {
		jo.AddString("rightsline", nl.RightsLine)
	}

	if nl.SeriesLine != "" {
		jo.AddString("seriesline", nl.SeriesLine)
	}

	jo.AddProperty(aj.KeywordLines)

	if nl.OutCue != "" {
		jo.AddString("outcue", nl.OutCue)
	}

	jo.AddProperty(aj.Persons)

	if nl.LocationLine != "" {
		jo.AddString("locationline", nl.LocationLine)
	}

	admin := aj.Xml.AdministrativeMetadata

	jo.AddProperty(aj.Provider)

	if admin.Creator != "" {
		jo.AddString("creator", admin.Creator)
	}

	jo.AddProperty(aj.Sources)
	jo.AddProperty(aj.Contributor)
	jo.AddProperty(aj.CanonicalLink)
	jo.AddProperty(aj.SourceMaterials)

	if admin.WorkflowStatus != "" {
		jo.AddString("workflowstatus", admin.WorkflowStatus)
	}

	jo.AddProperty(aj.TransmissionSources)
	jo.AddProperty(aj.ProductSources)
	jo.AddProperty(aj.ItemContentType)

	if admin.Workgroup != "" {
		jo.AddString("workgroup", admin.Workgroup)
	}

	jo.AddProperty(aj.DistributionChannels)

	if admin.ContentElement != "" {
		jo.AddString("editorialrole", admin.ContentElement)
	}

	jo.AddProperty(aj.Fixture)
	jo.AddProperty(aj.InPackages)
	jo.AddProperty(aj.Ratings)

	rm := aj.Xml.RightsMetadata

	if rm.Copyright.Holder != "" {
		jo.AddString("copyrightholder", rm.Copyright.Holder)
	}

	if rm.Copyright.Date > 0 {
		jo.AddInt("copyrightdate", rm.Copyright.Date)
	}

	jo.AddProperty(aj.UsageRights)

	desc := aj.Xml.DescriptiveMetadata

	jo.AddProperty(aj.Descriptions)

	has_geo := desc.DateLineLocation.LatitudeDD != 0 && desc.DateLineLocation.LongitudeDD != 0
	if desc.DateLineLocation.City != "" || desc.DateLineLocation.Country != "" || desc.DateLineLocation.CountryArea != "" || desc.DateLineLocation.CountryAreaName != "" || desc.DateLineLocation.CountryName != "" || has_geo {
		datelinelocation := json.JsonObject{}
		if desc.DateLineLocation.City != "" {
			datelinelocation.AddString("city", desc.DateLineLocation.City)
		}
		if desc.DateLineLocation.CountryArea != "" {
			datelinelocation.AddString("countryareacode", desc.DateLineLocation.CountryArea)
		}
		if desc.DateLineLocation.CountryAreaName != "" {
			datelinelocation.AddString("countryareaname", desc.DateLineLocation.CountryAreaName)
		}
		if desc.DateLineLocation.Country != "" {
			datelinelocation.AddString("countrycode", desc.DateLineLocation.Country)
		}
		if desc.DateLineLocation.CountryName != "" {
			datelinelocation.AddString("countryname", desc.DateLineLocation.CountryName)
		}
		if has_geo {
			coordinates := json.JsonArray{}
			coordinates.AddFloat(desc.DateLineLocation.LongitudeDD)
			coordinates.AddFloat(desc.DateLineLocation.LatitudeDD)
			geo := json.JsonObject{}
			geo.AddString("type", "Point")
			geo.AddArray("coordinates", &coordinates)

			datelinelocation.AddObject("geometry_geojson", &geo)
		}
	}

	if aj.Generators != nil && len(aj.Generators) > 0 {
		generators := json.JsonArray{}
		for _, g := range aj.Generators {
			generator := json.JsonObject{}
			generator.AddString("name", g.Name)
			generator.AddString("version", g.Version)
			generators.AddObject(&generator)
		}
		jo.AddArray("generators", &generators)
	}

	if aj.Categories != nil && len(aj.Categories) > 0 {
		categories, ok := codeNamesToJsonArray(aj.Categories)
		if ok {
			jo.AddArray("categories", categories)
		}
	}

	if aj.SuppCategories != nil && len(aj.SuppCategories) > 0 {
		suppcategories, ok := codeNamesToJsonArray(aj.SuppCategories)
		if ok {
			jo.AddArray("suppcategories", suppcategories)
		}
	}

	if !aj.AlertCategories.IsEmpty() {
		alertcategories := aj.AlertCategories.ToJson()
		jo.AddArray("alertcategories", alertcategories)
	}

	addSubjects(&aj.Subjects, &jo, "subject")
	addSubjects(&aj.Organizations, &jo, "organizations")

	return &jo, nil
}

func addSubjects(items *[]ApplSubject, jo *json.JsonObject, field string) {
	values := *items
	if values != nil && len(values) > 0 {
		subjects := json.JsonArray{}
		for _, sbj := range values {
			subject := json.JsonObject{}
			subject.AddString("name", sbj.Name)
			subject.AddString("scheme", "http://cv.ap.org/id/")
			subject.AddString("code", sbj.Code)
			if sbj.Creator != "" {
				subject.AddString("creator", sbj.Creator)
			}
			if !sbj.Rels.IsEmpty() {
				subject.AddArray("rels", sbj.Rels.ToJson())
			}
			if !sbj.ParentIds.IsEmpty() {
				subject.AddArray("parentids", sbj.ParentIds.ToJson())
			}
			subject.AddBoolean("topparent", sbj.TopParent)
			subjects.AddObject(&subject)
		}
		jo.AddArray(field, &subjects)
	}
}
