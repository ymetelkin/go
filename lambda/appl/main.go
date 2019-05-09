package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ymetelkin/go/appl"
	"github.com/ymetelkin/go/json"
)

func main() {
	lambda.Start(execute)
}

func execute(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		status int
		jo     json.Object
	)

	jo, err := appl.XMLToJSON(req.Body)
	if err != nil {
		status = http.StatusBadRequest
		jo = json.Object{}
		jo.AddString("error", err.Error())
	} else {
		status = http.StatusOK
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       jo.String(),
		Headers:    headers,
	}, nil
}
