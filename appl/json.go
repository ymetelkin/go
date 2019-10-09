package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

//JSON serializes APPL document to JSON
func (doc *Document) JSON() (jo json.Object) {
	jo.AddString("representationversion", "1.0")
	jo.AddString("representationtype", "full")

	//Identification
	jo.AddString("itemid", doc.ItemID)
	jo.AddString("recordid", doc.RecordID)
	jo.AddString("compositeid", doc.CompositeID)
	jo.AddString("compositiontype", doc.CompositionType)
	jo.AddString("type", doc.MediaType)
	addInt(&jo, "priority", doc.Priority)
	addString(&jo, "editorialpriority", doc.EditorialPriority)
	addString(&jo, "language", doc.Language)
	jo.AddInt("recordsequencenumber", doc.RSN)
	addString(&jo, "friendlykey", doc.FriendlyKey)
	addString(&jo, "referenceid", beautify(doc.ReferenceID))

	//Management
	jo.AddString("recordtype", doc.RecordType)
	jo.AddString("filingtype", doc.FilingType)
	addString(&jo, "changeevent", doc.ChangeEvent)
	addString(&jo, "itemkey", doc.ItemKey)
	if doc.ArrivalDateTime != nil {
		jo.AddString("arrivaldatetime", formatDate(*doc.ArrivalDateTime))
	}
	if doc.Created != nil {
		if doc.Created.Date != nil {
			jo.AddString("firstcreated", formatDate(*doc.Created.Date))
		}
		ua := doc.Created.User.json(doc.Created)
		if !ua.IsEmpty() {
			jo.AddObject("firstcreator", ua)
		}
	}
	if doc.Modified != nil {
		if doc.Modified.Date != nil {
			jo.AddString("lastmodifieddatetime", formatDate(*doc.Modified.Date))
		}
		if doc.Modified.User != nil {
			ua := doc.Modified.User.json(nil)
			if !ua.IsEmpty() {
				jo.AddObject("lastmodifier", ua)
			}
		}
	}
	addString(&jo, "pubstatus", doc.Status)
	if doc.ReleaseDateTime != nil {
		jo.AddString("releasedatetime", formatDate(*doc.ReleaseDateTime))
	}
	if doc.Associations != nil {
		var ja json.Array
		for _, v := range doc.Associations {
			ja.AddObject(v.json())
		}
		jo.AddArray("associations", ja)
	}
	addString(&jo, "refersto", doc.RefersTo)
	addStringArray(&jo, "outinginstructions", doc.Outs)
	addString(&jo, "specialinstructions", doc.SpecialInstructions)
	addStringArray(&jo, "editorialtypes", doc.EditorialTypes)
	addString(&jo, "editorialid", doc.EditorialID)
	if doc.Embargoed != nil {
		jo.AddString("embargoed", formatDate(*doc.Embargoed))
	}
	if doc.ItemStartDateTime != nil {
		jo.AddString("itemstartdatetime", formatDate(*doc.ItemStartDateTime))
	}
	if doc.ItemStartDateTimeActual != nil {
		jo.AddString("itemstartdatetimeactual", formatDate(*doc.ItemStartDateTimeActual))
	}
	if doc.ItemExpireDateTime != nil {
		jo.AddString("itemexpiredatetime", formatDate(*doc.ItemExpireDateTime))
	}
	if doc.SearchDateTime != nil {
		jo.AddString("searchdatetime", formatDate(*doc.SearchDateTime))
	}
	if doc.ItemEndDateTime != nil {
		jo.AddString("itemenddatetime", formatDate(*doc.ItemEndDateTime))
	}
	addString(&jo, "function", doc.Function)
	if doc.TimeRestrictions != nil {
		for _, tr := range doc.TimeRestrictions {
			if tr.ID != "" {
				jo.AddBool(tr.ID, tr.Include)
			}
		}
	}
	addStringArray(&jo, "signals", doc.Signals)

	//Newslines
	addString(&jo, "title", beautify(doc.Title))
	addString(&jo, "headline", beautify(doc.Headline))
	addString(&jo, "summary", beautify(doc.Summary))
	addString(&jo, "headline_extended", beautify(doc.ExtendedHeadline))
	if doc.Bylines != nil {
		var ja json.Array
		for _, bl := range doc.Bylines {
			ja.AddObject(bl.json())
		}
		jo.AddArray("bylines", ja)
	}
	if doc.Producer != nil {
		jo.AddObject("producer", doc.Producer.json())
	}
	if doc.Photographer != nil {
		jo.AddObject("photographer", doc.Photographer.json())
	}
	if doc.Captionwriter != nil {
		jo.AddObject("captionwriter", doc.Captionwriter.json())
	}
	if doc.Producer != nil {
		jo.AddObject("producer", doc.Producer.json())
	}
	if doc.Edits != nil {
		var ja json.Array
		for _, e := range doc.Edits {
			var jo json.Object
			jo.AddString("name", e)
			ja.AddObject(jo)
		}
		jo.AddArray("edits", ja)
	}
	addStringArray(&jo, "overlines", doc.Overlines)
	addString(&jo, "dateline", doc.Dateline)
	if doc.Creditline != nil {
		addString(&jo, "creditline", doc.Creditline.Name)
		addString(&jo, "creditlineid", doc.Creditline.Code)
	}
	if doc.Copyright != nil {
		addString(&jo, "copyrightnotice", beautify(doc.Copyright.Notice))
		addString(&jo, "copyrightholder", beautify(doc.Copyright.Holder))
		addInt(&jo, "copyrightdate", doc.Copyright.Year)
	}
	addString(&jo, "rightsline", doc.Rightsline)
	addString(&jo, "seriesline", doc.Seriesline)
	addStringArray(&jo, "keywordlines", doc.Keywordlines)
	addString(&jo, "outcue", doc.OutCue)
	addString(&jo, "locationline", doc.Locationline)

	//AdministrativeMetadata
	if doc.Provider != nil {
		jo.AddObject("provider", doc.Provider.json())
	}
	addString(&jo, "creator", doc.Creator)
	if doc.Sources != nil {
		var ja json.Array
		for _, s := range doc.Sources {
			ja.AddObject(s.json())
		}
		jo.AddArray("sources", ja)
	}
	addString(&jo, "contributor", doc.Contributor)
	if doc.SourceMaterials != nil {
		var ja json.Array
		for _, s := range doc.SourceMaterials {
			ja.AddObject(s.json())
		}
		jo.AddArray("sourcematerials", ja)
	}
	addString(&jo, "canonicallink", doc.CanonicalLink)
	addString(&jo, "workflowstatus", doc.WorkflowStatus)
	addStringArray(&jo, "transmissionsources", doc.TransmissionSources)
	addStringArray(&jo, "productsources", doc.ProductSources)
	if doc.ItemContentType != nil {
		jo.AddObject("itemcontenttype", doc.ItemContentType.json())
	}
	addString(&jo, "workgroup", doc.Workgroup)
	addStringArray(&jo, "distributionchannels", doc.DistributionChannels)
	addString(&jo, "editorialrole", doc.EditorialRole)
	if doc.Fixture != nil {
		jo.AddObject("fixture", doc.Fixture.json())
	}
	addStringArray(&jo, "inpackages", doc.InPackages)
	if doc.Ratings != nil {
		var ja json.Array
		for _, r := range doc.Ratings {
			ja.AddObject(r.json())
		}
		jo.AddArray("ratings", ja)
	}

	//RightsMetadata
	if doc.UsageRights != nil {
		var ja json.Array
		for _, ur := range doc.UsageRights {
			ja.AddObject(ur.json())
		}
		jo.AddArray("usagerights", ja)
	}

	//DescriptiveMetadata
	addStringArray(&jo, "descriptions", doc.Descriptions)
	if doc.DateLineLocation != nil {
		jo.AddObject("datelinelocation", doc.DateLineLocation.json())
	}
	addCodeNameArray(&jo, "generators", doc.Generators, jsongen)
	addCodeNameArray(&jo, "categories", doc.Categories, nil)
	addCodeNameArray(&jo, "suppcategories", doc.SuppCategories, nil)
	addStringArray(&jo, "alertcategories", doc.AlertCategories)
	if doc.Subjects != nil {
		var ja json.Array
		for _, s := range doc.Subjects {
			ja.AddObject(s.json())
		}
		jo.AddArray("subjects", ja)
	}
	if doc.Persons != nil {
		var ja json.Array
		for _, p := range doc.Persons {
			ja.AddObject(p.json())
		}
		jo.AddArray("persons", ja)
	}
	if doc.Organizations != nil {
		var ja json.Array
		for _, org := range doc.Organizations {
			ja.AddObject(org.json())
		}
		jo.AddArray("organizations", ja)
	}
	if doc.Companies != nil {
		var ja json.Array
		for _, c := range doc.Companies {
			ja.AddObject(c.json())
		}
		jo.AddArray("companies", ja)
	}
	if doc.Places != nil {
		var ja json.Array
		for _, p := range doc.Places {
			ja.AddObject(p.json())
		}
		jo.AddArray("places", ja)
	}
	if doc.Events != nil {
		var ja json.Array
		for _, e := range doc.Events {
			ja.AddObject(e.json())
		}
		jo.AddArray("events", ja)
	}
	if doc.Audiences != nil {
		var ja json.Array
		for _, a := range doc.Audiences {
			ja.AddObject(a.jsonaudience())
		}
		jo.AddArray("audiences", ja)
	}
	if doc.Services != nil {
		var ja json.Array
		for _, s := range doc.Services {
			ja.AddObject(s.jsonservice())
		}
		jo.AddArray("services", ja)
	}
	if doc.Perceptions != nil {
		var ja json.Array
		for _, p := range doc.Perceptions {
			ja.AddObject(p.json())
		}
		jo.AddArray("perceptions", ja)
	}
	if doc.ThirdParties != nil {
		var ja json.Array
		for _, tp := range doc.ThirdParties {
			ja.AddObject(tp.json())
		}
		jo.AddArray("thirdpartymeta", ja)
	}

	//FilingMetadata
	if doc.Filings != nil {
		var ja json.Array
		for _, f := range doc.Filings {
			ja.AddObject(f.json())
		}
		jo.AddArray("filings", ja)
	}

	//PublicationComponent
	if doc.Caption != nil {
		jo.AddObject("caption", doc.Caption.json())
	}
	if doc.Shotlist != nil {
		jo.AddObject("shotlist", doc.Shotlist.json())
	}
	if doc.Script != nil {
		jo.AddObject("script", doc.Script.json())
	}
	if doc.PublishableEditorNotes != nil {
		jo.AddObject("publishableeditornotes", doc.PublishableEditorNotes.json())
	}
	if doc.Story != nil {
		jo.AddObject("main", doc.Story.json())
	}
	renditions := doc.Renditions != nil
	if renditions {
		var ja json.Array
		for _, r := range doc.Renditions {
			ja.AddObject(r.json())
		}
		jo.AddArray("renditions", ja)
	}
	if doc.Shots != nil {
		last := len(doc.Shots) - 1
		if last >= 0 {
			if doc.Shots[last].EndTime == 0 && renditions {
				for _, r := range doc.Renditions {
					if r.MediaType == "video" && r.TotalDuration > 0 {
						doc.Shots[last].EndTime = r.TotalDuration
						break
					}
				}
			}
			var ja json.Array
			for _, shot := range doc.Shots {
				ja.AddObject(shot.json())
			}
			jo.AddArray("shots", ja)
		}
	}
	if doc.Parts != nil {
		var ja json.Array
		for _, r := range doc.Parts {
			ja.AddObject(r.json())
		}
		jo.AddArray("parts", ja)
	}

	return
}

