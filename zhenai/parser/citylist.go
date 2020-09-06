package parser

import (
	"GoCrawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

var cityRe = regexp.MustCompile(cityListRe)

/**
解析城市列表信息
*/
func ParseCityList(content []byte) engine.ParserResult {
	//re := regexp.MustCompile(cityListRe)
	matches := cityRe.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Item = append(result.Item, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		//fmt.Printf("City: %s, Url: %s\n", m[2], m[1])
		//for _, subMatch := range m {
		//	fmt.Printf("%s ",subMatch)
		//}
		//fmt.Printf("%s\n", m)
	}
	//fmt.Printf("Match2 found: %d\n", len(matches))
	return result
}
