package apql

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ymetelkin/go/json"
)

func TestProducts(t *testing.T) {
	data, err := ioutil.ReadFile("products.json")
	if err != nil {
		t.Error(err.Error())
		return
	}

	tr := New()

	hits, _ := json.ParseJSONArray(string(data))
	jos, _ := hits.GetObjects()
	for _, jo := range jos {
		id, _ := jo.GetString("_id")
		src, _ := jo.GetObject("_source")
		q1, _ := src.GetObject("product_query")
		apql, _ := src.GetString("product_apql")
		q2, err := tr.Query(apql)
		if err != nil {
			t.Error(err.Error())
		}

		s1 := q1.InlineString()
		s2 := q2.InlineString()
		if s1 != s2 {
			fmt.Printf("%s\tERROR!\n%s\n%s\n%s\n", id, apql, s1, s2)
			return
		}
	}
}
