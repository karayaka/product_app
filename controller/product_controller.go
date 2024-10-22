package controller

import (
	"net/http"
	"product_app/service"
	requestmodels "product_app/service/model/request_models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ps service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		ps: productService,
	}
}

func (pc *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/product/:id", pc.GetProductById)
	e.GET("/api/product", pc.GetAllProducts)
	e.POST("/api/product", pc.AddProduct)
}

func (pc *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)
	product, err := pc.ps.GetById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, pc.ps.GetAllProducts())
}

func (pc *ProductController) AddProduct(c echo.Context) error {
	var addProcductRequest requestmodels.ProductRequestModel
	err := c.Bind(&addProcductRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bir Hata Oluştu")
	}
	pc.ps.Add(addProcductRequest.ToModel())
	return c.JSON(http.StatusOK, "Product Kayıt edildi")
}
