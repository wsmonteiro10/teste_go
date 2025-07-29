package main

import (
	"teste_go/handlers"
	"teste_go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "teste_go/docs"
)

// @title Desafio Técnico | BU Sales & Marketing
// @version 1.0
// @description Desafio Técnico | BU Sales & Marketing.
// @securityDefinitions.apikey ClientSecret
// @in header
// @name Client-Secret
func main() {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Client-Secret")

	r.Use(cors.New(config))

	r.GET("/ping", handlers.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/rapido", middleware.ClientSecretMiddleware(), handlers.UploadFile)
	r.POST("/ultrarapido", middleware.ClientSecretMiddleware(), handlers.UltraUploadFile)
	r.GET("/listas_arquivos", middleware.ClientSecretMiddleware(), handlers.ListUploadFiles)

	r.Run(":8080")
}
