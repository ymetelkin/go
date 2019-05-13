package apql

import "github.com/ymetelkin/go/json"

type es struct {
	Name      string
	Type      string
	Transform []string
}

type field struct {
	Name    string
	Fields  []es
	And     and
	Replace json.Object
}

type and struct {
	Path  string
	Field string
	Value string
}

func newEs(nm string, tp string, tr []string) es {
	return es{
		Name:      nm,
		Type:      tp,
		Transform: tr,
	}
}

func getFields() map[string]field {
	fs := make(map[string]field)

	fs["alert"] = field{
		Fields: []es{
			newEs("alertcategories", "text", []string{"lowercase"}),
		}}

	fs["apcompany_id"] = field{
		Fields: []es{
			newEs("companies.code", "keyword", []string{"lowercase"}),
		}}

	fs["apcompany_name"] = field{
		Fields: []es{
			newEs("companies.name", "text", []string{"lowercase"}),
		}}

	fs["apcountry"] = field{
		Fields: []es{
			newEs("places.code", "keyword", []string{"lowercase"}),
		}}

	fs["apeditorialcountry"] = field{
		Fields: []es{
			newEs("places.code", "keyword", []string{"lowercase"}),
		}}

	fs["apeditorialregion"] = field{
		Fields: []es{
			newEs("places.code", "keyword", []string{"lowercase"}),
		}}

	fs["apeditorialsubject"] = field{
		Fields: []es{
			newEs("subjects.code", "keyword", []string{"lowercase"}),
		},
		And: and{
			Path:  "subjects",
			Field: "subjects.creator",
			Value: "editorial",
		}}

	fs["apevent_id"] = field{
		Fields: []es{
			newEs("events.code", "keyword", []string{"lowercase"}),
		}}

	fs["apevent_name"] = field{
		Fields: []es{
			newEs("events.name", "text", []string{"lowercase"}),
		}}

	fs["apgeography_id"] = field{
		Fields: []es{
			newEs("places.code", "keyword", []string{"lowercase"}),
		}}

	fs["apgeography_name"] = field{
		Fields: []es{
			newEs("places.name", "text", []string{"lowercase"}),
		}}

	fs["apindustry"] = field{
		Fields: []es{
			newEs("companies.industries.name", "text", []string{"lowercase"}),
		}}

	fs["apindustry_id"] = field{
		Fields: []es{
			newEs("companies.industries.code", "keyword", []string{"lowercase"}),
		}}

	fs["aporganization_id"] = field{
		Fields: []es{
			newEs("organizations.code", "keyword", []string{"lowercase"}),
		}}

	fs["aporganization_name"] = field{
		Fields: []es{
			newEs("organizations.name", "text", []string{"lowercase"}),
		}}

	fs["apregion"] = field{
		Fields: []es{
			newEs("places.name", "text", []string{"lowercase"}),
		}}

	fs["arrivaldatetime"] = field{
		Fields: []es{
			newEs("arrivaldatetime", "date", nil),
		}}

	fs["associatedstate"] = field{
		Fields: []es{
			newEs("persons.associatedstates.name", "text", []string{"lowercase"}),
		}}

	fs["associatedwith"] = field{
		Fields: []es{
			newEs("associations.itemid", "keyword", []string{"lowercase"}),
		}}

	fs["audience"] = field{
		Fields: []es{
			newEs("audiences.code", "keyword", []string{"lowercase"}),
		}}

	fs["audiencename"] = field{
		Fields: []es{
			newEs("audiences.name", "text", []string{"lowercase"}),
		}}

	fs["audiocutnumbername"] = field{
		Fields: []es{
			newEs("fixture.code", "text", []string{"lowercase"}),
		}}

	fs["backgroundcolor"] = field{
		Fields: []es{
			newEs("renditions.backgroundcolor", "text", []string{"lowercase"}),
		}}

	fs["body.content"] = field{
		Fields: []es{
			newEs("main.nitf", "text", nil),
		}}

	fs["breakingnews"] = field{
		Fields: []es{
			newEs("filings.breakingnews", "keyword", []string{"lowercase"}),
		}}

	fs["byline"] = field{
		Fields: []es{
			newEs("bylines.by", "text", nil),
			newEs("photographer.name", "text", nil),
			newEs("producer.name", "text", nil),
			newEs("captionwriter.name", "text", nil),
			newEs("editor.name", "text", nil),
		}}

	fs["bylinetitle"] = field{
		Fields: []es{
			newEs("bylines.code", "keyword", []string{"lowercase"}),
			newEs("photographer.code", "keyword", []string{"lowercase"}),
			newEs("producer.code", "keyword", []string{"lowercase"}),
			newEs("captionwriter.code", "keyword", []string{"lowercase"}),
			newEs("editor.code", "keyword", []string{"lowercase"}),
		}}

	fs["bylinetitlename"] = field{
		Fields: []es{
			newEs("bylines.title", "text", nil),
			newEs("photographer.title", "text", nil),
			newEs("captionwriter.title", "text", nil),
			newEs("editor.title", "text", nil),
		}}

	fs["caption"] = field{
		Fields: []es{
			newEs("caption.nitf", "text", nil),
		}}

	fs["category"] = field{
		Fields: []es{
			newEs("categories.code", "keyword", []string{"lowercase"}),
		}}

	fs["categoryname"] = field{
		Fields: []es{
			newEs("categories.name", "text", []string{"lowercase"}),
		}}

	fs["city"] = field{
		Fields: []es{
			newEs("datelinelocation.city", "text", []string{"lowercase"}),
		}}

	fs["colorspace"] = field{
		Fields: []es{
			newEs("renditions.colorspace", "text", []string{"lowercase"}),
		}}

	fs["comment"] = field{
		Fields: []es{
			newEs("services.apservice", "text", []string{"lowercase"}),
		}}

	fs["componentrole"] = field{
		Fields: []es{
			newEs("renditions.role", "keyword", []string{"lowercase"}),
		}}

	fs["compositiontype"] = field{
		Fields: []es{
			newEs("compositiontype", "keyword", []string{"lowercase"}),
		}}

	fs["consumerready"] = field{
		Replace: textQueryJSON("signals", "newscontent", true)}

	fs["content"] = field{
		Fields: []es{
			newEs("main.nitf", "text", nil),
			newEs("caption.nitf", "text", nil),
			newEs("script.nitf", "text", nil),
			newEs("shotlist.nitf", "text", nil),
		}}

	fs["contentelement"] = field{
		Fields: []es{
			newEs("editorialrole", "text", []string{"lowercase"}),
		}}

	fs["contentfileextension"] = field{
		Fields: []es{
			newEs("renditions.fileextension", "keyword", []string{"lowercase"}),
		}}

	fs["contentformat"] = field{
		Fields: []es{
			newEs("renditions.format", "text", []string{"lowercase"}),
		}}

	fs["contenthref"] = field{
		Fields: []es{
			newEs("renditions.href", "keyword", []string{"lowercase"}),
		}}

	fs["contentmimetype"] = field{
		Fields: []es{
			newEs("renditions.mimetype", "text", []string{"lowercase"}),
		}}

	fs["contentsizeinbytes"] = field{
		Fields: []es{
			newEs("renditions.sizeinbytes", "keyword", nil),
		}}

	fs["copyrightdate"] = field{
		Fields: []es{
			newEs("copyrightdate", "keyword", nil),
		}}

	fs["copyrightholder"] = field{
		Fields: []es{
			newEs("copyrightholder", "text", []string{"lowercase"}),
		}}

	fs["copyrightline"] = field{
		Fields: []es{
			newEs("copyrightnotice", "text", []string{"lowercase"}),
		}}

	fs["copyspace"] = field{
		Fields: []es{
			newEs("renditions.copyspace", "text", nil),
		}}

	fs["corporategroupid"] = field{
		Fields: []es{
			newEs("usagerights.groups.code", "keyword", []string{"lowercase"}),
		},
		And: and{
			Path:  "usagerights.groups",
			Field: "usagerights.groups.type",
			Value: "corporate",
		}}

	fs["corporategroupname"] = field{
		Fields: []es{
			newEs("usagerights.groups.name", "text", []string{"lowercase"}),
		},
		And: and{
			Path:  "usagerights.groups",
			Field: "usagerights.groups.type",
			Value: "corporate",
		}}

	fs["country"] = field{
		Fields: []es{
			newEs("datelinelocation.countrycode", "keyword", []string{"lowercase"}),
		}}

	fs["countryname"] = field{
		Fields: []es{
			newEs("datelinelocation.countryname", "text", []string{"lowercase"}),
		}}

	fs["creationdate"] = field{
		Fields: []es{
			newEs("firstcreated", "date", nil),
		}}

	fs["creationdatelower"] = field{
		Fields: []es{
			newEs("firstcreated", "date", nil),
		}}

	fs["creationdateupper"] = field{
		Fields: []es{
			newEs("firstcreated", "date", nil),
		}}

	fs["creationday"] = field{
		Fields: []es{
			newEs("firstcreated", "integer", nil),
		}}

	fs["creationmonth"] = field{
		Fields: []es{
			newEs("firstcreated", "integer", nil),
		}}

	fs["creationyear"] = field{
		Fields: []es{
			newEs("firstcreated", "integer", nil),
		}}

	fs["creator"] = field{
		Fields: []es{
			newEs("creator", "text", []string{"lowercase"}),
		}}

	fs["creditline"] = field{
		Fields: []es{
			newEs("creditline", "text", []string{"lowercase"}),
		}}

	fs["creditlinename"] = field{
		Fields: []es{
			newEs("creditlineid", "keyword", []string{"lowercase"}),
		}}

	fs["cycle"] = field{
		Fields: []es{
			newEs("cycle", "text", nil),
		}}

	fs["dateline"] = field{
		Fields: []es{
			newEs("dateline", "text", []string{"lowercase"}),
		}}

	fs["datedwithin"] = field{
		Fields: []es{
			newEs("firstcreated", "date", nil),
		}}

	fs["defaultlanguage"] = field{
		Fields: []es{
			newEs("language", "keyword", []string{"lowercase"}),
		}}

	fs["description"] = field{
		Fields: []es{
			newEs("descriptions", "text", nil),
		}}

	fs["distributionchannel"] = field{
		Fields: []es{
			newEs("distributionchannels", "text", []string{"lowercase"}),
		}}

	fs["distributionscope"] = field{
		Fields: []es{
			newEs("filings.distributionscope", "text", []string{"lowercase"}),
		}}

	fs["editorialid"] = field{
		Fields: []es{
			newEs("editorialid", "keyword", []string{"lowercase"}),
		}}

	fs["editorialpriority"] = field{
		Fields: []es{
			newEs("editorialpriority", "keyword", []string{"lowercase"}),
		}}

	fs["editorialslug"] = field{
		Fields: []es{
			newEs("filings.slugline", "text", []string{"lowercase"}),
		}}

	fs["editorialsubject"] = field{
		Fields: []es{
			newEs("subjects.code", "keyword", []string{"lowercase"}),
		},
		And: and{
			Path:  "subjects",
			Field: "subjects.creator",
			Value: "editorial",
		}}

	fs["editorialtype"] = field{
		Fields: []es{
			newEs("editorialtypes", "keyword", []string{"lowercase"}),
		}}

	fs["expandedfilingformat"] = field{
		Fields: []es{
			newEs("filings.format", "keyword", []string{"lowercase"}),
		}}

	fs["expandedgroupsidadd"] = field{
		Fields: []es{
			newEs("filings.routings.expandedgroupsidadds", "keyword", []string{"lowercase"}),
		}}

	fs["expandedgroupsidout"] = field{
		Fields: []es{
			newEs("filings.routings.expandedgroupsidouts", "keyword", []string{"lowercase"}),
		}}

	fs["expandedsidadd"] = field{
		Fields: []es{
			newEs("filings.routings.expandedsidadds", "keyword", []string{"lowercase"}),
		}}

	fs["expandedsidout"] = field{
		Fields: []es{
			newEs("filings.routings.expandedsidouts", "keyword", []string{"lowercase"}),
		}}

	fs["explicitwarning"] = field{
		Replace: textQueryJSON("signals", "explicitcontent", true)}

	fs["filingcategory"] = field{
		Fields: []es{
			newEs("filings.filingcategory", "keyword", []string{"lowercase"}),
		}}

	fs["filingcategoryencoded"] = field{
		Fields: []es{
			newEs("filings.filingcategory", "keyword", []string{"decode", "lowercase"}),
		}}

	fs["filingcountry"] = field{
		Fields: []es{
			newEs("filings.filingcountries", "text", []string{"lowercase"}),
		}}

	fs["filingformat"] = field{
		Fields: []es{
			newEs("filings.format", "keyword", []string{"lowercase"}),
		}}

	fs["filingonlinecode"] = field{
		Fields: []es{
			newEs("filings.filingonlinecode", "keyword", nil),
		}}

	fs["filingregion"] = field{
		Fields: []es{
			newEs("filings.filingregions", "text", []string{"lowercase"}),
		}}

	fs["filingslug"] = field{
		Fields: []es{
			newEs("filings.slugline", "text", []string{"lowercase"}),
		}}

	fs["filingsource"] = field{
		Fields: []es{
			newEs("filings.filingsource", "keyword", []string{"lowercase"}),
		}}

	fs["filingstyle"] = field{
		Fields: []es{
			newEs("filings.filingstyle", "text", []string{"lowercase"}),
		}}

	fs["filingsubject"] = field{
		Fields: []es{
			newEs("filings.filingsubjects", "text", []string{"lowercase"}),
		}}

	fs["filingsubsubject"] = field{
		Fields: []es{
			newEs("filings.filingsubjects", "text", []string{"lowercase"}),
		}}

	fs["filingtopic"] = field{
		Fields: []es{
			newEs("filings.filingtopics", "text", []string{"lowercase"}),
		}}

	fs["filingtype"] = field{
		Fields: []es{
			newEs("filingtype", "keyword", []string{"lowercase"}),
		}}

	fs["fixture"] = field{
		Fields: []es{
			newEs("fixture.name", "text", []string{"lowercase"}),
		}}

	fs["fixtureid"] = field{
		Fields: []es{
			newEs("fixture.code", "keyword", []string{"lowercase"}),
		}}

	fs["fixturename"] = field{
		Fields: []es{
			newEs("fixture.name", "text", []string{"lowercase"}),
		}}

	fs["friendlykey"] = field{
		Fields: []es{
			newEs("friendlykey", "keyword", nil),
		}}

	fs["function"] = field{
		Fields: []es{
			newEs("function", "keyword", []string{"lowercase"}),
		}}

	fs["geography"] = field{
		Fields: []es{
			newEs("usagerights.geography", "text", []string{"lowercase"}),
		}}

	fs["groupsidadd"] = field{
		Fields: []es{
			newEs("filings.routings.groupsidadds", "keyword", []string{"lowercase"}),
		}}

	fs["groupsidout"] = field{
		Fields: []es{
			newEs("filings.routings.groupsidouts", "keyword", []string{"lowercase"}),
		}}

	fs["headline"] = field{
		Fields: []es{
			newEs("headline", "text", nil),
		}}

	fs["height"] = field{
		Fields: []es{
			newEs("renditions.height", "keyword", nil),
		}}

	fs["hometownstate"] = field{
		Fields: []es{
			newEs("persons.associatedstates.name", "text", []string{"lowercase"}),
		}}

	fs["hue"] = field{
		Fields: []es{
			newEs("renditions.hue", "text", nil),
		}}

	fs["junkline"] = field{
		Fields: []es{
			newEs("filings.junkline", "text", []string{"lowercase"}),
		}}

	fs["importfolder"] = field{
		Fields: []es{
			newEs("filings.importfolder", "keyword", nil),
		}}

	fs["importwarnings"] = field{
		Fields: []es{
			newEs("filings.importwarnings", "keyword", nil),
		}}

	fs["indcode"] = field{
		Fields: []es{
			newEs("companies.industries.code", "keyword", []string{"lowercase"}),
		}}

	fs["inpackage"] = field{
		Fields: []es{
			newEs("inpackages", "text", []string{"lowercase"}),
		}}

	fs["isdigitized"] = field{
		Replace: not(textQueryJSON("signals", "isnotdigitized", true))}

	fs["itemcontenttype"] = field{
		Fields: []es{
			newEs("itemcontenttype.name", "text", []string{"lowercase"}),
		}}

	fs["itemid"] = field{
		Fields: []es{
			newEs("itemid", "keyword", []string{"lowercase"}),
		}}

	fs["keyword"] = field{
		Fields: []es{
			newEs("keywordlines", "text", []string{"lowercase"}),
		}}

	fs["keywordline"] = field{
		Fields: []es{
			newEs("keywordlines", "text", []string{"lowercase"}),
		}}

	fs["keywords"] = field{
		Fields: []es{
			newEs("keywordlines", "text", []string{"lowercase"}),
		}}

	fs["libraryrequestid"] = field{
		Fields: []es{
			newEs("filings.libraryrequestid", "keyword", []string{"lowercase"}),
		}}

	fs["libraryrequestlogin"] = field{
		Fields: []es{
			newEs("filings.libraryrequestlogin", "keyword", nil),
		}}

	fs["librarytwincheck"] = field{
		Fields: []es{
			newEs("filings.librarytwincheck", "keyword", []string{"lowercase"}),
		}}

	fs["limitations"] = field{
		Fields: []es{
			newEs("usagerights.limitations", "text", []string{"lowercase"}),
		}}

	fs["mediatype"] = field{
		Fields: []es{
			newEs("type", "keyword", []string{"lowercase"}),
		}}

	fs["locationline"] = field{
		Fields: []es{
			newEs("locationline", "text", []string{"lowercase"}),
		}}

	fs["nameline"] = field{
		Fields: []es{
			newEs("persons.name", "text", []string{"lowercase"}),
		},
		And: and{
			Path:  "persons",
			Field: "persons.rels",
			Value: "personfeatured",
		}}

	fs["newspowerdrivetimealaska"] = field{
		Fields: []es{
			newEs("newspowerdrivetimealaska", "boolean", nil),
		}}

	fs["newspowerdrivetimearizona"] = field{
		Fields: []es{
			newEs("newspowerdrivetimearizona", "boolean", nil),
		}}

	fs["newspowerdrivetimeatlantic"] = field{
		Fields: []es{
			newEs("newspowerdrivetimeatlantic", "boolean", nil),
		}}

	fs["newspowerdrivetimecentral"] = field{
		Fields: []es{
			newEs("newspowerdrivetimecentral", "boolean", nil),
		}}

	fs["newspowerdrivetimeeastern"] = field{
		Fields: []es{
			newEs("newspowerdrivetimeeastern", "boolean", nil),
		}}

	fs["newspowerdrivetimehawaii"] = field{
		Fields: []es{
			newEs("newspowerdrivetimehawaii", "boolean", nil),
		}}

	fs["newspowerdrivetimemountain"] = field{
		Fields: []es{
			newEs("newspowerdrivetimemountain", "boolean", nil),
		}}

	fs["newspowerdrivetimepacific"] = field{
		Fields: []es{
			newEs("newspowerdrivetimepacific", "boolean", nil),
		}}

	fs["newspowermorningalaska"] = field{
		Fields: []es{
			newEs("newspowermorningalaska", "boolean", nil),
		}}

	fs["newspowermorningarizona"] = field{
		Fields: []es{
			newEs("newspowermorningarizona", "boolean", nil),
		}}

	fs["newspowermorningatlantic"] = field{
		Fields: []es{
			newEs("newspowermorningatlantic", "boolean", nil),
		}}

	fs["newspowermorningcentral"] = field{
		Fields: []es{
			newEs("newspowermorningcentral", "boolean", nil),
		}}

	fs["newspowermorningeastern"] = field{
		Fields: []es{
			newEs("newspowermorningeastern", "boolean", nil),
		}}

	fs["newspowermorninghawaii"] = field{
		Fields: []es{
			newEs("newspowermorninghawaii", "boolean", nil),
		}}

	fs["newspowermorningmountain"] = field{
		Fields: []es{
			newEs("newspowermorningmountain", "boolean", nil),
		}}

	fs["newspowermorningpacific"] = field{
		Fields: []es{
			newEs("newspowermorningpacific", "boolean", nil),
		}}

	fs["normalizedslug"] = field{
		Fields: []es{
			newEs("filings.slugline", "text", []string{"lowercase"}),
		}}

	fs["onlinecode"] = field{
		Fields: []es{
			newEs("filings.filingonlinecode", "keyword", nil),
		}}

	fs["originalmediaid"] = field{
		Fields: []es{
			newEs("originalmediaid", "text", nil),
		}}

	fs["outinginstructions"] = field{
		Fields: []es{
			newEs("outinginstructions", "text", []string{"lowercase"}),
		}}

	fs["party"] = field{
		Fields: []es{
			newEs("persons.name", "text", []string{"lowercase"}),
		}}

	fs["partytype"] = field{
		Fields: []es{
			newEs("persons.types", "text", []string{"lowercase"}),
		}}

	fs["passcode"] = field{
		Fields: []es{
			newEs("filings.routings.passcodeadds", "keyword", []string{"lowercase"}),
		}}

	fs["passcodeout"] = field{
		Fields: []es{
			newEs("filings.routings.passcodeouts", "keyword", []string{"lowercase"}),
		}}

	fs["person"] = field{
		Fields: []es{
			newEs("persons.code", "keyword", []string{"lowercase"}),
		}}

	fs["person_name"] = field{
		Fields: []es{
			newEs("persons.name", "text", []string{"lowercase"}),
		}}

	fs["primaryticker"] = field{
		Fields: []es{
			newEs("companies.symbols.ticker", "keyword", []string{"lowercase"}),
		}}

	fs["priority"] = field{
		Fields: []es{
			newEs("priority", "integer", nil),
		}}

	fs["productid"] = field{
		Fields: []es{
			newEs("filings.products", "integer", nil),
		}}

	fs["productsource"] = field{
		Fields: []es{
			newEs("productsources", "keyword", []string{"lowercase"}),
		}}

	fs["profcompany"] = field{
		Fields: []es{
			newEs("companies.name", "text", []string{"lowercase"}),
		}}

	fs["profcompanyindustry"] = field{
		Fields: []es{
			newEs("companies.industries.code", "keyword", []string{"lowercase"}),
		}}

	fs["profcompanyticker"] = field{
		Fields: []es{
			newEs("companies.symbols.ticker", "keyword", []string{"lowercase"}),
		}}

	fs["profcountry"] = field{
		Fields: []es{
			newEs("places.name", "text", []string{"lowercase"}),
		}}

	fs["provider"] = field{
		Fields: []es{
			newEs("provider.name", "text", []string{"lowercase"}),
		}}

	fs["providerid"] = field{
		Fields: []es{
			newEs("provider.code", "keyword", []string{"lowercase"}),
		}}

	fs["providername"] = field{
		Fields: []es{
			newEs("provider.name", "text", []string{"lowercase"}),
		}}

	fs["providertype"] = field{
		Fields: []es{
			newEs("provider.type", "keyword", []string{"lowercase"}),
		}}

	fs["ratingvalue"] = field{
		Fields: []es{
			newEs("ratings.rating", "keyword", nil),
		}}

	fs["reach"] = field{
		Fields: []es{
			newEs("signals", "keyword", []string{"lowercase"}),
		}}

	fs["recordid"] = field{
		Fields: []es{
			newEs("recordid", "keyword", []string{"lowercase"}),
		}}

	fs["recordtype"] = field{
		Fields: []es{
			newEs("recordtype", "keyword", []string{"lowercase"}),
		}}

	fs["region"] = field{
		Fields: []es{
			newEs("places.code", "keyword", []string{"lowercase"}),
		}}

	fs["releasedatetime"] = field{
		Fields: []es{
			newEs("releasedatetime", "date", nil),
		}}

	fs["resolution"] = field{
		Fields: []es{
			newEs("renditions.resolution", "keyword", nil),
		}}

	fs["resolutionunits"] = field{
		Fields: []es{
			newEs("renditions.resolutionunit", "text", nil),
		}}

	fs["rightsline"] = field{
		Fields: []es{
			newEs("rightsline", "text", []string{"lowercase"}),
		}}

	fs["salesname"] = field{
		Fields: []es{
			newEs("services.apsales", "text", []string{"lowercase"}),
		}}

	fs["scene"] = field{
		Fields: []es{
			newEs("renditions.scene", "text", []string{"lowercase"}),
		}}

	fs["scenevalue"] = field{
		Fields: []es{
			newEs("renditions.scene", "text", []string{"lowercase"}),
		}}

	fs["selector"] = field{
		Fields: []es{
			newEs("filings.selector", "keyword", []string{"lowercase"}),
		}}

	fs["seriesline"] = field{
		Fields: []es{
			newEs("seriesline", "text", []string{"lowercase"}),
		}}

	fs["sidadd"] = field{
		Fields: []es{
			newEs("filings.routings.sidadds", "keyword", []string{"lowercase"}),
		}}

	fs["sidout"] = field{
		Fields: []es{
			newEs("filings.routings.sidouts", "keyword", []string{"lowercase"}),
		}}

	fs["signal"] = field{
		Fields: []es{
			newEs("signals", "keyword", []string{"lowercase"}),
		}}

	fs["slug"] = field{
		Fields: []es{
			newEs("filings.slugline", "text", []string{"lowercase"}),
		}}

	fs["source"] = field{
		Fields: []es{
			newEs("sources.name", "text", []string{"lowercase"}),
		}}

	fs["sourcecountryarea"] = field{
		Fields: []es{
			newEs("sources.countryarea", "text", []string{"lowercase"}),
		}}

	fs["sourceid"] = field{
		Fields: []es{
			newEs("sources.code", "keyword", []string{"lowercase"}),
		}}

	fs["sourcesubtype"] = field{
		Fields: []es{
			newEs("sources.subtype", "keyword", []string{"lowercase"}),
		}}

	fs["sourcetype"] = field{
		Fields: []es{
			newEs("sources.type", "keyword", []string{"lowercase"}),
		}}

	fs["specialfieldattn"] = field{
		Fields: []es{
			newEs("filings.specialfieldattn", "text", []string{"lowercase"}),
		}}

	fs["specialinstructions"] = field{
		Fields: []es{
			newEs("specialinstructions", "text", []string{"lowercase"}),
		}}

	fs["state"] = field{
		Fields: []es{
			newEs("datelinelocation.countryareacode", "keyword", []string{"lowercase"}),
		}}

	fs["statename"] = field{
		Fields: []es{
			newEs("datelinelocation.countryareaname", "keyword", []string{"lowercase"}),
		}}

	fs["status"] = field{
		Fields: []es{
			newEs("pubstatus", "keyword", []string{"lowercase"}),
		}}

	fs["subject"] = field{
		Fields: []es{
			newEs("subjects.code", "keyword", []string{"lowercase"}),
		}}

	fs["subjectid"] = field{
		Fields: []es{
			newEs("subjects.code", "keyword", []string{"lowercase"}),
		}}

	fs["subjectname"] = field{
		Fields: []es{
			newEs("subjects.name", "text", []string{"lowercase"}),
		}}

	fs["subjectvalue"] = field{
		Fields: []es{
			newEs("subjects.name", "text", []string{"lowercase"}),
		}}

	fs["suppcat"] = field{
		Fields: []es{
			newEs("suppcategories.code", "keyword", []string{"lowercase"}),
		}}

	fs["suppcatencoded"] = field{
		Fields: []es{
			newEs("suppcategories.code", "keyword", []string{"decode", "lowercase"}),
		}}

	fs["suppcatname"] = field{
		Fields: []es{
			newEs("suppcategories.name", "text", []string{"lowercase"}),
		}}

	fs["teragramsubject"] = field{
		Fields: []es{
			newEs("subjects.code", "keyword", []string{"lowercase"}),
		},
		And: and{
			Path:  "subjects",
			Field: "subjects.creator",
			Value: "teragram",
		}}

	fs["ticker"] = field{
		Fields: []es{
			newEs("companies.symbols.ticker", "keyword", []string{"lowercase"}),
		}}

	fs["title"] = field{
		Fields: []es{
			newEs("title", "text", nil),
		}}

	fs["tpmid"] = field{
		Fields: []es{
			newEs("thirdpartymeta.code", "keyword", []string{"lowercase"}),
		}}

	fs["tpmname"] = field{
		Fields: []es{
			newEs("thirdpartymeta.name", "text", []string{"lowercase"}),
		}}

	fs["tpmvocabulary"] = field{
		Fields: []es{
			newEs("thirdpartymeta.vocabulary", "keyword", []string{"lowercase"}),
		}}

	fs["tpmvocabularyowner"] = field{
		Fields: []es{
			newEs("thirdpartymeta.vocabularyowner", "keyword", nil),
		}}

	fs["transmissionfilename"] = field{
		Fields: []es{
			newEs("filings.transmissionfilename", "keyword", nil),
		}}

	fs["transmissionsource"] = field{
		Fields: []es{
			newEs("transmissionsources", "keyword", []string{"lowercase"}),
		}}

	fs["transref"] = field{
		Fields: []es{
			newEs("filings.transmissionreference", "keyword", []string{"lowercase"}),
		}}

	fs["transrefencoded"] = field{
		Fields: []es{
			newEs("filings.transmissionreference", "keyword", []string{"decode"}),
		}}

	fs["usagetype"] = field{
		Fields: []es{
			newEs("usagerights.usagetype", "text", []string{"lowercase"}),
		}}

	fs["words"] = field{
		Fields: []es{
			newEs("caption.words", "integer", nil),
			newEs("script.words", "integer", nil),
			newEs("shotlist.words", "integer", nil),
			newEs("main.words", "integer", nil),
		}}

	fs["workflowstatus"] = field{
		Fields: []es{
			newEs("workflowstatus", "keyword", []string{"lowercase"}),
		}}

	fs["workgroup"] = field{
		Fields: []es{
			newEs("workgroup", "keyword", []string{"lowercase"}),
		}}

	return fs
}
