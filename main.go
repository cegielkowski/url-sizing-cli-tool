package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

// Response represents the response from visiting a URL
type Response struct {
	URL     string
	BodyLen int64
}

// ByBodyLen is a type used to sort a slice of Responses by BodyLen
type ByBodyLen []Response

func (r ByBodyLen) Len() int {
	return len(r)
}

func (r ByBodyLen) Less(i, j int) bool {
	return r[i].BodyLen < r[j].BodyLen
}

func (r ByBodyLen) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// visitURL visits the given URL and returns its response body size
func visitURL(url string) (Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return Response{}, err
	}

	return Response{URL: url, BodyLen: int64(buf.Len())}, nil
}

// processURLs visits a list of URLs and returns a sorted list of responses
func processURLs(urls []string) ([]Response, error) {
	responses := make([]Response, len(urls))
	for i, url := range urls {
		resp, err := visitURL(url)
		if err != nil {
			return nil, err
		}
		responses[i] = resp
	}

	sort.Sort(ByBodyLen(responses))
	return responses, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the list of URLs (comma-separated): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		log.Fatalln("No URLs provided.")
	}

	urls := strings.Split(input, ",")
	responses, err := processURLs(urls)
	if err != nil {
		log.Fatalln(err)
	}
	for _, resp := range responses {
		log.Printf("%s: %d\n", resp.URL, resp.BodyLen)
	}
}
