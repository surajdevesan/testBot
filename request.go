package main

import (
	"fmt"
	"net/http"
)

type Header struct {
	name   string
	header string
}
type RequestFormat struct {
	url        string
	urlType    string
	headerData Header
}

func GetRequest(urlData RequestFormat) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(urlData.urlType, urlData.url, nil)
	req.Header.Add(urlData.headerData.name, urlData.headerData.header)
	fmt.Println("CallingGithub repo with url", urlData.url, urlData.urlType)
	res, err := client.Do(req)
	return res, err
}
