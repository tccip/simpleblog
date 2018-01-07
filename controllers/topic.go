package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	model "github.com/tccip/simpleblog/models"

	"github.com/gin-gonic/gin"
)

func Topic(c *gin.Context) {
	topics := model.GetAllTopic()
	fmt.Println(topics)
	c.HTML(http.StatusOK, "topic.html", gin.H{
		"Topics": topics,
	})
}

func TopicAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "topic_add.html", nil)
}

func TopicAddIn(c *gin.Context) {
	title, _ := c.GetPostForm("title")
	content, _ := c.GetPostForm("content")
	fmt.Println(title)
	fmt.Println(content)
	var err bool
	err = model.AddTopic(title, content)
	if err == false {
		fmt.Println("Add fail")
	}
	c.Redirect(302, "/topic")
}

func TopicDel(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("bad convert")
		return
	}
	model.DelTopic(id)
	c.Redirect(302, "/topic")
}

func TopicView(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("bad convert")
		return
	}
	topic := model.QueryTopic(id)
	c.HTML(http.StatusOK, "topic_view.html", gin.H{
		"Topic": topic,
	})
}

func TopicModify(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("bad convert")
		return
	}
	topic := model.QueryTopic(id)
	c.HTML(http.StatusOK, "topic_modify.html", gin.H{
		"Topic": topic,
	})
}

func TopicModifyIn(c *gin.Context) {
	postid, _ := c.GetPostForm("id")
	title, _ := c.GetPostForm("title")
	content, _ := c.GetPostForm("content")

	id, _ := strconv.Atoi(postid)
	model.ModifyTopic(id, title, content)
	c.Redirect(302, "/topic")
}
