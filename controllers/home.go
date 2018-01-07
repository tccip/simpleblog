package controllers

import "github.com/gin-gonic/gin"
import "net/http"
import model "tccip/test_templete/models"

func Home(c *gin.Context) {
	topics := model.GetAllTopic()
	c.HTML(http.StatusOK, "home.html", gin.H{"Topics": topics})
}