func (ua *UserAccount) json(fc *FirstCreated) (jo json.Object) {
	if fc != nil {
		addInt(&jo, "year", fc.Year)
		addInt(&jo, "month", fc.Month)
		addInt(&jo, "day", fc.Day)
		if fc.Time != "" {
			var sb strings.Builder
			if fc.Hour < 10 {
				sb.WriteByte('0')
			}
			sb.WriteString(strconv.Itoa(fc.Hour))
			sb.WriteByte(':')
			if fc.Minute < 10 {
				sb.WriteByte('0')
			}
			sb.WriteString(strconv.Itoa(fc.Minute))
			sb.WriteByte(':')
			if fc.Second < 10 {
				sb.WriteByte('0')
			}
			sb.WriteString(strconv.Itoa(fc.Second))
			jo.AddString("time", sb.String())
			if ua == nil {
				return
			}
		}
	}
	addString(&jo, "username", ua.Name)
	addString(&jo, "useraccount", ua.Account)
	addString(&jo, "useraccountsystem", ua.System)
	addString(&jo, "toolversion", ua.ToolVersion)
	addString(&jo, "userworkgroup", ua.Workgroup)
	addString(&jo, "userlocation", ua.Location)
	return
}

func (ass *Association) json() (jo json.Object) {
	if ass.Type != "" {
		jo.AddString("type", ass.Type)
	}
	jo.AddString("itemid", ass.ItemID)
	jo.AddString("representationtype", "partial")
	jo.AddInt("associationrank", ass.Rank)
	jo.AddInt("typerank", ass.TypeRank)
	return
}

