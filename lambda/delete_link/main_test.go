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
		LinkID:       "bce05557d0e14ea4a6367f0d31e0ca7c",
		UserID:       "YM",
	}

	res, _ := removeLink(req)
	bytes, _ := json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))
}
