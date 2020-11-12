package json

import (
	"fmt"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	s := `{"id":1,"code":"YM","name":"\"Yuri Metelkin\"", "cool":true, "obj":{"a":"b"}}`
	s = `{"took" : 12,"timed_out":false,"_shards":{"total":1,"successful":1,"skipped":0,"failed":0},"hits":{"total":38021,"max_score":null,"hits":[{"_index":"pronto-live-versions-index","_type":"doc","_id":"01b832e61a4499d04e3f197c13ada64f.-1.0","_score":null,"_source":{"tagsToStringNullable":null,"publishApiKeyWithVersion":"01b832e61a4499d04e3f197c13ada64f/services/pronto/01b832e61a4499d04e3f197c13ada64f.0.json?versionId=dvplFCzXL2goIpQSOCXfu90E4x9wz5PG","tagsToString":"null","itemTagsToString":"lrathke,pwhittle","notExcludedProducts":"20","SqsMessageId":"534723d0-b85c-46d9-a704-fbd7fee5720e","SqsMessageGroudId":"01b832e61a4499d04e3f197c13ada64f","SqsReceiveCount":"1","SqsReceiveHandle":"AQEBv75nzo1Ur55ymujPZvPZeSslrK0cKtEWPP1VInVIJdf9i/zcPrRc9TD23nI4vM5Sb4Nra3nXSUHs/ePMusV7d2KDnLNjeWPXFizP/LaubVTTl85qxhqImiz8WsTInnD8ojW7xO/TzkXkSwcWntMLecWKmHX5sZQ5oePQZYH23TZhwhrU32jo4sIL1lEzGKJhm96Dd4YhdUAZTOCcNkZLykf2PNd04Goed154W2g4ou96Vz8VEGiDHpTd+yMcmCTL1/oL9GnebMoA5FPI45y6NUFTFrKu4HrdHv+jSIHjt/4v3Nh70/6vHRwmf6EbIqvrisxX687k810aqBK+x5VZxg==","TraceId":"e48390f9c34f4f6bb3e3f5629e5b25b9","Size":5841,"LastModified":"2020-10-14T12:30:54Z","Version":0,"Generator":"pronto-prepublish","Products":"20","Url":"01b832e61a4499d04e3f197c13ada64f/services/pronto/01b832e61a4499d04e3f197c13ada64f.0.json","ETag":"\"667f2d104360d94e514920e7dae43722\"","Updated":"2020-10-14T12:30:53Z","ItemId":"01b832e61a4499d04e3f197c13ada64f","RSN":0,"ItemTags":["lrathke","pwhittle"],"ItemVersion":0,"ArrivalDateTime":"2020-10-14T12:30:53Z","ArrivalDateTimeSetByDb":false,"IsDelete":false,"Role":"pronto-prepublish","PublishApiKey":"01b832e61a4499d04e3f197c13ada64f/services/pronto/01b832e61a4499d04e3f197c13ada64f.0.json","EpochMilliSecondsTimestamp":1602678654500,"Tags":[],"AggEpochMilliSecondsTimestamp":1602678654500,"AppVersion":-1,"S3VersionId":"dvplFCzXL2goIpQSOCXfu90E4x9wz5PG","IsPublished":false,"IgnoreStatus":"None","OriginalPublishApiKey":"01b832e61a4499d04e3f197c13ada64f.-1.0.json","IsLatest":true},"sort":["01b832e61a4499d04e3f197c13ada64f.-1.0"]}]}}`
	//s = `{"a":{"b":null,"c":1},"d":2}`
	jo, err := ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		for _, jp := range jo.Properties {
			fmt.Printf("%s: %#v\n", jp.Name, jp.Value)
			_, ok := jo.GetValue(jp.Name)
			if !ok {
				t.Error("Failed to get value")
			}
		}

		fmt.Printf("%s\n", jo.String())

		jo.Set("name", String("SV"))
		jo.Set("id", Int(2))
		fmt.Printf("%s\n", jo.String())

		jo.Remove("name")
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"object":{},"array":[]}`
	jo, err = ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"test":3.14E+12}`
	jo, err = ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s && test != strings.ToLower(s) {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"test":3140000000000}`
	jo, err = ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"id":1,"name":"YM","success":true,"grades":[{"subject":"Math","grade":5},{"subject":"English","grade":3.74},5,3140000000000,"xyz"],"params":{"query":"test","size":100}}`
	jo, err = ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"query":{"bool":{"must":{"match":{"headline":"test"}},"filter":[{"term":{"type":"text"}},{"terms":{"filings.products":[1,2,3]}}]}},"size":100}`
	jo, err = ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"test":"APGBL\\dzelio"}`
	jo, err = ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"a":[null]}`
	jo, err = ParseObject([]byte(s))
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != `{"a":[]}` {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}
}
