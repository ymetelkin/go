package json

import (
	"fmt"
	"testing"
)

func TestObjectAdd(t *testing.T) {
	jo := JsonObject{}
	err := jo.Add("id", 1)
	if err != nil {
		t.Error(err.Error())
	} else {
		id, err := jo.GetInt("id")
		if err != nil {
			t.Error(err.Error())
		} else {
			fmt.Printf("id: %d\n", id)
		}
	}

	err = jo.Add("name", "YM")
	if err != nil {
		t.Error(err.Error())
	} else {
		name, err := jo.GetString("name")
		if err != nil {
			t.Error(err.Error())
		} else {
			fmt.Printf("name: %s\n", name)
		}
	}

	err = jo.Add("cool", true)
	if err != nil {
		t.Error(err.Error())
	} else {
		cool, err := jo.GetBoolean("cool")
		if err != nil {
			t.Error(err.Error())
		} else {
			fmt.Printf("cool: %t\n", cool)
		}
	}

	child := JsonObject{}
	child.Add("a", "b")
	err = jo.Add("child", child)
	if err != nil {
		t.Error(err.Error())
	} else {
		c, err := jo.GetObject("child")
		if err != nil {
			t.Error(err.Error())
		} else {
			fmt.Printf("child: %s\n", c.ToInlineString())
		}
	}

	ja := JsonArray{}
	ja.Add(1)
	ja.Add(2)
	err = jo.Add("products", ja)
	if err != nil {
		t.Error(err.Error())
	} else {
		products, err := jo.GetArray("products")
		if err != nil {
			t.Error(err.Error())
		} else {
			fmt.Printf("products: %s\n", products.ToInlineString())
		}
	}

	fmt.Printf("Is empty: %t\n", jo.IsEmpty())
	if jo.IsEmpty() {
		t.Error("Must not be empty")
	}

	fmt.Printf("%s\n", jo.ToString())
}

func TestObjectCopy(t *testing.T) {
	jo := JsonObject{}
	jo.Add("id", 1)
	jo.Add("name", "YM")
	jo.Add("cool", true)

	child := jo.Copy()
	child.Add("a", "b")
	jo.Add("child", *child)

	for key, value := range jo.Properties {
		fmt.Printf("%s: %#v\n", key, value)
		jv, err := jo.GetValue(key)
		if err != nil {
			t.Error(err.Error())
		} else {
			fmt.Printf("%s: %v\n", key, jv.Value)
		}
	}

	fmt.Printf("Is empty: %t\n", jo.IsEmpty())
	if jo.IsEmpty() {
		t.Error("Must not be empty")
	}

	fmt.Printf("%s\n", jo.ToString())
}

func TestObjectRemove(t *testing.T) {
	jo := JsonObject{}
	jo.Add("id", 1)
	jo.Add("name", "YM")

	jo.Remove("id")
	if len(jo.Properties) != 1 {
		t.Error("Failed to remove")
	}

	jo.Remove("foo")
	if len(jo.Properties) != 1 {
		t.Error("Failed to remove")
	}

	err := jo.Remove("")
	if err == nil {
		t.Error("Failed to fail to remove empty field")
	}

	fmt.Printf("%s\n", jo.ToString())
}
