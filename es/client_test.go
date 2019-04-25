package es

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	ec, err := newClient("http://proteus-qa-uno-esdata.aptechlab.com:9200")
	if err != nil {
		t.Error(err.Error())
	}

	sr := newSearchRequest("appl-thirty")
	sr.SetQuery(`{"query":{"match":{"headline":"messi"}}}`)
	sr.SetSource([]string{"type", "headline", "arrivaldatetime"})
	sr.SetSize(100)

	cr, err := ec.Search(sr)
	if err != nil {
		t.Error(err.Error())
	}

	defer cr.Close()

	bytes, err := ioutil.ReadAll(cr)
	if err != nil {
		t.Error(err.Error())
	}

	s := string(bytes)
	fmt.Println(s)

	var body map[string]interface{}

	if err = json.NewDecoder(strings.NewReader(s)).Decode(&body); err != nil {
		t.Error(err.Error())
	}

	for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
		fmt.Printf("%s  %s\n", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"].(map[string]interface{})["headline"])
	}
}
