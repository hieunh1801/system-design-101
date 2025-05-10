package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Print("STARTING_ON_PORT ", port)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.String(http.StatusOK, fmt.Sprintf("pong from %s", hostname))
	})

	r.Run(":" + port)
}
