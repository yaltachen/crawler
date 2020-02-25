package client

import (
	"crawler/crawler/engine"
	"crawler/crawler_distribute/rpcsupport"
	"log"
	"net/rpc"
)

// ItemSaver itemSaver channel
func ItemSaver(host string) (chan engine.Item, error) {
	var (
		client   *rpc.Client
		itemChan chan engine.Item
		err      error
	)
	if client, err = rpcsupport.NewSupport(host); err != nil {
		return nil, err
	}
	itemChan = make(chan engine.Item)

	go func() {
		var result string
		for {
			select {
			case item := <-itemChan:
				if err = client.Call("Service.Save", item, &result); err != nil {
					log.Printf("call failed, err: %v\r\n", err)
				} else {
					log.Printf("got item: %v\r\n", item)
					// log.Printf("result: %s", result)
				}
			}
		}
	}()

	return itemChan, nil
}
