package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", home)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome Home",
	})
}
