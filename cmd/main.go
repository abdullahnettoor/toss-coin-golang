package main

import (
	"github.com/abdullahnettoor/toss-coin-golang/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("./internal/view/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	r.GET("/toss-coin", handlers.TossCoin)

	r.Run(":3333")
}
