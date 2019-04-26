package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ymetelkin/go/links"
)

func execute(req links.LambdaLinkRequest) (links.LambdaResponse, error) {
	svc, err := links.New(os.Getenv("ES"))
	if err != nil {
		return links.LambdaResponse{Text: err.Error()}, nil
	}

	err = svc.RemoveLink(req.CollectionID, req.LinkID, req.UserID)
	if err != nil {
		return links.LambdaResponse{Text: err.Error()}, nil
	}
	col, docs, err := svc.GetCollection(req.CollectionID)
	if err != nil {
		return links.LambdaResponse{Text: err.Error()}, nil
	}

	return links.NewLambdaResponse(true, "Success", col, docs), nil
}

func main() {
	lambda.Start(execute)
}
