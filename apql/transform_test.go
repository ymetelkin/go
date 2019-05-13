package apql

import (
	"fmt"
	"testing"
)

func TestTransforms(t *testing.T) {
	tr := New()

	s := "@mediatype = text"
	jo, err := tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.ToString())
	}

	s = "@mediatype = text AND headline = trump"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.ToString())
	}

	s = "@mediatype = text byline = trump"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.ToString())
	}

	s = "@mediatype = text AND byline = trump"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.ToString())
	}
}
