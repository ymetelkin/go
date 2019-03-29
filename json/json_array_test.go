package json

import (
	"fmt"
	"testing"
)

func TestArrayAdd(t *testing.T) {
	ja := JsonArray{}
	i := ja.AddString("text")
	jv, err := ja.GetValue(i)
	if err != nil {
		t.Error(err.Error())
	} else {
		text, _ := jv.Get()
		fmt.Printf("%d\t%s\n", i, text)
	}

	i = ja.AddInt(1)
	jv, err = ja.GetValue(i)
	if err != nil {
		t.Error(err.Error())
	} else {
		number, _ := jv.Get()
		fmt.Printf("%d\t%d\n", i, number)
	}

	i = ja.AddBoolean(true)
	jv, err = ja.GetValue(i)
	if err != nil {
		t.Error(err.Error())
	} else {
		b, _ := jv.Get()
		fmt.Printf("%d\t%t\n", i, b)
	}

	jo := JsonObject{}
	jo.AddInt("id", 1)
	jo.AddString("name", "YM")
	products := JsonArray{}
	products.AddInt(1)
	products.AddInt(2)
	jo.AddArray("products", &products)
	i = ja.AddObject(&jo)
	jv, err = ja.GetValue(i)
	if err != nil {
		t.Error(err.Error())
	} else {
		v, _ := jv.Get()
		obj, _ := v.(JsonObjectValue)
		fmt.Printf("%d\t%s\n", i, obj.Value.ToInlineString())
	}

	fmt.Printf("Is empty: %t\n", ja.IsEmpty())
	if ja.IsEmpty() {
		t.Error("Must not be empty")
	}

	s := ja.ToString()
	fmt.Printf("%s\n", s)
}

func TestArrayCopy(t *testing.T) {
	ja := JsonArray{}
	ja.AddString("text")
	ja.AddInt(1)
	ja.AddBoolean(true)

	copy := ja.Copy()

	fmt.Printf("Is empty: %t\n", copy.IsEmpty())
	if copy.IsEmpty() {
		t.Error("Must not be empty")
	}

	if copy.Length() != ja.Length() {
		t.Error("Must have same size as source")
	}

	s := copy.ToString()
	fmt.Printf("%s\n", s)
}

func TestArrayRemove(t *testing.T) {
	ja := JsonArray{}
	ja.AddString("text")
	ja.AddInt(1)
	ja.AddBoolean(true)

	ja.Remove(1)

	if ja.Length() != 2 {
		t.Error("Must have size 2")
	}

	s := ja.ToString()
	fmt.Printf("%s\n", s)
}