func (cn *CodeName) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(cn.Code))
	addString(&jo, "name", cn.Name)
	return
}

func (cnt *CodeNameTitle) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(cnt.Code))
	addString(&jo, "name", cnt.Name)
	addString(&jo, "title", cnt.Title)
	return
}

func (bl *Byline) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(bl.Code))
	jo.AddString("by", bl.By)
	addString(&jo, "title", bl.Title)
	addString(&jo, "parametric", bl.Parametric)
	return
}

func (p *Provider) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(p.Code))
	addString(&jo, "type", p.Type)
	addString(&jo, "subtype", p.Subtype)
	addString(&jo, "name", p.Name)
	return
}

func (s *Source) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(s.Code))
	addString(&jo, "city", s.City)
	addString(&jo, "county", s.County)
	addString(&jo, "country", s.Country)
	addString(&jo, "countryarea", s.CountryArea)
	addString(&jo, "url", s.URL)
	addString(&jo, "type", s.Type)
	addString(&jo, "subtype", s.Subtype)
	addString(&jo, "name", s.Name)
	return
}

func (s *SourceMaterial) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(s.Code))
	addString(&jo, "type", s.Type)
	addString(&jo, "permissiongranted", s.PermissionGranted)
	addString(&jo, "name", s.Name)
	return
}

