package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ymetelkin/go/links"
)

func main() {
	lambda.Start(execute)
}

func execute(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	svc, err := links.New(os.Getenv("ES"))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	switch req.HTTPMethod {
	case "GET":
		return collection(req, svc), nil
	default:
		return link(req, svc), nil
	}
}

func link(req events.APIGatewayProxyRequest, svc links.Service) events.APIGatewayProxyResponse {
	var (
		rq links.LinkRequest
		rs links.LinkResponse
	)

	bytes := []byte(req.Body)
	if err := json.Unmarshal(bytes, &rq); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Invalid link request [%s]: %s", req.Body, err.Error()),
		}
	}

	switch req.HTTPMethod {
	case "PUT":
		rs = svc.AddLink(rq)
	case "POST":
		rs = svc.MoveLink(rq)
	case "DELETE":
		rs = svc.RemoveLink(rq)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       fmt.Sprintf("Method [%s] is not allowed on [%s]", req.HTTPMethod, req.Path),
		}
	}

	body, err := json.Marshal(rs)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: rs.Status,
		Body:       string(body),
	}
}

func collection(req events.APIGatewayProxyRequest, svc links.Service) events.APIGatewayProxyResponse {
	var (
		action, id string
		rq         links.GetCollectionRequest
		rs         links.GetCollectionResponse
	)

	tokens := strings.Split(req.Path, "?")
	path := strings.TrimPrefix(tokens[0], "/")
	tokens = strings.Split(path, "/")
	if len(tokens) > 1 {
		action = tokens[0]
		id = tokens[1]
	}

	if (action != "links" && action != "docs") || id == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Invalid collection request [%s]", req.Path),
		}
	}
	rq.CollectionID = id

	fields, _ := req.QueryStringParameters["fields"]
	if fields != "" {
		rq.Fields = strings.Split(fields, ",")
	}

	uid, _ := req.QueryStringParameters["uid"]
	if uid == "" {
		uid = "Anonymous"
	}
	rq.UserID = uid

	if action == "docs" {
		rs = svc.GetReversedCollection(rq)
	} else {
		rs = svc.GetCollection(rq)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: rs.Status,
		Body:       rs.ToString(),
	}
}
