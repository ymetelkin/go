package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ymetelkin/go/links"
)

func main() {
	lambda.Start(execute)
}

func execute(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	svc, err := links.New()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	switch req.HTTPMethod {
	case "GET":
		return collection(req, svc), nil
	case "POST":
		if req.Path == "/reset" {
			return reset(req, svc), nil
		}
		return link(req, svc), nil
	default:
		return link(req, svc), nil
	}
}

func link(req events.APIGatewayProxyRequest, svc links.Service) events.APIGatewayProxyResponse {
	var (
		rs links.LinkResponse
		e  string
	)

	switch req.HTTPMethod {
	case "PUT":
		rs, e = add(req, svc)
	case "POST":
		rs, e = move(req, svc, false)
	case "DELETE":
		rs, e = move(req, svc, true)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       fmt.Sprintf("Method [%s] is not allowed on [%s]", req.HTTPMethod, req.Path),
		}
	}

	if e != "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       e,
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

func add(req events.APIGatewayProxyRequest, svc links.Service) (links.LinkResponse, string) {
	var rq links.LinkRequest

	bytes := []byte(req.Body)
	if err := json.Unmarshal(bytes, &rq); err != nil {
		return links.LinkResponse{}, fmt.Sprintf("Invalid link add request [%s]: %s", req.Body, err.Error())
	}

	return svc.AddLink(rq), ""
}

func move(req events.APIGatewayProxyRequest, svc links.Service, rm bool) (links.LinkResponse, string) {
	var rq links.MoveRequest

	bytes := []byte(req.Body)
	if err := json.Unmarshal(bytes, &rq); err != nil {
		return links.LinkResponse{}, fmt.Sprintf("Invalid link move request [%s]: %s", req.Body, err.Error())
	}

	if rm {
		return svc.RemoveLink(rq), ""
	}

	return svc.MoveLink(rq), ""
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
		Body:       rs.String(),
	}
}

func reset(req events.APIGatewayProxyRequest, svc links.Service) events.APIGatewayProxyResponse {
	var (
		rq links.ResetRequest
		rs links.LinkResponse
	)

	bytes := []byte(req.Body)
	if err := json.Unmarshal(bytes, &rq); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Invalid reset request [%s]: %s", req.Body, err.Error()),
		}
	}

	rs = svc.ResetLinks(rq)

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
