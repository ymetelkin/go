package json

import (
	"fmt"
	"testing"
)

func TestObject(t *testing.T) {
	jo := New(
		Field("id", Int(1)),
		Field("name", String("YM")),
		Field("scores", Ints([]int{95, 98, 93, 100})),
	)

	if len(jo.Properties) != 3 {
		t.Error("Failed to create object")
	}

	s, ok := jo.GetString("name")
	if !ok {
		t.Error("Failed to read name")
	}
	fmt.Printf("%v %T\n", s, s)
	s, ok = jo.GetString("id")
	if !ok {
		t.Error("Failed to read id")
	}
	fmt.Printf("%v %T\n", s, s)
	i, ok := jo.GetInt("id")
	if !ok {
		t.Error("Failed to read id")
	}
	fmt.Printf("%v %T\n", i, i)
	f, ok := jo.GetFloat("id")
	if !ok {
		t.Error("Failed to read id")
	}
	fmt.Printf("%v %T\n", f, f)
	b, ok := jo.GetBool("id")
	if !ok {
		t.Error("Failed to read id")
	}
	fmt.Printf("%v %T\n", b, b)
	is, ok := jo.GetInts("scores")
	if !ok {
		t.Error("Failed to read scores")
	}
	fmt.Printf("%v %T\n", is, is)
	ss, ok := jo.GetStrings("scores")
	if !ok {
		t.Error("Failed to read scores")
	}
	fmt.Printf("%v %T\n", ss, ss)
}
