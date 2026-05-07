package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"

	_ "go-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Product API
// @version         1.0
// @description     API REST para gerenciamento de produtos.
// @host            localhost:8000
// @BasePath        /

func main() {
	server := gin.Default()

	dbConnetion, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnetion)
	//Camada de usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//Camada de controller
	ProductController := controller.NewProductControler(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messege": "pong",
		})
	})

	server.GET("/products", ProductController.Getproducts)
	server.POST("/product", ProductController.CreatProduct)
	server.GET("/product/:productId", ProductController.GetproductsById)
	server.PUT("/product/:productId", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")

}
