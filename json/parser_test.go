package json

import (
	"fmt"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	s := `{"id":1,"code":"YM","name":"\"Yuri Metelkin\"", "cool":true, "obj":{"a":"b"}}`
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
		if test != s {
			t.Error("Parsing failed")
		}
		fmt.Printf("%s\n", jo.String())
	}
}
