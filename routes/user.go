package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r gin.IRoutes) {
	r.POST("users", userCreate)
}

func userCreate(c *gin.Context) {
	fmt.Println("test1")
}
