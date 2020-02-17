package persist

import (
	"context"
	"crawler/crawler/engine"
	"errors"
	"log"

	"github.com/olivere/elastic/v7"
)

// ItemSaver 将item保存到elasticsearch
func ItemSaver(index string) (saverChan chan engine.Item, err error) {
	var (
		client *elastic.Client
	)

	if client, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL([]string{"http://192.168.94.30:9200"}...),
	); err != nil {
		log.Printf("new elastic client failed, err: %v\r\n", err)
		return nil, err
	}

	saverChan = make(chan engine.Item)

	go func() {
		var count = 0
		for {
			item := <-saverChan
			if err = Save(client, index, item); err != nil {
				log.Printf("save item:%s failed, err: %v\r\n", item.ID, err)
				continue
			}
			count++
			log.Printf("save item%d: %v\r\n", count, item)
		}
	}()

	return
}

// Save 保存
func Save(client *elastic.Client, index string, item engine.Item) (err error) {
	if item.ID == "" {
		return errors.New("item id is empty")
	}

	if _, err = client.Index().Index(index).
		Type(item.Type).
		Id(item.ID).
		BodyJson(item).Do(context.Background()); err != nil {
		log.Printf("save item: %s failed, err: %v\r\n", item.ID, err)
		return err
	}
	return nil
}
