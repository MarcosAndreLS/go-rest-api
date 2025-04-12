package main

import (
	"github.com/MarcosAndreLS/go-rest-api/controller"
	"github.com/MarcosAndreLS/go-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

func main(){
	
	server := gin.Default()
	ProductUseCase := usecase.NewProductUseCase()
	//camada de controllers
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8000")
}