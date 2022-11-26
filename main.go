package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// http://localhost:8080/ping

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "BYE mate",
    })
  })

  r.GET("/test", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "HELLO mate",
    })
  })

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}