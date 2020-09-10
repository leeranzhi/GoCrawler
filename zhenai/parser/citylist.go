package parser

import (
	"GoCrawler/engine"
	"regexp"
)

const cityListReString = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

var cityListRe = regexp.MustCompile(cityListReString)

/**
解析城市列表信息
*/
func ParseCityList(content []byte) engine.ParserResult {
	//re := regexp.MustCompile(cityListRe)
	matches := cityListRe.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}

	//limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
