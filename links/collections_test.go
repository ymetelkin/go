package links

import (
	"fmt"
	"testing"
)

func TestCollections(t *testing.T) {
	col := collection{ID: "A"}
	by := "YM"

	fmt.Println("Append three links")
	i, err := col.append("0", "/link", by)
	if err != nil {
		t.Error(err.Error())
	} else if i != 0 {
		t.Error("Invalid position")
	}

	i, err = col.append("1", "/link", by)
	if err != nil {
		t.Error(err.Error())
	} else if i != 1 {
		t.Error("Invalid position")
	}

	i, err = col.append("2", "/link", by)
	if err != nil {
		t.Error(err.Error())
	} else if i != 2 {
		t.Error("Invalid position")
	}

	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("Move 2 to the 0 position")
	_, err = col.move("2", 0, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("Move 2 to the back")
	_, err = col.move("2", 2, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
}
