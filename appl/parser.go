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
	Language             string
	ReferenceId          string
	PubStatus            PubStatus
	FirstCreated         string
	FirstCreatedYear     int
	Signals              UniqueStrings
	OutingInstructions   UniqueStrings
	EditorialTypes       UniqueStrings
	Embargoed            string
	TimeRestrictions     map[string]bool
	Associations         []ApplAssociation
	RefersTo             string
	Headline             string
	CopyrightNotice      string
	KeywordLines         UniqueStrings
	Bylines              []ApplByline
	Producer             ApplByline
	Photographer         ApplByline
	CaptionWriter        ApplByline
	Editor               ApplByline
	Persons              []ApplPerson
	Provider             ApplProvider
	Sources              []ApplSource
	Contributor          string
	CanonicalLink        string
	SourceMaterials      []ApplSourceMaterial
	TransmissionSources  UniqueStrings
	ProductSources       UniqueStrings
	ItemContentType      ApplItemContentType
	DistributionChannels UniqueStrings
	Fixture              ApplFixture
	InPackages           UniqueStrings
	UsageRights          []ApplUsageRights
	Filings              []ApplFiling
}

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

	if aj.Language != "" {
		jo.AddString("language", aj.Language)
	}

	if id.RecordSequenceNumber > 0 {
		jo.AddInt("recordsequencenumber", id.RecordSequenceNumber)
	}

	if id.FriendlyKey != "" {
		jo.AddString("friendlykey", id.FriendlyKey)
	}

	jo.AddString("referenceid", aj.ReferenceId)

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

	jo.AddString("firstcreated", aj.FirstCreated)

	if pm.LastModifiedDateTime != "" {
		jo.AddString("lastmodifieddatetime", pm.LastModifiedDateTime+"Z")
	}

	jo.AddString("pubstatus", string(aj.PubStatus))

	if pm.ReleaseDateTime != "" {
		jo.AddString("releasedatetime", pm.ReleaseDateTime+"Z")
	}

	if !aj.EditorialTypes.IsEmpty() {
		editorialtypes := aj.EditorialTypes.ToJson()
		jo.AddArray("editorialtypes", editorialtypes)
	}

	if aj.Associations != nil && len(aj.Associations) > 0 {
		associations := json.JsonArray{}
		for _, association := range aj.Associations {
			ass := json.JsonObject{}
			if association.Type != "" {
				ass.AddString("type", association.Type)
			}
			ass.AddString("itemid", association.ItemId)
			ass.AddString("representationtype", "partial")
			ass.AddInt("associationrank", association.Rank)
			ass.AddInt("typerank", association.TypeRank)
			associations.AddObject(&ass)
		}
		jo.AddArray("associations", &associations)
	}

	if aj.RefersTo != "" {
		jo.AddString("refersto", aj.RefersTo)
	}

	if !aj.OutingInstructions.IsEmpty() {
		outinginstructions := aj.OutingInstructions.ToJson()
		jo.AddArray("outinginstructions", outinginstructions)
	}

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

	if !aj.Signals.IsEmpty() {
		signals := aj.Signals.ToJson()
		jo.AddArray("signals", signals)
	}

	if aj.TimeRestrictions != nil {
		for k, v := range aj.TimeRestrictions {
			jo.AddBoolean(k, v)
		}
	}

	nl := aj.Xml.NewsLines

	if nl.Title != "" {
		jo.AddString("title", nl.Title)
	}

	if aj.Headline != "" {
		jo.AddString("headline", aj.Headline)
	}

	if nl.ExtendedHeadLine != "" {
		jo.AddString("headline_extended", nl.ExtendedHeadLine)
	}

	if len(nl.BodySubHeader) > 0 {
		jo.AddString("summary", nl.BodySubHeader[0])
	}

	if aj.Bylines != nil && len(aj.Bylines) > 0 {
		bylines := json.JsonArray{}
		for _, bl := range aj.Bylines {
			byline := json.JsonObject{}
			if bl.Code != "" {
				byline.AddString("code", bl.Code)
			}
			byline.AddString("by", bl.Name)
			if bl.Title != "" {
				byline.AddString("title", bl.Title)
			}
			if bl.Parametric != "" {
				byline.AddString("parametric", bl.Parametric)
			}
			bylines.AddObject(&byline)
		}
		jo.AddArray("bylines", &bylines)
	}

	if aj.Producer.Name != "" {
		producer := json.JsonObject{}
		if aj.Producer.Code != "" {
			producer.AddString("code", aj.Producer.Code)
		}
		producer.AddString("name", aj.Producer.Name)
		jo.AddObject("producer", &producer)
	}

	if aj.Photographer.Name != "" {
		photographer := json.JsonObject{}
		if aj.Photographer.Code != "" {
			photographer.AddString("code", aj.Photographer.Code)
		}
		photographer.AddString("name", aj.Photographer.Name)
		if aj.Photographer.Title != "" {
			photographer.AddString("title", aj.Photographer.Title)
		}
		jo.AddObject("photographer", &photographer)
	}

	if aj.CaptionWriter.Name != "" {
		captionwriter := json.JsonObject{}
		if aj.CaptionWriter.Code != "" {
			captionwriter.AddString("code", aj.CaptionWriter.Code)
		}
		captionwriter.AddString("name", aj.CaptionWriter.Name)
		if aj.CaptionWriter.Title != "" {
			captionwriter.AddString("title", aj.CaptionWriter.Title)
		}
		jo.AddObject("captionwriter", &captionwriter)
	}

	if aj.Editor.Name != "" {
		editor := json.JsonObject{}
		if aj.Editor.Code != "" {
			editor.AddString("code", aj.Editor.Code)
		}
		editor.AddString("name", aj.Editor.Name)
		if aj.Editor.Title != "" {
			editor.AddString("title", aj.Editor.Title)
		}
		jo.AddObject("editor", &editor)
	}

	if len(nl.OverLine) > 0 {
		overlines := json.JsonArray{}
		for _, s := range nl.OverLine {
			overlines.AddString(s)
		}
		jo.AddArray("overlines", &overlines)
	}

	if nl.DateLine != "" {
		jo.AddString("dateline", nl.DateLine)
	}

	if nl.CreditLine.Value != "" {
		jo.AddString("creditline", nl.CreditLine.Value)
	}

	if nl.CreditLine.Id != "" {
		jo.AddString("creditlineid", nl.CreditLine.Id)
	}

	if aj.CopyrightNotice != "" {
		jo.AddString("copyrightnotice", aj.CopyrightNotice)
	}

	if nl.RightsLine != "" {
		jo.AddString("rightsline", nl.RightsLine)
	}

	if nl.SeriesLine != "" {
		jo.AddString("seriesline", nl.SeriesLine)
	}

	if !aj.KeywordLines.IsEmpty() {
		keywordlines := aj.KeywordLines.ToJson()
		jo.AddArray("keywordlines", keywordlines)
	}

	if nl.OutCue != "" {
		jo.AddString("outcue", nl.OutCue)
	}

	if aj.Persons != nil && len(aj.Persons) > 0 {
		persons := json.JsonArray{}
		for _, person := range aj.Persons {
			p := json.JsonObject{}
			p.AddString("name", person.Name)
			if person.IsFeatured {
				rel := json.JsonArray{}
				rel.AddString("personfeatured")
				p.AddArray("rel", &rel)
			}
			p.AddString("creator", "Editorial")
			persons.Add(p)
		}
		jo.AddArray("person", &persons) //yes, singular!?
	}

	if nl.LocationLine != "" {
		jo.AddString("locationline", nl.LocationLine)
	}

	admin := aj.Xml.AdministrativeMetadata

	if !aj.Provider.IsEmpty {
		provider := json.JsonObject{}
		if aj.Provider.Code != "" {
			provider.AddString("code", aj.Provider.Code)
		}
		if aj.Provider.Type != "" {
			provider.AddString("type", aj.Provider.Type)
		}
		if aj.Provider.Subtype != "" {
			provider.AddString("subtype", aj.Provider.Subtype)
		}
		if aj.Provider.Name != "" {
			provider.AddString("name", aj.Provider.Name)
		}
		jo.AddObject("provider", &provider)
	}

	if admin.Creator != "" {
		jo.AddString("creator", admin.Creator)
	}

	if aj.Sources != nil && len(aj.Sources) > 0 {
		sources := json.JsonArray{}
		for _, src := range aj.Sources {
			source := json.JsonObject{}
			if src.City != "" {
				source.AddString("city", src.City)
			}
			if src.Country != "" {
				source.AddString("country", src.Country)
			}
			if src.Code != "" {
				source.AddString("code", src.Code)
			}
			if src.Url != "" {
				source.AddString("url", src.Url)
			}
			if src.Type != "" {
				source.AddString("type", src.Type)
			}
			if src.Subtype != "" {
				source.AddString("subtype", src.Subtype)
			}
			if src.Name != "" {
				source.AddString("name", src.Name)
			}
			sources.AddObject(&source)
		}
		jo.AddArray("sources", &sources)
	}

	if aj.Contributor != "" {
		jo.AddString("contributor", aj.Contributor)
	}

	if aj.CanonicalLink != "" {
		jo.AddString("canonicallink", aj.CanonicalLink)
	}

	if aj.SourceMaterials != nil && len(aj.SourceMaterials) > 0 {
		sourcematerials := json.JsonArray{}
		for _, src := range aj.SourceMaterials {
			sourcematerial := json.JsonObject{}
			if src.Name != "" {
				sourcematerial.AddString("name", src.Name)
			}
			if src.Code != "" {
				sourcematerial.AddString("code", src.Code)
			}
			if src.Type != "" {
				sourcematerial.AddString("type", src.Type)
			}
			if src.PermissionGranted != "" {
				sourcematerial.AddString("permissiongranted", src.PermissionGranted)
			}

			sourcematerials.AddObject(&sourcematerial)
		}
		jo.AddArray("sourcematerials", &sourcematerials)
	}

	if admin.WorkflowStatus != "" {
		jo.AddString("workflowstatus", admin.WorkflowStatus)
	}

	if !aj.TransmissionSources.IsEmpty() {
		transmissionsources := aj.TransmissionSources.ToJson()
		jo.AddArray("transmissionsources", transmissionsources)
	}

	if !aj.ProductSources.IsEmpty() {
		productsources := aj.ProductSources.ToJson()
		jo.AddArray("productsources", productsources)
	}

	if !aj.ItemContentType.IsEmpty {
		itemcontenttype := json.JsonObject{}
		if aj.ItemContentType.Creator != "" {
			itemcontenttype.AddString("creator", aj.ItemContentType.Creator)
		}
		if aj.ItemContentType.Code != "" {
			itemcontenttype.AddString("code", aj.ItemContentType.Code)
		}
		if aj.ItemContentType.Name != "" {
			itemcontenttype.AddString("name", aj.ItemContentType.Name)
		}
		jo.AddObject("itemcontenttype", &itemcontenttype)
	}

	if admin.Workgroup != "" {
		jo.AddString("workgroup", admin.Workgroup)
	}

	if !aj.DistributionChannels.IsEmpty() {
		distributionchannels := aj.DistributionChannels.ToJson()
		jo.AddArray("distributionchannels", distributionchannels)
	}

	if admin.ContentElement != "" {
		jo.AddString("editorialrole", admin.ContentElement)
	}

	if !aj.Fixture.IsEmpty {
		fixture := json.JsonObject{}
		if aj.Fixture.Code != "" {
			fixture.AddString("code", aj.Fixture.Code)
		}
		if aj.Fixture.Name != "" {
			fixture.AddString("name", aj.Fixture.Name)
		}
		jo.AddObject("fixture", &fixture)
	}

	if !aj.InPackages.IsEmpty() {
		inpackages := aj.InPackages.ToJson()
		jo.AddArray("inpackages", inpackages)
	}

	if admin.Rating != nil {
		ratings := json.JsonArray{}

		for _, r := range admin.Rating {
			if r.Value > 0 && r.ScaleMin > 0 && r.ScaleMax > 0 && r.ScaleUnit != "" {
				rating := json.JsonObject{}
				rating.AddInt("rating", r.Value)
				rating.AddInt("scalemin", r.ScaleMin)
				rating.AddInt("scalemax", r.ScaleMax)
				rating.AddString("scaleunit", r.ScaleUnit)
				if r.Raters > 0 {
					rating.AddInt("raters", r.Raters)
				}
				if r.RaterType != "" {
					rating.AddString("ratertype", r.RaterType)
				}
				ratings.AddObject(&rating)
			}
		}

		if ratings.Length() > 0 {
			jo.AddArray("ratings", &ratings)
		}
	}

	rm := aj.Xml.RightsMetadata

	if rm.Copyright.Holder != "" {
		jo.AddString("copyrightholder", rm.Copyright.Holder)
	}

	if rm.Copyright.Date > 0 {
		jo.AddInt("copyrightdate", rm.Copyright.Date)
	}

	if aj.UsageRights != nil && len(aj.UsageRights) > 0 {
		usagerights := json.JsonArray{}

		for _, ur := range aj.UsageRights {
			usageright := json.JsonObject{}
			if ur.UsageType != "" {
				usageright.AddString("usagetype", ur.UsageType)
			}
			if !ur.Geography.IsEmpty() {
				usageright.AddArray("geography", ur.Geography.ToJson())
			}
			if ur.RightsHolder != "" {
				usageright.AddString("rightsholder", ur.RightsHolder)
			}
			if !ur.Limitations.IsEmpty() {
				usageright.AddArray("limitations", ur.Limitations.ToJson())
			}
			if ur.StartDate != "" {
				usageright.AddString("startdate", ur.StartDate)
			}
			if ur.EndDate != "" {
				usageright.AddString("enddate", ur.EndDate)
			}
			if ur.Groups != nil && len(ur.Groups) > 0 {
				groups := json.JsonArray{}
				for _, g := range ur.Groups {
					group := json.JsonObject{}
					if g.Type != "" {
						group.AddString("type", g.Type)
					}
					if g.Code != "" {
						group.AddString("code", g.Code)
					}
					if g.Name != "" {
						group.AddString("name", g.Name)
					}
					groups.AddObject(&group)
				}
				usageright.AddArray("groups", &groups)
			}

			usagerights.AddObject(&usageright)
		}

		jo.AddArray("usagerights", &usagerights)
	}

	return &jo, nil
}
