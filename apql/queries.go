package apql

import (
	"strings"

	"github.com/ymetelkin/go/json"
)

func textQueryJSON(field string, value string, keyword bool) json.Object {
	text := json.Object{}
	text.AddString(field, value)
	jo := json.Object{}
	if keyword {
		jo.AddObject("term", text)
	} else {
		jo.AddObject("match", text)
	}
	return jo
}

func intQueryJSON(field string, value int) json.Object {
	term := json.Object{}
	term.AddInt(field, value)
	jo := json.Object{}
	jo.AddObject("term", term)
	return jo
}

func boolQueryJSON(field string, value bool) json.Object {
	term := json.Object{}
	term.AddBool(field, value)
	jo := json.Object{}
	jo.AddObject("term", term)
	return jo
}

func not(query json.Object) json.Object {
	not := json.Object{}
	not.AddObject("must_not", query)
	jo := json.Object{}
	jo.AddObject("bool", not)
	return jo
}

func textsQueryJSON(field string, values []string, keyword bool) json.Object {
	jo := json.Object{}

	if keyword {
		ja := json.Array{}
		ja.AddStrings(values)
		terms := json.Object{}
		terms.AddArray(field, ja)
		jo.AddObject("terms", terms)
	} else {
		v := strings.Join(values, " ")
		match := json.Object{}
		match.AddString(field, v)
		jo.AddObject("match", match)
	}

	return jo
}

func intsQueryJSON(field string, values []int) json.Object {
	ja := json.Array{}
	ja.AddInts(values)
	term := json.Object{}
	term.AddArray(field, ja)
	jo := json.Object{}
	jo.AddObject("terms", term)
	return jo
}
