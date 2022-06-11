package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/controller"
	"golang-rest-api/exception"
)

func NewRouter(categoryController controller.CategoryController, customerController controller.CustomerController, order_productController controller.OrderProductController, ordersController controller.OrdersController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	router.GET("/api/order_products", order_productController.FindAll)
	router.GET("/api/order_products/:order_productId", order_productController.FindById)
	router.POST("/api/order_products", order_productController.Create)
	router.PUT("/api/order_products/:order_productId", order_productController.Update)
	router.DELETE("/api/order_products/:order_productId", order_productController.Delete)

	router.GET("/api/orderss", ordersController.FindAll)
	router.GET("/api/orderss/:ordersId", ordersController.FindById)
	router.POST("/api/orderss", ordersController.Create)
	router.PUT("/api/orderss/:ordersId", ordersController.Update)
	router.DELETE("/api/orderss/:ordersId", ordersController.Delete)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
