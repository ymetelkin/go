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
	if doc.Priority > 0 {
		jo.AddInt("priority", doc.Priority)
	}
	if doc.EditorialPriority != "" {
		jo.AddString("editorialpriority", doc.EditorialPriority)
	}
	if doc.Language != "" {
		jo.AddString("language", doc.Language)
	}
	jo.AddInt("recordsequencenumber", doc.RSN)
	if doc.FriendlyKey != "" {
		jo.AddString("friendlykey", doc.FriendlyKey)
	}
	if doc.ReferenceID != "" {
		jo.AddString("referenceid", doc.ReferenceID)
	}

	//Management
	jo.AddString("recordtype", doc.RecordType)
	jo.AddString("filingtype", doc.FilingType)
	if doc.ChangeEvent != "" {
		jo.AddString("changeevent", doc.ChangeEvent)
	}
	if doc.ItemKey != "" {
		jo.AddString("itemkey", doc.ItemKey)
	}
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
	if doc.Status != "" {
		jo.AddString("pubstatus", doc.Status)
	}
	if doc.ReleaseDateTime != nil {
		jo.AddString("releasedatetime", formatDate(*doc.ReleaseDateTime))
	}
	if doc.Associations != nil && len(doc.Associations) > 0 {
		var ja json.Array
		for _, v := range doc.Associations {
			ja.AddObject(v.json())
		}
		jo.AddArray("associations", ja)
	}
	if doc.RefersTo != "" {
		jo.AddString("refersto", doc.RefersTo)
	}
	addStringArray("outinginstructions", doc.Outs, &jo)
	if doc.SpecialInstructions != "" {
		jo.AddString("specialinstructions", doc.SpecialInstructions)
	}
	addStringArray("editorialtypes", doc.EditorialTypes, &jo)
	if doc.EditorialID != "" {
		jo.AddString("editorialid", doc.EditorialID)
	}
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
	if doc.Function != "" {
		jo.AddString("function", doc.Function)
	}
	if doc.TimeRestrictions != nil {
		for _, tr := range doc.TimeRestrictions {
			if tr.ID != "" {
				jo.AddBool(tr.ID, tr.Include)
			}
		}
	}
	addStringArray("signals", doc.Signals, &jo)

	//Newslines
	if doc.Title != "" {
		jo.AddString("title", strings.TrimSpace(doc.Title))
	}
	jo.AddString("headline", strings.TrimSpace(doc.Headline))
	if doc.Summary != "" {
		jo.AddString("summary", strings.TrimSpace(doc.Summary))
	}
	if doc.ExtendedHeadline != "" {
		jo.AddString("headline_extended", strings.TrimSpace(doc.ExtendedHeadline))
	}
	if doc.Bylines != nil && len(doc.Bylines) > 0 {
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
	if doc.Edits != nil && len(doc.Edits) > 0 {
		var ja json.Array
		for _, e := range doc.Edits {
			var jo json.Object
			jo.AddString("name", e)
			ja.AddObject(jo)
		}
		jo.AddArray("edits", ja)
	}
	addStringArray("overlines", doc.Overlines, &jo)
	if doc.Dateline != "" {
		jo.AddString("dateline", doc.Dateline)
	}
	if doc.Creditline != nil {
		if doc.Creditline.Name != "" {
			jo.AddString("creditline", doc.Creditline.Name)
		}
		if doc.Creditline.Code != "" {
			jo.AddString("creditlineid", doc.Creditline.Code)
		}
	}
	if doc.Copyright != nil {
		if doc.Copyright.Notice != "" {
			jo.AddString("copyrightnotice", doc.Copyright.Notice)
		}
		if doc.Copyright.Holder != "" {
			jo.AddString("copyrightholder", doc.Copyright.Holder)
		}
		if doc.Copyright.Year > 0 {
			jo.AddInt("copyrightdate", doc.Copyright.Year)
		}
	}
	if doc.Rightsline != "" {
		jo.AddString("rightsline", doc.Rightsline)
	}
	if doc.Seriesline != "" {
		jo.AddString("seriesline", doc.Seriesline)
	}
	addStringArray("keywordlines", doc.Keywordlines, &jo)
	if doc.OutCue != "" {
		jo.AddString("outcue", doc.OutCue)
	}
	if doc.Locationline != "" {
		jo.AddString("locationline", doc.Locationline)
	}

	//AdministrativeMetadata
	if doc.Provider != nil {
		jo.AddObject("provider", doc.Provider.json())
	}
	if doc.Creator != "" {
		jo.AddString("creator", doc.Creator)
	}
	if doc.Sources != nil && len(doc.Sources) > 0 {
		var ja json.Array
		for _, s := range doc.Sources {
			ja.AddObject(s.json())
		}
		jo.AddArray("sources", ja)
	}
	if doc.Contributor != "" {
		jo.AddString("contributor", doc.Contributor)
	}
	if doc.SourceMaterials != nil && len(doc.SourceMaterials) > 0 {
		var ja json.Array
		for _, s := range doc.SourceMaterials {
			ja.AddObject(s.json())
		}
		jo.AddArray("sourcematerials", ja)
	}
	if doc.CanonicalLink != "" {
		jo.AddString("canonicallink", doc.CanonicalLink)
	}
	if doc.WorkflowStatus != "" {
		jo.AddString("workflowstatus", doc.WorkflowStatus)
	}
	addStringArray("transmissionsources", doc.TransmissionSources, &jo)
	addStringArray("productsources", doc.ProductSources, &jo)
	if doc.ItemContentType != nil {
		jo.AddObject("itemcontenttype", doc.ItemContentType.json())
	}
	if doc.Workgroup != "" {
		jo.AddString("workgroup", doc.Workgroup)
	}
	addStringArray("distributionchannels", doc.DistributionChannels, &jo)
	if doc.EditorialRole != "" {
		jo.AddString("editorialrole", doc.EditorialRole)
	}
	if doc.Fixture != nil {
		jo.AddObject("fixture", doc.Fixture.json())
	}
	addStringArray("inpackages", doc.InPackages, &jo)
	if doc.Ratings != nil && len(doc.Ratings) > 0 {
		var ja json.Array
		for _, r := range doc.Ratings {
			ja.AddObject(r.json())
		}
		jo.AddArray("ratings", ja)
	}

	//RightsMetadata
	if doc.UsageRights != nil && len(doc.UsageRights) > 0 {
		var ja json.Array
		for _, ur := range doc.UsageRights {
			ja.AddObject(ur.json())
		}
		jo.AddArray("usagerights", ja)
	}

	//DescriptiveMetadata
	addStringArray("descriptions", doc.Descriptions, &jo)
	if doc.DateLineLocation != nil {
		jo.AddObject("datelinelocation", doc.DateLineLocation.json())
	}
	addCodeNameArray("generators", doc.Generators, jsongen, &jo)
	addCodeNameArray("categories", doc.Categories, nil, &jo)
	addCodeNameArray("suppcategories", doc.SuppCategories, nil, &jo)
	addStringArray("alertcategories", doc.AlertCategories, &jo)
	if doc.Subjects != nil && len(doc.Subjects) > 0 {
		var ja json.Array
		for _, s := range doc.Subjects {
			ja.AddObject(s.json())
		}
		jo.AddArray("subjects", ja)
	}
	if doc.Persons != nil && len(doc.Persons) > 0 {
		var ja json.Array
		for _, p := range doc.Persons {
			ja.AddObject(p.json())
		}
		jo.AddArray("persons", ja)
	}
	if doc.Organizations != nil && len(doc.Organizations) > 0 {
		var ja json.Array
		for _, org := range doc.Organizations {
			ja.AddObject(org.json())
		}
		jo.AddArray("organizations", ja)
	}
	if doc.Companies != nil && len(doc.Companies) > 0 {
		var ja json.Array
		for _, c := range doc.Companies {
			ja.AddObject(c.json())
		}
		jo.AddArray("companies", ja)
	}
	if doc.Places != nil && len(doc.Places) > 0 {
		var ja json.Array
		for _, p := range doc.Places {
			ja.AddObject(p.json())
		}
		jo.AddArray("places", ja)
	}
	if doc.Events != nil && len(doc.Events) > 0 {
		var ja json.Array
		for _, e := range doc.Events {
			ja.AddObject(e.json())
		}
		jo.AddArray("events", ja)
	}
	if doc.Audiences != nil && len(doc.Audiences) > 0 {
		var ja json.Array
		for _, a := range doc.Audiences {
			ja.AddObject(a.jsonaudience())
		}
		jo.AddArray("audiences", ja)
	}
	if doc.Services != nil && len(doc.Services) > 0 {
		var ja json.Array
		for _, s := range doc.Services {
			ja.AddObject(s.jsonservice())
		}
		jo.AddArray("services", ja)
	}
	if doc.Perceptions != nil && len(doc.Perceptions) > 0 {
		var ja json.Array
		for _, p := range doc.Perceptions {
			ja.AddObject(p.json())
		}
		jo.AddArray("perceptions", ja)
	}
	if doc.ThirdParties != nil && len(doc.ThirdParties) > 0 {
		var ja json.Array
		for _, tp := range doc.ThirdParties {
			ja.AddObject(tp.json())
		}
		jo.AddArray("thirdpartymeta", ja)
	}

	//FilingMetadata
	if doc.Filings != nil && len(doc.Filings) > 0 {
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
	renditions := doc.Renditions != nil && len(doc.Renditions) > 0
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
	if doc.Parts != nil && len(doc.Parts) > 0 {
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
		if fc.Year > 0 {
			jo.AddInt("year", fc.Year)
		}
		if fc.Month > 0 {
			jo.AddInt("month", fc.Month)
		}
		if fc.Day > 0 {
			jo.AddInt("day", fc.Day)
		}
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
	if ua.Name != "" {
		jo.AddString("username", ua.Name)
	}
	if ua.Account != "" {
		jo.AddString("useraccount", ua.Account)
	}
	if ua.System != "" {
		jo.AddString("useraccountsystem", ua.System)
	}
	if ua.ToolVersion != "" {
		jo.AddString("toolversion", ua.ToolVersion)
	}
	if ua.Workgroup != "" {
		jo.AddString("userworkgroup", ua.Workgroup)
	}
	if ua.Location != "" {
		jo.AddString("userlocation", ua.Location)
	}

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
	if cn.Code != "" {
		jo.AddString("code", strings.ToLower(cn.Code))
	}
	if cn.Name != "" {
		jo.AddString("name", cn.Name)
	}
	return
}

func (cnt *CodeNameTitle) json() (jo json.Object) {
	if cnt.Code != "" {
		jo.AddString("code", strings.ToLower(cnt.Code))
	}
	if cnt.Name != "" {
		jo.AddString("name", cnt.Name)
	}
	if cnt.Title != "" {
		jo.AddString("title", cnt.Title)
	}
	return
}

func (bl *Byline) json() (jo json.Object) {
	if bl.Code != "" {
		jo.AddString("code", strings.ToLower(bl.Code))
	}
	jo.AddString("by", bl.By)
	if bl.Title != "" {
		jo.AddString("title", bl.Title)
	}
	if bl.Parametric != "" {
		jo.AddString("parametric", bl.Parametric)
	}
	return
}

func (p *Provider) json() (jo json.Object) {
	if p.Code != "" {
		jo.AddString("code", strings.ToLower(p.Code))
	}
	if p.Type != "" {
		jo.AddString("type", p.Type)
	}
	if p.Subtype != "" {
		jo.AddString("subtype", p.Subtype)
	}
	if p.Name != "" {
		jo.AddString("name", p.Name)
	}
	return
}

func (s *Source) json() (jo json.Object) {
	if s.Code != "" {
		jo.AddString("code", strings.ToLower(s.Code))
	}
	if s.City != "" {
		jo.AddString("city", s.City)
	}
	if s.County != "" {
		jo.AddString("county", s.County)
	}
	if s.Country != "" {
		jo.AddString("country", s.Country)
	}
	if s.CountryArea != "" {
		jo.AddString("countryarea", s.CountryArea)
	}
	if s.URL != "" {
		jo.AddString("url", s.URL)
	}
	if s.Type != "" {
		jo.AddString("type", s.Type)
	}
	if s.Subtype != "" {
		jo.AddString("subtype", s.Subtype)
	}
	if s.Name != "" {
		jo.AddString("name", s.Name)
	}
	return
}

func (s *SourceMaterial) json() (jo json.Object) {
	if s.Code != "" {
		jo.AddString("code", strings.ToLower(s.Code))
	}
	if s.Type != "" {
		jo.AddString("type", s.Type)
	}
	if s.PermissionGranted != "" {
		jo.AddString("permissiongranted", s.PermissionGranted)
	}
	if s.Name != "" {
		jo.AddString("name", s.Name)
	}
	return
}

func (ict *ItemContentType) json() (jo json.Object) {
	if ict.Code != "" {
		jo.AddString("code", strings.ToLower(ict.Code))
	}
	if ict.Creator != "" {
		jo.AddString("creator", ict.Creator)
	}
	if ict.Name != "" {
		jo.AddString("name", ict.Name)
	}
	return
}

func (r *Rating) json() (jo json.Object) {
	jo.AddInt("rating", r.Value)
	jo.AddInt("scalemin", r.ScaleMin)
	jo.AddInt("scalemax", r.ScaleMax)
	jo.AddString("scaleunit", r.ScaleUnit)
	if r.Raters > 0 {
		jo.AddInt("raters", r.Raters)
	}
	if r.RaterType != "" {
		jo.AddString("ratertype", r.RaterType)
	}
	if r.Creator != "" {
		jo.AddString("creator", r.Creator)
	}
	return
}

func (ur *UsageRights) json() (jo json.Object) {
	if ur.UsageType != "" {
		jo.AddString("usagetype", ur.UsageType)
	}
	addStringArray("geography", ur.Geography, &jo)
	if ur.RightsHolder != "" {
		jo.AddString("rightsholder", ur.RightsHolder)
	}
	addStringArray("limitations", ur.Limitations, &jo)
	if ur.StartDate != nil {
		jo.AddString("startdate", formatDate(*ur.StartDate))
	}
	if ur.EndDate != nil {
		jo.AddString("enddate", formatDate(*ur.EndDate))
	}
	if ur.Groups != nil && len(ur.Groups) > 0 {
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
	if loc.City != "" {
		jo.AddString("city", loc.City)
	}
	if loc.CountryAreaCode != "" {
		jo.AddString("countryareacode", loc.CountryAreaCode)
	}
	if loc.CountryAreaName != "" {
		jo.AddString("countryareaname", loc.CountryAreaName)
	}
	if loc.CountryCode != "" {
		jo.AddString("countrycode", loc.CountryCode)
	}
	if loc.CountryName != "" {
		jo.AddString("countryname", loc.CountryName)
	}
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
	addStringArray("rels", sbj.Rels, &jo)
	addStringArray("parentids", sbj.ParentIDs, &jo)
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
			addStringArray("rels", []string{"personfeatured"}, &jo)
			jo.AddString("person_featured", p.Name)
		}
		return
	}
	jo.AddString("scheme", "http://cv.ap.org/id/")
	jo.AddString("code", strings.ToLower(p.Code))
	jo.AddString("name", p.Name)
	if p.Creator != "" {
		jo.AddString("creator", p.Creator)
	}
	if p.IsFeatured {
		jo.AddString("person_featured", p.Name)
	}
	addStringArray("rels", p.Rels, &jo)
	addStringArray("types", p.Types, &jo)
	addCodeNameArray("teams", p.Teams, nil, &jo)
	addCodeNameArray("associatedevents", p.Events, nil, &jo)
	addCodeNameArray("associatedstates", p.States, nil, &jo)
	addStringArray("extids", p.IDs, &jo)
	return
}

func (c *Company) json() (jo json.Object) {
	jo.AddString("scheme", "http://cv.ap.org/id/")
	jo.AddString("code", strings.ToLower(c.Code))
	jo.AddString("name", c.Name)
	if c.Creator != "" {
		jo.AddString("creator", c.Creator)
	}
	addStringArray("rels", c.Rels, &jo)
	addCodeNameArray("industries", c.Industries, nil, &jo)
	if c.Symbols != nil && len(c.Symbols) > 0 {
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
	if p.Creator != "" {
		jo.AddString("creator", p.Creator)
	}
	addStringArray("rels", p.Rels, &jo)
	addStringArray("parentids", p.ParentIDs, &jo)
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
	if e.Code != "" {
		jo.AddString("code", strings.ToLower(e.Code))
	}
	jo.AddString("name", e.Name)
	if e.Creator != "" {
		jo.AddString("creator", e.Creator)
	}
	if e.ExternalIDs != nil && len(e.ExternalIDs) > 0 {
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
	if e.Properties != nil && len(e.Properties) > 0 {
		var ep json.Object
		for _, p := range e.Properties {
			ep.AddString(p.Code, p.Name)
		}
		jo.AddObject("eventproperties", ep)
	}
	return
}

func (p *Perception) json() (jo json.Object) {
	if p.Code != "" {
		jo.AddString("code", strings.ToLower(p.Code))
	}
	if p.Name != "" {
		jo.AddString("name", p.Name)
	}
	if p.Creator != "" {
		jo.AddString("creator", p.Creator)
	}
	if p.Rel != "" {
		var ja json.Array
		ja.AddString(p.Rel)
		jo.AddArray("rels", ja)
	}
	return
}

func (tp *ThirdParty) json() (jo json.Object) {
	if tp.Code != "" {
		jo.AddString("code", strings.ToLower(tp.Code))
	}
	if tp.Name != "" {
		jo.AddString("name", tp.Name)
	}
	if tp.Creator != "" {
		jo.AddString("creator", tp.Creator)
	}
	if tp.Vocabulary != "" {
		jo.AddString("vocabulary", tp.Vocabulary)
	}
	if tp.VocabularyOwner != "" {
		jo.AddString("vocabularyowner", tp.VocabularyOwner)
	}
	return
}

func (f *Filing) json() (jo json.Object) {
	if f.ID != "" {
		jo.AddString("filingid", f.ID)
	}
	if f.ArrivalDateTime != nil {
		jo.AddString("filingarrivaldatetime", formatDate(*f.ArrivalDateTime))
	}
	if f.Cycle != "" {
		jo.AddString("cycle", f.Cycle)
	}
	if f.TransmissionReference != "" {
		jo.AddString("transmissionreference", f.TransmissionReference)
	}
	if f.TransmissionFilename != "" {
		jo.AddString("transmissionfilename", f.TransmissionFilename)
	}
	if f.TransmissionContent != "" {
		jo.AddString("transmissioncontent", f.TransmissionContent)
	}
	if f.ServiceLevelDesignator != "" {
		jo.AddString("serviceleveldesignator", f.ServiceLevelDesignator)
	}
	if f.Selector != "" {
		jo.AddString("selector", f.Selector)
	}
	if f.Format != "" {
		jo.AddString("format", f.Format)
	}
	if f.Source != "" {
		jo.AddString("filingsource", f.Source)
	}
	if f.Category != "" {
		jo.AddString("filingcategory", f.Category)
	}
	if f.Routings != nil && len(f.Routings) > 0 {
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
	if f.Slugline != "" {
		jo.AddString("slugline", strings.TrimSpace(f.Slugline))
	}
	if f.OriginalMediaID != "" {
		jo.AddString("originalmediaid", f.OriginalMediaID)
	}
	if f.ImportFolder != "" {
		jo.AddString("importfolder", f.ImportFolder)
	}
	if f.ImportWarnings != "" {
		jo.AddString("importwarnings", f.ImportWarnings)
	}
	if f.LibraryTwinCheck != "" {
		jo.AddString("librarytwincheck", f.LibraryTwinCheck)
	}
	if f.LibraryRequestID != "" {
		jo.AddString("libraryrequestid", f.LibraryRequestID)
	}
	if f.SpecialFieldAttn != "" {
		jo.AddString("specialfieldattn", f.SpecialFieldAttn)
	}
	if f.Feedline != "" {
		jo.AddString("feedline", f.Feedline)
	}
	if f.LibraryRequestLogin != "" {
		jo.AddString("libraryrequestlogin", f.LibraryRequestLogin)
	}
	if f.Products != nil && len(f.Products) > 0 {
		var ja json.Array
		for _, p := range f.Products {
			ja.AddInt(p)
		}
		jo.AddArray("products", ja)
	}
	if f.Priorityline != "" {
		jo.AddString("priorityline", f.Priorityline)
	}
	if f.ForeignKeys != nil && len(f.ForeignKeys) > 0 {
		var ja json.Array
		for _, fk := range f.ForeignKeys {
			var o json.Object
			o.AddString(fk.Field, fk.Value)
			ja.AddObject(o)
		}
		jo.AddArray("foreignkeys", ja)
	}
	if f.Countries != nil && len(f.Countries) > 0 {
		var ja json.Array
		for _, c := range f.Countries {
			ja.AddString(c)
		}
		jo.AddArray("filingcountries", ja)
	}
	if f.Regions != nil && len(f.Regions) > 0 {
		var ja json.Array
		for _, r := range f.Regions {
			ja.AddString(r)
		}
		jo.AddArray("filingregions", ja)
	}
	if f.Subjects != nil && len(f.Subjects) > 0 {
		var ja json.Array
		for _, s := range f.Subjects {
			ja.AddString(s)
		}
		jo.AddArray("filingsubjects", ja)
	}
	if f.Topics != nil && len(f.Topics) > 0 {
		var ja json.Array
		for _, t := range f.Topics {
			ja.AddString(t)
		}
		jo.AddArray("filingtopics", ja)
	}
	if f.OnlineCode != "" {
		jo.AddString("filingonlinecode", f.OnlineCode)
	}
	if f.DistributionScope != "" {
		jo.AddString("distributionscope", f.DistributionScope)
	}
	if f.BreakingNews != "" {
		jo.AddString("breakingnews", f.BreakingNews)
	}
	if f.Style != "" {
		jo.AddString("filingstyle", f.Style)
	}
	if f.Junkline != "" {
		jo.AddString("junkline", f.Junkline)
	}
	return
}

func (text *Text) json() (jo json.Object) {
	jo.AddString("nitf", text.Body)
	if text.Words > 0 {
		jo.AddInt("words", text.Words)
	}
	return
}

func (r *Rendition) json() (jo json.Object) {
	jo.AddString("title", r.Title)
	if r.Rel != "" {
		jo.AddString("rel", r.Rel)
	}
	jo.AddString("code", strings.ToLower(r.Code))
	if r.MediaType != "" {
		jo.AddString("type", r.MediaType)
	}
	if r.FileExtension != "" {
		jo.AddString("fileextension", r.FileExtension)
	}
	if r.TapeNumber != "" {
		jo.AddString("tapenumber", r.TapeNumber)
	}
	if r.Attributes != nil {
		for k, v := range r.Attributes {
			jo.AddString(strings.ToLower(k), v)
		}
	}
	if r.ByteSize > 0 {
		jo.AddInt("sizeinbytes", r.ByteSize)
	}
	if r.Scene != "" {
		jo.AddString("scene", r.Scene)
	}
	if r.SceneID != "" {
		jo.AddString("sceneid", r.SceneID)
	}
	if r.BroadcastFormat != "" {
		jo.AddString("broadcastformat", r.BroadcastFormat)
	}
	if r.PresentationSystem != "" {
		jo.AddString("presentationsystem", r.PresentationSystem)
	}
	if r.PresentationFrame != "" {
		jo.AddString("presentationframe", r.PresentationFrame)
	}
	if r.PresentationFrameLocation != "" {
		jo.AddString("presentationframelocation", r.PresentationFrameLocation)
	}
	if r.Width > 0 {
		jo.AddInt("width", r.Width)
	}
	if r.Height > 0 {
		jo.AddInt("height", r.Height)
	}
	if r.Resolution > 0 {
		jo.AddInt("resolution", r.Resolution)
	}
	if r.ResolutionUnits != "" {
		jo.AddString("resolutionunits", r.ResolutionUnits)
	}
	if r.FrameRate > 0 {
		jo.AddFloat("framerate", r.FrameRate)
	}
	if r.TotalDuration > 0 {
		jo.AddInt("totalduration", r.TotalDuration)
	}
	if r.Characteristics != nil {
		for k, v := range r.Characteristics {
			jo.AddString(strings.ToLower(k), v)
		}
	}
	if r.ForeignKeys != nil && len(r.ForeignKeys) > 0 {
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
	if shot.Href != "" {
		jo.AddString("href", shot.Href)
	}
	if shot.Width > 0 {
		jo.AddInt("width", shot.Width)
	}
	if shot.Height > 0 {
		jo.AddInt("height", shot.Height)
	}
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
	if cnt.Code != "" {
		jo.AddString("code", strings.ToLower(cnt.Code))
	}
	if cnt.Name != "" {
		jo.AddString("name", cnt.Name)
	}
	if cnt.Title != "" {
		jo.AddString("type", cnt.Title)
	}
	return
}

func (cn *CodeName) jsonservice() (jo json.Object) {
	if cn.Code == "_apservice" && cn.Name != "" {
		jo.AddString("apservice", cn.Name)
	} else {
		if cn.Code != "" {
			jo.AddString("code", strings.ToLower(cn.Code))
		}
		if cn.Name != "" {
			jo.AddString("apsales", cn.Name)
		}
	}
	return
}

func addStringArray(name string, values []string, jo *json.Object) {
	if values != nil && len(values) > 0 {
		var ja json.Array
		for _, value := range values {
			clean := strings.TrimSpace(value)
			if clean != "" {
				ja.AddString(clean)
			}
		}
		jo.AddArray(name, ja)
	}
}

func addCodeNameArray(name string, values []CodeName, f func(CodeName) json.Object, jo *json.Object) {
	if values != nil && len(values) > 0 {
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
