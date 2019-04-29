package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ymetelkin/go/appl"
)

func main() {
	lambda.Start(execute)
}

func execute(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jo, err := appl.XMLToJSON(req.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       jo.ToString(),
		Headers:    headers,
	}, nil
}
