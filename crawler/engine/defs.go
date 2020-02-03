package engine

// Request request结构体
// 包含网页url，和对应的解析器
type Request struct {
	URL    string
	Parser ParserFunc
}

// ParserResult 解析结果
// 包含新的Request 和 解析结果
type ParserResult struct {
	Items    []interface{}
	Requests []*Request
}

// ParserFunc 解析器类型
// 接收字节数组，返回解析结果
type ParserFunc func([]byte) (*ParserResult, error)

// NilParser 空解析器
func NilParser([]byte) (*ParserResult, error) {
	return &ParserResult{}, nil
}
