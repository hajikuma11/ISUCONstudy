package main

import (
	"./db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")

	db.DbInit()

	engine.GET("/", func(contxt *gin.Context) {
		todos := db.DbGetAll()
		contxt.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	engine.POST("/new", func(contxt *gin.Context) {
		text := contxt.PostForm("text")
		status := contxt.PostForm("status")
		db.DbInsert(text, status)
		contxt.Redirect(302, "/")
	})

	engine.GET("/detail/:id", func(contxt *gin.Context) {
		n := contxt.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := db.DbGetOne(id)
		contxt.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	engine.POST("/update/:id", func(contxt *gin.Context) {
		n := contxt.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := contxt.PostForm("text")
		status := contxt.PostForm("status")
		db.DbUpdate(id, text, status)
		contxt.Redirect(302, "/")
	})

	//削除確認
	engine.GET("/delete_check/:id", func(contxt *gin.Context) {
		n := contxt.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := db.DbGetOne(id)
		contxt.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	engine.POST("/delete/:id", func(contxt *gin.Context) {
		n := contxt.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		db.DbDelete(id)
		contxt.Redirect(302, "/")

	})

	engine.Run()

}
