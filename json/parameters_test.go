package json

import (
	"fmt"
	"testing"
)

func TestJSONObjectWithParametersParsing(t *testing.T) {
	input := `{"id":"${id}","name":"${name}","params":{"id":1,"name":"YM"}}`
	expected := `{"id":1,"name":"YM"}`
	jo, err := ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params := jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test := jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}", "name":"${name}","params":{"name":"YM"}}`
	expected = `{"name":"YM"}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}", "name":"${name}","products":["${p1}","${p2}"]}`
	expected = `{"id":1,"name":"YM","products":[1,2]}`
	jo, err = ParseObject([]byte(input))
	fmt.Println(jo.String())
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}

	tmp, _ := ParseObject([]byte(`{"id":1,"name":"YM","index":"appl","size":20,"p1":1,"p2":2}`))
	params = tmp.Map()
	jo.SetParams(params, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}","name":"${name}", "child":{"name":"${name}", "age":"${age}","${extra_field}":"${extra_value}"}, "params":{"id":1,"name":"YM","age":13,"extra_field":"nick","extra_value":"Gusyonok"}}`
	expected = `{"id":1,"name":"YM","child":{"name":"YM","age":13,"nick":"Gusyonok"}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}","name":"${name}", "child":{"age":"${age}"}, "params":{"id":1,"name":"YM"}}`
	expected = `{"id":1,"name":"YM"}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"${id_prefix}_id":"${id}","name":"${name}", "child":{"name":"${name} Jr.", "age":"${age}"}, "params":{"id_prefix":"user","id":1,"name":"YM"}}`
	expected = `{"user_id":1,"name":"YM","child":{"name":"YM Jr."}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"${id_prefix}_id":"${id}","name":"${name}", "child":{"name":"${name} Jr.", "age":"${age}"}, "params":{"id_prefix1":"user","id":1,"name":"YM"}}`
	expected = `{"_id":1,"name":"YM","child":{"name":"YM Jr."}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"template":{"query":{"query_string":{"query":"${query}","fields":"${fields}"}}},"params":{"query":"test","fields":["head","body"]}}`
	expected = `{"template":{"query":{"query_string":{"query":"test","fields":["head","body"]}}}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"template":{"_source":["headline","${field1}","${field2}","${field3}"]},"params":{"field1":"type","field2":"date"}}`
	expected = `{"template":{"_source":["headline","type","date"]}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"first_name":"${first_name?Yuri}","last_name":"${last_name?Metelkine}","params":{"last_name":"Metelkin"}}`
	expected = `{"first_name":"Yuri","last_name":"Metelkin"}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"name":"${first_name?Yuri} ${last_name?Metelkine}","params":{"last_name":"Metelkin"}}`
	expected = `{"name":"Yuri Metelkin"}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"template":{"_source":["headline","${field1}","${field2}","${field3?test}"]},"params":{"field1":"type","field2":"date"}}`
	expected = `{"template":{"_source":["headline","type","date","test"]}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	params = jo.GetParams()
	for k, v := range params {
		s, _ := v.String()
		fmt.Printf("%s\t%s\n", k, s)
	}
	jo.SetParams(nil, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"bool":{"must":{"query_string":{"query":"${query}","fields":["head","body"]}},"filter":[{"terms":{"type":"${media_types}"}},{"terms":{"filing.products":"${include_products}"}}],"must_not":{"terms":{"filing.products":"${exclude_products}"}}}}}`
	expected = `{"query":{"bool":{"must":{"query_string":{"query":"ap","fields":["head","body"]}},"filter":[{"terms":{"type":["audio"]}},{"terms":{"filing.products":[1,2]}}],"must_not":{"terms":{"filing.products":[3]}}}}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	tmp, _ = ParseObject([]byte(`{"query":"ap","media_types":["audio"],"include_products":[1,2], "exclude_products":[3]}`))
	params = tmp.Map()
	jo.SetParams(params, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"bool":{"must":{"query_string":{"query":"${query}","fields":["head","body"]}},"filter":[{"terms":{"type":"${media_types}"}},{"terms":{"filing.products":"${include_products}"}}],"must_not":{"terms":{"filing.products":"${exclude_products}"}}}}}`
	expected = `{"query":{"bool":{"must":{"query_string":{"query":"ap","fields":["head","body"]}},"filter":[{"terms":{"type":["audio"]}}]}}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	tmp, _ = ParseObject([]byte(`{"query":"ap","media_types":["audio"]}`))
	params = tmp.Map()
	jo.SetParams(params, nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"bool":{"must":{"query_string":{"query":"${query}","fields":["head","body"]}},"filter":[{"terms":{"type":"${media_types}"}},{"terms":{"filing.products":"${include_products}"}}],"must_not":{"terms":{"filing.products":"${exclude_products}"}}}}}`
	expected = `{"query":{"bool":{"filter":[{"terms":{"type":["audio"]}}]}}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	tmp, _ = ParseObject([]byte(`{"media_types":["audio"]}`))
	params = tmp.Map()
	empty := map[string][]string{"query_string": []string{"query"}}
	jo.SetParams(params, empty)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"query_string":{"query":"${query}","fields":["head","body"]}}}`
	expected = `{"query":{"query_string":{"query":"head:putin","fields":["head","body"]}}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	tmp, _ = ParseObject([]byte(`{"query":"head:putin"}`))
	params = tmp.Map()
	empty = map[string][]string{"query_string": []string{"query"}}
	jo.SetParams(params, empty)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"query_string":{"query":"${query}","fields":"${fields}"}}}`
	expected = `{"query":{"query_string":{"query":"title:putin","fields":["body"]}}}`
	jo, err = ParseObject([]byte(input))
	if err != nil {
		t.Error(err.Error())
	}
	tmp, _ = ParseObject([]byte(`{"query":"title:putin","fields":["body"]}`))
	params = tmp.Map()
	empty = map[string][]string{"query_string": []string{"query"}}
	copy := jo.Copy()
	copy.SetParams(params, empty)
	test = copy.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}
}
