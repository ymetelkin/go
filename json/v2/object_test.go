package v2

import (
	"fmt"
	"testing"
)

func TestObjectParse(t *testing.T) {
	s := `{ "text": "abc", "number": 3.14, "flag": true, "array": [ 1, 2, 3 ], "object": { "a": "b" }}`
	p := newParser([]byte(s))
	err := p.SkipWS()
	if err != nil {
		t.Error(err.Error())
	}
	if p.Byte != '{' {
		t.Error("Failed to parse {")
	}
	v, err := p.ParseObject(false)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(v.String())
}

func TestObjectPointers(t *testing.T) {
	jo := New(Field("name", String("YM")))
	fmt.Println(jo.String())
	jo.Add("person", jo)
	fmt.Println(jo.String())
}
