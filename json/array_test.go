package json

import (
	"fmt"
	"testing"
)

func TestArrayAdd(t *testing.T) {
	ja := Array{}
	i := ja.AddString("text")
	s, ok := ja.GetString(i)
	if !ok {
		t.Error("Failed to get string")
	} else {
		fmt.Printf("%d\t%s\n", i, s)
	}

	i = ja.AddInt(1)
	number, ok := ja.GetInt(i)
	if !ok {
		t.Error("Failed to get number")
	} else {
		fmt.Printf("%d\t%d\n", i, number)
	}

	i = ja.AddBool(true)
	b, ok := ja.GetBool(i)
	if !ok {
		t.Error("Failed to get bool")
	} else {
		fmt.Printf("%d\t%t\n", i, b)
	}

	jo := Object{}
	jo.AddInt("id", 1)
	jo.AddString("name", "YM")
	products := Array{}
	products.AddInt(1)
	products.AddInt(2)
	jo.AddArray("products", products)
	i = ja.AddObject(jo)
	o, ok := ja.GetObject(i)
	if !ok {
		t.Error("Failed to get object")
	} else {
		fmt.Printf("%d\t%s\n", i, o.InlineString())
	}

	fmt.Printf("Is empty: %t\n", ja.IsEmpty())
	if ja.IsEmpty() {
		t.Error("Must not be empty")
	}

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

	if ja.Length() != 2 {
		t.Error("Must have size 2")
	}

	s := ja.String()
	fmt.Printf("%s\n", s)
}
