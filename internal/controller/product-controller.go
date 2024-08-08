package controller

import (
	"hexagonal-archteture/internal/model"
	"hexagonal-archteture/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	// Usecase
	ProductUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) *productController {
	return &productController{
		ProductUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.ProductUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{
			Message: "invalid id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Product id must be an number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUseCase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{
			Message: "invalid id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Product id must be an number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.ProductUseCase.DeleteProduct(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	response := model.Response{
		Message: "Product deleted successfully",
	}

	ctx.JSON(http.StatusOK, response)
}
