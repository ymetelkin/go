package main

import (
	"fmt"
	"testing"
)

func TestES(t *testing.T) {
	d := db{
		Table: "apnews-qa-associatedpressqa-us-east-1-BusinessObjects",
		Env:   "qa",
	}

	items, err := d.list()
	if err != nil {
		t.Error(err)
		return
	}

	es := newIndexer("http://proteus-qa-unp-esdata.aptechlab.com:9200")

	total, err := es.reindex(items)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Imported %d business items", total)
}
