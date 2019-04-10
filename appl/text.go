package appl

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

func (doc *document) ParseTextComponent(pc pubcomponent, parent *json.Object) error {
	nd := pc.Node.GetNode("DataContent")
	nd = nd.GetNode("nitf")
	if nd.Nodes == nil {
		return errors.New("[nitf] block is missing")
	}

	nd = getBodyContent(nd)
	if nd.Nodes == nil {
		return errors.New("[body.content] block is missing")
	}

	var sb strings.Builder

	for _, b := range nd.Nodes {
		if b.Name == "block" {
			s := b.ToInlineString()
			sb.WriteString(s)
		}
	}

	nitf := sb.String()
	if nitf == "" {
		return errors.New("[block] blocks are missing or empty")
	}

	jo := json.Object{}
	jo.AddString("nitf", nitf)

	nd = pc.Node.GetNode("Characteristics")
	nd = nd.GetNode("Words")
	i, err := strconv.Atoi(nd.Text)
	if err == nil {
		jo.AddInt("words", i)
	}

	parent.AddObject(strings.ToLower(pc.Role), jo)

	return nil
}

func getBodyContent(nd xml.Node) xml.Node {
	if nd.Nodes != nil {
		for _, n := range nd.Nodes {
			if n.Name == "body.content" {
				return n
			}

			test := getBodyContent(n)
			if test.Nodes != nil {
				return test
			}
		}
	}
	return xml.Node{}
}
