package json

import (
	"fmt"
	"testing"
)

/*
func TestPropertyNameWithParametersParsing(t *testing.T) {
	s := `{"${f}":1}`
	runes := []rune(s)
	size := len(runes)

	ps, _, err := parsePropertyNameWithParameters(runes, size, 1)
	if err != nil {
		t.Error(err.Error())
	} else if !ps.IsParameterized {
		t.Error("Must be parameterized")
	} else {
		p := ps.Parameters[0]
		fmt.Printf("%s\t%s\n", ps.Value, p.Name)
		if p.Name != "f" {
			t.Error("Must have name [f]")
		}
	}

	input = `{"${f?id}":1}`
	runes = []rune(s)
	size = len(runes)

	ps, _, err = parsePropertyNameWithParameters(runes, size, 1)
	if err != nil {
		t.Error(err.Error())
	} else if !ps.IsParameterized {
		t.Error("Must be parameterized")
	} else {
		p := ps.Parameters[0]
		fmt.Printf("%s\t%s (%s)\n", ps.Value, p.Name, p.Default)
		if p.Name != "f" {
			t.Error("Must have name [f]")
		}
		if p.Default != "id" {
			t.Error("Must have default [id]")
		}
	}

	input = `{"${prefix?user}_${suffix?id}":1}`
	runes = []rune(s)
	size = len(runes)

	ps, _, err = parsePropertyNameWithParameters(runes, size, 1)
	if err != nil {
		t.Error(err.Error())
	} else if !ps.IsParameterized {
		t.Error("Must be parameterized")
	} else if len(ps.Parameters) != 2 {
		t.Error("Must have 2 parameters")
	} else {
		p1 := ps.Parameters[0]
		p2 := ps.Parameters[1]
		fmt.Printf("%s\t%s (%s)\t%s (%s)\n", ps.Value, p1.Name, p1.Default, p2.Name, p2.Default)
		if p1.Name != "prefix" {
			t.Error("Must have name [prefix]")
		}
		if p1.Default != "user" {
			t.Error("Must have default [user]")
		}
		if p2.Name != "suffix" {
			t.Error("Must have name [suffix]")
		}
		if p2.Default != "id" {
			t.Error("Must have default [id]")
		}
	}
}

func TestValueWithParametersParsing(t *testing.T) {
	s := `"${v}"`
	runes := []rune(s)
	size := len(runes)

	ps, _, err := parseTextValueWithParameters(runes, size, 1)
	if err != nil {
		t.Error(err.Error())
	} else if !ps.IsParameterized {
		t.Error("Must be parameterized")
	} else {
		p := ps.Parameters[0]
		fmt.Printf("%s\t%s\n", ps.Value, p.Name)
		if p.Name != "v" {
			t.Error("Must have name [v]")
		}
	}

	input = `"${v?test}"`
	runes = []rune(s)
	size = len(runes)

	ps, _, err = parseTextValueWithParameters(runes, size, 1)
	if err != nil {
		t.Error(err.Error())
	} else if !ps.IsParameterized {
		t.Error("Must be parameterized")
	} else {
		p := ps.Parameters[0]
		fmt.Printf("%s\t%s (%s)\n", ps.Value, p.Name, p.Default)
		if p.Name != "v" {
			t.Error("Must have name [v]")
		}
		if p.Default != "test" {
			t.Error("Must have default [test]")
		}
	}

	input = `"prefix_${v?test}_suffix"`
	runes = []rune(s)
	size = len(runes)

	ps, _, err = parseTextValueWithParameters(runes, size, 1)
	if err != nil {
		t.Error(err.Error())
	} else if !ps.IsParameterized {
		t.Error("Must be parameterized")
	} else {
		p := ps.Parameters[0]
		fmt.Printf("%s\t%s (%s)\n", ps.Value, p.Name, p.Default)
		if p.Name != "v" {
			t.Error("Must have name [v]")
		}
		if p.Default != "test" {
			t.Error("Must have default [test]")
		}
	}

	input = `"This is ${prefix?user} ${suffix?id} xyz"`
	runes = []rune(s)
	size = len(runes)

	ps, _, err = parseTextValueWithParameters(runes, size, 1)
	if err != nil {
		t.Error(err.Error())
	} else if !ps.IsParameterized {
		t.Error("Must be parameterized")
	} else if len(ps.Parameters) != 2 {
		t.Error("Must have 2 parameters")
	} else {
		p1 := ps.Parameters[0]
		p2 := ps.Parameters[1]
		fmt.Printf("%s\t%s (%s)\t%s (%s)\n", ps.Value, p1.Name, p1.Default, p2.Name, p2.Default)
		if p1.Name != "prefix" {
			t.Error("Must have name [prefix]")
		}
		if p1.Default != "user" {
			t.Error("Must have default [user]")
		}
		if p2.Name != "suffix" {
			t.Error("Must have name [suffix]")
		}
		if p2.Default != "id" {
			t.Error("Must have default [id]")
		}
	}
}
*/

