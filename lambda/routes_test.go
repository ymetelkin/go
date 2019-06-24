package lambda

import (
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func f(params ...string) (events.APIGatewayProxyResponse, error) {
	var sb strings.Builder

	for _, p := range params {
		sb.WriteString(p)
		sb.WriteString(" ")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       sb.String(),
	}, nil
}

func f0(req Request) (events.APIGatewayProxyResponse, error) {
	return f(req.Method, req.Path)
}

func f1(req Request) (events.APIGatewayProxyResponse, error) {
	return f(req.Method, req.Path, req.PathParameters["id"])
}

func TestRoutes(t *testing.T) {
	rt := Router{}
	rt.Add("/", "GET", f0)
	rt.Add("/health", "GET", f0)
	rt.Add("/crud", "POST", f0)
	rt.Add("/crud", "DELETE", f0)
	rt.Add("/links/{id}", "GET", f1)
	rt.Add("/optional/{id?a}", "GET", f1)

	params := make(map[string]string)

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
	}
	rq, ok := rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected GET /"))
	}
	resp, _ := rq.Execute()
	if resp.Body != "GET / " {
		t.Error(errors.New("Expected GET / "))
	}

	params["proxy"] = "/health"
	req.PathParameters = params
	rq, ok = rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected GET /health"))
	}
	resp, _ = rq.Execute()
	if resp.Body != "GET /health " {
		t.Error(errors.New("Expected GET /health"))
	}

	params["proxy"] = "/crud"
	req.HTTPMethod = "POST"
	req.PathParameters = params
	rq, ok = rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected POST /crud"))
	}
	resp, _ = rq.Execute()
	if resp.Body != "POST /crud " {
		t.Error(errors.New("Expected POST /crud"))
	}

	req.HTTPMethod = "DELETE"
	rq, ok = rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected DELETE /crud"))
	}
	resp, _ = rq.Execute()
	if resp.Body != "DELETE /crud " {
		t.Error(errors.New("Expected DELETE /crud"))
	}

	params["proxy"] = "/links/a"
	req.HTTPMethod = "GET"
	req.PathParameters = params
	rq, ok = rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected GET /links/{id}"))
	}

	id, ok := rq.PathParameters["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /links/a"))
	}
	resp, _ = rq.Execute()
	if resp.Body != "GET /links/a a " {
		t.Error(errors.New("Expected GET /links/a"))
	}

	params["proxy"] = "/links/a/?q=test"
	req.PathParameters = params
	rq, ok = rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected GET /links/{id}"))
	}

	id, ok = rq.PathParameters["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /links/a"))
	}
	resp, _ = rq.Execute()
	if resp.Body != "GET /links/a a " {
		t.Error(errors.New("Expected GET /links/a"))
	}

	params["proxy"] = "/optional"
	req.HTTPMethod = "GET"
	req.PathParameters = params
	rq, ok = rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected GET /optional/{id?a}"))
	}

	id, ok = rq.PathParameters["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /optional/a"))
	}
	resp, _ = rq.Execute()
	if resp.Body != "GET /optional a " {
		t.Error(errors.New("Expected GET /optional/a"))
	}

	params["proxy"] = "/optional?q=v"
	req.HTTPMethod = "GET"
	req.PathParameters = params
	rq, ok = rt.GetRequest(req, true)
	if !ok {
		t.Error(errors.New("Expected GET /optional/{id?a}"))
	}

	id, ok = rq.PathParameters["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /optional/a"))
	}
	resp, _ = rq.Execute()
	if resp.Body != "GET /optional a " {
		t.Error(errors.New("Expected GET /optional/a"))
	}
}
