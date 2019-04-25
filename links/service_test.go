package links

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	svc, err := New("http://proteus-qa-uno-esdata.aptechlab.com:9200")
	if err != nil {
		t.Error(err.Error())
	}

	id := "1a087fa501d8445ab3d319fcbc72b709"
	err = svc.AddLink(id, "664259da4f1f429bab16307eea9a582f", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "b416041bc1de48799ff18894836e14c6", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, docs, err := svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s should have 2 links\n", id)
	for i, link := range col.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\t%s\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime(), docs[i].Type, docs[i].Date)
	}

	id = "bbade2c1b43b4184bf0bee9eebdf9dce"
	err = svc.AddLink(id, "b416041bc1de48799ff18894836e14c6", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.AddLink(id, "664259da4f1f429bab16307eea9a582f", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, docs, err = svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("\n%s should have 2 links\n", id)
	for i, link := range col.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\t%s\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime(), docs[i].Type, docs[i].Date)
	}

	id = "664259da4f1f429bab16307eea9a582f"
	rvs, docs, err := svc.GetReversedCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("\n%s should be linked to 2 docs\n", id)
	for i, link := range rvs.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\t%s\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime(), docs[i].Type, docs[i].Date)
	}
	id = "b416041bc1de48799ff18894836e14c6"
	rvs, docs, err = svc.GetReversedCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("\n%s should be linked to 2 docs\n", id)
	for i, link := range rvs.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\t%s\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime(), docs[i].Type, docs[i].Date)
	}

	id = "1a087fa501d8445ab3d319fcbc72b709"
	err = svc.RemoveLink(id, "664259da4f1f429bab16307eea9a582f", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "b416041bc1de48799ff18894836e14c6", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, docs, err = svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	if len(col.Links) > 0 {
		t.Errorf("\n %s should have no links", id)
	}

	id = "bbade2c1b43b4184bf0bee9eebdf9dce"
	err = svc.RemoveLink(id, "664259da4f1f429bab16307eea9a582f", id)
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.RemoveLink(id, "b416041bc1de48799ff18894836e14c6", id)
	if err != nil {
		t.Error(err.Error())
	}
	col, docs, err = svc.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	if len(col.Links) > 0 {
		t.Errorf("\n %s should have no links", id)
	}

	id = "b416041bc1de48799ff18894836e14c6"
	rvs, docs, err = svc.GetReversedCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	if len(col.Links) > 0 {
		t.Errorf("\n %s should have no links", id)
	}
}
