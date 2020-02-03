package zhenai

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"regexp"
)

var (
	ageParser = regexp.MustCompile(
		`<div[^>]*>([0-9]*岁)</div>`)
	heightParser = regexp.MustCompile(
		`<div[^>]*>([0-9]*cm)</div>`)
	weightParser = regexp.MustCompile(
		`<div[^>]*>([0-9]*kg)</div>`)
	incomeParser = regexp.MustCompile(
		`<div[^>]*>月收入:([^<]*)</div>`)
	marriageParser = regexp.MustCompile(
		`<div[^>]*>([未已离丧][婚异偶])</div>`)
	educationParser = regexp.MustCompile(
		`<div[^>]*>([小中高大硕博][中学专士][本及]{0,1}[科以]{0,1}[下]{0,1})</div>`)
	xingzuoParser = regexp.MustCompile(
		`<div[^>]*>(.{2,2}座)[^<]*</div>`)
	carParser = regexp.MustCompile(
		`<div[^>]*>(.{1,2}车)</div>`)
	houseParser = regexp.MustCompile(
		`<div[^>]*>(.{1,2}房)</div>`)
	jiguanParser = regexp.MustCompile(
		`<div[^>]*>籍贯:([^<]*)</div>`)
)

// PersonParser 用户资料解析器
func PersonParser(url string, id string, name, gender, loc string, content []byte) (parserResult *engine.ParserResult, err error) {
	var (
		profile   model.Profile
		age       string
		weight    string
		height    string
		house     string
		car       string
		jiguan    string
		income    string
		education string
		marriage  string
		xingzuo   string
	)
	age = getArgs(ageParser.FindStringSubmatch(string(content)))
	height = getArgs(heightParser.FindStringSubmatch(string(content)))
	weight = getArgs(weightParser.FindStringSubmatch(string(content)))
	house = getArgs(houseParser.FindStringSubmatch(string(content)))
	car = getArgs(carParser.FindStringSubmatch(string(content)))
	jiguan = getArgs(jiguanParser.FindStringSubmatch(string(content)))
	income = getArgs(incomeParser.FindStringSubmatch(string(content)))
	marriage = getArgs(marriageParser.FindStringSubmatch(string(content)))
	education = getArgs(educationParser.FindStringSubmatch(string(content)))
	xingzuo = getArgs(xingzuoParser.FindStringSubmatch(string(content)))
	profile = model.Profile{
		URL:  url,
		ID:   id,
		Type: "zhenai",
		Payload: model.ZhenaiPayLoad{
			Name:      name,
			Gender:    gender,
			Age:       age,
			Height:    height,
			Weight:    weight,
			Income:    income,
			Marriage:  marriage,
			Education: education,
			Location:  loc,
			JiGuan:    jiguan,
			XingZuo:   xingzuo,
			House:     house,
			Car:       car,
		},
	}
	parserResult = &engine.ParserResult{}
	parserResult.Items = append(parserResult.Items, profile)
	return
}

func getArgs(arg []string) string {
	if len(arg) == 0 {
		return "未知"
	}
	return arg[1]
}
