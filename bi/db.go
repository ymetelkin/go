package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type db struct {
	Table string
	Env   string
}

func (db *db) list() (map[string]string, error) {
	items := make(map[string]string)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", db.Env),
	})
	if err != nil {
		return items, err
	}

	dn := dynamodb.New(sess)

	table := db.Table

	params := &dynamodb.ScanInput{TableName: aws.String(table)}
	result, err := dn.Scan(params)
	if err != nil {
		return items, err
	}

	for _, item := range result.Items {
		obj := make(map[string]interface{})
		err = dynamodbattribute.UnmarshalMap(item, &obj)
		if err != nil {
			return items, err
		}

		t, ok := obj["BO_Schema"]
		if ok {
			s, ok := t.(string)
			if ok {
				bytes := []byte(s)

				var js map[string]interface{}

				if err := json.Unmarshal(bytes, &js); err != nil {
					return items, err
				}

				o, ok := js["boId"]
				if !ok {
					return items, fmt.Errorf("boId is missing")
				}

				id, ok := o.(string)
				if !ok {
					return items, fmt.Errorf("boId is not string")
				}

				bytes, err = json.Marshal(js)
				if err := json.Unmarshal(bytes, &js); err != nil {
					return items, err
				}

				s = string(bytes)
				items[id] = s
			}
		}
	}

	return items, err
}
