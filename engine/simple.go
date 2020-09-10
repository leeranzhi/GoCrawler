package engine

import (
	"GoCrawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seed ...Request) {
	var requests []Request

	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		paresResult, err := worker(r)
		if err != nil {
			continue
		}

		//加入新的url
		requests = append(requests, paresResult.Requests...)

		for _, item := range paresResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)

	//获取url的请求结果
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	//解析结果body
	return r.ParserFunc(body), nil
}
