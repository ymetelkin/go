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

	params := make(map[string]string)

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
	}
	h, ok := pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /"))
	}
	resp, _ := h.Execute()
	if resp.Body != "root" {
		t.Error(errors.New("Expected root"))
	}

	params["proxy"] = "/health"
	req.PathParameters = params
	h, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected GET /helath"))
	}
	resp, _ = h.Execute()
	if resp.Body != "health" {
		t.Error(errors.New("Expected health"))
	}

	params["proxy"] = "/crud"
	req.HTTPMethod = "POST"
	req.PathParameters = params
	h, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected POST /crud"))
	}
	resp, _ = h.Execute()
	if resp.Body != "update" {
		t.Error(errors.New("Expected update"))
	}

	req.HTTPMethod = "DELETE"
	req.PathParameters = params
	h, ok = pr.GetHandler(req)
	if !ok {
		t.Error(errors.New("Expected DELETE /crud"))
	}
	resp, _ = h.Execute()
	if resp.Body != "delete" {
		t.Error(errors.New("Expected delete"))
	}
}
