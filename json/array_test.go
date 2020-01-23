package json

import (
	"fmt"
	"testing"
)

func TestArrayAdd(t *testing.T) {
	ja := Array{}
	ja.AddString("text")
	s, ok := ja.GetString(0)
	if !ok {
		t.Error("Failed to get string")
	} else {
		fmt.Printf("%d\t%s\n", len(ja.Values), s)
	}

	ja.AddInt(1)
	number, ok := ja.GetInt(len(ja.Values) - 1)
	if !ok {
		t.Error("Failed to get number")
	} else {
		fmt.Printf("%d\t%d\n", len(ja.Values), number)
	}

	ja.AddBool(true)
	b, ok := ja.GetBool(len(ja.Values) - 1)
	if !ok {
		t.Error("Failed to get bool")
	} else {
		fmt.Printf("%d\t%t\n", len(ja.Values), b)
	}

	jo := Object{}
	jo.AddInt("id", 1)
	jo.AddString("name", "YM")
	products := Array{}
	products.AddInt(1)
	products.AddInt(2)
	jo.AddArray("products", products)
	ja.AddObject(jo)
	o, ok := ja.GetObject(len(ja.Values) - 1)
	if !ok {
		t.Error("Failed to get object")
	} else {
		fmt.Printf("%d\t%s\n", len(ja.Values), o.InlineString())
	}

	fmt.Printf("Is empty: %t\n", len(ja.Values) == 0)

	s = ja.String()
	fmt.Printf("%s\n", s)
}

func TestObjectArray(t *testing.T) {
	names := []string{"YM", "SV"}
	ja := Array{}
	for _, name := range names {
		jo := Object{}
		rels := Array{}
		rels.AddString(name)
		jo.AddArray("rels", rels)
		ja.AddObject(jo)
	}

	fmt.Println(ja.String())
}

func TestArrayRemove(t *testing.T) {
	ja := Array{}
	ja.AddString("text")
	ja.AddInt(1)
	ja.AddBool(true)

	ja.Remove(1)

	if len(ja.Values) != 2 {
		t.Error("Must have size 2")
	}

	s := ja.String()
	fmt.Printf("%s\n", s)
}
