package json

import (
	"fmt"
	"testing"
)

func TestValue(t *testing.T) {
	v := Value{
		Type: TypeInt,
		data: 1,
	}
	i, ok := v.Int()
	fmt.Printf("%d\n", i)
	if !ok || i != 1 {
		t.Error("Failed to Int int")
	}
	f, ok := v.Float()
	fmt.Printf("%f\n", f)
	if !ok || f != 1.0 {
		t.Error("Failed to Float int")
	}
	s, ok := v.String()
	fmt.Printf("%s\n", s)
	if !ok || s != "1" {
		t.Error("Failed to String int")
	}

	v = Value{
		Type: TypeFloat,
		data: 1.4,
	}
	i, ok = v.Int()
	fmt.Printf("%d\n", i)
	if !ok || i != 1 {
		t.Error("Failed to Int float")
	}
	f, ok = v.Float()
	fmt.Printf("%f\n", f)
	if !ok || f != 1.4 {
		t.Error("Failed to Float float")
	}
	s, ok = v.String()
	fmt.Printf("%s\n", s)
	if !ok || s != "1.4" {
		t.Error("Failed to String float")
	}

	v = Value{
		Type: TypeBool,
		data: true,
	}
	b, ok := v.Bool()
	fmt.Printf("%v\n", b)
	if !ok || !b {
		t.Error("Failed to Bool bool")
	}
	s, ok = v.String()
	fmt.Printf("%s\n", s)
	if !ok || s != "true" {
		t.Error("Failed to String bool")
	}

	v = Value{
		Type: TypeString,
		data: "xyz",
	}
	s, ok = v.String()
	fmt.Printf("%s\n", s)
	if !ok || s != "xyz" {
		t.Error("Failed to String string")
	}
	v = Value{
		Type: TypeString,
		data: "1",
	}
	i, ok = v.Int()
	fmt.Printf("%d\n", i)
	if !ok || i != 1 {
		t.Error("Failed to Int string")
	}
	v = Value{
		Type: TypeString,
		data: "1.4",
	}
	f, ok = v.Float()
	fmt.Printf("%f\n", f)
	if !ok || f != 1.4 {
		t.Error("Failed to Float string")
	}
	v = Value{
		Type: TypeString,
		data: "true",
	}
	b, ok = v.Bool()
	fmt.Printf("%v\n", b)
	if !ok || !b {
		t.Error("Failed to Bool string")
	}
	v = Value{
		Type: TypeNull,
	}
	s, ok = v.String()
	fmt.Printf("%s\n", s)
	if !ok || s != "null" {
		t.Error("Failed to String null string")
	}

	var jo Object
	jo.AddString("name", "xyz")
	v = Value{
		Type: TypeObject,
		data: jo,
	}
	jo, ok = v.Object()
	fmt.Printf("%v\n", jo)
	if !ok {
		t.Error("Failed to Object object")
	}
	s, ok = v.String()
	fmt.Printf("%s\n", s)
	if !ok || s != `{"name":"xyz"}` {
		t.Error("Failed to String object")
	}

	var ja Array
	ja.AddString("xyz")
	v = Value{
		Type: TypeArray,
		data: ja,
	}
	ja, ok = v.Array()
	fmt.Printf("%v\n", ja)
	if !ok {
		t.Error("Failed to Array array")
	}
	s, ok = v.String()
	fmt.Printf("%s\n", s)
	if !ok || s != `["xyz"]` {
		t.Error("Failed to String array")
	}
}
