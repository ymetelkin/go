package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		p, s string
		dn   db
		es   indexer
	)

	for {
		if p == "" {
			p = "\nBUSINESS ITEMS RE-FEEDER\n==================================================================\nSelect your environment (qa/prod/unp): "
		}

		fmt.Print(p)
		fmt.Scanln(&s)
		s = strings.ToLower(s)
		if s == "qa" {
			dn = db{Table: "apnews-qa-associatedpressqa-us-east-1-BusinessObjects", Env: "qa"}
			es = newIndexer("http://proteus-qa-uno-esdata.aptechlab.com:9200")
		} else if s == "prod" {
			dn = db{Table: "apnews-qa-associatedpressqa-us-east-1-BusinessObjects", Env: "prod"}
			es = newIndexer("http://proteus-prd-uno-esdata.associatedpress.com:9200")
		} else if s == "unp" {
			dn = db{Table: "apnews-qa-associatedpressqa-us-east-1-BusinessObjects", Env: "qa"}
			es = newIndexer("http://proteus-qa-unp-esdata.aptechlab.com:9200")
		} else {
			p = fmt.Sprintf("[%s] is invalid input. Try again: ", s)
		}

		if dn.Table != "" {
			fmt.Println("Importing business items")
			fmt.Printf("from %s\n", dn.Table)
			fmt.Printf("to %s/%s\n", es.Cluster, es.Index)
			fmt.Print("Do you want to proceed? (y/n):")

			fmt.Scanln(&s)
			s = strings.ToLower(s)
			if s == "y" {
				break
			}
			p = ""
		}
	}

	items, err := dn.list()
	if err != nil {
		fmt.Printf("Failed to load items from Dynamo: %s", err.Error())
		return
	}

	total, err := es.reindex(items)
	if err != nil {
		fmt.Printf("Failed to reindex business items: %s", err.Error())
	}

	fmt.Printf("Imported %d business items", total)
}
