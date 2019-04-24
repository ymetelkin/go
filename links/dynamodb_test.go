package links

import (
	"fmt"
	"testing"
)

func TestCRUD(t *testing.T) {
	id := "parent"
	by := "YM"
	col := Collection{ID: id}

	fmt.Println("Append three links")
	col.Append("d1", by)
	col.Append("d2", by)
	col.Append("d3", by)
	col.Insert("d0", 0, by)
	fmt.Println("In-memory collection")
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	db := newDb("LinkCollections")

	err := db.SaveCollection(col)
	if err != nil {
		t.Error(err.Error())
	}

	col, err = db.GetCollection(id)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("\nStored collection")
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\t%s:%v\t%s\n", link.Seq, link.ID, link.Updated.ID, link.Updated.Unix, link.Updated.DateTime())
	}
}