func (ict *ItemContentType) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(ict.Code))
	addString(&jo, "creator", ict.Creator)
	addString(&jo, "name", ict.Name)
	return
}

func (r *Rating) json() (jo json.Object) {
	jo.AddInt("rating", r.Value)
	jo.AddInt("scalemin", r.ScaleMin)
	jo.AddInt("scalemax", r.ScaleMax)
	jo.AddString("scaleunit", r.ScaleUnit)
	addInt(&jo, "raters", r.Raters)
	addString(&jo, "ratertype", r.RaterType)
	addString(&jo, "creator", r.Creator)
	return
}

func (ur *UsageRights) json() (jo json.Object) {
	addString(&jo, "usagetype", ur.UsageType)
	addStringArray(&jo, "geography", ur.Geography)
	addString(&jo, "rightsholder", ur.RightsHolder)
	addStringArray(&jo, "limitations", ur.Limitations)
	if ur.StartDate != nil {
		jo.AddString("startdate", formatDate(*ur.StartDate))
	}
	if ur.EndDate != nil {
		jo.AddString("enddate", formatDate(*ur.EndDate))
	}
	if ur.Groups != nil {
		var ja json.Array
		for _, group := range ur.Groups {
			var (
				g  json.Object
				ok bool
			)
			if group.Title != "" {
				ok = true
				g.AddString("type", group.Title)
			}
			if group.Code != "" {
				ok = true
				g.AddString("code", group.Code)
			}
			if group.Name != "" {
				ok = true
				g.AddString("name", group.Name)
			}
			if ok {
				ja.AddObject(g)
			}
		}
		jo.AddArray("groups", ja)
	}

	return
}

