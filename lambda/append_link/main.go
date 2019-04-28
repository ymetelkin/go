package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ymetelkin/go/links"
)

func execute(req links.LinkRequest) (links.LinkResponse, error) {
	svc, err := links.New(os.Getenv("ES"))
	if err != nil {
		return links.LinkResponse{Status: links.Failure, Code: links.ElasticsearchError, Result: err.Error()}, nil
	}

	res := svc.AddLink(req)
	return res, nil
}

func main() {
	lambda.Start(execute)
}
