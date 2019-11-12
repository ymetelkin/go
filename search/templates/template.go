package templates

import "github.com/ymetelkin/go/json"

//SearchTemplate helps to generate ES search request
type SearchTemplate struct {
	ID           string
	Name         string
	Description  string
	Index        string
	FieldAliases map[string]string
	Request      json.Object
}

//New tries to load a search template from a repository
func New(id string) (st SearchTemplate, err error) {

	return
}
