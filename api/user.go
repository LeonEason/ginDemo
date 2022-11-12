package api

import (
	"fmt"
	"ginDemo/dao"
	"ginDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "verification failed",
		})
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	fmt.Println(flag)
	if flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	}
	dao.AddUser(username, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successfully",
	})
}

func change(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "verification failed",
		})
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exist",
		})
		return
	}
	dao.AddUser(username, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "change password successfully",
	})
}
func find(c *gin.Context) {
	username := c.PostForm("username")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  200,
			"message": "user doesn't exist",
		})
		return
	}
	rightPassword := dao.GetPassword(username)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": rightPassword,
	})
}
func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "verification failed",
		})
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exist",
		})
		return
	}
	rightPassword := dao.GetPassword(username)
	if password != rightPassword {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successfully",
	})
}
