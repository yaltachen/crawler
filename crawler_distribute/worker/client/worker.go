package client

import (
	"crawler/crawler/engine"
	"crawler/crawler_distribute/worker"
	"log"
	"net/rpc"
)

// CreateProcessor 创建worker
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	var (
		// client *rpc.Client
		err error
	)
	// if client, err = rpcsupport.NewSupport(
	// 	fmt.Sprintf(":%s", config.WorkerPort0)); err != nil {
	// 	return nil, err
	// }

	return func(
		req *engine.Request) (*engine.ParserResult, error) {
		var (
			sReq    worker.Request
			sResult worker.ParserResult
			result  engine.ParserResult
		)
		sReq = worker.SerializeRequest(*req)
		c := <-clientChan
		if err = c.Call("CrawlerService.Process",
			sReq, &sResult); err != nil {
			log.Printf("call process failed, err: %v\r\n", err)
			return nil, err
		}
		if result, err = worker.DeserializeParserResult(sResult); err != nil {
			log.Printf("DeserializeParserResult failed, err: %v\r\n", err)
			return nil, err
		}
		return &result, err
	}

}
