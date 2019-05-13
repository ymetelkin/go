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
		fmt.Printf("\n%s\n%s\n\n", s, jo.String())
	}

	s = "@mediatype = text AND headline = trump"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.String())
	}

	s = "@mediatype = text byline = trump"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.String())
	}

	s = "@mediatype = text AND byline = trump"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.String())
	}

	s = "@mediatype=* AND @transmissionsource = NotMemberFeed AND @signal != Heartbeat"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.String())
	}

	s = "(@transmissionsource=livephoto or @transmissionsource=archivephoto) AND @transref != xgrbank AND @transmissionsource != MemberFeed"
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.String())
	}

	s = `@source = "Department of Defense" OR @source = NATO OR @source = Matched`
	jo, err = tr.Query(s)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("\n%s\n%s\n\n", s, jo.String())
	}
}
