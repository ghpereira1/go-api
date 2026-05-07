package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductControler(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) Getproducts(ctx *gin.Context) {

	products, err := p.productUsecase.Getproducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary      Criar um novo produto
// @Description  Recebe um JSON e cria um produto no banco de dados
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Dados do Produto"
// @Success      200      {object}  model.Product
// @Failure      400      {object}  map[string]string
// @Router       /product [post]
func (p *productController) CreatProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: o nome é obrigatório e o preço deve ser maior que zero",
		})
		return
	}

	insertedProduct, err := p.productUsecase.CreatProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetproductsById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Responde{
			Messege: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Responde{
			Messege: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetproductsById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Responde{
			Messege: "Produto nao encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary      Alterar um produto
// @Description  Recebe um JSON e altera o produto correspondente no banco de dados
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Dados do Produto"
// @Success      200      {object}  model.Product
// @Failure      400      {object}  map[string]string
// @Router       /product [put]
func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	idInt, err := strconv.Atoi(id) // Converte o ID da URL para número
	if err != nil {
		response := model.Responde{
			Messege: "Produto nao encontrado",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.productUsecase.UpdateProduct(idInt, product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary      Deletar um produto
// @Description  Recebe um ID e deleta o produto correspondente no banco de dados
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Dados do Produto"
// @Success      200      {object}  model.Product
// @Failure      400      {object}  map[string]string
// @Router       /product [delete]
func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response := model.Responde{
			Messege: "Produto nao encontrado",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUsecase.DeleteProduct(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produto deletado"})
}
