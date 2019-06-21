package main

import (
	"fmt"
	"testing"
)

func TestAdmin(t *testing.T) {
	s := "(categories.code:a OR categories.name:a) NOT CR"
	q, _ := fixQuery(s)
	fmt.Printf("%s\n%s\n\n", s, q)

	s = "Chicago AND NOT categories.code:s"
	q, _ = fixQuery(s)
	fmt.Printf("%s\n%s\n\n", s, q)

	s = "(Chicago AND NOT categories.code:s) OR NOT categories.code:j OR NOT categories.code:q"
	q, _ = fixQuery(s)
	fmt.Printf("%s\n%s\n\n", s, q)

	s = "filings.products:(45602)"
	q, _ = fixQuery(s)
	fmt.Printf("%s\n%s\n\n", s, q)

	s = "arrivaldatetime:{* TO now-3d}"
	q, _ = fixQuery(s)
	fmt.Printf("%s\n%s\n\n", s, q)

	s = "immigration AND NOT headline:(\"\"\"AP Top\"\"\")"
	q, _ = fixQuery(s)
	fmt.Printf("%s\n%s\n\n", s, q)

	s = "NOT sports, NOT weather, NOT lottery"
	q, _ = fixQuery(s)
	fmt.Printf("%s\n%s\n\n", s, q)
}
