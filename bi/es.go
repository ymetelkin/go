package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type indexer struct {
	Cluster string
	Index   string
	Size    int
}

type response struct {
	Items []struct {
		Index struct {
			ID    string `json:"_id"`
			Error struct {
				Reason   string `json:"reason"`
				CausedBy struct {
					Reason string `json:"reason"`
				} `json:"caused_by"`
			} `json:"error"`
		} `json:"index"`
	} `json:"items"`
}

func newIndexer(cluster string) indexer {
	return indexer{
		Cluster: strings.TrimRight(cluster, "/"),
		Index:   "business-index",
		Size:    1000,
	}
}

func (es *indexer) reindex(items map[string]string) (int, error) {
	var (
		i, total int
		bks      []string
		sb       strings.Builder
		rs       response
	)

	var lf rune = 10

	for k, v := range items {
		sb.WriteString(fmt.Sprintf(`{"index":{"_index":"%s","_type":"doc","_id":"%s"}}`, es.Index, k))
		sb.WriteRune(lf)
		sb.WriteString(v)
		sb.WriteRune(lf)

		i++
		if i == es.Size {
			bks = append(bks, sb.String())
			sb = strings.Builder{}
			i = 0
		}

		total++

		//fmt.Println(sb.String())
	}

	if i > 0 {
		bks = append(bks, sb.String())
	}

	url := es.Cluster + "/_bulk"
	client := &http.Client{}

	for _, bk := range bks {
		//fmt.Println(bk)
		bs := []byte(bk)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))
		if err != nil {
			return 0, err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()

		bs, _ = ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal(bs, &rs); err != nil {
			return 0, err
		}

		if rs.Items != nil && len(rs.Items) > 0 {
			for _, it := range rs.Items {
				if it.Index.Error.Reason != "" {
					total--
					fmt.Printf("Failed to import %s: %s: %s\n", it.Index.ID, it.Index.Error.Reason, it.Index.Error.CausedBy.Reason)
				}
			}
		}
	}

	return total, nil
}
