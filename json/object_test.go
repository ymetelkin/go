package json

import (
	"fmt"
	"testing"
)

func TestObjectAdd(t *testing.T) {
	jo := Object{}
	err := jo.AddInt("id", 1)
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

	err = jo.AddString("name", "YM")
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

	err = jo.AddBool("cool", true)
	if err != nil {
		t.Error(err.Error())
	} else {
		cool, err := jo.GetBool("cool")
		if err != nil {
			t.Error(err.Error())
		} else {
			fmt.Printf("cool: %t\n", cool)
		}
	}

	child := Object{}
	child.AddString("a", "b")
	err = jo.AddObject("child", child)
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

	ja := Array{}
	ja.AddInt(1)
	ja.AddInt(2)
	err = jo.AddArray("products", ja)
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

func TestObjectRemove(t *testing.T) {
	jo := Object{}
	jo.AddInt("id", 1)
	jo.AddString("name", "YM")

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
