package json

import (
	"fmt"
	"testing"
)

func TestArrayAdd(t *testing.T) {
	ja := JsonArray{}
	i, err := ja.Add("text")
	if err != nil {
		t.Error(err.Error())
	} else {
		jv, err := ja.GetValue(i)
		if err != nil {
			t.Error(err.Error())
		} else {
			text, err := jv.GetString()
			if err != nil {
				t.Error(err.Error())
			} else {
				fmt.Printf("%d\t%s\n", i, text)
			}
		}
	}

	i, err = ja.Add(1)
	if err != nil {
		t.Error(err.Error())
	} else {
		jv, err := ja.GetValue(i)
		if err != nil {
			t.Error(err.Error())
		} else {
			number, err := jv.GetInt()
			if err != nil {
				t.Error(err.Error())
			} else {
				fmt.Printf("%d\t%d\n", i, number)
			}
		}
	}

	i, err = ja.Add(true)
	if err != nil {
		t.Error(err.Error())
	} else {
		jv, err := ja.GetValue(i)
		if err != nil {
			t.Error(err.Error())
		} else {
			b, err := jv.GetBoolean()
			if err != nil {
				t.Error(err.Error())
			} else {
				fmt.Printf("%d\t%t\n", i, b)
			}
		}
	}

	jo := JsonObject{}
	jo.Add("id", 1)
	jo.Add("name", "YM")
	products := JsonArray{}
	products.Add(1)
	products.Add(2)
	jo.Add("products", products)
	i, err = ja.Add(jo)
	if err != nil {
		t.Error(err.Error())
	} else {
		jv, err := ja.GetValue(i)
		if err != nil {
			t.Error(err.Error())
		} else {
			obj, err := jv.GetObject()
			if err != nil {
				t.Error(err.Error())
			} else {
				fmt.Printf("%d\t%s\n", i, obj.ToInlineString())
			}
		}
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
	ja.Add("text")
	ja.Add(1)
	ja.Add(true)

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
	ja.Add("text")
	ja.Add(1)
	ja.Add(true)

	ja.Remove(1)

	if ja.Length() != 2 {
		t.Error("Must have size 2")
	}

	s := ja.ToString()
	fmt.Printf("%s\n", s)
}
