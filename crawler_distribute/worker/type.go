package worker

import (
	"crawler/crawler/engine"
	"crawler/crawler/parser/zhenai"
	"crawler/crawler_distribute/config"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// SerializedParser 序列化的parser
type SerializedParser struct {
	Name string      `json:"name"`
	Args interface{} `json:"args"`
}

// Request 可序列化的request
type Request struct {
	URL    string           `json:"url"`
	Parser SerializedParser `json:"parser"`
}

// ParserResult 可序列化的ParserResult
type ParserResult struct {
	Items    []engine.Item `json:"items"`
	Requests []Request     `json:"requests"`
}

// SerializeRequest 序列化engine.Request
func SerializeRequest(r engine.Request) Request {
	var (
		name string
		args interface{}
	)
	name, args = r.Parser.Serialize()
	return Request{
		URL:    r.URL,
		Parser: SerializedParser{Name: name, Args: args},
	}
}

// SerializeParserResult 序列化engine.SerializeParserResult
func SerializeParserResult(result engine.ParserResult) ParserResult {
	var (
		parserResult ParserResult
	)

	parserResult.Items = result.Items
	for _, request := range result.Requests {
		parserResult.Requests = append(parserResult.Requests, SerializeRequest(*request))
	}
	return parserResult
}

// DeserializeRequest 反序列化request
func DeserializeRequest(r Request) (engine.Request, error) {
	var (
		parser engine.Parser
		err    error
	)
	if parser, err = deserializeParser(r.Parser); err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		URL:    r.URL,
		Parser: parser,
	}, nil
}

// DeserializeParserResult 反序列化parserresult
func DeserializeParserResult(result ParserResult) (engine.ParserResult, error) {
	var (
		r   engine.ParserResult
		err error
	)
	r.Items = result.Items
	for _, request := range result.Requests {
		var req engine.Request
		if req, err = DeserializeRequest(request); err != nil {
			log.Printf("error deserializing "+
				"request: %v", err)
			continue
		}
		r.Requests = append(r.Requests, &req)
	}
	return r, nil
}

// deserializeParser 反序列化parser
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.CITYLISTPARSER:
		return zhenai.NewCityListParser(), nil
	case config.CITYPARSER:
		if loc, ok := p.Args.(string); ok {
			return zhenai.NewCityParser(loc), nil
		}
		return nil, fmt.Errorf("invalid arg: %v", p.Args)

	case config.PERSONPARSER:
		var person zhenai.PersonParser
		pstring, err := json.Marshal(p.Args.(map[string]interface{}))
		// log.Println("pstring", string(pstring))
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(pstring, &person)
		// log.Printf("%v", person)
		if err != nil {
			return nil, err
		}
		return zhenai.NewPersonParser(person.URL, person.ID, person.Name, person.Gender, person.Loc), nil

	default:
		return nil, errors.New("unknown parser name")
	}
}
