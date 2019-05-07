package apql

import (
	"fmt"
	"testing"
)

func TestSets(t *testing.T) {
	s := "@a = b"
	toks, _ := tokenize(s)
	set, err := booleanize(toks)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", set.string())
	}

	s = "@a = b AND @c=d"
	toks, _ = tokenize(s)
	set, err = booleanize(toks)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", set.string())
	}

	s = "(@a = b AND @c=d) (@e = f AND @x=1)"
	toks, _ = tokenize(s)
	set, err = booleanize(toks)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%s\n", set.string())
	}
}
