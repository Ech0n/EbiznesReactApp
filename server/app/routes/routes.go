package routes

import (
	"project/controllers"

	"github.com/labstack/echo/v4"
)

const productsPath = "/products"
const singleProductPath = "/products/:id"
const categoriesPath = "/categories"
const paymentsPath = "/payments"

func SetupRoutes(e *echo.Echo) {
	e.POST(productsPath, controllers.CreateProduct)
	e.GET(productsPath, controllers.GetProducts)
	e.GET(singleProductPath, controllers.GetProduct)
	e.PUT(singleProductPath, controllers.UpdateProduct)
	e.DELETE(singleProductPath, controllers.DeleteProduct)

	e.GET(paymentsPath, controllers.GetCategories)
	e.POST(paymentsPath, controllers.CreateCategory)

	e.POST(categoriesPath, controllers.ConfirmOrder)
}
