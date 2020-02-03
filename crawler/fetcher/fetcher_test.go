package fetcher

import (
	"testing"
)

func TestFetch(t *testing.T) {
	var testURL = "https://album.zhenai.com/u/1513404786"
	var start = "<!DOCTYPE html>"
	var end = "</html>"

	content, err := Fetch(testURL)
	if err != nil {
		t.Errorf("fetch failed, err: %v", err)
	}

	if string(content[0:len(start)]) != start {
		t.Errorf("fetch failed, should start with %s, but got %s", start, content[0:len(start)])
	}

	if string(content[len(content)-len(end):]) != end {
		t.Errorf("fetch failed, should end with %s, but got %s", end, string(content[len(content)-len(end):]))
	}

}
