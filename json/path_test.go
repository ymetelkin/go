package json

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	s := `{"id":1,"code":"YM","children":[{"id":10,"code":"O","parents":[{"id":1,"code":"YM","test":["a","b"]},{"id":2,"code":"SV","test":["c","d"]}]},{"id":11,"code":"V","parents":[{"id":1,"code":"YM","test":["e","f"]},{"id":2,"code":"SV","test":["j","k"]}]}]}`
	jo, _ := ParseObjectString(s)

	s, e := jo.PathString("code")
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Printf("%v\n", s)

	ss, e := jo.PathStrings("children.code")
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Printf("%v\n", ss)

	ss, e = jo.PathStrings("children.parents.code")
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Printf("%v\n", ss)

	ss, e = jo.PathStrings("children.parents.test")
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Printf("%v\n", ss)
}
