package zhenai

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"regexp"
	"strings"
)

// 城市列表正则表达式
// [1]: url
// [2]: loc
var cityReg = regexp.MustCompile(
	`href="(http://www.zhenai.com/zhenghun/[^"]*)"[^>]*>([^<]*)</a>`)

var cityList = make(map[string]bool)

// CityListParser 城市列表解析器
func CityListParser(content []byte) (parserResult *engine.ParserResult, err error) {
	var (
		citysMatch [][]string
		requests   []*engine.Request
		items      []interface{}
	)

	citysMatch = cityReg.FindAllStringSubmatch(string(content), -1)
	requests = make([]*engine.Request, 0)
	items = make([]interface{}, 0)

	count := 0

	for _, city := range citysMatch {
		if strings.Contains(city[2], "征婚") {
			continue
		}

		if cityList[city[1]] == true {
			continue
		}
		count++
		if count >= 5 {
			break
		}
		loc := city[2]
		requests = append(requests, &engine.Request{
			URL: city[1],
			Parser: func(content []byte) (*engine.ParserResult, error) {
				return CityParser(loc, content)
			},
		})

		cityList[city[1]] = true

		items = append(items, &model.CityListItem{URL: city[1], Loc: city[2]})
	}
	parserResult = &engine.ParserResult{
		Items:    items,
		Requests: requests,
	}
	return parserResult, err
}
