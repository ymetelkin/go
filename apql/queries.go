package apql

import (
	"strings"

	"github.com/ymetelkin/go/json"
)

func textQueryJSON(field string, value string, keyword bool, phrase bool) json.Object {
	if value == "*" {
		return existsQueryJSON(field)
	}

	text := json.Object{}
	text.AddString(field, value)
	jo := json.Object{}
	if keyword {
		jo.AddObject("term", text)
	} else if phrase {
		jo.AddObject("match_phrase", text)
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

func existsQueryJSON(field string) json.Object {
	f := json.Object{}
	f.AddString("field", field)
	jo := json.Object{}
	jo.AddObject("exists", f)
	return jo
}
