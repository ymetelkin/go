package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test(t *testing.T) {
	os.Setenv("ES", "http://proteus-int-all-esclient.aptechdevlab.com:9200")

	body := `
	{
		"doc_id": "9b2ca4c1f974e97ae156cd85d26cdea8",
		"link_id": "a237bc351d894948a00c8da9bcb7fe1e",
		"user_id": "YM",
		"seq": 0
	}`

	req := events.APIGatewayProxyRequest{
		Path:       "/",
		HTTPMethod: "PUT",
		Body:       body,
	}
	res, _ := execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ := json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))

	query := make(map[string]string)
	query["fields"] = "headline,type"
	query["uid"] = "YM"
	req = events.APIGatewayProxyRequest{
		Path:                  "/links/9b2ca4c1f974e97ae156cd85d26cdea8?fields=headline,type&uid=YM",
		QueryStringParameters: query,
		HTTPMethod:            "GET",
		Body:                  body,
	}
	res, _ = execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ = json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))

	req = events.APIGatewayProxyRequest{
		Path:                  "/docs/a237bc351d894948a00c8da9bcb7fe1e?fields=headline,type&uid=YM",
		QueryStringParameters: query,
		HTTPMethod:            "GET",
		Body:                  body,
	}
	res, _ = execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ = json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))

	req = events.APIGatewayProxyRequest{
		Path:       "/",
		HTTPMethod: "POST",
		Body:       body,
	}
	res, _ = execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ = json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))

	req = events.APIGatewayProxyRequest{
		Path:       "/",
		HTTPMethod: "DELETE",
		Body:       body,
	}
	res, _ = execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ = json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))
}