func (loc *Location) json() (jo json.Object) {
	addString(&jo, "city", beautify(loc.City))
	addString(&jo, "countryareacode", loc.CountryAreaCode)
	addString(&jo, "countryareaname", beautify(loc.CountryAreaName))
	addString(&jo, "countrycode", loc.CountryCode)
	addString(&jo, "countryname", beautify(loc.CountryName))
	if loc.Geo != nil {
		jo.AddObject("geometry_geojson", loc.Geo.json())
	}
	return
}

func (geo *Geo) json() (jo json.Object) {
	jo.AddString("type", "Point")
	var ja json.Array
	ja.AddFloat(geo.Longitude)
	ja.AddFloat(geo.Latitude)
	jo.AddArray("coordinates", ja)
	return
}

func (sbj *Subject) json() (jo json.Object) {
	jo.AddString("scheme", "http://cv.ap.org/id/")
	jo.AddString("code", strings.ToLower(sbj.Code))
	jo.AddString("name", sbj.Name)
	if sbj.Creator != "" {
		jo.AddString("creator", sbj.Creator)
		if sbj.Creator == "Editorial" {
			jo.AddString("editorial_subject", sbj.Name)
		}
	}
	addStringArray(&jo, "rels", sbj.Rels)
	addStringArray(&jo, "parentids", sbj.ParentIDs)
	if sbj.TopParent != nil {
		if *sbj.TopParent {
			jo.AddBool("topparent", true)
		} else {
			jo.AddBool("topparent", false)
		}
	}
	return
}

func (p *Person) json() (jo json.Object) {
	if p.IsNameline {
		jo.AddString("name", p.Name)
		jo.AddString("creator", "Editorial")
		if p.IsFeatured {
			addStringArray(&jo, "rels", []string{"personfeatured"})
			jo.AddString("person_featured", p.Name)
		}
		return
	}
	jo.AddString("scheme", "http://cv.ap.org/id/")
	jo.AddString("code", strings.ToLower(p.Code))
	jo.AddString("name", p.Name)
	addString(&jo, "creator", p.Creator)
	if p.IsFeatured {
		jo.AddString("person_featured", p.Name)
	}
	addStringArray(&jo, "rels", p.Rels)
	addStringArray(&jo, "types", p.Types)
	addCodeNameArray(&jo, "teams", p.Teams, nil)
	addCodeNameArray(&jo, "associatedevents", p.Events, nil)
	addCodeNameArray(&jo, "associatedstates", p.States, nil)
	addStringArray(&jo, "extids", p.IDs)
	return
}

func (c *Company) json() (jo json.Object) {
	jo.AddString("scheme", "http://cv.ap.org/id/")
	jo.AddString("code", strings.ToLower(c.Code))
	jo.AddString("name", c.Name)
	addString(&jo, "creator", c.Creator)
	addStringArray(&jo, "rels", c.Rels)
	addCodeNameArray(&jo, "industries", c.Industries, nil)
	if c.Symbols != nil {
		var (
			ja json.Array
			ok bool
		)
		for _, symbol := range c.Symbols {
			toks := strings.Split(symbol, ":")
			if len(toks) == 2 {
				var s json.Object
				s.AddString("ticker", toks[1])
				s.AddString("exchange", toks[0])
				s.AddString("instrument", symbol)
				ja.AddObject(s)
				ok = true
			}
		}
		if ok {
			jo.AddArray("symbols", ja)
		}
	}
	return
}

