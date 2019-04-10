package appl

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type texts struct {
	props []json.Property
}

func (txts *texts) ParseTextComponent(pc pubcomponent) error {
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

	jp := json.NewObjectProperty(strings.ToLower(pc.Role), jo)

	if txts.props == nil {
		txts.props = []json.Property{jp}
	} else {
		txts.props = append(txts.props, jp)
	}

	return nil
}

func (txts *texts) AddProperties(jo *json.Object) {
	if txts.props != nil {
		for _, jp := range txts.props {
			jo.AddProperty(jp)
		}
	}
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
