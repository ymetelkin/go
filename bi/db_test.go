package main

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	d := db{
		Table: "apnews-qa-associatedpressqa-us-east-1-BusinessObjects",
		Env:   "qa",
	}

	items, err := d.list()
	if err != nil {
		t.Error(err)
		return
	}

	for k, v := range items {
		fmt.Printf("%s\n%s\n\n", k, v)
	}
}
