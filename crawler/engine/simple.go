package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

// SimpleEngine 单机 单线程
type SimpleEngine struct{}

// Run 启动engine
func (SimpleEngine) Run(seeds ...*Request) {
	var (
		requests []*Request
		content  []byte
		result   *ParserResult
		count    int
		err      error
	)

	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		if content, err = fetcher.Fetch(r.URL); err != nil {
			log.Printf("Fetcher: error fetching %s: %v\r\n", r.URL, err)
		}

		if result, err = r.Parser.Parse(content); err != nil {
			log.Printf("Parser: error parsing %s: %v\r\n", r.URL, err)
		}

		for _, item := range result.Items {
			count++
			log.Printf("got item%d: %v", count, item)
		}

		requests = append(requests, result.Requests...)

	}
}
