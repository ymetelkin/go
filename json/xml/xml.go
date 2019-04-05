package xml

import (
	"errors"
	"fmt"
)

const (
	TOKEN_SPACE rune = 32
	TOKEN_CLOSE rune = 47
	TOKEN_LT    rune = 60
	TOKEN_GT    rune = 62
)

type Node struct {
	Name       string
	Text       string
	Nodes      []*Node
	Attributes []Attribute
}

type Attribute struct {
	Name  string
	Value string
}

func New(s string) (*Node, error) {
	if s == "" {
		return nil, errors.New("Missing XML")
	}

	runes := []rune(s)
	size := len(runes)
	i := 0
	for i < size {
		r := runes[i]
		if r == TOKEN_LT {
			nd, closed, i, err := parseNode(runes, i+1, size)
			if err != nil {
				return nil, err
			}
		}
	}

	return nil, nil
}

func parseNode(runes []rune, i int, size int) (*Node, bool, int, error) {
	start := i
	closed := false
	nd := Node{}
	var attr Attribute

	for i < size {
		r := runes[i]
		if r == TOKEN_SPACE {
			if nd.Name == "" {
				nd.Name = string(runes[start:i])
			} else {
				if attr.Name != "" {
					if nd.Attributes == nil {
						nd.Attributes = []Attribute{attr}
					} else {
						nd.Attributes = append(nd.Attributes, attr)
					}
					attr.Name = ""
				}
			}
		} else if r == TOKEN_GT {
			return &nd, closed, i + 1, nil
		} else if r == TOKEN_CLOSE {
			closed = true
		} else {

		}
	}

	if i > size {
		i = size
	}
	err := fmt.Sprintf("Failed to parse <%s>", string(runes[start:i]))
	return nil, closed, i, errors.New(err)
}
