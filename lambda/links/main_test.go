package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test(t *testing.T) {
	body := `
	{
		"doc": {
			"id": "c37efb6311d54ffa88546ea543a8835b",
			"href": "http://proteus-searchapi-us-east-1.aptechdevlab.com/api/appl/c37efb6311d54ffa88546ea543a8835b"
		},
		"link":{
			"id": "ec13ae10caca4760a75ce66f00b970e7",
			"href":"http://proteus-searchapi-us-east-1.aptechdevlab.com/api/appl/ec13ae10caca4760a75ce66f00b970e7"
		},
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

	req = events.APIGatewayProxyRequest{
		Path:       "/links/c37efb6311d54ffa88546ea543a8835b",
		HTTPMethod: "GET",
		Body:       body,
	}
	res, _ = execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ = json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))

	req = events.APIGatewayProxyRequest{
		Path:       "/docs/ec13ae10caca4760a75ce66f00b970e7",
		HTTPMethod: "GET",
		Body:       body,
	}
	res, _ = execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ = json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))

	body = `
	{
		"doc_id": "c37efb6311d54ffa88546ea543a8835b",
		"link_id": "ec13ae10caca4760a75ce66f00b970e7",
		"user_id": "YM",
		"seq": 0
	}`

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

func TestReset(t *testing.T) {
	body := `
	{
		"doc": {
			"id": "c37efb6311d54ffa88546ea543a8835b",
			"href": "http://proteus-searchapi-us-east-1.aptechdevlab.com/api/appl/c37efb6311d54ffa88546ea543a8835b"
		},
		"links": [
			{ 
				"id": "256817da7fee4afbae5dedbaa3991d33",
				"href":"http://proteus-searchapi-us-east-1.aptechdevlab.com/api/appl/256817da7fee4afbae5dedbaa3991d33"
			},
			{ 
				"id": "ec13ae10caca4760a75ce66f00b970e7",
				"href":"http://proteus-searchapi-us-east-1.aptechdevlab.com/api/appl/ec13ae10caca4760a75ce66f00b970e7"
			}
		],    
		"user_id": "YM"
	  }`

	req := events.APIGatewayProxyRequest{
		Path:       "/reset",
		HTTPMethod: "POST",
		Body:       body,
	}
	res, _ := execute(req)
	fmt.Printf("Status: %v\n", res.StatusCode)
	bytes, _ := json.MarshalIndent(res, "", "   ")
	fmt.Println(string(bytes))
}
