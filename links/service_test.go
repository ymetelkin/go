package links

import (
	"fmt"
	"testing"
)

func TestServiceCRUD(t *testing.T) {
	id := "YM"
	svc := New()
	err := svc.AddLink(id, "V")
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "O")
	if err != nil {
		t.Error(err.Error())
	}

	col, err := svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("Should have two links")
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	err = svc.RemoveLink(id, "V")
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "O")
	if err != nil {
		t.Error(err.Error())
	}

	col, err = svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	if len(col.Links) > 0 {
		t.Error("Should have no links")
	}
}
