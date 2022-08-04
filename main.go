package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var PTransport = &http.Transport{Proxy: nil}

	client := http.Client{Transport: PTransport}

	req, _ := http.NewRequest("GET", "http://169.254.169.254/metadata/instance", nil)
	req.Header.Add("Metadata", "True")

	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", "2021-02-01")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR> %v", err)
		log.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
}
