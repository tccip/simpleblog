package controllers

import "github.com/gin-gonic/gin"
import "net/http"

func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	autoLogin := c.PostForm("autoLogin")
	if username == "admin" && password == "111" {
		maxAge := 0
		if autoLogin == "on" {
			maxAge = 1<<31 - 1
		}
		user_cookie := &http.Cookie{
			Name:     "username",
			Value:    username,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   maxAge,
		}
		password_cookie := &http.Cookie{
			Name:     "password",
			Value:    password,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   maxAge,
		}
		http.SetCookie(c.Writer, user_cookie)
		http.SetCookie(c.Writer, password_cookie)
		c.Redirect(301, "/")
		return
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

func checkAccount(c *gin.Context) bool {
	ck, err := c.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value
	ck, err = c.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value
	if username == "admin" && password == "111" {
		return true
	}
	return false
}
