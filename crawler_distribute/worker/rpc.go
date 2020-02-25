package worker

import (
	"crawler/crawler/engine"
)

// CrawlerService 爬虫服务
type CrawlerService struct {
}

// Process 处理request
func (CrawlerService) Process(req Request, result *ParserResult) error {
	// log.Println("some one call crawler service ")
	// log.Printf("%v", req)
	eReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	// log.Printf("%v", eReq)
	eResult, err := engine.DoWork(&eReq)
	if err != nil {
		return err
	}
	*result = SerializeParserResult(*eResult)
	return nil
}
