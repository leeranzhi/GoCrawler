package parser

import (
	"GoCrawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1"+
			"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age:        34,
		Marriage:   "离异",
		Xingzuo:    "魔羯座(12.22-01.19)",
		Height:     163,
		Education:  "大学本科",
		Name:       "格格",
		Car:        "未买车",
		Income:     "3-5千",
		Occupation: "销售总监",
		Hokou:      "黑龙江绥化",
		House:      "已购房",
		Gender:     "女士",
	}
	if profile != expected {
		t.Errorf("profile: %v", profile)
		t.Errorf("expected: %v", expected)
	}
}
