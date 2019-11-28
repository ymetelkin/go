package json

import (
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	s := `{"id":1,"code":"YM","name":"\"Yuri Metelkin\"", "cool":true, "obj":{"a":"b"}}`
	jo, err := ParseObjectString(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		for key, value := range jo.Properties {
			fmt.Printf("%s: %#v\n", key, value)
			jv, ok := jo.getValue(key)
			if !ok {
				t.Error("Failrd to get value")
			} else {
				fmt.Printf("%s: %v\n", key, jv.Value)
			}
		}

		fmt.Printf("%s\n", jo.String())

		jo.SetString("name", "SV")
		jo.SetInt("id", 2)
		fmt.Printf("%s\n", jo.String())

		jo.Remove("name")
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"object":{},"array":[]}`
	jo, err = ParseObjectString(s)
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
	jo, err = ParseObjectString(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}

	s = `{"test":3140000000000}`
	jo, err = ParseObjectString(s)
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
	jo, err = ParseObjectString(s)
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
	jo, err = ParseObjectString(s)
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
	jo, err = ParseObjectString(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		test := jo.InlineString()
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}
}
