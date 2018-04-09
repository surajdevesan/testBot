package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func addToSpace(space SpaceStruct) string {
	return "Thank you for adding me"
}

func getTokenHeader(token string) string {
	return "token " + token
}

func getRepoName(repoInfo []GitHubRepoInfo) []string {
	repoNames := make([]string, len(repoInfo))
	for i, v := range repoInfo {
		repoNames[i] = v.Name
	}
	return repoNames
}

func messageFromUser(message MessageFormat) string {
	repoInfo := make([]GitHubRepoInfo, 0)
	if message.Text == "List all repos" {
		urlData := RequestFormat{
			url:     "https://api.github.com/users/surajdevesan/repos",
			urlType: "GET",
			headerData: Header{
				name:   "Authentication",
				header: getTokenHeader("Need to replace"),
			},
		}
		res, err := GetRequest(urlData)
		if res.StatusCode != 200 {
			fmt.Println("Error in request")
			return "Error in retrieving"
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("The response is res", res)
		err = json.Unmarshal([]byte(body), &repoInfo)
		if err != nil {
			fmt.Println("Error at unmarshalling", err)
		}
	}
	repoNames := getRepoName(repoInfo)
	names := strings.Join(repoNames, "\n")
	fmt.Println("Sucess", names)
	return names
}
