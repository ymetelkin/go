package lambda

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type canary struct {
	ActualVersion string
	TestVersion   string
	Error         error
}

//Success returns positive response
func Success(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "application/json"},
		StatusCode: 200,
		Body:       body,
	}, nil
}

//Failure returns response, error and logs ERROR
func Failure(status int, err error, fatal bool) (events.APIGatewayProxyResponse, error) {
	var extra string

	body := err.Error()

	if fatal {
		c := getCanary()
		if c.Error != nil {
			body = fmt.Sprintf("%s; Canary: %s", body, c.Error.Error())
		}

		if c.TestVersion != "" {
			extra = fmt.Sprintf("AdditionalVersion:%s ActiveVersion:%s ", c.TestVersion, c.ActualVersion)
		}
	} else {
		err = nil
	}

	log.Printf("ERROR %s%s\n", extra, strings.ReplaceAll(body, "\"", "\\\""))

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "application/json"},
		StatusCode: status,
		Body:       fmt.Sprintf(`{"error":"%s"}`, body),
	}, err
}

func getCanary() canary {
	f := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	if f == "" {
		return canary{
			Error: errors.New("Missing AWS_LAMBDA_FUNCTION_NAME env"),
		}
	}

	a := os.Getenv("alias")
	if a == "" {
		return canary{
			Error: errors.New("Missing alias env"),
		}
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := lambda.New(sess)
	input := &lambda.GetAliasInput{
		FunctionName: aws.String(f),
		Name:         aws.String(a),
	}

	result, err := svc.GetAlias(input)
	if err != nil {
		return canary{
			Error: err,
		}
	}

	if result.RoutingConfig != nil && result.RoutingConfig.AdditionalVersionWeights != nil {
		for k := range result.RoutingConfig.AdditionalVersionWeights {
			v := *result.FunctionVersion
			if k != v {
				return canary{
					ActualVersion: v,
					TestVersion:   k,
				}
			}
		}
	}

	return canary{}
}
