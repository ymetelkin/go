package links

import (
	"fmt"
	"testing"
)

func TestCRUD(t *testing.T) {
	id := "parent"
	col := Collection{ID: id}

	fmt.Println("Append three links")
	col.Append("d1")
	col.Append("d2")
	col.Append("d3")
	col.Insert("d0", 0)
	fmt.Println("In-memory collection")
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	db := newDb()

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
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
}
