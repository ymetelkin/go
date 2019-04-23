package links

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type db struct {
	ID    string
	Table string
}

func newDb() db {
	return db{ID: "ID", Table: "LinkCollections"}
}

func (db *db) SaveCollection(col Collection) error {
	item, err := dynamodbattribute.MarshalMap(col)
	if err != nil {
		return err
	}

	pi := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(db.Table),
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	_, err = svc.PutItem(pi)
	if err != nil {
		return err
	}

	return nil
}

func (db *db) GetCollection(id string) (Collection, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	col := Collection{}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.Table),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return col, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &col)
	if err != nil {
		return col, err
	}

	return col, nil
}
