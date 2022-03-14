package main

import (
	"github.com/gin-gonic/gin"
)

var (
	quoteService    service.QuoteService       = service.New()
	quoteController controller.QuoteController = controller.New(quoteService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery())

	server.GET("/quotes", func(ctx *gin.Context) {
		ctx.JSON(200, quoteController.GetAll())
	})

	server.POST("/quotes", func(ctx *gin.Context) {
		ctx.JSON(200, quoteController.Add(ctx))
	})

	server.Run(":8080")
}
