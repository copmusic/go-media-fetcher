package query_parsers

import "fmt"

type Query string
type MediaId string

type QueryParseError struct {
	error string
}

func (e *QueryParseError) Error() string {
	return fmt.Sprintf("%s", e.Error())
}

type QueryParser interface {
	Support(query Query) bool
	GetMediaId(query Query) (MediaId, error)
}
