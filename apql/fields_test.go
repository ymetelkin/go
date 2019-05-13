package apql

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestTemp(t *testing.T) {
	s := `
	{
		"alert": {
		  "fields": [
			{
			  "name": "alertcategories",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apcompany_id": {
		  "fields": [
			{
			  "name": "companies.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apcompany_name": {
		  "fields": [
			{
			  "name": "companies.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apcountry": {
		  "fields": [
			{
			  "name": "places.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apeditorialcountry": {
		  "fields": [
			{
			  "name": "places.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apeditorialregion": {
		  "fields": [
			{
			  "name": "places.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apeditorialsubject": {
		  "fields": [
			{
			  "name": "subjects.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ],
		  "and": {
			"path": "subjects",
			"query": {
			  "term": {
				"subjects.creator": "editorial"
			  }
			}
		  }
		},
		"apevent_id": {
		  "fields": [
			{
			  "name": "events.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apevent_name": {
		  "fields": [
			{
			  "name": "events.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apgeography_id": {
		  "fields": [
			{
			  "name": "places.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apgeography_name": {
		  "fields": [
			{
			  "name": "places.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apindustry": {
		  "fields": [
			{
			  "name": "companies.industries.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apindustry_id": {
		  "fields": [
			{
			  "name": "companies.industries.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"aporganization_id": {
		  "fields": [
			{
			  "name": "organizations.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"aporganization_name": {
		  "fields": [
			{
			  "name": "organizations.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"apregion": {
		  "fields": [
			{
			  "name": "places.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"arrivaldatetime": {
		  "fields": [
			{
			  "name": "arrivaldatetime",
			  "type": "date"
			}
		  ]
		},
		"associatedstate": {
		  "fields": [
			{
			  "name": "persons.associatedstates.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"associatedwith": {
		  "fields": [
			{
			  "name": "associations.itemid",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"audience": {
		  "fields": [
			{
			  "name": "audiences.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"audiencename": {
		  "fields": [
			{
			  "name": "audiences.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"audiocutnumbername": {
		  "fields": [
			{
			  "name": "fixture.code",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"backgroundcolor": {
		  "fields": [
			{
			  "name": "renditions.backgroundcolor",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"body.content": {
		  "fields": [
			{
			  "name": "main.nitf",
			  "type": "text"
			}
		  ]
		},
		"breakingnews": {
		  "fields": [
			{
			  "name": "filings.breakingnews",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"byline": {
		  "fields": [
			{
			  "name": "bylines.by",
			  "type": "text"
			},
			{
			  "name": "photographer.name",
			  "type": "text"
			},
			{
			  "name": "producer.name",
			  "type": "text"
			},
			{
			  "name": "captionwriter.name",
			  "type": "text"
			},
			{
			  "name": "editor.name",
			  "type": "text"
			}
		  ]
		},
		"bylinetitle": {
		  "fields": [
			{
			  "name": "bylines.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			},
			{
			  "name": "photographer.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			},
			{
			  "name": "producer.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			},
			{
			  "name": "captionwriter.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			},
			{
			  "name": "editor.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"bylinetitlename": {
		  "fields": [
			{
			  "name": "bylines.title",
			  "type": "text"
			},
			{
			  "name": "photographer.title",
			  "type": "text"
			},
			{
			  "name": "captionwriter.title",
			  "type": "text"
			},
			{
			  "name": "editor.title",
			  "type": "text"
			}
		  ]
		},
		"caption": {
		  "fields": [
			{
			  "name": "caption.nitf",
			  "type": "text"
			}
		  ]
		},
		"category": {
		  "fields": [
			{
			  "name": "categories.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"categoryname": {
		  "fields": [
			{
			  "name": "categories.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"city": {
		  "fields": [
			{
			  "name": "datelinelocation.city",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"colorspace": {
		  "fields": [
			{
			  "name": "renditions.colorspace",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"comment": {
		  "fields": [
			{
			  "name": "services.apservice",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"componentrole": {
		  "fields": [
			{
			  "name": "renditions.role",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"compositiontype": {
		  "fields": [
			{
			  "name": "compositiontype",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"consumerready": {
		  "replace": {
			"term": {
			  "signals": "newscontent"
			}
		  }
		},
		"content": {
		  "fields": [
			{
			  "name": "main.nitf",
			  "type": "text"
			},
			{
			  "name": "caption.nitf",
			  "type": "text"
			},
			{
			  "name": "script.nitf",
			  "type": "text"
			},
			{
			  "name": "shotlist.nitf",
			  "type": "text"
			}
		  ]
		},
		"contentelement": {
		  "fields": [
			{
			  "name": "editorialrole",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"contentfileextension": {
		  "fields": [
			{
			  "name": "renditions.fileextension",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"contentformat": {
		  "fields": [
			{
			  "name": "renditions.format",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"contenthref": {
		  "fields": [
			{
			  "name": "renditions.href",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"contentmimetype": {
		  "fields": [
			{
			  "name": "renditions.mimetype",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"contentsizeinbytes": {
		  "fields": [
			{
			  "name": "renditions.sizeinbytes",
			  "type": "keyword"
			}
		  ]
		},
		"copyrightdate": {
		  "fields": [
			{
			  "name": "copyrightdate",
			  "type": "keyword"
			}
		  ]
		},
		"copyrightholder": {
		  "fields": [
			{
			  "name": "copyrightholder",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"copyrightline": {
		  "fields": [
			{
			  "name": "copyrightnotice",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"copyspace": {
		  "fields": [
			{
			  "name": "renditions.copyspace",
			  "type": "text"
			}
		  ]
		},
		"corporategroupid": {
		  "fields": [
			{
			  "name": "usagerights.groups.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ],
		  "and": {
			"path": "usagerights.groups",
			"query": {
			  "term": {
				"usagerights.groups.type": "corporate"
			  }
			}
		  }
		},
		"corporategroupname": {
		  "fields": [
			{
			  "name": "usagerights.groups.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ],
		  "and": {
			"path": "usagerights.groups",
			"query": {
			  "term": {
				"usagerights.groups.type": "corporate"
			  }
			}
		  }
		},
		"country": {
		  "fields": [
			{
			  "name": "datelinelocation.countrycode",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"countryname": {
		  "fields": [
			{
			  "name": "datelinelocation.countryname",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"creationdate": {
		  "fields": [
			{
			  "name": "firstcreated",
			  "type": "date"
			}
		  ]
		},
		"creationdatelower": {
		  "fields": [
			{
			  "name": "firstcreated",
			  "type": "date"
			}
		  ]
		},
		"creationdateupper": {
		  "fields": [
			{
			  "name": "firstcreated",
			  "type": "date"
			}
		  ]
		},
		"creationday": {
		  "fields": [
			{
			  "name": "firstcreated",
			  "type": "integer"
			}
		  ]
		},
		"creationmonth": {
		  "fields": [
			{
			  "name": "firstcreated",
			  "type": "integer"
			}
		  ]
		},
		"creationyear": {
		  "fields": [
			{
			  "name": "firstcreated",
			  "type": "integer"
			}
		  ]
		},
		"creator": {
		  "fields": [
			{
			  "name": "creator",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"creditline": {
		  "fields": [
			{
			  "name": "creditline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"creditlinename": {
		  "fields": [
			{
			  "name": "creditlineid",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"cycle": {
		  "fields": [
			{
			  "name": "cycle",
			  "type": "text"
			}
		  ]
		},
		"dateline": {
		  "fields": [
			{
			  "name": "dateline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"datedwithin": {
		  "fields": [
			{
			  "name": "firstcreated",
			  "type": "date"
			}
		  ]
		},
		"defaultlanguage": {
		  "fields": [
			{
			  "name": "language",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"description": {
		  "fields": [
			{
			  "name": "descriptions",
			  "type": "text"
			}
		  ]
		},
		"distributionchannel": {
		  "fields": [
			{
			  "name": "distributionchannels",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"distributionscope": {
		  "fields": [
			{
			  "name": "filings.distributionscope",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"editorialid": {
		  "fields": [
			{
			  "name": "editorialid",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"editorialpriority": {
		  "fields": [
			{
			  "name": "editorialpriority",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"editorialslug": {
		  "fields": [
			{
			  "name": "filings.slugline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"editorialsubject": {
		  "fields": [
			{
			  "name": "subjects.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ],
		  "and": {
			"path": "subjects",
			"query": {
			  "term": {
				"subjects.creator": "editorial"
			  }
			}
		  }
		},
		"editorialtype": {
		  "fields": [
			{
			  "name": "editorialtypes",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"expandedfilingformat": {
		  "fields": [
			{
			  "name": "filings.format",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"expandedgroupsidadd": {
		  "fields": [
			{
			  "name": "filings.routings.expandedgroupsidadds",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"expandedgroupsidout": {
		  "fields": [
			{
			  "name": "filings.routings.expandedgroupsidouts",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"expandedsidadd": {
		  "fields": [
			{
			  "name": "filings.routings.expandedsidadds",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"expandedsidout": {
		  "fields": [
			{
			  "name": "filings.routings.expandedsidouts",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"explicitwarning": {
		  "replace": {
			"term": {
			  "signals": "explicitcontent"
			}
		  }
		},
		"filingcategory": {
		  "fields": [
			{
			  "name": "filings.filingcategory",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingcategoryencoded": {
		  "fields": [
			{
			  "name": "filings.filingcategory",
			  "type": "keyword",
			  "transform": [
				"decode",
				"lowercase"
			  ]
			}
		  ]
		},
		"filingcountry": {
		  "fields": [
			{
			  "name": "filings.filingcountries",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingformat": {
		  "fields": [
			{
			  "name": "filings.format",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingonlinecode": {
		  "fields": [
			{
			  "name": "filings.filingonlinecode",
			  "type": "keyword"
			}
		  ]
		},
		"filingregion": {
		  "fields": [
			{
			  "name": "filings.filingregions",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingslug": {
		  "fields": [
			{
			  "name": "filings.slugline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingsource": {
		  "fields": [
			{
			  "name": "filings.filingsource",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingstyle": {
		  "fields": [
			{
			  "name": "filings.filingstyle",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingsubject": {
		  "fields": [
			{
			  "name": "filings.filingsubjects",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingsubsubject": {
		  "fields": [
			{
			  "name": "filings.filingsubjects",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingtopic": {
		  "fields": [
			{
			  "name": "filings.filingtopics",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"filingtype": {
		  "fields": [
			{
			  "name": "filingtype",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"fixture": {
		  "fields": [
			{
			  "name": "fixture.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"fixtureid": {
		  "fields": [
			{
			  "name": "fixture.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"fixturename": {
		  "fields": [
			{
			  "name": "fixture.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"friendlykey": {
		  "fields": [
			{
			  "name": "friendlykey",
			  "type": "keyword"
			}
		  ]
		},
		"function": {
		  "fields": [
			{
			  "name": "function",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"geography": {
		  "fields": [
			{
			  "name": "usagerights.geography",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"groupsidadd": {
		  "fields": [
			{
			  "name": "filings.routings.groupsidadds",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"groupsidout": {
		  "fields": [
			{
			  "name": "filings.routings.groupsidouts",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"headline": {
		  "fields": [
			{
			  "name": "headline",
			  "type": "text"
			}
		  ]
		},
		"height": {
		  "fields": [
			{
			  "name": "renditions.height",
			  "type": "keyword"
			}
		  ]
		},
		"hometownstate": {
		  "fields": [
			{
			  "name": "persons.associatedstates.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"hue": {
		  "fields": [
			{
			  "name": "renditions.hue",
			  "type": "text"
			}
		  ]
		},
		"junkline": {
		  "fields": [
			{
			  "name": "filings.junkline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"importfolder": {
		  "fields": [
			{
			  "name": "filings.importfolder",
			  "type": "keyword"
			}
		  ]
		},
		"importwarnings": {
		  "fields": [
			{
			  "name": "filings.importwarnings",
			  "type": "keyword"
			}
		  ]
		},
		"indcode": {
		  "fields": [
			{
			  "name": "companies.industries.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"inpackage": {
		  "fields": [
			{
			  "name": "inpackages",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"isdigitized": {
		  "replace": {
			"bool": {
			  "must_not": {
				"term": {
				  "signals": "isnotdigitized"
				}
			  }
			}
		  }
		},
		"itemcontenttype": {
		  "fields": [
			{
			  "name": "itemcontenttype.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"itemid": {
		  "fields": [
			{
			  "name": "itemid",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"keyword": {
		  "fields": [
			{
			  "name": "keywordlines",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"keywordline": {
		  "fields": [
			{
			  "name": "keywordlines",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"keywords": {
		  "fields": [
			{
			  "name": "keywordlines",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"libraryrequestid": {
		  "fields": [
			{
			  "name": "filings.libraryrequestid",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"libraryrequestlogin": {
		  "fields": [
			{
			  "name": "filings.libraryrequestlogin",
			  "type": "keyword"
			}
		  ]
		},
		"librarytwincheck": {
		  "fields": [
			{
			  "name": "filings.librarytwincheck",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"limitations": {
		  "fields": [
			{
			  "name": "usagerights.limitations",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"mediatype": {
		  "fields": [
			{
			  "name": "type",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"locationline": {
		  "fields": [
			{
			  "name": "locationline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"nameline": {
		  "fields": [
			{
			  "name": "persons.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ],
		  "and": {
			"path": "persons",
			"query": {
			  "term": {
				"persons.rels": "personfeatured"
			  }
			}
		  }
		},
		"newspowerdrivetimealaska": {
		  "fields": [
			{
			  "name": "newspowerdrivetimealaska",
			  "type": "boolean"
			}
		  ]
		},
		"newspowerdrivetimearizona": {
		  "fields": [
			{
			  "name": "newspowerdrivetimearizona",
			  "type": "boolean"
			}
		  ]
		},
		"newspowerdrivetimeatlantic": {
		  "fields": [
			{
			  "name": "newspowerdrivetimeatlantic",
			  "type": "boolean"
			}
		  ]
		},
		"newspowerdrivetimecentral": {
		  "fields": [
			{
			  "name": "newspowerdrivetimecentral",
			  "type": "boolean"
			}
		  ]
		},
		"newspowerdrivetimeeastern": {
		  "fields": [
			{
			  "name": "newspowerdrivetimeeastern",
			  "type": "boolean"
			}
		  ]
		},
		"newspowerdrivetimehawaii": {
		  "fields": [
			{
			  "name": "newspowerdrivetimehawaii",
			  "type": "boolean"
			}
		  ]
		},
		"newspowerdrivetimemountain": {
		  "fields": [
			{
			  "name": "newspowerdrivetimemountain",
			  "type": "boolean"
			}
		  ]
		},
		"newspowerdrivetimepacific": {
		  "fields": [
			{
			  "name": "newspowerdrivetimepacific",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorningalaska": {
		  "fields": [
			{
			  "name": "newspowermorningalaska",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorningarizona": {
		  "fields": [
			{
			  "name": "newspowermorningarizona",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorningatlantic": {
		  "fields": [
			{
			  "name": "newspowermorningatlantic",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorningcentral": {
		  "fields": [
			{
			  "name": "newspowermorningcentral",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorningeastern": {
		  "fields": [
			{
			  "name": "newspowermorningeastern",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorninghawaii": {
		  "fields": [
			{
			  "name": "newspowermorninghawaii",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorningmountain": {
		  "fields": [
			{
			  "name": "newspowermorningmountain",
			  "type": "boolean"
			}
		  ]
		},
		"newspowermorningpacific": {
		  "fields": [
			{
			  "name": "newspowermorningpacific",
			  "type": "boolean"
			}
		  ]
		},
		"normalizedslug": {
		  "fields": [
			{
			  "name": "filings.slugline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"onlinecode": {
		  "fields": [
			{
			  "name": "filings.filingonlinecode",
			  "type": "keyword"
			}
		  ]
		},
		"originalmediaid": {
		  "fields": [
			{
			  "name": "originalmediaid",
			  "type": "text"
			}
		  ]
		},
		"outinginstructions": {
		  "fields": [
			{
			  "name": "outinginstructions",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"party": {
		  "fields": [
			{
			  "name": "persons.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"partytype": {
		  "fields": [
			{
			  "name": "persons.types",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"passcode": {
		  "fields": [
			{
			  "name": "filings.routings.passcodeadds",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"passcodeout": {
		  "fields": [
			{
			  "name": "filings.routings.passcodeouts",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"person": {
		  "fields": [
			{
			  "name": "persons.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"person_name": {
		  "fields": [
			{
			  "name": "persons.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"primaryticker": {
		  "fields": [
			{
			  "name": "companies.symbols.ticker",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"priority": {
		  "fields": [
			{
			  "name": "priority",
			  "type": "integer"
			}
		  ]
		},
		"productid": {
		  "fields": [
			{
			  "name": "filings.products",
			  "type": "integer"
			}
		  ]
		},
		"productsource": {
		  "fields": [
			{
			  "name": "productsources",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"profcompany": {
		  "fields": [
			{
			  "name": "companies.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"profcompanyindustry": {
		  "fields": [
			{
			  "name": "companies.industries.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"profcompanyticker": {
		  "fields": [
			{
			  "name": "companies.symbols.ticker",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"profcountry": {
		  "fields": [
			{
			  "name": "places.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"provider": {
		  "fields": [
			{
			  "name": "provider.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"providerid": {
		  "fields": [
			{
			  "name": "provider.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"providername": {
		  "fields": [
			{
			  "name": "provider.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"providertype": {
		  "fields": [
			{
			  "name": "provider.type",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"ratingvalue": {
		  "fields": [
			{
			  "name": "ratings.rating",
			  "type": "keyword"
			}
		  ]
		},
		"reach": {
		  "fields": [
			{
			  "name": "signals",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"recordid": {
		  "fields": [
			{
			  "name": "recordid",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"recordtype": {
		  "fields": [
			{
			  "name": "recordtype",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"region": {
		  "fields": [
			{
			  "name": "places.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"releasedatetime": {
		  "fields": [
			{
			  "name": "releasedatetime",
			  "type": "date"
			}
		  ]
		},
		"resolution": {
		  "fields": [
			{
			  "name": "renditions.resolution",
			  "type": "keyword"
			}
		  ]
		},
		"resolutionunits": {
		  "fields": [
			{
			  "name": "renditions.resolutionunit",
			  "type": "text"
			}
		  ]
		},
		"rightsline": {
		  "fields": [
			{
			  "name": "rightsline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"salesname": {
		  "fields": [
			{
			  "name": "services.apsales",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"scene": {
		  "fields": [
			{
			  "name": "renditions.scene",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"scenevalue": {
		  "fields": [
			{
			  "name": "renditions.scene",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"selector": {
		  "fields": [
			{
			  "name": "filings.selector",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"seriesline": {
		  "fields": [
			{
			  "name": "seriesline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"sidadd": {
		  "fields": [
			{
			  "name": "filings.routings.sidadds",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"sidout": {
		  "fields": [
			{
			  "name": "filings.routings.sidouts",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"signal": {
		  "fields": [
			{
			  "name": "signals",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"slug": {
		  "fields": [
			{
			  "name": "filings.slugline",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"source": {
		  "fields": [
			{
			  "name": "sources.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"sourcecountryarea": {
		  "fields": [
			{
			  "name": "sources.countryarea",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"sourceid": {
		  "fields": [
			{
			  "name": "sources.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"sourcesubtype": {
		  "fields": [
			{
			  "name": "sources.subtype",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"sourcetype": {
		  "fields": [
			{
			  "name": "sources.type",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"specialfieldattn": {
		  "fields": [
			{
			  "name": "filings.specialfieldattn",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"specialinstructions": {
		  "fields": [
			{
			  "name": "specialinstructions",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"state": {
		  "fields": [
			{
			  "name": "datelinelocation.countryareacode",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"statename": {
		  "fields": [
			{
			  "name": "datelinelocation.countryareaname",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"status": {
		  "fields": [
			{
			  "name": "pubstatus",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"subject": {
		  "fields": [
			{
			  "name": "subjects.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"subjectid": {
		  "fields": [
			{
			  "name": "subjects.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"subjectname": {
		  "fields": [
			{
			  "name": "subjects.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"subjectvalue": {
		  "fields": [
			{
			  "name": "subjects.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"suppcat": {
		  "fields": [
			{
			  "name": "suppcategories.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"suppcatencoded": {
		  "fields": [
			{
			  "name": "suppcategories.code",
			  "type": "keyword",
			  "transform": [
				"decode",
				"lowercase"
			  ]
			}
		  ]
		},
		"suppcatname": {
		  "fields": [
			{
			  "name": "suppcategories.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"teragramsubject": {
		  "fields": [
			{
			  "name": "subjects.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ],
		  "and": {
			"path": "subjects",
			"query": {
			  "term": {
				"subjects.creator": "teragram"
			  }
			}
		  }
		},
		"ticker": {
		  "fields": [
			{
			  "name": "companies.symbols.ticker",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"title": {
		  "fields": [
			{
			  "name": "title",
			  "type": "text"
			}
		  ]
		},
		"tpmid": {
		  "fields": [
			{
			  "name": "thirdpartymeta.code",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"tpmname": {
		  "fields": [
			{
			  "name": "thirdpartymeta.name",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"tpmvocabulary": {
		  "fields": [
			{
			  "name": "thirdpartymeta.vocabulary",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"tpmvocabularyowner": {
		  "fields": [
			{
			  "name": "thirdpartymeta.vocabularyowner",
			  "type": "keyword"
			}
		  ]
		},
		"transmissionfilename": {
		  "fields": [
			{
			  "name": "filings.transmissionfilename",
			  "type": "keyword"
			}
		  ]
		},
		"transmissionsource": {
		  "fields": [
			{
			  "name": "transmissionsources",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"transref": {
		  "fields": [
			{
			  "name": "filings.transmissionreference",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"transrefencoded": {
		  "fields": [
			{
			  "name": "filings.transmissionreference",
			  "type": "keyword",
			  "transform": [
				"decode"
			  ]
			}
		  ]
		},
		"usagetype": {
		  "fields": [
			{
			  "name": "usagerights.usagetype",
			  "type": "text",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"words": {
		  "fields": [
			{
			  "name": "caption.words",
			  "type": "integer"
			},
			{
			  "name": "script.words",
			  "type": "integer"
			},
			{
			  "name": "shotlist.words",
			  "type": "integer"
			},
			{
			  "name": "main.words",
			  "type": "integer"
			}
		  ]
		},
		"workflowstatus": {
		  "fields": [
			{
			  "name": "workflowstatus",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		},
		"workgroup": {
		  "fields": [
			{
			  "name": "workgroup",
			  "type": "keyword",
			  "transform": [
				"lowercase"
			  ]
			}
		  ]
		}
	  }
`
	jo, err := json.ParseJSONObject(s)
	if err != nil {
		t.Error(err.Error())
		return
	}

	var (
		sb strings.Builder
	)

	sb.WriteString("func getFields() map[string]field {\n")
	sb.WriteString("fs := make(map[string]field)\n")

	for _, k := range jo.Names() {
		sb.WriteString("\nfs[\"")
		sb.WriteString(k)
		sb.WriteString("\"]=field{\nFields:[]es{")

		o, _ := jo.GetObject(k)
		ja, _ := o.GetArray("fields")
		fs, _ := ja.GetObjects()
		for _, f := range fs {
			nm, _ := f.GetString("name")
			tp, _ := f.GetString("type")
			tr := "nil"
			ja, err := f.GetArray("transform")
			if err == nil && ja.Length() > 0 {
				var b strings.Builder
				b.WriteString("[]string{")
				ss, _ := ja.GetStrings()
				for i, t := range ss {
					if i > 0 {
						b.WriteString(", ")
					}
					b.WriteString("\"")
					b.WriteString(t)
					b.WriteString("\"")
				}
				b.WriteString("}")
				tr = b.String()
			}

			sb.WriteString("\nnewEs(\"")
			sb.WriteString(nm)
			sb.WriteString("\", \"")
			sb.WriteString(tp)
			sb.WriteString("\", ")
			sb.WriteString(tr)
			sb.WriteString("),")
		}
		sb.WriteString("\n}")

		and, err := o.GetObject("and")
		if err == nil {
			p, _ := and.GetString("path")
			q, _ := and.GetObject("query")
			t, _ := q.GetObject("term")
			f := t.Names()[0]
			v, _ := t.GetString(f)

			sb.WriteString(",\nAnd:and{\nPath:\"")
			sb.WriteString(p)
			sb.WriteString("\",\nField:\"")
			sb.WriteString(f)
			sb.WriteString("\",\nValue:\"")
			sb.WriteString(v)
			sb.WriteString("\",\n}")
		}

		sb.WriteString("}\n")
	}

	sb.WriteString("\nreturn fs\n}")

	fmt.Println(sb.String())
}
