package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	username := "jhj"
	password := "123456"
	token := strconv.FormatInt(time.Now().Unix(), 10)

	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		if c.Query("username") == username && c.Query("password") == password {
			c.JSON(200, gin.H{
				"token": token,
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid credentials",
			})
		}
	})

	name := "jhj"

	r.GET("/getUser", func(c *gin.Context) {
		if c.Query("token") == token {
			c.JSON(200, gin.H{
				"name": name,
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})
		}
	})

	r.GET("/logout", func(c *gin.Context) {
		if c.Query("token") == token {
			token = strconv.FormatInt(time.Now().Unix(), 10)
			c.JSON(200, gin.H{
				"message": "Logout successful",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})
		}
	
		type Login struct {
			UserName string `form:"username"json:"username"binding:"required"`
			Password string `form:"password"json:"password"binding:"required"`
		}
		
func main() {
			r := gin.Default()
			r.POST("/loginJSON", func(c *gin.Context) {
				var login Login
				if err := c.ShouldBind(&login); err == nil {
					c.JSON(http.StatusOK, gin.H{
						"message":  "login with json query",
						"username": login.UserName,
						"Password": login.Password,
					})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{
						"err_no": err.Error(),
					})
				}
			})
			r.POST("/loginForm", func(c *gin.Context) {
				var login Login
				if err := c.ShouldBind(&login); err == nil {
					c.JSON(http.StatusOK, gin.H{
						"message":  "login with form query",
						"username": login.UserName,
						"Password": login.Password,
					})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
				}
		
			})
			r.GET("/loginQuery", func(c *gin.Context) {
				var login Login
				if err := c.ShouldBind(&login); err == nil {
					c.JSON(http.StatusOK, gin.H{
						"message":  "login with querystring query",
						"username": login.UserName,
						"Password": login.Password,
					})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
				}
			})
			r.POST("/form", func(context *gin.Context) {
				types := context.DefaultPostForm("type", "post")
				username := context.PostForm("username")
				password := context.PostForm("password")
		
				context.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,types:%s", username, password, types))
			})})
	r.Run(":8000")
}
