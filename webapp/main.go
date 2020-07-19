package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")

	data := "Go."

	engine.GET("/", func(cntxt *gin.Context) {
		cntxt.HTML(200, "index.html", gin.H{"data": data})
	})

	engine.Run()
}
