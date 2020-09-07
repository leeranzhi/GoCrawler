package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var Headers = map[string]string{
	"user-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.19041",
	"Referer":    "http://album.zhenai.com",
	`cookie`:     `sid=463adf61-fa3a-46fb-a8d0-081554edc55a; FSSBBIl1UgzbN7NO=5ZhXgxf9igBFg0ZnYbBxPCRvN_a2NYdZb5tLfiSibgIol3SxqufmUf9R94hzrXVLoBcszsU4sT3LLAQaFmpIa5a; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1599062036; ec=sZPovhYL-1599322704680-d37388117fe52-2019756260; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1599497776; FSSBBIl1UgzbN7NP=5UTRbh2eKRPZqqqm0FUZNZaFyQ1a.FfRs1w86NdBSVHthj7TevhfVdBEyGYq6UaWaTD7bFPB4S7dsKJbjV9i1fL41gPVGn71XfqZEe_K8ApiRkPf8eBHINmJx1I7CzRf75lpvyfbwJp75g3I7lf.zx7Iov1Ts4665dwX9agvOy7SHFMOkLeZjDdRasiroWdaUmhm8tjs8dWekl1jH.31YZ7s.T5i0ptDca0FAPD9aipVbDXaMmXsvPwqGxhekGIi6A; _efmdata=XWaT53vM6fC0x%2BqZELGVjPYmrqwwLCIFiFSJ3sXbqRfuWCokBZcupqEm1pNVWQ64kuy%2F9ZhaLyA4uRl7flopbIT3YEMERc8XyCSBRLtvDCI%3D; _exid=9N4xVm%2FggNFUuzmJObPYV0jAyppty7qEvk4hJCx6ml0uzUU4O1T%2FfaG8rKhEQ2%2B%2BUfwVoPpU34rp4me56yUzCQ%3D%3D`,
}

func Fetch(url string) ([]byte, error) {
	client := http.Client{}

	newUrl := strings.Replace(url, "http://", "https://", 1)

	request, err := http.NewRequest(http.MethodGet, newUrl, nil)
	if err != nil {
		return nil, err
	}

	for index, header := range Headers {
		request.Header.Set(index, header)
	}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrongs status code: %d", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

/**
通过1024个字节去猜Encoding编码
*/
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8
	}
	determineEncoding, _, _ := charset.DetermineEncoding(bytes, "")
	return determineEncoding
}
