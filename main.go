package main

import (
	"github.com/gin-gonic/gin"
	"hello/conf"
	"net/http"
)

func main(){
	r := conf.Start()
	r.GET("/index", index)
	r.GET("/about", about)
	r.GET("/mycourse", mycourse)
	r.Run(":8002")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func about(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}

func mycourse(c *gin.Context) {
	c.HTML(http.StatusOK, "mycourse.html", nil)
}
