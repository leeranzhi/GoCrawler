package engine

import (
	"GoCrawler/fetcher"
	"log"
)

func Run(seed ...Request) {
	var requests []Request

	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)

		//获取url的请求结果
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher:error fetching url %s: %v", r.Url, err)
			continue
		}

		//解析结果body
		paresResult := r.ParserFunc(body)
		//加入新的url
		requests = append(requests, paresResult.Requests...)

		for _, item := range paresResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
