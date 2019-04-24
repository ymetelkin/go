package links

import (
	"fmt"
	"testing"
)

func TestServiceCRUD(t *testing.T) {
	svc := New()

	id := "YM"
	err := svc.AddLink(id, "V", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "O", id)
	if err != nil {
		t.Error(err.Error())
	}

	col, err := svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("YM Should have two links")
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime())
	}

	err = svc.RemoveLink(id, "V", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "O", id)
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

func TestService(t *testing.T) {
	svc := New()

	id := "YM"
	err := svc.AddLink(id, "O", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "V", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "A", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, err := svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("YM should have 3 links")
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime())
	}

	id = "SV"
	err = svc.AddLink(id, "V", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "O", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "A", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, err = svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("\nSV should have 3 links")
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime())
	}

	rvs, err := svc.GetReversedCollection("O")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("\nO should be linked to 2 docs")
	for _, link := range rvs.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime())
	}
	rvs, err = svc.GetReversedCollection("V")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("\nV should be linked to 2 docs")
	for _, link := range rvs.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime())
	}
	rvs, err = svc.GetReversedCollection("A")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("\nA should be linked to 2 docs")
	for _, link := range rvs.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime())
	}

	err = svc.RemoveLink(id, "V", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "O", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "A", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, err = svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	if len(col.Links) > 0 {
		t.Error("\n SV should have no links")
	}

	id = "YM"
	err = svc.RemoveLink(id, "V", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "O", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "A", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, err = svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	if len(col.Links) > 0 {
		t.Error("\nYM should have no links")
	}

	rvs, err = svc.GetReversedCollection("A")
	if err != nil {
		t.Error(err.Error())
	}
	if len(col.Links) > 0 {
		t.Error("\nA should not be linked")
	}
}
