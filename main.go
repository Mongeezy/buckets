package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")
  router.Static("/static", "./static")
  router.GET("/", func(context *gin.Context) {
    context.HTML(http.StatusOK, "index.html", gin.H{
      "title": "Buckets on buckets",
    })
  })
  router.Run(":8080")
}
