package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3doc struct {
	ID   string
	Body string
}

func getS3Docs(bucket string) ([]s3doc, error) {
	var (
		docs []s3doc
		err  error
	)

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	svc := s3.New(sess)

	prefix := "sstemplates/doc"

	req := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}

	err = svc.ListObjectsPages(req, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			input := &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(*obj.Key),
			}

			result, err := svc.GetObject(input)
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			defer result.Body.Close()

			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, result.Body); err != nil {
				fmt.Println(err.Error())
				break
			}
			bytes := buf.Bytes()

			var obj map[string]interface{}

			if err := json.Unmarshal(bytes, &obj); err != nil {
				fmt.Println(err.Error())
				break
			}

			v, ok := obj["id"]
			if !ok {
				fmt.Println("No ID found")
				break
			}
			id, _ := v.(string)

			bytes, err = json.Marshal(obj)
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			doc := s3doc{
				ID:   id,
				Body: string(bytes),
			}
			docs = append(docs, doc)
		}
		return true
	})

	if err != nil {
		fmt.Println("failed to list objects", err)
	}

	return docs, err
}
