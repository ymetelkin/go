package es

type searchRequest struct {
	Index  string
	Query  string
	Source []string
	Size   int
	From   int
}

func newSearchRequest(index string) searchRequest {
	return searchRequest{Index: index}
}

func (sr *searchRequest) SetQuery(query string) *searchRequest {
	sr.Query = query
	return sr
}

func (sr *searchRequest) SetSource(source []string) *searchRequest {
	sr.Source = source
	return sr
}

func (sr *searchRequest) SetSize(size int) *searchRequest {
	sr.Size = size
	return sr
}

func (sr *searchRequest) SetFrom(from int) *searchRequest {
	sr.From = from
	return sr
}