func (p *Place) json() (jo json.Object) {
	jo.AddString("scheme", "http://cv.ap.org/id/")
	jo.AddString("code", strings.ToLower(p.Code))
	jo.AddString("name", p.Name)
	addString(&jo, "creator", p.Creator)
	addStringArray(&jo, "rels", p.Rels)
	addStringArray(&jo, "parentids", p.ParentIDs)
	if p.TopParent != nil {
		if *p.TopParent {
			jo.AddBool("topparent", true)
		} else {
			jo.AddBool("topparent", false)
		}
	}
	if p.LocationType != nil {
		jo.AddObject("locationtype", p.LocationType.json())
	}
	if p.Geo != nil {
		jo.AddObject("geometry_geojson", p.Geo.json())
	}
	return
}

func (e *Event) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(e.Code))
	jo.AddString("name", e.Name)
	addString(&jo, "creator", e.Creator)
	if e.ExternalIDs != nil {
		var ja json.Array
		for _, p := range e.ExternalIDs {
			var ex json.Object
			name := strings.ToLower(p.Name)
			code := strings.ToLower(p.Code)
			ex.AddString("creator", name)
			ex.AddString("code", code)
			ex.AddString("creatorcode", fmt.Sprintf("%s:%s", name, code))
			ja.AddObject(ex)
		}
		jo.AddArray("externaleventids", ja)
	}
	if e.Properties != nil {
		var ep json.Object
		for _, p := range e.Properties {
			ep.AddString(p.Code, p.Name)
		}
		jo.AddObject("eventproperties", ep)
	}
	return
}

func (p *Perception) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(p.Code))
	addString(&jo, "name", p.Name)
	addString(&jo, "creator", p.Creator)
	if p.Rel != "" {
		var ja json.Array
		ja.AddString(p.Rel)
		jo.AddArray("rels", ja)
	}
	return
}

func (tp *ThirdParty) json() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(tp.Code))
	addString(&jo, "name", tp.Name)
	addString(&jo, "creator", tp.Creator)
	addString(&jo, "vocabulary", tp.Vocabulary)
	addString(&jo, "vocabularyowner", tp.VocabularyOwner)
	return
}

func (f *Filing) json() (jo json.Object) {
	addString(&jo, "filingid", f.ID)
	if f.ArrivalDateTime != nil {
		jo.AddString("filingarrivaldatetime", formatDate(*f.ArrivalDateTime))
	}
	addString(&jo, "cycle", f.Cycle)
	addString(&jo, "transmissionreference", f.TransmissionReference)
	addString(&jo, "transmissionfilename", f.TransmissionFilename)
	addString(&jo, "transmissioncontent", f.TransmissionContent)
	addString(&jo, "serviceleveldesignator", f.ServiceLevelDesignator)
	addString(&jo, "selector", f.Selector)
	addString(&jo, "format", f.Format)
	addString(&jo, "filingsource", f.Source)
	addString(&jo, "filingcategory", f.Category)
	if f.Routings != nil {
		var routings json.Object
		for k, v := range f.Routings {
			if v != nil {
				var ja json.Array
				for _, s := range v {
					ja.AddString(s)
				}
				routings.AddArray(k, ja)
			}
		}
		jo.AddObject("routings", routings)
	}
	addString(&jo, "slugline", beautify(f.Slugline))
	addString(&jo, "originalmediaid", f.OriginalMediaID)
	addString(&jo, "importfolder", f.ImportFolder)
	addString(&jo, "importwarnings", f.ImportWarnings)
	addString(&jo, "librarytwincheck", f.LibraryTwinCheck)
	addString(&jo, "libraryrequestid", f.LibraryRequestID)
	addString(&jo, "specialfieldattn", f.SpecialFieldAttn)
	addString(&jo, "feedline", f.Feedline)
	addString(&jo, "libraryrequestlogin", f.LibraryRequestLogin)
	if f.Products != nil {
		var ja json.Array
		for _, p := range f.Products {
			ja.AddInt(p)
		}
		jo.AddArray("products", ja)
	}
	addString(&jo, "priorityline", f.Priorityline)
	if f.ForeignKeys != nil {
		var ja json.Array
		for _, fk := range f.ForeignKeys {
			var o json.Object
			o.AddString(fk.Field, fk.Value)
			ja.AddObject(o)
		}
		jo.AddArray("foreignkeys", ja)
	}
	if f.Countries != nil {
		var ja json.Array
		for _, c := range f.Countries {
			ja.AddString(c)
		}
		jo.AddArray("filingcountries", ja)
	}
	if f.Regions != nil {
		var ja json.Array
		for _, r := range f.Regions {
			ja.AddString(r)
		}
		jo.AddArray("filingregions", ja)
	}
	if f.Subjects != nil {
		var ja json.Array
		for _, s := range f.Subjects {
			ja.AddString(s)
		}
		jo.AddArray("filingsubjects", ja)
	}
	if f.Topics != nil {
		var ja json.Array
		for _, t := range f.Topics {
			ja.AddString(t)
		}
		jo.AddArray("filingtopics", ja)
	}
	addString(&jo, "filingonlinecode", f.OnlineCode)
	addString(&jo, "distributionscope", f.DistributionScope)
	addString(&jo, "breakingnews", f.BreakingNews)
	addString(&jo, "filingstyle", f.Style)
	addString(&jo, "junkline", f.Junkline)
	return
}

