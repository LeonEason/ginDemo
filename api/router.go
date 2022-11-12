package api

import (
	middleware "ginDemo/api/middware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/change", change)
	r.POST("/find", find)
	r.Run(":8088")
}
