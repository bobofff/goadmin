package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goadmin/internal/admin/dto"
	"goadmin/internal/admin/service"
	"goadmin/pkg/binding"
	"goadmin/pkg/response"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (ctl *ProductController) List(c *gin.Context) {
	var request dto.ProductListRequest

	if err := c.ShouldBind(&request); err != nil {
		response.Error(c, http.StatusBadRequest, 4001, binding.ParseError(err, request, c))
		return
	}

	result, err := ctl.productService.GetList(c.Request.Context(), request)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 5001, err.Error())
		return
	}

	response.Success(c, result)
}
