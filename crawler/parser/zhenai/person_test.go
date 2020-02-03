package zhenai

import (
	"crawler/crawler/fetcher"
	"crawler/crawler/model"
	"testing"
)

func TestPersonParser(t *testing.T) {
	var (
		url     = "https://album.zhenai.com/u/1513404786"
		id      = "1513404786"
		name    = "心语"
		gender  = "女"
		loc     = "北京"
		content []byte
	)

	item := model.Profile{
		ID:   "1513404786",
		URL:  "https://album.zhenai.com/u/1513404786",
		Type: "zhenai", Payload: model.ZhenaiPayLoad{
			Name:      "心语",
			Gender:    "女",
			Age:       "44岁",
			Height:    "162cm",
			Weight:    "60kg",
			Income:    "3千以下",
			Marriage:  "离异",
			Education: "高中及以下",
			Location:  "北京",
			JiGuan:    "黑龙江齐齐哈尔",
			XingZuo:   "魔羯座",
			House:     "租房",
			Car:       "未买车"}}

	content, _ = fetcher.Fetch(url)
	result, err := PersonParser(url, id, name, gender, loc, content)
	if err != nil {
		t.Errorf("parse person failed, err: %v", err)
	}

	if item != result.Items[0] {
		t.Errorf("should get %#v, but got %#v", item, result.Items[0])
	}
}
