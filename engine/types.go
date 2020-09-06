package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Item     []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
