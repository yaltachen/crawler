package model

import "encoding/json"

// CityListItem 城市列表解析结果
type CityListItem struct {
	Loc string
	URL string
}

// CityItem 城市解析结果
type CityItem struct {
	ID     string
	Name   string
	Gender string
	URL    string
}

// ZhenaiPayLoad 珍爱网payLoad
type ZhenaiPayLoad struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       string `json:"age"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Income    string `json:"income"`
	Marriage  string `json:"marriage"`
	Education string `json:"education"`
	Location  string `json:"location"`
	JiGuan    string `json:"jiguan"`
	XingZuo   string `json:"xingzuo"`
	House     string `json:"house"`
	Car       string `json:"car"`
}

// FromJSONObj json to obj
func FromJSONObj(o interface{}) (ZhenaiPayLoad, error) {
	var profile ZhenaiPayLoad
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
