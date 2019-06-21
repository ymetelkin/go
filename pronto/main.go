package main

import (
	"flag"
	"fmt"
)

func main() {
	s3 := flag.String("s3", "proteus-int-apcapdevelopment-all-us-east-1-index", "S3 bucket")
	es := flag.String("es", "http://proteus-qa-uno-esdata.aptechlab.com:9200", "Elasticsearch cluster")
	flag.Parse()
	fmt.Println(*s3)
	fmt.Println(*es)

	docs, err := getS3Docs(*s3)
	if err != nil {
		fmt.Printf("Failed to load Pronto templates from S3: %s\n", err.Error())
		return
	}

	for i, doc := range docs {
		fmt.Printf("%d\t%s\n%s\n\n", i+1, doc.ID, doc.Body)
	}

	idx := newIndexer(*es)
	total, err := idx.reindex(docs)
	if err != nil {
		fmt.Printf("Failed to reindex Pronto templates: %s\n", err.Error())
		return
	}

	fmt.Printf("Imported %d Pronto templates\n", total)
}
