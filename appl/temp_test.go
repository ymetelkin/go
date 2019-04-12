package appl

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	s := `
	<Publication Version="5.3.1" xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">

  </Publication>`
	jo, _ := XMLToJSON(s)

	fmt.Printf("%s\n", jo.ToString())

}
