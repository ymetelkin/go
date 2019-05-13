package es

import (
	"fmt"
	"testing"
)

func TestApplGetDocuments(t *testing.T) {
	svc, err := NewApplService("http://proteus-qa-uno-esdata.aptechlab.com:9200")
	if err != nil {
		t.Error(err.Error())
	}

	ids := []string{
		"1a087fa501d8445ab3d319fcbc72b709",
		"bbade2c1b43b4184bf0bee9eebdf9dce",
		"664259da4f1f429bab16307eea9a582f",
		"b416041bc1de48799ff18894836e14c6",
	}
	docs, err := svc.GetDocuments(ids, []string{"itemid", "headline"})
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Items IDs:")
	for i, id := range ids {
		fmt.Printf("%d\t%s\n", i, id)
	}

	fmt.Println("\nDocuments:")
	fmt.Println(docs.String())
}
