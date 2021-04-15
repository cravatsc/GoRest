package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.Run()
}

type Prod struct {
	gorm.Model
}
