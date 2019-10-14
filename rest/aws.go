package rest

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

//APIGateway AWS APIGatewayProxy
type APIGateway struct {
	router Router
}

//AddRoute method adds handler to collection of routes
func (api *APIGateway) AddRoute(path string, method string, handler Execute) {
	api.router.Add(path, method, handler)
}

//Execute is wrapper around REST Execute
func (api *APIGateway) Execute(r events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	req := Request{
		HTTPMethod:            r.HTTPMethod,
		QueryStringParameters: r.QueryStringParameters,
	}

	var path string

	if r.RequestContext.ResourcePath == "/{proxy+}" && r.PathParameters != nil {
		path, _ = r.PathParameters["proxy"]
	} else {
		path = r.Path
		if r.RequestContext.Stage != "" {
			req.Path = strings.Replace(req.Path, r.RequestContext.Stage, "", 1)
		}
	}

	req.Path = strings.Split(path, "?")[0]

	resp := api.router.Execute(req)

<<<<<<< HEAD
	headers := resp.Headers
	if headers == nil {
		headers = make(map[string]string)
	}

	_, ok := headers["Content-Type"]
	if !ok {
		size := len(resp.Body) - 1
		var (
			i  int
			ct string
		)

		for {
			if i > size || i > 512 {
				break
			}

			c := resp.Body[i]
			switch c {
			case '{':
				ct = "application/json"
				break
			case '<':
				ct = "application/xml"
				break
			}
		}

		if ct == "" {
			ct = "text/plain"
		}
		headers["Content-Type"] = ct
	}

	return events.APIGatewayProxyResponse{
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
		Headers:    headers,
=======
	return events.APIGatewayProxyResponse{
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
		Headers:    resp.Headers,
>>>>>>> e40de20d1101308366e6c1267131230e4c431cf7
	}
}