func (text *Text) json() (jo json.Object) {
	jo.AddString("nitf", text.Body)
	addInt(&jo, "words", text.Words)
	return
}

func (r *Rendition) json() (jo json.Object) {
	jo.AddString("title", r.Title)
	addString(&jo, "rel", r.Rel)
	addString(&jo, "code", strings.ToLower(r.Code))
	addString(&jo, "type", r.MediaType)
	addString(&jo, "fileextension", r.FileExtension)
	addString(&jo, "tapenumber", r.TapeNumber)
	if r.Attributes != nil {
		for k, v := range r.Attributes {
			jo.AddString(strings.ToLower(k), v)
		}
	}
	addInt(&jo, "sizeinbytes", r.ByteSize)
	addString(&jo, "scene", r.Scene)
	addString(&jo, "sceneid", r.SceneID)
	addString(&jo, "broadcastformat", r.BroadcastFormat)
	addString(&jo, "presentationsystem", r.PresentationSystem)
	addString(&jo, "presentationframe", r.PresentationFrame)
	addString(&jo, "presentationframelocation", r.PresentationFrameLocation)
	addInt(&jo, "width", r.Width)
	addInt(&jo, "height", r.Height)
	addInt(&jo, "resolution", r.Resolution)
	addString(&jo, "resolutionunits", r.ResolutionUnits)
	if r.FrameRate > 0 {
		jo.AddFloat("framerate", r.FrameRate)
	}
	addInt(&jo, "totalduration", r.TotalDuration)
	if r.Characteristics != nil {
		for k, v := range r.Characteristics {
			jo.AddString(strings.ToLower(k), v)
		}
	}
	if r.ForeignKeys != nil {
		var ja json.Array
		for _, fk := range r.ForeignKeys {
			var o json.Object
			o.AddString(fk.Field, fk.Value)
		}
		jo.AddArray("foreignkeys", ja)
	}
	return
}

func (shot *PhotoShot) json() (jo json.Object) {
	jo.AddInt("seq", shot.Sequence)
	addString(&jo, "href", shot.Href)
	addInt(&jo, "width", shot.Width)
	addInt(&jo, "height", shot.Height)
	jo.AddString("start", formatTime(shot.StartTime))
	jo.AddString("end", formatTime(shot.EndTime))
	jo.AddString("timeunit", "normalplaytime")
	return
}

func jsongen(cn CodeName) (jo json.Object) {
	jo.AddString("name", cn.Name)
	jo.AddString("version", cn.Code)
	return
}

func (cnt *CodeNameTitle) jsonaudience() (jo json.Object) {
	addString(&jo, "code", strings.ToLower(cnt.Code))
	addString(&jo, "name", cnt.Name)
	addString(&jo, "type", cnt.Title)
	return
}

func (cn *CodeName) jsonservice() (jo json.Object) {
	if cn.Code == "_apservice" && cn.Name != "" {
		jo.AddString("apservice", cn.Name)
	} else {
		addString(&jo, "code", strings.ToLower(cn.Code))
		addString(&jo, "apsales", cn.Name)
	}
	return
}

