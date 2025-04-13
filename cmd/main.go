package main

import (
	"github.com/MarcosAndreLS/go-rest-api/controller"
	"github.com/MarcosAndreLS/go-rest-api/db"
	"github.com/MarcosAndreLS/go-rest-api/repository"
	"github.com/MarcosAndreLS/go-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

func main(){
	
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if (err != nil){
		panic(err)
	}
	//camda de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada de usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//camada de controllers
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)

	server.Run(":8000")
}