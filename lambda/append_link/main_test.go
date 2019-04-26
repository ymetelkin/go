package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Setenv("ES", "http://proteus-int-all-esclient.aptechdevlab.com:9200")

	req := request{
		CollectionID: "90c7709d40e24468b0a707377d58c1db",
		LinkID:       "887404dcb0774b539f2634b99d48300f",
		UserID:       "YM",
	}

	res, _ := addLink(req)
	bytes, _ := json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))
}
