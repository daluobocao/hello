package conf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start() *gin.Engine{
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.LoadHTMLGlob("views/*")
	return r
}
