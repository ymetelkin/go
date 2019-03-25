package json

import (
	"fmt"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	s := `{"id":1,"name":"YM", "cool":true, "obj":{"a":"b"}}`
	jo, err := ParseJsonObject(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		for key, value := range jo.Properties {
			fmt.Printf("%s: %#v\n", key, value)
			jv, err := jo.GetValue(key)
			if err != nil {
				t.Error(err.Error())
			} else {
				fmt.Printf("%s: %v\n", key, jv.Value)
			}
		}

		fmt.Printf("%s\n", jo.ToString())

		jo.Set("name", "SV")
		jo.Set("id", 2)
		fmt.Printf("%s\n", jo.ToString())

		jo.Remove("name")
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"test":3.14E+12}`
	jo, err = ParseJsonObject(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.ToInlineString()
		if test != strings.ToLower(s) {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"test":3140000000000}`
	jo, err = ParseJsonObject(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.ToInlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"id":1,"name":"YM","success":true,"grades":[{"subject":"Math","grade":5},{"subject":"English","grade":3.74},5,3140000000000,"xyz"],"params":{"query":"test","size":100}}`
	jo, err = ParseJsonObject(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.ToInlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"query":{"bool":{"must":{"match":{"headline":"test"}},"filter":[{"term":{"type":"text"}},{"terms":{"filings.products":[1,2,3]}}]}},"size":100}`
	jo, err = ParseJsonObject(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.ToInlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.ToString())
	}
}
