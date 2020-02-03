package model

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

// Profile 存储到es的内容
type Profile struct {
	ID      string      `json:"id"`
	URL     string      `json:"url"`
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