func TestJSONObjectWithParametersParsing(t *testing.T) {
	input := `{"id":"${id}","name":"${name}","params":{"id":1,"name":"YM"}}`
	expected := `{"id":1,"name":"YM"}`
	jo, err := ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps := jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test := jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}", "name":"${name}","params":{"name":"YM"}}`
	expected = `{"name":"YM"}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}", "name":"${name}","products":["${p1}","${p2}"]}`
	expected = `{"id":1,"name":"YM","products":[1,2]}`
	jo, err = ParseObjectString(input)
	fmt.Println(jo.String())
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}

	params, _ := ParseObjectString(`{"id":1,"name":"YM","index":"appl","size":20,"p1":1,"p2":2}`)
	jo.SetParams(params.Properties)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}","name":"${name}", "child":{"name":"${name}", "age":"${age}","${extra_field}":"${extra_value}"}, "params":{"id":1,"name":"YM","age":13,"extra_field":"nick","extra_value":"Gusyonok"}}`
	expected = `{"id":1,"name":"YM","child":{"name":"YM","age":13,"nick":"Gusyonok"}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"id":"${id}","name":"${name}", "child":{"age":"${age}"}, "params":{"id":1,"name":"YM"}}`
	expected = `{"id":1,"name":"YM"}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"${id_prefix}_id":"${id}","name":"${name}", "child":{"name":"${name} Jr.", "age":"${age}"}, "params":{"id_prefix":"user","id":1,"name":"YM"}}`
	expected = `{"user_id":1,"name":"YM","child":{"name":"YM Jr."}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"${id_prefix}_id":"${id}","name":"${name}", "child":{"name":"${name} Jr.", "age":"${age}"}, "params":{"id_prefix1":"user","id":1,"name":"YM"}}`
	expected = `{"_id":1,"name":"YM","child":{"name":"YM Jr."}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"template":{"query":{"query_string":{"query":"${query}","fields":"${fields}"}}},"params":{"query":"test","fields":["head","body"]}}`
	expected = `{"template":{"query":{"query_string":{"query":"test","fields":["head","body"]}}}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"template":{"_source":["headline","${field1}","${field2}","${field3}"],"params":{"field1":"type","field2":"date"}}}`
	expected = `{"template":{"_source":["headline","type","date"]}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"first_name":"${first_name?Yuri}","last_name":"${last_name?Metelkine}","params":{"last_name":"Metelkin"}}`
	expected = `{"first_name":"Yuri","last_name":"Metelkin"}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"name":"${first_name?Yuri} ${last_name?Metelkine}","params":{"last_name":"Metelkin"}}`
	expected = `{"name":"Yuri Metelkin"}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"template":{"_source":["headline","${field1}","${field2}","${field3?test}"],"params":{"field1":"type","field2":"date"}}}`
	expected = `{"template":{"_source":["headline","type","date","test"]}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	ps = jo.GetParams()
	for k, v := range ps {
		fmt.Printf("%s\t%s\n", k, v)
	}
	jo.SetParams(nil)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"bool":{"must":{"query_string":{"query":"${query}","fields":["head","body"]}},"filter":[{"terms":{"type":"${media_types}"}},{"terms":{"filing.products":"${include_products}"}}],"must_not":{"terms":{"filing.products":"${exclude_products}"}}}}}`
	expected = `{"query":{"bool":{"must":{"query_string":{"query":"ap","fields":["head","body"]}},"filter":[{"terms":{"type":["audio"]}},{"terms":{"filing.products":[1,2]}}],"must_not":{"terms":{"filing.products":[3]}}}}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	params, _ = ParseObjectString(`{"query":"ap","media_types":["audio"],"include_products":[1,2], "exclude_products":[3]}`)
	jo.SetParams(params.Properties)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"bool":{"must":{"query_string":{"query":"${query}","fields":["head","body"]}},"filter":[{"terms":{"type":"${media_types}"}},{"terms":{"filing.products":"${include_products}"}}],"must_not":{"terms":{"filing.products":"${exclude_products}"}}}}}`
	expected = `{"query":{"bool":{"must":{"query_string":{"query":"ap","fields":["head","body"]}},"filter":[{"terms":{"type":["audio"]}}]}}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	params, _ = ParseObjectString(`{"query":"ap","media_types":["audio"]}`)
	jo.SetParams(params.Properties)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}

	input = `{"query":{"bool":{"must":{"query_string":{"query":"${query}","fields":["head","body"]}},"filter":[{"terms":{"type":"${media_types}"}},{"terms":{"filing.products":"${include_products}"}}],"must_not":{"terms":{"filing.products":"${exclude_products}"}}}}}`
	expected = `{"query":{"bool":{"filter":[{"terms":{"type":["audio"]}}]}}}`
	jo, err = ParseObjectString(input)
	if err != nil {
		t.Error(err.Error())
	}
	params, _ = ParseObjectString(`{"media_types":["audio"]}`)
	jo.SetParams(params.Properties)
	test = jo.InlineString()
	fmt.Println(test)
	if test != expected {
		t.Error("Doesn't match!")
		fmt.Println(test)
	}
}
