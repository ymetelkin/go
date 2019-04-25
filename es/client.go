package es

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

type client struct {
	ClusterURL string
	Client     *elasticsearch.Client
}

func newClient(clusterURL string) (client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			clusterURL,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return client{}, err
	}

	return client{ClusterURL: clusterURL, Client: es}, nil
}

func (ec *client) Search(sr searchRequest) (io.ReadCloser, error) {
	req := []func(*esapi.SearchRequest){
		ec.Client.Search.WithContext(context.Background()),
		ec.Client.Search.WithIndex(sr.Index),
		ec.Client.Search.WithPretty(),
	}

	if sr.Query != "" {
		req = append(req, ec.Client.Search.WithBody(strings.NewReader(sr.Query)))
	}
	if sr.Source != nil {
		req = append(req, ec.Client.Search.WithSource(sr.Source...))
	}
	if sr.Size > 0 {
		req = append(req, ec.Client.Search.WithSize(sr.Size))
	}
	if sr.From > 0 {
		req = append(req, ec.Client.Search.WithFrom(sr.From))
	}

	res, err := ec.Client.Search(req...)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		defer res.Body.Close()

		var body map[string]interface{}

		if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		}

		return nil, fmt.Errorf("[%s] %s: %s",
			res.Status(),
			body["error"].(map[string]interface{})["type"],
			body["error"].(map[string]interface{})["reason"],
		)
	}

	return res.Body, nil
}
