package main

import (
  // "net/http"
  "github.com/gin-gonic/gin"
  // "fmt"
  "main/app/controllers"
)

// import . "main/app/controllers"

// http://localhost:8080/ping

func main() {
  r := gin.Default()
  // Group our routes:
  apiGroup := r.Group("/api")

  foodGroup := apiGroup.Group("/food")

  foodGroup.GET("/:foodid",  func(c *gin.Context) {
    controllers.GetFoods(c)
  })

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}