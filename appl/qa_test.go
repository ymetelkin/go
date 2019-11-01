package appl

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestXMLToJSON(t *testing.T) {
	files, e := ioutil.ReadDir("tests")
	if e != nil {
		t.Error(e.Error())
	}

	for _, file := range files {
		var (
			jl, jr      JSON
			x           XML
			left, right string
		)

		path := "tests/" + file.Name()
		toks := strings.Split(path, ".")
		if toks[1] == "xml" {
			b, e := ioutil.ReadFile(path)
			if e != nil {
				t.Error(e.Error())
				return
			}

			//fmt.Println(string(b))

			e = xml.Unmarshal(b, &x)
			if e != nil {
				t.Error(e.Error())
				return
			}

			//fmt.Println(string(js.String()))

			jl = x.JSON()
			left, e = jl.String()

			jr = JSON{}
			b, e = ioutil.ReadFile(toks[0] + ".json")
			if e != nil {
				t.Error(e.Error())
				return
			}
			e = json.Unmarshal(b, &jr)
			if e != nil {
				t.Error(e.Error())
				return
			}
			right, e = jr.String()

			fmt.Println(left)
			fmt.Println(right)

			fmt.Printf("[%s] SUCCESS\n", path)
		}
	}

}
