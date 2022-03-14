package main

import (
	"io"
	"net/http"
	"os"

	"github.com/InfoTCube/Seneca-the-Younger-API/src/controllers"
	"github.com/InfoTCube/Seneca-the-Younger-API/src/middlewares"
	"github.com/InfoTCube/Seneca-the-Younger-API/src/services"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	quoteService    services.QuoteService       = services.New()
	quoteController controllers.QuoteController = controllers.New(quoteService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/quotes", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, quoteController.GetAll())
	})

	server.POST("/quotes", func(ctx *gin.Context) {
		err := quoteController.Add(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Quote added sucessfully!"})
		}
	})

	server.Run(":8080")
}
