package json

import (
	"fmt"
	"testing"
)

func TestObjectAdd(t *testing.T) {
	jo := Object{}
	ok := jo.AddInt("id", 1)
	if !ok {
		t.Error("Failed to add int")
	} else {
		id, ok := jo.GetInt("id")
		if !ok {
			t.Error("Failed to get int")
		} else {
			fmt.Printf("id: %d\n", id)
		}
	}

	ok = jo.AddString("name", "YM")
	if !ok {
		t.Error("Failed to add string")
	} else {
		name, ok := jo.GetString("name")
		if !ok {
			t.Error("Failed to get string")
		} else {
			fmt.Printf("name: %s\n", name)
		}
	}

	ok = jo.AddBool("cool", true)
	if !ok {
		t.Error("Failed to add bool")
	} else {
		cool, ok := jo.GetBool("cool")
		if !ok {
			t.Error("Failed to get bool")
		} else {
			fmt.Printf("cool: %t\n", cool)
		}
	}

	child := Object{}
	child.AddString("a", "b")
	ok = jo.AddObject("child", child)
	if !ok {
		t.Error("Failed to add object")
	} else {
		c, ok := jo.GetObject("child")
		if !ok {
			t.Error("Failed to get object")
		} else {
			fmt.Printf("child: %s\n", c.InlineString())
		}
	}

	ja := Array{}
	ja.AddInt(1)
	ja.AddInt(2)
	ok = jo.AddArray("products", ja)
	if !ok {
		t.Error("Failed to add array")
	} else {
		products, ok := jo.GetArray("products")
		if !ok {
			t.Error("Failed to get array")
		} else {
			fmt.Printf("products: %s\n", products.InlineString())
		}
	}

	fmt.Printf("Is empty: %t\n", jo.IsEmpty())
	if jo.IsEmpty() {
		t.Error("Must not be empty")
	}

	fmt.Printf("%s\n", jo.String())
}

func TestObjectRemove(t *testing.T) {
	jo := Object{}
	jo.AddInt("id", 1)
	jo.AddString("name", "YM")

	jo.Remove("id")
	if len(jo.props) != 1 {
		t.Error("Failed to remove")
	}

	jo.Remove("foo")
	if len(jo.props) != 1 {
		t.Error("Failed to remove")
	}

	fmt.Printf("%s\n", jo.String())
}
