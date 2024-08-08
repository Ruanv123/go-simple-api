package main

import (
	"hexagonal-archteture/internal/controller"
	"hexagonal-archteture/internal/infra/db"
	"hexagonal-archteture/internal/repository"
	"hexagonal-archteture/internal/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()

	// camada de infra - banco
	dbConnection, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	// camada de repositorio
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada de usecases
	ProductUseCase := usecase.NewProductUseCase(*ProductRepository)

	// camada de controllers
	ProductController := controller.NewProductController(*ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &gin.H{"message": "pong"})
	})

	server.GET("/product", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetProductById)
	server.DELETE("/product/:id", ProductController.DeleteProduct)

	server.Run(":8000")
}
