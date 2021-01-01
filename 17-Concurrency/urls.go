package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://golang.org",
	"https://github.com",
	"https://123abc123.com",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func main() {
	resp := httpGet(urls)
	result(resp)
}

func httpGet(urls []string) []*HttpResponse {

	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}

	client := http.Client{}

	for _, url := range urls {
		go func(url1 string) {

			fmt.Println("Fetching Url", url1)
			resp, err := client.Get(url1)
			ch <- &HttpResponse{
				url:      url1,
				response: resp,
				err:      err,
			}

		}(url)
	}

	for {

		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			if r.err != nil {
				fmt.Println("with an error", r.err)
			}
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}

		case <-time.After(10 * time.Millisecond):
			fmt.Printf(".")
		}

	}

}

func result(results []*HttpResponse) {
	for _, res := range results {
		if res != nil && res.response != nil {
			fmt.Println(res.response.Status)
		} else {
			fmt.Println(res.url, "not found")
			fmt.Println(res)
		}



	}
}
