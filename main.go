package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("https://httpbin.org/post", url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Response: %v", string(body))
}
