package zhenai

import (
	"crawler/crawler/engine"
	"regexp"
	"strings"
)

// 城市列表正则表达式
var personReg = regexp.MustCompile(
	`(http://album.zhenai.com/u/[0-9]*[^"])"[^>]*>([^<]*)</a>[^性]*性别[^男女]*([男女])士`)

// 相关城市正则
var relateCityReg = regexp.MustCompile(
	`(http://www.zhenai.com/zhenghun/[^"/]*)">([^征<]*)征婚`)

// 下一页
var nextPage = regexp.MustCompile(
	`(http://www.zhenai.com/zhenghun/[^/]*/[0-9^"]*)">下一页`)

// CityParser 城市解析器
func CityParser(loc string, content []byte) (parserResult *engine.ParserResult, err error) {
	var (
		citysMatch       [][]string
		nextPageMatch    [][]string
		relateCitysMatch [][]string
		requests         []*engine.Request
		// items            []interface{}
	)

	requests = make([]*engine.Request, 0)
	// items = make([]interface{}, 0)

	// 城市解析
	citysMatch = personReg.FindAllStringSubmatch(string(content), -1)
	for _, city := range citysMatch {
		name := city[2]
		gender := city[3]
		url := city[1]
		id := url[strings.LastIndex(url, "/")+1:]
		requests = append(requests, &engine.Request{
			URL: url,
			Parser: func(content []byte) (*engine.ParserResult, error) {
				return PersonParser(url, id, name, gender, loc, content)
			},
		})
		// items = append(items, &model.CityItem{
		// 	ID:     id,
		// 	URL:    url,
		// 	Name:   name,
		// 	Gender: gender,
		// })
	}

	// 下一页
	nextPageMatch = nextPage.FindAllStringSubmatch(string(content), 1)
	for _, city := range nextPageMatch {
		url := city[1]
		requests = append(requests, &engine.Request{
			URL: url,
			Parser: func(content []byte) (*engine.ParserResult, error) {
				return CityParser(loc, content)
			},
		})
	}

	// 相关城市
	relateCitysMatch = relateCityReg.FindAllStringSubmatch(string(content), -1)
	for _, city := range relateCitysMatch {
		url := city[1]
		loc := city[2]
		requests = append(requests, &engine.Request{
			URL: url,
			Parser: func(content []byte) (*engine.ParserResult, error) {
				return CityParser(loc, content)
			},
		})
	}

	parserResult = &engine.ParserResult{
		Items:    make([]engine.Item, 0),
		Requests: requests,
	}

	return parserResult, nil
}
