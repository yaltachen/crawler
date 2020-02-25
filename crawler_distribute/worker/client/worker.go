package client

import (
	"crawler/crawler/engine"
	"crawler/crawler_distribute/config"
	"crawler/crawler_distribute/rpcsupport"
	"crawler/crawler_distribute/worker"
	"fmt"
	"net/rpc"
)

// CreateProcessor 创建worker
func CreateProcessor() (engine.Processor, error) {
	var (
		client *rpc.Client
		err    error
	)
	if client, err = rpcsupport.NewSupport(
		fmt.Sprintf(":%s", config.WorkerPort0)); err != nil {
		return nil, err
	}

	return func(
		req *engine.Request) (*engine.ParserResult, error) {
		var (
			sReq    worker.Request
			sResult worker.ParserResult
			result  engine.ParserResult
		)
		sReq = worker.SerializeRequest(*req)
		if err = client.Call("CrawlerService.Process",
			sReq, &sResult); err != nil {
			return nil, err
		}
		result, err = worker.DeserializeParserResult(sResult)
		return &result, err
	}, nil

}
