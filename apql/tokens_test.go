package apql

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	s := "@a = b"
	toks, err := tokenize(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%v\n", toks)
	}

	s = "@a = b AND @c=d"
	toks, err = tokenize(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%v\n", toks)
	}

	s = "(@a = b AND @c=d) (@e = f AND @x=1)"
	toks, err = tokenize(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("%v\n", toks)
	}
}
