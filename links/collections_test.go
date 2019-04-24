package links

import (
	"fmt"
	"testing"
)

func TestCollections(t *testing.T) {
	col := Collection{ID: "parent"}
	by := "YM"

	fmt.Println("Append three links")
	i, err := col.Append("d1", by)
	if err != nil {
		t.Error(err.Error())
	} else if i != 0 {
		t.Error("Invalid position")
	}

	i, err = col.Append("d2", by)
	if err != nil {
		t.Error(err.Error())
	} else if i != 1 {
		t.Error("Invalid position")
	}

	i, err = col.Append("d3", by)
	if err != nil {
		t.Error(err.Error())
	} else if i != 2 {
		t.Error("Invalid position")
	}

	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("Insert d0 to the begining")
	err = col.Insert("d0", 0, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("Move d3 to the second position")
	err = col.Move("d3", 1, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("Move d3 to the back")
	err = col.Move("d3", 3, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
}

func TestReversedCollections(t *testing.T) {
	col := Collection{ID: "d1"}
	by := "YM"

	fmt.Println("Add p1 at 0")
	_, err := col.AddReversed("p1", 0, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("\nAdd p2 at 0")
	_, err = col.AddReversed("p2", 0, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("\nAdd p3 at 1")
	_, err = col.AddReversed("p3", 1, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("\nMove p3 to the first position")
	_, err = col.AddReversed("p3", 0, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}

	fmt.Println("\nMove p1 to the last position")
	_, err = col.AddReversed("p1", 10, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
	fmt.Println("\nMove p1 to the first position")
	_, err = col.AddReversed("p1", 0, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
	fmt.Println("\nMove p1 to the last position")
	_, err = col.AddReversed("p1", 10, by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
	fmt.Println("\nRemove p2")
	col.RemoveReversed("p2", by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
	fmt.Println("\nRemove p1")
	col.RemoveReversed("p1", by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
	fmt.Println("\nRemove p3")
	col.RemoveReversed("p3", by)
	if err != nil {
		t.Error(err.Error())
	}
	for _, link := range col.Links {
		fmt.Printf("%d\t%s\n", link.Seq, link.ID)
	}
}
