package main

import (
	"github.com/gin-gonic/gin"
	"hello/conf"
	"net/http"
)

func main(){
	r := conf.Start()
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
