package persist

import (
	"crawler/crawler/engine"
	"crawler/crawler/persist"
	"log"

	"github.com/olivere/elastic/v7"
)

// Service persist service
type Service struct {
	Client *elastic.Client
	Index  string
}

// Save save
func (s *Service) Save(item engine.Item, result *string) (err error) {
	if err = persist.Save(s.Client, s.Index, item); err != nil {
		log.Printf("persist.save failed, err: %v\r\n", err)
		return err
	}
	*result = "ok"
	return nil
}
