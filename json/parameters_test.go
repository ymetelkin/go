package json

import (
	"fmt"
	"testing"
)

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

	s = `{"${f?id}":1}`
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

	s = `{"${prefix?user}_${suffix?id}":1}`
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

	s = `"${v?test}"`
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

	s = `"prefix_${v?test}_suffix"`
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

	s = `"This is ${prefix?user} ${suffix?id} xyz"`
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

func TestJSONObjectWithParametersParsing(t *testing.T) {
	s := `{"id":"${id}","name":"${name}","params":{"id":1,"name":"YM"}}`
	jo, err := ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"id":"${id}", "name":"${name}","params":{"name":"YM"}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"id":"${id}","name":"${name}", "child":{"name":"${name}", "age":"${age}","${extra_field}":"${extra_value}"}, "params":{"id":1,"name":"YM","age":13,"extra_field":"nick","extra_value":"Gusyonok"}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"id":"${id}","name":"${name}", "child":{"age":"${age}"}, "params":{"id":1,"name":"YM"}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"${id_prefix}_id":"${id}","name":"${name}", "child":{"name":"${name} Jr.", "age":"${age}"}, "params":{"id_prefix":"user","id":1,"name":"YM"}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"${id_prefix}_id":"${id}","name":"${name}", "child":{"name":"${name} Jr.", "age":"${age}"}, "params":{"id_prefix1":"user","id":1,"name":"YM"}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"template":{"query":{"query_string":{"query":"${query}","fields":"${fields}"}}},"params":{"query":"test","fields":["head","body"]}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"template":{"_source":["headline","${field1}","${field2}","${field3}"],"params":{"field1":"type","field2":"date"}}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"first_name":"${first_name?Yuri}","last_name":"${last_name?Metelkine}","params":{"last_name":"Metelkin"}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"name":"${first_name?Yuri} ${last_name?Metelkine}","params":{"last_name":"Metelkin"}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}

	s = `{"template":{"_source":["headline","${field1}","${field2}","${field3?test}"],"params":{"field1":"type","field2":"date"}}}`
	jo, err = ParseJSONObjectWithParameters(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", jo.ToString())
	}
}