func addStringArray(jo *json.Object, name string, values []string) {
	if values != nil {
		var ja json.Array
		for _, value := range values {
			clean := beautify(value)
			if clean != "" {
				ja.AddString(clean)
			}
		}
		jo.AddArray(name, ja)
	}
}

func addCodeNameArray(jo *json.Object, name string, values []CodeName, f func(CodeName) json.Object) {
	if values != nil {
		var ja json.Array
		for _, value := range values {
			var cn json.Object
			if f == nil {
				cn = value.json()
			} else {
				cn = f(value)
			}
			ja.AddObject(cn)
		}
		jo.AddArray(name, ja)
	}
}

func addString(jo *json.Object, name string, value string) {
	if value != "" {
		jo.AddString(name, value)
	}
}

func addInt(jo *json.Object, name string, value int) {
	if value > 0 {
		jo.AddInt(name, value)
	}
}

func beautify(s string) string {
	runes := []rune(s)
	size := len(runes) - 1

	var (
		sb     strings.Builder
		i      int
		sp, st bool
	)

	for {
		if i > size {
			break
		}

		add := true

		r := runes[i]
		if isWS(r) {
			sp = true
			add = false
		} else if r == '&' && i < size-2 {
			j := i + 1
			c := runes[j]
			if c == 'l' || c == 'g' || c == 'a' || c == '#' {
				ok, c, j := decode(c, runes, j, size)
				if ok {
					if isWS(c) {
						sp = true
						add = false
					} else {
						r = c
					}
					i = j
				}
			}
		}

		if add {
			if sp && st {
				sb.WriteByte(' ')
			}
			sb.WriteRune(r)
			st = true
			sp = false
		}

		i++
	}

	return sb.String()
}

func isWS(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '\r' || r == 160 || r == '\f' || r == '\v' || r == '\b'
}

func decode(r rune, runes []rune, i int, size int) (bool, rune, int) {
	if r == 'l' {
		ok, j := matchText(runes, []rune{'t', ';'}, i, size)
		if ok {
			return true, '<', j
		}
	} else if r == 'g' {
		ok, j := matchText(runes, []rune{'t', ';'}, i, size)
		if ok {
			return true, '>', j
		}
	} else if r == 'a' {
		ok, j := matchText(runes, []rune{'m', 'p', ';'}, i, size)
		if ok {
			r = '&'
			if j < size-2 {
				k := j + 1
				c := runes[k]
				if c == 'l' || c == 'g' || c == 'a' || c == '#' {
					ok, c, k := decode(c, runes, k, size)
					if ok {
						r = c
						j = k
					}
				}
			}
			return true, r, j
		}
	} else if r == '#' {
		ok, c, j := matchCode(runes, i+1, size)
		if ok {
			return true, c, j
		}
	}

	return false, r, i
}

func matchText(runes []rune, match []rune, i int, size int) (bool, int) {
	for _, r := range match {
		i++
		if i > size || runes[i] != r {
			return false, i - 1
		}
	}
	return true, i
}

func matchCode(runes []rune, i int, size int) (bool, rune, int) {
	var (
		r     rune
		val   int32
		pos   int
		start = i - 1
		x     bool
		xb    strings.Builder
	)
	for {
		if i > size {
			return false, r, i
		}

		r = runes[i]

		if x {
			if r == ';' {
				d, err := strconv.ParseInt(xb.String(), 16, 32)
				if err == nil {
					return true, rune(d), i
				}
				return false, r, start
			}
			xb.WriteRune(r)
		} else {
			var d int32
			if r > 47 && r < 58 {
				d = r - 48
				if pos == 0 {
					if d == 0 {
						return false, r, start
					}
					val = d
				} else {
					val = val*10 + d
				}
				pos++
			} else if r == ';' {
				if val > 0 {
					return true, rune(val), i
				}
				return false, r, start
			} else if r == 'x' {
				x = true
			} else {
				return false, r, start
			}
		}
		i++
	}
}
