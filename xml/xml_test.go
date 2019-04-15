package xml

import (
	"fmt"
	"testing"
)

func TestXml(t *testing.T) {
	s := `
	<p>
	     Before <a href="#">link</a>. after</p>`
	nd, _ := New(s)

	fmt.Printf("%s\n", nd.ToString())

}
