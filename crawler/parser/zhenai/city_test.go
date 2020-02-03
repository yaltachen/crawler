package zhenai

import (
	"crawler/crawler/fetcher"
	"testing"
)

func TestCityParser(t *testing.T) {
	var testURL = "http://www.zhenai.com/zhenghun/beijing"
	content, _ := fetcher.Fetch(testURL)
	result, err := CityParser("北京", content)
	if err != nil {
		t.Errorf("parse city failed, err: %v", err)
	}
	if len(result.Items) != 20 {
		t.Errorf("should got 20 items, but got %d", len(result.Items))
	}

}
