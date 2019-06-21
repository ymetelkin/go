package lambda

import (
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

type handler struct {
	body string
}

func (h *handler) Execute() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       h.body,
	}, nil
}

func TestRoutes(t *testing.T) {
	pr := ProxyRouter{}
	pr.Add("/", "GET", &handler{"root"})
	pr.Add("/health", "GET", &handler{"health"})
	pr.Add("/crud", "POST", &handler{"update"})
	pr.Add("/crud", "DELETE", &handler{"delete"})
	pr.Add("/links/{id}", "GET", &handler{"links"})
	pr.Add("/optional/{id?a}", "GET", &handler{"optional"})

	params := make(map[string]string)

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
	}
	rm, ok := pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /"))
	}
	resp, _ := rm.Handler.Execute()
	if resp.Body != "root" {
		t.Error(errors.New("Expected root"))
	}

	params["proxy"] = "/health"
	req.PathParameters = params
	rm, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /health"))
	}
	resp, _ = rm.Handler.Execute()
	if resp.Body != "health" {
		t.Error(errors.New("Expected health"))
	}

	params["proxy"] = "/crud"
	req.HTTPMethod = "POST"
	req.PathParameters = params
	rm, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected POST /crud"))
	}
	resp, _ = rm.Handler.Execute()
	if resp.Body != "update" {
		t.Error(errors.New("Expected update"))
	}

	req.HTTPMethod = "DELETE"
	rm, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected DELETE /crud"))
	}
	resp, _ = rm.Handler.Execute()
	if resp.Body != "delete" {
		t.Error(errors.New("Expected delete"))
	}

	params["proxy"] = "/links/a"
	req.HTTPMethod = "GET"
	req.PathParameters = params
	rm, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /links/{id}"))
	}

	id, ok := rm.Params["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /links/a"))
	}
	resp, _ = rm.Handler.Execute()
	if resp.Body != "links" {
		t.Error(errors.New("Expected links"))
	}

	params["proxy"] = "/links/a/?q=test"
	req.PathParameters = params
	rm, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /links/{id}"))
	}

	id, ok = rm.Params["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /links/a"))
	}
	resp, _ = rm.Handler.Execute()
	if resp.Body != "links" {
		t.Error(errors.New("Expected links"))
	}

	params["proxy"] = "/optional"
	req.HTTPMethod = "GET"
	req.PathParameters = params
	rm, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /optional/{id?a}"))
	}

	id, ok = rm.Params["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /optional/a"))
	}
	resp, _ = rm.Handler.Execute()
	if resp.Body != "optional" {
		t.Error(errors.New("Expected optional"))
	}

	params["proxy"] = "/optional?q=v"
	req.HTTPMethod = "GET"
	req.PathParameters = params
	rm, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /optional/{id?a}"))
	}

	id, ok = rm.Params["id"]
	if !ok || id != "a" {
		t.Error(errors.New("Expected GET /optional/a"))
	}
	resp, _ = rm.Handler.Execute()
	if resp.Body != "optional" {
		t.Error(errors.New("Expected optional"))
	}
}
