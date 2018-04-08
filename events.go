package main

import (
	"fmt"
)

func addToSpace(space SpaceStruct) string {
	return "Thank you for adding me"
}

func getTokenHeader(token string) string {
	return "token " + token
}

func messageFromUser(message MessageFormat) string {
	if message.Text == "List all repos" {
		urlData := RequestFormat{
			url:     "https://api.github.com/surajdevesan/repos",
			urlType: "GET",
			headerData: Header{
				name:   "Authentication",
				header: getTokenHeader("Have to be replaced"),
			},
		}
		res, err := GetRequest(urlData)
	}
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Sucess", res)
	}
	return "I am under development and these features are not yet developed. You can track the progress <https://github.com/surajdevesan/testBot|here>	"
}
