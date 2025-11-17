package controller

import (
	// Add necessary imports here
	"goadmin/internal/admin/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	// Add necessary fields here
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (ctl *ProductController) List(c *gin.Context) {

}
