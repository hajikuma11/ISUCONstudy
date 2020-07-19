package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")

	engine.GET("/", func(cntxt *gin.Context) {
		cntxt.HTML(200, "index.html", gin.H{})
	})

	engine.Run()
}
