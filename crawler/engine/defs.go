package engine

// Parser 接口
type Parser interface {
	Parse(content []byte) (result *ParserResult, err error)
	Serialize() (funcName string, args interface{})
}

// Request request结构体
// 包含网页url，和对应的解析器
type Request struct {
	URL    string
	Parser Parser
}

// ParserResult 解析结果
// 包含新的Request 和 解析结果
type ParserResult struct {
	Items    []Item
	Requests []*Request
}

// Item item
type Item struct {
	ID      string      `json:"id"`
	URL     string      `json:"url"`
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// ParserFunc 解析器类型
// 接收字节数组，返回解析结果
type ParserFunc func([]byte) (*ParserResult, error)

// NilParser 空解析器
// func NilParser([]byte) (*ParserResult, error) {
// 	return &ParserResult{}, nil
// }

// NilParser 空解析器
type NilParser struct {
}

// Parse 解析
func (p NilParser) Parse(content []byte) (result *ParserResult, err error) {
	// NilParser do nothing
	return nil, nil
}

// Serialize 序列化
func (p NilParser) Serialize() (funcName string, args interface{}) {
	return "NilParser", nil
}
