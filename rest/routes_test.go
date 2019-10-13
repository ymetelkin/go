package rest

import (
	"errors"
	"fmt"
	"testing"
)

func Echo(req Request) Response {
	return Response{
		StatusCode: 200,
		Body:       fmt.Sprintf("%s: %s %s", req.HTTPMethod, req.Path, req.Body),
	}
}

func EchoParams(req Request) Response {
	return Response{
		StatusCode: 200,
		Body:       fmt.Sprintf("%s: %s %v", req.HTTPMethod, req.Path, req.PathParameters),
	}
}

func TestRoutes(t *testing.T) {
	rt := Router{}
	rt.Add("/", "GET", Echo)
	rt.Add("/health", "GET", Echo)
	rt.Add("/crud", "POST", Echo)
	rt.Add("/crud", "DELETE", Echo)
	rt.Add("/links/{id}", "GET", EchoParams)
	rt.Add("/optional/{id?a}", "GET", EchoParams)
	rt.Add("/linking/crud", "DELETE", Echo)
	rt.Add("/linking/links/{id}", "GET", EchoParams)
	rt.Add("/linking/{area?links}/{id}/{action?test}", "GET", EchoParams)

	var req = Request{
		HTTPMethod: "GET",
		Path:       "/",
	}

	resp := rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected GET / "))
	}

	req = Request{
		HTTPMethod: "GET",
		Path:       "/health",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected GET /health "))
	}

	req = Request{
		HTTPMethod: "GET",
		Path:       "/foo",
	}

	resp = rt.Execute(req)
	if resp.StatusCode != 404 {
		t.Error(errors.New("Unexpected GET /foo "))
	} else {
		fmt.Println(resp.Body)
	}

	req = Request{
		HTTPMethod: "POST",
		Path:       "/crud",
		Body:       `{"id":1,"name":"YM"}`,
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected POST /crud "))
	}

	req = Request{
		HTTPMethod: "DELETE",
		Path:       "/crud",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected DELETE /crud "))
	}

	req = Request{
		HTTPMethod: "GET",
		Path:       "/links/ym",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected GET /links/ym "))
	}

	req = Request{
		HTTPMethod: "GET",
		Path:       "/optional/b/",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected GET /optional/b "))
	}

	req = Request{
		HTTPMethod: "GET",
		Path:       "/optional/",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected GET /optional/a "))
	}

	req = Request{
		HTTPMethod: "DELETE",
		Path:       "/linking/crud",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected DELETE /linking/crud "))
	}

	req = Request{
		HTTPMethod: "GET",
		Path:       "/linking/links/xyz/",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected GET /linking/links/xyz "))
	}

	req = Request{
		HTTPMethod: "GET",
		Path:       "/linking/s3/xyz/",
	}

	resp = rt.Execute(req)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	} else {
		t.Error(errors.New("Expected GET /linking/s3/xyz "))
	}
}
