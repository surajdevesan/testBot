package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SpaceStruct struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Type        string `json:"type"`
}

type Request struct {
	Type    string      `json:"type"`
	Token   string      `json:"token"`
	Space   SpaceStruct `json:"space"`
	Message struct {
		Name   string `json:"name"`
		Sender struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			AvatarURL   string `json:"avatarUrl"`
			Email       string `json:"email"`
		} `json:"sender"`
		Text   string `json:"text"`
		Thread struct {
			Name string `json:"name"`
		} `json:"thread"`
	} `json:"message"`
}

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		data := &Request{}
		response := "Test Default event"
		c.Bind(data)
		// x,_ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(data)
		switch eventType := data.Type; eventType {
		case "ADDED_TO_SPACE":
			response = addToSpace(data.Space)
		case "MESSAGE":
			response = messageFromUser()
		}
		// fmt.Printf("%s", string(x))
		c.JSON(200, gin.H{
			"text": response,
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test",
		})
	})
	r.Run()
}
