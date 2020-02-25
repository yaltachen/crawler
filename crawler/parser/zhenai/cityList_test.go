package zhenai

import (
	"crawler/crawler/fetcher"
	"testing"
)

func TestCityListParser(t *testing.T) {
	var testURL = "http://www.zhenai.com/zhenghun/"
	content, _ := fetcher.Fetch(testURL)
	result, err := cityListParser(content)
	if err != nil {
		t.Errorf("city list parse failed, err: %v", err)
	}
	if len(result.Items) != 470 {
		t.Errorf("should got 470 citys, but got %d citys", len(result.Items))
	}
}
