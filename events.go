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

func addToken(message MessageFormat) {
	token := "Need to be replaced"
	email := message.Sender.Email
	var userInfo = UserInfo{Token: token, Email: email}
	fmt.Println("Declared token and email")
	sess, err := connectToDb()
	defer sess.Close()
	db := setDB("test", sess)
	if err == nil {
		fmt.Println("Reached without error in declaring the database")
		err := insertToCollection(userInfo, db, "test")
		if err != nil {
			fmt.Println("Error in insertion", err)
			return
		}
		fmt.Println("Succesfully inserted to DB")
	} else {
		fmt.Println("Error in connection", err)
	}
}

func messageFromUser(message MessageFormat) string {
	repoInfo := make([]GitHubRepoInfo, 0)
	if message.Text == "List all repos" {
		sess, err := connectToDb()
		db := setDB("TestDB", sess)
		user, err := findFromCollection(message.Sender.Email, db.C("test"))
		if user.Email == "" || err != nil {
			fmt.Println("Error in acces token", user)
			return "Access token issue"
		}
		urlData := RequestFormat{
			url:     "https://api.github.com/users/surajdevesan/repos",
			urlType: "GET",
			headerData: Header{
				name:   "Authentication",
				header: getTokenHeader(user.Token),
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
		repoNames := getRepoName(repoInfo)
		names := strings.Join(repoNames, "\n")
		fmt.Println("Sucess", names)
		return names
	} else if message.Text == "Add token" {
		addToken(message)
	}
	return "Default"
}
