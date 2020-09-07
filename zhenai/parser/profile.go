package parser

import (
	"GoCrawler/engine"
	"GoCrawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`)
var nameRe = regexp.MustCompile(`class="nickName" [^>]*>([^<]+)</h1>`)
var genderRe = regexp.MustCompile(`"genderString":"(.[^"]+)"`)

func ParseProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extractStringSlice(contents, marriageRe, 0)

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Xingzuo = extractStringSlice(contents, xinzuoRe, 2)
	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Name = extractString(contents, nameRe)
	//profile.Name = name

	profile.Car = extractStringSlice(contents, carRe, 4)
	profile.Occupation = extractStringSlice(contents, occupationRe, 6)
	profile.Education = extractStringSlice(contents, educationRe, 7)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractStringSlice(contents, houseRe, 3)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)

	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractStringSlice(contents []byte, re *regexp.Regexp, index int) string {
	submatch := re.FindAllSubmatch(contents, -1)
	if len(submatch) >= 2 && len(submatch) > index {
		return string(submatch[index][1])
	} else {
		return ""
	}
}
