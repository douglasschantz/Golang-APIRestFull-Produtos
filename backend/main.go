package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schantz/web/go-api-produtos/backend/database"
	"github.com/schantz/web/go-api-produtos/backend/models/product"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	database.InitDB()
	dtbConn := database.ReturnDB()

	var (
		productRepository = product.NewRepository(dtbConn)
		/*employeeRepository = employee.NewRepository(databaseConnection)
		customerRepository = customer.NewRepository(databaseConnection)
		orderRepository    = order.NewRepository(databaseConnection)*/
	)

	var (
		productService product.Service
		//employeeService employee.Service
		//customerService customer.Service
		//orderService    order.Service
	)

	productService = product.NewService(productRepository)
	/*employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)*/

	r := chi.NewRouter()

	r.Mount("/products", product.MakeHttpHandler(productService))
	/*r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHTTPHandler(customerService))
	r.Mount("/orders", order.MakeHTTPHandler(orderService))*/

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))

	http.ListenAndServe(":3000", r)
}
