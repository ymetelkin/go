package apql

type set struct {
	And     bool
	Not     bool
	Queries []query
}

type query struct {
	Field    string
	Operator string
	Value    string
}
