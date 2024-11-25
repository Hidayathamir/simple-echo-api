package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/api/echo/:msg", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": c.Param("msg")})
	})
	err := ginEngine.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
