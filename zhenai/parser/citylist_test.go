package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}
	cityParseResult := ParseCityList(content)

	const resultSize = 470

	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectCities = []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	if len(cityParseResult.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"request; but had %d ", resultSize, len(cityParseResult.Requests))
	}

	for i, url := range expectedUrls {
		if cityParseResult.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but"+
				"was %s", i, cityParseResult.Requests[i].Url, url)
		}
	}
	for i, city := range expectCities {
		if cityParseResult.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but"+
				"was %s", i, cityParseResult.Items[i], city)
		}
	}

	if len(cityParseResult.Items) != resultSize {
		t.Errorf("result should have %d "+
			"Items; but had %d ", resultSize, len(cityParseResult.Items))
	}
}
