package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/ymetelkin/go/json"
)

func fixBI() error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamodb.New(sess)

	table := "apnews-dev-apcapdevelopment-us-east-1-BusinessObjects"

	params := &dynamodb.ScanInput{TableName: aws.String(table)}
	result, err := db.Scan(params)
	if err != nil {
		return err
	}

	var i int

	for _, item := range result.Items {
		obj := make(map[string]interface{})
		err = dynamodbattribute.UnmarshalMap(item, &obj)
		if err != nil {
			return err
		}

		t, ok := obj["BO_ID"]
		if ok {
			id, ok := t.(string)
			if ok {
				t, ok = obj["BO_Schema"]
				if ok {
					s, ok := t.(string)
					if ok {
						jo, err := json.ParseJSONObject(s)
						if err != nil {
							return err
						}

						q, err := jo.GetString("query")
						if q != "" {
							fx, ok := fixQuery(q)
							if ok {
								i++
								fmt.Printf("%d\t%s\n\t%s\n\n", i, q, fx)

								jo.SetString("query", fx)

								input := &dynamodb.UpdateItemInput{
									TableName: aws.String(table),
									Key: map[string]*dynamodb.AttributeValue{
										"BO_ID": {
											S: aws.String(id),
										},
									},
									UpdateExpression: aws.String("set BO_Schema = :bi"),
									ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
										":bi": {
											S: aws.String(jo.InlineString()),
										},
									},
								}

								_, err := db.UpdateItem(input)
								if err != nil {
									fmt.Println(err.Error())
									return err
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}
